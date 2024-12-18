package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var KnownCommands = map[string]int{"exit": 0, "echo": 1, "type": 2}

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
		case "type":
			checkType(split[1:])
		default:
			fmt.Printf("%s: command not found \n", command)
		}
	}
}

func checkType(args []string) {
	item := args[0]
	if _, exists := KnownCommands[item]; exists {
		class := "builtin"
		fmt.Fprintf(os.Stdout, "%v is a shell %v\n", item, class)
		return
	} else {
		paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			exec := filepath.Join(path, item)
			if _, err := os.Stat(exec); err == nil {
				fmt.Fprintf(os.Stdout, "%v is %v\n", item, exec)
				return
			}
		}
	}
	fmt.Fprintf(os.Stdout, "%v: not found\n", item)
}
