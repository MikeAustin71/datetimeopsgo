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

	WeekNanoSeconds = int64(7) * DayNanoSeconds

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

	// StdYearNanoSeconds - Number of Nanoseconds in a 365-day year
	StdYearNanoSeconds = DayNanoSeconds * 365
)

// TimesDto - used for transmitting
// time elements.
type TimesDto struct {
	Years        int64
	Months       int64
	Weeks        int64
	Days         int64
	Hours        int64
	Minutes      int64
	Seconds      int64
	Milliseconds int64
	Microseconds int64
	Nanoseconds  int64
}

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
	Weeks         int64
	WeekDays      int64
	Days          int64
	Hours         int64
	Minutes       int64
	Seconds       int64
	Milliseconds  int64
	Microseconds  int64
	Nanoseconds   int64
	// NanosecStr - Example: 2-Days 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds
	NanosecStr string

	// DurationStr - Example: 12-Years 3-Months 2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
	DurationStr string

	// YearsMthsWeeksStr - Example: 12-Years 3-Months 2-Weeks 1-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
	YearsMthsWeeksStr string

	// CumWeeksStr - Example: 126-Weeks 1-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
	CumWeeksStr string

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

// CalcDurationFromTimes - Calculate a duration from
// 'startTime' and 'endTime' values passed to this method.
// Results are stored int he current DurationUtility data
// structure.
func (du *DurationUtility) CalcDurationFromTimes(startTime time.Time, endTime time.Time) error {

	if startTime.IsZero() {
		return errors.New("DurationUtility.CalcDurationFromTimes() Error: startTime has ZERO time value!")
	}

	if endTime.IsZero() {
		return errors.New("DurationUtility.CalcDurationFromTimes() Error: endTime has ZERO time value!")
	}

	if startTime == endTime {
		return errors.New("DurationUtility.CalcDurationFromTimes() Error: starTime and endTime are EQUAL!")
	}

	sTime := startTime
	eTime := endTime

	if endTime.Before(startTime) {
		sTime = endTime
		eTime = startTime
	}

	du.StartDateTime = sTime

	du.EndDateTime = eTime

	du.TimeDuration = du.EndDateTime.Sub(du.StartDateTime)

	du.CalcDurationElements()

	du.CalcDurationStrings()

	return nil
}

// CalcDurationFromElements -This method is used to calculate durations based on the
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
//               du.CalcDurationFromElements()
//////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (du *DurationUtility) CalcDurationFromElements() {

	if du.StartDateTime.IsZero() {
		du.StartDateTime = time.Now().UTC()
	}

	startDate := du.StartDateTime
	addTime := (du.Hours * HourNanoSeconds) + (du.Minutes * MinuteNanoSeconds) +
		(du.Seconds * SecondNanoseconds) + (du.Milliseconds * MilliSecondNanoseconds) +
		(du.Microseconds * MicroSecondNanoseconds) + du.Nanoseconds

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

	du.CalcDurationElements()
	du.CalcDurationStrings()

	return
}

// CalcDurationStrings - Calculates Duration breakdowns
// in a variety of display formats based on the Time Duration
// value stored in the existing DurationUtility.
func (du *DurationUtility) CalcDurationStrings() {

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

		du.YearsMthsWeeksStr = du.DurationStr

		du.NanosecStr = hoursElement + minutesElement +

			secondsElement + nanosecondsElement

		du.DaysStr = "0-Days " +
			hoursElement + minutesElement +
			secondsElement + millisecondsElement +
			microsecondsElement + nanosecondsElement

		du.CumWeeksStr = "0-Weeks " + du.DaysStr

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

	rn := (du.Milliseconds * MilliSecondNanoseconds) +
		(du.Microseconds * MicroSecondNanoseconds) +
		du.Nanoseconds

	totalNanoseconds := fmt.Sprintf("%v-Nanoseconds", rn)
	du.NanosecStr = yearsElement + monthsElement + daysElement +
		hoursElement + minutesElement + secondsElement +
		totalNanoseconds

	millisecondsElement = fmt.Sprintf("%v-Milliseconds ", du.Milliseconds)

	microsecondsElement = fmt.Sprintf("%v-Microseconds ", du.Microseconds)

	nanosecondsElement = fmt.Sprintf("%v-Nanoseconds", du.Nanoseconds)

	du.DurationStr = yearsElement + monthsElement + daysElement +
		hoursElement + minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	du.DaysStr, _ = du.CalcCumDaysDurationStr()

	du.HoursStr, _ = du.CalcCumHoursDurationStr()

	tDto := TimesDto{}

	du.YearsMthsWeeksStr, tDto, _ = du.CalcYearsMthsWeeksStr()

	du.Weeks = tDto.Weeks

	du.WeekDays = tDto.Days

	du.CumWeeksStr, _ = du.CalcCumWeeksDurationStr()

	return
}

