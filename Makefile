.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./code/main
	
build:
	GOOS=linux GOARCH=amd64 go build -o code/main ./code