package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func CreatePingOrPong(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == "play" {
		s.ChannelMessageSend(m.ChannelID, "Playing!")
	}
	if m.Content == "stop" {
		s.ChannelMessageSend(m.ChannelID, "Stopped!")
	}
	if m.Content == "mkk" {
		s.ChannelMessageSend(m.ChannelID, "skm")
	}

	if m.Content == "good bot" {
		s.ChannelMessageSend(m.ChannelID, "eyw")
	}
}

var (
	base64img   string
	contentType string
	AvatarFile  string
	AvatarURL   string
)

func CreateChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!create channel") {
		// get the name of the channel
		channelName := strings.TrimPrefix(m.Content, "!create channel")
		// create the channel
		channel, err := s.GuildChannelCreate(GuildID, channelName, discordgo.ChannelTypeGuildText)
		if err != nil {
			fmt.Println("error creating channel,", err)
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Channel created!")

		// send a message to the channel
		s.ChannelMessageSend(channel.ID, "Hello World!")
	}
}

func DeleteChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!delete channel ") {
		// get the name of the channel
		channelName := strings.TrimPrefix(m.Content, "!delete channel ")
		// get the channel id
		channels, err := s.GuildChannels(GuildID)
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
		// delete the channel
		_, err = s.ChannelDelete(channelID)
		if err != nil {
			fmt.Println("error deleting channel,", err)
			return
		}

	}
}

func ChangeAvatar(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!change avatar") {
		// get the name of the avatar
		AvatarURL := strings.TrimPrefix(m.Content, "!change avatar ")
		if AvatarURL != "" {

			resp, err := http.Get(AvatarURL)
			if err != nil {
				fmt.Println("Error retrieving the file, ", err)
				return
			}
			defer func() {
				_ = resp.Body.Close()
			}()

			img, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading the response, ", err)
				return
			}

			contentType = http.DetectContentType(img)
			base64img = base64.StdEncoding.EncodeToString(img)
		}
		// change the avatar
		avatar := fmt.Sprintf("data:%s;base64,%s", contentType, base64img)
		_, err := s.UserUpdate("", avatar)
		if err != nil {
			fmt.Println("error changing avatar,", err)
			return
		}
		s.ChannelMessageSend(m.ChannelID, "Avatar changed!")
	}
}

func ClearMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!clear") {
		// get the number of messages to delete
		numMessages := strings.TrimPrefix(m.Content, "!clear")
		// get the messages
		messages, err := s.ChannelMessages(m.ChannelID, 100, "", "", "")
		if err != nil {
			fmt.Println("error getting messages,", err)
			return
		}
		// delete the messages
		for _, message := range messages {
			err = s.ChannelMessageDelete(m.ChannelID, message.ID)
			if err != nil {
				fmt.Println("error deleting message,", err)
				return
			}
		}
		fmt.Println("deleted", numMessages, "messages")
	}
}
