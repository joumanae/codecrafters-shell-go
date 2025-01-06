package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cmd := strings.TrimSpace(line)
		if strings.HasPrefix(cmd, "echo") {
			cmd = "echo"
		}
		switch cmd {
		case "exit 0":
			os.Exit(0)
		case "echo":
			fmt.Printf("%v", EchoCommand(line))
		default:
			fmt.Printf("%s: command not found \n", cmd)

		}

	}

}

func EchoCommand(cmd string) string {
	resp := strings.TrimPrefix(cmd, "echo ")
	return resp
}
