package main

import (
	"fmt"
	"github.com/cornelk/hashmap"
)

func main() {
	m := hashmap.New(1)
	m.Set("amount1", 123)
	m.Set("amount2", 456)
	fmt.Println(m.Get("amount1"))

	fmt.Println(m.Get("amount2"))
}
