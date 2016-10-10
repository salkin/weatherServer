all: arm package

test:
	docker run --rm -v "$PWD":/usr/src/weatherServer -v "${GOPATH}":/go -w /usr/src/weatherServer golang:1.6 go test -v

build:
	docker run --rm -v "${PWD}":/usr/src/weatherServer -v ${GOPATH}:/go -w /usr/src/weatherServer golang:1.6 go build -v

arm:
	docker run --rm -v "${PWD}":/usr/src/weatherServer -v ${GOPATH}:/go -e GOARM=7 -e GOOS=linux -e GOARCH=arm -w /usr/src/weatherServer rpi-golang go build -v

package:
	docker build -t  weatherServer ${PWD}
	docker save -o weatherServer

