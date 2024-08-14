package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := NewServer("127.0.0.1", "8081", newKeycloak())
	fmt.Println("Listening on port 8081")

	s.listen()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
