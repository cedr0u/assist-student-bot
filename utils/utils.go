package utils

import "github.com/bwmarrin/discordgo"

var (
	BannedExpressions = []string{}

	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "help",
			Description: "Commande vous affichant les commandes existantes",
		},
		{
			Name:        "uwu",
			Description: "Ça vous répond un joli gentil message >.<",
		},
		{
			Name:        "sourcecode",
			Description: "Renvoie le lien vers le code source du bot Discord.",
		},
	}
)
