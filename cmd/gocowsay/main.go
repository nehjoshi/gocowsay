package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/nehjoshi/gocowsay/cowsay"
)

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	lines = cowsay.TabsToSpaces(lines)
	maxwidth := cowsay.CalculateMaxWidth(lines)
	messages := cowsay.NormalizeStringsLength(lines, maxwidth)
	balloon := cowsay.BuildBalloon(messages, maxwidth)
	fmt.Println(balloon)
	fmt.Println(cow)
	fmt.Println()
}
