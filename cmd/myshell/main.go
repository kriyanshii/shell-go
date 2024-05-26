package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// // You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	stdin := bufio.NewReader(os.Stdin)
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, _ := stdin.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		split := strings.Split(cmd, " ")
		command := split[0]
		switch command {
		case "exit":
			code, _ := strconv.Atoi(split[1])
			os.Exit(code)
		case "echo":
			fmt.Printf("%s\n", strings.Join(split[1:], " "))
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
