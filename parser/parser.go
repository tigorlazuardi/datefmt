package parser

import (
	"time"
)

// Parser handles the parsing from time.Time to various language formats.
type Parser interface {
	// ParseTime parses the time.
	ParseTime(time.Time)
	// Digit day of month
	Day() uint8
	// Digit day of month with leading 0 if there's only one digit.
	PaddedDay() string
	// Digit month of a year
	Month() uint8
	// Digit month of a year with leading 0 if there's only one digit.
	PaddedMonth() string
	// 4 digit year
	Year() uint16
	// Last two digit of a Year. e.g. year of 1990 will return 90
	YearShort() uint8
	// Hour of the day
	Hour() uint8
	// Hour of the day with leading 0
	PaddedHour() string
	// Minute of an Hour
	Minute() uint8
	// Minute of an Hour with leading 0
	PaddedMinute() string
	// Second of a minute
	Second() uint8
	// Second from unix
	SecondUnix() int64
	// Second of a minute with leading 0
	PaddedSecond() string
	// Long Day. Like `Monday` or `Wednesday`
	DayLong() string
	// Short day. Like `Mon` or `Wed`
	DayShort() string
	// Full Month. Like `January` or `December`
	MonthLong() string
	// Short Month. Like `Jan` or `Dec`
	MonthShort() string
}

// Shared struct between parsers
type baseParser struct {
	t            time.Time
	year         uint16
	yearShort    uint8
	day          uint8
	paddedDay    string
	month        uint8
	paddedMonth  string
	hour         uint8
	paddedHour   string
	minute       uint8
	paddedMinute string
	second       uint8
	secondUnix   int64
	paddedSecond string
}

func newBaseParser(t time.Time) baseParser {
	timeStr := t.String()
	return baseParser{
		t:            t,
		day:          uint8(t.Day()),
		paddedDay:    timeStr[8:10],
		hour:         uint8(t.Hour()),
		paddedHour:   timeStr[11:13],
		minute:       uint8(t.Minute()),
		paddedMinute: timeStr[14:16],
		second:       uint8(t.Second()),
		paddedSecond: timeStr[17:19],
		secondUnix:   t.Unix(),
		month:        uint8(t.Month()),
		paddedMonth:  timeStr[5:7],
		year:         uint16(t.Year()),
		yearShort:    uint8(t.Year() % 100),
	}
}

func (b *baseParser) ParseTime(t time.Time) {
	*b = newBaseParser(t)
}

func (b baseParser) Day() uint8 {
	return b.day
}

func (b baseParser) PaddedDay() string {
	return b.paddedDay
}

func (b baseParser) Month() uint8 {
	return b.month
}

func (b baseParser) PaddedMonth() string {
	return b.paddedMonth
}

func (b baseParser) Year() uint16 {
	return b.year
}

func (b baseParser) YearShort() uint8 {
	return b.yearShort
}

func (b baseParser) Hour() uint8 {
	return b.hour
}

func (b baseParser) PaddedHour() string {
	return b.paddedHour
}

func (b baseParser) Minute() uint8 {
	return b.minute
}

func (b baseParser) PaddedMinute() string {
	return b.paddedMinute
}

func (b baseParser) Second() uint8 {
	return b.second
}

func (b baseParser) PaddedSecond() string {
	return b.paddedSecond
}

func (b baseParser) SecondUnix() int64 {
	return b.secondUnix
}
