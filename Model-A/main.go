package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

const profileFile = "profile.mem"

func startMemoryProfiling() {
	f, err := os.Create(profileFile)
	if err != nil {
		log.Fatal("could not create profile file:", err)
	}
	pprof.WriteHeapProfile(f)
	f.Close()
	fmt.Println("Memory profile written to", profileFile)
}

func main() {
	// ... existing code ...

	startMemoryProfiling()

	// ... code to be profiled ...

	doSomeWork()

	startMemoryProfiling()
}

func doSomeWork() {
	// Simulate memory-intensive operations
	var largeSlice = make([]int, 1_000_000)
	for i := range largeSlice {
		largeSlice[i] = i
	}
	time.Sleep(2 * time.Second)
}
