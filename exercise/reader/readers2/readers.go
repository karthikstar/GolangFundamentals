//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When thpe user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	count := 0
	nonBlankLines := 0

	for {
		input, _ := r.ReadString('\n')
		n := strings.TrimSpace(input)
		nonBlankLines += 1
		if n == "" {
			continue
		}

		count += 1
		if n == "hello" {
			fmt.Println("hi there!")
		}

		if n == "bye" {
			fmt.Println("byeee!!")
		}

		if n == "Q" || n == "q" {
			break
		}

	}
	fmt.Println("count", count)
	fmt.Println("non blank lines", nonBlankLines)

}
