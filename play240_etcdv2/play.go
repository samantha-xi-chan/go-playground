package play240_etcdv2

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"time"
)

const (
	HB_TIMEOUT = 10
	KEY        = "/key001"
	END_POINT  = "http://192.168.31.7:2379"
)

func Play02() {
	go func() {
		time.Sleep(time.Second * 10)
		PutVal(KEY, "new-value")
		time.Sleep(time.Second * 10)
		PutVal(KEY, "finish")
	}()

	waitValue()
}

func waitValue() {
	etcdClient := etcd.NewClient([]string{END_POINT})

	fmt.Println("Watching for changes...")
	watch := make(chan *etcd.Response)
	go etcdClient.Watch(KEY, 0, true, watch, nil)

	for {
		select {
		case resp := <-watch:
			fmt.Printf("Event received! Key: %s, Value: %s\n", resp.Node.Key, resp.Node.Value)
			if resp.Node.Value == "finish" {
				log.Println("match. going to break")
				break
			}
		case <-time.After(HB_TIMEOUT * time.Second):
			fmt.Printf("No changes in %d seconds.", HB_TIMEOUT)
		}

		log.Println("i am in for loop")
	}

	log.Println("for loop end")
}

func PutVal(key string, newValue string) {
	// Create a new etcd client
	client := etcd.NewClient([]string{END_POINT})

	_, err := client.Set(key, newValue, 0)
	if err != nil {
		log.Fatalf("Failed to update key %s: %v", key, err)
	}

	fmt.Printf("Updated key %s with value %s\n", key, newValue)
}
