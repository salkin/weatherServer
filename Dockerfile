FROM resin/rpi-raspbian
MAINTAINER niklas.wik@iki.fi

RUN mkdir -p /opt/weatherServer/templates

COPY bin/weatherServer /opt/weatherServer
RUN chmod +x /opt/weatherServer/weatherServer
COPY templates/*.tmpl /opt/weatherServer/templates/
COPY config.yaml /opt/weatherServer

ENTRYPOINT ["/opt/weatherServer/weatherServer"]
