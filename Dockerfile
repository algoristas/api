# Golang version
FROM golang:1.8

MAINTAINER rafaelrendonpablo@gmail.com

# Get the dependencies
RUN go get github.com/onsi/ginkgo/ginkgo &&  go get github.com/onsi/gomega

# Get & Build API
RUN go get github.com/algoristas/api
RUN env GOOS=linux GOARCH=amd64 go build -o bin/api github.com/algoristas/api

# Create application folder
RUN mkdir -p /app

# Set working directory
WORKDIR /app

ENV APP_ROOT /app

COPY . /app

EXPOSE 8080

CMD ["./bin/deploy.sh"]
#CMD ["./bin/api"]
