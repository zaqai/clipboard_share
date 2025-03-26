FROM alpine:3.15.5
WORKDIR /app
ARG TARGETOS
ARG TARGETARCH
COPY build/clipboardshare_${TARGETOS}_${TARGETARCH} ./clipboardshare
RUN chmod +x ./clipboardshare
COPY index.html ./
CMD ["./clipboardshare"]  