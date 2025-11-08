package meetup

import "time"

type WeekSchedule int

const (
	First = WeekSchedule(iota)
	Second
	Third
	Fourth
	Teenth
	Last
)

const timezone = "Europe/Athens"

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	loc, _ := time.LoadLocation(timezone)
	date := time.Date(year, month, 1, 0, 0, 0, 0, loc) // Start from 1st of month
	wdayMatches := 0                                   // Counter
	var day int
	for {
		if date.Weekday() == wDay {
			wdayMatches++
			day = date.Day() // Remember day of last weekday match
			if (wSched == First && wdayMatches == 1) ||
				(wSched == Second && wdayMatches == 2) ||
				(wSched == Third && wdayMatches == 3) ||
				(wSched == Fourth && wdayMatches == 4) ||
				(wSched == Teenth && day >= 13 && day <= 19) {
				return day
			}
		}
		nextDate := date.Add(time.Hour * 24)  // Add one day
		if nextDate.Month() != date.Month() { // Month changed
			if wSched == Last {
				return day
			}
			panic("should never happen")
		}
		date = nextDate
	}
}
