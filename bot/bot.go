package bot

import (
	"botdiscord/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "festa" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "só se for em ipanema meu amor")
	}

	if m.Content == "Festa" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "só se for em ipanema meu amor")
	}

	if m.Content == "palminha" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Palminha! Palminha! É palma na asinha!")
	}

	if m.Content == "Palminha" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Palminha! Palminha! É palma na asinha!")
	}

	if m.Content == "ih" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Meteu essa?")
	}

	if m.Content == "Ih" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Meteu essa?")
	}

	if m.Content == "bluzin" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "simplesmente o melhor bot ever")
	}

	if m.Content == "Bluzin" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "simplesmente o melhor bot ever")
	}

	if m.Content == "bom dia" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "bom dia meu")
	}

	if m.Content == "Bom dia" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "bom dia meu")
	}

	if m.Content == "boa tarde" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "boa tarde")
	}

	if m.Content == "Boa tarde" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "boa tarde")
	}

	if m.Content == "boa noite" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "boa noite")
	}

	if m.Content == "Boa noite" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "boa noite")
	}

	if m.Content == "amém" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Amém!")
	}

	if m.Content == "Amém" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Amém!")
	}

	if m.Content == "versículo" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `por enquanto eu só sei Deuteronômio 6:5: Amarás o SENHOR, teu Deus, com todo o coração, com toda a tua alma e com todas as tuas forças.`)
	}

	if m.Content == "Versículo" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `por enquanto eu só sei Deuteronômio 6:5: Amarás o SENHOR, teu Deus, com todo o coração, com toda a tua alma e com todas as tuas forças.`)
	}

	if m.Content == "versiculo" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `por enquanto eu só sei Deuteronômio 6:5: Amarás o SENHOR, teu Deus, com todo o coração, com toda a tua alma e com todas as tuas forças.`)
	}

	if m.Content == "versiculo" {
		_, _ = s.ChannelMessageSend(m.ChannelID, `por enquanto eu só sei Deuteronômio 6:5: Amarás o SENHOR, teu Deus, com todo o coração, com toda a tua alma e com todas as tuas forças.`)
	}

}
