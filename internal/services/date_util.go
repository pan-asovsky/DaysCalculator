package services

import (
	"fmt"
	"log"
	"math"
	conv "strconv"
	"time"
)

const startDay = 1
const startMonth = 1
const hour = 0
const min = 0
const sec = 0
const nsec = 0

func GetDateFromString(date string) (time.Time, error) {

	receivedYear, err := conv.Atoi(date)
	if err != nil {
		log.Println("Error converting string to int:", err)
		return time.Time{}, err
	}

	receivedDate := time.Date(receivedYear, startMonth, startDay, hour, min, sec, nsec, time.UTC)
	return receivedDate, nil
}

func GetDateDifference(receivedDate time.Time) string {

	now := time.Now()
	currentDay := now.Day()
	currentMonth := now.Month()
	currentYear := now.Year()

	currentDate := time.Date(currentYear, currentMonth, currentDay, hour, min, sec, nsec, time.UTC)

	difference := currentDate.Sub(receivedDate)
	diffDays := difference.Hours() / 24

	if diffDays > 0 {
		return fmt.Sprintf("Days gone: %d", int(diffDays))
	} else {
		diffDays := int(math.Abs(diffDays))
		return fmt.Sprintf("Days left: %d", diffDays)
	}
}
