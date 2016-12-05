#!/bin/bash

SERVICES=("guestbook" "greeter")
#SERVICES=("greeter")

function init {
    echo "Add base and third-party package to golang1.6.0 image"
    cd /home/vagrant/projects/guessbook-go/
    docker build --rm --force-rm -t golang:1.6.0 .
}

function deploy_service {
    echo "Create or update service "${service}
    sudo su - vagrant
    cd ~/projects/guessbook-go/src/${service}

    # compile and package image then publish image to docker
    VERSION=$1 REGISTRY="192.168.1.10:5000" make release
    VERSION=$1 REGISTRY="192.168.1.10:5000" make clean

    # roll update service web image version
    cd ~/projects/guessbook-go
    sed -i 's/{version}/'$1'/g' ${service}-deployment.yaml

    kubectl apply -f ./config/${service}-deployment.yaml
    kubectl apply -f ./config/${service}-service.yaml
}

init

for service in ${SERVICES[@]}
do
    deploy_service
done