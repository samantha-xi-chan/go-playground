package play009_file_stdout

import (
	"fmt"
	"os"
	"strings"
)

func PlayC() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stdout,
		},
	}

	cmd := fmt.Sprintf("date")
	strArr := strings.Fields(strings.TrimSpace(cmd))

	process, err := os.StartProcess("/bin/sh", strArr, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}

	process.Wait()

}
