package tests

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	in "github.com/pan-asovsky/DaysCalculator/internal/services"
	"math"
	"testing"
	"time"
)

func TestGetDateFromString(t *testing.T) {

	t.Run("FirstCorrectDate", func(t *testing.T) {
		firstCorrectDateParam := "2012"
		firstCorrectDate := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)

		firstDate, err := in.GetDateFromString(firstCorrectDateParam)
		assert.Equal(t, err, nil)
		assert.Equal(t, firstDate, firstCorrectDate)
	})

	t.Run("SecondCorrectDate", func(t *testing.T) {
		secondCorrectDateParam := "2034"
		secondCorrectDate := time.Date(2034, 1, 1, 0, 0, 0, 0, time.UTC)

		firstDate, err := in.GetDateFromString(secondCorrectDateParam)
		assert.Equal(t, err, nil)
		assert.Equal(t, firstDate, secondCorrectDate)
	})

	t.Run("IncorrectDate", func(t *testing.T) {
		incorrectDateParam := "abc"
		_, err := in.GetDateFromString(incorrectDateParam)
		assert.NotEqual(t, err, nil)
	})

}

func TestGetDateDifference(t *testing.T) {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	currentDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	t.Run("PastDate", func(t *testing.T) {
		pastDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		diff := int(currentDate.Sub(pastDate).Hours() / 24)

		difference := in.GetDateDifference(pastDate)
		assert.Equal(t, difference, fmt.Sprintf("Days gone: %d", diff))
	})

	t.Run("FutureDate", func(t *testing.T) {
		pastDate := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
		diff := int(math.Abs(currentDate.Sub(pastDate).Hours() / 24))

		difference := in.GetDateDifference(pastDate)
		assert.Equal(t, difference, fmt.Sprintf("Days left: %d", diff))
	})

}
