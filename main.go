package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Expected exactly one argument")
	}
	arg := os.Args[1]
	LeitstellenKennungServer := "test-leitstand"
	addrServer := ":80"
	LeitstellenKennungClient := "test-drehscheibe"
	addrClient := ":81"

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	switch arg {
	case "client":
		fmt.Println("Starting client...")
		server := NewVdvServer(LeitstellenKennungClient, addrClient)
		server.Start()
		server.Subscribe(LeitstellenKennungServer, "TestAboID")
		<-stop
		server.Stop()
	case "server":
		fmt.Println("Starting server...")
		server := NewVdvServer(LeitstellenKennungServer, addrServer)
		server.AddLeitstelle(LeitstellenKennungClient, addrClient)
		server.Start()
		<-stop
		server.Stop()
	default:
		fmt.Println("Unknown command ", arg)
	}
}
