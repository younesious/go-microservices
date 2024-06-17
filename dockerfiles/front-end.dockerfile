FROM alpine:latest

RUN mkdir /app

COPY ../front-end/frontApp /app

CMD [ "/app/frontApp" ]
