# Default dir gateway
DIR ?=gateway

# Generate service proto
SER ?=buzz

GREEN := \033[0;32m
RESET := \033[0m
YELLOW := \033[0;33m

run_gateway:
	@echo 'Run gateway service'
	cd gateway && go run ./cmd/main.go


run_buzz:
	@echo 'Run buzz grpc service'
	cd service_buzz && go run cmd/*.go

# SERVICE_PROTO=buzz | fuzz
gen_go:
	@echo 'Generate go files by protos'
	@echo "$(YELLOW) DIR: $(DIR) $(RESET)"
	@echo "$(GREEN) SERVICE: $(SER) $(RESET)"
	make gen

# Execution by gateway dir
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
gen:
	@echo 'Generate types'
	cd $(DIR) && protoc --proto_path=../protobuf "../protobuf/$(SER).proto" \
		--go_out=./pkg/genproto --go_opt=paths=source_relative \
  	--go-grpc_out=./pkg/genproto --go-grpc_opt=paths=source_relative

run_services:
	@echo 'Run services'
	docker-compose up --build -d

stop_services:
	@echo 'Stop services'
	docker-compose down