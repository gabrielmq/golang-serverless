.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/InsertProduct cmd/lambda/InsertProduct/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/ListProducts cmd/lambda/ListProducts/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy -s local --verbose
