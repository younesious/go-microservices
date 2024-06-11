FROM alpine:latest

RUN mkdir /app

copy loggerApp /app

CMD ["/app/loggerApp"]
