BINARY_NAME=git-vendor
MAIN_PATH=cmd

build:
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}-mac-arm ${MAIN_PATH}/main.go
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-mac ${MAIN_PATH}/main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ${MAIN_PATH}/main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows ${MAIN_PATH}/main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}-mac-arm
	rm ${BINARY_NAME}-mac
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test:
	go test ./...

dep:
	go mod download

vet:
	go vet .../

lint:
	golangci-lint run --enable-all