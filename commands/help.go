package commands

import (
	"AssistantEtudiants/utils"

	"github.com/bwmarrin/discordgo"
)

func HelpCommand(discord *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "help" {

		embedfield := ""
		for i, cmd := range utils.Commands {
			sep := ""
			if len(utils.Commands)-1 != i {
				sep = "\n\n"
			}
			embedfield += "**/" + cmd.Name + "**\n→ " + cmd.Description + sep
		}

		discord.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "",
				Embeds: []*discordgo.MessageEmbed{
					&discordgo.MessageEmbed{
						Title:       ":books: Commandes disponibles",
						Description: "Pour obtenir plus de détail sur une commande, faites `/help [commande]`",
						Color:       utils.DefaultEmbedColor,
						Footer: &discordgo.MessageEmbedFooter{
							Text: "Assistant Étudiants, v1.0 - GOLANG ❤️",
						},
						Image:     nil,
						Thumbnail: nil,
						Video:     nil,
						Provider:  nil,
						Author:    nil,
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:   "Commandes",
								Value:  embedfield,
								Inline: false,
							},
						},
					},
				},
			},
		})
	}
}
