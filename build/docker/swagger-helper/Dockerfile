FROM golang:1.10.1-alpine
RUN apk update && apk add git protobuf-dev bash && \
    go get github.com/jteeuwen/go-bindata/... && \
    go get github.com/elazarl/go-bindata-assetfs/...

