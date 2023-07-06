package play009_file_stdout

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

const (
	STRING_BUF_SIZE = 1024
)

func PlayF(cmdStr string, stdOut chan string, stdErr chan string) int {
	strArr := strings.Fields(strings.TrimSpace(cmdStr))
	log.Println(strArr)
	cmd := exec.Command("/usr/local/bin/docker", "logs", "--follow", "nginx001")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("StdoutPipe: ", err)
		return 0
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("StderrPipe: ", err)
		return 0
	}

	err = cmd.Start()
	if err != nil {
		log.Println("cmd.Start: ", err)
		return 0
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			stdOut <- line
		}
	}()

	go func() {
		for {
			buffer := make([]byte, STRING_BUF_SIZE)
			n, err := stderr.Read(buffer)
			if err != nil {
				log.Println("stderr.Read err:", err)
				break
			}
			x := string(buffer[:n])
			stdErr <- x
		}
	}()

	err = cmd.Wait()
	if err != nil {
		log.Println("cmd.Wait err: ", err)
		return 0
	}

	log.Println("cmd.Wait end")
	return 0
}
