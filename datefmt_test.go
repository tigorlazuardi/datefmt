package datefmt_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tigorlazuardi/datefmt"
	"github.com/tigorlazuardi/datefmt/parser"
)

func ExampleFormatter_Render() {
	t, _ := time.Parse(time.RFC3339, "2021-10-01T13:05:31Z")
	formatter := datefmt.New(t, parser.NewEnglish())

	str, err := formatter.Render("New York, {{.DayLong}} {{.Day}} {{.MonthLong}} {{.Year}} {{.PaddedHour}}:{{.PaddedMinute}}")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)

	// Output: New York, Friday 1 October 2021 13:05
}

func ExampleFormatter_RenderWriter() {
	t, _ := time.Parse(time.RFC3339, "2021-10-01T13:05:31Z")
	formatter := datefmt.New(t, parser.NewIndonesian())

	writer := &bytes.Buffer{}
	err := formatter.RenderWriter(writer, "New York, {{.DayLong}} {{.Day}} {{.MonthLong}} {{.Year}} {{.PaddedHour}}:{{.PaddedMinute}}")
	if err != nil {
		panic(err)
	}
	fmt.Println(writer.String())

	// Output: New York, Friday 1 October 2021 13:05
}

func ExampleFormatter_MustRender() {
	t, _ := time.Parse(time.RFC3339, "2021-10-01T13:05:31Z")
	formatter := datefmt.New(t, parser.NewIndonesian())

	// Panics on failure to render (bad format string)
	str := formatter.MustRender("New York, {{.DayLong}} {{.Day}} {{.MonthLong}} {{.Year}} {{.PaddedHour}}:{{.PaddedMinute}}")
	fmt.Println(str)

	// Output: New York, Friday 1 October 2021 13:05
}

func Test_RenderFail(test *testing.T) {
	t, _ := time.Parse(time.RFC3339, "2021-10-01T13:05:31Z")
	formatter := datefmt.New(t, parser.NewEnglish())

	_, err := formatter.Render("Jakarta, {{if .DayLong}}")
	require.NotNil(test, err)
}

func Test_RenderWriter_Fail_File_Closed(test *testing.T) {
	t, _ := time.Parse(time.RFC3339, "2021-10-01T13:05:31Z")
	formatter := datefmt.New(t, parser.NewEnglish())

	file, err := os.CreateTemp(os.TempDir(), "Test_RenderWriter_EOF")
	require.Nil(test, err)
	defer os.Remove(file.Name())
	file.Close()

	err = formatter.RenderWriter(file, "Jakarta, {{.DayLong}} {{.Day}} {{.MonthLong}} {{.Year}} Pukul {{.PaddedHour}}:{{.PaddedMinute}} WIB")
	require.NotNil(test, err)
}
