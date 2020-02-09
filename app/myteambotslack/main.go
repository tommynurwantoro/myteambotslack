package main

import (
	"fmt"

	"github.com/bot/myteambotslack/app"
	"github.com/bot/myteambotslack/app/method"
	"github.com/bot/myteambotslack/app/utility"
	"github.com/nlopes/slack"
)

func main() {
	bot := app.Bot
	if bot == nil {
		panic("BOT ERROR")
	}

	app.RTM = bot.NewRTM()
	go app.RTM.ManageConnection()

	for msg := range app.RTM.IncomingEvents {
		switch event := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Println("Infos:", event.Info)
			fmt.Println("Connection counter:", event.ConnectionCount)

		case *slack.MessageEvent:
			if utility.IsFromBot(event, app.RTM) {
				continue
			}

			m := method.NewMethod(event)

			if !utility.IsBotMentioned(event, app.RTM) && !utility.IsDirectMessage(event) {
				go m.HandleMessage()
				continue
			}

			go m.HandleCommand()

		case *slack.MemberJoinedChannelEvent:
			app.RTM.SendMessage(app.RTM.NewOutgoingMessage(utility.GreetingNewJoinedUser(event.User), event.Channel))

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", event.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", event.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:
		}
	}
}
