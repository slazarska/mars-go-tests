# Stage 1: builder
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Copy all source files *before* tidy
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./test ./test

# Now that all Go files are present, tidy will catch all dependencies
RUN go mod tidy

# Run tests
RUN go test ./... -v

# Stage 2: runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/cmd /app/cmd
COPY --from=builder /app/internal /app/internal

CMD ["/bin/sh"]
