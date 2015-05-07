package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) {
		println(c.Args()[0])
		println("Hello friend!")
		Execute(c.Args())
	}

	app.Run(os.Args)
}

func Execute(args []string) {
	fmt.Println(args)
	fmt.Printf("PASTE code snippet (First 120 characts are used): \n%s\n", strings.Repeat("*", 80))
	code, err := ReadInput(bufio.NewReader(os.Stdin))

	if err != nil {
		fmt.Println("Error receiving input ")
		os.Exit(0)
	}
	fmt.Println(FormatString(strings.TrimSpace(strings.Trim(code, "^"))))
	fmt.Printf(strings.Repeat("*", 80))
}
