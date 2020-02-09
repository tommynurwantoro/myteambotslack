package utility

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

const (
	slackBotUser        = "USLACKBOT"
	userMentionFormat   = "<@%s>"
	directChannelMarker = "D"
)

func IsFromBot(event *slack.MessageEvent, rtm *slack.RTM) bool {
	info := rtm.GetInfo()
	return len(event.User) == 0 || event.User == slackBotUser || event.User == info.User.ID || len(event.BotID) > 0
}

func IsBotMentioned(event *slack.MessageEvent, rtm *slack.RTM) bool {
	info := rtm.GetInfo()
	return strings.Contains(event.Text, fmt.Sprintf(userMentionFormat, info.User.ID))
}

func IsDirectMessage(event *slack.MessageEvent) bool {
	return strings.HasPrefix(event.Channel, directChannelMarker)
}

func GetArgsByRegex(args string) string {
	rgx := regexp.MustCompile(`\[(.*)\]`)

	return rgx.FindString(args)
}

func IsValidParameter(args string) bool {
	splitArgs := strings.Split(args, " ")

	if len(splitArgs) < 3 {
		return false
	}

	if !strings.HasPrefix(splitArgs[2], "[") {
		return false
	}

	return true
}

func GetArgsParameter(args string) string {
	rgx := regexp.MustCompile(`\[(.*)\]`)
	parameter := rgx.FindString(args)

	return parameter[1 : len(parameter)-1]
}
