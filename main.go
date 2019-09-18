package main

import (
	"github.com/katsuomi/gin-gorm-practice/db"
	"github.com/katsuomi/gin-gorm-practice/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
