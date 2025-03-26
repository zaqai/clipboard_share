FROM alpine:3.15.5
ARG TARGETOS
ARG TARGETARCH
COPY build/clipboardshare_${TARGETOS}_${TARGETARCH} /usr/local/bin/clipboardshare
RUN chmod +x /usr/local/bin/clipboardshare