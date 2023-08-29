package play013_redis

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	t.Log("TestRedis")

	KEY := time.Now().String()
	redisManager := RedisManager{
		Address: "localhost:6379",
		Client:  nil,
		MaxSize: 5,
	}

	redisManager.Init()
	log.Println("tick: ", time.Now().String())
	for i := 0; i < 1_0_000; i++ {
		redisManager.NewLog(context.Background(), false, KEY, fmt.Sprintf("line %d", i))
	}

	//J_MAX := 1000
	//BAT_SIZE := 1_00
	//for j := 0; j < J_MAX; j++ {
	//	array := make([]string, BAT_SIZE)
	//	for i := 0; i < BAT_SIZE; i++ {
	//		array = append(array, fmt.Sprintf("line %d", j*BAT_SIZE+i))
	//		redisManager.NewLogMulti(context.Background(), false, KEY, array)
	//	}
	//}

	log.Println("tick: ", time.Now().String())
	redisManager.Traversal(context.Background(), true, KEY, true)

	log.Println("tick: ", time.Now().String())
}
