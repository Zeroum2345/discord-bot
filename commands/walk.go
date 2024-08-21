package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Walk(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "You found something!",
					Color: 1,
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							CustomID: "fight",
							Label:    "Fight",
							Style:    discordgo.DangerButton,
						},
						discordgo.Button{
							CustomID: "run",
							Label:    "Run",
							Style:    discordgo.DangerButton,
						},
					},
				},
			},
		},
	})
}
