package times

import (
	"fmt"
	"time"
)

func Russian(t time.Time) (a *ru) {
	a = new(ru)
	a.t = t
	return
}
func RussianUnix(t int64) (a *ru) {
	a = new(ru)
	a.t = time.Unix(t, 0)
	return
}

func RussianNow() (a *ru) {
	a = new(ru)
	a.t = time.Now()
	return
}

type ru struct {
	t time.Time
}

var russianmonth = []string{
	"",
	"января",
	"февраля",
	"марта",
	"апреля",
	"мая",
	"июня",
	"июля",
	"августа",
	"сентября",
	"октября",
	"ноября",
	"декабря",
}

//1 октября 2019, 14:42
func (a *ru) Full() string {
	m := russianmonth[a.t.Month()]
	return fmt.Sprintf("%d %s %d, %s", a.t.Day(), m, a.t.Year(), a.t.Format("15:04"))
}

//1 октября 2019
func (a *ru) Date() string {
	m := russianmonth[a.t.Month()]
	return fmt.Sprintf("%d %s %d", a.t.Day(), m, a.t.Year())
}

//1 октября 2019
func (a *ru) DateOnly() string {
	m := russianmonth[a.t.Month()]
	return fmt.Sprintf("%d %s", a.t.Day(), m)
}
