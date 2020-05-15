package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/4726/twitch-chat-logger/app"
	"github.com/4726/twitch-chat-logger/config"
	"github.com/spf13/cobra"
)

func serverCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Starts the server",
		Run:   serverCmdFunc,
	}
}

func serverCmdFunc(cmd *cobra.Command, args []string) {
	conf := config.Load(cfgFile)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(c)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := <-c
		log.Println("")
		log.Println(sig)
		cancel()
	}()

	s, err := app.NewServer(conf)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := s.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	s.Close()
}
