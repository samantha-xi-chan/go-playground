package play024_log

import (
	"log"
	"os"
)

func Play() {

	for i := 0; i < 200; i++ {
		log.Println("")
	}

	logFile, err := os.Create("tm2.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// 记录一些日志消息
	log.Println("This is a log message written to the file.")

}
