package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
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
	main1()
	// discordBot, err := NewDiscordBot()
	// if err != nil {
	// 	fmt.Println("error creating discord bot,", err)
	// 	return
	// }
	// err = discordBot.Start()
	// if err != nil {
	// 	fmt.Println("error starting discord bot,", err)
	// 	return
	// }
	// members, err := initializeMembers(discordBot)
	// printMembers(members)
	//
	// if err != nil {
	// 	fmt.Println("error getting members,", err)
	// 	return
	// }
	//
	// eventLoop(discordBot.discord)
}
