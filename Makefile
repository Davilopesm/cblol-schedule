# build:
# 	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/main.go

# clean:
# 	@go clean
# 	@rm -rf ./bin

# start:
# 	sls offline --host 0.0.0.0

# # run:
# # 	@templ generate
# # 	@go run cmd/main.go

# format:
# 	gofmt -s -w .
.PHONY: build clean deploy

build:
	@templ generate
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose