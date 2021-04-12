package main

import (
	"fmt"
	"os"
	"github.com/kanatti/just-go/books/go-by-example/p01_switch"
	"github.com/kanatti/just-go/books/go-by-example/p03_error"
	"github.com/kanatti/just-go/books/go-by-example/p04_goroutines"
	"github.com/kanatti/just-go/books/go-by-example/p05_channels"
	"github.com/kanatti/just-go/books/go-by-example/p06_select"
)

func main() {
	chapterMappings := map[string]func() {
		"p01_switch": p01_switch.Run,
		"p03_error": p03_error.Run,
		"p04_goroutines": p04_goroutines.Run,
		"p05_channels": p05_channels.Run,
		"p06_select": p06_select.Run,
	}


	if len(os.Args) <= 1 {
		fmt.Println("Please provide chapter name\nAvailable Chapters:")
		chapterNames := getKeys(chapterMappings);
		for i, chapter := range chapterNames {
			fmt.Printf("%d %s\n", i+1, chapter)
		}
		return
	}

	chapter := os.Args[1]
	if chapterRun, ok := chapterMappings[chapter]; ok {
		chapterRun()
	} else {
		fmt.Println("Unknown chapter")
	}
}

func getKeys(m map[string]func()) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}