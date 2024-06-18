FROM alpine:latest
RUN mkdir /app

COPY ../listener-service/listenerApp /app

CMD [ "/app/listenerApp" ]
