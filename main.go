package main

import (
	"golang-log-server/internals/app"
	"golang-log-server/internals/cfg"
	"log"
	"os"
	"os/signal"
)

func main() {
	config := cfg.NewConfig()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	server := app.NewAppServer(config)

	go func() {
		oscall := <-ch
		log.Println("system call:", oscall)
		server.Shutdown()
	}()

	server.Serve()
}
