/*
Package prometheus implements a collector of DSMR metrics for Prometheus.
*/
package prometheus

import (
	"log"
	"strconv"

	"github.com/basvdlei/gotsmart/dsmr"
	"github.com/prometheus/client_golang/prometheus"
)

// DSMRCollector implements the Prometheus Collector interface.
type DSMRCollector struct {
	metrics []prometheus.Metric
}

// Collect implements part of the prometheus.Collector interface.
func (dc *DSMRCollector) Collect(ch chan<- prometheus.Metric) {
	for _, m := range dc.metrics {
		ch <- m
	}
}

// Describe implements part of the prometheus.Collector interface.
func (dc *DSMRCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, mb := range metricBuilders {
		ch <- mb.Desc
	}
}

// Update all the metrics to the values of the given frame.
func (dc *DSMRCollector) Update(f dsmr.Frame) {
	var metrics []prometheus.Metric
	for _, obj := range f.Objects {
		if mb, found := metricBuilders[obj.ID]; found {
			if !mb.CheckUnit(obj.Unit) {
				log.Printf("unit in object does not meet spec: %s\n", obj)
				continue
			}
			value, err := strconv.ParseFloat(obj.Value, 64)
			if err != nil {
				log.Printf("could not parse value to float64 for %s\n", obj)
				continue
			}
			m, err := prometheus.NewConstMetric(
				mb.Desc,
				mb.ValueType,
				value,
				f.EquipmentID, f.Version, //labels
			)
			if err != nil {
				log.Printf("could not create prometheus metric for %s\n", obj)
				continue
			}
			metrics = append(metrics, m)
		} else {
			continue
		}
	}
	dc.metrics = metrics
}
