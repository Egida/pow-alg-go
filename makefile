run:
	cd build && docker-compose up

test:
	go test ./...

.PHONY: tools
tools:
	go generate tools/tools.go