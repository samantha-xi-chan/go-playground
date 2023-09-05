package play601_socks5

// export GOPROXY=https://goproxy.io
// go get github.com/armon/go-socks5

import (
	"github.com/armon/go-socks5"
	"log"
	"os"
)

func Play() {
	// 创建一个SOCKS5服务器配置
	config := &socks5.Config{}

	// 创建一个新的SOCKS5服务器
	server, err := socks5.New(config)
	if err != nil {
		log.Fatalf("Error creating SOCKS5 server: %v", err)
		os.Exit(1)
	}

	// 启动SOCKS5服务器并监听指定的地址和端口
	addr := "0.0.0.0:1080" // 代理服务器地址和端口
	if err := server.ListenAndServe("tcp", addr); err != nil {
		log.Fatalf("Error starting SOCKS5 server: %v", err)
		os.Exit(1)
	}

	log.Printf("SOCKS5 server is listening on %s", addr)

	// 阻塞程序以保持服务器运行
	select {}
}
