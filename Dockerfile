
FROM golang:1.8-jessie

ADD aceleratomo /opt/

CMD /opt/aceleratomo

EXPOSE 6969

HEALTHCHECK CMD curl --fail http://localhost:6969/ || exit 1