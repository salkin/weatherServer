all: build package

test:
	docker run --rm -v "$PWD":/go/src/github.com/salkin/weatherServer -v "${GOPATH}":/go -w /go/src/github.com/salkin/weatherServer golang:1.6 go test -v

build:
	docker run --rm --privileged -v "${PWD}":/go/src/github.com/salkin/weatherServer  -w /go/src/github.com/salkin/weatherServer golang:1.6 /bin/bash -c "go get && go build -v"

arm:
	docker run --rm -v "${PWD}":/go/src/github.com/salkin/weatherServer -v ${GOPATH}:/go -e GOARM=7 -e GOOS=linux -e GOARCH=arm -w /go/src/github.com/salkin/weatherServer golang:1.6 go build -v
	mv weatherServer bin/

package:
	docker build -t  nwik/weather-server ${PWD}

