// +build !darwin,!linux,!windows,!solaris,!freebsd

package gateway

import (
	"net"
)

func discoverGatewayOSSpecific(arg string) (ip net.IP, err error) {
	return ip, errNotImplemented
}

func discoverGatewayInterfaceOSSpecific(arg string) (ip net.IP, err error) {
	return nil, errNotImplemented
}
