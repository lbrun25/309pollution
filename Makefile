BINARY_NAME := 309pollution

.PHONY: all
all: build

build:
	go build -o ${BINARY_NAME}