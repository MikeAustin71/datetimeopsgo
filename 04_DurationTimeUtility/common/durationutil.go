package common

import (
	"errors"
	"fmt"
	"math"
	"time"
)

const (
	// MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond
	MicroSecondNanoseconds = int64(time.Microsecond)
	// MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
	MilliSecondNanoseconds = int64(time.Millisecond)
	// SecondNanoseconds - Number of Nanoseconds in a Second
	SecondNanoseconds = int64(time.Second)
	// MinuteNanoSeconds - Number of Nanoseconds in a minute
	MinuteNanoSeconds = int64(time.Minute)
	// HourNanoSeconds - Number of Nanoseconds in an hour
	HourNanoSeconds = int64(time.Hour)
	// DayNanoSeconds - Number of Nanoseconds in a 24-hour day
	DayNanoSeconds = int64(24) * HourNanoSeconds
	// StdYearNanoSeconds - Number of Nanoseconds in a 365-day year

	/*
		For the Gregorian calendar the average length of the calendar year
		(the mean year) across the complete leap cycle of 400 years is 365.2425 days.
		The Gregorian Average Year is therefore equivalent to 365 days, 5 hours,
		49 minutes and 12 seconds.
		Sources:
		https://en.wikipedia.org/wiki/Year
		Source: https://en.wikipedia.org/wiki/Gregorian_calendar
	*/

	GregorianYearNanoSeconds = int64(31556952000000000)

	// Year durations are calculated using a standard 365-day year consisting
	// of 24-hour days.
	StdYearNanoSeconds = DayNanoSeconds * 365
)

// DurationUtility - holds elements of
// time duration
type DurationUtility struct {
	StartDateTime time.Time
	EndDateTime   time.Time
	TimeDuration  time.Duration
	Years         int64
	YearsNanosecs int64
	Months        int64
	MonthNanosecs int64
	Days          int64
	Hours         int64
	Minutes       int64
	Seconds       int64
	MilliSeconds  int64
	MicroSeconds  int64
	NanoSeconds   int64
	// NanosecStr - Example: 2-Days 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds
	NanosecStr string

	// DurationStr - Example: 2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
	DurationStr string

	// DaysStr = Example 97-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
	DaysStr string

	// HoursStr - Example 152-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
	HoursStr string

	// DefaultStr - Example: 61h26m46.864197832s - format provided by 'go' library
	DefaultStr string
}

// AddDurationToThis - Add Duration to the existing value of this
// DurationUtility data structure
func (du *DurationUtility) AddDurationToThis(duration time.Duration) {

	durPlus := du.TimeDuration + duration

	elapsedDuration := du.GetDurationBreakDown(du.StartDateTime, durPlus)

	du.CopyToThis(elapsedDuration)
}

// AddToThis - Add another DurationUtility data structure
// to the existing DurationUtility data structure
func (du *DurationUtility) AddToThis(duIn DurationUtility) {

	durPlus := du.TimeDuration + duIn.TimeDuration

	elapsedDuration := du.GetDurationBreakDown(du.StartDateTime, durPlus)

	du.CopyToThis(elapsedDuration)

}

// CalculateDurationFromTimes - Calculate a duration from
// 'startTime' and 'endTime' values passed to this method.
// Results are stored int he current DurationUtility data
// structure.
func (du *DurationUtility) CalculateDurationFromTimes(startTime time.Time, endTime time.Time) error {

	if endTime.Before(startTime) {
		return errors.New("DurationUtility.CalculateDurationFromTimes() Error: endTime less than startTime")
	}

	if startTime.IsZero() && endTime.IsZero() {
		return errors.New("DurationUtility.CalculateDurationFromTimes() Error: Both startTime and endTime are zero time values!")
	}

	du.StartDateTime = startTime

	du.EndDateTime = endTime

	du.TimeDuration = endTime.Sub(startTime)

	du.CalculateDurationElements()

	du.CalculateDurationStrings()

	return nil
}

