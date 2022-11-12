package models

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type IPv4 struct {
	octets [4]uint8
}

func NewIPv4(ipv4Str string) (*IPv4, error) {
	// Syntactic validation
	ipv4Pattern := `^(\d{1,3}\.){3}\d{1,3}$`
	if matched, _ := regexp.MatchString(ipv4Pattern, ipv4Str); !matched {
		return nil, errors.New("invalid IPv4 address")
	}

	// Symantic validation
	octets := [4]uint8{}
	for i, s := range strings.Split(ipv4Str, ".") {
		if n, _ := strconv.Atoi(s); 0 <= n && n <= 255 {
			octets[i] = uint8(n)
		} else {
			return nil, errors.New("invalid IPv4 address")
		}
	}

	return &IPv4{octets}, nil
}

type Netmask struct {
	cidr uint8
	mask [4]uint8
	bits [4][8]uint8
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

	// Get mask
	mask := [4]uint8{}
	for i, octetBits := range bits {
		value := 0
		for power, bit := range octetBits {
			if bit == 1 {
				value += int(math.Pow(2, float64(power)))
			}
		}
		mask[i] = uint8(value)
	}

	return &Netmask{cidr, mask, bits}, nil
}
