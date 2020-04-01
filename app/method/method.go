package method

import (
	"strings"

	"github.com/bot/myteambotslack/app"
	"github.com/bot/myteambotslack/app/utility"
	"github.com/bot/myteambotslack/app/utility/repository"
	"github.com/nlopes/slack"
)

type Method struct {
	message *slack.MessageEvent
}

func NewMethod(message *slack.MessageEvent) *Method {
	return &Method{message: message}
}

func (m *Method) HandleCommand() {
	var command *repository.Command
	var responses []string
	split := strings.Split(m.message.Text, " ")

	if len(split) < 2 {
		app.RTM.SendMessage(app.RTM.NewOutgoingMessage(utility.InvalidCommand(), m.message.Channel))
		return
	}

	splitCommand := split[1]

	switch splitCommand {
	case command.Start().Name:
		responses = append(responses, utility.Start())

	case command.Help().Name:
		responses = append(responses, utility.Help(repository.GenerateAllCommands()))

	case command.TitipReview().Name:
		responses = append(responses, TitipReview(m.message.Channel, m.message.Text))

	case command.AntrianReview().Name:
		responses = append(responses, AntrianReview(m.message.Channel)...)

	case command.SudahDireview().Name:
		responses = append(responses, SudahDireview(m.message.Channel, m.message.User, m.message.Text, false))

	case command.SudahDireviewSemua().Name:
		responses = append(responses, SudahDireview(m.message.Channel, m.message.User, m.message.Text, true))

	case command.TambahUserReview().Name:
		responses = append(responses, TambahUserReview(m.message.Channel, m.message.Text))

	case command.HapusReview().Name:
		responses = append(responses, HapusReview(m.message.Channel, m.message.Text))

	case command.SiapQA().Name:
		responses = append(responses, SiapQA(m.message.Channel, m.message.Text))

	case command.AntrianQA().Name:
		responses = append(responses, AntrianQA(m.message.Channel)...)

	case command.SudahDites().Name:
		responses = append(responses, SudahDites(m.message.Channel, m.message.Text))

	case command.SimpanCustomCommand().Name:
		responses = append(responses, SimpanCustomCommand(m.message.Channel, m.message.User, m.message.Text))

	case command.ListCustomCommand().Name:
		responses = append(responses, ListCustomCommand(m.message.Channel, m.message.User))

	case command.UbahCustomCommand().Name:
		responses = append(responses, UbahCustomCommand(m.message.Channel, m.message.User, m.message.Text))

	case command.HapusCustomCommand().Name:
		responses = append(responses, HapusCustomCommand(m.message.Channel, m.message.User, m.message.Text))
	}

	for _, response := range responses {
		app.RTM.SendMessage(app.RTM.NewOutgoingMessage(response, m.message.Channel))
	}
}

func (m *Method) HandleMessage() {
	response := RespondCustomCommand(m.message.Channel, m.message.Text)
	if response != "" {
		app.RTM.SendMessage(app.RTM.NewOutgoingMessage(response, m.message.Channel))
	}
}