// CalculateDurationFromElements -This method is used to calculate durations based on the
// existing time element values of the DurationUtility data structure (i.e. Years, Months,
// Days, Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds) To use this
// method properly first set the data structure's StartDateTime and then fill in the time
// element values as needed (Years, Months, Days, Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds). !!!! StartDateTime !!! is required!!!! If you fail to
// to populate DurationUtility.StartDateTime before calling this method, the starting
// date time will default to time.Now().UTC()!!!
//
// Note: You can set time elements to negative values. In this case, StartDateTime will be
// converted to EndDateTime and StartDateTime will be re-calculated accordingly.
//
// Example call: du := DurationUtility{StartDateTime: t1, Years: 2, Months: 3}
//               du.CalculateDurationFromElements()
//////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (du *DurationUtility) CalculateDurationFromElements() {

	if du.StartDateTime.IsZero() {
		du.StartDateTime = time.Now().UTC()
	}

	startDate := du.StartDateTime
	addTime :=  (du.Hours * HourNanoSeconds) + (du.Minutes * MinuteNanoSeconds) +
		(du.Seconds * SecondNanoseconds) + (du.MilliSeconds * MilliSecondNanoseconds) +
		(du.MicroSeconds * MicroSecondNanoseconds) + du.NanoSeconds

	endDate := startDate.AddDate(int(du.Years), int(du.Months), int(du.Days)).Add(time.Duration(addTime))

	du.Empty()

	if endDate.Before(startDate) {
		du.StartDateTime = endDate
		du.EndDateTime = startDate
	} else {
		du.StartDateTime = startDate
		du.EndDateTime = endDate
	}

	du.TimeDuration = du.EndDateTime.Sub(du.StartDateTime)

	du.CalculateDurationElements()
	du.CalculateDurationStrings()

	return
}

// CalculateDurationStrings - Calculates Duration breakdowns
// in a variety of display formats based on the Time Duration
// value stored in the existing DurationUtility.
func (du *DurationUtility) CalculateDurationStrings() {

	yearsElement := ""
	monthsElement := ""
	daysElement := ""
	hoursElement := "0-Hours "
	minutesElement := "0-Minutes "
	secondsElement := "0-Seconds "
	millisecondsElement := "0-Milliseconds "
	microsecondsElement := "0-Microseconds "
	nanosecondsElement := "0-Nanoseconds"

	rd := int64(du.TimeDuration)
	du.DefaultStr = fmt.Sprintf("%v", du.TimeDuration)
	if rd == 0 {
		du.DurationStr = hoursElement + minutesElement +
			secondsElement + millisecondsElement +
			microsecondsElement + nanosecondsElement
		du.NanosecStr = hoursElement + minutesElement +

			secondsElement + nanosecondsElement

		du.DaysStr = "0-Days " +
			hoursElement + minutesElement +
			secondsElement + millisecondsElement +
			microsecondsElement + nanosecondsElement

		du.HoursStr = hoursElement + minutesElement +
			secondsElement + millisecondsElement +
			microsecondsElement + nanosecondsElement

		return
	}

	if du.Years > 0 {
		yearsElement = fmt.Sprintf("%v-Years ", du.Years)
		monthsElement = "0-Months "
		daysElement = "0-Days "
	}

	if du.Months > 0 {
		monthsElement = fmt.Sprintf("%v-Months ", du.Months)
		daysElement = "0-Days "
	}

	if du.Days > 0 {
		daysElement = fmt.Sprintf("%v-Days ", du.Days)
	}

	hoursElement = fmt.Sprintf("%v-Hours ", du.Hours)

	minutesElement = fmt.Sprintf("%v-Minutes ", du.Minutes)

	secondsElement = fmt.Sprintf("%v-Seconds ", du.Seconds)

	rn := (du.MilliSeconds * MilliSecondNanoseconds) +
		(du.MicroSeconds * MicroSecondNanoseconds) +
		du.NanoSeconds

	totalNanoseconds := fmt.Sprintf("%v-Nanoseconds", rn)
	du.NanosecStr = yearsElement + monthsElement + daysElement +
		hoursElement + minutesElement + secondsElement +
		totalNanoseconds

	millisecondsElement = fmt.Sprintf("%v-Milliseconds ", du.MilliSeconds)

	microsecondsElement = fmt.Sprintf("%v-Microseconds ", du.MicroSeconds)

	nanosecondsElement = fmt.Sprintf("%v-Nanoseconds", du.NanoSeconds)

	du.DurationStr = yearsElement + monthsElement + daysElement +
		hoursElement + minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	consolDays := (du.YearsNanosecs + du.MonthNanosecs +
		(du.Days * DayNanoSeconds)) / DayNanoSeconds

	du.DaysStr = fmt.Sprintf("%v-Days ", consolDays) +
		hoursElement + minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	consolHours := (consolDays * 24) + du.Hours

	du.HoursStr = fmt.Sprintf("%v-Hours ", consolHours) +
		minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	return
}