func (du *DurationUtility) CalcYearsMthsWeeksStr() (string, TimesDto, error) {

	rd := int64(du.TimeDuration)

	if rd == 0 {
		return "", TimesDto{} ,fmt.Errorf("DurationUtility.CalcYearsMthsWeeksStr() ERROR - Duration Equals Zero!")
	}


	tDto := TimesDto{}

	tDto.Years = du.Years
	rd -= du.YearsNanosecs

	tDto.Months = du.Months
	rd -= du.MonthNanosecs

	if rd >= WeekNanoSeconds {
		tDto.Weeks = rd / WeekNanoSeconds
		rd -= tDto.Weeks * WeekNanoSeconds
	}

	if rd >= DayNanoSeconds {
		tDto.Days = rd / DayNanoSeconds
		rd -= DayNanoSeconds * tDto.Days
	}

	if rd >= HourNanoSeconds {
		tDto.Hours = rd / HourNanoSeconds
		rd -= HourNanoSeconds * tDto.Hours
	}

	if rd >= MinuteNanoSeconds {
		tDto.Minutes = rd / MinuteNanoSeconds
		rd -= MinuteNanoSeconds * tDto.Minutes
	}

	if rd >= SecondNanoseconds {
		tDto.Seconds = rd / SecondNanoseconds
		rd -= SecondNanoseconds * tDto.Seconds
	}

	if rd >= MilliSecondNanoseconds {
		tDto.Milliseconds = rd / MilliSecondNanoseconds
		rd -= MilliSecondNanoseconds * tDto.Milliseconds
	}

	if rd >= MicroSecondNanoseconds {
		tDto.Microseconds = rd / MicroSecondNanoseconds
		rd -= MicroSecondNanoseconds * tDto.Microseconds
	}

	tDto.Nanoseconds = rd

	yearsElement := fmt.Sprintf("%v-Years ", tDto.Years)

	monthsElement := fmt.Sprintf("%v-Months ", tDto.Months)

	weeksElement := fmt.Sprintf("%v-Weeks ", tDto.Weeks)

	daysElement := fmt.Sprintf("%v-Days ", tDto.Days)

	hoursElement := fmt.Sprintf("%v-Hours ", tDto.Hours)

	minutesElement := fmt.Sprintf("%v-Minutes ", tDto.Minutes)

	secondsElement := fmt.Sprintf("%v-Seconds ", tDto.Seconds)

	millisecondsElement := fmt.Sprintf("%v-Milliseconds ", tDto.Milliseconds)

	microsecondsElement := fmt.Sprintf("%v-Microseconds ", tDto.Microseconds)

	nanosecondsElement := fmt.Sprintf("%v-Nanoseconds", tDto.Nanoseconds)

	yearsMthsWeeksStr := yearsElement + monthsElement +
		weeksElement + daysElement +
		hoursElement + minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	return yearsMthsWeeksStr, tDto, nil

}

