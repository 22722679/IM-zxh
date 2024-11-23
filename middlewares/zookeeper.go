package middlewares

import (
	"encoding/json"
	//"im/middlewares"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

const (
    RandomStr = "random"
)

type ZooClient struct {
    ServerAddr []string
}


type MircoServerInfo struct{
    ServerName      string
    Host            string
    Port            int
}


//服务注册
func (cli *ZooClient)Register(mirco *MircoServerInfo) error {
    //创建连接
    conn, _, err := zk.Connect(cli.ServerAddr, time.Second)

    if err != nil {
        return err
    }

    path := mirco.ServerName + "/"+ RandomStr
    exists, _, err := conn.Exists(path)
    if err != nil {
        return err
    }

    if !exists {
        data , _ := json.Marshal(mirco)
        
        _, err := conn.CreateProtectedEphemeralSequential(path, data, zk.WorldACL(zk.PermAll))
        if err != nil {
            return err
        }
    }

    return nil
}


//服务发现
func (cli *ZooClient) GetNodes(serverName string) ([]*MircoServerInfo, error) {
    conn, _, err := zk.Connect(cli.ServerAddr, time.Second)
    if err != nil {
        return nil, err
    }

    childs, _, err := conn.Children(serverName)
    if err != nil {
        if err == zk.ErrNoNode {
            return []*MircoServerInfo{},nil
        }
        return nil, err
    } 

    
    nodes := make([]*MircoServerInfo, 0)
    for _, child := range childs {
        fullPath := serverName + "/" +child
        data, _, err := conn.Get(fullPath)
        if err != nil {
            if err == zk.ErrNoNode{
                continue
            }
            return nil, err
        }

        node := new(MircoServerInfo)
        if err = json.Unmarshal(data, node); err != nil {
            return nil, err
        }

        nodes =append(nodes, node)
    }

    return nodes, nil
}

