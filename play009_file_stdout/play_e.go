package play009_file_stdout

import (
	"fmt"
	"github.com/nxadm/tail"
	"log"
	"os"
	"strings"
)

func getLogFileV3(taskId string) string {
	return "./log.txt"
}

func PlayE() {
	go trackV3(nil, "")

	const DOCKER_BIN_ON_MACOS = "/usr/local/bin/docker"

	file, err := os.Create(getLogFileV3(""))

	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			file,
			file,
		},
	}

	cmd := fmt.Sprintf("docker logs --follow %s", "mynginx")
	//cmd := fmt.Sprintf("date")
	strArr := strings.Fields(strings.TrimSpace(cmd))

	//process, err := os.StartProcess("/bin/sh", strArr, procAttr)
	process, err := os.StartProcess(DOCKER_BIN_ON_MACOS, strArr, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}

	process.Wait()

}

func trackV3(chanSig chan int, taskId string) (exit int) {
	chanText := make(chan string)
	t, err := tail.TailFile(getLogFileV3(taskId), tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		return 0
	}

	go func() {
		for line := range t.Lines {
			chanText <- line.Text
		}
	}()

	for {
		select {
		case tt := <-chanText:
			log.Println("line: ", tt)

		case sig := <-chanSig:
			return sig
		}
	}
}
