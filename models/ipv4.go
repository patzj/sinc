package models

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/patzj/network-calculator/utils"
)

type IPv4 struct {
	octets Octets
	bits   [4][8]uint8
}

func NewIPv4(ipv4Str string) (*IPv4, error) {
	// Syntactic validation
	ipv4Pattern := `^(\d{1,3}\.){3}\d{1,3}$`
	if matched, _ := regexp.MatchString(ipv4Pattern, ipv4Str); !matched {
		return nil, errors.New("invalid IPv4 address")
	}

	// Symantic validation
	octets := Octets{}
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

func (ipv4 IPv4) Octets() Octets {
	return ipv4.octets
}

func (ipv4 IPv4) IsBefore(other IPv4) bool {
	for i, octet := range ipv4.Octets() {
		if octet > other.octets[i] {
			return false
		}
	}
	return true
}

func (ipv4 IPv4) IsAfter(other IPv4) bool {
	for i, octet := range ipv4.Octets() {
		if octet < other.octets[i] {
			return false
		}
	}
	return true
}

func (ipv4 IPv4) IsEqual(other IPv4) bool {
	for i, octet := range ipv4.Octets() {
		if octet != other.octets[i] {
			return false
		}
	}
	return true
}

func (ipv4 IPv4) Subnet(netmask Netmask) (*IPv4, *IPv4, error) {
	hostOctetsMin := ipv4.Octets()
	hostOctetsMax := ipv4.Octets()

	for i, octet := range netmask.Octets() {
		if octet != 255 {
			hostOctetsMin[i] = 0
			hostOctetsMax[i] = 255
		}
	}

	hostAddrMin, _ := NewIPv4(hostOctetsMin.String())
	hostAddrMax, _ := NewIPv4(hostOctetsMax.String())

	hosts := netmask.Hosts()
	networkAddr := *hostAddrMin

	for {
		if networkAddr.IsAfter(*hostAddrMax) {
			return nil, nil, errors.New("subnet not found")
		}

		broadcastAddr := offset(networkAddr, hosts-1)
		withinLowerBounds := ipv4.IsAfter(networkAddr) || ipv4.IsEqual(networkAddr)
		withinUpperBounds := ipv4.IsBefore(broadcastAddr) || ipv4.IsEqual(broadcastAddr)

		if withinLowerBounds && withinUpperBounds {
			return &networkAddr, &broadcastAddr, nil
		}

		networkAddr = offset(networkAddr, hosts)
	}
}

func offset(ipv4 IPv4, n uint) IPv4 {
	octets := ipv4.Octets()
	previousOctet := uint(octets[3])
	previousOctet += n

	pos := cap(octets) - 1
	for pos >= 0 {
		octets[pos] = uint8(previousOctet % 256)

		var carry uint
		if previousOctet > 255 {
			carry = previousOctet / 256
		}

		if pos > 0 {
			previousOctet = uint(octets[pos-1])
			previousOctet += carry
		}

		pos--
	}

	newIpv4Str := fmt.Sprintf("%d.%d.%d.%d", octets[0], octets[1], octets[2], octets[3])
	newIpv4, _ := NewIPv4(newIpv4Str)
	return *newIpv4
}
