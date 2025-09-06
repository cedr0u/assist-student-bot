package handlers

import (
	"AssistantEtudiants/utils"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Automod(discord *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == discord.State.User.ID {
		return
	}
	for _, exp := range utils.BannedExpressions {
		if strings.Compare(exp, msg.Content) == 0 {
			_, err := discord.ChannelMessageSendComplex(msg.ChannelID, &discordgo.MessageSend{
				Content: "",
				Reference: &discordgo.MessageReference{
					ChannelID: msg.ChannelID,
					MessageID: msg.ID,
				},
				Embed: &discordgo.MessageEmbed{
					Title:       ":warning: // Modération Automatique",
					Description: "Il se trouve que tu as envoyé un message potentiellement risqué. En effet, tu as dit quelque chose qui est susceptible de créer des tensions. Ce sont des choses à éviter pour la bonne entente avec tout le monde.",
					Color:       utils.DefaultEmbedColor,
					Footer: &discordgo.MessageEmbedFooter{
						Text: "Assistant Étudiants, v1.0 - GOLANG ❤️",
					},
				},
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}
