# cdncheck
检查给定 IP 是否属于已知 CDN 范围（akamai、cloudflare、incapsula、sucuri 和 leaseweb）的帮助程序库,下面是基础使用示例:

```go
package main

import (
	"log"
	"net"

	"github.com/projectdiscovery/cdncheck"
)

func main() {
	// uses projectdiscovery endpoint with cached data to avoid ip ban
	// Use cdncheck.New() if you want to scrape each endpoint (don't do it too often or your ip can be blocked)
	client, err := cdncheck.NewWithCache()
	if err != nil {
		log.Fatal(err)
	}
	if found, result, err := client.Check(net.ParseIP("173.245.48.12")); found && err == nil {
		log.Println("ip is part of cdn, result:", result)
	}
}
```