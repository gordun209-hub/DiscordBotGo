package main

// Database sqlite3 //
import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	/* Create a new DiscordBot */
	discordBot, err := NewDiscordBot()
	if err != nil {
		log.Fatal(err)
	}

	/* Start the DiscordBot */
	discordBot.Start()

	msg := discordBot.GetLastMessage()
	if msg != nil {
		fmt.Println(msg.Content)
	}
	fmt.Println(msg)

	// write to the database

	discordBot.EventLoop()
}
