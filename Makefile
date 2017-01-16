all: arm package

test:
	docker run --rm -v "$PWD":/usr/src/weatherServer -v "${GOPATH}":/go -w /usr/src/weatherServer golang:1.6 go test -v

build:
	docker run --rm -v "${PWD}":/go/src/github.com/salkin/weatherServer  -w /go/src/github.com/salkin/weatherServer golang:1.7 /bin/bash -c "go get && go build -v"

arm:
	docker run --rm -v "${PWD}":/usr/src/weatherServer -v ${GOPATH}:/go -e GOARM=7 -e GOOS=linux -e GOARCH=arm -w /usr/src/weatherServer golang:1.6 go build -v
	mv weatherServer bin/

package:
	docker build -t  nwik/weather-server ${PWD}

