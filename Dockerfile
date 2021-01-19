FROM golang:1.14-alpine

RUN apk add --no-cache bash \
    bash-doc \
    bash-completion \
    git \
    make

WORKDIR /go/app

RUN go get -d -v ./...
RUN go install -v ./...