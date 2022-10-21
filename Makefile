LOGGER_APP=loggerApp
BROKER_APP=brokerApp
LISTENER_APP=listenerApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_logger build_broker build_listener
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd ./logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_APP} ./cmd/api
	@echo "Done!"

build_broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_APP} ./cmd/api
	@echo "Done!"

build_listener:
	@echo "Building listener binary..."
	cd ./listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_APP} ./cmd/api
	@echo "Done!"

php-shell:
	docker exec -it php sh
