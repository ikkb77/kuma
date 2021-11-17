#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

echo "Building Envoy for Linux"

mkdir -p "$(dirname ${BINARY_PATH})"

SOURCE_DIR="${SOURCE_DIR}" "${KUMA_DIR:-.}/tools/envoy/fetch_sources.sh"

BUILD_CMD=${BUILD_CMD:-"BAZEL_BUILD_EXTRA_OPTIONS=\"${BAZEL_BUILD_EXTRA_OPTIONS:-}\" ./ci/do_ci.sh bazel.release.server_only"}

ENVOY_BUILD_SHA=$(curl --fail --location --silent https://raw.githubusercontent.com/envoyproxy/envoy/"${ENVOY_TAG}"/.bazelrc | grep envoyproxy/envoy-build-ubuntu | sed -e 's#.*envoyproxy/envoy-build-ubuntu:\(.*\)#\1#'| uniq)
ENVOY_BUILD_IMAGE="envoyproxy/envoy-build-ubuntu:${ENVOY_BUILD_SHA}"
LOCAL_BUILD_IMAGE="envoy-builder:${ENVOY_TAG}"

echo "BUILD_CMD=${BUILD_CMD}"

docker build -t "${LOCAL_BUILD_IMAGE}" --progress=plain \
  --build-arg ENVOY_BUILD_IMAGE="${ENVOY_BUILD_IMAGE}" \
  --build-arg BUILD_CMD="${BUILD_CMD}" \
  -f tools/envoy/Dockerfile.build-ubuntu "${SOURCE_DIR}"

# copy out the binary
id=$(docker create "${LOCAL_BUILD_IMAGE}")
docker cp "$id":/envoy-sources/linux/amd64/build_envoy_release_stripped/envoy "${BINARY_PATH}"
docker rm -v "$id"
