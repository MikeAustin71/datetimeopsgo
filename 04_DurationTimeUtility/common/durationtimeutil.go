package common

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	// MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond
	MicroSecondNanoseconds = int64(1000)
	// MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
	MilliSecondNanoseconds = int64(1000 * 1000)
	// SecondNanoseconds - Number of Nanoseconds in a Second
	SecondNanoseconds = int64(1000 * 1000 * 1000)
	// MinuteNanoSeconds - Number of Nanoseconds in a minute
	MinuteNanoSeconds = int64(1000 * 1000 * 1000 * 60)
	// HourNanoSeconds - Number of Nanoseconds in an hour
	HourNanoSeconds = int64(1000 * 1000 * 1000 * 60 * 60)
	// DayNanoSeconds - Number of Nanoseconds in a 24-hour day
	DayNanoSeconds = int64(1000 * 1000 * 1000 * 60 * 60 * 24)
	// YearNanoSeconds - Number of Nanoseconds in a 365-day year
	YearNanoSeconds = int64(1000 * 1000 * 1000 * 60 * 60 * 24 * 365)
)

// ElapsedDuration - holds elements of
// time duration
type ElapsedDuration struct {
	TimeDuration time.Duration
	Years        int64
	Days         int64
	Hours        int64
	Minutes      int64
	Seconds      int64
	MilliSeconds int64
	MicroSeconds int64
	NanoSeconds  int64
	// NanosecStr - Example: 2-Days 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds
	NanosecStr string
	// DurationStr - Example: 2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
	DurationStr string
	// DefaultStr - Example: 61h26m46.864197832s - format provided by 'go' library
	DefaultStr string
}

// GetDuration - Returns a time.Duration structure defining the duration between
// input parameters startTime and endTime
func (eld ElapsedDuration) GetDuration(startTime time.Time, endTime time.Time) (time.Duration, error) {

	def := time.Duration(0)

	if startTime.Equal(endTime) {
		return def, nil
	}

	if endTime.Before(startTime) {
		return def, errors.New("DateTimeUtility.GetDuration() Error: endTime less than startTime")
	}

	return endTime.Sub(startTime), nil
}

// GetDurationBreakDown - Receives a Duration type
// and returns a breakdown of duration by years,
// days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.
func (eld ElapsedDuration) GetDurationBreakDown(d time.Duration) ElapsedDuration {
	str := ""
	ed := ElapsedDuration{TimeDuration: d}
	ed.DefaultStr = fmt.Sprintf("%v", d)
	firstEle := false
	rd := int64(d)

	if rd >= YearNanoSeconds {
		ed.Years = rd / YearNanoSeconds
		rd -= YearNanoSeconds * ed.Years
	}

	if ed.Years > 0 {
		str = fmt.Sprintf("%v-Years ", ed.Years)
		firstEle = true
	}

	if rd >= DayNanoSeconds {
		ed.Days = rd / DayNanoSeconds
		rd -= DayNanoSeconds * ed.Days
	}

	if ed.Days > 0 || firstEle {
		str += fmt.Sprintf("%v-Days ", ed.Days)
		firstEle = true
	}

	if rd >= HourNanoSeconds {
		ed.Hours = rd / HourNanoSeconds
		rd -= HourNanoSeconds * ed.Hours
	}

	if ed.Hours > 0 || firstEle {
		str += fmt.Sprintf("%v-Hours ", ed.Hours)
		firstEle = true
	}

	if rd >= MinuteNanoSeconds {
		ed.Minutes = rd / MinuteNanoSeconds
		rd -= MinuteNanoSeconds * ed.Minutes
	}

	if ed.Minutes > 0 || firstEle {
		str += fmt.Sprintf("%v-Minutes ", ed.Minutes)
		firstEle = true
	}

	if rd >= SecondNanoseconds {
		ed.Seconds = rd / SecondNanoseconds
		rd -= SecondNanoseconds * ed.Seconds
	}

	if ed.Seconds > 0 || firstEle {
		str += fmt.Sprintf("%v-Seconds ", ed.Seconds)
		firstEle = true
	}

	ed.NanosecStr = str + fmt.Sprintf("%v-Nanoseconds", rd)

	if rd >= MilliSecondNanoseconds {
		ed.MilliSeconds = rd / MilliSecondNanoseconds
		rd -= MilliSecondNanoseconds * ed.MilliSeconds
	}

	if ed.MilliSeconds > 0 || firstEle {
		str += fmt.Sprintf("%v-Milliseconds ", ed.MilliSeconds)
		firstEle = true
	}

	if rd >= MicroSecondNanoseconds {
		ed.MicroSeconds = rd / MicroSecondNanoseconds
		rd -= MicroSecondNanoseconds * ed.MicroSeconds
	}

	if ed.MicroSeconds > 0 || firstEle {
		str += fmt.Sprintf("%v-Microseconds ", ed.MicroSeconds)
		firstEle = true
	}

	ed.NanoSeconds = rd

	if ed.NanoSeconds > 0 || firstEle {
		str += fmt.Sprintf("%v-Nanoseconds", ed.NanoSeconds)
		firstEle = true
	}

	ed.DurationStr = strings.TrimRight(str, " ")

	return ed

}

// GetElapsedTime - calculates the elapsed time
// between input parameters startTime and endTime.
// The result is returned in an ElapsedDuration
// structure.
func (eld ElapsedDuration) GetElapsedTime(startTime time.Time, endTime time.Time) (ElapsedDuration, error) {

	dur, err := eld.GetDuration(startTime, endTime)

	if err != nil {
		s := "DateTimeUtility-GetElapsedTime Error: " + err.Error()

		return ElapsedDuration{}, errors.New(s)
	}

	return eld.GetDurationBreakDown(dur), nil

}
