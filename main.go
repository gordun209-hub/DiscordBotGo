package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token   = "MTAzODQ2MzA0OTQ5NDk1ODIyMA.GMEDuY.lfL6ctUI3syjt01CmOWriIQdH79Zo08HszmlXE"
	appID   = "1038463049494958220"
	guildID = "1038468460654645310"
)

func initializeDiscord() *discordgo.Session {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating discord session", err)
		return nil
	}
	return discord
}

type DiscordBot struct {
	discord *discordgo.Session
}

func NewDiscordBot(token string) (*DiscordBot, error) {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &DiscordBot{discord: discord}, nil
}

func (d *DiscordBot) Start() error {
	d.discord.AddHandler(createChannel)
	d.discord.AddHandler(deleteChannel)
	d.discord.AddHandler(CreatePingOrPong)
	err := d.discord.Open()
	if err != nil {
		return err
	}
	return nil
}

func (d *DiscordBot) Stop() error {
	return d.discord.Close()
}

func main() {
	discordBot, err := NewDiscordBot(token)
	if err != nil {
		fmt.Println("error creating discord bot,", err)
		return
	}
	err = discordBot.Start()
	if err != nil {
		fmt.Println("error starting discord bot,", err)
		return
	}
	eventLoop(discordBot.discord)
}

func eventLoop(dc *discordgo.Session) {
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dc.Close()
}

func createChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "create channel") {
		// get the name of the channel
		channelName := strings.TrimPrefix(m.Content, "create channel")
		// create the channel
		channel, err := s.GuildChannelCreate(guildID, channelName, discordgo.ChannelTypeGuildText)
		if err != nil {
			fmt.Println("error creating channel,", err)
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Channel created!")

		// send a message to the channel
		s.ChannelMessageSend(channel.ID, "Hello World!")
	}
}

func deleteChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "delete channel") {
		// get the name of the channel
		channelName := strings.TrimPrefix(m.Content, "delete channel ")
		// get the channel id
		channels, err := s.GuildChannels(guildID)
		if err != nil {
			fmt.Println("error getting channels,", err)
			return
		}
		var channelID string
		for _, channel := range channels {
			cnn := strings.TrimSpace(channel.Name)
			if cnn == channelName {
				channelID = channel.ID
				break
			}
		}
		s.ChannelMessageSend(m.ChannelID, "Channel deleted!")
		// delete the channel
		_, err = s.ChannelDelete(channelID)
		if err != nil {
			fmt.Println("error deleting channel,", err)
			return
		}

	}
}
