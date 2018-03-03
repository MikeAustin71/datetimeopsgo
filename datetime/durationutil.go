package datetime

import (
	"errors"
	"fmt"
	"time"
	"strings"
)

/*
	Library Location
  ================
  The Duration Utilities Library is located in source code repository:

		https://github.com/MikeAustin71/datetimeopsgo.git

	You will find this source file, 'durationutil.go' in the subdirectory:

			datetimeopsgo\datetime


	Dependencies
	============

	(1) 'timezonedto.go'
	(2)	'durationtimedto.go'



	Overview and Usage
	==================

	The principal component of this library is the DurationUtility. This
	type plus associated methods are used to manage, calculate and analyze
	time duration calculations.

	Usage generally involves providing start time and and end time as inputs. Thereafter
	the 'DurationUtility' calculates time duration and provides multiple formats
  to displaying the resulting time duration.

	When providing start times and end times, methods usually require standardized
	time zones or time zone locations.  If duration calculations are performed using
	start and end times with differing time zone locations, the calculation could
	produce unexpected results. are provided with are provided with


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
// 	Usage:
//	======
//
// du, err := DurationUtility{}.NewStartEndTimes(startDateTime, endDateTime)
//
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting time
//
// endDateTime		time.Time - Ending time
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time comparisons.
//
// 														Time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (du DurationUtility) NewStartEndTimes(startDateTime,
				endDateTime time.Time, timeZoneLocation, dateTimeFmtStr string ) (DurationUtility, error) {

	ePrefix := "DurationUtility.NewStartTimeDuration() "

	du2 := DurationUtility{}

	err := du2.SetStartEndTimes(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)

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
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting Date Time for duration calculation
//
// duration time.Duration 	- Time Duration added to 'startDatTime' in order to
//														compute Ending Date Time
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time comparisons.
//
// 														Time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (du DurationUtility) NewStartTimeDuration(startDateTime time.Time, duration time.Duration,
						timeZoneLocation, dateTimeFmtStr string ) (DurationUtility, error) {


	ePrefix := "DurationUtility.NewStartTimeDuration() "

	du2 := DurationUtility{}

	err := du2.SetStartTimeDuration(startDateTime, duration, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DurationUtility{}, fmt.Errorf(ePrefix + "Error returned from du2.SetStartTimeDuration(startDateTime, duration).\nError='%v'", err)
	}

	return du2, nil
}


// NewStartTimeMinusTime - Returns a new DurationUtility based on two input parameters,
// 'startDateTime' and 'timeDto'. 'timeDto' is an instance of TimeDto which is
// subtracted from 'startDateTime' in order to calculate time duration. The user is
// required to provide Time Zone Location as an input parameter in order to ensure that
// time duration calculations are performed using equivalent time zones.
//
// Usage:
//
// du, err := DurationUtility{}.NewStartTimeMinusTime(startDateTime, timeDto)
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting time
//
// minusTimeDto		  TimeDto - Provides time values which will be subtracted from
//														'startDateTime' in order to calculate duration.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time comparisons.
//
// 														Time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (du DurationUtility) NewStartTimeMinusTime(startDateTime time.Time, minusTimeDto TimeDto,
														timeZoneLocation, dateTimeFmtStr string) (DurationUtility, error){

ePrefix := "DurationUtility.NewStartTimeMinusTime() "

	du2 := DurationUtility{}

	err := du2.SetStartTimeMinusTime(startDateTime, minusTimeDto, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DurationUtility{}, fmt.Errorf(ePrefix + "Error returned from du2.SetStartTimeMinusTime(startDateTime, minusTimeDto).\nError='%v'", err)
	}

	return du2, nil
}

// NewStartTimePlusTime - Returns a New DurationUtility based on 'startDateTime'
// and DurationDto input parameters.
//
// Usage:
// du, err := DurationUtility{}.NewStartTimePlusTime(startDateTime, timeDto)
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting time
//
// plusTimeDto		  TimeDto - Provides time values which will be subtracted from
//														'startDateTime' in order to calculate duration.
//									type TimeDto struct {
//										Years          int // Number of Years
//										Months         int // Number of Months
//										Weeks          int // Number of Weeks
//										WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//										DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//										Hours          int // Number of Hours.
//										Minutes        int // Number of Minutes
//										Seconds        int // Number of Seconds
//										Milliseconds   int // Number of Milliseconds
//										Microseconds   int // Number of Microseconds
//										Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//										TotNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//																			// 	plus remaining Nanoseconds
//									}
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time comparisons.
//
// 														Time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (du DurationUtility) NewStartTimePlusTime(startDateTime time.Time, plusTimeDto TimeDto,
														timeZoneLocation, dateTimeFmtStr string) (DurationUtility, error) {

	ePrefix := "DurationUtility.NewStartTimePlusTime() "

	du2 := DurationUtility{}

	err := du2.SetStartTimePlusTime(startDateTime, plusTimeDto, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DurationUtility{}, fmt.Errorf(ePrefix + "Error returned from du2.SetStartTimePlusTime(startDateTime, plusTimeDto).\nError='%v'", err)
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
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting Date Time for duration calculation
//
// duration time.Duration 	- Time Duration added to 'startDatTime' in order to
//														compute Ending Date Time
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time comparisons.
//
// 														Time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (du *DurationUtility) SetStartTimeDuration(startDateTime time.Time,
				duration time.Duration, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "DurationUtility.SetStartTimeDuration() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error - Start Time is Zero!")
	}

	tzLoc := du.preProcessTimeZoneLocation(timeZoneLocation)
	dtFormat := du.preProcessDateFormatStr(dateTimeFmtStr)


	x := int64(duration)

	du.Empty()

	if x < 0 {
		eTimeTzu, err := TimeZoneDto{}.New(startDateTime,tzLoc, dtFormat)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(startDateTime,\"Local\"). startDateTime='%v'\nError='%v'", startDateTime, err.Error())
		}

		du.EndTimeTzu = eTimeTzu.CopyOut()

		du.StartTimeTzu, err = TimeZoneDto{}.NewAddDuration(eTimeTzu, duration, dtFormat)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDuration(eTimeTzu, duration). Error='%v'", err.Error())
		}

		du.TimeDuration = time.Duration(x * -1)

	} else if x == 0 {

		return fmt.Errorf(ePrefix + "Error - Input parameter 'duration' is Zero!")

	} else {

		var err error

		du.StartTimeTzu, err = TimeZoneDto{}.New(startDateTime,tzLoc, dtFormat)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned from TimeZoneDto{}.New(startDateTime,\"Local\") Error='%v'", err.Error())
		}

		du.EndTimeTzu, err = TimeZoneDto{}.NewAddDuration(du.StartTimeTzu, duration, dtFormat)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned from TimeZoneDto{}.NewAddDuration(du.StartTimeDateTz, duration). Error='%v'", err.Error())
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
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting time
//
// endDateTime		time.Time - Ending time
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time comparisons.
//
// 														Time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (du *DurationUtility) SetStartEndTimes(startDateTime,
					endDateTime time.Time, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "DurationUtility.SetStartEndTimes() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'startDateTime' is ZERO!")
	}

	if endDateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'endDateTime' is ZERO!")
	}

	tzLoc := du.preProcessTimeZoneLocation(timeZoneLocation)
	dtFormat := du.preProcessDateFormatStr(dateTimeFmtStr)

	du.Empty()

	sTime, err := TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(startDateTime, \"Local\"). Error='%v'", err.Error())
	}

	eTime, err := TimeZoneDto{}.New(endDateTime, tzLoc, dtFormat)

	if eTime.TimeOut.DateTime.Before(sTime.TimeOut.DateTime) {
		s2 := sTime.CopyOut()
		sTime = eTime.CopyOut()
		eTime = s2.CopyOut()
	}

	du.StartTimeTzu = sTime.CopyOut()

	du.EndTimeTzu = eTime.CopyOut()

	duration := du.EndTimeTzu.TimeOut.DateTime.Sub(du.StartTimeTzu.TimeOut.DateTime)

	du.TimeDuration = duration

	err = du.IsDurationBaseDataValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "ERROR: Duration Base Data is INVALID! Error='%v'", err.Error())
	}

	return nil
}

// SetStartTimeMinusTime - Calculate duration values based on a Starting Date Time and
// time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'timeDto' parameter. The time values in the 'timeDto' parameter are subtracted
// from 'StartTimeDateTz'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time.
//
// As a result. true values for StartTimeDateTz, EndTimeDateTz and
// TimeDuration are stored in the DurationUtility data structure.
// In other words, the input 'startDateTime' becomes the EndTimeDateTz and
// 'startDateTime' is calculated.
//
func (du *DurationUtility) SetStartTimeMinusTime(startDateTime time.Time,
		minusTimeDto TimeDto, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "DurationUtility.SetStartTimeMinusTime() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'startDateTime' is ZERO VALUE")
	}

	tzLoc := du.preProcessTimeZoneLocation(timeZoneLocation)
	dtFormat := du.preProcessDateFormatStr(dateTimeFmtStr)

	du.Empty()

	minusTimeDto.ConvertToNegativeValues()

	dur := DurationDto{}
	var err error

	dur.StartTimeTzu, err = TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat)

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


	sTime, err := TimeZoneDto{}.NewAddDuration(dur.StartTimeTzu, tDur, dtFormat)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDuration(dur.StartTimeDateTz, tDur). Error='%v'", err.Error())
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


// SetStartTimePlusTime - Calculate duration values based on a Starting Date Time and
// time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'times' parameter. The 'timeDto' parameter is added to
// 'StartTimeDateTz'.
//
// Values in the 'timeDto' parameter are automatically converted to positive
// numeric values before being added to 'StartTimeDateTz'.
//
// True values for StartTimeDateTz, EndTimeDateTz and TimeDuration are
// then stored in the DurationUtility data structure.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting time
//
// plusTimeDto		  TimeDto - Provides time values which will be subtracted from
//														'startDateTime' in order to calculate duration.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time comparisons.
//
// 														Time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (du *DurationUtility) SetStartTimePlusTime(startDateTime time.Time, plusTimeDto TimeDto,
	timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "DurationUtility.SetStartTimePlusTime() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'startDateTime' is ZERO VALUE!")
	}

	tzLoc := du.preProcessTimeZoneLocation(timeZoneLocation)
	dtFormat := du.preProcessDateFormatStr(dateTimeFmtStr)

	du.Empty()

	plusTimeDto.ConvertToAbsoluteValues()

	var err error
	dur := DurationDto{}

	dur.StartTimeTzu, err = TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat)

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

	du.EndTimeTzu, err = TimeZoneDto{}.NewAddDuration(du.StartTimeTzu, du.TimeDuration, dtFormat)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDuration(du.StartTimeDateTz, du.TimeDuration). Error='%v'", err.Error())
	}

	err = du.IsDurationBaseDataValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: Duration Base Data INVALID! Error='%v'", err.Error())
	}

	return nil
}

// IsDurationBaseDataValid - Validates DurationUtility.TimeDuration
// DurationUtility.StartTimeDateTz and DurationUtility.EndTimeDateTz.
// Note: if DurationUtility.StartTimeDateTz and DurationUtility.EndTimeDateTz
// have zero values, DurationUtility.StartTimeDateTz will be defaulted to
// time.Now().UTC()
func (du *DurationUtility) IsDurationBaseDataValid() error {

	ePrefix := "DurationUtility.IsDurationBaseDataValid() "

	rd := int64(du.TimeDuration)

	err := du.StartTimeTzu.IsTimeZoneDtoValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error. StartTimeDateTz is INVALID! Error='%v'", err.Error())
	}

	err = du.EndTimeTzu.IsTimeZoneDtoValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error. EndTimeDateTz is INVALID! Error='%v'", err.Error())
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

	for yearDateTimeTzu.TimeOut.DateTime.Before(dDto.EndTimeTzu.TimeOut.DateTime) {

		i++

		yearDateTimeTzu, err = TimeZoneDto{}.NewAddDate(dDto.StartTimeTzu,i,0,0, FmtDateTimeYrMDayFmtStr)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDate(dDto.StartTimeDateTz,i,0,0). Error='%v'", err.Error())
		}

	}

	i -= 1

	if i > 0 {

		dDto.Years = int64(i)

		yearDateTimeTzu, err = TimeZoneDto{}.NewAddDate(dDto.StartTimeTzu, i, 0, 0, FmtDateTimeYrMDayFmtStr)

		duration := yearDateTimeTzu.TimeOut.DateTime.Sub(dDto.StartTimeTzu.TimeOut.DateTime)

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
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDuration(dDto.StartTimeDateTz,time.Duration(dDto.YearsNanosecs)). Error='%v'", err.Error())
	}

	mthDateTimeTzu := yearDateTimeTzu.CopyOut()

	for mthDateTimeTzu.TimeOut.DateTime.Before(dDto.EndTimeTzu.TimeOut.DateTime) {

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
			return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDate(dDto.StartTimeDateTz, int(dDto.Years), int(dDto.Months), 0). Error='%v'", err.Error())
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
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDate(dDto.StartTimeDateTz, int(dDto.Years), 0, 0 ). Error='%v'", err.Error())
	}

	duration, err := yearDateTimeTzu.Sub(dDto.StartTimeTzu)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by yearDateTimeTzu.Sub(dDto.StartTimeDateTz). Error='%v'", err.Error())
	}

	dDto.YearsNanosecs = int64(duration)

	return nil

}

func (du *DurationUtility) calcFromMonths(dDto *DurationDto) error {

	ePrefix := "DurationUtility.calcFromMonths() "

	mthStartDateTimeTzu, err := TimeZoneDto{}.NewAddDuration(dDto.StartTimeTzu,
															time.Duration(dDto.YearsNanosecs), FmtDateTimeYrMDayFmtStr)
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.NewAddDuration(dDto.StartTimeDateTz,	time.Duration(dDto.YearsNanosecs)). Error='%v'", err.Error())
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

func (du *DurationUtility) preProcessDateFormatStr(dateTimeFmtStr string) string {

	if len(dateTimeFmtStr) == 0 {
		return FmtDateTimeYrMDayFmtStr
	}

	return dateTimeFmtStr
}


func (du *DurationUtility) preProcessTimeZoneLocation(timeZoneLocation string) string {

	if len(timeZoneLocation) == 0 {
		return TzIanaUTC
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		return "Local"
	}

	return timeZoneLocation
}