package gateway

import (
	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	core_xds "github.com/kumahq/kuma/pkg/core/xds"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	envoy_listeners "github.com/kumahq/kuma/pkg/xds/envoy/listeners"
	envoy_names "github.com/kumahq/kuma/pkg/xds/envoy/names"
)

// TODO(jpeach) It's a lot to ask operators to tune these defaults,
// and we probably would never do that. However, it would be convenient
// to be able to update them for performance testing and benchmarking,
// so at some point we should consider making these settings available,
// perhaps on the Gateway or on the Dataplane.

// Buffer defaults.
const DefaultConnectionBuffer = 32 * 1024

// ListenerGenerator generates Kuma gateway listeners.
type ListenerGenerator struct{}

func (*ListenerGenerator) SupportsProtocol(p mesh_proto.Gateway_Listener_Protocol) bool {
	switch p {
	case mesh_proto.Gateway_Listener_HTTP,
		mesh_proto.Gateway_Listener_HTTPS:
		return true
	default:
		return false
	}
}

func (*ListenerGenerator) GenerateHost(ctx xds_context.Context, info *GatewayResourceInfo) (*core_xds.ResourceSet, error) {
	// TODO(jpeach) what we really need to do here is build the
	// listener once, then generate a HTTP filter chain for each
	// host on the same HTTPConnectionManager. Each HTTP filter
	// chain should be wrapped in a matcher that selects it for
	// only the host's domain name. This will give us consistent
	// per-host HTTP filter chains for both HTTP and HTTPS
	// listeners.
	//
	// https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/matching/matching_api
	if info.Resources.Listener != nil {
		return nil, nil
	}

	// A new listener gets a new filter chain.
	info.Resources.FilterChain = nil

	port := info.Listener.Port
	protocol := info.Listener.Protocol
	address := info.Dataplane.Spec.GetNetworking().Address

	log.V(1).Info("generating listener",
		"address", address,
		"port", port,
		"protocol", protocol,
	)

	info.Resources.Listener = envoy_listeners.NewListenerBuilder(info.Proxy.APIVersion).
		Configure(
			envoy_listeners.InboundListener(
				envoy_names.GetGatewayListenerName(info.Gateway.Meta.GetName(), protocol.String(), port),
				address, port, core_xds.SocketAddressProtocolTCP),
			// Limit default buffering for edge connections.
			envoy_listeners.ConnectionBufferLimit(DefaultConnectionBuffer),
			// Roughly balance incoming connections.
			envoy_listeners.EnableReusePort(true),
			// Always sniff for TLS.
			envoy_listeners.TLSInspector(),
		)

	// TODO(jpeach) if proxy protocol is enabled, add the proxy protocol listener filter.

	return nil, nil
}
