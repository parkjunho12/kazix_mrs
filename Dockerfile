FROM ubuntu:latest

RUN apt-get update;apt-get install -y ca-certificates;

COPY . /app/

RUN chmod 777 -R /app/

ENV APP_HOME=/app/config/

ENTRYPOINT [ "/app" ]

EXPOSE 2020


