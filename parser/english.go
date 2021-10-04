package parser

// Parses datefmt time to Indonesian language date
func NewEnglish() Parser {
	return &English{}
}

type English struct {
	baseParser
}

func (e English) DayLong() string {
	return e.t.Weekday().String()
}

func (e English) DayShort() string {
	return e.DayLong()[0:3]
}

func (e English) MonthLong() string {
	return e.t.Month().String()
}

// Short Month. Like `Jan` or `Dec`
func (e English) MonthShort() string {
	return e.MonthLong()[0:3]
}
