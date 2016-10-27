metablog: main.go handler.go
	go build -o $@ $^

deps:
	go get

test:
	go test

clean:
	go clean

.PHONY: deps test clean
