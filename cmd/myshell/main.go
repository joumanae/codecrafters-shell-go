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
		switch cmd {
		case "exit 0":
			os.Exit(0)
		default:
			fmt.Printf("%s: command not found \n", cmd)

		}

	}

}
