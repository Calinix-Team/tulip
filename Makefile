build:
	go build -v -ldflags="-s -w" -o tulip

run:
	go run main.go

install:
	go build -v -ldflags="-s -w" -o tulip
	install -Dm755 tulip /usr/bin/tulip