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

run_fuzz:
	@echo 'Run fuzz grpc service'
	cd service_fuzz && yarn start

# SERVICE_PROTO=buzz | fuzz
gen_go:
	@echo 'Generate go files by protos'
	@echo "$(YELLOW) DIR: $(DIR) $(RESET)"
	@echo "$(GREEN) SERVICE: $(SER) $(RESET)"
	make gengo

gen_ts:
	@echo 'Generate typescript files by protos'
	@echo "$(YELLOW) DIR: $(DIR) $(RESET)"
	@echo "$(GREEN) SERVICE: $(SER) $(RESET)"
	make gents

# Execution by gateway dir
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
gengo:
	@echo 'Generate go files'
	cd $(DIR) && protoc --proto_path=../protobuf "../protobuf/$(SER).proto" \
		--go_out=./pkg/genproto --go_opt=paths=source_relative \
  	--go-grpc_out=./pkg/genproto --go-grpc_opt=paths=source_relative

gents:
	@echo 'Generate typescript files'
	cd $(DIR) && protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto \
       --ts_proto_out=./src/genproto \
       --ts_proto_opt=outputServices=grpc-js \
       --proto_path=../protobuf \
       ../protobuf/$(SER).proto

run_services:
	@echo 'Run services'
	docker-compose up --build -d

run_services_rrh:
	@echo "Run services"
	docker-compose -f docker-compose-rrh.yaml up --build -d

stop_services:
	@echo 'Stop services'
	docker-compose down

stop_services_rrh:
	@echo "Run services"
	docker-compose -f docker-compose-rrh.yaml down


# 5000 parallels connections, duration 20 seconds
load_buzz:
	@echo "Run load $(GREEN) buzz$(RESET)"
	autocannon -c 5000 -d 20 -m POST -H "Content-Type: application/json" -b '{"str": "Ping"}' http://localhost:3000/buzz

load_fuzz:
	@echo "Run load $(GREEN) fuzz$(RESET)"
	autocannon -c 5000 -d 20 -m POST -H "Content-Type: application/json" -b '{"str": "Ping"}' http://localhost:3000/fuzz