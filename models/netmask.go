package models

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Octets [4]uint8

func (octets Octets) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", octets[0], octets[1], octets[2], octets[3])
}

type Netmask struct {
	cidr   uint8
	bits   [4][8]uint8
	octets Octets
}

func NewNetmask(cidrStr string) (*Netmask, error) {
	// Syntactic validation
	cidrPattern := `^\d{1,2}$`
	if matched, _ := regexp.MatchString(cidrPattern, cidrStr); !matched {
		return nil, errors.New("invalid subnet")
	}

	// Symantic validation
	var cidr uint8
	if n, _ := strconv.Atoi(cidrStr); 0 <= n && n <= 32 {
		cidr = uint8(n)
	} else {
		return nil, errors.New("invalid subnet")
	}

	// Get bits
	bits := [4][8]uint8{}
	for i := 0; i < int(cidr); i++ {
		bits[int(i/8)][i%8] = 1
	}

	// Get octets
	octets := Octets{}
	for i, octetBits := range bits {
		value := 0
		for power, bit := range octetBits {
			if bit == 1 {
				value += int(math.Pow(2, float64(power)))
			}
		}
		octets[i] = uint8(value)
	}

	return &Netmask{cidr, bits, octets}, nil
}

func (netmask Netmask) Bits() [4][8]uint8 {
	return netmask.bits
}

func (netmask Netmask) Octets() Octets {
	return netmask.octets
}

func (netmask Netmask) Hosts() uint {
	pow := 32 - netmask.cidr
	return uint(math.Pow(2, float64(pow)))
}
