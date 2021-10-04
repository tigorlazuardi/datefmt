package parser

// Parses datefmt time to Indonesian language date
func NewIndonesian() Parser {
	return &Indonesian{}
}

type Indonesian struct {
	baseParser
}

func (i Indonesian) DayLong() string {
	switch int(i.t.Weekday()) {
	case 0:
		return "Minggu"
	case 1:
		return "Senin"
	case 2:
		return "Selasa"
	case 3:
		return "Rabu"
	case 4:
		return "Kamis"
	case 5:
		return "Jumat"
	case 6:
		return "Sabtu"
	default:
		return "Minggu"
	}
}

func (i Indonesian) DayShort() string {
	return i.DayLong()[0:3]
}

func (i Indonesian) MonthLong() string {
	switch i.month {
	case 1:
		return "Januari"
	case 2:
		return "Februari"
	case 3:
		return "Maret"
	case 4:
		return "April"
	case 5:
		return "Mei"
	case 6:
		return "Juni"
	case 7:
		return "Juli"
	case 8:
		return "Agustus"
	case 9:
		return "September"
	case 10:
		return "Oktober"
	case 11:
		return "November"
	case 12:
		return "Desember"
	default:
		return "Januari"
	}
}

// Short Month. Like `Jan` or `Dec`
func (i Indonesian) MonthShort() string {
	return i.MonthLong()[0:3]
}
