FROM resin/rpi-raspbian
MAINTAINER niklas.wik@iki.fi

RUN mkdir -p /opt/weatherServer/templates

COPY weatherServer /opt/weatherServer
COPY templates/*.tmpl /opt/weatherServer/templates/
COPY config.yml /opt/weatherServer
