package utils

import (
	"github.com/patzj/sinc/models"
)

func Ipv4Hosts(networkAddr, broadcastAddr *models.IPv4) []models.IPv4 {
	hosts := []models.IPv4{}
	host := *networkAddr

	for {
		host = models.Offset(host, 1)
		if host.IsBefore(*broadcastAddr) {
			hosts = append(hosts, host)
		} else {
			break
		}
	}

	return hosts
}
