package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func echo(args []string, options map[string]bool) {
	if options["newline"] {
		_echo_newline(args[1:])
	} else if options["linenum"] {
		_echo_linenum(args[1:])
	} else {
		_echo_default(args)
	}
}

func _echo_newline(args []string) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func _echo_linenum(args []string) {
	for i, arg := range args {
		fmt.Println(strconv.Itoa(i) + ": " + arg)
	}
}

func _echo_default(args []string) {
	fmt.Println(strings.Join(args, " "))
}