func (du *DurationUtility) CalcCumDaysDurationStr() (string, error) {
	tDto := TimesDto{}

	rd := int64(du.TimeDuration)

	if rd == 0 {
		return "", fmt.Errorf("DurationUtility.CalcCumDaysDurationStr() ERROR - Duration Equals Zero!")
	}

	if rd >= DayNanoSeconds {
		tDto.Days = rd / DayNanoSeconds
		rd -= DayNanoSeconds * tDto.Days
	}

	if rd >= HourNanoSeconds {
		tDto.Hours = rd / HourNanoSeconds
		rd -= HourNanoSeconds * tDto.Hours
	}

	if rd >= MinuteNanoSeconds {
		tDto.Minutes = rd / MinuteNanoSeconds
		rd -= MinuteNanoSeconds * tDto.Minutes
	}

	if rd >= SecondNanoseconds {
		tDto.Seconds = rd / SecondNanoseconds
		rd -= SecondNanoseconds * tDto.Seconds
	}

	if rd >= MilliSecondNanoseconds {
		tDto.Milliseconds = rd / MilliSecondNanoseconds
		rd -= MilliSecondNanoseconds * tDto.Milliseconds
	}

	if rd >= MicroSecondNanoseconds {
		tDto.Microseconds = rd / MicroSecondNanoseconds
		rd -= MicroSecondNanoseconds * tDto.Microseconds
	}

	tDto.Nanoseconds = rd

	daysElement := fmt.Sprintf("%v-Days ", tDto.Days)

	hoursElement := fmt.Sprintf("%v-Hours ", tDto.Hours)

	minutesElement := fmt.Sprintf("%v-Minutes ", tDto.Minutes)

	secondsElement := fmt.Sprintf("%v-Seconds ", tDto.Seconds)

	millisecondsElement := fmt.Sprintf("%v-Milliseconds ", tDto.Milliseconds)

	microsecondsElement := fmt.Sprintf("%v-Microseconds ", tDto.Microseconds)

	nanosecondsElement := fmt.Sprintf("%v-Nanoseconds", tDto.Nanoseconds)

	cumDaysStr := daysElement + hoursElement +
		minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	return cumDaysStr, nil

}

func (du *DurationUtility) CalcCumHoursDurationStr() (string, error) {
	tDto := TimesDto{}
	rd := int64(du.TimeDuration)

	if rd == 0 {
		return "", fmt.Errorf("DurationUtility.CalcCumHoursDurationStr() ERROR - Duration Equals Zero!")
	}

	if rd >= HourNanoSeconds {
		tDto.Hours = rd / HourNanoSeconds
		rd -= HourNanoSeconds * tDto.Hours
	}

	if rd >= MinuteNanoSeconds {
		tDto.Minutes = rd / MinuteNanoSeconds
		rd -= MinuteNanoSeconds * tDto.Minutes
	}

	if rd >= SecondNanoseconds {
		tDto.Seconds = rd / SecondNanoseconds
		rd -= SecondNanoseconds * tDto.Seconds
	}

	if rd >= MilliSecondNanoseconds {
		tDto.Milliseconds = rd / MilliSecondNanoseconds
		rd -= MilliSecondNanoseconds * tDto.Milliseconds
	}

	if rd >= MicroSecondNanoseconds {
		tDto.Microseconds = rd / MicroSecondNanoseconds
		rd -= MicroSecondNanoseconds * tDto.Microseconds
	}

	tDto.Nanoseconds = rd

	hoursElement := fmt.Sprintf("%v-Hours ", tDto.Hours)

	minutesElement := fmt.Sprintf("%v-Minutes ", tDto.Minutes)

	secondsElement := fmt.Sprintf("%v-Seconds ", tDto.Seconds)

	millisecondsElement := fmt.Sprintf("%v-Milliseconds ", tDto.Milliseconds)

	microsecondsElement := fmt.Sprintf("%v-Microseconds ", tDto.Microseconds)

	nanosecondsElement := fmt.Sprintf("%v-Nanoseconds", tDto.Nanoseconds)

	cumHoursStr := hoursElement + minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	return cumHoursStr, nil

}

