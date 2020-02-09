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
		response = m.TitipReview()

	case command.AntrianReview().Name:
		response = m.AntrianReview()

	case command.SudahDireview().Name:
		response = m.SudahDireview()

	case command.SudahDireviewSemua().Name:
		response = m.SudahDireviewSemua()

	case command.TambahUserReview().Name:
		response = m.TambahUserReview()

	case command.SiapQA().Name:
		response = m.SiapQA()

	case command.AntrianQA().Name:
		response = m.AntrianQA()

	case command.SudahDites().Name:
		response = m.SudahDites()

	case command.SimpanCommand().Name:
		response = m.SimpanCommand()

	case command.ListCommand().Name:
		response = m.ListCommand()

	case command.UbahCommand().Name:
		response = m.UbahCommand()

	case command.HapusCommand().Name:
		response = m.HapusCommand()
	}

	app.RTM.SendMessage(app.RTM.NewOutgoingMessage(response, m.message.Channel))
}

func (m *Method) HandleMessage() {
	response := m.RespondAllText()
	if response != "" {
		app.RTM.SendMessage(app.RTM.NewOutgoingMessage(response, m.message.Channel))
	}
}

func (m *Method) TitipReview() string {
	if !utility.IsValidParameter(m.message.Text) {
		return utility.InvalidParameter()
	}

	return AddReview(m.message.Channel, m.message.Text)
}

func (m *Method) AntrianReview() string {
	return GetReviewQueue(m.message.Channel)
}

func (m *Method) SudahDireview() string {
	if !utility.IsValidParameter(m.message.Text) {
		return utility.InvalidParameter()
	}

	return UpdateDoneReview(m.message.Channel, m.message.User, m.message.Text, false)
}

func (m *Method) SudahDireviewSemua() string {
	if !utility.IsValidParameter(m.message.Text) {
		return utility.InvalidParameter()
	}

	return UpdateDoneReview(m.message.Channel, m.message.User, m.message.Text, true)
}

func (m *Method) TambahUserReview() string {
	if !utility.IsValidParameter(m.message.Text) {
		return utility.InvalidParameter()
	}

	return AddUserReview(m.message.Channel, m.message.Text)
}

func (m *Method) SiapQA() string {
	if !utility.IsValidParameter(m.message.Text) {
		return utility.InvalidParameter()
	}

	return UpdateReadyQA(m.message.Channel, m.message.Text)
}

func (m *Method) AntrianQA() string {
	return GetQAQueue(m.message.Channel)
}

func (m *Method) SudahDites() string {
	if !utility.IsValidParameter(m.message.Text) {
		return utility.InvalidParameter()
	}

	return UpdateDoneQA(m.message.Channel, m.message.Text)
}

func (m *Method) SimpanCommand() string {
	if !utility.IsValidParameter(m.message.Text) {
		return utility.InvalidParameter()
	}

	return SaveCustomCommandGroup(m.message.Channel, m.message.User, m.message.Text)
}

func (m *Method) ListCommand() string {
	return ListCustomCommandGroup(m.message.Channel, m.message.User)
}

func (m *Method) UbahCommand() string {
	if !utility.IsValidParameter(m.message.Text) {
		return utility.InvalidParameter()
	}

	return UpdateCustomCommandGroup(m.message.Channel, m.message.User, m.message.Text)
}

func (m *Method) HapusCommand() string {
	if !utility.IsValidParameter(m.message.Text) {
		return utility.InvalidParameter()
	}

	return DeleteCustomCommandGroup(m.message.Channel, m.message.User, m.message.Text)
}

func (m *Method) RespondAllText() string {
	return RespondCustomCommandGroup(m.message.Channel, m.message.Text)
}

// Private functions
func trimAndLower(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return text
}
