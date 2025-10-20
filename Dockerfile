FROM golang:1.25-alpine AS builder

# 设置 Go 代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /main .
CMD ["./main"]