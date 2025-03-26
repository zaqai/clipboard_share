FROM alpine:3.15.5
ARG TARGETOS
ARG TARGETARCH
COPY build/clipboardshare_${TARGETOS}_${TARGETARCH} /app/clipboardshare
RUN chmod +x /app/clipboardshare
COPY index.html /app/index.html
CMD ["/app/clipboardshare"]  # 添加默认启动命令