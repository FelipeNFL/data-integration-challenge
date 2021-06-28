FROM golang:1.16-alpine

ENV SRC_PATH $GOPATH/yawoen
ENV GO111MODULE=on

WORKDIR $SRC_PATH

RUN apk --no-cache add git gcc libc-dev

COPY . $SRC_PATH

RUN go build apis/matching_api.go
RUN go build apis/integration_api.go
RUN go build scripts/populate_db.go

EXPOSE 8080
EXPOSE 8081