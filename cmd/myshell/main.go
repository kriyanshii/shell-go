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

	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")
	// Wait for user input
	cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s: command not found\n", strings.TrimRight(cmd, "\n"))
}
