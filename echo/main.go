package main

import (
	"os"
)

func main() {
	var flag = os.Args[1]
	args := make(map[string]bool)

	switch flag {
	case "--newline":
		args["newline"] = true
	case "--linenum":
		args["linenum"] = true
	}

	echo(os.Args[1:], args)
}
