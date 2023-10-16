run:
	cd build && docker-compose up

test:
	go test ./...

lint: tools
	bin/golangci-lint run

.PHONY: tools
tools:
	go generate tools/tools.go