package main

import (
	"fmt"
	"net"
	"utils/iputil"
)

func main() {
	fmt.Println(iputil.GetFreePort())
	//domain := "wxapi.soyoung.com"
	domain := "www.baidu.com"
	ip, _ := net.ResolveIPAddr("ip", domain)
	fmt.Println(ip.String())

	port, _ := net.LookupPort("tcp", "sftp") // 查看telnet服务器使用的端口
	fmt.Println(port)
}
