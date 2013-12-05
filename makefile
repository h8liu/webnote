.PHONY: all fmt

all:
	go build
	go vet

fmt:
	go fmt
