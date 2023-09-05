GENERATED_DIR=pkg/generated
PROTO_DIR=proto

protoc:
	@if [ -d "$(GENERATED_DIR)" ]; then \
		rm -rf $(GENERATED_DIR); \
	fi
	mkdir -p $(GENERATED_DIR)

	protoc --go_out=$(GENERATED_DIR) \
    --go-grpc_out=$(GENERATED_DIR) \
    $(PROTO_DIR)/*.proto

setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc