FROM golang:1.19 as builder
COPY . /builddir
WORKDIR /builddir

RUN make build

FROM alpine:3.17 as local
COPY --from=builder /builddir/syncthing_exporter /usr/bin/syncthing_exporter

EXPOSE 9093
ENTRYPOINT ["syncthing_exporter"] 


FROM alpine:3.17 as ghactions
ARG LOCAL_FOLDER
COPY .build/${LOCAL_FOLDER}/syncthing_exporter /usr/bin/syncthing_exporter

EXPOSE 9093
ENTRYPOINT ["syncthing_exporter"]
