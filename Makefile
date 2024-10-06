build:
	go build -o main .

run: build
	./main


docker-run: docker-build
	docker compose up --build

.PHONY: build run test docker-build docker-run
