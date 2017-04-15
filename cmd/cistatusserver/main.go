package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := configFromEnv()
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		log.Println("Exiting")
		os.Exit(1)
	}

	server := config.NewServer()
	err = http.ListenAndServe(config.HTTPAddress, server)
	if err != nil {
		log.Println(err)
	}

	sig := <-exitChan()
	log.Printf("Received signal: %v\n", sig)
	log.Println("Exiting")
}

func exitChan() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	return ch
}
