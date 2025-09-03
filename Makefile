.PHONY: build run docker-build docker-run clean

build:
	go build -o bin/report-service ./cmd/server

run: build
	./bin/report-service

generate:
	./generate.sh

docker-build:
	docker build -t report-service .

docker-run:
	docker run -p 50051:50051 report-service

clean:
	rm -rf bin/ gen/