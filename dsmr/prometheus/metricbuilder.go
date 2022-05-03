package prometheus

import "github.com/prometheus/client_golang/prometheus"

const (
	namespace = "gotsmart"
)

var (
	defaultLabels = []string{"device", "version"}
)

// MetricBuilder holds the information needed to create a Prometheus metrics.
type MetricBuilder struct {
	ValueType  prometheus.ValueType
	Desc       *prometheus.Desc
	Unit       string
	MetricFunc func(value float64) (prometheus.Metric, error)
}

func (mb MetricBuilder) String() string {
	return mb.Desc.String()
}

// CheckUnit verifies if the given unit is expected for this object.
func (mb MetricBuilder) CheckUnit(unit string) bool {
	return mb.Unit == unit
}

// MetricBuilders contains builders for all object types in a DSMR frame.
var metricBuilders = map[string]MetricBuilder{
	// The first 3 objects from the spec are parsed as part of the frame:
	/*
		// Version information for P1 output 1-3:0.2.8.255 2 1 Data S2, tag 9
		"1-3:0.2.8": MetricBuilder{
			ValueType: prometheus.UntypedValue,
			Desc: prometheus.NewDesc(
				namespace+"_p1_version",
				"version information of the last P1 output",
				defaultLabels,
				prometheus.Labels{},
			),
		},
		// Date-time stamp of the P1 message 0-0:1.0.0.255 2 8 TST YYMMDDhhmmssX
		"0-0:1.0.0": MetricBuilder{
			ValueType: prometheus.CounterValue,
			Desc: prometheus.NewDesc(
				namespace+"_p1_timestamp",
				"date-time stamp of the last P1 message",
				defaultLabels,
				prometheus.Labels{},
			),
		},
		// Equipment identifier 0-0:96.1.1.255 2 Value 1 Data Sn (n=0..96), tag 9
		"0-0:96.1.1": MetricBuilder{
			ValueType: prometheus.UntypedValue,
			Desc: prometheus.NewDesc(
				namespace+"_equipment_identifier",
				"equipment identifier",
				defaultLabels,
				prometheus.Labels{},
			),
		},
	*/
	// Meter Reading electricity delivered to client (Tariff 1) in 0,001
	// kWh 1-0:1.8.1.255 2 Value 3 Register F9(3,3), tag 6 kWh
	"1-0:1.8.1": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_electricity_delivered_to_client_tariff_1_kwh",
			"meter reading electricity delivered to client (tariff 1) in 0,001 kwh",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kWh",
	},
	// Meter Reading electricity delivered to client (Tariff 2) in 0,001
	// kWh 1-0:1.8.2.255 2 Value 3 Register F9(3,3), tag 6 kWh
	"1-0:1.8.2": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_electricity_delivered_to_client_tariff_2_kwh",
			"meter reading electricity delivered to client (tariff 2) in 0,001 kwh",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kWh",
	},
	// Meter Reading electricity delivered by client (Tariff 1) in 0,001
	// kWh 1-0:2.8.1.255 2 Value 3 Register F9(3,3), tag 6 kWh
	"1-0:2.8.1": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_electricity_delivered_by_client_tariff_1_kwh",
			"meter reading electricity delivered by client (tariff 1) in 0,001 kwh",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kWh",
	},
	// Meter Reading electricity delivered by client (Tariff 2) in 0,001
	// kWh 1-0:2.8.2.255 2 Value 3 Register F9(3,3), tag 6 kWh
	"1-0:2.8.2": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_electricity_delivered_by_client_tariff_2_kwh",
			"meter reading electricity delivered by client (tariff 2) in 0,001 kwh",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kWh",
	},
	// Tariff indicator electricity.  The tariff indicator can also be used
	// to switch tariff dependent loads e.g boilers. This is the
	// responsibility of the P1 user 0-0:96.14.0.255 2 Value 1 Data S4, tag
	// 9
	"0-0:96.14.0": MetricBuilder{
		ValueType: prometheus.UntypedValue,
		Desc: prometheus.NewDesc(
			namespace+"_tariff_indicator_electricity",
			"tariff indicator electricity",
			defaultLabels,
			prometheus.Labels{},
		),
	},
	// Actual electricity power delivered (+P) in 1 Watt resolution
	// 1-0:1.7.0.255 2 Value 3 Register F5(3,3), tag 18 kW
	"1-0:1.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_electricity_power_delivered_kw",
			"actual electricity power delivered (+p) in 1 watt resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kW",
	},
	// Actual electricity power received (-P) in 1 Watt resolution
	// 1-0:2.7.0.255 2 Value 3 Register F5(3,3), tag 18 kW
	"1-0:2.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_electricity_power_received_kw",
			"actual electricity power received (-p) in 1 watt resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kW",
	},
	// The actual threshold Electricity in kW 0-0:17.0.0.255 3 Threshold
	// active 71 Limiter Class F4(1,1), tag 18 kW
	"0-0:17.0.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_threshold_electricity_kw",
			"the actual threshold electricity in kw",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kW",
	},
	// Switch position Electricity (in/out/enabled).  0-0:96.3.10.255 3
	// Control State 70 Disconnector Control I1, tag 22
	"0-0:96.3.10": MetricBuilder{
		ValueType: prometheus.UntypedValue,
		Desc: prometheus.NewDesc(
			namespace+"_switch_position_electricity",
			"switch position electricity (in/out/enabled)",
			defaultLabels,
			prometheus.Labels{},
		),
	},
	// Number of power failures in any phase 0-0:96.7.21.255 2 Value 1 Data
	// F5(0,0), tag 18
	"0-0:96.7.21": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_power_failures_total",
			"number of power failures in any phase",
			defaultLabels,
			prometheus.Labels{},
		),
	},
	// Number of long power failures in any phase 0-0:96.7.9.255 2 Value 1
	// Data F5(0,0), tag 18
	"0-0:96.7.9": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_long_power_failures_total",
			"number of long power failures in any phase",
			defaultLabels,
			prometheus.Labels{},
		),
	},
	// XXX Should implement handling of this special log datatype.
	// Power Failure Event Log (long power failures) 1-0:99.97.0.255 2
	// Buffer 7 Profile Generic TST, F10(0,0) - tag 6 Format applicable for
	// the value within the log (OBIS code 0- 0:96.7.19.255) Timestamp (end
	// of failure) –duration in seconds
	/*
		"1-0:99.97.0": MetricBuilder{
			ValueType: prometheus.CounterValue,
			Desc: prometheus.NewDesc(
				namespace+"_",
				"power failure event log (long power failures)",
				defaultLabels,
				prometheus.Labels{},
			),
		},
	*/
	// Number of voltage sags in phase L1 1-0:32.32.0.255 2 Value 1 Data
	// F5(0,0), tag 18
	"1-0:32.32.0": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_voltage_sags_in_phase_l1_total",
			"number of voltage sags in phase l1",
			defaultLabels,
			prometheus.Labels{},
		),
	},
	// Number of voltage sags in phase L2 (polyphase meters only)
	// 1-0:52.32.0.255 2 Value 1 Data F5(0,0), tag 18
	"1-0:52.32.0": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_voltage_sags_in_phase_l2_total",
			"number of voltage sags in phase l2",
			defaultLabels,
			prometheus.Labels{},
		),
	},
	// Number of voltage sags in phase L3 (polyphase meters only)
	// 1-0:72:32.0.255 2 Value 1 Data F5(0,0), tag 18
	"1-0:72:32.0": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_voltage_sags_in_phase_l3_total",
			"number of voltage sags in phase l3",
			defaultLabels,
			prometheus.Labels{},
		),
	},
	// Number of voltage swells in phase L1 1-0:32.36.0.255 2 Value 1 Data
	// F5(0,0), tag 18
	"1-0:32.36.0": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_voltage_swells_in_phase_l1_total",
			"number of voltage swells in phase l1",
			defaultLabels,
			prometheus.Labels{},
		),
	},
	// Number of voltage swells in phase L2 (polyphase meters only)
	// 1-0:52.36.0.255 2 Value 1 Data F5(0,0), tag 18
	"1-0:52.36.0": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_voltage_swells_in_phase_l2_total",
			"number of voltage swells in phase l2",
			defaultLabels,
			prometheus.Labels{},
		),
	},
	// Number of voltage swells in phase L3 (polyphase meters only)
	// 1-0:72.36.0.255 2 Value 1 Data F5(0,0), tag 18
	"1-0:72.36.0": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_voltage_swells_in_phase_l3_total",
			"number of voltage swells in phase l3",
			defaultLabels,
			prometheus.Labels{},
		),
	},
	// Text message codes: numeric 8 digits 0-0:96.13.1.255 2 Value 1 Data
	// Sn (n=0..16),, tag 9
	// TODO Not implemented

	// Text message max 1024 characters.  0-0:96.13.0.255 2 Value 1 Data Sn
	// (n=0..2048), tag 9
	// TODO Not implemented

	// Device-Type  0-n:24.1.0.255  9 Device type 72 M-Bus client F3(0,0),
	// tag 17
	// TODO Not implemented

	// Instantaneous current L1 in A resolution.  1-0:31.7.0.255  2 Value 3
	// Register F3(0,0), tag 18  A
	"1-0:31.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_current_l1_a",
			"instantaneous current l1 in a resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "A",
	},
	// Instantaneous current L2 in A resolution.  1-0:51.7.0.255  2 Value 3
	// Register F3(0,0), tag 18  A
	"1-0:51.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_current_l2_a",
			"instantaneous current l2 in a resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "A",
	},
	// Instantaneous current L3 in A resolution.  1-0:71.7.0.255  2 Value 3
	// Register F3(0,0), tag 18  A
	"1-0:71.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_current_l3_a",
			"instantaneous current l3 in a resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "A",
	},
	// Instantaneous voltage L1 in V resolution.  1-0:32.7.0.255  2 Value 3
	// Register F4(1,1), tag 18  V
	"1-0:32.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_voltage_l1_v",
			"instantaneous voltage l1 in v resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "V",
	},
	// Instantaneous voltage L2 in V resolution.  1-0:52.7.0.255  2 Value 3
	// Register F4(1,1), tag 18  V
	"1-0:52.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_voltage_l2_v",
			"instantaneous voltage l2 in v resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "V",
	},
	// Instantaneous voltage L3 in V resolution.  1-0:72.7.0.255  2 Value 3
	// Register F4(1,1), tag 18  V
	"1-0:72.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_voltage_l3_v",
			"instantaneous voltage l3 in v resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "V",
	},
	// Instantaneous active power L1 (+P) in W resolution 1-0:21.7.0.255  2
	// Value 3 Register F5(3,3), tag 18  kW
	"1-0:21.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_active_power_delivered_l1_kw",
			"instantaneous active power l1 (+p) in w resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kW",
	},
	// Instantaneous active power L2 (+P) in W resolution 1-0:41.7.0.255  2
	// Value 3 Register F5(3,3), tag 18  kW
	"1-0:41.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_active_power_delivered_l2_kw",
			"instantaneous active power l2 (+p) in w resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kW",
	},
	// Instantaneous active power L3 (+P) in W resolution 1-0:61.7.0.255  2
	// Value 3 Register F5(3,3), tag 18  kW
	"1-0:61.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_active_power_delivered_l3_kw",
			"instantaneous active power l3 (+p) in w resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kW",
	},
	// Instantaneous active power L1 (-P) in W resolution 1-0:22.7.0.255  2
	// Value 3 Register F5(3,3), tag 18  kW
	"1-0:22.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_active_power_received_l1_kw",
			"instantaneous active power l1 (-p) in w resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kW",
	},
	// Instantaneous active power L2 (-P) in W resolution 1-0:42.7.0.255  2
	// Value 3 Register F5(3,3), tag 18  kW
	"1-0:42.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_active_power_received_l2_kw",
			"instantaneous active power l2 (-p) in w resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kW",
	},
	// Instantaneous active power L3 (-P) in W resolution 1-0:62.7.0.255  2
	// Value 3 Register F5(3,3), tag 18  kW
	"1-0:62.7.0": MetricBuilder{
		ValueType: prometheus.GaugeValue,
		Desc: prometheus.NewDesc(
			namespace+"_active_power_received_l3_kw",
			"instantaneous active power l3 (-p) in w resolution",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "kW",
	},

	// Switch position Gas

	"0-1:24.4.0": MetricBuilder{
		ValueType: prometheus.UntypedValue,
		Desc: prometheus.NewDesc(
			namespace+"_gas_switch",
			"gas switch",
			defaultLabels,
			prometheus.Labels{},
		),
	},

	// Reading from natural gas meter (timestamp) (value)

	"0-1:24.2.3": MetricBuilder{
		ValueType: prometheus.CounterValue,
		Desc: prometheus.NewDesc(
			namespace+"_gas_m3",
			"actual gas volume delivered",
			defaultLabels,
			prometheus.Labels{},
		),
		Unit: "m3",
	},

	// TODO The types below are Smart Meter extensions like Gas meter, etc.

	// Device-Type  0-n:24.1.0.255  9 Device type 72 M-Bus client F3(0,0), tag 17

	// Equipment identifier (Gas) 0-n:96.1.0.255  2 Value 1 Data Sn
	// (n=0..96), tag 9

	// Last hourly value (temperature converted), gas delivered to client
	// in m3, including decimal values and capture time  0-n:24.2.1.255 5
	// Capture time 4 Extended Register TST 0-n:24.2.1.255  2 Value 4
	// Extended Register F8(2,2)/F8(3,3), tag 18  (See note 2) m3

	// Valve position Gas (on/off/released).  (See Note 3) 0-n:24.4.0.255
	// 3 Control state 70 Disconnect Control I1, tag 22

	// Equipment identifier (Thermal:  Heat or Cold) 0-n:96.1.0.255  2
	// Value 1 Data Sn (n=0..96), tag 9

	// Last hourly Meter reading Heat or Cold in 0,01 GJ and capture time
	// 0-n:24.2.1.255 5 Capture time 4 Extended Register  TST
	// 0-n:24.2.1.255 2 Value 4 Extended Register Fn(2,2)  (See note 1) GJ

	// Valve position (on/off/released).  (See Note 3) 0-n:24.4.0.255  3
	// Control state 70 Disconnect Control I1, tag 22

	// Device-Type  0-n:24.1.0.255  9 Device type 72 M-Bus client F3(0,0),
	// tag 17

	// Equipment identifier (Water) 0-n:96.1.0.255  2 Value 1 Data Sn
	// (n=0..96), tag 9

	// Last hourly Meter  0-n:24.2.1.255  5  4  TST   reading in 0,001 m3
	// and capture time Capture time Extended Register 0-n:24.2.1.255 2
	// Value 4 Extended Register Fn(3,3)  (See Note 1) m3

	// Valve position (on/off/released).  (See Note 3) 0-n:24.4.0.255  3
	// Control state 70 Disconnect Control I1, tag 22

	// Device-Type  0-n:24.1.0.255  9 Device type 72 M-Bus client F3(0,0),
	// tag 17

	// Equipment identifier   0-n:96.1.0.255  2 Value 1 Data Sn (n=0..96),
	// tag 9

	// Last hourly Meter reading and capture time (e.g. slave E meter)
	// 0-n:24.2.1.255 5 Capture time 4 Extended Register TST 0-n:24.2.1.255
	// 2 Value 4 Extended Register Fn(3,3)  (See Note 1) kWh

	// Valve/Switch position (on/off/released).  (See Note 3)
	// 0-n:24.4.0.255  3 Control state 70 Disconnect Control I1, tag 22
}
