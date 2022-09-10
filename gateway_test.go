package gateway

import (
	"fmt"
	"testing"
)

func TestDiscoverGateway(t *testing.T) {
	//devinfo,err := GetAllDevInfo()
	//if err!=nil {
	//	t.Fatal(err)
	//}
	//
	//for ip,dev := range devinfo {
	//	fmt.Println(ip,dev)
	//}

	//macos 根据选择的网卡名称,获取网关IP
	//fmt.Println(DiscoverInterface(""))
	//fmt.Println(DiscoverGateway("en2"))

	//linux 根据选择的网卡名称,获取网关IP
	//fmt.Println(DiscoverInterface(""))
	//fmt.Println(DiscoverGateway("ens38"))

	//windows 根据网卡IP获取网关IP
	//fmt.Println(DiscoverInterface(""))
	fmt.Println(DiscoverGateway(""))
}
