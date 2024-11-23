package service

import (
	"fmt"
	"im/middlewares"
	"time"
)

func MessageFind() {
	cli := middlewares.ZooClient{ServerAddr: []string{"192.168.28.131:2181"}}

	//确定的是存储的元素的根路径，使用一个常量来表示
	server1 := middlewares.MircoServerInfo{
		ServerName: "/mirco",
		Host:       "192.168.28.131",
		Port:       2181,
	}
	for t := 1; t == 1; {
		fmt.Println("-1")
		if MS != nil {
			if err := cli.Register(&server1); err != nil {
				fmt.Println("error", err)
				continue
			}
			fmt.Println("zookeeper add success!!!")
			MS = nil
		}
		time.Sleep(time.Second * 20)
	}
	panic("timeout")
}
