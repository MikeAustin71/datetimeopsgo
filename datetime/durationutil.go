package datetime

import (
	"errors"
	"fmt"
	"time"
)

/*
	Library Location
  ================
  The Duration Utilities Library is located in source code repository:

		https://github.com/MikeAustin71/datetimeopsgo.git

	You will find this source file, 'durationutil.go' in the subdirectory:

			datetimeopsgo\DurationTimeUtility


	Dependencies
	============

	'gotimezonedto.go'




	Overview and Usage
	==================

	The principal component of this library is the DurationUtility. This
	type plus associated methods is used to manage and control time
	duration calculations.

	Usage requires two operations:
	1. You must first initialize the DurationUtility type using one of the
		 four 'Set' methods shown below:
		 	a. SetStartTimeDuration() 		also NewStartTimeDuration()
		 	b. SetStartEndTimes()					also NewStartEndTimes()
		 	c. SetStartTimePlusTime()			also NewStartTimePlus()
		 	d. SetStartTimeMinusTime()		also NewStartTimeMinusTime()

	2. After the DurationUtility is initialized in step one above, you are free
		 to call any of the following 'Get' methods in order to return
		 formatted time durations. A call to any of these methods will
		 return a DurationDto which contains a record of the duration
		 calculation broken down by years, months, weeks, days, hours,
		 minutes, seconds, milliseconds, microseconds and nanoseconds.
		 In addition, the DurationDto contains a field named, 'DisplayStr'
		 which contains the formatted text version of the duration output.

			a. GetYearsMthDays()
			b. GetYearsMthsWeeksTime()
			c. GetWeeksDaysTime()
			d. GetDaysTime()
			e. GetHoursTime()
			f. GetYrMthsWkDayHourSecNanosecDuration()
			g. GetNanosecondsDuration()
			h. GetDefaultDuration()
			i. GetGregorianYearDuration()


*/



// DurationDto - This type holds duration
// values as time components.
type DurationDto struct {
	StartTimeTzu         TimeZoneDto
	EndTimeTzu           TimeZoneDto
	TimeDuration         time.Duration
	Years                int64
	YearsNanosecs        int64
	Months               int64
	MonthsNanosecs       int64
	Weeks                int64
	WeeksNanosecs        int64
	Days                 int64
	DaysNanosecs         int64
	Hours                int64
	HoursNanosecs        int64
	Minutes              int64
	MinutesNanosecs      int64
	Seconds              int64
	SecondsNanosecs      int64
	Milliseconds         int64
	MillisecondsNanosecs int64
	Microseconds         int64
	MicrosecondsNanosecs int64
	Nanoseconds          int64
	NanosecondsNanosecs  int64
	DisplayStr           string
}


// CalcTotalNanoSecs - Adds up all the time elements
// of the current DurationDto struct and converts
// the value to nanoseconds.
func (dDto *DurationDto) CalcTotalNanoSecs() int64 {
	ns := dDto.YearsNanosecs
	ns += dDto.MonthsNanosecs
	ns += dDto.WeeksNanosecs
	ns += dDto.DaysNanosecs
	ns += dDto.HoursNanosecs
	ns += dDto.MinutesNanosecs
	ns += dDto.SecondsNanosecs
	ns += dDto.MillisecondsNanosecs
	ns += dDto.MicrosecondsNanosecs
	ns += dDto.NanosecondsNanosecs

	return ns
}

// InitializeTime - Initializes a DurationDto Structure
// from input parameter type TimeDto.
func (dDto *DurationDto) InitializeTime(tDto TimeDto) {
	dDto.EmptyTimeValues()
	dDto.Years = int64(tDto.Years)
	dDto.Months = int64(tDto.Months)
	dDto.Weeks = int64(tDto.Weeks)
	dDto.Days = int64(tDto.WeekDays)
	dDto.Hours = int64(tDto.Hours)
	dDto.Minutes = int64(tDto.Minutes)
	dDto.Seconds = int64(tDto.Seconds)
	dDto.Milliseconds = int64(tDto.Milliseconds)
	dDto.Microseconds = int64(tDto.Microseconds)
	dDto.Nanoseconds = int64(tDto.Nanoseconds)
}

// Copy - Makes a deep copy of the current
// DurationDto struct and returns it as a new
// DurationDto instance.
func (dDto *DurationDto) Copy() DurationDto {
	d := DurationDto{}

	d.StartTimeTzu = dDto.StartTimeTzu
	d.EndTimeTzu = dDto.EndTimeTzu
	d.TimeDuration = dDto.TimeDuration
	d.Years = dDto.Years
	d.YearsNanosecs = dDto.YearsNanosecs
	d.Months = dDto.Months
	d.MonthsNanosecs = dDto.MonthsNanosecs
	d.Weeks = dDto.Weeks
	d.WeeksNanosecs = dDto.WeeksNanosecs
	d.Days = dDto.Days
	d.DaysNanosecs = dDto.DaysNanosecs
	d.Hours = dDto.Hours
	d.HoursNanosecs = dDto.HoursNanosecs
	d.Minutes = dDto.Minutes
	d.MinutesNanosecs = dDto.MinutesNanosecs
	d.Seconds = dDto.Seconds
	d.SecondsNanosecs = dDto.SecondsNanosecs
	d.Milliseconds = dDto.Milliseconds
	d.MillisecondsNanosecs = dDto.MillisecondsNanosecs
	d.Microseconds = dDto.Microseconds
	d.MicrosecondsNanosecs = dDto.MicrosecondsNanosecs
	d.Nanoseconds = dDto.Nanoseconds
	d.NanosecondsNanosecs = dDto.MicrosecondsNanosecs
	d.DisplayStr = dDto.DisplayStr

	return d
}

// Empty() - Resets all the time elements
// of the current DurationDto struct to their
// initial or 'zero' values.
func (dDto *DurationDto) Empty() {
	dDto.StartTimeTzu = TimeZoneDto{}
	dDto.EndTimeTzu = TimeZoneDto{}
	dDto.TimeDuration = time.Duration(0)
	dDto.Years = 0
	dDto.YearsNanosecs = 0
	dDto.Months = 0
	dDto.MonthsNanosecs = 0
	dDto.Weeks = 0
	dDto.WeeksNanosecs = 0
	dDto.Days = 0
	dDto.DaysNanosecs = 0
	dDto.Hours = 0
	dDto.HoursNanosecs = 0
	dDto.Minutes = 0
	dDto.MinutesNanosecs = 0
	dDto.Seconds = 0
	dDto.SecondsNanosecs = 0
	dDto.Milliseconds = 0
	dDto.MillisecondsNanosecs = 0
	dDto.Microseconds = 0
	dDto.MicrosecondsNanosecs = 0
	dDto.Nanoseconds = 0
	dDto.NanosecondsNanosecs = 0
	dDto.DisplayStr = ""
}

// EmptyTimeValues - Resets only the time values
// in the current DurationDto struct to their
// initial or 'zero' values.
func (dDto *DurationDto) EmptyTimeValues() {

	dDto.Years = 0
	dDto.YearsNanosecs = 0
	dDto.Months = 0
	dDto.MonthsNanosecs = 0
	dDto.Weeks = 0
	dDto.WeeksNanosecs = 0
	dDto.Days = 0
	dDto.DaysNanosecs = 0
	dDto.Hours = 0
	dDto.HoursNanosecs = 0
	dDto.Minutes = 0
	dDto.MinutesNanosecs = 0
	dDto.Seconds = 0
	dDto.SecondsNanosecs = 0
	dDto.Milliseconds = 0
	dDto.MillisecondsNanosecs = 0
	dDto.Microseconds = 0
	dDto.MicrosecondsNanosecs = 0
	dDto.Nanoseconds = 0
	dDto.NanosecondsNanosecs = 0
	dDto.DisplayStr = ""
}

// EmptyNanosecs - Resets all the Nanosecond
// values in the current DurationDto struct to
// zero.
func (dDto *DurationDto) EmptyNanosecs() {
	dDto.YearsNanosecs = 0
	dDto.MonthsNanosecs = 0
	dDto.WeeksNanosecs = 0
	dDto.DaysNanosecs = 0
	dDto.HoursNanosecs = 0
	dDto.MinutesNanosecs = 0
	dDto.SecondsNanosecs = 0
	dDto.MillisecondsNanosecs = 0
	dDto.MicrosecondsNanosecs = 0
	dDto.NanosecondsNanosecs = 0
	dDto.DisplayStr = ""
}

// Equal - Determines whether the input DurationDto is
// equal to the current DurationDto structure.
func (dDto *DurationDto) Equal(dto2 DurationDto) bool {

	if dDto.StartTimeTzu != dto2.StartTimeTzu ||
		dDto.EndTimeTzu != dto2.EndTimeTzu ||
		dDto.TimeDuration != dto2.TimeDuration ||
		dDto.Years != dto2.Years ||
		dDto.YearsNanosecs != dto2.YearsNanosecs ||
		dDto.Months != dto2.Months ||
		dDto.MonthsNanosecs != dto2.MonthsNanosecs ||
		dDto.Weeks != dto2.Weeks ||
		dDto.WeeksNanosecs != dto2.WeeksNanosecs ||
		dDto.Days != dto2.Days ||
		dDto.DaysNanosecs != dto2.DaysNanosecs ||
		dDto.Hours != dto2.Hours ||
		dDto.HoursNanosecs != dto2.HoursNanosecs ||
		dDto.Minutes != dto2.Minutes ||
		dDto.MinutesNanosecs != dto2.MinutesNanosecs ||
		dDto.Seconds != dto2.Seconds ||
		dDto.SecondsNanosecs != dto2.SecondsNanosecs ||
		dDto.Milliseconds != dto2.Milliseconds ||
		dDto.MillisecondsNanosecs != dto2.MillisecondsNanosecs ||
		dDto.Microseconds != dto2.Microseconds ||
		dDto.MicrosecondsNanosecs != dto2.MicrosecondsNanosecs ||
		dDto.Nanoseconds != dto2.Nanoseconds ||
		dDto.NanosecondsNanosecs != dto2.NanosecondsNanosecs ||
		dDto.DisplayStr != dto2.DisplayStr {

		return false
	}

	return true
}

// DurationUtility - holds elements of
// time duration
type DurationUtility struct {
	StartTimeTzu TimeZoneDto
	EndTimeTzu   TimeZoneDto
	TimeDuration time.Duration
}

// CopyIn - Receives and incoming DurationUtility data
// structure and copies the values to the current DurationUtility
// data structure.
func (du *DurationUtility) CopyToThis(duIn DurationUtility) {
	du.Empty()
	du.TimeDuration = duIn.TimeDuration
	du.StartTimeTzu = duIn.StartTimeTzu
	du.EndTimeTzu = duIn.EndTimeTzu

	return
}

// Copy - Returns a deep copy of the current
// DurationUtility data fields.
func (du *DurationUtility) Copy() DurationUtility {
	duOut := DurationUtility{}
	duOut.TimeDuration = du.TimeDuration
	duOut.StartTimeTzu = du.StartTimeTzu
	duOut.EndTimeTzu = du.EndTimeTzu

	return duOut
}

// Equal - This method may be used to determine if two
// DurationUtility data structures are equivalent.
func (du *DurationUtility) Equal(duIn DurationUtility) bool {

	if du.TimeDuration != duIn.TimeDuration ||
		du.StartTimeTzu != duIn.StartTimeTzu ||
		du.EndTimeTzu != duIn.EndTimeTzu {

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
	du.StartTimeTzu = TimeZoneDto{}
	du.EndTimeTzu = TimeZoneDto{}
}

// GetYearMthDaysTimeAbbrv - Abbreviated formatting of Years, Months,
// WeekDays, Hours, Minutes, Seconds, Milliseconds, Microseconds and
// Nanoseconds. At a minimum only Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds.
// Abbreviated Years Mths WeekDays Time Duration - Example Return:
// 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetYearMthDaysTimeAbbrv() (DurationDto, error) {

	ePrefix := "DurationUtility.GetYearMthDaysTimeAbbrv() "

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	err := du.calcYearsFromDuration(&dDto)

	if err != nil {
		return DurationDto{}, fmt.Errorf(ePrefix + "Error returned by du.calcYearsFromDuration(&dDto). Error='%v'", err.Error())
	}


	err = du.calcMonthsFromDuration(&dDto)

	if err != nil {
		return DurationDto{}, fmt.Errorf(ePrefix + "Error returned by du.calcMonthsFromDuration(&dDto). Error='%v'", err.Error())
	}

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	yearsElement := ""

	monthsElement := ""

	daysElement := ""

	if dDto.Years > 0 {
		yearsElement = fmt.Sprintf("%v-Years ", dDto.Years)
	}

	if dDto.Months > 0 {
		monthsElement = fmt.Sprintf("%v-Months ", dDto.Months)
	}

	if dDto.Days > 0 {
		daysElement = fmt.Sprintf("%v-WeekDays ", dDto.Days)
	}

	str := fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = yearsElement + monthsElement + daysElement + str

	return dDto, nil

}

// GetYearMthDaysTime - Calculates Duration and breakdowns
// time elements by Years, Months, days, hours, minutes,
// seconds, milliseconds, microseconds and nanoseconds.
// Example DisplayStr:
// Years Mths WeekDays Time Duration - Example Return:
// 12-Years 3-Months 2-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetYearMthDaysTime() (DurationDto, error) {

	ePrefix := "DurationUtility.GetYearMthDaysTime() "

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	err := du.calcYearsFromDuration(&dDto)

	if err != nil {
		return DurationDto{}, fmt.Errorf(ePrefix + "Error returned by du.calcYearsFromDuration(&dDto). Error='%v'", err.Error())
	}

	err = du.calcMonthsFromDuration(&dDto)

	if err != nil {
		return DurationDto{}, fmt.Errorf(ePrefix + "Error returned by du.calcMonthsFromDuration(&dDto). Error='%v'", err.Error())
	}

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	str := fmt.Sprintf("%v-Years ", dDto.Years)

	str += fmt.Sprintf("%v-Months ", dDto.Months)

	str += fmt.Sprintf("%v-WeekDays ", dDto.Days)

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil
}

// GetYearsMthsWeeksTimeAbbrv - Abbreviated formatting of Years, Months,
// Weeks, WeekDays, Hours, Minutes, Seconds, Milliseconds, Microseconds,
// Nanoseconds. At a minimum only Hours, Minutes, Seconds, Milliseconds,
// Microseconds, Nanoseconds are displayed. Example return when Years,
// Months, Weeks and WeekDays are zero:
// 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetYearsMthsWeeksTimeAbbrv() (DurationDto, error) {

	ePrefix := "DurationUtility.GetYearsMthsWeeksTimeAbbrv() "
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	err := du.calcYearsFromDuration(&dDto)

	if err != nil {
		return DurationDto{}, fmt.Errorf(ePrefix + "Error returned by du.calcYearsFromDuration(&dDto). Error='%v'", err.Error())
	}

	err = du.calcMonthsFromDuration(&dDto)

	if err != nil {
		return DurationDto{}, fmt.Errorf(ePrefix + "Error returned by du.calcMonthsFromDuration(&dDto). Error='%v'", err.Error())
	}

	du.calcWeeksFromDuration(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	yearsElement := ""

	monthsElement := ""

	weeksElement := ""

	daysElement := ""

	if dDto.Years > 0 {
		yearsElement = fmt.Sprintf("%v-Years ", dDto.Years)
	}

	if dDto.Months > 0 {
		monthsElement = fmt.Sprintf("%v-Months ", dDto.Months)
	}

	if dDto.Weeks > 0 {
		weeksElement = fmt.Sprintf("%v-Weeks ", dDto.Weeks)
	}

	if dDto.Days > 0 {
		daysElement = fmt.Sprintf("%v-WeekDays ", dDto.Days)
	}

	hoursElement := fmt.Sprintf("%v-Hours ", dDto.Hours)

	minutesElement := fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	secondsElement := fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	millisecondsElement := fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	microsecondsElement := fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	nanosecondsElement := fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = yearsElement + monthsElement +
		weeksElement + daysElement +
		hoursElement + minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	return dDto, nil

}

