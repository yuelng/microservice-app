#!/usr/bin/env bash
cd /home/vagrant/projects/guessbook-go/src/web
docker run golang:1.6.0 mkdir -p /go/src/guestbook
docker run golang:1.6.0 go get -v github.com/codegangsta/negroni
docker run golang:1.6.0 go get -v github.com/gorilla/mux
docker run golang:1.6.0 go get -v github.com/xyproto/simpleredis
docker commit $(docker ps -lq) golang:1.6.0

# publish docker
VERSION=$1 REGISTRY="192.168.1.10:5000" make release

# roll update service
cd /home/vagrant/projects/guessbook-go
sed -i 's/{version}/'$1'/g' guestbook-deployment.yaml
kubectl apply -f .
#kubectl rolling-update guestbook -f guestbook-controller.json
#kubectl rolling-update guestbook --image=192.168.1.10:5000/guestbook:$1