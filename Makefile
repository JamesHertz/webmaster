MODULE  := $(shell go list -m)/server
BIN_DIR :=  bin
BIN := $(BIN_DIR)/master

all: $(BIN)

$(BIN): ./**/*.go
	go build -o $@ $(MODULE)

.PHONNY: clean
clean:
	@rm -rf $(BIN_DIR)