package main

import (
	"fmt"
	"os"
)

type User struct {
	name  string
	ID    string
	point int
	level int
}

type Users []User

type DB struct {
	users    []*User
	fileName string
}

func newDB(fileName string) *DB {
	return &DB{fileName: fileName}
}

func (db *DB) InitializeDB() {
	db.fileName = "db.txt"
	// create the file
	file, err := os.Create(db.fileName)
	if err != nil {
		fmt.Println("error creating file,", err)
		return
	}

	defer file.Close()
}

func (db *DB) AddUser(u *User) {
	db.users = append(db.users, u)
}

func (db *DB) InitializeUsers(dg *DiscordBot) {
	// read the file
	file, err := os.Open(db.fileName)
	if err != nil {
		fmt.Println("error opening file,", err)
		return
	}
	defer file.Close()
}

func (u *Users) IncreasePointOfUser(name string) {
	for _, user := range *u {
		if user.name == name {
			user.point = user.point + 1
		}
	}
}