func (du *DurationUtility) CalcCumWeeksDurationStr() (string, error) {
	tDto := TimesDto{}

	rd := int64(du.TimeDuration)

	if rd == 0 {
		return "", fmt.Errorf("DurationUtility.CalcCumWeeksDurationStr() ERROR - Duration Equals Zero!")
	}

	if rd >= WeekNanoSeconds {
		tDto.Weeks = rd / WeekNanoSeconds
		rd -= WeekNanoSeconds * tDto.Weeks
	}

	if rd >= DayNanoSeconds {
		tDto.Days = rd / DayNanoSeconds
		rd -= DayNanoSeconds * tDto.Days
	}

	if rd >= HourNanoSeconds {
		tDto.Hours = rd / HourNanoSeconds
		rd -= HourNanoSeconds * tDto.Hours
	}

	if rd >= MinuteNanoSeconds {
		tDto.Minutes = rd / MinuteNanoSeconds
		rd -= MinuteNanoSeconds * tDto.Minutes
	}

	if rd >= SecondNanoseconds {
		tDto.Seconds = rd / SecondNanoseconds
		rd -= SecondNanoseconds * tDto.Seconds
	}

	if rd >= MilliSecondNanoseconds {
		tDto.Milliseconds = rd / MilliSecondNanoseconds
		rd -= MilliSecondNanoseconds * tDto.Milliseconds
	}

	if rd >= MicroSecondNanoseconds {
		tDto.Microseconds = rd / MicroSecondNanoseconds
		rd -= MicroSecondNanoseconds * tDto.Microseconds
	}

	tDto.Nanoseconds = rd

	weeksElement := fmt.Sprintf("%v-Weeks ", tDto.Weeks)

	daysElement := fmt.Sprintf("%v-Days ", tDto.Days)

	hoursElement := fmt.Sprintf("%v-Hours ", tDto.Hours)

	minutesElement := fmt.Sprintf("%v-Minutes ", tDto.Minutes)

	secondsElement := fmt.Sprintf("%v-Seconds ", tDto.Seconds)

	millisecondsElement := fmt.Sprintf("%v-Milliseconds ", tDto.Milliseconds)

	microsecondsElement := fmt.Sprintf("%v-Microseconds ", tDto.Microseconds)

	nanosecondsElement := fmt.Sprintf("%v-Nanoseconds", tDto.Nanoseconds)

	cumWeeksStr := weeksElement + daysElement +
		hoursElement + minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	return cumWeeksStr, nil
}

// CalcDurationElements - Breaks down the duration
// value into Years, Days, Hours, Minutes, Seconds,
// Milliseconds, Microseconds and Nanoseconds. Calculation
// is based on the 'TimeDuration' value currently stored
// in the DurationUtility data structure.
func (du *DurationUtility) CalcDurationElements() {

	rd := int64(du.TimeDuration)

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
		du.Milliseconds = rd / MilliSecondNanoseconds
		rd -= MilliSecondNanoseconds * du.Milliseconds
	}

	if rd >= MicroSecondNanoseconds {
		du.Microseconds = rd / MicroSecondNanoseconds
		rd -= MicroSecondNanoseconds * du.Microseconds
	}

	du.Nanoseconds = rd

	return
}

// CalcGregorianYearDuration - Returns a string showing the
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

