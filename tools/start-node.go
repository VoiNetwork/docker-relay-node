package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	done := make(chan bool)

	var cmd *exec.Cmd
	var err error

	go func() {
		// Start algod
		cmd = exec.Command("/node/bin/algod", "-d", "/algod/data")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run() // Run is blocking
		if err != nil {
			log.Fatalf("Failed to start algod: %v", err)
		}

		// Signal that algod has finished
		done <- true
	}()

	// Wait for algod to start
	time.Sleep(5 * time.Second)

	// Execute catch-catchpoint
	cmd = exec.Command("/node/bin/catch-catchpoint")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute catch-catchpoint: %v", err)
	}

	<-done // Wait for algod to finish

}
