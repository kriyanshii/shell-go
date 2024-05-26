package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// // You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	stdin := bufio.NewReader(os.Stdin)
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		cmd, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		cmds := strings.Split(cmd, " ")
		if cmds[0] == "exit" {
			os.Exit(0)
		} else if cmds[0] == "echo" {
			fmt.Println(strings.Join(cmds[1:], " "))
		} else {
			fmt.Println(strings.TrimSpace(cmds[0]) + ": command not found")
		}
	}
}