func (du *DurationUtility) CalcGregorianYearDuration() string {

	rd := int64(du.TimeDuration)

	yearsElement := "0-Gregorian Years "
	daysElement := "0-Days "
	hoursElement := "0-Hours "
	minutesElement := "0-Minutes "
	secondsElement := "0-Seconds "
	millisecondsElement := "0-Milliseconds "
	microsecondsElement := "0-Microseconds "
	nanosecondsElement := "0-Nanoseconds"

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
		microsecondsElement = fmt.Sprintf("%v-Microseconds ", du.Microseconds)

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

// CalcTargetTimeFromPlusDuration - Calculates End Date Time by adding
// duration components to a Start Date Time. Enter all date time element
// values as positive numbers.
func (du *DurationUtility) CalcTargetTimeFromPlusDuration(startDateTime time.Time,
	times TimesDto) {

	aYears := int(math.Abs(float64(times.Years)))
	aMonths := int(math.Abs(float64(times.Months)))
	aWeeks := int(math.Abs(float64(times.Weeks)))
	aDays := int(math.Abs(float64(times.Days + (int64(aWeeks * 7)))))
	aHours := int64(math.Abs(float64(times.Hours)))
	aMinutes := int64(math.Abs(float64(times.Minutes)))
	aSeconds := int64(math.Abs(float64(times.Seconds)))
	aMilliseconds := int64(math.Abs(float64(times.Milliseconds)))
	aMicroseconds := int64(math.Abs(float64(times.Microseconds)))
	aNanoseconds := int64(math.Abs(float64(times.Nanoseconds)))

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

	du.CalcDurationElements()
	du.CalcDurationStrings()

	return

}

// CalcTargetTimeFromMinusDuration - Calculates Start Date Time by subtracting
// duration components from an End Date Time. Enter date time element values
// as positive numbers.
func (du *DurationUtility) CalcTargetTimeFromMinusDuration(endDateTime time.Time,
	times TimesDto) {

	aYears := int(math.Abs(float64(times.Years)))
	aMonths := int(math.Abs(float64(times.Months)))
	aWeeks := int(math.Abs(float64(times.Weeks)))
	aDays := int(math.Abs(float64(times.Days + (int64(aWeeks * 7)))))
	aHours := int64(math.Abs(float64(times.Hours)))
	aMinutes := int64(math.Abs(float64(times.Minutes)))
	aSeconds := int64(math.Abs(float64(times.Seconds)))
	aMilliseconds := int64(math.Abs(float64(times.Milliseconds)))
	aMicroseconds := int64(math.Abs(float64(times.Microseconds)))
	aNanoseconds := int64(math.Abs(float64(times.Nanoseconds)))
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
	du.CalcDurationElements()
	du.CalcDurationStrings()

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
	du.Weeks = duIn.Weeks
	du.WeekDays = duIn.WeekDays
	du.Days = duIn.Days
	du.Hours = duIn.Hours
	du.Minutes = duIn.Minutes
	du.Seconds = duIn.Seconds
	du.Milliseconds = duIn.Milliseconds
	du.Microseconds = duIn.Microseconds
	du.Nanoseconds = duIn.Nanoseconds
	du.NanosecStr = duIn.NanosecStr
	du.YearsMthsWeeksStr = duIn.YearsMthsWeeksStr
	du.CumWeeksStr = duIn.CumWeeksStr
	du.DaysStr = duIn.DaysStr
	du.HoursStr = duIn.HoursStr
	du.DurationStr = duIn.DurationStr
	du.DefaultStr = duIn.DefaultStr

	return
}

// Copy - Returns a deep copy of the current
// DurationUtility data fields.
func (du *DurationUtility) Copy() DurationUtility {
	duOut := DurationUtility{}
	duOut.TimeDuration = du.TimeDuration
	duOut.StartDateTime = du.StartDateTime
	duOut.EndDateTime = du.EndDateTime
	duOut.Years = du.Years
	duOut.YearsNanosecs = du.YearsNanosecs
	duOut.Months = du.Months
	duOut.MonthNanosecs = du.MonthNanosecs
	duOut.Weeks = du.Weeks
	duOut.WeekDays = du.WeekDays
	duOut.Days = du.Days
	duOut.Hours = du.Hours
	duOut.Minutes = du.Minutes
	duOut.Seconds = du.Seconds
	duOut.Milliseconds = du.Milliseconds
	duOut.Microseconds = du.Microseconds
	duOut.Nanoseconds = du.Nanoseconds
	duOut.NanosecStr = du.NanosecStr
	duOut.YearsMthsWeeksStr = du.YearsMthsWeeksStr
	duOut.CumWeeksStr = du.CumWeeksStr
	duOut.DaysStr = du.DaysStr
	duOut.HoursStr = du.HoursStr
	duOut.DurationStr = du.DurationStr
	duOut.DefaultStr = du.DefaultStr

	return duOut
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
		du.Weeks != duIn.Weeks ||
		du.WeekDays != duIn.WeekDays ||
		du.Days != duIn.Days ||
		du.Hours != duIn.Hours ||
		du.Minutes != duIn.Minutes ||
		du.Seconds != duIn.Seconds ||
		du.Milliseconds != duIn.Milliseconds ||
		du.Microseconds != duIn.Microseconds ||
		du.Nanoseconds != duIn.Nanoseconds ||
		du.NanosecStr != duIn.NanosecStr ||
		du.HoursStr != duIn.HoursStr ||
		du.DaysStr != duIn.DaysStr ||
		du.YearsMthsWeeksStr != duIn.YearsMthsWeeksStr ||
		du.CumWeeksStr != duIn.CumWeeksStr ||
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
	du.Weeks = 0
	du.WeekDays = 0
	du.Days = 0
	du.Hours = 0
	du.Minutes = 0
	du.Seconds = 0
	du.Milliseconds = 0
	du.Microseconds = 0
	du.Nanoseconds = 0
	du.NanosecStr = ""
	du.HoursStr = ""
	du.DaysStr = ""
	du.YearsMthsWeeksStr = ""
	du.CumWeeksStr = ""
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

	d2 := DurationUtility{StartDateTime: startTime, EndDateTime: endTime, TimeDuration: endTime.Sub(startTime)}
	d2.CalcDurationElements()
	d2.CalcDurationStrings()
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

	durationUtility := DurationUtility{StartDateTime: startTime, TimeDuration: d}
	durationUtility.CalcDurationElements()
	durationUtility.CalcDurationStrings()
	return durationUtility

}

// GetDurationFromElapsedTime - Receives an incoming DurationUtility and extracts passed
// Years, Days, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds to compute
// time.Duration value.
func (du DurationUtility) GetDurationFromElapsedTime(elapsedTime DurationUtility) (time.Duration, error) {
	var dns int64

	dns = elapsedTime.YearsNanosecs
	dns += elapsedTime.MonthNanosecs
	dns += elapsedTime.Days * DayNanoSeconds
	dns += elapsedTime.Hours * HourNanoSeconds
	dns += elapsedTime.Minutes * MinuteNanoSeconds
	dns += elapsedTime.Seconds * SecondNanoseconds
	dns += elapsedTime.Milliseconds * MilliSecondNanoseconds
	dns += elapsedTime.Microseconds * MicroSecondNanoseconds
	dns += elapsedTime.Nanoseconds

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

// SetStartTimeDuration - Receives a starting date time and
// a time duration. The method then calculates the ending
// date time, duration and populates the DurationUtility
// data fields.
//
// The Method will except negative time durations. This means
// that the duration will be subtracted from the starting
// date time.
func (du *DurationUtility) SetStartTimeDuration(startDateTime time.Time, duration time.Duration) error {

	if startDateTime.IsZero() {
		return fmt.Errorf("DurationUtility.SetStartTimeDuration() Error - Start Time is Zero!")
	}

	x := int64(duration)

	du.Empty()

	if x < 0 {
		du.StartDateTime = startDateTime.Add(duration)
		du.EndDateTime = startDateTime
		du.TimeDuration = time.Duration(int64(math.Abs(float64(int64(duration)))))
	} else {
		du.StartDateTime = startDateTime
		du.EndDateTime = startDateTime.Add(duration)
		du.TimeDuration = duration
	}

	du.CalcDurationElements()
	du.CalcDurationStrings()

	return nil

}

// SetStartEndTimes - Calculate duration values and save the results in the DurationUtility
// data fields. Calculations are based on a starting date time and an ending date time passed
// to the method.
func (du *DurationUtility) SetStartEndTimes(startDateTime time.Time, endDateTime time.Time) error {

	err := du.CalcDurationFromTimes(startDateTime, endDateTime)

	if err != nil {
		return fmt.Errorf("DurationUtility.SetStartEndTimes() ERROR - %v", err.Error())
	}

	return nil
}

// SetStartTimePlusTimes - Calculate duration values based on a Starting Date Time and
// time values (years, months, weeks, days, hours, minutes etc.) passed to the method
// in the 'times' parameter.
func (du *DurationUtility) SetStartTimePlusTimes(startDateTime time.Time, times TimesDto) {

	du.CalcTargetTimeFromPlusDuration(startDateTime, times)
}

// SetStartTimePlusTimes - Calculate duration values based on a Starting Date Time and
// time values (years, months, weeks, days, hours, minutes etc.) passed to the method
// in the 'times' parameter.
func (du *DurationUtility) SetStartTimeMinusTimes(startDateTime time.Time, times TimesDto) {

	du.CalcTargetTimeFromMinusDuration(startDateTime, times)

}
