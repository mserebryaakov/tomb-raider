DOCKER_IMAGE=tomb-raider

run:
	go run ./cmd/main.go

docker-build:
	docker build -t DOCKER_IMAGE .