# Copyright 2016 The Kubernetes Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Build the greeter_server example

# Usage:
#   [VERSION=v3] [REGISTRY="gcr.io/google_containers"] make build
VERSION?=v3
REGISTRY?=192.168.1.10:5000

release: clean build push clean

# builds a docker image that builds the app and packages it into a minimal docker image
build:
	docker build --rm --force-rm -t ${REGISTRY}/greeter-builder .
	docker run --rm ${REGISTRY}/greeter-builder | docker build -t "${REGISTRY}/greeter:${VERSION}" -

# push the image to an registry
push:
	docker push ${REGISTRY}/greeter:${VERSION}

# remove previous images and containers
clean:
	docker rm -f ${REGISTRY}/greeter-builder 2> /dev/null || true
	docker rmi -f ${REGISTRY}/greeter-builder || true
	docker rmi -f "${REGISTRY}/greeter:${VERSION}" || true

.PHONY: release clean build push
