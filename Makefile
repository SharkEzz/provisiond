BUILD_DIR ?= ./build
APP_NAME = provisiond
BUILD_TAG = $(shell git describe --tags --abbrev=0)

.PHONY: build
build: $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME):
	go build -trimpath -ldflags="-X 'main.version=$(BUILD_TAG)'" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/provisiond.go

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)