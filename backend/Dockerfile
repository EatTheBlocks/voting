FROM golang:1.17-alpine AS builder
RUN apk add gcc g++ git libc-dev make
WORKDIR /app
COPY . .
RUN go get ./...
RUN go build -o hub  ./cmd/hub

FROM alpine
COPY --from=builder /app/hub /hub
ENTRYPOINT ["/hub"]