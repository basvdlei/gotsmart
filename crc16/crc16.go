package crc16

const (
	ANSI_REFL = 0xA001
)

func Update(crc uint16, b byte) uint16 {
	crc = crc ^ uint16(b)
	for i := 0; i < 8; i++ {
		if (crc & 0x0001) > 0 {
			crc = (crc >> 1) ^ ANSI_REFL
		} else {
			crc = (crc >> 1)
		}
	}
	return crc
}

func Checksum(data []byte) uint16 {
	c := uint16(0x0000)
	for _, b := range data {
		c = Update(c, b)
	}
	return c
}
