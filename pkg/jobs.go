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

		states := make(map[int8]int)
		states[ipp.JobStatePending] = 0
		states[ipp.JobStateHeld] = 0
		states[ipp.JobStateProcessing] = 0
		states[ipp.JobStateStopped] = 0
		states[ipp.JobStateCanceled] = 0
		states[ipp.JobStateAborted] = 0
		states[ipp.JobStateCompleted] = 0

		for _, attr := range jobs {

			states[int8(attr["job-state"][0].Value.(int))]++
		}

		ch <- prometheus.MustNewConstMetric(e.jobStateTotal, prometheus.GaugeValue, float64(states[ipp.JobStatePending]), printer, "pending")
		ch <- prometheus.MustNewConstMetric(e.jobStateTotal, prometheus.GaugeValue, float64(states[ipp.JobStateHeld]), printer, "held")
		ch <- prometheus.MustNewConstMetric(e.jobStateTotal, prometheus.GaugeValue, float64(states[ipp.JobStateProcessing]), printer, "processing")
		ch <- prometheus.MustNewConstMetric(e.jobStateTotal, prometheus.GaugeValue, float64(states[ipp.JobStateStopped]), printer, "stopped")
		ch <- prometheus.MustNewConstMetric(e.jobStateTotal, prometheus.GaugeValue, float64(states[ipp.JobStateCanceled]), printer, "canceled")
		ch <- prometheus.MustNewConstMetric(e.jobStateTotal, prometheus.GaugeValue, float64(states[ipp.JobStateAborted]), printer, "aborted")
		ch <- prometheus.MustNewConstMetric(e.jobStateTotal, prometheus.GaugeValue, float64(states[ipp.JobStateCompleted]), printer, "completed")

	}

	return nil
}
