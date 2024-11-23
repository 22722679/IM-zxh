package main

import (
	"fmt"
	"time"

	//"im/config"
	"im/middlewares"

	//"github.com/go-playground/validator/v10"
	//"github.com/samuel/go-zookeeper/zk"
)


// 1.了解了zookeeper的数据输出结构
// 2.了解了如何实现服务注册发现
func main() {
	cli := middlewares.ZooClient{ServerAddr: []string{"192.168.28.131:2181"}}

	//确定的是存储的元素的根路径，使用一个常量来表示
	server1 := middlewares.MircoServerInfo{
		ServerName: "/mirco",
		Host:       "192.168.28.131",
		Port:       2181,
	}

	server2 := middlewares.MircoServerInfo{
		ServerName: "/mirco",
		Host:       "192.168.28.131",
		Port:       2181,
	}

	if err := cli.Register(&server1); err != nil {
		fmt.Println("error", err)
	}

	if err := cli.Register(&server2); err != nil {
		fmt.Println("error", err)
	}

	if node, err := cli.GetNodes(server1.ServerName); err != nil {
		fmt.Println("get  nodes failure", err)
	} else {
		fmt.Println(node)
		for _, v := range node {
			fmt.Println(*v)
		}
	}



	time.Sleep(time.Second * 20)
}