// GetYearsMthsWeeksTime - Example Return:
// 12-Years 3-Months 2-Weeks 1-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetYearsMthsWeeksTime() (DurationDto, error) {

	ePrefix := "DurationUtility.GetYearsMthsWeeksTime() "

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	err := du.calcYearsFromDuration(&dDto)

	if err != nil {
		return DurationDto{}, fmt.Errorf(ePrefix + "Error returned by du.calcYearsFromDuration(&dDto). Error='%v'", err.Error())
	}

	err = du.calcMonthsFromDuration(&dDto)

	if err != nil {
		return DurationDto{}, fmt.Errorf(ePrefix + "Error returned by du.calcMonthsFromDuration(&dDto). Error='%v'", err.Error())
	}


	du.calcWeeksFromDuration(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	yearsElement := fmt.Sprintf("%v-Years ", dDto.Years)

	monthsElement := fmt.Sprintf("%v-Months ", dDto.Months)

	weeksElement := fmt.Sprintf("%v-Weeks ", dDto.Weeks)

	daysElement := fmt.Sprintf("%v-WeekDays ", dDto.Days)

	hoursElement := fmt.Sprintf("%v-Hours ", dDto.Hours)

	minutesElement := fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	secondsElement := fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	millisecondsElement := fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	microsecondsElement := fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	nanosecondsElement := fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = yearsElement + monthsElement +
		weeksElement + daysElement +
		hoursElement + minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	return dDto, nil

}

// GetWeeksDaysTime - Example DisplayStr
// 126-Weeks 1-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetWeeksDaysTime() (DurationDto, error) {
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcWeeksFromDuration(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	str := ""

	str += fmt.Sprintf("%v-Weeks ", dDto.Weeks)

	str += fmt.Sprintf("%v-WeekDays ", dDto.Days)

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil

}

// GetDaysTime - Returns duration formatted as
// days, hours, minutes, seconds, milliseconds, microseconds,
// and nanoseconds.
// Example: 97-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetDaysTime() (DurationDto, error) {
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	str := ""

	str += fmt.Sprintf("%v-WeekDays ", dDto.Days)

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil
}

// GetHoursTime - Returns duration formatted as hours,
// minutes, seconds, milliseconds, microseconds, nanoseconds.
// Example: 152-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetHoursTime() (DurationDto, error) {
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	str := ""

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil

}

// GetYrMthWkDayHrMinSecNanosecs - Returns duration formatted
// as Year, Month, Day, Hour, Second and Nanoseconds.
// Example: 3-Years 2-Months 3-Weeks 2-WeekDays 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds
func (du *DurationUtility) GetYrMthWkDayHrMinSecNanosecs() (DurationDto, error) {

	ePrefix := "DurationUtility.GetYrMthWkDayHrMinSecNanosecs() "

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	err:= du.calcYearsFromDuration(&dDto)

	if err != nil {
		return DurationDto{}, fmt.Errorf(ePrefix + "Error returned by err:= du.calcYearsFromDuration(&dDto). Error='%v'", err.Error())
	}

	err = du.calcMonthsFromDuration(&dDto)

	if err != nil {
		return DurationDto{}, fmt.Errorf(ePrefix + "Error returned by du.calcMonthsFromDuration(&dDto). Error='%v'", err.Error())
	}

	du.calcWeeksFromDuration(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcNanoseconds(&dDto)

	str := ""

	str += fmt.Sprintf("%v-Years ", dDto.Years)

	str += fmt.Sprintf("%v-Months ", dDto.Months)

	str += fmt.Sprintf("%v-Weeks ", dDto.Weeks)

	str += fmt.Sprintf("%v-WeekDays ", dDto.Days)

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil
}

// GetNanosecondsDuration - Returns duration formatted as
// Nonseconds. DisplayStr shows Nanoseconds expressed as a
// 64-bit integer value.
func (du *DurationUtility) GetNanosecondsDuration() (DurationDto, error) {
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcNanoseconds(&dDto)

	str := fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil

}

// GetDefaultDuration returns duration formatted
// as nanoseconds. The DisplayStr shows the default
// string value for duration.
// Example: 61h26m46.864197832s
func (du *DurationUtility) GetDefaultDuration() (DurationDto, error) {

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	dDto.Nanoseconds = rd

	dur := time.Duration(rd)

	dDto.DisplayStr = fmt.Sprintf("%v", dur)

	return dDto, nil
}

// GetGregorianYearDuration - Returns a string showing the
// breakdown of duration by Gregorian Years, WeekDays, Hours, Minutes,
// Seconds, Milliseconds, Microseconds and Nanoseconds. Unlike
// other calculations which use a Standard 365-day year consisting
// of 24-hour days, a Gregorian Year consists of 365 days, 5-hours,
// 59-minutes and 12 Seconds. For the Gregorian calendar the
// average length of the calendar year (the mean year) across
// the complete leap cycle of 400 Years is 365.2425 days.
// Sources:
// https://en.wikipedia.org/wiki/Year
// Source: https://en.wikipedia.org/wiki/Gregorian_calendar
//
func (du *DurationUtility) GetGregorianYearDuration() (DurationDto, error) {

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	if rd > GregorianYearNanoSeconds {
		dDto.Years = rd / GregorianYearNanoSeconds
		dDto.YearsNanosecs = dDto.Years * GregorianYearNanoSeconds
	}

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	str := fmt.Sprintf("%v-Gregorian Years ", dDto.Years)

	str += fmt.Sprintf("%v-WeekDays ", dDto.Days)

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil

}

// GetDurationFromStartEndTimes - Computes the duration
// by subtracting Starting Date Time from the Ending Date
// time. No changes are made to or stored in the
// existing DurationUtility data structures.
func (du *DurationUtility) GetDurationFromStartEndTimes(startDateTime time.Time, endDateTime time.Time) (time.Duration, error) {

	ePrefix := "DurationUtility.GetDurationFromStartEndTimes() "

	if startDateTime.IsZero() {
		return time.Duration(0), errors.New(ePrefix + "ERROR - startDateTime is ZERO!")
	}

	if endDateTime.IsZero() {
		return time.Duration(0), errors.New(ePrefix + "ERROR - endDateTime is ZERO!")
	}

	startDateTimeTzu, err := TimeZoneDto{}.New(startDateTime, "Local", FmtDateTimeYrMDayFmtStr)

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix + "Error returned from TimeZoneDto{}.New(startDateTime, \"Local\"). Error='%v'", err.Error())
	}

	endDateTimeTzu, err := TimeZoneDto{}.New(endDateTime, "Local", FmtDateTimeYrMDayFmtStr)

	if err!= nil {
		return time.Duration(0), fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(endDateTime, \"Local\"). Error='%v'", err.Error())
	}

	duration, err := endDateTimeTzu.Sub(startDateTimeTzu)

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix + "Error returned by endDateTimeTzu.Sub(startDateTimeTzu). Error='%v'", err.Error())
	}

	return duration, nil
}

