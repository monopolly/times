package times

import (
	"fmt"
	"strings"
	"time"

	"github.com/monopolly/now"
)

var formats = []string{
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RubyDate,
	"20060102",
	"20060201",
	"20060102T030405Z",
	"01.02.2006",
	"02 Jan 06 15:04 -0700",
	"02 Jan 06 15:04 MST",
	"02-Jan-2006 03:04:05 MST",
	"02-Jan-2006",
	"02-nov-2006",
	"02.01.06",
	"02.01.2006",
	"1-2",
	"1.2.2006",
	"1/2/2006 15:4:5",
	"1/2/2006",
	"15",
	"15:4",
	"15:4:5 Jan 2, 2006 MST ",
	"15:4:5 Jan 2, 2006 MST",
	"15:4:5",
	"2006-01-02 03:04:05",
	"2006-01-02 15:04:05.000000000 -0700",
	"2006-01-02 15:04:05.999999999 -0700 MST",
	"2006-01-02 3:4:5",
	"2006-01-02",
	"2006-01-02T03:04:05-07:00",
	"2006-01-02T03:04:05.0Z",
	"2006-01-02T03:04:05Z",
	"2006-01-02T15:04:05.999999999Z07:00",
	"2006-01-02T15:04:05Z07:00",
	"2006-02-01T03:04:05Z",
	"2006-1-2 ",
	"2006-1-2 03:04:05",
	"2006-1-2 15:4 ",
	"2006-1-2 15:4",
	"2006-1-2 15:4:5",
	"2006-1-2",
	"2006.01.02",
	"2006/01/02",
	"3:04PM",
	"before Jan-2006",
	"Jan _2 15:04:05",
	"Jan _2 15:04:05.000",
	"Jan _2 15:04:05.000000",
	"Jan _2 15:04:05.000000000",
	"Mon Jan 02 15:04:05 -0700 2006",
	"Mon Jan _2 15:04:05 2006",
	"Mon Jan _2 15:04:05 MST 2006",
	"Mon, 02 Jan 2006 15:04:05 -0700",
	"Mon, 02 Jan 2006 15:04:05 MST",
	"Monday, 02-Jan-06 15:04:05 MST",
}

func Parse(v string) (t *time.Time) {
	v = strings.TrimSpace(v)
	tm, err := now.Parse(v)
	if err == nil && !tm.IsZero() {
		fmt.Println("")
		return &tm
	}

	for _, x := range formats {
		tm, err := time.Parse(x, v)
		if err != nil {
			continue
		}
		if tm.IsZero() {
			continue
		}
		if tm.Year() == 1 {
			continue
		}

		return &tm
	}

	return nil
}
