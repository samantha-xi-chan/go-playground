package play013_redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"log"
	"time"
)

// go get "github.com/go-redis/redis/v8"

func Play() {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with the address of your Redis server
		Password: "",               // No password if not set
		DB:       0,                // Use default DB
	})

	// Ping the Redis server to check the connection
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error pinging Redis server:", err)
		return
	}
	fmt.Println("Redis server responded:", pong)

	// Set a key-value pair
	err = client.Set(context.Background(), "mykey", "myvalue", 8*time.Second).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
		return
	}
	fmt.Println("Key-value pair set successfully.")

	time.Sleep(5 * time.Second)
	// Get the value of the key
	val, err := client.Get(context.Background(), "mykey").Result()
	if err != nil {
		fmt.Println("Error getting key:", err)
		return
	}
	fmt.Println("Value of 'mykey' is:", val)

	time.Sleep(5 * time.Second)
	// Get the value of the key
	val, err = client.Get(context.Background(), "mykey").Result()
	if err != nil {
		fmt.Println("Error getting key:", err)
		return
	}
	fmt.Println("Value of 'mykey' is:", val)

	// Wait for a few seconds before closing the connection
	time.Sleep(5 * time.Second)

	// Close the Redis client connection
	err = client.Close()
	if err != nil {
		fmt.Println("Error closing Redis connection:", err)
		return
	}
	fmt.Println("Redis connection closed.")
}

func Sub() {
	options := &redis.Options{
		Addr: "localhost:6379",
		DB:   0, // Your preferred database number
	}

	client := redis.NewClient(options)

	pubsub := client.Subscribe(context.Background(), "__keyevent@0__:expired")
	channel := pubsub.Channel()

	for msg := range channel {
		// Do something when a key expires
		println("Key expired:", msg.Payload)
	}
}

type RedisManager struct {
	Address string
	Client  *redis.Client
	MaxSize int64
}

func (mgr *RedisManager) Init() (e error) {
	// Create a new Redis client
	mgr.Client = redis.NewClient(&redis.Options{
		Addr:     mgr.Address, // Replace with the address of your Redis server
		Password: "",          // No password if not set
		DB:       0,           // Use default DB
	})

	// Ping the Redis server to check the connection
	pong, err := mgr.Client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error pinging Redis server:", err)
		return
	}
	fmt.Println("Redis server responded:", pong)

	return nil
}

func (mgr *RedisManager) NewLog(ctx context.Context, trim bool, key string, val string) (e error) {

	mgr.Client.LPush(ctx, key, val)

	if trim {
		_, e = mgr.Client.LTrim(ctx, key, 0, mgr.MaxSize-1).Result()
		if e != nil {
			log.Println("e: ", e)
		}
	}

	return nil
}

func (mgr *RedisManager) NewLogMulti(ctx context.Context, trim bool, key string, vals []string) (e error) {

	mgr.Client.LPush(ctx, key, vals)

	if trim {
		_, e = mgr.Client.LTrim(ctx, key, 0, mgr.MaxSize-1).Result()
		if e != nil {
			log.Println("e: ", e)
		}
	}

	return nil
}

func (mgr *RedisManager) Traversal(ctx context.Context, trim bool, key string, startFromRear bool) (e error) {

	if trim {
		_, e = mgr.Client.LTrim(ctx, key, 0, mgr.MaxSize-1).Result()
		if e != nil {
			log.Println("e: ", e)
			return e
		}
	}

	elements, err := mgr.Client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		log.Println("Error:", err)
		return
	}

	if startFromRear {
		log.Println("List elements:")
		for _, element := range elements {
			log.Println(element)
		}
	} else {
		for i := len(elements) - 1; i >= 0; i-- {
			fmt.Println(elements[i])
		}
	}

	return nil
}
