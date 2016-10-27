metablog: main.go handler.go
	go build -o $@ $^

test:
	go test

clean:
	go clean

.PHONY: default
