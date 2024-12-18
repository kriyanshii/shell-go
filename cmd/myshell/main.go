package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var KnownCommands = map[string]int{"exit": 0, "echo": 1, "type": 2, "pwd": 3, "cd": 4}

func main() {
	// // You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	stdin := bufio.NewReader(os.Stdin)
	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, _ := stdin.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		cmds := strings.Split(cmd, " ")
		command := cmds[0]
		switch command {
		case "exit":
			code, _ := strconv.Atoi(cmds[1])
			os.Exit(code)
		case "echo":
			fmt.Printf("%s\n", strings.Join(cmds[1:], " "))
		case "type":
			handleType(cmds[1:])
		case "pwd":
			pwd, _ := os.Getwd()
			fmt.Println(pwd)
			continue
		case "cd":
			handleCd(cmds[1:])
		default:
			command := exec.Command(cmds[0], cmds[1:]...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout

			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found \n", command)
			}
		}
	}
}

func handleCd(args []string) {
	command := args[0]
	if err := os.Chdir(command); err != nil {
		fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", command)
	}
}

func handleType(args []string) {
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
