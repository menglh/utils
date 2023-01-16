package main

import (
	"fmt"
	"github.com/projectdiscovery/fdmax"

	_ "github.com/projectdiscovery/fdmax/autofdmax"
)

func main() {
	fmt.Println("test")
	fdmax.Set(1)
	l, _ := fdmax.Get()

	fmt.Println(l.Max, l.Current)
}
