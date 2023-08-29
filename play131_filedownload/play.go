package play131_filedownload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const LOCAL_PATH = "./test/"
const URL_PATH = "/download/"

func Play() {
	// 指定要监听的地址和端口
	address := "0.0.0.0"
	port := "8080"

	// 设置文件下载目录

	// 创建文件下载目录
	if err := os.MkdirAll(LOCAL_PATH, os.ModePerm); err != nil {
		panic(err)
	}

	// 注册文件下载处理器
	http.HandleFunc(URL_PATH, func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.Path[len(URL_PATH):]
		filePath := filepath.Join(LOCAL_PATH, fileName)

		// 打开文件
		file, err := os.Open(filePath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer file.Close()

		// 设置响应头，告知浏览器要下载文件
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
		w.Header().Set("Content-Type", "application/octet-stream")
		fileInfo, _ := file.Stat()
		w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

		// 将文件内容写入响应流
		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "Error serving file", http.StatusInternalServerError)
			return
		}
	})

	// 启动HTTP服务器
	fmt.Printf("Server listening on %s:%s\n", address, port)
	http.ListenAndServe(address+":"+port, nil)
}
