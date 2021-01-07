FROM golang:1.13 as builder

WORKDIR /build
ADD . /build
RUN go build -o cups_exporter main.go

FROM golang:1.13

LABEL maintainer="lenny.consuegra@camptocamp.com"

ENV CUPS_URI "https://localhost:631"

COPY --from=builder /build/cups_exporter /bin/cups_exporter

EXPOSE 9329

CMD [ "/bin/sh", "-c", "/bin/cups_exporter -cups.uri ${CUPS_URI}" ]
