APP_NAME=simple-rest-api
APP_SRC=$(wildcard *.go)


build:
	go build -o $(APP_NAME) $(APP_SRC)

run: build
	./$(APP_NAME)

install:
	go mod download

clean:
	rm -f $(APP_NAME)

.PHONY: build run install clean
