package dsmr

import "testing"

var (
	frame = `/XMX5LGBBFG1009421637

1-3:0.2.8(42)
0-0:1.0.0(161001135304S)
0-0:96.1.1(4530303331303033323232333733303136)
1-0:1.8.1(000093.179*kWh)
1-0:1.8.2(000056.684*kWh)
1-0:2.8.1(000000.000*kWh)
1-0:2.8.2(000000.000*kWh)
0-0:96.14.0(0001)
1-0:1.7.0(00.372*kW)
1-0:2.7.0(00.000*kW)
0-0:96.7.21(00001)
0-0:96.7.9(00000)
1-0:99.97.0(0)(0-0:96.7.19)
1-0:32.32.0(00000)
1-0:32.36.0(00000)
0-0:96.13.1()
0-0:96.13.0()
1-0:31.7.0(002*A)
1-0:21.7.0(00.372*kW)
1-0:22.7.0(00.000*kW)
!`
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
		t.Errorf("value does not match %q != %q", d.Value, want)
	}

	want = ""
	if d.Unit != want {
		t.Errorf("unit does not match %q != %q", d.Unit, want)
	}
}

func TestParseFrame(t *testing.T) {
	f, err := ParseFrame(frame)
	t.Logf("got %d objects", len(f.Objects))
	if err != nil {
		t.Error(err)
	}

	t.Logf("Frame: %v", f)
}
