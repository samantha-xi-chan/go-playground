package play008_file

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Play() {
	// 创建输出文件
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建子进程并设置输出文件
	cmd := exec.Command("echo", "Hello, world!")
	cmd.Stdout = file

	// 启动子进程
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// 等待子进程结束
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("子进程执行完毕！")

	// 在子进程执行完毕之后修改输出的目标
	// 这里不能再修改 cmd.Stdout 字段了

	// 重新打开另一个输出文件
	anotherFile, err := os.Create("another_output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer anotherFile.Close()

	// 修改输出目标为另一个文件
	cmd.Stdout = anotherFile

	// 启动子进程
	cmd = exec.Command("echo", "Hello, world!")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// 等待子进程结束
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("修改后的子进程执行完毕！")
}
