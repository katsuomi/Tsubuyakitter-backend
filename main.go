package main

import (
	"./db"
	"./server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}