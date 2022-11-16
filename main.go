package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"

	"github.com/gordun209-hub/webapp/db"
)

type Members struct {
	members []*discordgo.Member
}

func (m *Members) String() string {
	var s string
	for _, member := range m.members {
		s += "\n" + member.User.Username + "\n discriminator :  " + member.User.Discriminator + " \n ID : " + member.User.ID
	}
	return s
}

type Discord struct {
	*discordgo.Session
}

func (dg *DiscordBot) initializeMembers() (*Members, error) {
	// get the members
	members, err := dg.discord.GuildMembers(guildID, "", 1000)
	return &Members{members: members}, err
}

func printMembers(m *Members) {
	for _, member := range m.members {
		fmt.Println(member.User.Username)
	}
}

func findMember(m *Members, name string) *discordgo.Member {
	for _, member := range m.members {
		if member.User.Username == name {
			return member
		}
	}
	return nil
}

func main() {
	discordBot, err := NewDiscordBot()
	if err != nil {
		fmt.Println("error creating discord bot,", err)
		return
	}

	discordBot.Start()
	members, err := discordBot.initializeMembers()
	if err != nil {
		fmt.Println("error getting members,", err)
		return
	}
	fmt.Println("members:", members)
	client, _ := ConnectDB()

	// get users from db

	// disconnect db
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	eventLoop(discordBot.discord)
}

func ConnectDB() (*db.PrismaClient, context.Context) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	return client, ctx
	// create a post
}
