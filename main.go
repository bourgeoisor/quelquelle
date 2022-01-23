package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
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
		Since: 0,
		Game: discordgo.Activity{
			Name: "top 10,000 anime",
			Type: discordgo.ActivityTypeListening,
		},
		Status: "wat.",
		AFK:    false,
	}

	err = dg.Open()
	if err != nil {
		log.Fatalln("fatal: error opening connection,", err)
	}

	log.Println("Bot has started.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "~ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "~help" {
		s.ChannelMessageSend(m.ChannelID, "I can do `~ping`, `~lay egg`, and `~rej`")
	}

	if m.Content == "~lay egg" {
		s.ChannelMessageSend(m.ChannelID, "_Lays an egg._")
	}

	rejPictures := [...]string{
		"https://i.imgur.com/90EpyFH.png",
		"https://i.imgur.com/gsiMusB.png",
		"https://i.imgur.com/IbHwHgp.png",
		"https://i.imgur.com/zGVIEek.png",
		"https://i.imgur.com/bIktemr.png",
		"https://i.imgur.com/KJeu9m9.png",
		"https://i.imgur.com/nHFM5MF.png",
		"https://i.imgur.com/uJeAmnJ.png",
		"https://i.imgur.com/ZEwgpbN.png",
		"https://i.imgur.com/MdRTOTN.png",
		"https://i.imgur.com/FL2JSPg.png",
	}

	if m.Content == "~rej" {
		s.ChannelMessageSendEmbed(m.ChannelID,
			&discordgo.MessageEmbed{
				Color: 15548997,
				Type:  "image",
				Image: &discordgo.MessageEmbedImage{URL: rejPictures[rand.Intn(len(rejPictures))]},
			})
	}
}

