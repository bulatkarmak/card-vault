PROTO_DIR = api/card-vault
GEN_DIR = api/generated

generate:
	@echo "Generating gRPC code..."
	protoc --go_out=$(GEN_DIR) --go-grpc_out=$(GEN_DIR) $(PROTO_DIR)/card-vault.proto

.PHONY: generate