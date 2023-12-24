FROM golang:1.19-alpine AS builder

WORKDIR /go/src/github.com/IlyaZayats/restus
COPY . .

RUN go build -o ./bin/restus ./cmd/restus

FROM alpine:latest AS runner

COPY --from=builder /go/src/github.com/IlyaZayats/restus/bin/restus /app/restus

RUN apk -U --no-cache add bash ca-certificates \
    && chmod +x /app/restus

WORKDIR /app
ENTRYPOINT ["/app/restus"]
