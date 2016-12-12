FROM golang:1.7

RUN rm -rf /go/src/greeter
ADD . /go/src/greeter
WORKDIR /go/src

RUN cd greeter && go get -v && go build -o ../bin/greeter
RUN cp greeter/greeter/Dockerfile .

CMD tar czf - .