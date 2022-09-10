// +build darwin

package gateway

import (
	"errors"
	"net"
	"os/exec"
	"strings"
)

func discoverGatewayOSSpecific(arg string) (net.IP, error) {
	//routeCmd := exec.Command("/sbin/route", "-n", "get", "0.0.0.0")
	//output, err := routeCmd.CombinedOutput()
	//if err != nil {
	//	return nil, err
	//}
	//
	//return parseDarwinRouteGet(output,arg)

	routeCmd := exec.Command("netstat", "-rn")
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return parseDarWinNetstat(output, arg)

}

//获取本机IP
func discoverGatewayInterfaceOSSpecific(arg string) (ip net.IP, err error) {
	interfaceAddrs, err := GetInterFaceInfo()
	if err != nil {
		return nil, err
	}
	//获取本机网关IP
	gateWayIP, err := discoverGatewayOSSpecific(arg)
	if err != nil {
		return nil, err
	}

	for ipaddr, _ := range interfaceAddrs {
		ips1 := strings.Split(ipaddr, ".")
		ips2 := strings.Split(gateWayIP.String(), ".")
		if ips1[0] == ips2[0] && ips1[1] == ips2[1] && ips1[2] == ips2[2] {
			return net.ParseIP(ipaddr), nil
		}
	}

	return nil, errors.New("获取本机网卡失败")
}

func parseDarWinNetstat(output []byte, arg string) (net.IP, error) {
	//	Routing tables
	//
	//Internet:
	//	Destination        Gateway            Flags        Netif Expire
	//	default            192.168.208.2      UGSc           en0
	//	default            192.168.1.1        UGScI          en2
	//	127                127.0.0.1          UCS            lo0
	//	127.0.0.1          127.0.0.1          UH             lo0
	//	169.254            link#4             UCS            en0      !
	//	169.254            link#5             UCSI           en2      !
	//	169.254.252.182    48:89:e7:e9:5a:2   UHLSW          en2      !
	//	192.168.0/22       link#5             UCS            en2      !
	//	192.168.0.107/32   link#5             UCS            en2      !
	//	192.168.0.232      0:be:d5:5c:d3:cd   UHLWI          en2    788
	//	192.168.1.1/32     link#5             UCS            en2      !
	//	192.168.1.1        44:f9:71:f3:82:a8  UHLWIir        en2   1199
	//	192.168.1.118      50:2b:73:a9:1d:cc  UHLWIi         en2   1136
	//	192.168.1.130      48:89:e7:e9:5a:2   UHLWI          en2    223
	//	192.168.1.200      e8:6f:38:b9:d2:a9  UHLWI          en2   1143
	//	192.168.2.50       da:a:ec:24:e1:fc   UHLWI          en2    596

	outputLines := strings.Split(string(output), "\n")
	for _, line := range outputLines {
		fields := strings.Fields(line)
		if len(fields) >= 3 && fields[0] == "default" {
			if arg != "" {
				if fields[3] == arg {
					return net.ParseIP(fields[1]), nil
				}

			} else {
				ip := net.ParseIP(fields[1])
				if ip != nil {
					return ip, nil
				}
			}

		}
	}

	return nil, errNoGateway
}
