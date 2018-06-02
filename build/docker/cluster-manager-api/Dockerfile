# Initial preparations
FROM golang:1.10.2-alpine
RUN apk update && apk add git make

# Let's add the code in
ADD . /go/src/github.com/samsung-cnct/cluster-manager-api

# Time to intall and rund ep
WORKDIR /go/src/github.com/samsung-cnct/cluster-manager-api

# Build package in container
RUN make -f build/Makefile container-build

# Now that we have build a static binary, we can stuff it in a bare-bones alpine image
FROM alpine
COPY --from=0 /go/src/github.com/samsung-cnct/cluster-manager-api/cluster-manager-api /cluster-manager-api

ENTRYPOINT ["/cluster-manager-api"]
