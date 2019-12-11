# build stage
FROM golang:alpine AS builder
RUN apk --no-cache add build-base git gcc
ADD . /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o go-team

# final stage
FROM scratch
WORKDIR /app
COPY --from=builder /src/go-team /app/
EXPOSE 8080
ENTRYPOINT ./go-team