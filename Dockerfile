# 使用多阶段构建确保最小化镜像
# 第一阶段：构建二进制（可省略，因为已在 GitHub Actions 中构建）
# 第二阶段：运行
FROM alpine:latest

WORKDIR /app

# 根据 TARGETARCH 复制对应二进制文件
ARG TARGETARCH
COPY clipboardshare-linux-$TARGETARCH /app/clipboardshare
COPY README.md index.html ./

# 解决 Alpine 兼容性问题（如需动态库）
RUN apk add --no-cache libc6-compat

EXPOSE 9090
USER nobody:nobody

ENTRYPOINT ["/app/clipboardshare"]