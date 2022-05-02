// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: mesh/options.proto

package mesh

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KumaResourceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the Kuma resource struct.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name and value of the modelResourceType constant.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// True if this resource has global scope. Otherwise it will be mesh scope.
	Global bool `protobuf:"varint,3,opt,name=global,proto3" json:"global,omitempty"`
	// Name of the resource's Go package.
	Package string `protobuf:"bytes,4,opt,name=package,proto3" json:"package,omitempty"`
	// Whether to generate a no-op validation function.
	SkipValidation bool `protobuf:"varint,5,opt,name=skip_validation,json=skipValidation,proto3" json:"skip_validation,omitempty"`
	// Whether to skip type registration for this resource.
	SkipRegistration bool            `protobuf:"varint,6,opt,name=skip_registration,json=skipRegistration,proto3" json:"skip_registration,omitempty"`
	Kds              *KumaKdsOptions `protobuf:"bytes,10,opt,name=kds,proto3" json:"kds,omitempty"`
	Ws               *KumaWsOptions  `protobuf:"bytes,7,opt,name=ws,proto3" json:"ws,omitempty"`
	// Whether scope is "Namespace"; Otherwise to "Cluster".
	ScopeNamespace bool `protobuf:"varint,11,opt,name=scope_namespace,json=scopeNamespace,proto3" json:"scope_namespace,omitempty"`
	// Whether to skip generation of native API helper functions.
	SkipKubernetesWrappers bool `protobuf:"varint,12,opt,name=skip_kubernetes_wrappers,json=skipKubernetesWrappers,proto3" json:"skip_kubernetes_wrappers,omitempty"`
	// Whether to generate Inspect API endpoint
	AllowToInspect bool `protobuf:"varint,13,opt,name=allow_to_inspect,json=allowToInspect,proto3" json:"allow_to_inspect,omitempty"`
	// If resource has more than one version, then the flag defines which version
	// is used in the storage. All other versions must be convertible to it.
	StorageVersion bool `protobuf:"varint,14,opt,name=storage_version,json=storageVersion,proto3" json:"storage_version,omitempty"`
}

func (x *KumaResourceOptions) Reset() {
	*x = KumaResourceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KumaResourceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KumaResourceOptions) ProtoMessage() {}

func (x *KumaResourceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KumaResourceOptions.ProtoReflect.Descriptor instead.
func (*KumaResourceOptions) Descriptor() ([]byte, []int) {
	return file_mesh_options_proto_rawDescGZIP(), []int{0}
}

func (x *KumaResourceOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KumaResourceOptions) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *KumaResourceOptions) GetGlobal() bool {
	if x != nil {
		return x.Global
	}
	return false
}

func (x *KumaResourceOptions) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *KumaResourceOptions) GetSkipValidation() bool {
	if x != nil {
		return x.SkipValidation
	}
	return false
}

func (x *KumaResourceOptions) GetSkipRegistration() bool {
	if x != nil {
		return x.SkipRegistration
	}
	return false
}

func (x *KumaResourceOptions) GetKds() *KumaKdsOptions {
	if x != nil {
		return x.Kds
	}
	return nil
}

func (x *KumaResourceOptions) GetWs() *KumaWsOptions {
	if x != nil {
		return x.Ws
	}
	return nil
}

func (x *KumaResourceOptions) GetScopeNamespace() bool {
	if x != nil {
		return x.ScopeNamespace
	}
	return false
}

func (x *KumaResourceOptions) GetSkipKubernetesWrappers() bool {
	if x != nil {
		return x.SkipKubernetesWrappers
	}
	return false
}

func (x *KumaResourceOptions) GetAllowToInspect() bool {
	if x != nil {
		return x.AllowToInspect
	}
	return false
}

func (x *KumaResourceOptions) GetStorageVersion() bool {
	if x != nil {
		return x.StorageVersion
	}
	return false
}

type KumaWsOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the policy for resource name usage in path.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Plural is only to be set if the plural of the resource is irregular (not
	// just adding a 's' at the end).
	Plural string `protobuf:"bytes,2,opt,name=plural,proto3" json:"plural,omitempty"`
	// ReadOnly if the resource is read only.
	ReadOnly bool `protobuf:"varint,3,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	// AdminOnly whether this entity requires admin auth to access these
	// endpoints.
	AdminOnly bool `protobuf:"varint,4,opt,name=admin_only,json=adminOnly,proto3" json:"admin_only,omitempty"`
}

func (x *KumaWsOptions) Reset() {
	*x = KumaWsOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KumaWsOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KumaWsOptions) ProtoMessage() {}

func (x *KumaWsOptions) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KumaWsOptions.ProtoReflect.Descriptor instead.
func (*KumaWsOptions) Descriptor() ([]byte, []int) {
	return file_mesh_options_proto_rawDescGZIP(), []int{1}
}

func (x *KumaWsOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KumaWsOptions) GetPlural() string {
	if x != nil {
		return x.Plural
	}
	return ""
}

func (x *KumaWsOptions) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

func (x *KumaWsOptions) GetAdminOnly() bool {
	if x != nil {
		return x.AdminOnly
	}
	return false
}

type KumaKdsOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SendToGlobal whether this entity will be sent from zone cp to global cp
	SendToGlobal bool `protobuf:"varint,1,opt,name=send_to_global,json=sendToGlobal,proto3" json:"send_to_global,omitempty"`
	// SendToZone whether this entity will be sent from global cp to zone cp
	SendToZone bool `protobuf:"varint,2,opt,name=send_to_zone,json=sendToZone,proto3" json:"send_to_zone,omitempty"`
}

func (x *KumaKdsOptions) Reset() {
	*x = KumaKdsOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KumaKdsOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KumaKdsOptions) ProtoMessage() {}

func (x *KumaKdsOptions) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KumaKdsOptions.ProtoReflect.Descriptor instead.
func (*KumaKdsOptions) Descriptor() ([]byte, []int) {
	return file_mesh_options_proto_rawDescGZIP(), []int{2}
}

func (x *KumaKdsOptions) GetSendToGlobal() bool {
	if x != nil {
		return x.SendToGlobal
	}
	return false
}

func (x *KumaKdsOptions) GetSendToZone() bool {
	if x != nil {
		return x.SendToZone
	}
	return false
}

var file_mesh_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*KumaResourceOptions)(nil),
		Field:         43534533,
		Name:          "kuma.mesh.resource",
		Tag:           "bytes,43534533,opt,name=resource",
		Filename:      "mesh/options.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional kuma.mesh.KumaResourceOptions resource = 43534533;
	E_Resource = &file_mesh_options_proto_extTypes[0] // 'kuma'
)

var File_mesh_options_proto protoreflect.FileDescriptor

var file_mesh_options_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd2, 0x03, 0x0a, 0x13, 0x4b, 0x75, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x6b,
	0x69, 0x70, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11,
	0x73, 0x6b, 0x69, 0x70, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x6b, 0x69, 0x70, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x03, 0x6b, 0x64, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x4b, 0x75, 0x6d, 0x61, 0x4b, 0x64, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x03, 0x6b, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x02, 0x77, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x4b,
	0x75, 0x6d, 0x61, 0x57, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x02, 0x77, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x73, 0x6b, 0x69,
	0x70, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x73, 0x6b, 0x69,
	0x70, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x57, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x6f, 0x5f,
	0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x6f, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x77, 0x0a, 0x0d, 0x4b, 0x75, 0x6d, 0x61, 0x57, 0x73,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6c, 0x75, 0x72, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75,
	0x72, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4f, 0x6e, 0x6c, 0x79, 0x22,
	0x58, 0x0a, 0x0e, 0x4b, 0x75, 0x6d, 0x61, 0x4b, 0x64, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x54,
	0x6f, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x6f, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73,
	0x65, 0x6e, 0x64, 0x54, 0x6f, 0x5a, 0x6f, 0x6e, 0x65, 0x3a, 0x5e, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc5, 0x91, 0xe1, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x4b, 0x75, 0x6d, 0x61,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x68, 0x71, 0x2f, 0x6b,
	0x75, 0x6d, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mesh_options_proto_rawDescOnce sync.Once
	file_mesh_options_proto_rawDescData = file_mesh_options_proto_rawDesc
)

func file_mesh_options_proto_rawDescGZIP() []byte {
	file_mesh_options_proto_rawDescOnce.Do(func() {
		file_mesh_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_mesh_options_proto_rawDescData)
	})
	return file_mesh_options_proto_rawDescData
}

var file_mesh_options_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mesh_options_proto_goTypes = []interface{}{
	(*KumaResourceOptions)(nil),         // 0: kuma.mesh.KumaResourceOptions
	(*KumaWsOptions)(nil),               // 1: kuma.mesh.KumaWsOptions
	(*KumaKdsOptions)(nil),              // 2: kuma.mesh.KumaKdsOptions
	(*descriptorpb.MessageOptions)(nil), // 3: google.protobuf.MessageOptions
}
var file_mesh_options_proto_depIdxs = []int32{
	2, // 0: kuma.mesh.KumaResourceOptions.kds:type_name -> kuma.mesh.KumaKdsOptions
	1, // 1: kuma.mesh.KumaResourceOptions.ws:type_name -> kuma.mesh.KumaWsOptions
	3, // 2: kuma.mesh.resource:extendee -> google.protobuf.MessageOptions
	0, // 3: kuma.mesh.resource:type_name -> kuma.mesh.KumaResourceOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mesh_options_proto_init() }
func file_mesh_options_proto_init() {
	if File_mesh_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mesh_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KumaResourceOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mesh_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KumaWsOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mesh_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KumaKdsOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mesh_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_mesh_options_proto_goTypes,
		DependencyIndexes: file_mesh_options_proto_depIdxs,
		MessageInfos:      file_mesh_options_proto_msgTypes,
		ExtensionInfos:    file_mesh_options_proto_extTypes,
	}.Build()
	File_mesh_options_proto = out.File
	file_mesh_options_proto_rawDesc = nil
	file_mesh_options_proto_goTypes = nil
	file_mesh_options_proto_depIdxs = nil
}
