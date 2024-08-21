package commandsInfo

import (
	"github.com/bwmarrin/discordgo"
)

type CommandInfo struct {
	*discordgo.ApplicationCommand
	DetailedDescription string
	CommandPermission   int
}

var commandsList = []CommandInfo{
	{
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "walk",
			Description: "Walk through the fantasy world",
		},
		DetailedDescription: "Walking means that you can explore and find valuable treasures, and creatures, after that you will be able to interact with them for achievements, powerful items and rewards",
		CommandPermission:   0,
	},
	{
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "help",
			Description: "List of all commands with detailed information",
		},
		DetailedDescription: "Lists all the commands that you have permission to use, showing its descriptions with more detailed content",
		CommandPermission:   0,
	},
	{
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "test",
			Description: "Debugging command",
		},
		DetailedDescription: "Dev only",
		CommandPermission:   0,
	},
}

func GetList() []CommandInfo {
	return commandsList
}
