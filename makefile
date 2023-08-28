run:
	go run main.go

build:
	export GOOS=linux 
	export GOARCH=amd64
	mkdir -p ./bin
	go build -o ./bin/exporter
