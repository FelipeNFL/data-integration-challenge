FROM golang:1.16-alpine

ENV SRC_PATH $GOPATH/yawoen
ENV GO111MODULE=on

WORKDIR $SRC_PATH

RUN apk --no-cache add git gcc libc-dev

COPY . $SRC_PATH

# RUN go get -d -v

# RUN go build main.go