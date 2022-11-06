package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix  = "!"
	token   string
	appID   = "1038463049494958220"
	guildID = "1038468460654645310"
)

type Config struct {
	Prefix        string `json:"prefix"`
	ServiceURL    string `json:"service_url"`
	BotToken      string `json:"bot_token"`
	OwnerID       string `json:"owner_id"`
	UseSharding   bool   `json:"use_sharding"`
	ShardID       int    `json:"shard_id"`
	ShardCount    int    `json:"shard_count"`
	DefaultStatus string `json:"default_status"`
}

type DiscordBot struct {
	discord  *discordgo.Session
	guild    *discordgo.Guild
	User     *discordgo.User
	Channels []*discordgo.Channel
	Message  *discordgo.Message
	Youtube  *Config
}

func NewDiscordBot() (*DiscordBot, error) {
	token := getToken()
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	user, err := discord.User("@me")
	if err != nil {
		return nil, err
	}
	guild, err := discord.Guild(guildID)
	if err != nil {
		return nil, err
	}
	channels, err := discord.GuildChannels(guildID)
	if err != nil {
		return nil, err
	}

	return &DiscordBot{discord: discord, User: user, guild: guild, Channels: channels}, nil
}

func (d *DiscordBot) Start() error {
	d.discord.AddHandler(createChannel)
	d.discord.AddHandler(deleteChannel)
	d.discord.AddHandler(CreatePingOrPong)
	d.discord.AddHandler(changeAvatar)
	d.discord.AddHandler(clearMessages)
	err := d.discord.Open()
	if err != nil {
		return err
	}
	return nil
}

func (d *DiscordBot) Stop() error {
	return d.discord.Close()
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

func getToken() string {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		fmt.Println("No token provided")
		os.Exit(1)
	}
	return token
}
