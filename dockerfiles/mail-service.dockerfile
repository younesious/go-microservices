FROM alpine:latest

RUN mkdir /app
RUN mkdir /templates

COPY ../mail-service/mailerApp /app
COPY ../mail-service/templates /templates

CMD ["/app/mailerApp"]
