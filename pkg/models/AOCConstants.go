package models

// AOC URL constants.
const (
	// AOCTimezone is the timezone for the AOC API.
	AOCTimezone = "America/New_York"
	// BaseURL is the base URL for the AOC API.
	BaseURL = "https://adventofcode.com"
)

var (
	// YearURL is the URL for the year.
	YearURL = BaseURL + "/%d"
	// DayURL is the URL for the day.
	DayURL = YearURL + "/day/%d"
	// inputURL is the URL for the input data.
	InputURL = DayURL + "/input"
	// AnswerURL is the URL post the answer.
	AnswerURL = DayURL + "/answer"
)
