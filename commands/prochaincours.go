package commands

import (
	"AssistantEtudiants/utils"
	"log"
	"os"
	"strings"
	"time"

	"github.com/apognu/gocal"
	"github.com/bwmarrin/discordgo"
)

func NextCourseCommand(discord *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "prochaincours" {
		return
	}

	data, err := os.Open("ADECal.ics")
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	p := gocal.NewParser(data)

	now := time.Now()

	if err := p.Parse(); err != nil {
		log.Fatal(err)
	}

	var next *gocal.Event

	for _, event := range p.Events {
		if event.Start.After(now) {
			if next == nil || event.Start.Before(*next.Start) {
				next = &event // On link "next" à l'adresse d'"event"
			}
		}
	}

	var start *time.Time = next.Start
	until := start.Sub(time.Now()).String()

	// Reformattage complet du until pour que ce soit joli lors de l'envoi
	replacer := strings.NewReplacer(
		"h", "h ",
		"m", "m ",
	)

	until = strings.TrimSpace(replacer.Replace(until))
	until = strings.Split(until, ".")[0] + "s" // On retire les décimales (ok j'ai pas arrondi mais pas grave c'est que des secondes)

	// Péparer l'envoi

	course := strings.Split(strings.ReplaceAll(next.Description, "\\n", "\n"), "\n")
	course_name := course[2]
	course_group := course[3]
	course_teacher := course[4]

	discord.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       ":clock9:Prochain cours dans " + until,
					Description: ":books: **" + course_name + "**\n:teacher: Professeur : __" + course_teacher + "__" + "\n:student: Pour les : " + course_group,
					Timestamp:   "",
					Color:       utils.DefaultEmbedColor,
				},
			},
		},
	})

}
