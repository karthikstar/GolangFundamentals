package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin) // read from the standard input - inputting data on a terminal

	sum := 0

	for {
		// infinite loop

		// read string function
		// create a slice of strings which are seperated by '
		input, inputErr := r.ReadString(' ') // will create a slice of strings, which are seperated by ' '
		n := strings.TrimSpace(input)

		if n == "" {
			continue //if n happens to be a empty string we just keep going
		}
		// convert from a string to a number
		num, convErr := strconv.Atoi(n)
		//if have conversion error
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}

		if inputErr == io.EOF {
			// check if we ran out of data - then we can exit
			// iof error doesnt mean its a error
			break
		}

		// check if an actl error occured
		if inputErr != nil {
			fmt.Println("Error reading stdin:", inputErr)
		}
	}
	fmt.Printf("sum: %v\n", sum)

}

// to run it type e.g 1 1 1 in terminal, then control d to send the end of file signal
