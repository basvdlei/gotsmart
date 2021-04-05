package dsmr

import "testing"

var (
	frame = `/XMX5LGBBFG1009421637

1-3:0.2.8(50)
0-0:1.0.0(161001135304S)
0-0:96.1.1(4530303331303033323232333733303136)
1-0:1.8.1(000093.179*kWh)
1-0:1.8.2(000056.684*kWh)
1-0:2.8.1(000000.000*kWh)
1-0:2.8.2(000000.000*kWh)
0-0:96.14.0(0001)
1-0:1.7.0(00.372*kW)
1-0:2.7.0(00.000*kW)
0-0:96.7.21(00019)
0-0:96.7.9(00002)
1-0:99.97.0(2)(0-0:96.7.19)(190122122823W)(0000017050*s)(190129101540W)(0000000809*s)
1-0:32.32.0(00006)
1-0:52.32.0(00007)
1-0:72.32.0(00005)
1-0:32.36.0(00000)
1-0:52.36.0(00000)
1-0:72.36.0(00000)
0-0:96.13.0()
1-0:32.7.0(233.0*V)
1-0:52.7.0(233.0*V)
1-0:72.7.0(235.0*V)
1-0:31.7.0(000*A)
1-0:51.7.0(000*A)
1-0:71.7.0(000*A)
1-0:21.7.0(00.129*kW)
1-0:41.7.0(00.000*kW)
1-0:61.7.0(00.067*kW)
1-0:22.7.0(00.000*kW)
1-0:42.7.0(00.060*kW)
1-0:62.7.0(00.000*kW)
0-1:24.1.0(003)
0-1:96.1.0(4730303331303033323232333733303136)
0-1:24.2.1(210405101508S)(01024.212*m3)`
)

func TestParseObject(t *testing.T) {
	do := "1-0:1.8.1(000084.276*kWh)"
	d, err := ParseObject(do)
	if err != nil {
		t.Error(err)
	}
	t.Logf("DO: %v", d)

	want := "000084.276"
	if d.Value != want {
		t.Errorf("value does not match %q != %q", d.Value, want)
	}

	want = "kWh"
	if d.Unit != want {
		t.Errorf("unit does not match %q != %q", d.Unit, want)
	}

	//Double last OBIS digit
	do = "0-0:96.7.21(00001)"
	d, err = ParseObject(do)
	if err != nil {
		t.Error(err)
	}
	t.Logf("DO: %v", d)

	want = "00001"
	if d.Value != want {
		t.Errorf("value does not match: %q != %q", d.Value, want)
	}

	want = ""
	if d.Unit != want {
		t.Errorf("unit does not match: %q != %q", d.Unit, want)
	}

	do = "0-1:24.2.1(210405101508S)(01024.212*m3)"
	d, err = ParseObject(do)
	if err != nil {
		t.Error(err)
	}
	t.Logf("DO: %v", d)

	want = "01024.212"
	if d.Value != want {
		t.Errorf("value does not match: %q != %q", d.Value, want)
	}

	want = "m3"
	if d.Unit != want {
		t.Errorf("unit does not match: %q != %q", d.Unit, want)
	}

	t.Log(d.Unit)
}

func TestParseFrame(t *testing.T) {
	f, err := ParseFrame(frame)
	t.Logf("got %d objects", len(f.Objects))
	if err != nil {
		t.Error(err)
	}

	t.Logf("Frame: %v", f)

	want := "4530303331303033323232333733303136"

	if f.EquipmentID != want {
		t.Errorf("EquipmentID does not match: %q != %q", f.EquipmentID, want)
	}

	assertions := make(map[string]string)
	assertions["1-0:1.8.1"] = "000093.179"
	assertions["1-0:1.8.2"] = "000056.684"
	assertions["0-1:24.2.1"] = "01024.212"

	for id, want := range assertions {
		dataObject, exists := f.Objects[id]

		if !exists {
			t.Errorf("Id %s not set in frame", id)
		}

		value := dataObject.Value

		if value != want {
			t.Errorf("Value for %s does not match: %q != %q", id, value, want)
		}
	}
}
