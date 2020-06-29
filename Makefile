# -include .env

# VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
# GOBASE := $(shell pwd)
GOPATH := $(HOME)/go
GOBIN := $(GOPATH)/bin
PATH := $(PATH):$(GOBIN)

# Go parameters
GOCMD=go
GOMOD=$(GOCMD) mod
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

BINARY_NAME=fitfile-tool

.DEFAULT_GOAL := help

## clean: Clean
clean:
	# @echo no clean actions
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GOCLEAN) ./..
	rm -rf build

build:
	 @GOPATH=$(GOPATH) $(GOINSTALL) ./...

run:
	@GOPATH=$(GOPATH) $(GORUN) -v ./...

.PHONY: help build

$(OBJDIR):
	mkdir -p $(OBJDIR)

all: help

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
