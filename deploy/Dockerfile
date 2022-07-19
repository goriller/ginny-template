FROM alpine:latest

RUN apk update && apk upgrade
RUN apk --no-cache --update add build-base openssh git curl bash tzdata

# Time Zone
RUN apk add --no-cache bash curl ca-certificates tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

ADD ./dist/<appName>-linux-amd64 app
ADD ./configs/pro.yml .
ADD ./deploy/wait-for .

# CMD ["./app", "-conf", "./pro.yml"]