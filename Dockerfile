# 使用预编译二进制的最小化镜像
FROM alpine:latest

WORKDIR /app

# 从构建上下文复制预编译的二进制文件
ARG TARGETARCH
COPY clipboardshare-linux-$TARGETARCH /app/clipboardshare
COPY index.html README.md ./

EXPOSE 9090
USER nobody:nobody

ENTRYPOINT ["/app/clipboardshare"]