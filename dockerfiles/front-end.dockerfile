FROM alpine:latest
RUN mkdir /app

COPY front-end/frontendLinux /app

CMD [ "/app/frontendLinux" ]
