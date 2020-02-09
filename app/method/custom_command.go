package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambotslack/app/utility"
	"github.com/bot/myteambotslack/app/utility/repository"
)

func SaveCustomCommandGroup(channelID string, username, args string) string {
	parameter := utility.GetArgsParameter(args)
	split := strings.Split(parameter, "][")

	if len(split) < 2 {
		return utility.InvalidParameter()
	}

	repository.InsertCustomCommand(channelID, split[0], split[1])

	return utility.SuccessInsertData()
}

func ListCustomCommandGroup(channelID string, username string) string {
	customCommands := repository.GetAllCustomCommandsByChannelID(channelID)

	if len(customCommands) == 0 {
		return utility.CustomCommandNotFound()
	}

	return fmt.Sprintf("Ini list command tim kamu:\n%s", repository.GenerateCustomCommands(customCommands))
}

func UpdateCustomCommandGroup(channelID string, username, args string) string {
	parameter := utility.GetArgsParameter(args)
	split := strings.Split(parameter, "][")

	if len(split) < 2 {
		return utility.InvalidParameter()
	}

	sequence, err := strconv.Atoi(split[0])
	if err != nil {
		return utility.InvalidParameter()
	}

	repository.UpdateCustomCommand(channelID, sequence, split[1])

	return utility.SuccessUpdateData()
}

func DeleteCustomCommandGroup(channelID string, username, args string) string {
	parameter := utility.GetArgsParameter(args)

	sequences := strings.Split(parameter, " ")
	repository.DeleteCustomCommand(channelID, sequences)

	return utility.SuccessUpdateData()
}

func RespondCustomCommandGroup(channelID string, args string) string {
	commands := repository.GetAllCustomCommandsByChannelID(channelID)
	if commands != nil {
		for _, c := range commands {
			if strings.Contains(args, c.Command) {
				return c.Message
			}
		}
	}

	return ""
}
