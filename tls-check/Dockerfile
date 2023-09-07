FROM golang:1.21.0-alpine3.18 AS builder

WORKDIR /app
COPY . .
RUN apk add upx && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -trimpath -o bin/main main.go && \
    upx --best --lzma bin/main

FROM scratch

EXPOSE 5555
COPY --from=builder /app/bin/main ./main
CMD ["./main"]