// GetDurationFromSeconds - returns a time Duration value
// based on the number of seconds passed to this method.
// No changes are made to or stored in the existing
// DurationUtility data structures.
func (du *DurationUtility) GetDurationFromSeconds(seconds int64) time.Duration {

	return time.Duration(seconds) * time.Second

}

// GetDurationFromMinutes - returns a time Duration value
// based on the number of minutes passed to this method.
// No changes are made to or stored in the existing
// DurationUtility data structures.
func (du DurationUtility) GetDurationFromMinutes(minutes int64) time.Duration {

	return time.Duration(minutes) * time.Minute

}

// NewStartEndTimes - Returns a New DurationUtility based on two input
// parameters, startDateTime and endDateTime.
//
// Usage:
//
// du, err := DurationUtility{}.NewStartEndTimes(startDateTime, endDateTime)
//
func (du DurationUtility) NewStartEndTimes(startDateTime time.Time, endDateTime time.Time) (DurationUtility, error) {

	ePrefix := "DurationUtility.NewStartTimeDuration() "

	du2 := DurationUtility{}

	err := du2.SetStartEndTimes(startDateTime, endDateTime)

	if err != nil {
		return DurationUtility{}, fmt.Errorf(ePrefix + "Error returned from du2.SetStartEndTimes(startDateTime, endDateTime).\nError='%v'", err)
	}

	return du2, nil

}

// NewStartTimeDuration - Returns a New DurationUtility based on 'startDateTime'
// and time.Duration input parameters.
//
// Usage:
//
// du, err := DurationUtility{}.NewStartTimeDuration(startDateTime, timeDto)
//
func (du DurationUtility) NewStartTimeDuration(startDateTime time.Time, duration time.Duration) (DurationUtility, error) {

	ePrefix := "DurationUtility.NewStartTimeDuration() "

	du2 := DurationUtility{}

	err := du2.SetStartTimeDuration(startDateTime, duration)

	if err != nil {
		return DurationUtility{}, fmt.Errorf(ePrefix + "Error returned from du2.SetStartTimeDuration(startDateTime, duration).\nError='%v'", err)
	}

	return du2, nil
}


// NewStartTimeMinusTime - Returns a new DurationUtility based on two input parameters,
// 'startDateTime' and 'timeDto'.
//
// Usage:
//
// du, err := DurationUtility{}.NewStartTimeMinusTime(startDateTime, timeDto)
//
func (du DurationUtility) NewStartTimeMinusTime(startDateTime time.Time, timeDto TimeDto) (DurationUtility, error){

ePrefix := "DurationUtility.NewStartTimeMinusTime() "

	du2 := DurationUtility{}

	err := du2.SetStartTimeMinusTime(startDateTime, timeDto)

	if err != nil {
		return DurationUtility{}, fmt.Errorf(ePrefix + "Error returned from du2.SetStartTimeMinusTime(startDateTime, timeDto).\nError='%v'", err)
	}

	return du2, nil
}

// NewStartTimePlusTime - Returns a New DurationUtility based on 'startDateTime'
// and DurationDto input parameters.
//
// Usage:
// du, err := DurationUtility{}.NewStartTimePlusTime(startDateTime, timeDto)
//
func (du DurationUtility) NewStartTimePlusTime(startDateTime time.Time, timeDto TimeDto) (DurationUtility, error) {

	ePrefix := "DurationUtility.NewStartTimePlusTime() "

	du2 := DurationUtility{}

	err := du2.SetStartTimePlusTime(startDateTime, timeDto)

	if err != nil {
		return DurationUtility{}, fmt.Errorf(ePrefix + "Error returned from du2.SetStartTimePlusTime(startDateTime, timeDto).\nError='%v'", err)
	}

	return du2, nil
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

	ePrefix := "DurationUtility.SetStartTimeDuration() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error - Start Time is Zero!")
	}

	x := int64(duration)

	du.Empty()

	if x < 0 {
		eTimeTzu, err := TimeZoneDto{}.New(startDateTime,"Local", FmtDateTimeYrMDayFmtStr)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(startDateTime,\"Local\"). startDateTime='%v'\nError='%v'", startDateTime, err.Error())
		}

		du.EndTimeTzu = eTimeTzu.CopyOut()

		du.StartTimeTzu, err = TimeZoneDto{}.NewAddDuration(eTimeTzu, duration, FmtDateTimeYrMDayFmtStr)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDuration(eTimeTzu, duration). Error='%v'", err.Error())
		}

		du.TimeDuration = time.Duration(x * -1)

	} else if x == 0 {

		return fmt.Errorf(ePrefix + "Error - Input parameter 'duration' is Zero!")

	} else {

		var err error

		du.StartTimeTzu, err = TimeZoneDto{}.New(startDateTime,"Local", FmtDateTimeYrMDayFmtStr)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned from TimeZoneDto{}.New(startDateTime,\"Local\") Error='%v'", err.Error())
		}

		du.EndTimeTzu, err = TimeZoneDto{}.NewAddDuration(du.StartTimeTzu, duration, FmtDateTimeYrMDayFmtStr)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned from TimeZoneDto{}.NewAddDuration(du.StartTimeTzu, duration). Error='%v'", err.Error())
		}

		du.TimeDuration = duration
	}

	err := du.IsDurationBaseDataValid()

	if err != nil {
		return fmt.Errorf("DurationUtility.SetStartEndTimes() ERROR - %v", err.Error())
	}

	return nil
}

