FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
LOGGER_BINARY=loggerApp
MAIL_BINARY=mailerApp
LISTENER_BINARY=listenerApp

FRONT_END_VERSION=1.0.0
AUTH_VERSION=1.0.0
BROKER_VERSION=1.0.0
LISTENER_VERSION=1.0.0
MAIL_VERSION=1.0.0
LOGGER_VERSION=1.0.0

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_logger build_mail build_listener
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## build_dockerfiles: builds all dockerfile images
build_dockerfiles: build_auth build_broker build_listener build_logger build_mail front_end_linux
	@echo "Building dockerfiles..."
	docker build -f dockerfiles/front-end.dockerfile -t younesious/front-end-service:${FRONT_END_VERSION} .
	docker build -f dockerfiles/authentication-service.dockerfile -t younesious/authentication-service:${AUTH_VERSION} .
	docker build -f dockerfiles/broker-service.dockerfile -t younesious/broker-service:${BROKER_VERSION} .
	docker build -f dockerfiles/listener-service.dockerfile -t younesious/listener-service:${LISTENER_VERSION} .
	docker build -f dockerfiles/mail-service.dockerfile -t younesious/mail-service:${MAIL_VERSION} .
	docker build -f dockerfiles/logger-service.dockerfile -t younesious/logger-service:${LOGGER_VERSION} .

## push_dockerfiles: pushes tagged versions to docker hub
push_dockerfiles: build_dockerfiles
	docker push younesious/front-end-service:${FRONT_END_VERSION}
	docker push younesious/authentication-service:${AUTH_VERSION}
	docker push younesious/broker-service:${BROKER_VERSION}
	docker push younesious/listener-service:${LISTENER_VERSION}
	docker push younesious/mail-service:${MAIL_VERSION}
	docker push younesious/logger-service:${LOGGER_VERSION}
	@echo "Done!"

## front_end_linux: builds linux executable for front end
front_end_linux:
	@echo "Building linux version of front end..."
	cd front-end && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o frontendLinux ./cmd/web
	@echo "Done!"

## build_auth: builds the authentication binary as a linux executable
build_auth:
	@echo "Building authentication binary.."
	cd authentication-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Authentication binary built!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd logger-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Logger binary built!"

## build_mail: builds the mail binary as a linux executable
build_mail:
	@echo "Building mailer binary..."
	cd mail-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/api
	@echo "Mailer binary built!"

## build_listener: builds the listener binary as a linux executable
build_listener:
	@echo "Building listener binary..."
	cd listener-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${LISTENER_BINARY} .
	@echo "Listener binary built!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"

## test: runs all tests
test:
	@echo "Testing..."
	go test -v ./...

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@cd broker-service && rm -f ${BROKER_BINARY}
	@cd broker-service && go clean
	@cd listener-service && rm -f ${LISTENER_BINARY}
	@cd listener-service && go clean
	@cd authentication-service && rm -f ${AUTH_BINARY}
	@cd authentication-service && go clean
	@cd mail-service && rm -f ${MAIL_BINARY}
	@cd mail-service && go clean
	@cd logger-service && rm -f ${LOGGER_BINARY}
	@cd logger-service && go clean
	@cd front-end && go clean
	@cd front-end && rm -f ${FRONT_END_BINARY}
	@echo "Cleaned!"

## help: displays help
help: Makefile
	@echo " Choose a command:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
