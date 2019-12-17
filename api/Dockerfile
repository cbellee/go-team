# build stage
FROM golang:alpine AS builder
RUN apk --no-cache add build-base git gcc
ADD . /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o go-team

# Create a "nobody" non-root user for the next image by crafting an /etc/passwd
# file that the next image can copy in. This is necessary since the next image
# is based on scratch, which doesn't have adduser, cat, echo, or even sh.
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

# Final stage in scratch container
FROM scratch
WORKDIR /app
COPY --from=builder /src/go-team /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the /etc_passwd file we created in the builder stage into /etc/passwd in
# the target stage. This creates a new non-root user as a security best
# practice.
COPY --from=builder /etc_passwd /etc/passwd

# added a comment
# expose port 
EXPOSE 8080

# Run as the new non-root by default
USER nobody
CMD ["./go-team"]
