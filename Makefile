.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/pdf-converter main.go

clean:
	rm -rf ./bin

test: build
	go test -v ./...
