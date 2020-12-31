FROM golang:1.15 as builder

# Setup project directory
WORKDIR /go/src/syncthing_exporter
COPY . .

# Build
RUN make build


FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/syncthing_exporter/syncthing_exporter /usr/local/bin

# Create user so the exporter does not run as root.
RUN addgroup -S exporter && adduser -S exporter -G exporter
USER exporter

ENTRYPOINT ["syncthing_exporter"]
