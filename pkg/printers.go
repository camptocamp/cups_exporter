package pkg

import (
	"github.com/phin1x/go-ipp"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *Exporter) printerMetrics(ch chan<- prometheus.Metric) error {
	printers, err := e.client.GetPrinters([]string{"printer-state"})
	if err != nil {
		e.log.Error(err, "failed to fetch printers")
		return err
	}

	ch <- prometheus.MustNewConstMetric(e.printersTotal, prometheus.GaugeValue, float64(len(printers)))

	for _, attr := range printers {

		printer := attr["printer-name"][0].Value.(string)

		states := make(map[int8]int)
		states[ipp.PrinterStateIdle] = 0
		states[ipp.PrinterStateProcessing] = 0
		states[ipp.PrinterStateStopped] = 0

		states[int8(attr["printer-state"][0].Value.(int))]++

		ch <- prometheus.MustNewConstMetric(e.printerStateTotal, prometheus.GaugeValue, float64(states[ipp.PrinterStateIdle]), printer, "idle")
		ch <- prometheus.MustNewConstMetric(e.printerStateTotal, prometheus.GaugeValue, float64(states[ipp.PrinterStateProcessing]), printer, "processing")
		ch <- prometheus.MustNewConstMetric(e.printerStateTotal, prometheus.GaugeValue, float64(states[ipp.PrinterStateStopped]), printer, "stopped")

	}

	return nil
}
