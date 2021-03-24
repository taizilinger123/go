package main

import (
	"fmt"
	"go_dev/day7/example/example1/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	//insts := make([]*balance.Instance)
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}

	//var balancer balance.Balancer
	//var conf = "random"
	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}
	/*
		if conf == "random" {
			balancer = &balance.RandomBalance{}
			fmt.Println("use random balancer")
		} else if conf == "roundrobin" {
			balancer = &balance.RoundRobinBalance{}
			fmt.Println("use roundrobin balancer")
		}
	*/
	//balancer := &balance.RandomBalance{}  //实现负载均衡算法
	//balancer := &balance.RoundRobinBalance{} //实现rr轮循算法
	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}

/*
cmd里
D:\project>go  build  go_dev/day7/example/example1/main
D:\project>main.exe
192.168.143.133:8080
192.168.16.138:8080
192.168.119.151:8080
192.168.237.37:8080
192.168.167.131:8080

ctrl+b打开和显示和隐藏左边工程project
*/
