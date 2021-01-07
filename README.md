# Prometheus CUPS Exporter

Prometheus exporter for CUPS server, forked from ["phin1x/cups_exporter"](https://github.com/phin1x/cups_exporter)

# Running

Download the latest release and run the exporter locally :

Examples:
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


# Metrics

| Metric | Meaning | Labels |
| ------ | ------- | ------ |
| cups_up | Was the last scrape of cups successful | |
| cups_job_state_total | Number of current print jobs per state | state, printer |
| cups_job_total | Total number of print jobs per printer | printer |
| cups_printer_state_total | Number of printers per state | state, printer |
| cups_printer_total | Total number of available printers | |
| cups_scrape_duration_seconds |  Duration of the last scrape in seconds | |


Examples:
```
# HELP cups_job_state_total Number of jobs per state
# TYPE cups_job_state_total gauge
cups_job_state_total{printer="CUPS_Printer_1",state="aborted"} 0
cups_job_state_total{printer="CUPS_Printer_1",state="canceled"} 0
cups_job_state_total{printer="CUPS_Printer_1",state="completed"} 2
cups_job_state_total{printer="CUPS_Printer_1",state="held"} 0
cups_job_state_total{printer="CUPS_Printer_1",state="pending"} 0
cups_job_state_total{printer="CUPS_Printer_1",state="processing"} 1
cups_job_state_total{printer="CUPS_Printer_1",state="stopped"} 0
# HELP cups_job_total Total number of print jobs
# TYPE cups_job_total counter
cups_job_total{printer="CUPS_Printer_1"} 3
# TYPE cups_printer_state_total gauge
cups_printer_state_total{printer="CUPS_Printer_1",state="idle"} 1
cups_printer_state_total{printer="CUPS_Printer_1",state="processing"} 0
cups_printer_state_total{printer="CUPS_Printer_1",state="stopped"} 0
# TYPE cups_printer_total gauge
cups_printer_total 2
```

# Licence

Apache Licence Version 2.0
