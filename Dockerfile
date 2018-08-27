# Golang version
FROM golang:1.8

# Maintainer
MAINTAINER rafaelrendonpablo@gmail.com

# Get the dependencies
RUN go get github.com/onsi/ginkgo/ginkgo &&  go get github.com/onsi/gomega && go get -u golang.org/x/lint/golint

# Will download our package to $GOPATH/src/github.com/algoristas/api
RUN go get github.com/algoristas/api

# Set working directory
ENV APP_ROOT $GOPATH/src/github.com/algoristas/api
WORKDIR $GOPATH/src/github.com/algoristas/api
COPY . $GOPATH/src/github.com/algoristas/api

# Get dependencies
RUN go get

# Build
RUN env GOOS=linux GOARCH=amd64 go build -o bin/api github.com/algoristas/api

# Configure port
EXPOSE 8080

# Run deploy and test it
CMD ["./bin/deploy.sh"]
