# 构建阶段
FROM golang:1.20-alpine3.18 AS builder

WORKDIR /app

# 设置 GOPROXY
ENV GOPROXY=https://goproxy.cn,direct

# 复制 go mod 和 sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 运行阶段
FROM alpine:3.18

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

#HEALTHCHECK --timeout=10s --start-period=60s --interval=60s \
#  CMD wget --spider -q http://localhost:8080/hello/healthz

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]