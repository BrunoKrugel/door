all:
	go build ./cmd/door
	go run ./cmd/door

clean:
	go mod tidy
