package main

import (
	"github.com/katsuomi/gin-like-twitter-api/db"
	"github.com/katsuomi/gin-like-twitter-api/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
