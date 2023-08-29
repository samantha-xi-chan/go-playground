package play120_viper

import (
	"fmt"
	"github.com/spf13/viper"
)

// go get github.com/spf13/viper

func Play() {
	// 初始化 Viper
	viper.SetConfigFile("configs/app.yaml")
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 读取配置信息
	host := viper.GetString("app.host")
	port := viper.GetInt("app.port")
	debug := viper.GetBool("app.debug")

	// 使用配置信息
	fmt.Printf("Host: %s\n", host)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Debug: %v\n", debug)
}
