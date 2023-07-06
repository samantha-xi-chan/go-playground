package play009_file_stdout

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os/exec"
)

// docker rm -f nginx001;
// docker run --name nginx001 -p 80:80 nginx:latest;

const (
	FUNC_ERROR_CODE_UNKNOW = -1
)

type Callback func(int)

func StartProcBlo(stdOut chan string, stdErr chan string, cb Callback, cmdName string, cmdArg ...string) (funcErrCode int, procErrCode int, e error) {
	//strArr := strings.Fields(strings.TrimSpace(cmdStr))
	//log.Println(strArr)
	cmd := exec.Command(cmdName, cmdArg...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("StdoutPipe: ", err)
		return FUNC_ERROR_CODE_UNKNOW, 0, nil
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("StderrPipe: ", err)
		return FUNC_ERROR_CODE_UNKNOW, 0, nil
	}

	err = cmd.Start()
	if err != nil {
		log.Println("cmd.Start: ", err)
		return FUNC_ERROR_CODE_UNKNOW, 0, nil
	}
	pid := cmd.Process.Pid
	fmt.Println("Process ID:", pid)
	cb(pid)

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			stdOut <- scanner.Text()
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			stdErr <- scanner.Text()
		}
	}()

	err = cmd.Wait()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode := exitErr.ExitCode()
			fmt.Println("exec.ExitError ExitCode: ", exitCode)
			return 0, exitCode, nil
		} else {
			fmt.Println("exec.ExitError: ", err)
			return 0, 0, errors.Wrap(err, "")
		}
	}
	log.Println("cmd.Wait end ok")
	return FUNC_ERROR_CODE_UNKNOW, 0, nil
}
