package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambotslack/app/utility"
	"github.com/bot/myteambotslack/app/utility/repository"
)

// AntrianReview _
func AntrianReview(channelID string) []string {
	var antrianReviews []string
	reviews := repository.GetAllNeedReview(channelID)

	if len(reviews) == 0 {
		return append([]string{"Gak ada antrian review nih 👍🏻"})
	}

	antrianReviews = repository.GenerateContentReview(reviews)
	antrianReviews = append([]string{"Ini antrian review tim kamu:\n"}, antrianReviews...)

	return antrianReviews
}

// AntrianQA _
func AntrianQA(channelID string) []string {
	var antrianQA []string
	reviews := repository.GetAllNeedQA(channelID)

	if len(reviews) == 0 {
		return append([]string{"Gak ada antrian QA nih 👍🏻"})
	}

	antrianQA = repository.GenerateContentReview(reviews)
	antrianQA = append([]string{"Ini antrian QA tim kamu:\n"}, antrianQA...)

	return antrianQA
}

// TitipReview _
func TitipReview(channelID string, args string) string {
	if !utility.IsValidParameter(args) {
		return utility.InvalidParameter()
	}

	parameter := utility.GetArgsParameter(args)
	split := strings.Split(parameter, "][")
	title := ""
	url := ""
	users := ""

	for i, s := range split {
		// If Title
		if i == 0 {
			title = s
		}
		// If URL
		if i == 1 {
			url = s
			if strings.HasPrefix(url, "<") {
				url = url[1:]
			}

			if strings.HasSuffix(url, ">") {
				url = url[:len(url)-1]
			}
		}
		// If Users
		if i == 2 {
			users = s
			if strings.HasPrefix(users, ">") {
				users = users[1:]
			}

			if strings.HasSuffix(users, "<") {
				users = url[:len(url)-1]
			}
		}
	}

	if len(split) < 3 {
		return utility.InvalidParameter()
	}

	repository.InsertReview(title, url, users, channelID)

	return utility.SuccessInsertData()
}

// SudahDireview _
func SudahDireview(channelID string, username, args string, force bool) string {
	if !utility.IsValidParameter(args) {
		return utility.InvalidParameter()
	}

	parameter := utility.GetArgsParameter(args)

	sequences := strings.Split(parameter, " ")
	success := repository.UpdateToDoneReview(sequences, channelID, fmt.Sprintf("<@%s>", username), force)

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), AntrianReview(channelID))
	}

	return utility.InvalidSequece()
}

// HapusReview _
func HapusReview(channelID string, args string) string {
	if !utility.IsValidParameter(args) {
		return utility.InvalidParameter()
	}

	parameter := utility.GetArgsParameter(args)

	sequences := strings.Split(parameter, " ")
	success := repository.DeleteReview(sequences, channelID)

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), AntrianReview(channelID))
	}

	return utility.InvalidSequece()
}

// SiapQA _
func SiapQA(channelID string, args string) string {
	if !utility.IsValidParameter(args) {
		return utility.InvalidParameter()
	}

	parameter := utility.GetArgsParameter(args)

	sequences := strings.Split(parameter, " ")
	success := repository.UpdateToReadyQA(sequences, channelID)

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), AntrianReview(channelID))
	}

	return utility.InvalidSequece()
}

// SudahDites _
func SudahDites(channelID string, args string) string {
	if !utility.IsValidParameter(args) {
		return utility.InvalidParameter()
	}

	parameter := utility.GetArgsParameter(args)

	sequences := strings.Split(parameter, " ")
	success := repository.UpdateToDoneQA(sequences, channelID)

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), AntrianQA(channelID))
	}

	return utility.InvalidSequece()
}

// TambahUserReview _
func TambahUserReview(channelID string, args string) string {
	if !utility.IsValidParameter(args) {
		return utility.InvalidParameter()
	}

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
			return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), AntrianReview(channelID))
		}
	}

	return utility.InvalidSequece()
}
