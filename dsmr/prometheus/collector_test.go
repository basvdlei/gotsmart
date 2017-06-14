package prometheus

import (
	"testing"
	"time"

	"github.com/basvdlei/gotsmart/dsmr"
	"github.com/prometheus/client_golang/prometheus"
)

var frame = dsmr.Frame{
	Header:      "",
	Version:     "",
	EquipmentID: "",
	Timestamp:   time.Now(),

	Objects: map[string]dsmr.DataObject{
		"1-0:1.8.1": dsmr.DataObject{
			ID:    "1-0:1.8.1",
			Value: "000093.179",
			Unit:  "kWh",
		},
		"0-0:96.14.0": dsmr.DataObject{
			ID:    "0-0:96.14.0",
			Value: "0001",
		},
	},
}

func TestDSMRCollector(t *testing.T) {
	dc := &DSMRCollector{}
	dc.Update(frame)
	ch := make(chan prometheus.Metric)
	go dc.Collect(ch)
	t.Logf("Metric: %s\n", <-ch)
	t.Logf("Metric: %s\n", <-ch)
}
