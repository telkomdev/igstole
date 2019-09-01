.PHONY : test build clean format

build:
	go build github.com/telkomdev/igstole/cmd/igstole

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

clean:
	rm igstole