FROM ubuntu:latest

USER root

RUN apt-get update;apt-get install -y ca-certificates;

COPY ./dist/ /app/

RUN cd /app/

RUN chmod 777 -R /app/

ENV APP_HOME=/app/config/

ENTRYPOINT [ "/app/kazix" ]

EXPOSE 2020


