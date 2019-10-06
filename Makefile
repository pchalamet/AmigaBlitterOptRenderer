# GOFILES := $(wildcard *.go)

# debug:
# 	echo $(GOPATH)
# 	go build -gcflags="-m"  -ldflags="-s -w" 

all: build-all test-all


convert-app:
	go build -o bin/convert.exe cmd/convert/main.go

build-all:
	go build ./...

test-all:
	go test ./...	

