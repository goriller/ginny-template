FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/<appName>-linux-amd64 app
ADD ./configs/pro.yml .
ADD ./scripts/wait-for .