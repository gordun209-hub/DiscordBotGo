package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	base64img   string
	contentType string
)

var (
	AvatarFile string
	AvatarURL  string
	prefix     = "!"
	token      = "MTAzODQ2MzA0OTQ5NDk1ODIyMA.GMEDuY.lfL6ctUI3syjt01CmOWriIQdH79Zo08HszmlXE"
	appID      = "1038463049494958220"
	guildID    = "1038468460654645310"
)

func initializeDiscord() *discordgo.Session {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating discord session", err)
		return nil
	}
	return discord
}

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
	discord     *discordgo.Session
	guild       *discordgo.Guild
	User        *discordgo.User
	TextChannel *discordgo.Channel
	Message     *discordgo.Message
	Youtube     *Config
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
