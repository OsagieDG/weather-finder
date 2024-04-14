PROTO_DIR := api/v1
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)
GO_OUT_DIR := $(PROTO_DIR)
GRPC_OUT_DIR := $(PROTO_DIR)
GO_OUT := $(patsubst $(PROTO_DIR)/%.proto,$(GO_OUT_DIR)/%.pb.go,$(PROTO_FILES))
GRPC_OUT := $(patsubst $(PROTO_DIR)/%.proto,$(GRPC_OUT_DIR)/%.pb.gw.go,$(PROTO_FILES))

.DEFAULT_GOAL := build

compile: $(GO_OUT) $(GRPC_OUT)

$(GO_OUT): $(PROTO_FILES)
	protoc $< --go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative --proto_path=$(PROTO_DIR)

$(GRPC_OUT): $(PROTO_FILES)
	protoc $< --go-grpc_out=$(GRPC_OUT_DIR) --go-grpc_opt=paths=source_relative --proto_path=$(PROTO_DIR)


build: compile
	go build -o bin/app ./cmd/finder/

run:
	@./bin/app

clean:
	rm -f $(GO_OUT) $(GRPC_OUT) bin/app

.PHONY: compile build run clean



