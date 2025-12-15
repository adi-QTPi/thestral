FROM golang:1.25-alpine AS builder

ARG APP_NAME=thestral

RUN apk add --no-cache git make

ENV CGO_ENABLED=0 \
GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

# Running stage
FROM alpine:latest

ARG APP_NAME=thestral

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/bin/${APP_NAME} ./binary

ENV PORT=7007
ENV ADMIN_PORT=7008

EXPOSE 7007
EXPOSE 7008

CMD ["./binary"]