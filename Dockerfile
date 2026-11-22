# ---------- 构建阶段：编译 cert-monitor ----------
FROM golang:1.16 AS builder

# 腾讯 Go 模块代理（国内构建环境）；CGO 关闭后为纯静态二进制，不依赖运行镜像的 libc
ENV GOPROXY=https://mirrors.cloud.tencent.com/go/ \
    CGO_ENABLED=0

WORKDIR /src

COPY . .

RUN go build -o bin/cert-monitor cmd/main.go

# ---------- 运行阶段：最小运行时（debian-slim + acme.sh）----------
FROM actors315/go:acme

LABEL maintainer="actors315 <actors315@gmail.com>"

WORKDIR /usr/local/qcloud-tools

COPY --from=builder /src/bin/cert-monitor ./cert-monitor
COPY config/config.simple.yaml ./config/config.yaml
COPY config/issue-template.tpl ./config/issue-template.tpl
COPY Dockerstart /start

RUN chmod +x /start ./cert-monitor

EXPOSE 80

CMD ["/start"]
