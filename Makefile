all: build package

test:
	docker run --rm -v "$PWD":/usr/src/weatherServer -v "${GOPATH}":/go -w /usr/src/weatherServer golang:1.6 go test -v

build:
	docker run --rm -v "${PWD}":/usr/src/weatherServer -v ${GOPATH}:/go -w /usr/src/weatherServer golang:1.6 go build -v

arm:
	docker run --rm -v "${PWD}":/usr/src/weatherServer -v ${GOPATH}:/go -e GOOS=linux -e GOARCH=arm64 -w /usr/src/weatherServer golang:1.6 go build -v



