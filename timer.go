package gotime

import (
	"time"
)

func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	nt := StartOfMonth(t).AddDate(0, 1, 0)
	return time.Unix(nt.Unix()-1, 0)
}

func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

func EndOfYear(t time.Time) time.Time {
	nt := StartOfYear(t).AddDate(1, 0, 0)
	return time.Unix(nt.Unix()-1, 0)
}

func AddDays(t time.Time, d int) time.Time {
	return t.AddDate(0, 0, d)
}

func AddMonths(t time.Time, m int) time.Time {
	return t.AddDate(0, m, 0)
}

func AddYears(t time.Time, y int) time.Time {
	return t.AddDate(y, 0, 0)
}
