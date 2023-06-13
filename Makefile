MODULE  := $(shell go list -m)/server
BIN_DIR :=  bin
BIN := $(BIN_DIR)/master

all: $(BIN)

$(BIN): $(BIN_DIR) ./**/*.go
	go build -o $@ $(MODULE)

$(BIN_DIR):
	mkdir -p $@