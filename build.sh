#!/bin/bash

#SERVICES=("guestbook" "greeter")
#SERVICES=("greeter")
SERVICES=("guestbook")
VERSION=$1
# upgrade_package

function upgrade_package {
    # if not upgrade or not install new package don't execute it
    cd /home/vagrant/projects/guessbook-go/
    docker build --rm --force-rm -t golang:1.7 .
    # squash docker layer
    docker-squash -f a99621b7e319 -t golang:1.7 golang:1.7
}

function deploy_service {
    echo "Create or update service "${service}

    cd /home/vagrant/projects/guessbook-go/src/${service}

    # compile and package image then publish image to docker
    VERSION=${VERSION} REGISTRY="192.168.1.10:5000" make release

    # roll update service web image version
    cd /home/vagrant/projects/guessbook-go/config
    sed -i 's/{version}/'${VERSION}'/g' ${service}-deployment.yaml

    kubectl apply -f ${service}-deployment.yaml
    kubectl apply -f ${service}-service.yaml
}

function delete_images {
    docker rmi $(docker images|grep "192.168.1.10" |awk '{print $3}')
}

for service in ${SERVICES[@]}
do
    deploy_service
done

# delete_images