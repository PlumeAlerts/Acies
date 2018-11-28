FROM golang:1.11.2-alpine

RUN apk add --update netcat-openbsd bash && rm -rf /var/cache/apk/*

WORKDIR /

ENV GO111MODULE=on
ENV MYSQL_HOST=mysql
ENV MYSQL_PORT=3306

COPY Acies /Acies
COPY wait-mysql.sh /wait-mysql.sh

ENTRYPOINT "./wait-mysql.sh"