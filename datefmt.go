package datefmt

import (
	"bytes"
	"io"
	"text/template"
	"time"

	"github.com/tigorlazuardi/datefmt/parser"
)

type Formatter struct {
	parser.Parser
	Day          uint8
	DayLong      string
	DayShort     string
	Hour         uint8
	Minute       uint8
	Month        uint8
	MonthLong    string
	MonthShort   string
	PaddedDay    string
	PaddedHour   string
	PaddedMinute string
	PaddedMonth  string
	PaddedSecond string
	Second       uint8
	SecondUnix   int64
	Year         uint16
	YearShort    uint8
	template     *template.Template
}

/*
Turns the given string to replaced format.
Use `MustRender` if the given string format is hard coded so you don't have to handle errors.

Example:
	str, err := formatter.Render("Jakarta, {{.Day}} {{.MonthLong}} {{.Year}} Pukul {{.PaddedHour}}:{{.PaddedMinute}} WIB")
	// str will have output
	// Output: Jakarta, 1 Oktober 2021 Pukul 13:05 WIB

Available APIs:
	{{.Day}}          // day of month
	{{.DayLong}}      // full day name
	{{.DayShort}}     // three letter day name
	{{.Hour}}         // hour of day
	{{.Minute}}       // minute of hour
	{{.Month}}        // month of year
	{{.MonthLong}}    // full month name
	{{.MonthShort}}   // three letter month name
	{{.PaddedDay}}    // day of month with leading 0
	{{.PaddedHour}}   // hour with leading 0
	{{.PaddedMinute}} // minute with leading 0
	{{.PaddedMonth}}  // month with leading 0
	{{.PaddedSecond}} // second with leading 0
	{{.Second}}       // second of minute
	{{.SecondUnix}}   // second from epoch
	{{.Year}}         // 4 digit year
	{{.YearShort}}    // last 2 digit of year
*/
func (f Formatter) Render(s string) (string, error) {
	t, err := f.template.Parse(s)
	if err != nil {
		return "", err
	}
	b := &bytes.Buffer{}
	_ = t.Execute(b, f)
	return b.String(), err
}

/*
Turns the given string to replaced format and write the result to the given writer

Example:
	buf := &bytes.Buffer{}
	str, err := formatter.RenderWriter(buf, "Jakarta, {{.Day}} {{.MonthLong}} {{.Year}} Pukul {{.PaddedHour}}:{{.PaddedMinute}} WIB")
	// buf.String() will have output
	// Output: Jakarta, 1 Oktober 2021 Pukul 13:05 WIB

Available APIs:
	{{.Day}}          // day of month
	{{.DayLong}}      // full day name
	{{.DayShort}}     // three letter day name
	{{.Hour}}         // hour of day
	{{.Minute}}       // minute of hour
	{{.Month}}        // month of year
	{{.MonthLong}}    // full month name
	{{.MonthShort}}   // three letter month name
	{{.PaddedDay}}    // day of month with leading 0
	{{.PaddedHour}}   // hour with leading 0
	{{.PaddedMinute}} // minute with leading 0
	{{.PaddedMonth}}  // month with leading 0
	{{.PaddedSecond}} // second with leading 0
	{{.Second}}       // second of minute
	{{.SecondUnix}}   // second from epoch
	{{.Year}}         // 4 digit year
	{{.YearShort}}    // last 2 digit of year
*/
func (f Formatter) RenderWriter(writer io.Writer, s string) error {
	t, err := f.template.Parse(s)
	if err != nil {
		return err
	}
	err = t.Execute(writer, f)
	return err
}

/*
Turns the given string to replaced format, function panics when failed to parse.

Example:
	str := formatter.MustRender("Jakarta, {{.Day}} {{.MonthLong}} {{.Year}} Pukul {{.PaddedHour}}:{{.PaddedMinute}} WIB")
	// str will have output
	// Output: Jakarta, 1 Oktober 2021 Pukul 13:05 WIB

Available APIs:
	{{.Day}}          // day of month
	{{.DayLong}}      // full day name
	{{.DayShort}}     // three letter day name
	{{.Hour}}         // hour of day
	{{.Minute}}       // minute of hour
	{{.Month}}        // month of year
	{{.MonthLong}}    // full month name
	{{.MonthShort}}   // three letter month name
	{{.PaddedDay}}    // day of month with leading 0
	{{.PaddedHour}}   // hour with leading 0
	{{.PaddedMinute}} // minute with leading 0
	{{.PaddedMonth}}  // month with leading 0
	{{.PaddedSecond}} // second with leading 0
	{{.Second}}       // second of minute
	{{.SecondUnix}}   // second from epoch
	{{.Year}}         // 4 digit year
	{{.YearShort}}    // last 2 digit of year
*/
func (f Formatter) MustRender(s string) string {
	t := template.Must(f.template.Parse(s))
	b := &bytes.Buffer{}
	_ = t.Execute(b, f)
	return b.String()
}

// Updates current formatter time with new time
func (f *Formatter) ParseTime(t time.Time) {
	f.Parser.ParseTime(t)
	f.Day = f.Parser.Day()
	f.DayLong = f.Parser.DayLong()
	f.DayShort = f.Parser.DayShort()
	f.Hour = f.Parser.Hour()
	f.Minute = f.Parser.Minute()
	f.Month = f.Parser.Month()
	f.MonthLong = f.Parser.MonthLong()
	f.MonthShort = f.Parser.MonthShort()
	f.PaddedDay = f.Parser.PaddedDay()
	f.PaddedHour = f.Parser.PaddedHour()
	f.PaddedMinute = f.Parser.PaddedMinute()
	f.PaddedSecond = f.Parser.PaddedSecond()
	f.Second = f.Parser.Second()
	f.SecondUnix = f.Parser.SecondUnix()
	f.Year = f.Parser.Year()
	f.YearShort = f.Parser.YearShort()
}

/*
Creates new formatter for given time with specific parser.

Example:

	formatter := datefmt.New(time.Now(), parser.NewIndonesian())
*/
func New(t time.Time, parser parser.Parser) Formatter {
	f := Formatter{Parser: parser}
	f.ParseTime(t)
	f.template = template.New("datefmt/formatter")
	return f
}
