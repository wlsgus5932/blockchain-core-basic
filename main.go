package main

import (
	"blockchain-core/app"
	"blockchain-core/config"
	"flag"
)

var (
	configFlag = flag.String("environment", "./environment.toml", "environment toml file not found")
	difficulty = flag.Int("difficulty", 12, "difficulty err")
)

func main() {
	flag.Parse()

	c := config.NewConfig(*configFlag)
	app.NewApp(c)
}
