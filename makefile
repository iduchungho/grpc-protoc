protoc:
	protoc --go_out=pkg/plugin \
    --go-grpc_out=pkg/plugin \
    pkg/proto/greeter.proto