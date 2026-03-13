FROM golang:1.25-alpine AS builder
WORKDIR /build

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest
WORKDIR /app

ENV HOST="0.0.0.0"

RUN apk add --no-cache ca-certificates

COPY --from=builder /build/app .

CMD ["./app"]