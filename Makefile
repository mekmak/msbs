.DEFAULT_GOAL := build

export GOBIN=${HOME}/bin/
export GO111MODULE=on

build:
	go build -o bin/msbs ./cmd/msbs

install:
	go install ./cmd/msbs

clean:
	go clean ./cmd/msbs