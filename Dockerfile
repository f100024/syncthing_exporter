FROM quay.io/prometheus/golang-builder:1.17-base as builder

COPY . /builddir
WORKDIR /builddir

RUN make build


FROM debian:stable-slim

ARG USERNAME=syncthing_exporter_user
ARG USERGROUP=${USERNAME}

RUN addgroup --system --gid 10001 ${USERGROUP}
RUN adduser --system --uid 10000 --gid 10001 --home /home/${USERNAME} ${USERNAME}

COPY --from=builder /builddir/syncthing_exporter /usr/bin/syncthing_exporter

WORKDIR /home/${USERNAME}
RUN chown -R ${USERNAME}:${USERNAME} /usr/bin/syncthing_exporter

USER ${USERNAME}

EXPOSE 9093
ENTRYPOINT ["syncthing_exporter"] 

