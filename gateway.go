package gateway

import (
	"errors"
	"net"
	"runtime"
)

var (
	errNoGateway      = errors.New("no gateway found")
	errCantParse      = errors.New("can't parse string output")
	errNotImplemented = errors.New("not implemented for OS: " + runtime.GOOS)
)

// DiscoverGateway is the OS independent function to get the default gateway
//获取网关ip
func DiscoverGateway(arg string) (ip net.IP, err error) {
	return discoverGatewayOSSpecific(arg)
}

// DiscoverInterface is the OS independent function to call to get the default network interface IP that uses the default gateway
//获取本机IP
func DiscoverInterface(arg string) (ip net.IP, err error) {
	return discoverGatewayInterfaceOSSpecific(arg)
}
