.PHONY: all run vendor

all:
	go install ./...

run: all
	api

vendor:
	dep ensure
