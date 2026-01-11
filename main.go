package main

import (
	"log"
	"my-echo-app/config"
	route "my-echo-app/routes"
)

func main() {
	config.RedisInit()
	defer config.CloseRedis() // Ensure Redis closes on shutdown

	// 3. Start your Echo server
	log.Printf("🚀 Starting server on port %s", config.GetConfig().Port)

	e := route.InitHttp()

	e.Logger.Fatal(e.Start(":" + config.GetConfig().Port))
}
