#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

THIS_DIRECTORY=$(dirname "${BASH_SOURCE}")
PROJECT_DIRECTORY=${THIS_DIRECTORY}/../..
DESTINATION_LIBRARY=pkg/generated/api
SWAGGER_DESTINATION=assets/generated/swagger

echo "Ensuring Destination Directory ( ${DESTINATION_LIBRARY} ) Exists..."
mkdir -p ${PROJECT_DIRECTORY}/${DESTINATION_LIBRARY}
echo "Ensuring Swagger Asset Direcotry ( ${SWAGGER_DESTINATION} ) Exists..."
mkdir -p ${PROJECT_DIRECTORY}/${SWAGGER_DESTINATION}

echo
echo "generating api stubs..."
echo "protoc ${PROJECT_DIRECTORY}/api/api.proto -I ${PROJECT_DIRECTORY}/api/  --go_out=plugins=grpc:${PROJECT_DIRECTORY}/${DESTINATION_LIBRARY}/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway -I$GOPATH/src"
protoc ${PROJECT_DIRECTORY}/api/api.proto -I ${PROJECT_DIRECTORY}/api/  --go_out=plugins=grpc:${PROJECT_DIRECTORY}/${DESTINATION_LIBRARY}/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway -I$GOPATH/src

echo
echo "generating REST gateway stubs..."
echo "protoc ${PROJECT_DIRECTORY}/api/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com//grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:${PROJECT_DIRECTORY}/${DESTINATION_LIBRARY}"
protoc ${PROJECT_DIRECTORY}/api/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com//grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:${PROJECT_DIRECTORY}/${DESTINATION_LIBRARY}

echo
echo "generating swagger docs..."
echo "protoc ${PROJECT_DIRECTORY}/api/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:${PROJECT_DIRECTORY}/${SWAGGER_DESTINATION}"
protoc ${PROJECT_DIRECTORY}/api/api.proto -I /usr/local/include/ -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:${PROJECT_DIRECTORY}/${SWAGGER_DESTINATION}

echo
echo "generating api docs..."
echo "protoc ${PROJECT_DIRECTORY}/api/api.proto -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway ${PROJECT_DIRECTORY}/api/api.proto --doc_out ${PROJECT_DIRECTORY}/docs/api-generated --doc_opt=markdown,api.md"
protoc ${PROJECT_DIRECTORY}/api/api.proto -I ${PROJECT_DIRECTORY}/api/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway ${PROJECT_DIRECTORY}/api/api.proto --doc_out ${PROJECT_DIRECTORY}/docs/api-generated --doc_opt=markdown,api.md
