GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test
DEPCMD=dep

# Variables for binary version information
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_TAG=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

# Associating version information variables with data
VERSION_PACKAGE=github.com/samsung-cnct/cluster-manager-api/pkg/version
LD_DATA_GIT_COMMIT=-X ${VERSION_PACKAGE}.gitCommit=${GIT_COMMIT}
LD_DATA_GIT_TAG=-X ${VERSION_PACKAGE}.gitVersion=${GIT_TAG}
LD_DATA_GIT_DIRTY=-X ${VERSION_PACKAGE}.gitTreeState=${GIT_DIRTY}
LD_DATA_BUILD_DATE=-X ${VERSION_PACKAGE}.buildDate=${BUILD_DATE}
LD_DATA=${LD_DATA_GIT_COMMIT} ${LD_DATA_GIT_TAG} ${LD_DATA_GIT_DIRTY} ${LD_DATA_BUILD_DATE}

# Creating static binary
STATIC_FLAGS=-w -extldflags "-static"

# Generating ldflag string
LDFLAGS = '${STATIC_FLAGS} ${LD_DATA}'

# Location of
CMA_BINARY_NAME=cluster-manager-api
CMA_SOURCE_FILE=cmd/cluster-manager-api/main.go

# build binary command
BUILD_CMA_BINARY_CMD=${GOBUILD} -a -tags netgo -ldflags ${LDFLAGS} -o ${CMA_BINARY_NAME} ${CMA_SOURCE_FILE}

# protoc variables
PROTOC_VERSION=3.6.1
PROTOC_DARWIN_FILENAME=protoc-${PROTOC_VERSION}-osx-x86_64.zip
PROTOC_DARWIN_DOWNLOAD_URL=https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_DARWIN_FILENAME}

# By default we will most likely want to build a linux build of the application
all: new-dep cma-bin-linux-amd64

container-build: install-dep dep-ensure cma-bin-linux-amd64

# Clean will clean up dep and remove any binaries
clean: remove-binary

# Installs the required tools
install-tools: install-dep
	go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elazarl/go-bindata-assetfs/...


# Installs tools for a mac
install-tools-darwin: install-protoc-darwin install-tools

install-protoc-darwin:
	curl -L ${PROTOC_DARWIN_DOWNLOAD_URL} > /tmp/${PROTOC_DARWIN_FILENAME}
	sudo unzip -o /tmp/${PROTOC_DARWIN_FILENAME} -d /usr/local/
	rm /tmp/${PROTOC_DARWIN_FILENAME}


# Install dep will install dep in $GOROOT/bin
install-dep:
	go get -u github.com/golang/dep/cmd/dep

# New Dep will just blow away Gopkg, vendor and do a dep init
new-dep:
	rm -rf vendor Gopkg.toml Gopkg.lock
	$(DEPCMD) init

# For builds where we're just going to go along with whatever Gopkg is saying
dep-ensure:
	$(DEPCMD) ensure

# Run the code generators for the user
generators:
	./build/generators/api.sh
	./build/generators/swagger-dist-adjustment.sh
	./build/generators/bindata.sh

# Build a darwin (mac) binary
cma-bin-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ${BUILD_CMA_BINARY_CMD}

# Build a linux amd64 binary
cma-bin-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${BUILD_CMA_BINARY_CMD}

# Delete an existing binary / cleanup
remove-binary:
	rm ${CMA_BINARY_NAME}