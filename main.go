package main

import (
	"fmt"
	"go-playground/play201_kafka"
)

//
//func level3(arg ...string) string {
//	result := ""
//	for _, str := range arg {
//		result += str
//	}
//
//	return result
//}
//
//func level2(arg ...string) string {
//	//strs := []string{}
//	//
//	//for _, str := range arg {
//	//	strs = append(strs, str)
//	//}
//	//
//	//return level3(strs...)
//
//	return level3(arg...)
//}

func PrintResult(result int) {
	fmt.Println("操作结果:", result)
}

const SIZE = 1024

func main() {
	//s := level2("a", " ", "c")
	//log.Println(s)

	//log.Printf("hello world")
	//play001_error.Test()
	//play002_error.Test()
	//play003_error.Test()
	//
	//x := play005_float.AddFloat(0.300, 0.6000)
	//log.Println(x)
	//
	//y := play005_float.AddInt(300000, 60000)
	//log.Println(y)
	//
	//play006_tail.Test()
	//log.Println(y)
	//
	//play007_cron.TestCron()
	//play008_file.Play()
	//play009_file_stdout.PlayA()
	//play009_file_stdout.PlayB()
	//play009_file_stdout.PlayBB()
	//play009_file_stdout.PlayC()

	//play009_file_stdout.PlayD()
	//play009_file_stdout.PlayE()

	// docker rm -f nginx001; docker run --name nginx001 -p 80:80 nginx:latest

	//play010_gin.Play()
	//play011_exec.Play()

	//chanStdOut := make(chan string, SIZE)
	//chanStdErr := make(chan string, SIZE)
	//go func() {
	//	for {
	//		select {
	//		case x := <-chanStdOut:
	//			log.Println("chanStdOut: ", x)
	//		case y := <-chanStdErr:
	//			log.Println("chanStdErr: ", y)
	//		}
	//	}
	//}()
	//play009_file_stdout.StartProcBlo(chanStdOut, chanStdErr, PrintResult, "/usr/local/bin/docker", "logs", "--follow", "nginx001")

	//play012_gopark.Play()

	//play901_pprof.Play()
	//play101_docker.Play()
	//play110_sse.Play()

	play201_kafka.Consume()
	//time.Sleep(time.Second)

	//play201_kafka.Produce()

}
