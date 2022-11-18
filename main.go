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

	db := newDB("db.txt")
	db.InitializeDB()
	// db.InitializeUsers(discordBot)
	/* Start the DiscordBot */
	discordBot.Start()

	membrs := discordBot.InitializeMembers()

	for _, member := range membrs {
		fmt.Println(member.level + 1)
	}

	// write to the database
	discordBot.EventLoop()
}
