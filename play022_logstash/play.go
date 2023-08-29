package play022_logstash

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

func Play() {
	// 设置 Logrus 为 JSON 格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 创建 Logstash 连接
	conn, err := net.Dial("tcp", "192.168.31.8:5000") // 替换为您的 Logstash 主机和端口
	if err != nil {
		logrus.Fatal(err)
	}
	defer conn.Close()

	// 设置 Logrus 的输出为 Logstash 连接
	logrus.SetOutput(conn)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Hostname:", hostname)

	// 示例日志
	for i := 1; i <= 999; i++ {
		logrus.WithFields(logrus.Fields{
			"time":     time.Now().Format(time.RFC1123),
			"hostname": hostname,
			"hour":     time.Now().Hour(),
			"minute":   time.Now().Minute(),
		}).Info("ELK")
		//time.Sleep(time.Second)
	}
}
