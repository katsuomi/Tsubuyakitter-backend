package main

import (
	"github.com/katsuomi/LikeTwitterApp-backend/db"
	"github.com/katsuomi/LikeTwitterApp-backend/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
