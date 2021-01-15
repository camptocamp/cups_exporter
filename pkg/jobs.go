package pkg

import (
	"github.com/phin1x/go-ipp"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
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

		states := map[int8]int{}

		for _, attr := range jobs {

			if len(attr["job-state"]) == 1 {

				value := int8(attr["job-state"][0].Value.(int))

				if value <= 9 && value >= 3 {
					states[value]++
				} else {
					e.log.Info("Unknow job state : " + strconv.Itoa(int(value)))
				}
			} else {
				e.log.Info("job state attribute missing")
			}
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
