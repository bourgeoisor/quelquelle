package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("fatal: The BOT_TOKEN env is not set")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln("fatal: could not create Discord session,", err)
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{
			Name: "top 10,000 anime",
			Type: discordgo.ActivityTypeListening,
		},
	}

	err = dg.Open()
	if err != nil {
		log.Fatalln("fatal: error opening connection,", err)
	}

	log.Println("Bot has started.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = dg.Close()
	if err != nil {
		log.Fatalln("fatal: error closing connection,", err)
	}
}
