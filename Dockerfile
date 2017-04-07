FROM golang:1.8

MAINTAINER rafaelrendonpablo@gmail.com

RUN mkdir -p /app
WORKDIR /app

ENV APP_ROOT /app

COPY . /app

EXPOSE 8080

CMD ["./bin/api"]
