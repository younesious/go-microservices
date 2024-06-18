FROM alpine:latest

RUN mkdir /app

COPY ../broker-service/brokerApp /app

CMD ["/app/brokerApp"]
