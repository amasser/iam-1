PROJECT = iam
VET_REPORT = vet.report
TEST_REPORT = tests.xml
GOARCH = amd64

VERSION?=1.0.0
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

GITHUB_USERNAME=maurofran
BUILD_DIR=${GOPATH}/src/github.com/${GITHUB_USERNAME}/${PROJECT}
CURRENT_DIR=$(shell pwd)
TARGET_DIR=target

LD_FLAGS = -ldflags "-X main.Version=${VERSION} -X main.Commit=${COMMIT} -X main.Branch=${BRANCH}"

# Build all the project
all: clean protoc linux darwin windows

protoc:
	mkdir -p internal/app/ports/adapter/grpc; \
	protoc -I/usr/local/include -Iapi -I${GOPATH}/src --go_out=plugins=grpc:internal/app/ports/adapter/grpc api/iam.proto; \

linux: protoc
	GOOS=linux GOARCH=${GOARCH} go build ${LD_FLAGS} -o ${TARGET_DIR}/iamd-linux-${GOARCH} cmd/iamd/main.go; \
	GOOS=linux GOARCH=${GOARCH} go build ${LD_FLAGS} -o ${TARGET_DIR}/iam-linux-${GOARCH} cmd/iam/main.go; \

darwin: protoc
	GOOS=darwin GOARCH=${GOARCH} go build ${LD_FLAGS} -o ${TARGET_DIR}/iamd-darwin-${GOARCH} cmd/iamd/main.go; \
	GOOS=darwin GOARCH=${GOARCH} go build ${LD_FLAGS} -o ${TARGET_DIR}/iam-darwin-${GOARCH} cmd/iam/main.go; \

windows: protoc
	GOOS=windows GOARCH=${GOARCH} go build ${LD_FLAGS} -o ${TARGET_DIR}/iamd-windows-${GOARCH}.exe cmd/iamd/main.go; \
	GOOS=windows GOARCH=${GOARCH} go build ${LD_FLAGS} -o ${TARGET_DIR}/iam-windows-${GOARCH}.exe cmd/iam/main.go; \

clean:
	-rm -f ${TARGET_DIR}

.PHONY: linux darwin windows protoc clean	