// SetStartEndTimes - Calculate duration values and save the results in the DurationUtility
// data fields. Calculations are based on a starting date time and an ending date time passed
// to the method.
func (du *DurationUtility) SetStartEndTimes(startDateTime time.Time, endDateTime time.Time) error {

	ePrefix := "DurationUtility.SetStartEndTimes() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'startDateTime' is ZERO!")
	}

	if endDateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'endDateTime' is ZERO!")
	}


	du.Empty()

	sTime, err := TimeZoneDto{}.New(startDateTime, "Local", FmtDateTimeYrMDayFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(startDateTime, \"Local\"). Error='%v'", err.Error())
	}

	eTime, err := TimeZoneDto{}.New(endDateTime, "Local", FmtDateTimeYrMDayFmtStr)

	if eTime.TimeLocal.DateTime.Before(sTime.TimeLocal.DateTime) {
		s2 := sTime.CopyOut()
		sTime = eTime.CopyOut()
		eTime = s2.CopyOut()
	}

	du.StartTimeTzu = sTime.CopyOut()

	du.EndTimeTzu = eTime.CopyOut()

	duration, err := du.EndTimeTzu.Sub(du.StartTimeTzu)

	if err != nil {
		return fmt.Errorf("Error returned by du.EndTimeTzu.Sub(du.StartTimeTzu). Error='%v'", err.Error())
	}

	du.TimeDuration = duration

	err = du.IsDurationBaseDataValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "ERROR: Duration Base Data is INVALID! Error='%v'", err.Error())
	}

	return nil
}

// SetStartTimePlusTime - Calculate duration values based on a Starting Date Time and
// time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'times' parameter. The 'timeDto' parameter is added to
// 'StartTimeTzu'.
//
// Values in the 'timeDto' parameter are automatically converted to positive
// numeric values before being added to 'StartTimeTzu'.
//
// True values for StartTimeTzu, EndTimeTzu and TimeDuration are
// then stored in the DurationUtility data structure.
func (du *DurationUtility) SetStartTimePlusTime(startDateTime time.Time, plusTimeDto TimeDto) error {

	ePrefix := "DurationUtility.SetStartTimePlusTime() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'startDateTime' is ZERO VALUE!")
	}

	du.Empty()

	plusTimeDto.ConvertToAbsoluteValues()

	var err error
	dur := DurationDto{}

	dur.StartTimeTzu, err = TimeZoneDto{}.New(startDateTime, "Local", FmtDateTimeYrMDayFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(startDateTime, \"Local\"). Error='%v'", err.Error())
	}

	dur.InitializeTime(plusTimeDto)

	err = du.calcFromYears(&dur)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by du.calcFromYears(&dur). Error='%v'", err.Error())
	}

	err = du.calcFromMonths(&dur)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by err = du.calcFromMonths(&dur). Error='%v'", err.Error())
	}

	du.calcFromWeeks(&dur)
	du.calcFromDays(&dur)
	du.calcFromHoursMinSecs(&dur)
	du.calcFromMilliSecs(&dur)
	du.calcFromMicroSecs(&dur)
	du.calcFromNanoSecs(&dur)
	du.StartTimeTzu = dur.StartTimeTzu.CopyOut()
	du.TimeDuration = time.Duration(dur.CalcTotalNanoSecs())

	du.EndTimeTzu, err = TimeZoneDto{}.NewAddDuration(du.StartTimeTzu, du.TimeDuration, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDuration(du.StartTimeTzu, du.TimeDuration). Error='%v'", err.Error())
	}

	err = du.IsDurationBaseDataValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: Duration Base Data INVALID! Error='%v'", err.Error())
	}

	return nil
}

// SetStartTimeMinusTime - Calculate duration values based on a Starting Date Time and
// time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'timeDto' parameter. The time values in the 'timeDto' parameter are subtracted
// from 'StartTimeTzu'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time.
//
// As a result. true values for StartTimeTzu, EndTimeTzu and
// TimeDuration are stored in the DurationUtility data structure.
// In other words, the input 'startDateTime' becomes the EndTimeTzu and
// 'startDateTime' is calculated.
//
func (du *DurationUtility) SetStartTimeMinusTime(startDateTime time.Time, minusTimeDto TimeDto) error {

	ePrefix := "DurationUtility.SetStartTimeMinusTime() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'startDateTime' is ZERO VALUE")
	}

	du.Empty()
	minusTimeDto.ConvertToNegativeValues()

	dur := DurationDto{}
	var err error

	dur.StartTimeTzu, err = TimeZoneDto{}.New(startDateTime, "Local", FmtDateTimeYrMDayFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(startDateTime, \"Local\"). Error='%v'", err.Error())
	}

	dur.InitializeTime(minusTimeDto)

	err =	du.calcFromYears(&dur)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by du.calcFromYears(&dur). Error='%v'", err.Error())
	}

	err = du.calcFromMonths(&dur)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by du.calcFromMonths(&dur). Error='%v'", err.Error())
	}


	du.calcFromWeeks(&dur)
	du.calcFromDays(&dur)
	du.calcFromHoursMinSecs(&dur)
	du.calcFromMilliSecs(&dur)
	du.calcFromMicroSecs(&dur)
	du.calcFromNanoSecs(&dur)

	ns := dur.CalcTotalNanoSecs()
	tDur := time.Duration(ns)


	sTime, err := TimeZoneDto{}.NewAddDuration(dur.StartTimeTzu, tDur, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDuration(dur.StartTimeTzu, tDur). Error='%v'", err.Error())
	}

	du.StartTimeTzu = sTime.CopyOut()

	// Convert duration to positive value
	du.TimeDuration = time.Duration(ns * -1)
	du.EndTimeTzu = dur.StartTimeTzu.CopyOut()

	err = du.IsDurationBaseDataValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: Duration Base Data INVALID! Error='%v'", err.Error())
	}

	return nil
}

