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
	split := strings.Split(m.message.Text, " ")

	if len(split) < 2 {
		app.RTM.SendMessage(app.RTM.NewOutgoingMessage(utility.InvalidCommand(), m.message.Channel))
		return
	}

	splitCommand := split[1]
	response := ""

	switch splitCommand {
	case command.Start().Name:
		response = utility.Start()

	case command.Help().Name:
		response = utility.Help(repository.GenerateAllCommands())

	case command.TitipReview().Name:
		response = TitipReview(m.message.Channel, m.message.Text)

	case command.AntrianReview().Name:
		response = AntrianReview(m.message.Channel)

	case command.SudahDireview().Name:
		response = SudahDireview(m.message.Channel, m.message.User, m.message.Text, false)

	case command.SudahDireviewSemua().Name:
		response = SudahDireview(m.message.Channel, m.message.User, m.message.Text, true)

	case command.TambahUserReview().Name:
		response = TambahUserReview(m.message.Channel, m.message.Text)

	case command.HapusReview().Name:
		response = HapusReview(m.message.Channel, m.message.Text)

	case command.SiapQA().Name:
		response = SiapQA(m.message.Channel, m.message.Text)

	case command.AntrianQA().Name:
		response = AntrianQA(m.message.Channel)

	case command.SudahDites().Name:
		response = SudahDites(m.message.Channel, m.message.Text)

	case command.SimpanCustomCommand().Name:
		response = SimpanCustomCommand(m.message.Channel, m.message.User, m.message.Text)

	case command.ListCustomCommand().Name:
		response = ListCustomCommand(m.message.Channel, m.message.User)

	case command.UbahCustomCommand().Name:
		response = UbahCustomCommand(m.message.Channel, m.message.User, m.message.Text)

	case command.HapusCustomCommand().Name:
		response = HapusCustomCommand(m.message.Channel, m.message.User, m.message.Text)
	}

	app.RTM.SendMessage(app.RTM.NewOutgoingMessage(response, m.message.Channel))
}

func (m *Method) HandleMessage() {
	response := RespondCustomCommand(m.message.Channel, m.message.Text)
	if response != "" {
		app.RTM.SendMessage(app.RTM.NewOutgoingMessage(response, m.message.Channel))
	}
}
