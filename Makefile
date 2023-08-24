all: build

SHELL := env ENV=$(ENV) $(SHELL)
ENV ?= dev

BIN_DIR = $(PWD)/bin

clean:
	rm -rf bin/*

dependencies:
	go mod download

go-build:
	go build -o ./bin/app ./api/main.go

build: dependencies go-build



