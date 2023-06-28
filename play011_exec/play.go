package play011_exec

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sync"
)

func Play() {
	cmd := exec.Command("find", "/", "-name", "*.sum")

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("无法获取标准输出管道: %s\n", err)
		return
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("无法获取标准错误输出管道: %s\n", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printOutput(stdoutPipe, "stdout")
	}()

	go func() {
		defer wg.Done()
		printOutput(stderrPipe, "stderr")
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Printf("无法启动进程: %s\n", err)
		return
	}

	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("进程执行错误: %s\n", err)
		return
	}
}

func printOutput(reader io.Reader, name string) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("[%s] %s\n", name, scanner.Text())
	}
}
