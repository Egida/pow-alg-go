run:
	cd build && docker-compose up

test:
	go test ./...