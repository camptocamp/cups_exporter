package pkg

import (
	"github.com/go-logr/logr"
	"github.com/phin1x/go-ipp"
	"github.com/prometheus/client_golang/prometheus"
	"net/url"
	"strconv"
	"time"
)

const (
	namespace = "cups"
)

func NewExporter(cupsUri string, log logr.Logger) (*Exporter, error) {
	cu, err := url.Parse(cupsUri)
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(cu.Port())
	if err != nil {
		return nil, err
	}

	password, _ := cu.User.Password()

	client := ipp.NewCUPSClient(cu.Hostname(), port, cu.User.Username(), password, cu.Scheme == "https")

	return &Exporter{
		log:    log,
		client: client,

		lastJobId:          0,
		lastCompletedJobId: 0,

		cupsUp: prometheus.NewDesc(
			"cups_up",
			"Was the last scrape of cups successful",
			nil,
			nil,
		),
		scrapeDurationSeconds: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "scrape_duration_seconds"),
			"Duration of the last scrape in seconds",
			nil,
			nil,
		),
		jobsTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "job", "total"),
			"Total number of print jobs",
			[]string{"printer"},
			nil,
		),
		printersTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "printer", "total"),
			"Number of available printers",
			nil,
			nil,
		),
		printerStateTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "printer", "state_total"),
			"Number of printers per state",
			[]string{"printer", "state"},
			nil,
		),
	}, nil
}

type Exporter struct {
	log    logr.Logger
	client *ipp.CUPSClient

	lastJobId          int
	lastCompletedJobId int

	cupsUp                *prometheus.Desc
	scrapeDurationSeconds *prometheus.Desc
	jobsTotal             *prometheus.Desc
	printersTotal         *prometheus.Desc
	printerStateTotal     *prometheus.Desc
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.cupsUp
	ch <- e.scrapeDurationSeconds
	ch <- e.jobsTotal
	ch <- e.printersTotal
	ch <- e.printerStateTotal
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	cupsUp := float64(1)
	scrapeStartTime := time.Now()

	if err := e.jobsMetrics(ch); err != nil {
		e.log.Error(err, "failed to get job metrics")
		cupsUp = 0
	}

	if err := e.printerMetrics(ch); err != nil {
		e.log.Error(err, "failed to get printer metrics")
		cupsUp = 0
	}

	scrapeDuration := time.Since(scrapeStartTime)

	ch <- prometheus.MustNewConstMetric(e.scrapeDurationSeconds, prometheus.GaugeValue, scrapeDuration.Seconds())
	ch <- prometheus.MustNewConstMetric(e.cupsUp, prometheus.GaugeValue, cupsUp)
}
