package repository

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bot/myteambotslack/app/models"
	"github.com/bot/myteambotslack/app/utility"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// GetAllNeedReview _
func GetAllNeedReview(channelID string) []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_reviewed = ? AND channel_id = ?", false, channelID), qm.OrderBy("created_at")).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// GetAllNeedQA _
func GetAllNeedQA(channelID string) []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_reviewed = ? AND is_tested = ? AND channel_id = ?", true, false, channelID), qm.OrderBy("created_at")).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// InsertReview _
func InsertReview(title, url, users string, channelID string) {
	var review models.Review

	review.ChannelID = channelID
	review.URL = url
	review.Title = title
	review.IsReviewed = false
	review.IsTested = false
	review.Users = users

	err := review.InsertG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

func UpdateReview(ID uint, title, url, users string) {
	review, err := models.Reviews(qm.Where("id = ?", ID)).OneG()
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	review.URL = url
	review.IsReviewed = false
	review.IsTested = false
	review.Title = title
	review.Users = users

	err = review.UpdateG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

// UpdateToDoneReview _
func UpdateToDoneReview(sequences []string, channelID string, user string, force bool) bool {
	successToUpdate := false
	reviews := GetAllNeedReview(channelID)

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		for i, review := range reviews {
			if i == sequence-1 {
				if force {
					review.Users = ""
				} else {
					review.Users = removeAvailableUsers(review.Users, user)
				}

				err := review.UpdateG(boil.Infer())
				if err != nil {
					panic(err)
				}

				successToUpdate = true
				break
			}
		}
	}

	return successToUpdate
}

func DeleteReview(sequences []string, channelID string) bool {
	successToDelete := false
	reviews := GetAllNeedReview(channelID)

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		for i, review := range reviews {
			if i+1 == sequence {
				err := review.DeleteG()
				if err != nil {
					panic(err)
				}

				successToDelete = true
				break
			}
		}
	}

	return successToDelete
}

func UpdateToReadyQA(sequences []string, channelID string) bool {
	successToUpdate := false
	reviews := GetAllNeedReview(channelID)

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		for i, review := range reviews {
			if i+1 == sequence {
				review.IsReviewed = true

				err := review.UpdateG(boil.Infer())
				if err != nil {
					panic(err)
				}

				successToUpdate = true
				break
			}
		}
	}

	return successToUpdate
}

func UpdateToDoneQA(sequences []string, channelID string) bool {
	successToUpdate := false
	reviews := GetAllNeedQA(channelID)

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		for i, review := range reviews {
			if i == sequence-1 {
				review.IsTested = true

				err := review.UpdateG(boil.Infer())
				if err != nil {
					panic(err)
				}

				successToUpdate = true
				break
			}
		}
	}

	return successToUpdate
}

// GenerateContentReview _
func GenerateContentReview(reviews []*models.Review) []string {
	var buffer bytes.Buffer
	var allReviews []string

	for i, review := range reviews {
		if review.Title == "" {
			buffer.WriteString(fmt.Sprintf("%d. <%s|Belum ada title> %s\n", i+1, review.URL, review.Users))
		} else {
			buffer.WriteString(fmt.Sprintf("%d. <%s|%s> %s\n", i+1, review.URL, review.Title, review.Users))
		}

		if (i > 0 && i%10 == 0) || i == len(reviews)-2 {
			fmt.Println("APPEND")
			allReviews = append(allReviews, buffer.String())
			buffer.Reset()
		}
	}

	return allReviews
}

// Private functions
func removeAvailableUsers(users, deleteUser string) string {
	splitUsers := utility.GetUsersFromArgs(users)
	var newUsers []string

	for _, user := range splitUsers {
		if user == deleteUser {
			continue
		} else {
			newUsers = append(newUsers, user)
		}
	}

	return strings.Join(newUsers, " ")
}
