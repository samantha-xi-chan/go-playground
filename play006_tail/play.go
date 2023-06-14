package play006_tail

import (
	"github.com/nxadm/tail"
	"log"
)

func Test() {
	// 创建一个Tail对象，传入要监视的文件路径
	t, err := tail.TailFile("~/Desktop/a.txt", tail.Config{
		Follow: true, // 设置为true，表示持续监视文件的新增内容
	})

	if err != nil {
		log.Fatal(err)
	}

	// 通过循环读取Tail对象的Lines通道，可以获取到新增的内容
	for line := range t.Lines {
		// 处理每一行的日志内容
		// 这里可以根据实际需求进行处理
		log.Println(line.Text)
	}
}