// IsDurationBaseDataValid - Validates DurationUtility.TimeDuration
// DurationUtility.StartTimeTzu and DurationUtility.EndTimeTzu.
// Note: if DurationUtility.StartTimeTzu and DurationUtility.EndTimeTzu
// have zero values, DurationUtility.StartTimeTzu will be defaulted to
// time.Now().UTC()
func (du *DurationUtility) IsDurationBaseDataValid() error {

	ePrefix := "DurationUtility.IsDurationBaseDataValid() "

	rd := int64(du.TimeDuration)

	err := du.StartTimeTzu.IsTimeZoneDtoValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error. StartTimeTzu is INVALID! Error='%v'", err.Error())
	}

	err = du.EndTimeTzu.IsTimeZoneDtoValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error. EndTimeTzu is INVALID! Error='%v'", err.Error())
	}

	if rd < 0 {
		return fmt.Errorf(ePrefix + "Error. Duration is less than zero. DurationUtility.TimeDuration='%v'", rd)
	}

	return nil
}

// calcBaseData - Validates Time Duration
func (du *DurationUtility) calcBaseData(dDto *DurationDto) error {
	dDto.TimeDuration = du.TimeDuration
	dDto.StartTimeTzu = du.StartTimeTzu.CopyOut()
	dDto.EndTimeTzu = du.EndTimeTzu.CopyOut()
	return nil
}


// calcYearsFromDuration - Calculates the absolute number of years
// based on duration.
func (du *DurationUtility) calcYearsFromDuration(dDto *DurationDto) error {

	ePrefix := "DurationUtility.calcYearsFromDuration() "
	yearDateTimeTzu := dDto.StartTimeTzu.CopyOut()

	i := 0

	var err error

	for yearDateTimeTzu.TimeLocal.DateTime.Before(dDto.EndTimeTzu.TimeLocal.DateTime) {

		i++

		yearDateTimeTzu, err = TimeZoneDto{}.NewAddDate(dDto.StartTimeTzu,i,0,0, FmtDateTimeYrMDayFmtStr)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDate(dDto.StartTimeTzu,i,0,0). Error='%v'", err.Error())
		}

	}

	i -= 1

	if i > 0 {

		dDto.Years = int64(i)

		yearDateTimeTzu, err = TimeZoneDto{}.NewAddDate(dDto.StartTimeTzu, i, 0, 0, FmtDateTimeYrMDayFmtStr)

		duration, err := yearDateTimeTzu.Sub(dDto.StartTimeTzu)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by yearDateTimeTzu.Sub(dDto.StartTimeTzu). Erro='%v'", err.Error())
		}

		dDto.YearsNanosecs = int64(duration)
	}

	return nil
}

func (du *DurationUtility) calcMonthsFromDuration(dDto *DurationDto) error {

	ePrefix := "DurationUtility.calcMonthsFromDuration() "

	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs

	i := 0

	yearDateTimeTzu, err := TimeZoneDto{}.NewAddDuration(dDto.StartTimeTzu,time.Duration(dDto.YearsNanosecs),FmtDateTimeYrMDayFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDuration(dDto.StartTimeTzu,time.Duration(dDto.YearsNanosecs)). Error='%v'", err.Error())
	}

	mthDateTimeTzu := yearDateTimeTzu.CopyOut()

	for mthDateTimeTzu.TimeLocal.DateTime.Before(dDto.EndTimeTzu.TimeLocal.DateTime) {

		i++

		mthDateTimeTzu, err = TimeZoneDto{}.NewAddDate(yearDateTimeTzu, 0, i, 0, FmtDateTimeYrMDayFmtStr)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDate(yearDateTimeTzu, 0, i, 0). i='%v'  Error='%v'", i, err.Error())
		}

	}

	i -= 1

	if i > 0 {

		dDto.Months = int64(i)

		mthDateTimeTzu, err = TimeZoneDto{}.NewAddDate(yearDateTimeTzu, 0, int(dDto.Months), 0, FmtDateTimeYrMDayFmtStr)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDate(dDto.StartTimeTzu, int(dDto.Years), int(dDto.Months), 0). Error='%v'", err.Error())
		}

		duration, err := mthDateTimeTzu.Sub(yearDateTimeTzu)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by mthDateTimeTzu.Sub(yearDateTimeTzu). Error='%v'", err.Error())
		}

		dDto.MonthsNanosecs = int64(duration)

	}

	return nil
}

func (du *DurationUtility) calcWeeksFromDuration(dDto *DurationDto) {
	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs

	if rd >= WeekNanoSeconds {
		dDto.Weeks = rd / WeekNanoSeconds
		dDto.WeeksNanosecs = dDto.Weeks * WeekNanoSeconds
	}

	return
}

