package play009_file_stdout

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"
)

func PlayB() {
	stdinPipeReader, stdinPipeWriter := io.Pipe()
	stdoutPipeReader, stdoutPipeWriter := io.Pipe()

	cmd := exec.Command("cat")
	cmd.Stdin = stdinPipeReader
	cmd.Stdout = stdoutPipeWriter

	// 启动子进程
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			inputData := []byte("11 Hello, child process! \n 22 Hello, child process! \n ")
			_, err = stdinPipeWriter.Write(inputData)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}

		log.Println("end of for")
		// 关闭输入管道，表示写入完成
		stdinPipeWriter.Close()
	}()

	go func() {
		for {
			// 从输出管道读取数据
			outputData := make([]byte, 200)
			n, err := stdoutPipeReader.Read(outputData)
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}

			// 输出子进程的输出结果
			fmt.Printf("Output: %s\n", outputData[:n])
		}

		log.Println("end of for")
	}()

	// 等待子进程结束
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Child process completed!")
}
