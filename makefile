GENERATED_DIR=internal/generated
PROTO_DIR=internal/proto

protoc:
	@if [ -d "$(GENERATED_DIR)" ]; then \
		rm -rf $(GENERATED_DIR); \
	fi
	mkdir -p $(GENERATED_DIR)

	protoc --go_out=internal/generated \
    --go-grpc_out=internal/generated \
    internal/proto/*.proto

setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc