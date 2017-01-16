FROM golang:1.7

RUN mkdir -p /opt/weatherServer/templates
COPY weatherServer /opt/weatherServer
COPY config.yaml /opt/weatherServer
COPY templates /opt/weatherServer/templates
ENTRYPOINT ["/opt/weatherServer/weatherServer"]

