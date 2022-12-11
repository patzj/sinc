package models

type NetAddressV4 interface {
	Bits() [4][8]uint8
	Octets() Octets
}
