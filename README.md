# Prometheus CUPS Exporter

Prometheus exporter for CUPS server, forked from ["phin1x/cups_exporter"](https://github.com/phin1x/cups_exporter)

# Running

### Download the latest release and run the exporter locally :

```bash
$ ./cups_exporter
INFO	cups_exporter/main.go:30	starting cups exporter
INFO	cups_exporter/main.go:49	listening on :9628
```

The following arguments are supported :
```bash
$ ./cups_exporter -h
Usage of ./cups_exporter:
  -cups.uri string
    	uri under which the cups server is available, including username and password it required (default "https://localhost:631")
  -web.listen-address string
    	address on which to expose metrics and web interface (default ":9628")
  -web.telemetry-path string
    	path under which to expose metrics (default "/metrics")
```

### With docker :

```bash
docker run -e CUPS_URI=https://cups.my ghcr.io/camptocamp/cups_exporter:0.0.8
```

Or if you want to test on your machine :

```bash
$ docker run --rm --network="host" ghcr.io/camptocamp/cups_exporter:0.0.8
```

# Metrics

| Metric | Meaning | Labels |
| ------ | ------- | ------ |
| cups_up | Was the last scrape of cups successful | |
| cups_job_total | Total number of print jobs per printer | printer |
| cups_printer_state_total | Number of printers per state | printer, state |
| cups_printer_total | Total number of available printers | |
| cups_scrape_duration_seconds |  Duration of the last scrape in seconds | |


Examples:
```
# HELP cups_job_total Total number of print jobs
# TYPE cups_job_total counter
cups_job_total{printer="CUPS_Printer_1"} 3
# HELP cups_printer_state_total Number of printers per state
# TYPE cups_printer_state_total gauge
cups_printer_state_total{printer="CUPS_Printer_1",state="idle"} 1
cups_printer_state_total{printer="CUPS_Printer_1",state="processing"} 0
cups_printer_state_total{printer="CUPS_Printer_1",state="stopped"} 0
# HELP cups_printer_total Number of available printers
# TYPE cups_printer_total gauge
cups_printer_total 2
# HELP cups_scrape_duration_seconds Duration of the last scrape in seconds
# TYPE cups_scrape_duration_seconds gauge
cups_scrape_duration_seconds 0.036579092
# HELP cups_up Was the last scrape of cups successful
# TYPE cups_up gauge
cups_up 1
```

# Licence

Apache Licence Version 2.0
