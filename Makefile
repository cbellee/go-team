VERSION := $(shell git rev-parse --short HEAD)

build:
	docker build -t belstarr/go-team:${VERSION} .

tag:
	docker tag belstarr/go-team:${VERSION} belstarr/go-team:latest

run:
	docker run -p 8080:8080 belstarr/go-team:${VERSION}

push:
	docker push belstarr/go-team:latest

check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	GO111MODULE=on go mod vendor  && GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.yaml