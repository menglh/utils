package main

import (
	"fmt"
	"utils/iputil"
)

func main() {
	ip, _ := iputil.GetSourceIP("8.8.8.8")
	fmt.Println("------", ip.String())
	ip4, _, _ := iputil.GetOutboundIPs()
	i, _ := iputil.GetInterfaceByIp(ip4)
	fmt.Println(i.HardwareAddr)
	fmt.Println(i.Name)
	fmt.Println(i.Index)
	fmt.Println(i.Flags)
	fmt.Println(i.MTU)
	fmt.Println(i.MulticastAddrs())

}