func (du *DurationUtility) calcDaysFromDuration(dDto *DurationDto) {

	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs +
		dDto.WeeksNanosecs

	if rd >= DayNanoSeconds {
		dDto.Days = rd / DayNanoSeconds
		dDto.DaysNanosecs = DayNanoSeconds * dDto.Days
	}
}

func (du *DurationUtility) calcHoursMinSecs(dDto *DurationDto) {

	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs +
		dDto.WeeksNanosecs + dDto.DaysNanosecs

	if rd >= HourNanoSeconds {
		dDto.Hours = rd / HourNanoSeconds
		dDto.HoursNanosecs = HourNanoSeconds * dDto.Hours
		rd -= dDto.HoursNanosecs
	}

	if rd >= MinuteNanoSeconds {
		dDto.Minutes = rd / MinuteNanoSeconds
		dDto.MinutesNanosecs = MinuteNanoSeconds * dDto.Minutes
		rd -= dDto.MinutesNanosecs
	}

	if rd >= SecondNanoseconds {
		dDto.Seconds = rd / SecondNanoseconds
		dDto.SecondsNanosecs = SecondNanoseconds * dDto.Seconds
		rd -= dDto.SecondsNanosecs
	}

}

func (du *DurationUtility) calcMilliSeconds(dDto *DurationDto) {
	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs +
		dDto.WeeksNanosecs + dDto.DaysNanosecs + dDto.HoursNanosecs +
		dDto.MinutesNanosecs + dDto.SecondsNanosecs

	if rd >= MilliSecondNanoseconds {
		dDto.Milliseconds = rd / MilliSecondNanoseconds
		dDto.MillisecondsNanosecs = MilliSecondNanoseconds * dDto.Milliseconds
	}

}

func (du *DurationUtility) calcMicroSeconds(dDto *DurationDto) {
	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs +
		dDto.WeeksNanosecs + dDto.DaysNanosecs + dDto.HoursNanosecs +
		dDto.MinutesNanosecs + dDto.SecondsNanosecs +
		dDto.MillisecondsNanosecs

	if rd >= MicroSecondNanoseconds {
		dDto.Microseconds = rd / MicroSecondNanoseconds
		dDto.MicrosecondsNanosecs = MicroSecondNanoseconds * dDto.Microseconds
	}

}

func (du *DurationUtility) calcNanoseconds(dDto *DurationDto) {

	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs +
		dDto.WeeksNanosecs + dDto.DaysNanosecs + dDto.HoursNanosecs +
		dDto.MinutesNanosecs + dDto.SecondsNanosecs +
		dDto.MillisecondsNanosecs + dDto.MicrosecondsNanosecs

	dDto.Nanoseconds = rd

}

func (du *DurationUtility) calcFromYears(dDto *DurationDto) error {

	ePrefix := "DurationUtility.calcFromYears() "


	yearDateTimeTzu, err := TimeZoneDto{}.NewAddDate(dDto.StartTimeTzu, int(dDto.Years), 0, 0, FmtDateTimeYrMDayFmtStr )

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDate(dDto.StartTimeTzu, int(dDto.Years), 0, 0 ). Error='%v'", err.Error())
	}

	duration, err := yearDateTimeTzu.Sub(dDto.StartTimeTzu)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by yearDateTimeTzu.Sub(dDto.StartTimeTzu). Error='%v'", err.Error())
	}

	dDto.YearsNanosecs = int64(duration)

	return nil

}

func (du *DurationUtility) calcFromMonths(dDto *DurationDto) error {

	ePrefix := "DurationUtility.calcFromMonths() "

	mthStartDateTimeTzu, err := TimeZoneDto{}.NewAddDuration(dDto.StartTimeTzu,
															time.Duration(dDto.YearsNanosecs), FmtDateTimeYrMDayFmtStr)
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDuration(dDto.StartTimeTzu,	time.Duration(dDto.YearsNanosecs)). Error='%v'", err.Error())
	}

	mthEndDateTimeTzu, err := TimeZoneDto{}.NewAddDate (mthStartDateTimeTzu,0, int(dDto.Months), 0, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDate (mthStartDateTimeTzu,0, int(dDto.Months), 0). dDto.Months='%v'  Error='%v'", dDto.Months, err.Error())
	}

	duration, err := mthEndDateTimeTzu.Sub(mthStartDateTimeTzu)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by mthEndDateTimeTzu.Sub(mthStartDateTimeTzu). Error='%v'", err.Error())
	}

	dDto.MonthsNanosecs = int64(duration)

	return nil
}

func (du *DurationUtility) calcFromWeeks(dDto *DurationDto) {

	dDto.WeeksNanosecs = dDto.Weeks * WeekNanoSeconds
}

func (du *DurationUtility) calcFromDays(dDto *DurationDto) {
	dDto.DaysNanosecs = dDto.Days * DayNanoSeconds
}

func (du *DurationUtility) calcFromHoursMinSecs(dDto *DurationDto) {
	dDto.HoursNanosecs = dDto.Hours * HourNanoSeconds
	dDto.MinutesNanosecs = dDto.Minutes * MinuteNanoSeconds
	dDto.SecondsNanosecs = dDto.Seconds * SecondNanoseconds
}

func (du *DurationUtility) calcFromMilliSecs(dDto *DurationDto) {
	dDto.MillisecondsNanosecs = dDto.Milliseconds * MilliSecondNanoseconds
}

func (du *DurationUtility) calcFromMicroSecs(dDto *DurationDto) {
	dDto.MicrosecondsNanosecs = dDto.Microseconds * MicroSecondNanoseconds
}

func (du *DurationUtility) calcFromNanoSecs(dDto *DurationDto) {
	dDto.NanosecondsNanosecs = dDto.Nanoseconds
}
