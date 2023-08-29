package play020_elk

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func Play() {
	// 设置 Logrus 日志格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 设置 Logstash 的地址和端口
	logstashHook, err := NewLogstashHook("192.168.31.45:5000")
	if err != nil {
		logrus.Error(err)
	}

	// 添加 Logstash 钩子到 Logrus
	logrus.AddHook(logstashHook)

	// 记录日志
	for true {
		logrus.Info("This is an info log.")
		logrus.Warn("This is a warning log.")
		logrus.Error("This is an error log.")

		time.Sleep(10 * time.Second)
	}
}

// LogstashHook 定义 Logstash 钩子结构
type LogstashHook struct {
	Addr string
}

// NewLogstashHook 创建 Logstash 钩子
func NewLogstashHook(addr string) (*LogstashHook, error) {
	return &LogstashHook{
		Addr: addr,
	}, nil
}

// Fire 实现 Logrus 钩子接口的 Fire 方法
func (hook *LogstashHook) Fire(entry *logrus.Entry) error {
	// 创建 Logstash 日志条目
	logstashEntry := map[string]interface{}{
		"@timestamp": entry.Time.Format(time.RFC3339),
		"message":    entry.Message,
		"level":      entry.Level.String(),
		"fields":     entry.Data,
	}

	// 发送日志到 Logstash
	// 在这里，你可以使用合适的方法将 logstashEntry 发送到 Logstash 服务
	fmt.Println("Sending log to Logstash:", logstashEntry)

	return nil
}

// Levels 实现 Logrus 钩子接口的 Levels 方法
func (hook *LogstashHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
