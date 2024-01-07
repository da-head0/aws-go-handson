.PHONY: build clean deploy

build:
	go get github.com/aws/aws-lambda-go/lambda
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
