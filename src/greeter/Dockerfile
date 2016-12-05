FROM golang:1.6.0

RUN  mkdir -p /go/src/greeter
ADD . /go/src/greeter
WORKDIR /go/src

RUN cd greeter && go get && go build -o ../bin/greeter
RUN cp greeter/greeter/Dockerfile .

CMD tar czf - .
