FROM golang:alpine as builder
RUN apk add --no-cache git ca-certificates && update-ca-certificates
ENV UID=10001
RUN adduser \
	--disabled-password \
	--gecos "" \
	--home "/nonexistent" \
	--shell "/sbin/nologin" \
	--no-create-home \
	--uid 10001 \
	d2-prometheus-exporter

WORKDIR /d2-prometheus-exporter
COPY d2-prometheus-exporter.go metrics.go go.mod go.sum  /d2-prometheus-exporter/
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o d2-prometheus-exporter
RUN chown -R d2-prometheus-exporter:d2-prometheus-exporter /d2-prometheus-exporter

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /d2-prometheus-exporter /d2-prometheus-exporter
WORKDIR /d2-prometheus-exporter
USER d2-prometheus-exporter:d2-prometheus-exporter
EXPOSE 3000
ENTRYPOINT ["/d2-prometheus-exporter/d2-prometheus-exporter"]
