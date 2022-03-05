build:
	go build -v -ldflags="-s -w" -o tulip

run:
	go run main.go