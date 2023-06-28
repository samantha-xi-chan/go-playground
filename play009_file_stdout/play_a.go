package play009_file_stdout

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"
)

func PlayA() {
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
			scanner := bufio.NewScanner(stdoutPipeReader) // 使用bufio.Scanner来读取管道数据
			for scanner.Scan() {
				line := scanner.Text() // 获取每行的文本
				fmt.Println("Line:", line)
			}

			if err := scanner.Err(); err != nil {
				fmt.Println("Error:", err)
			}
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
