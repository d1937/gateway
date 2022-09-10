package gateway

import (
	"net"
)

//type DeviceFlagsInfo struct {
//	Name        string
//	Ipaddr      string
//	DeviceName  string
//	Description string
//	Flags       string
//}

func GetInterFaceInfo() (map[string]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	data := make(map[string]string, 0)
	for _, inter := range interfaces {
		address, err := inter.Addrs()
		if err != nil {
			continue
		}
		if len(address) == 2 {
			ip, _, err := net.ParseCIDR(address[1].String())
			if err != nil {
				return nil, err
			}
			if ip.To4() != nil {
				if ip.String() == "127.0.0.1" {
					continue
				}
				data[ip.String()] = inter.Name

			}
		}

	}

	return data, nil
}

//
//func GetAllDevInfo() (map[string]*DeviceFlagsInfo, error) {
//	devices, err := pcap.FindAllDevs()
//	if err != nil {
//		return nil, err
//	}
//
//	info, err := GetInterFaceInfo()
//	if err != nil {
//		return nil, err
//	}
//
//	dinfo := make(map[string]*DeviceFlagsInfo)
//	for _, device := range devices {
//		if len(device.Addresses) == 2 {
//			netaddr := device.Addresses[1]
//
//			if netaddr.IP.To4() != nil {
//				if netaddr.IP.String() != "" {
//					name := ""
//					if v, ok := info[netaddr.IP.String()]; ok {
//						name = v
//					}
//
//					devinfo := &DeviceFlagsInfo{
//						Name:        name,
//						Ipaddr:      netaddr.IP.String(),
//						DeviceName:  device.Name,
//						Description: device.Description,
//						Flags:       fmt.Sprintf("%d", device.Flags),
//					}
//					dinfo[netaddr.IP.String()] = devinfo
//				}
//
//			}
//		}
//	}
//	return dinfo, nil
//}
