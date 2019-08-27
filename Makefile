# GOFILES := $(wildcard *.go)

# debug:
# 	echo $(GOPATH)
# 	go build -gcflags="-m"  -ldflags="-s -w" 

convert-app:
	go build -o bin/convert.exe cmd/convert/main.go

all: convert-app
