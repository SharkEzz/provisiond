BUILD_DIR ?= ./build
APP_NAME = provisiond
BUILD_COMMIT_SHORT = $(shell git rev-parse --short HEAD)

.PHONY: build
build: $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME):
	go build -trimpath -o $(BUILD_DIR)/$(APP_NAME) ./cmd/provisiond.go

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)