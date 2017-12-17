package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	api := slack.New(os.Getenv("API_TOKEN"))

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello World", ev.Channel))
			}
		}
	}
}
