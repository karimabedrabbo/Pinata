package managers

import (
	"github.com/phuslu/geoip"
	"net"
)

type IpFilter struct {
}

var ipFilter *IpFilter

func SetupIpFilter() *IpFilter{
	return &IpFilter{}
}

func InitIpFilter() {
	ipFilter = SetupIpFilter()
}

func GetIpFilter() *IpFilter {
	return ipFilter
}

func (f *IpFilter) IpAllowed(ip string) bool {

	//only allow US traffic (for now)
	allowed := string(geoip.Country(net.ParseIP(ip))) == "US"

	//special case localhost
	if !allowed &&
		(ip == "::1" || ip == "127.0.0.1") {
		allowed = true
	}

	return allowed
}
