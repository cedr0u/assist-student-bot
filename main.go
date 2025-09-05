package main

import (
	"AssistantEtudiants/commands"
	"AssistantEtudiants/handlers"
	"AssistantEtudiants/utils"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	discord.Open()
	if err != nil {
		fmt.Println("error creating Discord session,", err)
	}

	setupBannedExpressions()

	discord.AddHandler(handlers.Automod)
	discord.AddHandler(commands.HelpCommand)

	/* Enregister les commandes */
	for _, value := range utils.Commands {
		fmt.Println("Enregistrement de la commande : /" + value.Name)
		discord.ApplicationCommandCreate(discord.State.User.ID, os.Getenv("GUILD_ID"), value)
	}

	/* ... */

	discord.UpdateGameStatus(0, "vous assister au quotiden ! (/help)")

	fmt.Println("Le bot est en route ! Ctrl+C pour l'arrêter.")

	// Programme l'extinction lors de l'intervention de l'administrateur
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Extinction du bot...")
	for _, value := range utils.Commands {
		fmt.Println("Désenregistrement de la commande : /" + value.Name)
		discord.ApplicationCommandDelete(discord.State.User.ID, os.Getenv("GUILD_ID"), value.ID)
	}
	discord.Close()
}

// Enregister dans la liste BannedExpressions l'entièreté des expressions bannies contenues dans le fichier "banned-expressions".
func setupBannedExpressions() {

	data, err := os.ReadFile("banned-expressions")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)                  // Ca c'est pour éviter les problèems inutiles avec les espaces et les tabulations
		if line == "" || strings.HasPrefix(line, "#") { // Si la ligne est vide ou qu'elle débute par # on ignore
			continue
		}
		line = strings.ToLower(line)
		utils.BannedExpressions = append(utils.BannedExpressions, line)
	}

}
