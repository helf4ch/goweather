APP_NAME := goweather
BUILD_DIR := build
BUILD_PATH := $(BUILD_DIR)/$(APP_NAME)
INSTALL_DIR := $(HOME)/.local/bin/
INSTALL_PATH := $(HOME)/.local/bin/$(APP_NAME)
SRC := src

all: build

build:
	mkdir $(BUILD_DIR)
	go build -o $(BUILD_PATH) $(SRC)/main.go

install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BUILD_PATH) $(INSTALL_PATH)

clean:
	rm -rf ./$(BUILD_DIR)
