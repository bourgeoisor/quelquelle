package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	helpStrings := []string{
		"`~ping`: test bot's reactivity",
		"`~rej`: draws a random photo of Rej",
	}

	if m.Content == "~help" {
		s.ChannelMessageSendEmbed(m.ChannelID,
			&discordgo.MessageEmbed{
				Type:        "rich",
				Title:       "Quelquelle's guide",
				Description: strings.Join(helpStrings, "\n"),
			})
	}

	if m.Content == "~ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "~rej" {
		rejPictures := []string{
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
		randomIndex := rand.Intn(len(rejPictures))

		s.ChannelMessageSendEmbed(m.ChannelID,
			&discordgo.MessageEmbed{
				Color: 15548997,
				Type:  "image",
				Image: &discordgo.MessageEmbedImage{URL: rejPictures[randomIndex]},
				Footer: &discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf("Rej Picture Pack - Card %d/%d", randomIndex+1, len(rejPictures)),
				},
			})
	}
}
