package main

import (
	"os"
	"os/signal"
	"syscall"

	"log"

	"net/http"

	"tantalic.com/cistatus"
	"tantalic.com/cistatus/gitlab"
)

func main() {
	config, err := configFromEnv()
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		log.Println("Exiting")
		os.Exit(1)
	}

	fetcher := gitlab.NewClient(config.GitLabBaseURL, config.GitLabAPIToken)
	server := cistatus.Server{
		Fetcher:      fetcher,
		Logger:       cistatus.NewVerboseLogger(os.Stdout),
		JWTAlgorithm: config.JWTAlgorithm,
		JWTSecret:    config.JWTSecret,
	}
	server.StartFetching(config.GitLabRefreshInterval)

	err = http.ListenAndServe(config.HTTPAddress, &server)
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
