package utils

func OctetBits(octet uint8) [8]uint8 {
	bits := [8]uint8{}
	pos := cap(bits) - 1

	for pos >= 0 {
		rem := octet % 2
		bits[pos] = rem

		pos--
		octet /= 2
	}

	return bits
}
