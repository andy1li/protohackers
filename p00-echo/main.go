package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/andy1li/protohackers/p00-echo/echo"
)

func main() {
	echoServer := echo.NewEchoServer()
	go func() {
		if err := echoServer.Start(); err != nil {
			log.Fatalf("❌ Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("✅ Echo server started. Press Ctrl+C to stop.")
	<-sigChan
	fmt.Println("\n👋 Shutting down...")
}
