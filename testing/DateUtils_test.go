package testing

import (
	"testing"
	"time"

	"github.com/sabidos/utils"
)

func TestGetEndOfCurrentWeekAtEndOfWeek(t *testing.T) {

	utils.TimeNow = func() time.Time {
		t, _ := time.Parse("2006-01-02T15:04:05.000Z", "2021-01-30T14:43:26.371Z")
		return t
	}

	result := utils.GetEndOfCurrentWeek()

	if result != 1612051199 {
		t.Errorf("TestGetEndOfCurrentWeekAtEndOfWeek() failed, expected %v, got %v", 1612051199, result)
	}

}

func TestGetEndOfCurrentWeekAtBeginningOfWeek(t *testing.T) {

	utils.TimeNow = func() time.Time {
		t, _ := time.Parse("2006-01-02T15:04:05.000Z", "2021-01-31T00:00:00.000Z")
		return t
	}

	result := utils.GetEndOfCurrentWeek()

	if result != 1612655999 {
		t.Errorf("TestGetEndOfCurrentWeekAtBeginningOfWeek() failed, expected %v, got %v", 1612655999, result)
	}

}

func TestGetEndOfCurrentWeekAtMiddleOfWeek(t *testing.T) {

	utils.TimeNow = func() time.Time {
		t, _ := time.Parse("2006-01-02T15:04:05.000Z", "2021-02-03T17:10:00.000Z")
		return t
	}

	result := utils.GetEndOfCurrentWeek()

	if result != 1612655999 {
		t.Errorf("TestGetEndOfCurrentWeekAtBeginningOfWeek() failed, expected %v, got %v", 1612655999, result)
	}

}
