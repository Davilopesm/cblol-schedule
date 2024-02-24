.PHONY: build clean deploy

build:
	@templ generate
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-games/main cmd/get-games/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/create-games/main cmd/create-games/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose