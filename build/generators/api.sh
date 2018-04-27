#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

THIS_DIRECTORY=$(dirname "${BASH_SOURCE}")
PROJECT_DIRECTORY=${THIS_DIRECTORY}/../..

echo
echo "generating api stubs..."
echo "protoc ${PROJECT_DIRECTORY}/api/api.proto -I ${PROJECT_DIRECTORY}/api/  --go_out=plugins=grpc:${PROJECT_DIRECTORY}/pkg/api/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway -I$GOPATH/src"
protoc ${PROJECT_DIRECTORY}/api/api.proto -I ${PROJECT_DIRECTORY}/api/  --go_out=plugins=grpc:${PROJECT_DIRECTORY}/pkg/api/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway -I$GOPATH/src

echo
echo "generating REST gateway stubs..."
echo "protoc ${PROJECT_DIRECTORY}/api/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com//grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com//grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:${PROJECT_DIRECTORY}/pkg/api"
protoc ${PROJECT_DIRECTORY}/api/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com//grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com//grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:${PROJECT_DIRECTORY}/pkg/api

echo
echo "generating swagger docs..."
echo "protoc ${PROJECT_DIRECTORY}/api/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:${PROJECT_DIRECTORY}/api/swagger"
protoc ${PROJECT_DIRECTORY}/api/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:${PROJECT_DIRECTORY}/api/swagger

echo
echo "generating api docs..."
echo "protoc ${PROJECT_DIRECTORY}/api/api.proto -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway ${PROJECT_DIRECTORY}/api/api.proto --doc_out ${PROJECT_DIRECTORY}/docs/api-generated --doc_opt=markdown,api.md"
protoc ${PROJECT_DIRECTORY}/api/api.proto -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway ${PROJECT_DIRECTORY}/api/api.proto --doc_out ${PROJECT_DIRECTORY}/docs/api-generated --doc_opt=markdown,api.md
