package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	// biz here
	Value int
}

var instance *singleton

var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			Value: 42,
		}
	})

	return instance
}

func main() {
	instance1 := GetInstance()
	fmt.Println(instance1.Value) // 输出: 42

	instance2 := GetInstance()
	fmt.Println(instance2.Value) // 输出: 42

	if instance1 == instance2 {
		fmt.Println("instance1 和 instance2 是同一个实例")
	}
}