// CalculateDurationElements - Breaks down the duration
// value into Years, Days, Hours, Minutes, Seconds,
// Milliseconds, Microseconds and Nanoseconds. Calculation
// is based on the 'TimeDuration' value currently stored
// in the DurationUtility data structure.
func (du *DurationUtility) CalculateDurationElements() {

	rd := int64(du.TimeDuration)
	du.DefaultStr = fmt.Sprintf("%v", du.TimeDuration)

	if rd == 0 {
		return
	}

	if du.StartDateTime.IsZero() {
		du.StartDateTime = time.Now().UTC()
	}

	du.EndDateTime = du.StartDateTime.Add(du.TimeDuration)

	yearDateTime := du.StartDateTime
	i := 0
	var prevDateTime time.Time
	for yearDateTime.Before(du.EndDateTime) {
		prevDateTime = yearDateTime
		i++
		yearDateTime = du.StartDateTime.AddDate(i, 0, 0)
	}

	i -= 1

	if i > 0 {
		yearDateTime = prevDateTime
		du.Years = int64(i)
		du.YearsNanosecs = int64(yearDateTime.Sub(du.StartDateTime))
		rd -= du.YearsNanosecs
	}

	i = 0

	mthDateTime := yearDateTime

	for mthDateTime.Before(du.EndDateTime) {
		prevDateTime = mthDateTime
		i++
		mthDateTime = yearDateTime.AddDate(0, i, 0)
	}

	i -= 1

	if i > 0 {
		mthDateTime = prevDateTime
		du.Months = int64(i)
		du.MonthNanosecs = int64(mthDateTime.Sub(yearDateTime))
		rd -= du.MonthNanosecs

	}

	if rd >= DayNanoSeconds {
		du.Days = rd / DayNanoSeconds
		rd -= DayNanoSeconds * du.Days
	}

	if rd >= HourNanoSeconds {
		du.Hours = rd / HourNanoSeconds
		rd -= HourNanoSeconds * du.Hours
	}

	if rd >= MinuteNanoSeconds {
		du.Minutes = rd / MinuteNanoSeconds
		rd -= MinuteNanoSeconds * du.Minutes
	}

	if rd >= SecondNanoseconds {
		du.Seconds = rd / SecondNanoseconds
		rd -= SecondNanoseconds * du.Seconds
	}

	if rd >= MilliSecondNanoseconds {
		du.MilliSeconds = rd / MilliSecondNanoseconds
		rd -= MilliSecondNanoseconds * du.MilliSeconds
	}

	if rd >= MicroSecondNanoseconds {
		du.MicroSeconds = rd / MicroSecondNanoseconds
		rd -= MicroSecondNanoseconds * du.MicroSeconds
	}

	du.NanoSeconds = rd

	return
}

// CalculateGregorianYearDuration - Returns a string showing the
// breakdown of duration by Gregorian Years, Days, Hours, Minutes,
// Seconds, Milliseconds, Microseconds and Nanoseconds. Unlike
// other calculations which use a Standard 365-day year consisting
// of 24-hour days, a Gregorian Year consists of 365 days, 5-hours,
// 59-minutes and 12 Seconds. For the Gregorian calendar the
// average length of the calendar year (the mean year) across
// the complete leap cycle of 400 years is 365.2425 days.
// Sources:
// https://en.wikipedia.org/wiki/Year
// Source: https://en.wikipedia.org/wiki/Gregorian_calendar

