package main

import (
	"fmt"
	"go_dev/day7/example/example1/balance"
	"math/rand"
)

func main() {
	insts := make([]*balance.Instance)
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, port)
		insts = append(insts)
	}
}
