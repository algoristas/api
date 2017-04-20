FROM golang:1.8

MAINTAINER rafaelrendonpablo@gmail.com

RUN go get github.com/onsi/ginkgo/ginkgo &&  go get github.com/onsi/gomega

RUN mkdir -p /app

WORKDIR /app

ENV APP_ROOT /app

COPY . /app

EXPOSE 8080

#CMD ["./bin/deploy.sh"]
CMD ["./bin/api"]
