package repository

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/bot/myteambotslack/app/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func FindCutomCommand(commandID int64) *models.CustomCommand {
	command, err := models.CustomCommands(qm.Where("id = ?", commandID)).OneG()
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	return command
}

func GetAllCustomCommandsByChannelID(channelID string) []*models.CustomCommand {
	commands, err := models.CustomCommands(qm.Where("channel_id = ?", channelID), qm.OrderBy("created_at")).AllG()
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	return commands
}

func InsertCustomCommand(channelID string, com, message string) {
	var command models.CustomCommand
	command.ChannelID = channelID
	command.Command = com
	command.Message = message

	err := command.InsertG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

func UpdateCustomCommand(channelID string, sequence int, message string) {
	commands := GetAllCustomCommandsByChannelID(channelID)

	for i, c := range commands {
		if i == sequence-1 {
			c.Message = message
			c.UpdateG(boil.Infer())
		}
	}
}

func DeleteCustomCommand(channelID string, sequences []string) bool {
	successToDelete := false
	commands := GetAllCustomCommandsByChannelID(channelID)

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		for i, c := range commands {
			if i+1 == sequence {
				c.DeleteG()
				successToDelete = true
				break
			}
		}
	}

	return successToDelete
}

func GenerateCustomCommands(commands []*models.CustomCommand) string {
	var buffer bytes.Buffer

	for i, command := range commands {
		buffer.WriteString(fmt.Sprintf("%d. %s\n", i+1, command.Command))
	}

	return buffer.String()
}
