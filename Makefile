build:
	mkdir -p internal/app/ports/adapter/grpc
	protoc -I/usr/local/include -Iapi -I${GOPATH}/src --go_out=plugins=grpc:internal/app/ports/adapter/grpc api/iam.proto
