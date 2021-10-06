run:
	go build --ldflags "-X main.envPath=./test.env" -o bin/main main.go && ./bin/main

test:
	go test ./...