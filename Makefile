.PHONY: build run test clean docker-*

# Variáveis
APP_NAME=sar_api
DOCKER_IMAGE=$(APP_NAME)

build:
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/api ./cmd/api

run:
    go run ./cmd/api/main.go

test:
    go test -v ./...

clean:
    rm -rf bin/

docker-build:
    docker build -t $(DOCKER_IMAGE) -f docker/api/Dockerfile .

docker-up: docker-build
    docker-compose up -d

docker-down:
    docker-compose down -v

docker-logs:
    docker-compose logs -f

docker-restart: docker-down docker-up

dev: docker-up run

# Comandos para depuração
docker-sh:
    docker exec -it $(APP_NAME) sh

mongo-sh:
    docker exec -it mongodb mongo