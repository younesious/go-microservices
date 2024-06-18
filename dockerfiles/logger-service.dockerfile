FROM alpine:latest

RUN mkdir /app

copy ../logger-service/loggerApp /app

CMD ["/app/loggerApp"]
