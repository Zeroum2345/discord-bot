package main

import (
	"fmt"
	"main/commands"
	"main/commands/commandsInfo"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")
	appID := os.Getenv("ID")

	app, _ := discordgo.New("Bot " + token)

	loadCommands(app, appID)

	app.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if interaction.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if interaction.ApplicationCommandData().Name == "walk" {
			commands.Walk(session, interaction)
		}

		if interaction.ApplicationCommandData().Name == "help" {
			commands.Help(session, interaction)
		}

		if interaction.ApplicationCommandData().Name == "test" {
			commands.Test(session, interaction)
		}
	})

	app.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Bot Online: %s", r.User.String())
	})

	app.Open()
	defer app.Close()

	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func loadCommands(app *discordgo.Session, appID string) {
	list := commandsInfo.GetList()
	commandsList := []*discordgo.ApplicationCommand{}

	for _, cmd := range list {
		applicationCommand := cmd.ApplicationCommand
		commandsList = append(commandsList, applicationCommand)
	}

	_, err := app.ApplicationCommandBulkOverwrite(appID, "", commandsList)
	if err != nil {
		fmt.Println(err)
		return
	}
}
