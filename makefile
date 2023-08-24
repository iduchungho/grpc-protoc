protoc:
	protoc --go_out=pkg/generated \
    --go-grpc_out=pkg/generated \
    pkg/proto/greeter.proto