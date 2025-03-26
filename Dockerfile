FROM alpine:3.15.5
ARG TARGETOS
ARG TARGETARCH
COPY build/clipboardshare_${TARGETOS}_${TARGETARCH} /usr/local/bin/clipboardshare
RUN chmod +x /usr/local/bin/clipboardshare
COPY --from=builder /clipboardshare /app/clipboardshare
# 复制配置文件等（确保index.html存在于构建上下文）
COPY index.html ./  