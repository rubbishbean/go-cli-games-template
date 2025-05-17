run:
	go run main.go play

tidy:
	go mod tidy

build:
	go build -o go-cli-games main.go