package play240_etcdv2

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"time"
)

func Play() {
	// 创建etcd客户端
	client := etcd.NewClient([]string{"http://192.168.31.45:2379"})

	// 设置键值对
	//_, err := client.Set("/key001", "your_value", 0)
	//if err != nil {
	//	fmt.Println("Error setting key:", err)
	//	return
	//}

	// 获取键值
	response, err := client.Get("/key001", false, false)
	if err != nil {
		fmt.Println("Error getting key:", err)
		return
	}

	// 打印获取到的键值信息
	fmt.Println("Key:", response.Node.Key)
	fmt.Println("Value:", response.Node.Value)

	Play02()
}

func Play02() {
	etcdClient := etcd.NewClient([]string{"http://192.168.31.45:2379"})

	fmt.Println("Watching for changes...")
	watch := make(chan *etcd.Response)
	go etcdClient.Watch("/key001", 0, true, watch, nil)

	for {
		select {
		case resp := <-watch:
			fmt.Printf("Event received! Key: %s, Value: %s\n", resp.Node.Key, resp.Node.Value)
		case <-time.After(1000 * time.Second):
			fmt.Println("No changes in 10 seconds.")
		}
	}
}
