# Dockerfile for Syncthing Exporter
FROM golang:1.14

WORKDIR /app

ENV WEB_LISTEN_PORT="9112"
ENV WEB_LISTEN_ADDRESS=":${PORT}"
ENV WEB_METRIC_PATH="/metrics"
ENV ST_URI="http://127.0.0.1:8384"
ENV ST_TOKEN="123"
ENV ST_TIMEOUT="5s"

RUN git clone https://github.com/adenau/syncthing_exporter.git

WORKDIR /app/syncthing_exporter

RUN go build .

CMD ./syncthing_exporter
EXPOSE ${PORT}
