package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/4726/twitch-chat-logger/app"
	"github.com/4726/twitch-chat-logger/config"
)

func main() {
	conf := config.Load("config_dev.yml")

	run := make(chan error, 1)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(c)

	s, err := app.NewServer(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	go func() {
		run <- s.Run()
	}()

	select {
	case sig := <-c:
		log.Print(sig.String())
	case err := <-run:
		log.Print(err)
	}
}
