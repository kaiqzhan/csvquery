package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kaiqzhan/csvquery/cli"
)

func main() {
	c := cli.CLI{
		DataPath: "./data",
	}

	fmt.Println("Welcome to the CSV query demo! Please enter command:")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()
		command := scanner.Text()

		table, err := c.Query(command)
		if err != nil {
			fmt.Print(err)
		}

		fmt.Print(table.String())
	}
}
