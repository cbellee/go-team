VERSION := $(shell git rev-parse --short HEAD)

build:
	docker build -t belstarr/go-team-api:${VERSION} .

tag:
	docker tag belstarr/go-team-api:${VERSION} belstarr/go-team-api:latest

run:
	docker run -p 8080:8080 belstarr/go-team-api:${VERSION}

push:
	docker push belstarr/go-team-api:${VERSION}

check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	GO111MODULE=on go mod vendor  && GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.yaml