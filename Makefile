BINARY_NAME=qrcode-generator
BUILD_DIR=build

.PHONY: all build darwin linux clean

all: darwin linux

build:
	go build -o $(BINARY_NAME) .

darwin: $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64

linux: $(BUILD_DIR)/$(BINARY_NAME)-linux-386 $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64

$(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o $@ .

$(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o $@ .

$(BUILD_DIR)/$(BINARY_NAME)-linux-386:
	GOOS=linux GOARCH=386 go build -o $@ .

$(BUILD_DIR)/$(BINARY_NAME)-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $@ .

clean:
	rm -rf $(BUILD_DIR) $(BINARY_NAME)
