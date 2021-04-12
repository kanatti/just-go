package p04_goroutines

import (
	"time"
	"fmt"
)

func Run() {
	start := time.Now()
	runTimed(func() {
		f("main-thread") // blocks
	})

	runTimed(func() {
		go runTimed(func() { f("routine-1"); }) // non blocking
		go runTimed(func() { f("routine-2"); })
		go runTimed(func() { f("routine-3"); })
		go runTimed(func() { f("routine-4"); })
		go runTimed(func() { f("routine-5"); })
		go runTimed(func() { f("routine-6"); })
		go runTimed(func() { f("routine-7"); })
		go runTimed(func() { f("routine-8"); })
	})

	time.Sleep(time.Second) // better to use waitGroups
	fmt.Println("done")
	end := time.Now()
	fmt.Println("Total time: %v", end.Sub(start))
}

func f(id string) {
	for i := 0; i < 3; i++ {
		fmt.Println(id, ":", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func runTimed(fn func()) {
	start := time.Now()
	fn()
	end := time.Now()
	fmt.Println("Elapsed time: %v", end.Sub(start))
}