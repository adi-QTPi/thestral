FROM golang:1.25-alpine AS builder

ARG APP_NAME = thestral

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

ARG APP_NAME = thestral

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/bin/${APP_NAME} ./binary

#to be supplied in run command , backup value 8080,
ENV PORT=8080 
EXPOSE $PORT

CMD ["./binary"]