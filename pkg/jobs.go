package pkg

import (
	"github.com/phin1x/go-ipp"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *Exporter) jobsMetrics(ch chan<- prometheus.Metric) error {

	printers, err := e.client.GetPrinters([]string{"printer-state"})

	if err != nil {
		e.log.Error(err, "failed to fetch printers")
		return err
	}

	for _, attr := range printers {

		printer := attr["printer-name"][0].Value.(string)

		jobs, err := e.client.GetJobs(printer, "", ipp.JobStateFilterAll, false, 0, 0, []string{"job-state"})
		if err != nil {
			e.log.Error(err, "failed to fetch all jobs states")
			return err
		}

		ch <- prometheus.MustNewConstMetric(e.jobsTotal, prometheus.CounterValue, float64(len(jobs)), printer)
	}

	return nil
}
