FROM alpine:latest

RUN apk update && apk upgrade
RUN apk --no-cache --update add build-base openssh git curl bash tzdata

# Time Zone
RUN apk add --no-cache bash curl ca-certificates tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /data

ENV LOG_PATH=logs/log.log
ENV LOG_LEVEL=info
ENV PRE_STOP="80s"

ADD ./deploy/app.bin app.bin
ADD ./configs/config.yml config.yml

CMD ["/data/app.bin", "-conf", "./config.yml"]

