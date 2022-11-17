package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var events = map[string]func(*discordgo.Session, *discordgo.MessageCreate){
	"createChannel":    CreateChannel,
	"deleteChannel":    DeleteChannel,
	"CreatePingOrPong": CreatePingOrPong,
	"changeAvatar":     ChangeAvatar,
	"clearMessages":    ClearMessages,
}

var (
	prefix  = "!"
	token   = "MTAzODQ2MzA0OTQ5NDk1ODIyMA.GBNLBN.k0s7ZfFRmlw0eT_ZPk14D_USECtTHamtgg7nd8"
	appID   = "1038463049494958220"
	GuildID = "1038468460654645310"
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

type Members struct {
	members []*discordgo.Member
	point   int
	level   int
}

func (dg *DiscordBot) GetMembers() (*Members, error) {
	// get the members
	members, err := dg.discord.GuildMembers(GuildID, "", 1000)
	return &Members{members: members}, err
}

func (m *Members) String() string {
	var s string
	for _, member := range m.members {
		s += "\n" + member.User.Username + "\n discriminator :  " + member.User.Discriminator + " \n ID : " + member.User.ID
	}
	return s
}

func (d *DiscordBot) EventLoop() {
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	d.discord.Close()
}

func NewDiscordBot() (*DiscordBot, error) {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	user, err := discord.User("@me")
	if err != nil {
		return nil, err
	}
	guild, err := discord.Guild(GuildID)
	if err != nil {
		return nil, err
	}
	channels, err := discord.GuildChannels(GuildID)
	if err != nil {
		return nil, err
	}

	return &DiscordBot{discord: discord, User: user, guild: guild, Channels: channels}, nil
}

func (d *DiscordBot) Start() {
	for _, f := range events {
		d.discord.AddHandler(f)
	}
	err := d.discord.Open()
	if err != nil {
		panic(1)
	}
}

func (d *DiscordBot) Stop() error {
	return d.discord.Close()
}

func (d *DiscordBot) findMember(name string) *discordgo.Member {
	for _, member := range d.guild.Members {
		if member.User.Username == name {
			return member
		}
	}
	return nil
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
