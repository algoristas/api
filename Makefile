export APP_ROOT=$(PWD)
build:
	go fmt
	go vet
	go build -o bin/api

run:
	./bin/api
