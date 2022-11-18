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
	"GetLastMessage":   GetLastMessage,
}

var (
	prefix  = "!"
	appID   = "1038463049494958220"
	GuildID = "1038468460654645310"
)

type DiscordBot struct {
	discord  *discordgo.Session
	guild    *discordgo.Guild
	User     *discordgo.User
	Channels []*discordgo.Channel
	Members  []*discordgo.Member
}

type Members struct {
	name     string
	memberID string
	point    int
	level    int
}

// For getting members as Member struct
func (dg *DiscordBot) InitializeMembers() []*Members {
	members := make([]*Members, 0)
	for _, member := range dg.Members {
		members = append(members, &Members{
			name:     member.User.Username,
			memberID: member.User.ID,
			point:    0,
			level:    0,
		})
	}
	return members
}

// Get members as Member format
func (dg *DiscordBot) GetMembers() (*Members, error) {
	// get the members

	return &Members{
		name:     "",
		memberID: "",
		point:    0,
		level:    0,
	}, nil
}

func (dg *DiscordBot) IncreasePoint(name string) {
	member := dg.findMember(name)
	fmt.Println(member)
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
	discord, err := discordgo.New("Bot " + getToken())
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
	members, err := discord.GuildMembers(GuildID, "", 1000)
	if err != nil {
		return nil, err
	}

	channels, err := discord.GuildChannels(GuildID)
	if err != nil {
		return nil, err
	}

	return &DiscordBot{
		discord:  discord,
		guild:    guild,
		User:     user,
		Channels: channels,
		Members:  members,
	}, nil
}

func (u *User) String() string {
	return fmt.Sprintf("%s %s %d %d", u.name, u.ID, u.point, u.level)
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
