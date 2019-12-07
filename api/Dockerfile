# build stage
FROM golang:alpine AS builder
RUN apk --no-cache add build-base git gcc
ADD . /src
RUN cd /src && go build -o go-team-api

# final stage
FROM alpine
WORKDIR /app
COPY --from=builder /src/go-team /app/
EXPOSE 8080
ENTRYPOINT ./go-team-api