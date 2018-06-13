/*
Package dsmr implements DSMR P1 frame reading.

The DSMR 4 specification can be found at: http://www.netbeheernederland.nl/

A frame is formatted as:

  / X X X 5 Identification CR LF CR LF Data ! CRC CR LF

*/
package dsmr

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// DateTimeFormat used in a frame, YYMMDDhhmmssX localtime with last as S/W for
// Summer/Winter
const DateTimeFormat = "060102150405"

var (
	// Regexp that matches most of the objects groups with 2 groups:
	//  - OBIS Reduced ID-code eg `1-0:1.8.1`
	//  - Value eg `(000084.276*kWh)`
	objectRegexp = regexp.MustCompile("([0-9]+-[0-9]+:[0-9]+.[0-9]+.[0-9]+)\\((.*)\\)")
	// defaultValueRegexp extract value and unit with 2 groups:
	//  - Value eg `000084.276`
	//  - Unit (optional) eg `kWh`
	defaultValueRegexp = regexp.MustCompile("([^*]+)\\*?(.*)")
)

// Frame represents a DSMR4 frame as send from a P1 port.
type Frame struct {
	Header      string
	Version     string
	EquipmentID string
	Timestamp   time.Time

	Objects map[string]DataObject
}

// DataObject represents a line in the DSMR frame.
type DataObject struct {
	ID    string
	Value string
	Unit  string
}

func (do DataObject) String() string {
	if do.Unit == "" {
		return fmt.Sprintf("%s(%s)", do.ID, do.Value)
	}
	return fmt.Sprintf("%s(%s*%s)", do.ID, do.Value, do.Unit)
}

// ParseFrame returns a frame from the text respresentation.
func ParseFrame(frame string) (f Frame, err error) {
	f.Objects = make(map[string]DataObject)

	for _, s := range strings.Split(frame, "\n") {
		s = strings.TrimSpace(s)

		// skip lines without objects
		if s == "" || s[0] == '!' {
			continue
		} else if s[0] == '/' {
			f.Header = s
			continue
		}

		obj, err := ParseObject(s)
		if err != nil {
			continue
		}

		switch obj.ID {
		// Version of P1 output
		case "1-3:0.2.8":
			f.Version = obj.Value
		// Date-Time of P1 output
		case "0-0:1.0.0":
			if len(obj.Value) > 2 {
				// Remove S/W from timestamp
				timestamp := obj.Value[:len(obj.Value)-1]
				//daylight := obj.Value[len(obj.Value)-1]
				loc, err := time.LoadLocation("Europe/Amsterdam")
				if err != nil {
					continue
				}
				t, err := time.ParseInLocation(DateTimeFormat, timestamp, loc)
				if err != nil {
					continue
				}
				f.Timestamp = t
			}
		case "0-0:96.1.1":
			f.EquipmentID = obj.Value
		default:
			f.Objects[obj.ID] = obj
		}
	}
	return f, nil
}

// ParseObject returns a object for a given line in a frame.
func ParseObject(line string) (DataObject, error) {
	m := objectRegexp.FindStringSubmatch(strings.TrimSpace(line))
	if m == nil || len(m) < 3 {
		return DataObject{}, fmt.Errorf("no object found in string")
	}

	id := m[1]
	rawValue := m[2]

	m = defaultValueRegexp.FindStringSubmatch(rawValue)
	if m == nil {
		return DataObject{
			ID:    id,
			Value: rawValue,
		}, nil
	}
	if len(m) > 2 {
		return DataObject{
			ID:    id,
			Value: m[1],
			Unit:  m[2],
		}, nil
	}
	return DataObject{
		ID:    id,
		Value: m[1],
	}, nil
}
