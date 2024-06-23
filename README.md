Go Microservices Project
========================

About
-----

This project was designed with educational purposes for the final project of the IUST software engineering course and aims to cover a wide range of software engineering topics such as microservice architecture, communication between services (API protocols like JSON, RPC, and gRPC), software development platforms such as Kubernetes and Docker Swarm, test writing for the authentication-service in the `tests` branch, and monitoring tools such as Prometheus, Grafana, Jaeger, and Pyroscope should be reviewed.

Project Overview
----------------

Microservices, also known as the microservice architecture, are an architectural style that structures an application as a loosely coupled collection of smaller applications. The microservice architecture allows for the rapid and reliable delivery of large, complex applications. Key features of microservices include:

-   **Maintainable and testable**
-   **Loosely coupled**
-   **Independently deployable**
-   **Organized around business capabilities**
-   **Owned by small teams**

In this project, I develop a number of small, self-contained, loosely coupled microservices that communicate with one another and a simple front-end application using REST API, RPC, gRPC, and AMQP (Advanced Message Queuing Protocol). The microservices we build include the following functionality:

All services are written in Go and will be deployed using Docker Swarm and Kubernetes.

Working with Microservices in Go
--------------------------------

This project consists of several loosely coupled microservices, all written in Go:

-   **front-end-service**: Displays web pages.
-   **broker-service**: An optional single entry point into the microservice cluster. to connect to all services from one place (accepts JSON; sends JSON, makes calls via gRPC, and pushes to RabbitMQ).
-   **authentication-service**: Authenticates users against a Postgres database (accepts JSON).
-   **logger-service**: Logs important events to a MongoDB database (accepts RPC, gRPC, and JSON).
-   **listener-service**: Consumes messages from AMQP (RabbitMQ) and initiates actions based on the payload (sends via RPC).
-   **mail-service**:Converts a JSON payload into a formatted email and sends it with Mailhog.


In addition to the microservices, the included `docker-compose.yml` at the root level of the project starts the following services:

-   **PostgreSQL**: Used by the authentication-service to store user accounts.
-   **MongoDB**: Used by the logger-service to save logs from all services.
-   **Mailhog**: Used as a fake mail server to work with the mail service. (Mailhog dashboard is available on `http://localhost:8025`).


Monitoring Tools
----------------

Access the following monitoring tools at these URLs:

-   **Prometheus**: <http://localhost:9090>
-   **Grafana**: <http://localhost:3000>
-   **Jaeger**: <http://localhost:16686>
-   **Pyroscope**: <http://localhost:4040>

Running the Project
-------------------

Fortunately, I wrote some `Makefile` for this project to make life easier for me and you. So From the root level of the project, execute this command (this assumes that you have GNU Make and a recent version of Docker installed on your machine):

```shell
make up_build
```

If the code has not changed, subsequent runs can just be:

```shell
make up
```

Then start the front end:

```shell
make start
```

Hit the front end with your web browser at <http://localhost:8081>.

To stop everything:

```shell
make stop
make down
```

Make Commands
-------------------

here is the complete list of `Make` commands:

```shell
 Choose a command:
  up                  starts all containers in the background without forcing build
  down                stop docker compose
  up_build            stops docker-compose (if running), builds all projects and starts docker compose
  build_dockerfiles   builds all dockerfile images
  push_dockerfiles    pushes tagged versions to docker hub
  front_end_linux     builds linux executable for front end
  build_auth          builds the authentication binary as a linux executable
  build_broker        builds the broker binary as a linux executable
  build_logger        builds the logger binary as a linux executable
  build_mail          builds the mail binary as a linux executable
  build_listener      builds the listener binary as a linux executable
  build_front         builds the frone end binary
  start               starts the front end
  stop                stop the front end
  test                runs all tests
  clean               runs go clean and deletes binaries
  help                displays help
```

License
-------

This project is licensed under the MIT License - see the [LICENSE](https://github.com/younesious/go-microservices/blob/master/LICENSE) file for details.

### Contributing

Feel free to contribute and I'll be happy to see you :)
