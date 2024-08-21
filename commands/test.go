package commands

import (
	"main/item/items"

	"github.com/bwmarrin/discordgo"
)

func Test(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

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
		Title:  "**TESTING**",
		Color:  0xffffff,
		Fields: listInfoAsFields(),
	})

	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: contentTxt,
		},
	})
}

func listInfoAsFields() []*discordgo.MessageEmbedField {
	item := items.DragonClaw
	testList := []*discordgo.MessageEmbedField{}

	testList = append(testList, &discordgo.MessageEmbedField{
		Name:  "Name",
		Value: item.Name,
	})

	testList = append(testList, &discordgo.MessageEmbedField{
		Name:  "Rarity",
		Value: item.LootRarity,
	})

	return testList
}
