package commands

import "github.com/bwmarrin/discordgo"

func HelpCommand(discord *discordgo.Session, i *discordgo.InteractionCreate) {
	discord.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Je ferai demain j'suis crevé.",
		},
	})
}
