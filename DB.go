package main

import (
	"context"
	"log"

	"github.com/gordun209-hub/webapp/db"
)

func ConnectDB() (*db.PrismaClient, context.Context) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	return client, ctx
	// create a post
}

func DisconnectDB(client *db.PrismaClient) {
	if err := client.Prisma.Disconnect(); err != nil {
		panic(err)
	}
}
