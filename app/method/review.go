package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambotslack/app/utility"
	"github.com/bot/myteambotslack/app/utility/repository"
)

// GetReviewQueue _
func GetReviewQueue(channelID string) string {
	reviews := repository.GetAllNeedReview(channelID)

	if len(reviews) == 0 {
		return "Gak ada antrian review nih 👍🏻"
	}

	return fmt.Sprintf("Ini antrian review tim kamu:\n%s", repository.GenerateContentReview(reviews))
}

// GetQAQueue _
func GetQAQueue(channelID string) string {
	reviews := repository.GetAllNeedQA(channelID)

	if len(reviews) == 0 {
		return "Gak ada antrian QA nih 👍🏻"
	}

	return fmt.Sprintf("Ini antrian QA tim kamu:\n%s", repository.GenerateContentReview(reviews))
}

// AddReview _
func AddReview(channelID string, args string) string {
	parameter := utility.GetArgsParameter(args)
	split := strings.Split(parameter, "][")

	if len(split) < 3 {
		return utility.InvalidParameter()
	}

	repository.InsertReview(split[0], split[1][1:len(split[1])-1], split[2], channelID)

	return utility.SuccessInsertData()
}

// UpdateDoneReview _
func UpdateDoneReview(channelID string, username, args string, force bool) string {
	parameter := utility.GetArgsParameter(args)

	sequences := strings.Split(parameter, " ")
	success := repository.UpdateToDoneReview(sequences, channelID, fmt.Sprintf("<@%s>", username), force)

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), GetReviewQueue(channelID))
	}

	return utility.InvalidSequece()
}

// UpdateReadyQA _
func UpdateReadyQA(channelID string, args string) string {
	parameter := utility.GetArgsParameter(args)

	sequences := strings.Split(parameter, " ")
	success := repository.UpdateToReadyQA(sequences, channelID)

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), GetReviewQueue(channelID))
	}

	return utility.InvalidSequece()
}

// UpdateDoneQA _
func UpdateDoneQA(channelID string, args string) string {
	parameter := utility.GetArgsParameter(args)

	sequences := strings.Split(parameter, " ")
	success := repository.UpdateToDoneQA(sequences, channelID)

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), GetQAQueue(channelID))
	}

	return utility.InvalidSequece()
}

func AddUserReview(channelID string, args string) string {
	parameter := utility.GetArgsParameter(args)
	split := strings.Split(parameter, "][")

	sequence, err := strconv.Atoi(split[0])

	if len(split) < 2 || err != nil {
		return utility.InvalidParameter()
	}

	reviews := repository.GetAllNeedReview(channelID)

	for i, review := range reviews {
		if i+1 == sequence {
			repository.UpdateReview(review.ID, review.Title, review.URL, fmt.Sprintf("%s %s", review.Users, split[1]))
			return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), GetReviewQueue(channelID))
		}
	}

	return utility.InvalidSequece()
}