func (du *DurationUtility) CalculateGregorianYearDuration() string {

	rd := int64(du.TimeDuration)

	yearsElement := "0-Gregorian Years "
	daysElement := "0-Days "
	hoursElement := "0-Hours "
	minutesElement := "0-Minutes "
	secondsElement := "0-Seconds "
	millisecondsElement := "0-Milliseconds "
	microsecondsElement := "0-Microseconds "
	nanosecondsElement := "0-NanoSeconds"

	if rd == 0 {
		return yearsElement +
			daysElement +
			hoursElement +
			minutesElement +
			secondsElement +
			millisecondsElement +
			microsecondsElement +
			nanosecondsElement
	}

	if rd > GregorianYearNanoSeconds {
		y := rd / GregorianYearNanoSeconds
		yearsElement = fmt.Sprintf("%v-Gregorian Years ", y)
		rd -= GregorianYearNanoSeconds * y
	}

	if rd >= DayNanoSeconds {
		days := rd / DayNanoSeconds
		rd -= DayNanoSeconds * days
		daysElement = fmt.Sprintf("%v-Days ", days)
	}

	if rd >= HourNanoSeconds {
		hrs := rd / HourNanoSeconds
		rd -= HourNanoSeconds * hrs
		hoursElement = fmt.Sprintf("%v-Hours ", hrs)
	}

	if rd >= MinuteNanoSeconds {
		min := rd / MinuteNanoSeconds
		rd -= MinuteNanoSeconds * min
		minutesElement = fmt.Sprintf("%v-Minutes ", min)

	}

	if rd >= SecondNanoseconds {
		secs := rd / SecondNanoseconds
		rd -= SecondNanoseconds * secs
		secondsElement = fmt.Sprintf("%v-Seconds ", secs)
	}

	if rd >= MilliSecondNanoseconds {
		milSec := rd / MilliSecondNanoseconds
		rd -= MilliSecondNanoseconds * milSec
		millisecondsElement = fmt.Sprintf("%v-Milliseconds ", milSec)
	}

	if rd >= MicroSecondNanoseconds {
		microSecs := rd / MicroSecondNanoseconds
		rd -= MicroSecondNanoseconds * microSecs
		microsecondsElement = fmt.Sprintf("%v-Microseconds ", du.MicroSeconds)

	}

	nanosecondsElement = fmt.Sprintf("%v-Nanoseconds", rd)

	return yearsElement +
		daysElement +
		hoursElement +
		minutesElement +
		secondsElement +
		millisecondsElement +
		microsecondsElement +
		nanosecondsElement
}

// CalculateTargetTimeFromPlusDuration - Calculates End Date Time by adding
// duration components to a Start Date Time.
func (du *DurationUtility) CalculateTargetTimeFromPlusDuration(startDateTime time.Time,
	years int64, months int64, days int64, hours int64, minutes int64, seconds int64,
	milliseconds int64, microseconds int64, nanoseconds int64) {

	aYears := int(math.Abs(float64(years)))
	aMonths := int(math.Abs(float64(months)))
	aDays := int(math.Abs(float64(days)))
	aHours := int64(math.Abs(float64(hours)))
	aMinutes := int64(math.Abs(float64(minutes)))
	aSeconds := int64(math.Abs(float64(seconds)))
	aMilliseconds := int64(math.Abs(float64(milliseconds)))
	aMicroseconds := int64(math.Abs(float64(microseconds)))
	aNanoseconds := int64(math.Abs(float64(nanoseconds)))

	intermediateDate := time.Time{}

	if aYears != 0 || aMonths != 0 || aDays != 0 {
		intermediateDate = startDateTime.AddDate(aYears, aMonths, aDays)
	} else {
		intermediateDate = startDateTime
	}

	iDur := intermediateDate.Sub(startDateTime)

	var dns int64

	dns = aHours * HourNanoSeconds
	dns += aMinutes * MinuteNanoSeconds
	dns += aSeconds * SecondNanoseconds
	dns += aMilliseconds * MilliSecondNanoseconds
	dns += aMicroseconds * MicroSecondNanoseconds
	dns += aNanoseconds
	dns += int64(iDur)

	du.Empty()

	du.TimeDuration = time.Duration(dns)
	du.StartDateTime = startDateTime
	du.EndDateTime = du.StartDateTime.Add(du.TimeDuration)

	du.CalculateDurationElements()
	du.CalculateDurationStrings()

	return

}

// CalculateTargetTimeFromMinusDuration - Calculates Start Date Time by subtracting
// duration components from an End Date Time. You must enter date time elements
// as positive numbers.
func (du *DurationUtility) CalculateTargetTimeFromMinusDuration(endDateTime time.Time,
	years int64, months int64, days int64, hours int64, minutes int64, seconds int64,
	milliseconds int64, microseconds int64, nanoseconds int64) {

	aYears := int(math.Abs(float64(years)))
	aMonths := int(math.Abs(float64(months)))
	aDays := int(math.Abs(float64(days)))
	aHours := int64(math.Abs(float64(hours)))
	aMinutes := int64(math.Abs(float64(minutes)))
	aSeconds := int64(math.Abs(float64(seconds)))
	aMilliseconds := int64(math.Abs(float64(milliseconds)))
	aMicroseconds := int64(math.Abs(float64(microseconds)))
	aNanoseconds := int64(math.Abs(float64(nanoseconds)))

	intermediateDate := time.Time{}

	if aYears != 0 || aMonths != 0 || aDays != 0 {
		intermediateDate = endDateTime.AddDate(-aYears, -aMonths, -aDays)
	} else {
		intermediateDate = endDateTime
	}

	iDur := endDateTime.Sub(intermediateDate)

	var dns int64

	dns = aHours * HourNanoSeconds
	dns += aMinutes * MinuteNanoSeconds
	dns += aSeconds * SecondNanoseconds
	dns += aMilliseconds * MilliSecondNanoseconds
	dns += aMicroseconds * MicroSecondNanoseconds
	dns += aNanoseconds
	dns += int64(iDur)

	du.Empty()
	du.TimeDuration = time.Duration(dns)
	du.EndDateTime = endDateTime
	du.StartDateTime = du.EndDateTime.Add(-du.TimeDuration)
	du.CalculateDurationElements()
	du.CalculateDurationStrings()

	return
}

