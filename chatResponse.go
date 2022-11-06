package main

import "github.com/bwmarrin/discordgo"

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
}
