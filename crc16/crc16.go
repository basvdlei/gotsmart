// Package crc16 implements CRC-16-IBM/CRC-16-ANSI reversed.
package crc16

// polynomial, reversed: x16 + x15 + x2 + 1
const polynomial = 0xA001

// Update returns a new checksum of crc with additional byte b.
func Update(crc uint16, b byte) uint16 {
	crc = crc ^ uint16(b)
	for i := 0; i < 8; i++ {
		if (crc & 0x0001) > 0 {
			crc = (crc >> 1) ^ polynomial
		} else {
			crc = (crc >> 1)
		}
	}
	return crc
}

// Checksum returns the CRC16 checksum of data.
func Checksum(data []byte) uint16 {
	c := uint16(0x0000)
	for _, b := range data {
		c = Update(c, b)
	}
	return c
}
