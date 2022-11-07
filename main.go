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

func initializeMembers(dg *DiscordBot) (*Members, error) {
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
	err = discordBot.Start()
	if err != nil {
		fmt.Println("error starting discord bot,", err)
		return
	}
	members, err := initializeMembers(discordBot)
	if err != nil {
		fmt.Println("error getting members,", err)
		return
	}
	client, _ := ConnectDB()
	// put members to db
	for _, member := range members.members {
		_, errr := client.User.CreateOne(
			db.User.Name.Set(member.User.Username),
		).Exec(context.Background())
		if errr != nil {
			fmt.Println("error putting member to db,", err)
			return
		}
	}
	// get users from db
	users, err := client.User.FindMany().Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user.Score)
	}
	// change users score
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
