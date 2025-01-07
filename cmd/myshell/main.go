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
		if strings.HasPrefix(cmd, "type") {
			cmd = "type"
		}

		switch cmd {
		case "exit 0":
			os.Exit(0)
		case "echo":
			fmt.Printf("%v", EchoCommand(line))
		case "type":
			fmt.Printf("%v", IsShellBuiltin(line))
		default:
			fmt.Printf("%s: command not found \n", cmd)

		}

	}

}

func EchoCommand(cmd string) string {
	resp := strings.TrimPrefix(cmd, "echo ")
	return resp
}

func IsShellBuiltin(cmd string) string {

	builtin := strings.TrimSpace(strings.TrimPrefix(cmd, "type "))

	builtins := map[string]bool{
		"echo": true,
		"exit": true,
		"type": true,
	}

	if builtins[builtin] {
		return fmt.Sprintf("%s is a shell builtin\n", builtin)
	}
	return fmt.Sprintf("%s: not found\n", builtin)
}
