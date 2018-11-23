FROM golang:1.11.2-stretch
USER root
RUN ls -al
ADD . /go/src/github.com/felts94/QueuesUsingStacks
WORKDIR /go/src/github.com/felts94/QueuesUsingStacks/app/
RUN go get -d ../...
RUN go build
ENTRYPOINT ["./app"]