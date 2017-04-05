export APP_ROOT=$(PWD)
build:
	go build -o bin/api

run:
	./bin/api
