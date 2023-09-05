package play602_http_dl

import (
	"log"
	"net/http"
)

func Play() {
	// 设置文件服务器的根目录
	//fs := http.FileServer(http.Dir("/path/to/your/files")) // 可以使用绝对路径 ， 也可以使用 相对路径
	fs := http.FileServer(http.Dir("static")) // 支持文件目录列表 以及文件下载服务

	// 使用默认的路由处理函数来处理文件请求
	http.Handle("/", fs)

	// 启动服务器并监听指定端口
	port := "2080"
	log.Printf("Server started on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
