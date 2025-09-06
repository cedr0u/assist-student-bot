package commands

import "github.com/bwmarrin/discordgo"

func UWUCommand(discord *discordgo.Session, msg *discordgo.InteractionCreate) {
	if msg.ApplicationCommandData().Name != "uwu" {
		return
	}

	discord.InteractionRespond(msg.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "uwu :3",
		},
	})

}