// CopyToThis - Receives and incoming DurationUtility data
// structure and copies the values to the current DurationUtility
// data structure.
func (du *DurationUtility) CopyToThis(duIn DurationUtility) {
	du.Empty()
	du.TimeDuration = duIn.TimeDuration
	du.StartDateTime = duIn.StartDateTime
	du.EndDateTime = duIn.EndDateTime
	du.Years = duIn.Years
	du.YearsNanosecs = duIn.YearsNanosecs
	du.Months = duIn.Months
	du.MonthNanosecs = duIn.MonthNanosecs
	du.Days = duIn.Days
	du.Hours = duIn.Hours
	du.Minutes = duIn.Minutes
	du.Seconds = duIn.Seconds
	du.MilliSeconds = duIn.MilliSeconds
	du.MicroSeconds = duIn.MicroSeconds
	du.NanoSeconds = duIn.NanoSeconds
	du.NanosecStr = duIn.NanosecStr
	du.DaysStr = duIn.DaysStr
	du.HoursStr = duIn.HoursStr
	du.DurationStr = duIn.DurationStr
	du.DefaultStr = duIn.DefaultStr

	return
}

// Equal - This method may be used to determine if two
// DurationUtility data structures are equivalent.
func (du *DurationUtility) Equal(duIn DurationUtility) bool {

	if du.TimeDuration != duIn.TimeDuration ||
		du.StartDateTime != duIn.StartDateTime ||
		du.EndDateTime != duIn.EndDateTime ||
		du.Years != duIn.Years ||
		du.YearsNanosecs != duIn.YearsNanosecs ||
		du.Months != duIn.Months ||
		du.MonthNanosecs != duIn.MonthNanosecs ||
		du.Days != duIn.Days ||
		du.Hours != duIn.Hours ||
		du.Minutes != duIn.Minutes ||
		du.Seconds != duIn.Seconds ||
		du.MilliSeconds != duIn.MilliSeconds ||
		du.MicroSeconds != duIn.MicroSeconds ||
		du.NanoSeconds != duIn.NanoSeconds ||
		du.NanosecStr != duIn.NanosecStr ||
		du.HoursStr != duIn.HoursStr ||
		du.DaysStr != duIn.DaysStr ||
		du.DurationStr != duIn.DurationStr ||
		du.DefaultStr != duIn.DefaultStr {
		return false
	}

	return true

}

// Empty - This method initializes
// all of the fields in this
// DurationUtility structure to their
// zero values.
func (du *DurationUtility) Empty() {
	du.TimeDuration = time.Duration(0)
	du.StartDateTime = time.Time{}
	du.EndDateTime = time.Time{}
	du.Years = 0
	du.YearsNanosecs = 0
	du.Months = 0
	du.MonthNanosecs = 0
	du.Days = 0
	du.Hours = 0
	du.Minutes = 0
	du.Seconds = 0
	du.MilliSeconds = 0
	du.MicroSeconds = 0
	du.NanoSeconds = 0
	du.NanosecStr = ""
	du.HoursStr = ""
	du.DaysStr = ""
	du.DurationStr = ""
	du.DefaultStr = ""

}

// EmptyCalcs - Zero's all elements of the
// DurationUtility structure except duration
// time data elements.
func (du *DurationUtility) EmptyCalcs() {
	du.NanosecStr = ""
	du.HoursStr = ""
	du.DaysStr = ""
	du.DurationStr = ""
	du.DefaultStr = ""
}

