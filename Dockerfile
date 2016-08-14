FROM resin/rpi-raspbian
MAINTAINER niklas.wik@iki.fi

RUN mkdir -p /opt/weatherServer/templates

ADD weatherServer /opt/weatherServer
ADD templates/*.tmpl /opt/weatherServer/templates/
