#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

THIS_DIRECTORY=$(dirname "$BASH_SOURCE")
PROJECT_DIRECTORY=$THIS_DIRECTORY/../..

echo "Ensuring directories exist"
mkdir -p ${PROJECT_DIRECTORY}/pkg/generated/ui/data/swagger \
         ${PROJECT_DIRECTORY}/pkg/generated/ui/data/protobuf \
         ${PROJECT_DIRECTORY}/pkg/generated/ui/data/swaggerjson \
         ${PROJECT_DIRECTORY}/pkg/generated/ui/data/homepage \
         ${PROJECT_DIRECTORY}/assets/generated/swagger

echo "generating swagger-ui bindata file..."
go-bindata-assetfs -pkg swagger -o ${PROJECT_DIRECTORY}/pkg/generated/ui/data/swagger/bindata.go -prefix ${PROJECT_DIRECTORY}/third_party ${PROJECT_DIRECTORY}/third_party/swagger-ui

echo "generating protobuf files..."
go-bindata-assetfs -pkg protobuf -o ${PROJECT_DIRECTORY}/pkg/generated/ui/data/protobuf/bindata.go -prefix ${PROJECT_DIRECTORY}/api ${PROJECT_DIRECTORY}/api

echo "generating swagger.json files..."
go-bindata-assetfs -pkg swaggerjson -o ${PROJECT_DIRECTORY}/pkg/generated/ui/data/swaggerjson/bindata.go -prefix ${PROJECT_DIRECTORY}/assets/generated/swagger ${PROJECT_DIRECTORY}/assets/generated/swagger

echo "generating homepage files..."
go-bindata-assetfs -pkg homepage -o ${PROJECT_DIRECTORY}/pkg/generated/ui/data/homepage/bindata.go -prefix ${PROJECT_DIRECTORY}/assets/homepage ${PROJECT_DIRECTORY}/assets/homepage
