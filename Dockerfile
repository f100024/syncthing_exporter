FROM golang:1.24 as builder
COPY . /builddir
WORKDIR /builddir

RUN make build

FROM alpine:3.21 as local
COPY --from=builder /builddir/syncthing_exporter /usr/bin/syncthing_exporter

EXPOSE 9093
ENTRYPOINT ["syncthing_exporter"] 

FROM alpine:3.21 as ghactions
ARG TARGETOS TARGETARCH TARGETVARIANT
COPY .build/${TARGETOS}-${TARGETARCH}${TARGETVARIANT}/syncthing_exporter /usr/bin/syncthing_exporter

EXPOSE 9093
ENTRYPOINT ["syncthing_exporter"]
