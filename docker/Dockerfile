FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS builder

# Docker fills TARGETARCH automatically
ARG TARGETARCH

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=$TARGETARCH

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p bin
RUN go build -o bin/thestral ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/thestral ./thestral

CMD ["./thestral"]