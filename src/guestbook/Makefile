# Usage:
#   [VERSION=v3] [REGISTRY="gcr.io/google_containers"] make build
VERSION?=v3
REGISTRY?=192.168.1.10:5000

release: clean build push clean

# builds a docker image that builds the app and packages it into a minimal docker image
build:
	docker build --rm --force-rm -t ${REGISTRY}/guestbook-builder .
	docker run --rm ${REGISTRY}/guestbook-builder | docker build -t "${REGISTRY}/guestbook:${VERSION}" -

# push the image to an registry
push:
	docker push ${REGISTRY}/guestbook:${VERSION}
	
# remove previous images and containers
clean: 
	docker rm -f ${REGISTRY}/guestbook-builder 2> /dev/null || true
	docker rmi -f ${REGISTRY}/guestbook-builder || true
	docker rmi -f "${REGISTRY}/guestbook:${VERSION}" || true

.PHONY: release clean build push
