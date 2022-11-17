package main

import (
	"fmt"
)

func main() {
	/* Create a new DiscordBot */
	discordBot, err := NewDiscordBot()
	if err != nil {
		fmt.Println("error creating discord bot,", err)
		return
	}

	discordBot.Start()
	/* Connect to db*/
	client, _ := ConnectDB()

	DisconnectDB(client)

	fmt.Println(discordBot.GetMembers())
	discordBot.EventLoop()
}
