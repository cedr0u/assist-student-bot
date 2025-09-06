package commands

import (
	"AssistantEtudiants/utils"

	"github.com/bwmarrin/discordgo"
)

func SourceCodeCommand(discord *discordgo.Session, msg *discordgo.InteractionCreate) {
	if msg.ApplicationCommandData().Name != "sourcecode" {
		return
	}
	discord.InteractionRespond(msg.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       ":computer: Code source disponible sur GitHub",
					Description: "Vous pouvez retrouver le **code source** de l'<@" + discord.State.User.ID + ">" + " en cliquant [ici](https://github.com/cedr0u/assist-student-bot)",
					Color:       utils.DefaultEmbedColor,
				},
			},
		},
	})
}
