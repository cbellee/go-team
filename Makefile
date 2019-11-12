VERSION := $(shell git rev-parse --short HEAD)

build:
	docker build -t belstarr/go-team:${VERSION} .

tag:
	docker tag belstarr/go-team:${VERSION} belstarr/go-team:latest

run:
	docker run -p 8080:8080 belstarr/go-team:${VERSION

push:
	docker push belstarr/go-team:latest