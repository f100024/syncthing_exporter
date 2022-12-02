FROM golang:1.18 as builder

COPY . /builddir
WORKDIR /builddir

RUN make build

FROM alpine:3.17
COPY --from=builder /builddir/syncthing_exporter /usr/bin/syncthing_exporter

EXPOSE 9093
ENTRYPOINT ["syncthing_exporter"] 
