FROM golang:1.19-alpine AS builder

WORKDIR /go/src/github.com/IlyaZayats/faculus
COPY . .

RUN go build -o ./bin/faculus ./cmd/faculus

FROM alpine:latest AS runner

COPY --from=builder /go/src/github.com/IlyaZayats/faculus/bin/faculus /app/faculus

RUN apk -U --no-cache add bash ca-certificates \
    && chmod +x /app/faculus

WORKDIR /app
ENTRYPOINT ["/app/faculus"]
