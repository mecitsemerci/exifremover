build:
	go build -o exifremover .

run: build
	./exifremover


docker-run: docker-build
	docker compose up --build

.PHONY: build run test docker-build docker-run
