package main

import (
	"fmt"
	"gostudy/oblsn/src/day7/exp1_balance/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	var insts []*balance.Instance
	//create 20 instance
	for i := 0; i < 20; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		fmt.Printf("Instance%d:%s create succ\n", i, one.String())
		insts = append(insts, one)
	}

	var balanceName = "random"
	//命令行第一个字符
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}
	//第一种比较蠢的方法 根据指令在运行时确定用那种算法
	/*
		    	var balanceName = "random"
				var balancer balance.Balancer
				//命令行第一个字符
				if len(os.Args) > 1 {
					balanceName = os.Args[1]
				}
				if balanceName == "random" {
					balancer = &balance.RandomBalance{}
					fmt.Println("use random ")
				} else if balanceName == "poll" {
					balancer = &balance.PollBalance{}
					fmt.Println("use poll")
				}*/

	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			//fmt.Println("do balance err:", err)
			fmt.Fprintf(os.Stdout, "do balance err\n")
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
