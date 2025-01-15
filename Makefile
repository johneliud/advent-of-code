APP_NAME := advent-of-code
SRC := $(shell find . -type f -name '*.go')

.PHONY: all format run clean

all: format run

format:
	@gofmt -w $(SRC)

run:
	@go run .

clean:
	@go clean
