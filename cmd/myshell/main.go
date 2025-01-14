package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		if strings.HasPrefix(cmd, "cd") {
			cmd = "cd"
		}

		args := strings.TrimPrefix(line, "cd")

		newargs := strings.Fields(args)

		switch cmd {
		case "exit 0":
			os.Exit(0)
		case "echo":
			fmt.Printf("%v", EchoCommand(line))
		case "type":
			fmt.Printf("%v", IsShellBuiltin(line))
		case "pwd":
			pwd, _ := os.Getwd()
			fmt.Println(pwd)
		case "cd":
			CDCommand(newargs)
		default:
			RunProgram(line)

		}

	}

}

func EchoCommand(cmd string) string {
	resp := strings.TrimPrefix(cmd, "echo ")
	return resp
}

func CDCommand(args []string) {
	command := strings.TrimSpace(args[0])
	if err := os.Chdir(command); err != nil {
		fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", command)
	}
}

func IsShellBuiltin(cmd string) string {
	builtin := strings.TrimSpace(strings.TrimPrefix(cmd, "type "))

	builtins := map[string]bool{
		"echo": true,
		"exit": true,
		"type": true,
		"pwd":  true,
	}

	if builtins[builtin] {
		return fmt.Sprintf("%s is a shell builtin\n", builtin)
	} else {
		inputs := strings.Split(strings.TrimSpace(cmd), " ")
		fpaths := FindPaths(inputs[1:])
		if fpaths == inputs[0] {
			return fmt.Sprintf("%s: not found", inputs[0])
		}
		return fmt.Sprintf("%s\n", fpaths)
	}

}

func FindPaths(args []string) string {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		fp := filepath.Join(path, args[0])
		if _, err := os.Stat(fp); err == nil {
			return fp
		}
	}
	return fmt.Sprintf("%s: not found", args[0])
}

func RunProgram(cmd string) {
	cmds := strings.Split(strings.TrimSpace(cmd), " ")
	command := exec.Command(cmds[0], cmds[1:]...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", cmds[0])
	}
}
