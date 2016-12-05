FROM golang:1.6.0

# add custom base package
RUN rm -rf /go/src/base
RUN  mkdir -p /go/src/base
ADD ./src/base /go/src/base

# add package to base image golang1.6.0
RUN  go get -v github.com/codegangsta/negroni && \
     go get -v github.com/gorilla/mux  && \
     go get -v github.com/xyproto/simpleredis
     #go get -v google.golang.org/grpc