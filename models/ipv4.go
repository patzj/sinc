package models

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/patzj/network-calculator/utils"
)

type IPv4 struct {
	octets [4]uint8
	bits   [4][8]uint8
}

func NewIPv4(ipv4Str string) (*IPv4, error) {
	// Syntactic validation
	ipv4Pattern := `^(\d{1,3}\.){3}\d{1,3}$`
	if matched, _ := regexp.MatchString(ipv4Pattern, ipv4Str); !matched {
		return nil, errors.New("invalid IPv4 address")
	}

	// Symantic validation
	octets := [4]uint8{}
	bits := [4][8]uint8{}

	for i, s := range strings.Split(ipv4Str, ".") {
		if n, _ := strconv.Atoi(s); 0 <= n && n <= 255 {
			octets[i] = uint8(n)
			bits[i] = utils.OctetBits(octets[i])
		} else {
			return nil, errors.New("invalid IPv4 address")
		}
	}

	return &IPv4{octets, bits}, nil
}

func (ipv4 IPv4) Bits() [4][8]uint8 {
	return ipv4.bits
}

func (ipv4 IPv4) Octets() [4]uint8 {
	return ipv4.octets
}
