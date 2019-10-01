package main

import (
	"github.com/katsuomi/gin-gorm-twitter-api/db"
	"github.com/katsuomi/gin-gorm-twitter-api/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
