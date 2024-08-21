package commands

import (
	"main/commands/commandsInfo"

	"github.com/bwmarrin/discordgo"
)

func Help(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	member := interaction.Member
	userID := ""
	contentTxt := ""

	if member == nil {
		userID = interaction.User.ID
		contentTxt = "There is the command list: "
	} else {
		userID = member.User.ID
		contentTxt = "Command list sent to your DM!"
	}

	channel, _ := session.UserChannelCreate(userID)

	session.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed{
		Title:  "**No worries! Here is the list of all commands and their respective descriptions!**",
		Color:  0x971E4E,
		Fields: listCommandsAsFields(),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Note: only commands that you have permission to use are shown",
		}},
	)

	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: contentTxt,
		},
	})
}

func listCommandsAsFields() []*discordgo.MessageEmbedField {
	// TODO: Read the user permission (for later implementation)
	list := commandsInfo.GetList()
	helpList := []*discordgo.MessageEmbedField{}

	for _, cmd := range list {

		newCmd := &discordgo.MessageEmbedField{
			Name:  "/" + cmd.ApplicationCommand.Name,
			Value: cmd.DetailedDescription + "\n",
		}

		helpList = append(helpList, newCmd)
	}

	return helpList
}
