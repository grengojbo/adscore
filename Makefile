MAINTAINER=oleg.dolya@gmail.com
HOMEPAGE=https://github.com/grengojbo/ads
DESCRIPTION="Core Advertising System"

RELEASE_DIR=release
RELEASE_BUILD=builds

OSNAME=$(shell uname)

GO=$(shell which go)

CUR_TIME=$(shell date '+%Y-%m-%d_%H:%M:%S')
# Program version
VERSION=$(shell cat VERSION)

# Binary name for bintray
#BIN_NAME=$(shell basename $(abspath ./))
BIN_NAME=adscore

# Project name for bintray
PROJECT_NAME=$(shell basename $(abspath ./))
PROJECT_DIR=$(shell pwd)

# Project url used for builds
# examples: github.com, bitbucket.org
REPO_HOST_URL=github.com.org

# Grab the current commit
GIT_COMMIT="$(shell git rev-parse HEAD)"

# Check if there are uncommited changes
GIT_DIRTY="$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)"

# Add the godep path to the GOPATH
#GOPATH=$(shell godep path):$(shell echo $$GOPATH)

default: help

help:
	@echo "..............................................................."
	@echo "Project: $(PROJECT_NAME) | current dir: $(PROJECT_DIR)"
	@echo "version: $(VERSION) GIT_DIRTY: $(GIT_DIRTY)\n"
	@#echo "Autocomplete exec -> PROG=$(BIN_NAME) source ./autocomplete/bash_autocomplete\n"
	@echo "make init    - Load godep"
	@echo "make save    - Save project libs"
	@echo "make install - Install packages"
	@echo "make clean   - Clean .orig, .log files"
	@echo "make test     - Run project debug mode"
	@echo "make docs"   - Project documentation
	@echo "...............................................................\n"

init:
	@go get github.com/tools/godep

save:
	@godep save

install:
	@#go get -v -u
	@go get -v -u github.com/jackc/pgx
	@go get -v -u github.com/smartystreets/goconvey/convey
	@#go get -v -u gopkg.in/inconshreveable/log15.v2
	@#go get -v -u github.com/gorilla/sessions
	@#go get -v -u github.com/gin-gonic/gin
	@#go get -v -u github.com/gin-gonic/contrib/sessions
	@#go get -v -u github.com/codegangsta/cli
	@#go get -v -u github.com/nu7hatch/gouuid
	@#go get -v -u github.com/mssola/user_agent
	@#go get -v -u github.com/azumads/faker

clean:
	@test ! -e ./${BIN_NAME} || rm ./${BIN_NAME}
	@git gc --prune=0 --aggressive
	@find . -name "*.orig" -type f -delete
	@find . -name "*.log" -type f -delete

test:
	@echo "Start testing..."
	@go test -v ./...
	@#API_PATH=$(PROJECT_DIR) ginkgo -v -r

push:
	@git add -A
	@git ci -am "new release v$(VERSION) COMMIT: $(GIT_COMMIT)"
	@git push

docs:
	godoc -http=:6060 -index
