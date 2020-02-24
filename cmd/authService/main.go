package main

import (
	"flag"
	"log"
	"promotioner/cmd/authService/manager"
	authServer "promotioner/cmd/authService/server"
)

var port = flag.String("port", ":8001", "server port")
var redisAddr = flag.String("redisAddr", "localhost:6379", "redis addr")
var redisPassword = flag.String("redisPassword", "", "redis password")
var redisDb = flag.Int("redisDb", 0, "redis db")

func main() {
	redisManager := manager.NewRedisManager(*redisAddr, *redisPassword, *redisDb)
	server := authServer.NewServer(*port, redisManager)
	log.Fatalln(server.Run())
}
