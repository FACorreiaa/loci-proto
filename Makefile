SERVICE_NAME = $(shell basename $(shell pwd))

.PHONY: help info
.PHONY: proto-setup proto-lint proto-gen
.PHONY: go-lint

help: ## Displays a list of available makefile command and their uses
	@grep -E '^[A-z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s \n", $$1, $$2}'

info: ## Display the current service context
	@echo "SERVICE_NAME: $(SERVICE_NAME)"

proto-setup: ## Fetches deps for building .pb.go files
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

proto-lint: ## Runs protolint against all .proto files
	@protolint lint --config_path=.config/proto.yml proto

proto-gen: proto-lint ## Generates .pb.go files from .proto files
	@for protofile in ./proto/*.proto; do \
		name=$$(basename $$protofile .proto); \
		mkdir -p ./modules/$$name/generated; \
		protoc --proto_path=proto \
			--go_out=. \
			--go_opt=M$$name.proto=./modules/$$name/generated \
			--go-grpc_out=. \
			--go-grpc_opt=M$$name.proto=./modules/$$name/generated \
			$$protofile; \
	done

go-lint: ## Runs linter for .go files
	@golangci-lint run --config .config/go.yml
	@echo "Go lint passed successfully"
