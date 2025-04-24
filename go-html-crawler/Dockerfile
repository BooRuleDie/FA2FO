# Stage 1: Build binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install CA certificates (needed later in scratch stage)
RUN apk add --no-cache ca-certificates

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Build statically-linked Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /goapp *.go

# Stage 2: Minimal scratch image
FROM scratch

# Copy binary and CA certificates from builder
COPY --from=builder /goapp /goapp
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


ENTRYPOINT ["/goapp"]