// GetDuration - Returns a time.Duration structure defining the duration between
// input parameters startTime and endTime
func (du *DurationUtility) GetDuration(startTime time.Time, endTime time.Time) (DurationUtility, error) {

	if startTime.Equal(endTime) {
		return DurationUtility{}, nil
	}

	if endTime.Before(startTime) {
		return DurationUtility{}, errors.New("DurationUtility.GetDuration() Error: endTime less than startTime")
	}

	d2 := DurationUtility{StartDateTime:startTime, EndDateTime:endTime, TimeDuration: endTime.Sub(startTime)}
	d2.CalculateDurationElements()
	d2.CalculateDurationStrings()
	return d2, nil
}

func (du DurationUtility) GetDurationBySeconds(seconds int64) time.Duration {

	return time.Duration(seconds) * time.Second

}

func (du DurationUtility) GetDurationByMinutes(minutes int64) time.Duration {

	return time.Duration(minutes) * time.Minute

}

func (du DurationUtility) GenerateDuration(duIn DurationUtility) (time.Duration, error) {
	return du.GetDurationFromElapsedTime(duIn)
}

// GetDurationBreakDown - Receives a Duration type
// and returns a breakdown of duration by years,
// days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.
// NOTE: Years is arbitrarily set to the number of
// nanoseconds in a standard 365-day year. The Years
// calculation does NOT take Leap Years into account.
func (du DurationUtility) GetDurationBreakDown(startTime time.Time, d time.Duration) DurationUtility {

	if startTime.IsZero() {
		startTime = time.Now().UTC()
	}

	durationUtility := DurationUtility{StartDateTime:startTime, TimeDuration: d}
	durationUtility.CalculateDurationElements()
	durationUtility.CalculateDurationStrings()
	return durationUtility

}

// GetDurationFromElapsedTime - Receives an incoming DurationUtility and extracts passed
// Years, Days, Minutes, Seconds, Milliseconds, MicroSeconds and NanoSeconds to compute
// time.Duration value.
func (du DurationUtility) GetDurationFromElapsedTime(elapsedTime DurationUtility) (time.Duration, error) {
	var dns int64

	dns = elapsedTime.YearsNanosecs
	dns += elapsedTime.MonthNanosecs
	dns += elapsedTime.Days * DayNanoSeconds
	dns += elapsedTime.Hours * HourNanoSeconds
	dns += elapsedTime.Minutes * MinuteNanoSeconds
	dns += elapsedTime.Seconds * SecondNanoseconds
	dns += elapsedTime.MilliSeconds * MilliSecondNanoseconds
	dns += elapsedTime.MicroSeconds * MicroSecondNanoseconds
	dns += elapsedTime.NanoSeconds

	timeDuration := time.Duration(dns)

	return timeDuration, nil

}

// GetElapsedTime - calculates the elapsed time
// between input parameters startTime and endTime.
// The result is returned in an DurationUtility
// structure.
func (du DurationUtility) GetElapsedTime(startTime time.Time, endTime time.Time) (DurationUtility, error) {

	dur, err := du.GetDuration(startTime, endTime)

	if err != nil {
		s := "DateTimeUtility-GetElapsedTime Error: " + err.Error()

		return DurationUtility{}, errors.New(s)
	}

	return dur, nil

}

// GetTimePlusDuration - Returns time plus input duration as a time.Time type.
func (du *DurationUtility) GetTimePlusDuration(tStartTime time.Time, duration time.Duration) time.Time {

	return tStartTime.Add(duration)
}

// GetTimeMinusDuration - Returns time minus input duration as a time.Type type.
func (du DurationUtility) GetTimeMinusDuration(tStartTime time.Time, duration time.Duration) time.Time {

	return tStartTime.Add(-duration)
}

func (du *DurationUtility) SetStartTimeDuration(startDateTime time.Time, duration time.Duration)  error {

	if startDateTime.IsZero() {
		return fmt.Errorf("DurationUtility.SetStartTimeDuration() Error - Start Time is Zero!")
	}

	x := int64(duration)

	du.Empty()

	if x < 0 {
		du.StartDateTime =  startDateTime.Add(duration)
		du.EndDateTime = startDateTime
		du.TimeDuration =  time.Duration(int64(math.Abs(float64(int64(duration)))))
	} else {
		du.StartDateTime = startDateTime
		du.EndDateTime = startDateTime.Add(duration)
		du.TimeDuration = duration
	}

	du.CalculateDurationElements()
	du.CalculateDurationStrings()

	return nil

}

