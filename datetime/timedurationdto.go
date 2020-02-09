package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)


// TimeDurationDto
//
// Overview and Usage
//
// The 'TimeDurationDto' Type is used to calculate, store and transmit
// time duration information. It is therefore designed to work with
// incremental time or time duration.
//
// The Type includes a starting date time, ending date time, a time
// duration value specifying the time span between starting and ending
// date time and a breakdown of time duration by a series of time
// components including years, months, weeks, days, hours, minutes,
// seconds, milliseconds, microseconds and nanoseconds.
//
// Dependencies
//
// Starting and ending date times are stored as 'DateTzDto' types. The
// 'DateTzDto' type may be found in source code file:
//   MikeAustin71\datetimeopsgo\datetime\datetzdto.go
//
// Source Code Location
//
// This source file is located in source code repository:
//   https://github.com/MikeAustin71/datetimeopsgo.git
//
// The location of this source file is:
//   MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
type TimeDurationDto struct {
	StartTimeDateTz DateTzDto     // Starting Date Time with Time Zone info
	EndTimeDateTz   DateTzDto     // Ending Date Time with Time Zone info
	TimeDuration    time.Duration // Elapsed time or duration between starting and ending date time
	CalcType        TDurCalcType  // The calculation Type. This controls the allocation of time
	// 		duration over years, months, weeks, days and hours.
	Years                int64 // Number of Years
	YearsNanosecs        int64 // Number of Years in Nanoseconds
	Months               int64 // Number of Months
	MonthsNanosecs       int64 // Number of Months in Nanoseconds
	Weeks                int64 // Number of Weeks: Date Days / 7
	WeeksNanosecs        int64 // Number of Weeks in Nanoseconds
	WeekDays             int64 // WeekDays = DateDays - (Weeks * 7)
	WeekDaysNanosecs     int64 // Equivalent WeekDays in NanoSeconds
	DateDays             int64 // Day Number in Month (1-31)
	DateDaysNanosecs     int64 // DateDays in equivalent nanoseconds
	Hours                int64 // Number of Hours
	HoursNanosecs        int64 // Number of Hours in Nanoseconds
	Minutes              int64 // Number of Minutes
	MinutesNanosecs      int64 // Number of Minutes in Nanoseconds
	Seconds              int64 // Number of Seconds
	SecondsNanosecs      int64 // Number of Seconds in Nanoseconds
	Milliseconds         int64 // Number of Milliseconds
	MillisecondsNanosecs int64 // Number of Milliseconds in Nanoseconds
	Microseconds         int64 // Number of Microseconds
	MicrosecondsNanosecs int64 // Number of Microseconds in Nanoseconds
	Nanoseconds          int64 // Number of Nanoseconds (Remainder after Milliseconds & Microseconds)
	TotSubSecNanoseconds int64 // Equivalent Nanoseconds for Milliseconds + Microseconds + Nanoseconds
	TotDateNanoseconds   int64 // Equal to Years + Months + DateDays in equivalent nanoseconds.
	TotTimeNanoseconds   int64 // Equal to Hours + Seconds + Milliseconds + Microseconds + Nanoseconds in
	//                             in equivalent nanoseconds

	lock                 sync.Mutex
}

// CopyIn - Receives a TimeDurationDto as an input parameters
// and proceeds to set all data fields of the current TimeDurationDto
// equal to the incoming TimeDurationDto.
//
// When this method completes, the current TimeDurationDto will
// equal in all respects to the incoming TimeDurationDto.
func (tDur *TimeDurationDto) CopyIn(t2Dur TimeDurationDto) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.copyIn(
		tDur,
		&t2Dur,
		"TimeDurationDto.CopyIn() ")
}

// copyOut - Returns a deep copy of the current
// TimeDurationDto instance.
func (tDur *TimeDurationDto) CopyOut() TimeDurationDto {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.copyOut(
							tDur,
						"TimeDurationDto.CopyOut() ")
}

// Empty - Resets all of the current TimeDurationDto
// data fields to their zero or uninitialized values.
func (tDur *TimeDurationDto) Empty() {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()
	
	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.empty(
		tDur,
		"TimeDurationDto.Empty() ")
}

// EmptyTimeFields - Sets all of the data fields
// associated with time duration allocation to zero.
func (tDur *TimeDurationDto) EmptyTimeFields() {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()
	
	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.emptyTimeFields(
		tDur,
		"TimeDurationDto.EmptyTimeFields() ")

}

// Equal - Compares two TimeDurationDto instances to determine
// if they are equivalent.
func (tDur *TimeDurationDto) Equal(t2Dur TimeDurationDto) bool {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	if !tDur.StartTimeDateTz.Equal(t2Dur.StartTimeDateTz) ||
		!tDur.EndTimeDateTz.Equal(t2Dur.EndTimeDateTz) ||
		tDur.TimeDuration != t2Dur.TimeDuration ||
		tDur.CalcType != t2Dur.CalcType ||
		tDur.Years != t2Dur.Years ||
		tDur.YearsNanosecs != t2Dur.YearsNanosecs ||
		tDur.Months != t2Dur.Months ||
		tDur.MonthsNanosecs != t2Dur.MonthsNanosecs ||
		tDur.Weeks != t2Dur.Weeks ||
		tDur.WeeksNanosecs != t2Dur.WeeksNanosecs ||
		tDur.WeekDays != t2Dur.WeekDays ||
		tDur.WeekDaysNanosecs != t2Dur.WeekDaysNanosecs ||
		tDur.DateDays != t2Dur.DateDays ||
		tDur.DateDaysNanosecs != t2Dur.DateDaysNanosecs ||
		tDur.Hours != t2Dur.Hours ||
		tDur.HoursNanosecs != t2Dur.HoursNanosecs ||
		tDur.Minutes != t2Dur.Minutes ||
		tDur.MinutesNanosecs != t2Dur.MinutesNanosecs ||
		tDur.Seconds != t2Dur.Seconds ||
		tDur.SecondsNanosecs != t2Dur.SecondsNanosecs ||
		tDur.Milliseconds != t2Dur.Milliseconds ||
		tDur.MillisecondsNanosecs != t2Dur.MillisecondsNanosecs ||
		tDur.Microseconds != t2Dur.Microseconds ||
		tDur.MicrosecondsNanosecs != t2Dur.MicrosecondsNanosecs ||
		tDur.Nanoseconds != t2Dur.MillisecondsNanosecs ||
		tDur.TotSubSecNanoseconds != t2Dur.TotSubSecNanoseconds ||
		tDur.TotDateNanoseconds != t2Dur.TotDateNanoseconds ||
		tDur.TotTimeNanoseconds != t2Dur.TotTimeNanoseconds {

		return false
	}

	return true

}

// IsEmpty() - Returns 'true' if the current TimeDurationDto
// instance is uninitialized and consists entirely of zero values.
func (tDur *TimeDurationDto) IsEmpty() bool {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	if tDur.StartTimeDateTz.IsEmpty() &&
		tDur.EndTimeDateTz.IsEmpty() &&
		tDur.TimeDuration == 0 &&
		tDur.Years == 0 &&
		tDur.YearsNanosecs == 0 &&
		tDur.Months == 0 &&
		tDur.MonthsNanosecs == 0 &&
		tDur.Weeks == 0 &&
		tDur.WeeksNanosecs == 0 &&
		tDur.WeekDays == 0 &&
		tDur.WeekDaysNanosecs == 0 &&
		tDur.DateDays == 0 &&
		tDur.DateDaysNanosecs == 0 &&
		tDur.Hours == 0 &&
		tDur.HoursNanosecs == 0 &&
		tDur.Minutes == 0 &&
		tDur.MinutesNanosecs == 0 &&
		tDur.Seconds == 0 &&
		tDur.SecondsNanosecs == 0 &&
		tDur.Milliseconds == 0 &&
		tDur.MillisecondsNanosecs == 0 &&
		tDur.Microseconds == 0 &&
		tDur.MicrosecondsNanosecs == 0 &&
		tDur.Nanoseconds == 0 &&
		tDur.TotSubSecNanoseconds == 0 &&
		tDur.TotDateNanoseconds == 0 &&
		tDur.TotTimeNanoseconds == 0 {

		tDur.CalcType = TDurCalcType(0).None()
		return true
	}

	return false
}

// IsValid - Returns an error value signaling whether
// the current TimeDurationDto data fields are valid.
func (tDur *TimeDurationDto) IsValid() error {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.IsValid() "

	tDurDtoUtil := timeDurationDtoUtility{}
	
	return tDurDtoUtil.isValid(tDur, ePrefix)
}

// GetDurationFromTime - Calculates and returns a cumulative duration based on
// input parameters consisting of time elements.
func (tDur TimeDurationDto) GetDurationFromTime(
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int) time.Duration {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	dtMech :=DTimeMechanics{}

	return dtMech.GetDurationFromTimeComponents(
		0,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds)
}

// GetDurationFromDays - returns a time Duration value
// based on the number of days passed into this method.
// No changes are made to or stored in the existing
// TimeDurationDto data structures.
func (tDur TimeDurationDto) GetDurationFromDays(days int64) time.Duration {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return time.Duration(days*24) * time.Hour

}

// GetDurationFromHours - returns a time Duration value
// based on the number of hours passed into this method.
// No changes are made to or stored in the existing
// TimeDurationDto data structures.
func (tDur TimeDurationDto) GetDurationFromHours(hours int64) time.Duration {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return time.Duration(hours) * time.Hour

}

// GetDurationFromMinutes - returns a time Duration value
// based on the number of minutes passed into this method.
// No changes are made to or stored in the existing
// TimeDurationDto data structures.
func (tDur TimeDurationDto) GetDurationFromMinutes(minutes int64) time.Duration {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return time.Duration(minutes) * time.Minute

}

// GetDurationFromSeconds - returns a time Duration value
// based on the number of seconds passed to this method.
// No changes are made to or stored in the existing
// TimeDurationDto data structures.
func (tDur TimeDurationDto) GetDurationFromSeconds(seconds int64) time.Duration {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return time.Duration(seconds) * time.Second
}

// GetElapsedTimeStr - Provides a quick means for formatting Years, Months,
// DateDays, Hours, Minutes, Seconds, Milliseconds, Microseconds and
// Nanoseconds. At a minimum only Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds.
//
// This method only returns date time elements with value greater than zero.
// As a minimum, the string will display Nanoseconds.
//
// Example Return:
//
//  864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetElapsedTimeStr() string {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetElapsedTimeStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	if t2Dur.CalcType != TDurCalcType(0).StdYearMth() {
		err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).StdYearMth())
		if err != nil {
			return fmt.Sprintf(ePrefix +
				"\nError returned by t2Dur.ReCalcTimeDurationAllocation("+
				"TDurCalcType(0).StdYearMth())\nError='%v'\n", err.Error())
		}
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.DateDays > 0 || str != "" {
		str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)
	}

	if t2Dur.Hours > 0 || str != "" {

		str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	}

	if t2Dur.Minutes > 0 || str != "" {

		str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	}

	if t2Dur.Seconds > 0 || str != "" {

		str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	}

	if t2Dur.Milliseconds > 0 || str != "" {

		str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	}

	if t2Dur.Microseconds > 0 || str != "" {

		str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	}

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str

}

// GetElapsedMinutesStr - Provides a quick means for formatting Years, Months,
// DateDays, Hours, Minutes, Seconds, Milliseconds, Microseconds and
// Nanoseconds. At a minimum only Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds.
//
// This method only returns years, months, days or hours if those values
// are greater than zero.
//
// As a minimum the display string will show minutes, seconds, milliseconds,
// microseconds and nanoseconds.
//
// Example Return:
//
//  0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetElapsedMinutesStr() string {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetElapsedMinutesStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	if t2Dur.CalcType != TDurCalcType(0).StdYearMth() {
		err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).StdYearMth())

		if err != nil {
			return fmt.Sprintf(ePrefix +
				"\nError returned by t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).StdYearMth()).\n"+
				"Error='%v'\n", err.Error())
		}

	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.DateDays > 0 || str != "" {
		str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)
	}

	if t2Dur.Hours > 0 || str != "" {

		str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	}

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str

}

// GetYearMthDaysTimeAbbrvStr - Abbreviated formatting of Years, Months,
// DateDays, Hours, Minutes, Seconds, Milliseconds, Microseconds and
// Nanoseconds. At a minimum only Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds.
//
// Abbreviated Years Mths DateDays Time Duration - Example Return:
//
// 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYearMthDaysTimeAbbrvStr() string {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetYearMthDaysTimeAbbrvStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	if t2Dur.CalcType != TDurCalcType(0).StdYearMth() {

		err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).StdYearMth())

		if err != nil {
			return fmt.Sprintf(ePrefix +
				"\nError returned by t2Dur." +
				"ReCalcTimeDurationAllocation(" +
				"TDurCalcType(0).StdYearMth()).\n"+
				"Error='%v'\n", err.Error())
		}
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.DateDays > 0 || str != "" {
		str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)
	}

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str

}

// GetYearMthDaysTimeStr - Calculates Duration and breakdowns
// time elements by Years, Months, Date Days, hours, minutes,
// seconds, milliseconds, microseconds and nanoseconds.
//
// Example DisplayStr
// ==================
//
// Years Months DateDays Time Duration - Example Return:
//
// 12-Years 3-Months 2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
// If Years, Months and Days have a zero value, only the time components will be displayed.
// Example:
//		13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (tDur *TimeDurationDto) GetYearMthDaysTimeStr() string {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetYearMthDaysTimeStr() "
	
	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	if t2Dur.CalcType != TDurCalcType(0).StdYearMth() {

		err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).StdYearMth())

		if err != nil {
			return fmt.Sprintf(ePrefix +
				"\nError returned by t2Dur.ReCalcTimeDurationAllocation(" +
				"TDurCalcType(0).StdYearMth()).\n"+
				"Error='%v'\n", err.Error())
		}
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.DateDays > 0 || str != "" {
		str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)
	}

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetYearsMthsWeeksTimeAbbrvStr - Abbreviated formatting of Years, Months,
// Weeks, WeekDays, Hours, Minutes, Seconds, Milliseconds, Microseconds,
// Nanoseconds.
//
// At a minimum only Hours, Minutes, Seconds, Milliseconds, Microseconds
// Nanoseconds are displayed. Example return when Years, Months, Weeks
// and WeekDays are zero:
//
// 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYearsMthsWeeksTimeAbbrvStr() string {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetYearsMthsWeeksTimeAbbrvStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	if t2Dur.CalcType != TDurCalcType(0).StdYearMth() {
		err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).StdYearMth())

		if err != nil {
			return fmt.Sprintf(ePrefix +
				"Error returned by t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).StdYearMth()). "+
				"Error='%v' ", err.Error())
		}

	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.Weeks > 0 || str != "" {
		str += fmt.Sprintf("%v-Weeks ", t2Dur.Weeks)
	}

	if t2Dur.WeekDays > 0 || str != "" {
		str += fmt.Sprintf("%v-WeekDays ", t2Dur.WeekDays)
	}

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetYearsMthsWeeksTimeStr - Example Return:
// 12-Years 3-Months 2-Weeks 1-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
// At a minimum only Weeks, WeekDays, Hours, Minutes, Seconds,
// Milliseconds, Microseconds and Nanoseconds are displayed.
//
// Example return when Years, and Months are zero:
//
// 3-Weeks 2-WeekDays 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYearsMthsWeeksTimeStr() string {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetYearsMthsWeeksTimeStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	if t2Dur.CalcType != TDurCalcType(0).StdYearMth() {
		err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).StdYearMth())

		if err != nil {
			return fmt.Sprintf(ePrefix +
				"\nError returned by t2Dur.ReCalcTimeDurationAllocation(" +
				"TDurCalcType(0).StdYearMth()).\n"+
				"Error='%v'\n", err.Error())
		}
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	str += fmt.Sprintf("%v-Weeks ", t2Dur.Weeks)

	str += fmt.Sprintf("%v-WeekDays ", t2Dur.WeekDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetCumDaysCalcDto - Returns a new TimeDurationDto which re-calculates
// the values of the current TimeDurationDto and stores them in a
// 'cumulative days' format. This format always shows zero years and
// zero months. It consolidates years, months and days and presents them
// as cumulative days.
func (tDur *TimeDurationDto) GetCumDaysCalcDto() (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto) GetCumDaysCalcDto() "

	if int64(tDur.TimeDuration) == 0 {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix + "Error: Time Duration is equal to zero!")
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumDays())

	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix +
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).CumDays())\n"+
			" Error='%v'\n", err.Error())
	}

	return t2Dur, nil
}

// GetCumDaysTimeStr - Returns duration formatted as
// days, hours, minutes, seconds, milliseconds, microseconds,
// and nanoseconds. Years, months and weeks are always excluded and
// included in cumulative 'days'.
//
// Example:
//
// 97-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetCumDaysTimeStr() (string, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()
	
	ePrefix := "TimeDurationDto) GetCumDaysTimeStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumDays())

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).CumDays())\n"+
			" Error='%v'\n", err.Error())
	}

	str := ""

	str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetCumHoursCalcDto - Returns a new TimeDurationDto. The time
// values of the current TimeDurationDto are recalculated for
// 'cumulative hours'.

// This means that years, months and days are ignored and set to
// a zero value.  Instead, years, months, days and hours are
// consolidated and stored as cumulative hours.
//
func (tDur *TimeDurationDto) GetCumHoursCalcDto() (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto) GetCumHoursCalcDto() "

	if int64(tDur.TimeDuration) == 0 {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"\nError: Time Duration is ZERO value!\n")
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumHours())

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).CumHours())\n"+
			"Error='%v'\n", err.Error())
	}

	return t2Dur, nil
}

// GetCumHoursTimeStr - Returns duration formatted as hours,
// minutes, seconds, milliseconds, microseconds, nanoseconds.
//
// Years, months and days are ignored and set to a zero value.
// Instead, years, months, days and hours are consolidated and
// presented as cumulative hours.
//
// Example: 152-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (tDur *TimeDurationDto) GetCumHoursTimeStr() (string, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumHoursTimeStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumHours())

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).CumHours())\n"+
			"Error='%v'\n", err.Error())
	}

	str := ""

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetCumMinutesStr - Returns a new TimeDurationDto calculated and configured
// for cumulative minutes. This means that years, months, days and hours are
// set to a zero value.
//
// Instead, years, months, days, hours and minutes are all consolidated
// and presented as minutes.
//
// Example:
//	"527-Minutes 37-Seconds 18-Milliseconds 256-Microseconds 852-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumMinutesCalcDto() (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumMinutesCalcDto() "

	if int64(tDur.TimeDuration) == 0 {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix + 
				"\nError: Time Duration is ZERO!\n")
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumMinutes())

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).CumMinutes())\n"+
			" Error='%v'\n", err.Error())
	}

	return t2Dur, nil
}

// GetCumMinutesStr - Returns duration formatted as cumulative
// minutes. This format ignores years, months, days and hours.
//
// Instead, years, months, days, hours and minutes are all consolidated
// and presented as minutes.
//
// Example:
//	"527-Minutes 37-Seconds 18-Milliseconds 256-Microseconds 852-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumMinutesStr() (string, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumMinutesStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumMinutes())

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(TDurCalcType(0).CumMinutes())\n"+
			" Error='%v'\n", err.Error())
	}

	str := ""

	str += fmt.Sprintf("%v-Minutes ", tDur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", tDur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetCumMonthsCalcDto - Returns a new TimeDurationDto calculated
// for 'cumulative months'.
//
// The time values of the current TimeDurationDto are re-calculated and
// returned in the new TimeDurationDTo as 'cumulative months'.
// This means that Years are ignored and assigned a zero value. Instead,
// Years and Months are consolidated and presented as 'cumulative months'.
//
func (tDur *TimeDurationDto) GetCumMonthsCalcDto() (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumMonthsCalcDto() "

	if int64(tDur.TimeDuration) == 0 {
		return TimeDurationDto{}, nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumMonths())

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(TDurCalcType(0).CumMonths())\n"+
			"Error='%v'\n", err.Error())
	}

	return t2Dur, nil

}

// GetCumMonthsDaysTimeStr - Returns Cumulative Months Display
// showing Months, Days, Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds.
//
// Years are ignored and assigned a zero value. Instead, years and
// months are consolidated and presented as cumulative months.
//
func (tDur *TimeDurationDto) GetCumMonthsDaysTimeStr() (string, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumMonthsDaysTimeStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumMonths())

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).CumMonths())\n"+
			"Error='%v'\n", err.Error())
	}

	str := ""

	str += fmt.Sprintf("%v-Months ", t2Dur.Months)

	str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil

}

// GetCumSecondsCalcDto - Returns a new TimeDurationDto calculated
// for 'cumulative seconds'.
//
// The time values of the current TimeDurationDto are re-calculated and
// returned in the new TimeDurationDTo as 'cumulative seconds'.
// This means that Years, months, weeks, week days, date days, hours,
// and minutes are ignored and assigned a zero value. Instead,
// time duration is consolidated and presented as 'cumulative seconds'
// including seconds, milliseconds, microseconds and nanoseconds.
//
func (tDur *TimeDurationDto) GetCumSecondsCalcDto() (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumSecondsCalcDto() "

	if int64(tDur.TimeDuration) == 0 {
		return TimeDurationDto{}, nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumSeconds())

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).CumSeconds())\n"+
			"Error='%v'\n", err.Error())
	}

	return t2Dur, nil
}

// GetCumSecondsTimeStr - Returns a formatted time string presenting
// time duration as cumulative seconds. The display shows Seconds,
// Milliseconds, Microseconds and Nanoseconds.
func (tDur *TimeDurationDto) GetCumSecondsTimeStr() (string, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumSecondsTimeStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumSeconds())

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).CumSeconds())\n"+
			"Error='%v'\n", err.Error())
	}

	str := ""

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetCumNanosecondsDurationStr - Returns duration formatted as
// Nanoseconds. DisplayStr shows Nanoseconds expressed as a
// 64-bit integer value.
func (tDur *TimeDurationDto) GetCumNanosecondsDurationStr() string {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	str := fmt.Sprintf("%v-Nanoseconds", int64(tDur.TimeDuration))

	return str

}

// GetCumWeeksCalcDto - Returns a new TimeDurationDto re-calculated for 'Cumulative Weeks'.
// The time values of the current TimeDurationDto are converted to cumulative weeks and
// stored in the returned TimeDurationDto.
//
// 'Cumulative Weeks' means that Years and Months are ignored and assigned zero values.
// Instead, Years, Months and Weeks are consolidated and stored as cumulative Weeks,
// WeekDays, Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds.
//
func (tDur *TimeDurationDto) GetCumWeeksCalcDto() (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumWeeksCalcDto() "

	if int64(tDur.TimeDuration) == 0 {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix + "Error: Time Duration is Zero")
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumWeeks())

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).CumWeeks())\n"+
			" Error='%v'\n", err.Error())
	}

	return t2Dur, nil
}

// GetCumWeeksDaysTimeStr - Returns time duration expressed as Weeks, WeekDays, Hours,
// Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds. Years, Months and
// Days are ignored and assigned a zero value.
//
// Instead, Years, Months, Days and Hours are consolidated and presented as cumulative
// Hours.
//
// Example DisplayStr
// 126-Weeks 1-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetCumWeeksDaysTimeStr() (string, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumWeeksDaysTimeStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumWeeks())

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).CumWeeks())\n"+
			" Error='%v'\n", err.Error())
	}

	str := ""

	str += fmt.Sprintf("%v-Weeks ", t2Dur.Weeks)

	str += fmt.Sprintf("%v-WeekDays ", t2Dur.WeekDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetYrMthWkDayHrMinSecNanosecsStr - Returns duration formatted
// as Year, Month, Day, Hour, Second and Nanoseconds.
// Example: 3-Years 2-Months 3-Weeks 2-WeekDays 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds
func (tDur *TimeDurationDto) GetYrMthWkDayHrMinSecNanosecsStr() string {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetYrMthWkDayHrMinSecNanosecsStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	if t2Dur.CalcType != TDurCalcType(0).StdYearMth() {
		err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).StdYearMth())

		if err != nil {
			return fmt.Sprintf(ePrefix +
				"\nError returned by t2Dur." +
				"ReCalcTimeDurationAllocation(" +
				"TDurCalcType(0).StdYearMth()).\n"+
				"Error='%v'\n", err.Error())
		}
	}

	str := ""

	str += fmt.Sprintf("%v-Years ", t2Dur.Years)

	str += fmt.Sprintf("%v-Months ", t2Dur.Months)

	str += fmt.Sprintf("%v-Weeks ", t2Dur.Weeks)

	str += fmt.Sprintf("%v-WeekDays ", t2Dur.WeekDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.TotSubSecNanoseconds)

	return str
}

// GetDefaultDurationStr - Returns duration formatted
// as nanoseconds. The DisplayStr shows the default
// string value for duration.
// Example: 61h26m46.864197832s
func (tDur *TimeDurationDto) GetDefaultDurationStr() string {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return fmt.Sprintf("%v", tDur.TimeDuration)
}

// GetGregorianYearCalcDto - Returns a new TimeDurationDto in which years are
// calculated as 'Gregorian Years'.
//
// Unlike other calculations which use a Standard 365-day year consisting of
// 24-hour days which take into account correct leap years, a Gregorian Year
// consists of 365 days, 5-hours, 59-minutes and 12 Seconds. For the Gregorian
// calendar the average length of the calendar year (the mean year) across
// the complete leap cycle of 400 Years is 365.2425 days.
//
// Sources:
// https://en.wikipedia.org/wiki/Year
// Source: https://en.wikipedia.org/wiki/Gregorian_calendar
//
func (tDur *TimeDurationDto) GetGregorianYearCalcDto() (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetGregorianYearCalcDto() "

	if int64(tDur.TimeDuration) == 0 {
		return TimeDurationDto{}, nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).GregorianYears())

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).GregorianYears())\n"+
			"Error='%v'\n", err.Error())
	}

	return t2Dur, nil
}

// GetGregorianYearDurationStr - Returns a string showing the
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
func (tDur *TimeDurationDto) GetGregorianYearDurationStr() (string, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetGregorianYearDurationStr() "

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := t2Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).GregorianYears())

	if err != nil {
		return "", fmt.Errorf(ePrefix+
			"\nError returned by ReCalcTimeDurationAllocation(" +
			"TDurCalcType(0).GregorianYears())\n"+
			"Error='%v'\n", err.Error())
	}

	str := fmt.Sprintf("%v-Gregorian Years ", t2Dur.Years)

	str += fmt.Sprintf("%v-Months ", t2Dur.Months)

	str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// New - Creates and returns a new TimeDurationDto based on starting
// and ending date times.  Because, time zone location is crucial to
// completely accurate duration calculations, the time zone of the
// starting date time, 'startDateTime' is applied to parameter,
// 'endDateTime' before making the duration calculation.
//
// Note:  This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//        This means that duration is allocated over years, months, weeks, weekdays, date days,
//        hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//        For details, see Type 'TDurCalcType' in this source file: 
//          MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// Input Parameters:
// =================
//
// startDateTime time.Time
//            - Starting date time
//
// endDateTime  time.Time
//            - Ending date time
//
// dateTimeFmtStr string  
//             - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//               If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.New(
//                                   startTime,
//                                   endTime,
//                                   FmtDateTimeYrMDayFmtStr)
//
//  Note: FmtDateTimeYrMDayFmtStr' is a constant available in source file:
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) New(
	startDateTime,
	endDateTime time.Time,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.New() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimesCalcTz(
		&tDur2,
		startDateTime,
		endDateTime,
		TDurCalc.StdYearMth(),
		startDateTime.Location().String(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewAutoEnd - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'timeZoneLocation' input parameters.
//
// The input parameter 'startDateTime' is first converted to the Time Zone specified
// by the input parameter, 'timeZoneLocation'.
//
// The ending date time is calculated automatically by assigning the date time returned
// by calling time.Now(). The ending date time that is assigned by time.Now() is converted
// to the time zone specified by input parameter, 'timeZoneLocation'.
//
// If the calculated ending date time is prior to 'startDateTime', the values are
// reversed and ending date time is assigned to 'startDateTime' while 'startDateTime'
// is assigned to ending date time.
//
// Note:   This method applies the standard Time Duration allocation, calculation type
//         'TDurCalcType(0).StdYearMth()'. This means that time duration is allocated over years,
//         months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
//         microseconds and nanoseconds.  For details, see Type 'TDurCalcType' in this source
//         file: MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// Input Parameters:
// =================
//
// startDateTime  time.Time
//  - Starting date time
//
// timeZoneLocation  string
//  - Designates the standard Time Zone location by which
//    time duration will be compared. This ensures that
//    'oranges are compared to oranges and apples are compared
//    to apples' with respect to start time and end time duration
//    calculations.
//
//    If 'timeZoneLocation' is passed as an empty string, it
//    will be automatically defaulted to the 'UTC' time zone.
//    Reference Universal Coordinated Time: 
//       https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//    Time zone location must be designated as one of three
//    types of time zones.
//
//   (1) The time zone "Local", which Golang accepts as
//       the time zone currently configured on the host
//       computer.
//
//   (2) IANA Time Zone - A valid IANA Time Zone from the
//       IANA database.
//       See https://golang.org/pkg/time/#LoadLocation
//       and https://www.iana.org/time-zones to ensure that
//       the IANA Time Zone Database is properly configured
//       on your system.
//
//       IANA Time Zone Examples:
//         "America/New_York"
//         "America/Chicago"
//         "America/Denver"
//         "America/Los_Angeles"
//         "Pacific/Honolulu"
//         "Etc/UTC" = GMT or UTC
//
//   (3) A Military Time Zone
//       Reference:
//         https://en.wikipedia.org/wiki/List_of_military_time_zones
//         http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//         https://www.timeanddate.com/time/zones/military
//         https://www.timeanddate.com/worldclock/timezone/alpha
//         https://www.timeanddate.com/time/map/
//
//        Examples:
//          "Alpha"   or "A"
//          "Bravo"   or "B"
//          "Charlie" or "C"
//          "Delta"   or "D"
//          "Zulu"    or "Z"
//
//          If the time zone "Zulu" is passed to this method, it will be
//          classified as a Military Time Zone.
//
// dateTimeFmtStr string    - A date time format string which will be used
//                            to format and display 'dateTime'. Example:
//                            "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                            If 'dateTimeFmtStr' is submitted as an
//                            'empty string', a default date time format
//                            string will be applied. The default date time
//                            format string is:
//                            FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewAutoEnd(
//                                     startTime,
//                                     TZones.US.Central(),
//                                     FmtDateTimeYrMDayFmtStr)
//
//   Note:
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewAutoEnd(
	startDateTime time.Time,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewAutoEnd() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimesCalcTz(
		&tDur2,
		startDateTime,
		time.Now().UTC(),
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewAutoStart - Creates and returns a new TimeDurationDto instance. Starting date time is
// automatically initialized by calling time.Now(). Afterwards, start date time is converted
// to the Time Zone specified in input parameter, 'timeZoneLocation'.
//
// This method will set ending date time to the same value as starting date time resulting in
// a time duration value of zero.
//
// In order to compute the final time duration value, call the method TimeDurationDto.SetAutoEnd()
// when ready.  At that point, ending date time will be set by a call to time.Now().
//
// Use of these two methods, 'NewAutStart' and 'SetAutoEnd', constitutes a stop watch feature which
// can be triggered to measure elapsed time.
//
// Note:  This method applies the standard Time Duration allocation, calculation type
//        'TDurCalcType(0).StdYearMth()'. This means that time duration is allocated over years,
//        months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
//        microseconds and nanoseconds. For details, see Type 'TDurCalcType' in this source
//        file: MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// Input Parameters:
// =================
//
// timeZoneLocation string
//        - Designates the standard Time Zone location by which
//          time duration will be compared. This ensures that
//          'oranges are compared to oranges and apples are compared
//          to apples' with respect to start time and end time duration
//          calculations.
//
//          If 'timeZoneLocation' is passed as an empty string, it
//          will be automatically defaulted to the 'UTC' time zone.
//          Reference Universal Coordinated Time:
//             https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//           Time zone location must be designated as one of three types
//           of time zones.
//
//           (1) The time zone "Local", which Golang accepts as
//               the time zone currently configured on the host
//               computer.
//
//           (2) IANA Time Zone - A valid IANA Time Zone from the
//               IANA database.
//               See https://golang.org/pkg/time/#LoadLocation
//               and https://www.iana.org/time-zones to ensure that
//               the IANA Time Zone Database is properly configured
//               on your system.
//
//               IANA Time Zone Examples:
//                 "America/New_York"
//                 "America/Chicago"
//                 "America/Denver"
//                 "America/Los_Angeles"
//                 "Pacific/Honolulu"
//                 "Etc/UTC" = GMT or UTC
//
//            (3) A Military Time Zone
//                Reference:
//                 https://en.wikipedia.org/wiki/List_of_military_time_zones
//                 http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                 https://www.timeanddate.com/time/zones/military
//                 https://www.timeanddate.com/worldclock/timezone/alpha
//                 https://www.timeanddate.com/time/map/
//
//                Examples:
//                  "Alpha"   or "A"
//                  "Bravo"   or "B"
//                  "Charlie" or "C"
//                  "Delta"   or "D"
//                  "Zulu"    or "Z"
//
//                  If the time zone "Zulu" is passed to this method, it will be
//                  classified as a Military Time Zone.
//
// dateTimeFmtStr string
//             - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//               If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewAutoStart(
//                   TZones.US.Central(),
//                   FmtDateTimeYrMDayFmtStr)
//
//   Note:
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewAutoStart(
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewAutoStart() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	startEndTime := time.Now().UTC()

	err := tDurDtoUtil.setStartEndTimesCalcTz(
		&tDur2,
		startEndTime,
		startEndTime,
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndDateTzDto - Creates and returns a new TimeDurationDto populated with
// time duration data. The input parameters for starting and ending date time are
// submitted as type 'DateTzDto'.
//
// The Time Zone Location is extracted from input parameter, 'startDateTz'. The extracted
// time zone is used to configure the ending date time thereby providing a common basis for
// subsequent time duration calculations.
//
// This method automatically applies the standard Time Duration allocation, calculation type
// 'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years, months, weeks,
//  weekdays, date days, hours, minutes, seconds, milliseconds,	microseconds and nanoseconds.
// For details, see Type 'TDurCalcType' in this source file:
//        MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// Input Parameters
// ================
//
// startDateTz DateTzDto
//            - Contains the starting date time used in time duration calculations
//
// endDateTz   DateTzDto
//            - Contains the ending date time used in time duration calculations.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//               FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndDateTzDto(
//                                   startDateTz,
//                                   endDateTz,
//                                   FmtDateTimeYrMDayFmtStr)
//
//  Note: 'FmtDateTimeYrMDayFmtStr' is a constant defined in source
//        file, "constantsdatetime.go".
//
func (tDur TimeDurationDto) NewStartEndDateTzDto(
	startDateTz,
	endDateTz DateTzDto,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndDateTzDto() "

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.isValidDateTzDto(
		&startDateTz,
		ePrefix)

	if err != nil {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "startDateTz",
				inputParameterValue: "",
				errMsg: fmt.Sprintf(
					"Input Parameter 'startDateTz' is INVALID!\n" +
						"Validation Error='%v'", err.Error()),
				err:                 nil,
			}
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	timeZoneLocation :=
		startDateTz.GetBestConvertibleTimeZone().GetMilitaryOrStdTimeZoneName()

	err = tDurDtoUtil.setStartEndTimesDateDtoCalcTz(
		&tDur2,
		startDateTz,
		endDateTz,
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix,
		)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndDateTzDtoCalcTz - Creates and returns a new TimeDurationDto populated with
// time duration data. The input parameters for starting and ending date time are
// submitted as type 'DateTzDto'.
//
// The user is required to specify a Time Zone Location for use in converting date
// times to a common frame of reference used in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. For details see
// Type 'TDurCalcType' which is located in source file:
//       MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// Input Parameters:
// =================
//
// startDateTz     DateTzDto
//     - Contains the starting date time used in time duration calculations
//
// endDateTz       DateTzDto
//     - Contains the ending date time used in time duration calculations.
//
// tDurCalcType TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration:
//
//       TDurCalcType(0).StdYearMth() - Default - standard year, month week,
//                                      day time calculation.
//
//       TDurCalcType(0).CumMonths()  - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()   - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()    - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()   - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes() - Computes cumulative minutes. No Years, months, weeks, days
//                                      or hours.
//
//       TDurCalcType(0).CumSeconds() - Computes cumulative seconds. No Years, months, weeks, days,
//                                      hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//                    - Computes Years based on average length of a Gregorian Year
//                      Used for very large duration values.
//
//              Type 'TDurCalcType' is located in source file:
//                 MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string 
//     - Designates the standard Time Zone location Name by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end times used in
//       time duration calculations.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
// dateTimeFmtStr string  
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndDateTzDtoCalcTz(
//                                    startDateTz,
//                                    endDateTz,
//                                    TDurCalcType(0).StdYearMth(),
//                                    TZones.US.Central(),
//                                    FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartEndDateTzDtoCalcTz(
	startDateTz,
	endDateTz DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()
	
	defer tDur.lock.Unlock()
		
	ePrefix := "TimeDurationDto.NewStartEndDateTzDtoCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimesDateDtoCalcTz(
		&tDur2,
		startDateTz,
		endDateTz,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndDateTzDtoTz - Creates and returns a new TimeDurationDto populated with
// time duration data. The input parameters for starting and ending date time are
// submitted as type 'DateTzDto'.
//
// The user is required to specify a Time Zone Location for use in converting date
// times to a common frame of reference used in subsequent time duration calculations.
//
// This method automatically applies the standard Time Duration allocation, calculation type
// 'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years, months, weeks,
//  weekdays, date days, hours, minutes, seconds, milliseconds,	microseconds and nanoseconds.
// For details, see Type 'TDurCalcType' in this source file:
//        MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// Input Parameters:
// =================
//
// startDateTz  DateTzDto  - Contains the starting date time used in time duration calculations
//
// endDateTz    DateTzDto - Contains the ending date time used in time duration calculations.
//
// timeZoneLocation  string  
//              - Designates the standard Time Zone location by which
//                time duration will be compared. This ensures that
//                'oranges are compared to oranges and apples are compared
//                to apples' with respect to start time and end times used in
//                time duration calculations.
//
//                If 'timeZoneLocation' is passed as an empty string, it
//                will be automatically defaulted to the 'UTC' time zone.
//                Reference Universal Coordinated Time: 
//                   https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//                Time zone location must be designated as one of three types
//                of time zones.
//
//                (1) The time zone "Local", which Golang accepts as
//                    the time zone currently configured on the host
//                    computer.
//                
//                (2) IANA Time Zone - A valid IANA Time Zone from the
//                    IANA database.
//                    See https://golang.org/pkg/time/#LoadLocation
//                    and https://www.iana.org/time-zones to ensure that
//                    the IANA Time Zone Database is properly configured
//                    on your system.
//                
//                    IANA Time Zone Examples:
//                      "America/New_York"
//                      "America/Chicago"
//                      "America/Denver"
//                      "America/Los_Angeles"
//                      "Pacific/Honolulu"
//                      "Etc/UTC" = GMT or UTC
//                
//                (3) A Military Time Zone
//                    Reference:
//                      https://en.wikipedia.org/wiki/List_of_military_time_zones
//                      http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                      https://www.timeanddate.com/time/zones/military
//                      https://www.timeanddate.com/worldclock/timezone/alpha
//                      https://www.timeanddate.com/time/map/
//                
//                     Examples:
//                       "Alpha"   or "A"
//                       "Bravo"   or "B"
//                       "Charlie" or "C"
//                       "Delta"   or "D"
//                       "Zulu"    or "Z"
//                
//                       If the time zone "Zulu" is passed to this method, it will be
//                       classified as a Military Time Zone.
//
// dateTimeFmtStr string    
//              - A date time format string which will be used
//                to format and display 'dateTime'. Example:
//                "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                If 'dateTimeFmtStr' is submitted as an
//                'empty string', a default date time format
//                string will be applied. The default date time
//                format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndDateTzDtoTz(
//                       startDateTz,
//                       endDateTz,
//                       TZones.US.Central(),
//                       FmtDateTimeYrMDayFmtStr)
//
//      Note:
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartEndDateTzDtoTz(
	startDateTz,
	endDateTz DateTzDto,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndDateTzDtoTz() "

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	var err error

	err = startDateTz.IsValid()

	if err != nil {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "startDateTz",
				inputParameterValue: "",
				errMsg:   fmt.Sprintf(
					"Input Parameter 'startDateTz' is INVALID!\n" +
						"Validation Error='%v'", err.Error()),
				err:                 nil,
			}
	}

	err = endDateTz.IsValid()

	if err != nil {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "endDateTz",
				inputParameterValue: "",
				errMsg:   fmt.Sprintf(
					"Input Parameter 'endDateTz' is INVALID!\n" +
						"Validation Error='%v'", err.Error()),
				err:                 nil,
			}
	}

	var tzSpec TimeZoneSpecification

	tzSpec,
		err = TimeZoneDefinition{}.NewTzSpecFromTzName(
		startDateTz.dateTimeValue,
		timeZoneLocation,
		TzConvertType.Relative())

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix+
				"\nError: 'timeZoneLocation' input parameter is INVALID! "+
				"'timeZoneLocation'='%v'\n" +
				"Error='%v'\n",
				timeZoneLocation, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartEndTimesCalcTz(
		startDateTz.dateTimeValue,
		endDateTz.dateTimeValue,
		TDurCalcType(0).StdYearMth(),
		tzSpec.locationName,
		dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix+
				"Error returned by SetStartEndTimesCalcTz(). Error='%v' ",
				err.Error())
	}

	return t2Dur, nil
}

// NewStartEndTimesTz - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters.
// The user is required to specify a common Time Zone Location for use in converting
// date times to a common frame of reference to subsequent time duration calculations.
//
// Note:    This method applies the standard Time Duration allocation, calculation type
//             'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years,
//             months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
//             microseconds and nanoseconds.   See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime   time.Time  - Starting time
//
// endDateTime      time.Time - Ending time
//
// timeZoneLocation   string   
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time: 
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//              
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//              
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//              
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//              
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//              
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string      
//             - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//               If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesTz(
//                                    startTime,
//                                    endTime,
//                                    TZones.US.Central(),
//                                    FmtDateTimeYrMDayFmtStr)
//
//      Note:
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartEndTimesTz(
	startDateTime, 
	endDateTime time.Time,
	timeZoneLocation, 
	dateTimeFmtStr string) (TimeDurationDto, error) {
	
	tDur.lock.Lock()
	
	defer tDur.lock.Unlock()
	
	ePrefix := "TimeDurationDto.NewStartEndTimesTz() "
	
	if startDateTime.IsZero() && endDateTime.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + 
				"\nError: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!\n")
	}

	if endDateTime.Before(startDateTime) {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "endDateTime",
				inputParameterValue: "",
				errMsg:              "'endDateTime' is less than 'startDateTime'!",
				err:                 nil,
			}
	}
	
	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)
	
	var tzSpec TimeZoneSpecification
	var err error
	
	tzSpec,
		err = TimeZoneDefinition{}.NewTzSpecFromTzName(
		startDateTime,
		timeZoneLocation,
		TzConvertType.Relative())

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix+
				"\nError: 'timeZoneLocation' input parameter is INVALID! "+
				"'timeZoneLocation'='%v'\n" +
				"Error='%v'\n",
				timeZoneLocation, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartEndTimesCalcTz(
		startDateTime, 
		endDateTime, 
		TDurCalcType(0).StdYearMth(),
		tzSpec.locationName, 
		dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix+"Error returned from "+
			"SetStartEndTimesCalcTz(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)."+
			"Error='%v'", err.Error())
	}

	return t2Dur, nil
}

// NewStartEndTimesCalc - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// The user is required to submit input parameters for date time calculation type. The
// Time Zone Location is extracted from the 'startDateTime' parameter
//
// Input Parameters:
// =================
//
// startDateTime    time.Time - Starting time
//
// endDateTime      time.Time - Ending time
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//              TDurCalcType(0).StdYearMth()        - Default - standard year, month week,
//                                                                       day time calculation.
//
//              TDurCalcType(0).CumMonths()         - Computes cumulative months - no Years.
//
//              TDurCalcType(0).CumWeeks()          - Computes cumulative weeks. No Years or months
//
//              TDurCalcType(0).CumDays()           - Computes cumulative days. No Years, months or weeks.
//
//              TDurCalcType(0).CumHours()          - Computes cumulative hours. No Years, months, weeks or days.
//
//              TDurCalcType(0).CumMinutes()         - Computes cumulative minutes. No Years, months, weeks, days
//                                                     or hours.
//
//              TDurCalcType(0).CumSeconds()         - Computes cumulative seconds. No Years, months, weeks, days,
//                                                     hours or minutes.
//
//              TDurCalcType(0).GregorianYears()     - Computes Years based on average length of a Gregorian Year
//                                                     Used for very large duration values.
//
//              Type 'TDurCalcType' is located in source file:
//                  MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesCalcTz(
//                    startTime,
//                    endTime,
//                    TDurCalcType(0).StdYearMth(),
//                    FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartEndTimesCalc(
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesCalc() "

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!")
	}

	if endDateTime.Before(startDateTime) {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "endDateTime",
				inputParameterValue: "",
				errMsg:              "'endDateTime' is LESS than 'startDateTime'",
				err:                 nil,
			}
	}

	if tDurCalcType < TDurCalc.StdYearMth() ||
		tDurCalcType > TDurCalc.GregorianYears() {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "tDurCalcType",
				inputParameterValue: tDurCalcType.String(),
				errMsg:  "Input Parameter 'tDurCalcType' is INVALID!",
				err:                 nil,
			}
	}

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	var tzSpec TimeZoneSpecification

	var err error

	tzSpec,
		err = TimeZoneDefinition{}.NewTzSpecFromTzName(
		startDateTime,
		startDateTime.Location().String(),
		TzConvertType.Relative())

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix+
				"\nError: 'timeZoneLocation' input parameter is INVALID! "+
				"'timeZoneLocation'='%v'\n" +
				"Error='%v'\n",
				startDateTime.Location().String(), err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartEndTimesCalcTz(
		startDateTime, 
		endDateTime, 
		tDurCalcType, tzSpec.locationName,
		dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix+"Error returned from "+
			"SetStartEndTimesCalcTz(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)."+
			"Error='%v'", err.Error())
	}

	return t2Dur, nil
}

// NewStartEndTimesCalcTz - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters.
//
// The user is required to specify a common Time Zone Location for use in converting
// date times to a common frame of reference for use in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. For details see
// Type 'TDurCalcType' which is located in source file:
//          MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// Input Parameters:
// =================
//
// startDateTime   time.Time - Starting time
//
// endDateTime     time.Time - Ending time
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//               TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                                  day time calculation.
//
//               TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//               TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//               TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//               TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//               TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                                  or hours.
//
//               TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                                  hours or minutes.
//
//               TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                                  Used for very large duration values.
//
//                               Type 'TDurCalcType' is located in source file:
//                                    MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
// timeZoneLocation   string   
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time: 
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//              
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//              
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//              
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//              
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//              
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string      
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesCalcTz(
//                    startTime,
//                    endTime,
//                    TDurCalcType(0).StdYearMth(),
//                    TZones.US.Central(),
//                    FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartEndTimesCalcTz(
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesCalcTz() "

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + 
				"\nError: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!\n")
	}

	if endDateTime.Before(startDateTime) {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "endDateTime",
				inputParameterValue: "",
				errMsg:              "'endDateTime' is LESS than 'startDateTime'!",
				err:                 nil,
			}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	var err error

	if tDurCalcType < TDurCalc.StdYearMth() ||
		tDurCalcType > TDurCalc.GregorianYears() {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "tDurCalcType",
				inputParameterValue: tDurCalcType.String(),
				errMsg:  "Input Parameter 'tDurCalcType' is INVALID!",
				err:                 nil,
			}
	}

	var tzSpec TimeZoneSpecification
	
	tzSpec,
		err = TimeZoneDefinition{}.NewTzSpecFromTzName(
		startDateTime,
		timeZoneLocation,
		TzConvertType.Relative())

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix+
				"\nError: 'timeZoneLocation' input parameter is INVALID! "+
				"'timeZoneLocation'='%v'\n" +
				"Error='%v'\n",
				timeZoneLocation, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartEndTimesCalcTz(
		startDateTime, 
		endDateTime, 
		tDurCalcType, 
		tzSpec.locationName,
		dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix+"Error returned from "+
			"SetStartEndTimesCalcTz(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)."+
			"Error='%v'", err.Error())
	}

	return t2Dur, nil
}

// NewStartEndTimesDateDto - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters. The
// 'startDateTime' and 'endDateTime' parameters are of type DateTzDto.
//
// Time Zone Location is derived from input parameter 'startDateTime' and provides a
// common frame of reference for use in subsequent time duration calculations.
//
// Note: This method applies the standard Time Duration allocation, calculation type
//       'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years,
//       months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
//       microseconds and nanoseconds.   See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime    DateTzDto - Starting date time
//
// endDateTime      DateTzDto - Ending date time
//
// dateTimeFmtStr string      
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesDateDto(
//                                    startTime, 
//                                    endTime, 
//                                    FmtDateTimeYrMDayFmtStr)
//
//      Note:
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartEndTimesDateDto(
	startDateTime,
	endDateTime DateTzDto,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesDateDto() "
	
	var err error
	
	err = startDateTime.IsValid()

	if err != nil {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "startDateTime",
				inputParameterValue: "",
				errMsg:   fmt.Sprintf(
					"Input Parameter 'startDateTime' is INVALID!\n" +
						"Validation Error='%v'", err.Error()),
				err:                 nil,
			}
	}

	err = endDateTime.IsValid()

	if err != nil {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "endDateTz",
				inputParameterValue: "",
				errMsg:   fmt.Sprintf(
					"Input Parameter 'endDateTz' is INVALID!\n" +
						"Validation Error='%v'", err.Error()),
				err:                 nil,
			}
	}

	if startDateTime.dateTimeValue.IsZero() &&
		  endDateTime.dateTimeValue.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + 
				"\nError: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!\n")
	}

	if endDateTime.dateTimeValue.Before(startDateTime.dateTimeValue) {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "endDateTime",
				inputParameterValue: "",
				errMsg:              "'endDateTime' is LESS than 'startDateTime'!",
				err:                 nil,
			}
	}
	
	tzSpec := startDateTime.GetBestConvertibleTimeZone()

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartEndTimesDateDtoCalcTz(
		startDateTime,
		endDateTime,
		TDurCalcType(0).StdYearMth(),
		tzSpec.locationName,
		dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix+
			"\nError returned from "+
			"SetStartEndTimesDateDtoCalcTz(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr).\n"+
			"Error='%v'\n", err.Error())
	}

	return t2Dur, nil
}

// NewStartEndTimesDateDtoCalc - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters. The
// 'startDateTime' and 'endDateTime' parameters are of type DateTzDto.
//
// Time Zone Location is derived from input parameter 'startDateTime'. If the 'endDateTime'
// time zone is NOT equivalent to 'startDateTime', the 'endDateTime' will be converted to
// the time zone provided by the 'startDateTime' parameter. This will provide a common basis
// for use in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime   DateTzDto   - Starting date time
//
// endDateTime     DateTzDto - Ending date time
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//               TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                                  day time calculation.
//
//               TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//               TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//               TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//               TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//               TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                                  or hours.
//
//               TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                                  hours or minutes.
//
//               TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                                  Used for very large duration values.
//
//               Type 'TDurCalcType' is located in source file:
//                  MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesDateDtoCalc(
//                            startTime,
//                            endTime,
//                            TDurCalcType(0).StdYearMth(),
//                            FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartEndTimesDateDtoCalc(
	startDateTime,
	endDateTime DateTzDto,
	tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesDateDtoCalc() "

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	var err error

	err = startDateTime.IsValid()

	if err != nil {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "startDateTime",
				inputParameterValue: "",
				errMsg:   fmt.Sprintf(
					"Input Parameter 'startDateTime' is INVALID!\n" +
						"Validation Error='%v'", err.Error()),
				err:                 nil,
			}
	}

	err = endDateTime.IsValid()

	if err != nil {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "endDateTz",
				inputParameterValue: "",
				errMsg:   fmt.Sprintf(
					"Input Parameter 'endDateTz' is INVALID!\n" +
						"Validation Error='%v'", err.Error()),
				err:                 nil,
			}
	}

	if startDateTime.dateTimeValue.IsZero() &&
			endDateTime.dateTimeValue.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!")
	}

	if endDateTime.dateTimeValue.Before(startDateTime.dateTimeValue) {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "endDateTime",
				inputParameterValue: "",
				errMsg:              "'endDateTime' is LESS than 'startDateTime'!",
				err:                 nil,
			}
	}

	if tDurCalcType < TDurCalc.StdYearMth() ||
		tDurCalcType > TDurCalc.GregorianYears() {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "tDurCalcType",
				inputParameterValue: tDurCalcType.String(),
				errMsg:  "Input Parameter 'tDurCalcType' is INVALID!",
				err:                 nil,
			}
	}

	t2Dur := TimeDurationDto{}

	tzSpec := startDateTime.GetBestConvertibleTimeZone()

	err = t2Dur.SetStartEndTimesDateDtoCalcTz(
		startDateTime,
		endDateTime,
		tDurCalcType,
		tzSpec.locationName,
		dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix+
			"\nError returned from "+
			"SetStartEndTimesDateDtoCalcTz(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr).\n"+
			"Error='%v'\n", err.Error())
	}

	return t2Dur, nil
}

// NewStartEndTimesDateTzDtoCalcTz - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters. The
// 'startDateTime' and 'endDateTime' parameters are of type DateTzDto.
//
// The user is required to specify a specific Time Zone Location used to convert date times
// to a common frame of reference for use in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime DateTzDto - Starting date time
//
// endDateTime   DateTzDto - Ending date time
//
// tDurCalcType TDurCalcType 
//   - Specifies the calculation type to be used in allocating
//     time duration:
//
//     TDurCalcType(0).StdYearMth()   - Default - standard year, month week,
//                    day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//     Type 'TDurCalcType' is located in source file:
//       MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
// timeZoneLocation string 
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//               calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time: 
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//              
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//              
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//              
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//              
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//              
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string  
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesDateTzDtoCalc(
//                                   startTime,
//                                   endTime,
//                                   TZones.US.Central(),
//                                   TDurCalcType(0).StdYearMth(),
//                                   FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartEndTimesDateTzDtoCalcTz(
	startDateTz,
	endDateTz DateTzDto, 
	tDurCalcType TDurCalcType, 
	timeZoneLocation string,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()
	
	defer tDur.lock.Unlock()
	
	ePrefix := "TimeDurationDto.NewStartEndTimesDateTzDtoCalcTz() "

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	var err error

	err = startDateTz.IsValid()

	if err != nil {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "startDateTz",
				inputParameterValue: "",
				errMsg:   fmt.Sprintf(
					"Input Parameter 'startDateTz' is INVALID!\n" +
						"Validation Error='%v'", err.Error()),
				err:                 nil,
			}
	}

	err = endDateTz.IsValid()

	if err != nil {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "endDateTz",
				inputParameterValue: "",
				errMsg:   fmt.Sprintf(
					"Input Parameter 'endDateTz' is INVALID!\n" +
						"Validation Error='%v'", err.Error()),
				err:                 nil,
			}
	}

	if startDateTz.dateTimeValue.IsZero() &&
		endDateTz.dateTimeValue.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix +
				"\nError: Both 'startDateTz' and 'endDateTz' " +
				"input parameters are ZERO!\n")
	}

	if endDateTz.dateTimeValue.Before(startDateTz.dateTimeValue) {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "endDateTz",
				inputParameterValue: "",
				errMsg:              "'endDateTz' is LESS than 'startDateTz'!",
				err:                 nil,
			}
	}

	if tDurCalcType < TDurCalc.StdYearMth() ||
		tDurCalcType > TDurCalc.GregorianYears() {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "tDurCalcType",
				inputParameterValue: tDurCalcType.String(),
				errMsg:  "Input Parameter 'tDurCalcType' is INVALID!",
				err:                 nil,
			}
	}

	var tzSpec TimeZoneSpecification

	tzSpec,
		err = TimeZoneDefinition{}.NewTzSpecFromTzName(
		startDateTz.dateTimeValue,
		timeZoneLocation,
		TzConvertType.Relative())

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix+
				"\nError: 'timeZoneLocation' input parameter is INVALID! "+
				"'timeZoneLocation'='%v'\n" +
				"Error='%v'\n",
				timeZoneLocation, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartEndTimesDateDtoCalcTz(startDateTz,
		endDateTz,
		tDurCalcType,
		tzSpec.locationName,
		dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{}, 
		fmt.Errorf(ePrefix+"Error returned from "+
			"SetStartEndTimesDateDtoCalcTz(startDateTz, endDateTz, timeZoneLocation, dateTimeFmtStr).\n"+
			"Error='%v'\n", err.Error())
	}

	return t2Dur, nil
}

// NewStartTimeDurationTz - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTime' and time duration. 'startDateTime' is used to derive Time Zone Location.
// The time duration value is added to 'startDateTime' in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time. The
// actual starting date time is computed by subtracting duration from ending date time.
//
// Note:  This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//        This means that duration is allocated over years, months, weeks, weekdays, date days,
//        hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//        See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime time.Time
//            - Starting date time for the duration calculation.
//              Note: The Time Zone extracted from 'startDateTime'
//              is used in calculating both starting date time and
//              ending date time.
//
// duration  time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationTz(
//                                     startTime,
//                                     duration,
//                                     FmtDateTimeYrMDayFmtStr)
//      Note:
//       'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//       'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartTimeDuration(
	startDateTime time.Time,
	duration time.Duration,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDuration() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimeDurationCalcTz(
		&tDur2,
		startDateTime,
		duration,
		TDurCalc.StdYearMth(),
		startDateTime.Location().String(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDurationTz - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTime', time duration and 'timeZoneLocation'. 'startDateTime' is converted to the
// specified 'timeZoneLocation'. The duration value is added to 'startDateTime' in order to
// compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and the
//  actual starting date time is computed by subtracting duration.
//
// Note: This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//       This means that duration is allocated over years, months, weeks, weekdays, date days,
//       hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//       See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime time.Time
//            - Starting date time for the duration calculation
//
// duration  time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationTz(
//                                    startTime,
//                                    duration,
//                                    TZones.US.Central(),
//                                    FmtDateTimeYrMDayFmtStr)
//
//     Note:
//       'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'
//
//       'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartTimeDurationTz(
	startDateTime time.Time,
	duration time.Duration,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimeDurationCalcTz(
		&tDur2,
		startDateTime,
		duration,
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDurationCalcTz - Creates and returns a new TimeDurationDto based on input
// parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation type.
// 'startDateTime' is converted to the specified 'timeZoneLocation' and the duration value
// is added to it in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and
// the	actual starting date time is computed by subtracting duration.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime time.Time
//            - Starting date time for the duration calculation
//
// duration  time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//              Type 'TDurCalcType' is located in source file:
//                MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationCalcTz(
//                                   startTime,
//                                   duration,
//                                   TZones.US.Central(),
//                                   TDurCalcType(0).StdYearMth(),
//                                   FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//       'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'
//
//       'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartTimeDurationCalcTz(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationCalcTz() "

	if startDateTime.IsZero() && duration == 0 {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'duration' " +
				"input parameters are ZERO!")
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	var err error

	if tDurCalcType < TDurCalc.StdYearMth() ||
		tDurCalcType > TDurCalc.GregorianYears() {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "tDurCalcType",
				inputParameterValue: tDurCalcType.String(),
				errMsg:  "Input Parameter 'tDurCalcType' is INVALID!",
				err:                 nil,
			}
	}

	var tzSpec TimeZoneSpecification

	tzSpec,
		err = TimeZoneDefinition{}.NewTzSpecFromTzName(
		startDateTime,
		timeZoneLocation,
		TzConvertType.Relative())

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix+
				"\nError: 'timeZoneLocation' input parameter is INVALID! "+
				"'timeZoneLocation'='%v'\n" +
				"Error='%v'\n",
				timeZoneLocation, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartTimeDurationCalcTz(
		startDateTime,
		duration,
		tDurCalcType,
		tzSpec.locationName,
		dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix+
				"Error returned by t2Dur.SetStartTimeDurationCalcTz(...) Error='%v'",
				err.Error())
	}

	return t2Dur, nil
}

// NewStartTimeDurationCalc - Creates and returns a new TimeDurationDto based on input
// parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation type.
//
// The duration value is added to 'startDateTime' in order to compute the ending date time.
//
// If 'duration' is a negative value, 'startDateTime' is converted to ending date time and
// the actual starting date time is computed by subtracting duration.
//
// The time zone location applied to both 'startDateTime' and ending date time is derived
// from input parameter, 'startDateTime'.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime   time.Time
//            - Starting date time for the duration calculation
//
// duration    time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                 FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationCalc(
//                                   startTime,
//                                   duration,
//                                   TDurCalc.StdYearMth(),
//                                   FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//        standard year month day time duration allocation.
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartTimeDurationCalc(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationCalc() "

	timeZoneLocation := startDateTime.Location().String()

	t2Dur := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartTimeDurationCalcTz(
		&t2Dur,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, nil
	}

	return t2Dur, nil
}

// NewStartTimeDurationDateDto - Creates and returns a new TimeDurationDto based on input
// parameters 'startDateTime' and time duration. 'startDateTime' is of type DateTzDto.
//
// The time duration value is added to 'startDateTime' in order to compute the ending date time.
// If 'duration' is a negative value, 'startDateTime' is converted to ending date time. The
// actual starting date time is computed by subtracting duration from ending date time.
//
// Time Zone location is derived from 'startDateTime'.
//
// Note: This method applies the standard Time Duration allocation calculation type,
//       'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years,
//       months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
//       microseconds and nanoseconds. See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime DateTzDto
//            - Starting date time for the duration calculation
//
// duration  time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDto(
//                   startTime,
//                   duration,
//                   FmtDateTimeYrMDayFmtStr)
//
//  Note: 'FmtDateTimeYrMDayFmtStr' is a constant value available in source
//        file, constantsdatetime.go
//
func (tDur TimeDurationDto) NewStartTimeDurationDateDto(
	startDateTimeTz DateTzDto,
	duration time.Duration,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationTz() "

	timeZone := startDateTimeTz.GetBestConvertibleTimeZone()

	var timeZoneLocation string

	if timeZone.timeZoneType == TzType.Military() {
		timeZoneLocation = timeZone.militaryTimeZoneName
	} else {
		timeZoneLocation = timeZone.locationName
	}

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartTimeDurationDateDtoCalcTz(
						&tDur2,
						startDateTimeTz,
						duration,
						TDurCalc.StdYearMth(),
						timeZoneLocation,
						dateTimeFmtStr,
						ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDurationDateDtoTz - Creates and returns a new TimeDurationDto based on input
// parameters 'startDateTime', time duration and 'timeZoneLocation'. 'startDateTime' is
// converted to the specified 'timeZoneLocation'. The 'duration' value is added to 'startDateTime'
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and the
// actual starting date time is computed by subtracting duration.
//
// The user is required to specify a common Time Zone Location for use in converting
// date times to a common frame of reference for use in subsequent time duration calculations.
//
// Note: This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//       This means that duration is allocated over years, months, weeks, weekdays, date days,
//       hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//       See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime time.Time
//            - Starting date time
//
// duration  time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                 FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDtoTz(startTime, duration,
//                   TZones.US.Central(), FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartTimeDurationDateDtoTz(
	startDateTime DateTzDto,
	duration time.Duration,
	timeZoneLocation string,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationDateDtoTz() "

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartTimeDurationDateDtoCalcTz(
		&tDur2,
		startDateTime,
		duration,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDurationDateDtoTzCalc - Creates and returns a new TimeDurationDto based
// on input parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation
// type.
//
// Input parameter, 'startDateTime' is of Type DateTzDto.
//
// 'startDateTime' is converted to the specified 'timeZoneLocation' and the duration
// value is added to it in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and
// the actual starting date time is computed by subtracting duration.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime DateTzDto
//            - Starting date time for the duration calculation
//
// duration  time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()   - Default - standard year, month week,
//                    day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDtoTzCalc(startTime, duration,
//           TZones.US.Central(), TDurCalcType(0).StdYearMth(), FmtDateTimeYrMDayFmtStr)
//
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartTimeDurationDateDtoTzCalc(
	startDateTime DateTzDto,
	duration time.Duration,
	timeZoneLocation string,
	tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationDateDtoTzCalc() "

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartTimeDurationDateDtoCalcTz(
												&tDur2,
												startDateTime,
												duration,
												tDurCalcType,
												timeZoneLocation,
												dateTimeFmtStr,
												ePrefix)

	if err != nil {
		return TimeDurationDto{}, nil
	}

	return tDur2, nil
}

// NewStartTimeDurationDateDtoCalc - Creates and returns a new TimeDurationDto based on input
// parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation type.
//
// Input parameter 'startDateTime' is of Type DateTzDto.
//
// The duration value will be added to 'startDateTime' in order to compute the ending date time.
//
// If 'duration' is a negative value, 'startDateTime' is converted to ending date time and
// the actual starting date time will be computed by subtracting duration.
//
// The time zone location applied to both 'startDateTime' and ending date time is derived
// from input parameter, 'startDateTime'.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime   DateTz 
//            - Starting date time for the duration calculation
//
// duration time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// tDurCalcType TDurCalcType- Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// dateTimeFmtStr string  - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDtoCalc(startTime,
//            duration,
//             TDurCalcType(0).StdYearMth(),
//              FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//      standard year month day time duration allocation.
//
//      'FmtDateTimeYrMDayFmtStr' is a constant defined in the source
//       file, constantsdatetime.go.
//
func (tDur TimeDurationDto) NewStartTimeDurationDateDtoCalc(
	startDateTime DateTzDto,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationDateDtoCalc() "

	tDurDtoUtil := timeDurationDtoUtility{}

	timeZone := startDateTime.GetBestConvertibleTimeZone()

	var timeZoneLocationName string

	if timeZone.timeZoneType == TzType.Military() {
		timeZoneLocationName = timeZone.militaryTimeZoneName
	} else {
		timeZoneLocationName = timeZone.locationName
	}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimeDurationDateDtoCalcTz(
		&tDur2,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocationName,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimePlusTimeDto - Creates and returns a new TimeDurationDto setting
// the start date time, end date time and duration based on a starting date time
// and the time components contained in a TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// For the purposes of this time duration calculation, the Time Zone Location is
// extracted from the input parameter, 'startDateTime'.
//
// Note: This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//       This means that duration is allocated over years, months, weeks, weekdays, date days,
//       hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//       See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime time.Time
//            - Starting date time. The ending date time will be computed
//              by adding the time components of the 'plusTimeDto' to
//              'startDateTime'.
//
// plusTimeDto  TimeDto
//            - Time components (Years, months, weeks, days, hours etc.)
//              which will be added to 'startDateTime' to compute
//              time duration and ending date time.
//
//              type TimeDto struct {
//               Years          int // Number of Years
//               Months         int // Number of Months
//               Weeks          int // Number of Weeks
//               WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//               DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//               Hours          int // Number of Hours.
//               Minutes        int // Number of Minutes
//               Seconds        int // Number of Seconds
//               Milliseconds   int // Number of Milliseconds
//               Microseconds   int // Number of Microseconds
//               Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//               TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                        //  plus remaining Nanoseconds
//              }
//
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimePlusTimeDto(
//                    startTime,
//                    plusTimeDto,
//                    FmtDateTimeYrMDayFmtStr)
//
//  Note: 'FmtDateTimeYrMDayFmtStr' is a constant available in constantsdatetime.go
//
func (tDur TimeDurationDto) NewStartTimePlusTimeDto(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimePlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimePlusTimeDtoCalcTz(
										&tDur2,
										startDateTime,
										plusTimeDto,
										TDurCalcType(0).StdYearMth(),
										startDateTime.Location().String(),
										dateTimeFmtStr,
										ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, err
}

// NewStartTimePlusTimeDtoCalcTz - Creates and returns a new TimeDurationDto setting
// the start date time, end date time and duration based on a starting date time
// and the time components contained in a TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// The user is required to submit input parameters for time zone location and
// date time calculation type.
//
// Input Parameters:
// =================
//
// startDateTime time.Time
//            - Starting date time. The ending date time will be computed
//              by adding the time components of the 'plusTimeDto' to
//              'startDateTime'.
//
// plusTimeDto  TimeDto
//            - Time components (Years, months, weeks, days, hours etc.)
//              which will be added to 'startDateTime' to compute
//              time duration and ending date time.
//
//              type TimeDto struct {
//               Years          int // Number of Years
//               Months         int // Number of Months
//               Weeks          int // Number of Weeks
//               WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//               DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//               Hours          int // Number of Hours.
//               Minutes        int // Number of Minutes
//               Seconds        int // Number of Seconds
//               Milliseconds   int // Number of Milliseconds
//               Microseconds   int // Number of Microseconds
//               Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//               TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                        //  plus remaining Nanoseconds
//              }
//
//               Type 'TimeDto' is located in source file:
//                MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimePlusTimeDtoCalcTz(
//                    startTime,
//                    plusTimeDto,
//                    TDurCalcType(0).StdYearMth(),
//                    TZones.US.Central(),
//                    FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartTimePlusTimeDtoCalcTz(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimePlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimePlusTimeDtoCalcTz(
		&tDur2,
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewEndTimeMinusTimeDtoTz - Creates and returns a new TimeDurationDto setting
// start date time, end date time and duration based on an ending date time
// and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// The Time Zone Location used in the subsequent date time duration calculations
// is extracted from input parameter, 'endDateTime'. This Time Zone Location is
// applied to both the starting and ending date times.
//
// Note: This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//       This means that duration is allocated over years, months, weeks, weekdays, date days,
//       hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//       See Type 'TDurCalcType' for details.
//
// Input Parameters
// ================
//
// endDateTime time.Time
//            - Ending date time. The starting date time will be computed
//              by subtracting minusTimeDto from 'endDateTime'
//
// minusTimeDto TimeDto
//            - Time components (Years, months, weeks, days, hours etc.)
//              which will be subtracted from 'endDateTime' to compute
//              time duration and starting date time.
//
//              type TimeDto struct {
//               Years          int // Number of Years
//               Months         int // Number of Months
//               Weeks          int // Number of Weeks
//               WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//               DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//               Hours          int // Number of Hours.
//               Minutes        int // Number of Minutes
//               Seconds        int // Number of Seconds
//               Milliseconds   int // Number of Milliseconds
//               Microseconds   int // Number of Microseconds
//               Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//               TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                        //  plus remaining Nanoseconds
//              }
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewEndTimeMinusTimeDtoTz(
//                           endTime,
//                           minusTimeDto,
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note: 'FmtDateTimeYrMDayFmtStr' is a constant defined in source file,
//        constantsdatetime.go.
//
func (tDur TimeDurationDto) NewEndTimeMinusTimeDto(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewEndTimeMinusTimeDtoTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setEndTimeMinusTimeDtoCalcTz(
		&tDur2,
		endDateTime,
		minusTimeDto,
		TDurCalcType(0).StdYearMth(),
		endDateTime.Location().String(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, err
}

// NewEndTimeMinusTimeDtoCalcTz - Creates and returns a new TimeDurationDto setting
// start date time, end date time and duration based on an ending date time
// and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// The user is required to submit input parameters for time zone location and
// date time calculation type.
//
// Input Parameters:
// =================
//
// endDateTime time.Time
//            - Ending date time. The starting date time will be computed
//              by subtracting minusTimeDto from 'endDateTime'
//
// minusTimeDto TimeDto
//            - Time components (Years, months, weeks, days, hours etc.)
//              which will be subtracted from 'endDateTime' to compute
//              time duration and starting date time.
//
//              type TimeDto struct {
//               Years          int // Number of Years
//               Months         int // Number of Months
//               Weeks          int // Number of Weeks
//               WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//               DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//               Hours          int // Number of Hours.
//               Minutes        int // Number of Minutes
//               Seconds        int // Number of Seconds
//               Milliseconds   int // Number of Milliseconds
//               Microseconds   int // Number of Microseconds
//               Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//               TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                        //  plus remaining Nanoseconds
//              }
//
//              Type 'TimeDto' is located in source file:
//                MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time comparisons.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//               FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewEndTimeMinusTimeDtoCalcTz(endTime,
//                    minusTimeDto,
//                    TDurCalcType(0).StdYearMth(),
//                    TZones.US.Central()
//                    FmtDateTimeYrMDayFmtStr)
//
//  Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewEndTimeMinusTimeDtoCalcTz(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewEndTimeMinusTimeDtoTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := TimeDurationDto{}

	err := tDurDtoUtil.setEndTimeMinusTimeDtoCalcTz(
								&t2Dur,
								endDateTime,
								minusTimeDto,
								tDurCalcType,
								timeZoneLocation,
								dateTimeFmtStr,
								ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return t2Dur, nil
}

// ReCalcTimeDurationAllocation - Re-calculates and allocates time duration for the current
// TimeDurationDto instance over the various time components (years, months, weeks, weekdays,
// datedays, hour, minutes, seconds, milliseconds, microseconds and nanoseconds) depending
// on the value of the 'TDurCalcType' input parameter.
//
// Input Parameters
// ================
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
func (tDur *TimeDurationDto) ReCalcTimeDurationAllocation(
	tDurCalcType TDurCalcType) error {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.ReCalcTimeDurationAllocation() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.reCalcTimeDurationAllocation(
		tDur,
		tDurCalcType,
		ePrefix)
}

// ReCalcEndDateTimeToNow - Recomputes time duration values for the
// current TimeDurationDto by setting ending date time to time.Now().
// This is useful in stop watch applications.
//
// The Time Zone Location is derived from the existing starting date
// time, 'tDur.StartTimeDateTz'.  The Calculation type is taken from
// the existing calculation type, 'tDur.CalcType'.
//
func (tDur *TimeDurationDto) ReCalcEndDateTimeToNow() error {

	tDur.lock.Lock()
	
	defer tDur.lock.Unlock()
	
	ePrefix := "TimeDurationDto.ReCalcEndDateTimeToNow() "
	
	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.reCalcEndDateTimeToNow(tDur, ePrefix)
}

// SetAutoEnd - When called, this method automatically sets the ending date
// time and re-calculates the time duration for the current TimeDurationDto
// instance.
//
// Ending date time is assigned the value returned by time.Now(). This ending
// date time is converted to the specified Time Zone specified by the Time Zone
// Location associated with the current starting date time value.
//
// When used together, the two methods 'NewAutoStart' and this method, 'SetAutoEnd'
// function as a stop watch feature. Simply calling these functions can set
// the starting date time and later the ending date time to measure elapsed time, or
// time duration.
//
// The time duration calculation type is taken from the current TimeDurationDto
// calculation type setting.
//
func (tDur *TimeDurationDto) SetAutoEnd() error {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetAutoEnd() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setAutoEnd(
							tDur,
							ePrefix)
}

// SetEndTimeMinusTimeDtoCalcTz - Sets start date time, end date time and duration
// based on an ending date time and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// Input Parameters:
// =================
//
// endDateTime time.Time
//            - Ending date time. The starting date time will be computed
//              by subtracting minusTimeDto from 'endDateTime'
//
// minusTimeDto TimeDto
//            - Time components (Years, months, weeks, days, hours etc.)
//              which will be subtracted from 'endDateTime' to compute
//              time duration and starting date time.
//
//              type TimeDto struct {
//               Years          int // Number of Years
//               Months         int // Number of Months
//               Weeks          int // Number of Weeks
//               WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//               DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//               Hours          int // Number of Hours.
//               Minutes        int // Number of Minutes
//               Seconds        int // Number of Seconds
//               Milliseconds   int // Number of Milliseconds
//               Microseconds   int // Number of Microseconds
//               Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//               TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                        //  plus remaining Nanoseconds
//              }
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                    Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time comparisons.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetEndTimeMinusTimeDtoCalcTz(
	endDateTime time.Time,
	minusTimeDto TimeDto, 
	tDurCalcType TDurCalcType, 
	timeZoneLocation, 
	dateTimeFmtStr string) error {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetEndTimeMinusTimeDtoCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setEndTimeMinusTimeDtoCalcTz(
		tDur,
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartEndTimesDateDtoCalcTz - Sets data field values for the current
// TimeDurationDto instance using a Start Date Time, End Date Time and a
// time zone specification.
//
// The Starting Date Time and Ending Date Time are submitted as type 'DateTzDto'
//
// First, 'startDateTime' and 'endDateTime' are converted to the designate Time
// Zone Location. Next, 'startDateTime' is subtracted from 'endDateTime' to compute
// time duration.
//
// The user is required to submit input parameters for time zone location and
// date time calculation type.
//
// All data fields in the current TimeDurationDto instance are overwritten with
// the new time duration values.
//
// Input Parameters:
// =================
//
// startDateTime DateTzDto - Starting date time
//
// endDateTime   DateTzDto - Ending date time
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                    Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string 
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string  
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                 FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartEndTimesDateDtoCalcTz(
	startDateTimeTz,
	endDateTimeTz DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	tDur.lock.Lock()
	
	defer tDur.lock.Unlock()
	
	ePrefix := "TimeDurationDto.SetStartEndTimesDateDtoCalcTz() "
	
	tDurDtoUtil := timeDurationDtoUtility{}
	
	return tDurDtoUtil.setStartEndTimesDateDtoCalcTz(
			tDur,
			startDateTimeTz,
			endDateTimeTz,
			tDurCalcType,
			timeZoneLocation,
			dateTimeFmtStr,
			ePrefix)
}

// SetStartEndTimesCalcTz - Sets data field values for the current TimeDurationDto
// instance using a Start Date Time, End Date Time and a time zone specification.
// First, 'startDateTime' and 'endDateTime' are converted to the designate Time
// Zone Location. Next, 'startDateTime' is subtracted from 'endDateTime' to compute
// time duration.
//
// All data fields in the current TimeDurationDto instance are overwritten with
// the new time duration values.
//
// Input Parameters:
// =================
//
// startDateTime time.Time - Starting time
//
// endDateTime  time.Time - Ending time
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                     or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//             MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartEndTimesCalcTz(
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartEndTimesCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartEndTimesCalcTz(
							tDur,
							startDateTime,
							endDateTime,
							tDurCalcType,
							timeZoneLocation,
							dateTimeFmtStr,
							ePrefix)
}

// SetStartTimeDurationCalcTz - Sets start time, end time and duration for the
// current TimeDurationDto instance. 'startDateTime' is converted to the
// specified 'timeZoneLocation' and the duration value is added to it
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to
// ending date time and the actual starting date time is computed by
// subtracting duration.
//
// Input Parameters:
// =================
//
// startDateTime time.Time
//            - Starting date time for the duration calculation
//
// duration  time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()   - Default - standard year, month week,
//                                      day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartTimeDurationCalcTz(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()
	
	ePrefix := "TimeDurationDto.SetStartTimeDurationCalcTz() "
	
	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimeDurationCalcTz(
						tDur,
						startDateTime,
						duration,
						tDurCalcType,
						timeZoneLocation,
						dateTimeFmtStr,
						ePrefix)
}

// SetStartTimeDurationDateDtoCalcTz - Sets start time, end time and
// duration for the current TimeDurationDto instance.
//
// The input parameter, 'startDateTime', is of type DateTzDto. It is
// converted to the specified 'timeZoneLocation' and the duration value
// is added to in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to
// ending date time and the actual starting date time is computed by
// subtracting duration.
//
// Input Parameters:
// =================
//
// startDateTime DateTzDto
//            - Provides starting date time for the duration calculation
//
// duration  time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//             - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                 FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
func (tDur *TimeDurationDto) SetStartTimeDurationDateDtoCalcTz(
	startDateTimeTz DateTzDto,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartTimeDurationDateDtoCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimeDurationDateDtoCalcTz(
						tDur,
						startDateTimeTz,
						duration,
						tDurCalcType,
						timeZoneLocation,
						dateTimeFmtStr,
						ePrefix)
}

// SetStartTimePlusTimeDtoCalcTz - Sets start date time, end date time and duration
// based on a starting date time and the time components contained in a TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// Input Parameters:
// =================
//
// startDateTime time.Time
//            - Starting date time. The ending date time will be computed
//              by adding the time components of the 'plusTimeDto' to
//              'startDateTime'.
//
// plusTimeDto  TimeDto
//            -  Time components (Years, months, weeks, days, hours etc.)
//               which will be added to 'startDateTime' to compute
//               time duration and ending date time.
//
//               type TimeDto struct {
//                Years          int // Number of Years
//                Months         int // Number of Months
//                Weeks          int // Number of Weeks
//                WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//                DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//                Hours          int // Number of Hours.
//                Minutes        int // Number of Minutes
//                Seconds        int // Number of Seconds
//                Milliseconds   int // Number of Milliseconds
//                Microseconds   int // Number of Microseconds
//                Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//                TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                         //  plus remaining Nanoseconds
//               }
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears()
//                                      - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time comparisons.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartTimePlusTimeDtoCalcTz(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartTimePlusTimeDtoCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimePlusTimeDtoCalcTz(
								tDur,
								startDateTime,
								plusTimeDto,
								tDurCalcType,
								timeZoneLocation,
								dateTimeFmtStr,
								ePrefix)
}

// XX
// calcTimeDurationAllocations - Examines the input parameter 'calcType' and
// then determines which type of time duration allocation calculation will be
// applied to the data fields of the current TimeDurationDto instance.
func (tDur *TimeDurationDto) calcTimeDurationAllocations(calcType TDurCalcType) error {

	ePrefix := "TimeDurationDto.calcTimeDurationAllocations() "

	switch calcType {

	case TDurCalcType(0).StdYearMth():
		return tDur.calcTypeSTDYEARMTH()

	case TDurCalcType(0).CumMonths():
		return tDur.calcTypeCUMMONTHS()

	case TDurCalcType(0).CumWeeks():
		return tDur.calcTypeCUMWEEKS()

	case TDurCalcType(0).CumDays():
		return tDur.calcTypeCUMDays()

	case TDurCalcType(0).CumHours():
		return tDur.calcTypeCUMHours()

	case TDurCalcType(0).CumMinutes():
		return tDur.calcTypeCUMMINUTES()

	case TDurCalcType(0).CumSeconds():
		return tDur.calcTypeCUMSECONDS()

	case TDurCalcType(0).GregorianYears():
		return tDur.calcTypeGregorianYears()
	}

	return fmt.Errorf(ePrefix+
		"Error: Invalid TDurCalcType. calcType='%v'", calcType.String())
}

// calcTypeCUMDays - Calculates Cumulative Days. Years, months and weeks are consolidated
// and counted as cumulative days. The Data Fields for years, months, weeks and week days
// are set to zero.  All cumulative days are allocated to the data field, 'DateDays'.
func (tDur *TimeDurationDto) calcTypeCUMDays() error {

	ePrefix := "TimeDurationDto.calcTypeCUMDays() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumDays()

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	if rd >= DayNanoSeconds {
		tDur.DateDays = rd / DayNanoSeconds
		tDur.DateDaysNanosecs = tDur.DateDays * DayNanoSeconds
	}

	err := tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}

// calcTypeCUMHours - Calculates Cumulative Hours. Years, months, weeks, week days,
// date days and hours are consolidated and included in cumulative hours. Values for years,
// months, weeks, week days and date days are ignored and set to zero. Time duration is
// allocated over cumulative hours plus minutes, seconds, milliseconds, microseconds and
// nanoseconds.
func (tDur *TimeDurationDto) calcTypeCUMHours() error {

	ePrefix := "TimeDurationDto.calcTypeCUMHours() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumHours()

	if tDur.TimeDuration == 0 {
		return nil
	}

	err := tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}

// calcTypeGregorianYears - Allocates Years using the number of nanoseconds in a
// standard or average GregorianYear
func (tDur *TimeDurationDto) calcTypeGregorianYears() error {
	ePrefix := "TimeDurationDto.calcTypeGregorianYears() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).GregorianYears()

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	if rd >= GregorianYearNanoSeconds {
		tDur.Years = rd / GregorianYearNanoSeconds
		tDur.YearsNanosecs = tDur.Years * GregorianYearNanoSeconds
	}

	err := tDur.calcMonthsFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcMonthsFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcDateDaysWeeksFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcDateDaysWeeksFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}

// calcTypeCUMMINUTES - Calculates Cumulative Minutes. Years, months, weeks, week days,
// date days, hours and minutes are consolidated and included in cumulative minutes.
// Values for years, months, weeks, week days, date days and hours are ignored and set
// to zero. Time duration is allocated over cumulative minutes plus seconds, milliseconds,
// microseconds and nanoseconds.
func (tDur *TimeDurationDto) calcTypeCUMMINUTES() error {

	ePrefix := "TimeDurationDto.calcTypeCUMHours() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumMinutes()

	if tDur.TimeDuration == 0 {
		return nil
	}

	rd := int64(tDur.TimeDuration)

	if rd >= MinuteNanoSeconds {
		tDur.Minutes = rd / MinuteNanoSeconds
		tDur.MinutesNanosecs = tDur.Minutes * MinuteNanoSeconds
		rd -= tDur.MinutesNanosecs
	}

	if rd >= SecondNanoseconds {
		tDur.Seconds = rd / SecondNanoseconds
		tDur.SecondsNanosecs = tDur.Seconds * SecondNanoseconds
		rd -= tDur.SecondsNanosecs
	}

	if rd >= MilliSecondNanoseconds {
		tDur.Milliseconds = rd / MilliSecondNanoseconds
		tDur.MillisecondsNanosecs = tDur.Milliseconds * MilliSecondNanoseconds
		rd -= tDur.MillisecondsNanosecs
	}

	if rd >= MicroSecondNanoseconds {
		tDur.Microseconds = rd / MicroSecondNanoseconds
		tDur.MillisecondsNanosecs = tDur.Microseconds * MicroSecondNanoseconds
		rd -= tDur.MicrosecondsNanosecs
	}

	tDur.Nanoseconds = rd

	err := tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil

}

// Data Fields for Years is always set to Zero. Years
// and months are consolidated and counted as cumulative
// months.
func (tDur *TimeDurationDto) calcTypeCUMMONTHS() error {

	ePrefix := "TimeDurationDto.calcTypeCUMWEEK() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumMonths()

	err := tDur.calcMonthsFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcMonthsFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcDateDaysWeeksFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcDateDaysWeeksFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil

}

// calcTypeCUMSECONDS - Calculates Cumulative Seconds of
// time duration.
//
// tDur.CalcType = TDurCalcType(0).CumSeconds()
//
// Years, months, weeks, weekdays, date days, hours and
// minutes are ignored and set to zero. Time is accumulated
// in seconds, milliseconds, microseconds and nanoseconds.
//
func (tDur *TimeDurationDto) calcTypeCUMSECONDS() error {

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.emptyTimeFields(
		tDur,
		"TimeDurationDto.calcTypeCUMSECONDS() ")

	tDur.CalcType = TDurCalcType(0).CumSeconds()

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	if rd >= SecondNanoseconds {
		tDur.Seconds = rd / SecondNanoseconds
		tDur.SecondsNanosecs = SecondNanoseconds * tDur.Seconds
		rd -= tDur.SecondsNanosecs
	}

	if rd >= MilliSecondNanoseconds {
		tDur.Milliseconds = rd / MilliSecondNanoseconds
		tDur.MillisecondsNanosecs = MilliSecondNanoseconds * tDur.Milliseconds
		rd -= tDur.MillisecondsNanosecs
	}

	if rd >= MicroSecondNanoseconds {
		tDur.Microseconds = rd / MicroSecondNanoseconds
		tDur.MicrosecondsNanosecs = MicroSecondNanoseconds * tDur.Microseconds
		rd -= tDur.MicrosecondsNanosecs
	}

	tDur.Nanoseconds = rd

	return nil
}

// calcTypeSTDYEARMTH - Performs Duration calculations for
// TDurCalcType == TDurCalcType(0).StdYearMth()
//
// TDurCalcTypeYEARMTH - Standard Year, Month, Weeks, Days calculation.
// All data fields in the TimeDto are populated in the duration
// allocation.
func (tDur *TimeDurationDto) calcTypeSTDYEARMTH() error {

	ePrefix := "TimeDurationDto.calcTypeSTDYEARMTH() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).StdYearMth()

	err := tDur.calcYearsFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcYearsFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcMonthsFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcMonthsFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcDateDaysWeeksFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcDateDaysWeeksFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}

// calcTypeCUMWEEKS - Data Fields for Years and Months are always set to zero.
// Years and Months are consolidated and counted as equivalent Weeks.
func (tDur *TimeDurationDto) calcTypeCUMWEEKS() error {

	ePrefix := "TimeDurationDto.calcTypeCUMWEEKS() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumWeeks()

	rd := int64(tDur.TimeDuration)

	if rd >= WeekNanoSeconds {

		tDur.Weeks = rd / WeekNanoSeconds
		tDur.WeeksNanosecs = tDur.Weeks * WeekNanoSeconds
		rd -= tDur.WeeksNanosecs

	}

	if rd >= DayNanoSeconds {
		tDur.WeekDays = rd / DayNanoSeconds
		tDur.WeekDaysNanosecs = tDur.WeekDays * DayNanoSeconds
		rd -= tDur.WeekDaysNanosecs
	}

	tDur.DateDays = tDur.Weeks * int64(7)
	tDur.DateDays += tDur.WeekDays
	tDur.DateDaysNanosecs = tDur.WeeksNanosecs + tDur.WeekDaysNanosecs

	err := tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	// For Cumulative Weeks calculation and presentations, set Date Days to zero
	tDur.DateDays = 0
	tDur.DateDaysNanosecs = 0

	return nil
}

// calcYearsFromDuration - Calculates number of years duration and nanoseconds
// represented by years duration using input parameters 'tDur.StartTimeDateTz' and
// 'tDur.EndTimeDateTz'.
//
// NOTE: Before calling this method, ensure that tDur.StartTimeDateTz,
//       tDur.EndTimeDateTz and tDur.TimeDuration are properly initialized.
//
func (tDur *TimeDurationDto) calcYearsFromDuration() error {

	ePrefix := "TimeDurationDto.calcYearsFromDuration() "

	years := int64(0)
	yearNanosecs := int64(0)
	startTime := tDur.StartTimeDateTz.GetDateTimeValue()
	endTime := tDur.EndTimeDateTz.GetDateTimeValue()

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix+"Error: 'startTime' and 'endTime' Time Zone Location do NOT match! "+
			"startTimeZoneLocation='%v'  endTimeZoneLocation='%v'",
			startTime.Location().String(), endTime.Location().String())
	}

	yearDateTime := startTime

	i := 0

	for yearDateTime.Before(endTime) {

		i++

		yearDateTime = startTime.AddDate(i, 0, 0)

	}

	i--

	if i > 0 {

		years = int64(i)

		yearDateTime = startTime.AddDate(i, 0, 0)

		duration := yearDateTime.Sub(startTime)

		yearNanosecs = int64(duration)

	} else {

		years = 0

		yearNanosecs = 0
	}

	tDur.Years = years
	tDur.YearsNanosecs = yearNanosecs

	return nil
}

// calcMonthsFromDuration - calculates the months duration
// using the start and end dates, 'tDur.StartTimeDateTz' and
// 'tDur.EndTimeDateTz.DateTime'.
//
// NOTE: (1) Before calling this method, ensure that tDur.StartTimeDateTz,
//           tDur.EndTimeDateTz and tDur.TimeDuration are properly initialized.
//
//       (2) Before calling this method, ensure that the following method is called
//           first:
//              TimeDurationDto.calcYearsFromDuration
//
func (tDur *TimeDurationDto) calcMonthsFromDuration() error {

	ePrefix := "TimeDurationDto.calcMonthsFromDuration() "

	startTime := tDur.StartTimeDateTz.GetDateTimeValue()
	endTime := tDur.EndTimeDateTz.GetDateTimeValue()

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix+"Error: 'startTime' and 'endTime' Time Zone Location do NOT match! "+
			"startTimeZoneLocation='%v'  endTimeZoneLocation='%v'",
			startTime.Location().String(), endTime.Location().String())
	}

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	rd -= tDur.YearsNanosecs

	i := 0

	yearDateTime := startTime.Add(time.Duration(tDur.YearsNanosecs))

	mthDateTime := yearDateTime

	for mthDateTime.Before(endTime) {

		i++

		mthDateTime = yearDateTime.AddDate(0, i, 0)

	}

	i -= 1

	if i > 0 {

		tDur.Months = int64(i)

		mthDateTime = yearDateTime.AddDate(0, i, 0)

		tDur.MonthsNanosecs = int64(mthDateTime.Sub(yearDateTime))

	} else {
		tDur.Months = 0
		tDur.MonthsNanosecs = 0
	}

	return nil
}

// calcDateDaysWeeksFromDuration - Calculates the Days associated
// with the duration for this TimeDurationDto.
//
// Calculates 'tDur.DateDays', 'tDur.DateDaysNanosecs', 'tDur.Weeks', 'tDur.WeeksNanosecs',
// 'tDur.WeekDays' and 'tDur.WeekDaysNanosecs'.
//
// NOTE: (1) Before calling this method, ensure that TimeDurationDto.StartTimeDateTz,
//           TimeDurationDto.EndTimeDateTz and tDur.TimeDuration are properly initialized.
//
//       (2) Before calling this method, ensure that the following methods are called
//           first, in sequence:
//             TimeDurationDto.calcYearsFromDuration
//             TimeDurationDto.calcMonthsFromDuration
//
func (tDur *TimeDurationDto) calcDateDaysWeeksFromDuration() error {

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs

	// Calculate DateDays
	tDur.DateDays = 0
	tDur.DateDaysNanosecs = 0

	if rd >= DayNanoSeconds {
		tDur.DateDays = rd / DayNanoSeconds
		tDur.DateDaysNanosecs = DayNanoSeconds * tDur.DateDays
	}

	// Calculate Weeks and WeekDays
	tDur.Weeks = 0
	tDur.WeeksNanosecs = 0
	tDur.WeekDays = 0
	tDur.WeekDaysNanosecs = 0

	if tDur.DateDays > 0 {

		if tDur.DateDays >= 7 {

			tDur.Weeks = tDur.DateDays / int64(7)
			tDur.WeeksNanosecs = WeekNanoSeconds * tDur.Weeks

		}

		tDur.WeekDays = tDur.DateDays - (tDur.Weeks * 7)
		tDur.WeekDaysNanosecs = tDur.WeekDays * DayNanoSeconds

	}

	return nil
}

// calcHoursMinSecs - Calculates Hours, Minute, and
// Seconds of duration using startTime, tDur.StartTimeDateTz,
// and endTime, tDur.EndTimeDateTz.DateTime.
//
//
// NOTE: (1) Before calling this method, ensure that tDur.StartTimeDateTz,
//           TimeDurationDto.EndTimeDateTz and TimeDurationDto.TimeDuration
//           are properly initialized.
//
//       (2) Before calling this method, ensure that the following methods are called
//           first, in sequence:
//             TimeDurationDto.calcYearsFromDuration
//             TimeDurationDto.calcMonthsFromDuration
//             TimeDurationDto.calcDateDaysWeeksFromDuration
//
func (tDur *TimeDurationDto) calcHoursMinSecs() error {

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	if tDur.DateDays > 0 {
		rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs +
			tDur.DateDaysNanosecs
	} else {
		rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs +
			tDur.WeeksNanosecs + tDur.WeekDaysNanosecs
	}

	tDur.Hours = 0
	tDur.HoursNanosecs = 0
	tDur.Minutes = 0
	tDur.MinutesNanosecs = 0
	tDur.Seconds = 0
	tDur.SecondsNanosecs = 0

	if rd >= HourNanoSeconds {
		tDur.Hours = rd / HourNanoSeconds
		tDur.HoursNanosecs = HourNanoSeconds * tDur.Hours
		rd -= tDur.HoursNanosecs
	}

	if rd >= MinuteNanoSeconds {
		tDur.Minutes = rd / MinuteNanoSeconds
		tDur.MinutesNanosecs = MinuteNanoSeconds * tDur.Minutes
		rd -= tDur.MinutesNanosecs
	}

	if rd >= SecondNanoseconds {
		tDur.Seconds = rd / SecondNanoseconds
		tDur.SecondsNanosecs = SecondNanoseconds * tDur.Seconds
		rd -= tDur.SecondsNanosecs
	}

	return nil
}

// calcNanoseconds - Calculates 'tDur.Milliseconds', 'tDur.MillisecondsNanosecs',
// 'tDur.Microseconds', 'tDur.MicrosecondsNanosecs',  and 'tDur.Nanoseconds'.
//
// NOTE: (1) Before calling this method, ensure that tDur.StartTimeDateTz,
//           TimeDurationDto.EndTimeDateTz and TimeDurationDto.TimeDuration
//           are properly initialized.
//
//       (2) Before calling this method, ensure that the following methods are called
//           first, in sequence:
//             TimeDurationDto.calcYearsFromDuration
//             TimeDurationDto.calcMonthsFromDuration
//             TimeDurationDto.calcDateDaysWeeksFromDuration
//             TimeDurationDto.calcHoursMinSecs
//
func (tDur *TimeDurationDto) calcNanoseconds() error {

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs

	if tDur.DateDaysNanosecs > 0 {
		rd -= tDur.DateDaysNanosecs
	} else {
		rd -= tDur.WeeksNanosecs + tDur.WeekDaysNanosecs
	}

	rd -= tDur.HoursNanosecs +
		tDur.MinutesNanosecs + tDur.SecondsNanosecs

	tDur.Milliseconds = 0
	tDur.MillisecondsNanosecs = 0
	tDur.Microseconds = 0
	tDur.MicrosecondsNanosecs = 0
	tDur.Nanoseconds = 0

	if rd >= MilliSecondNanoseconds {
		tDur.Milliseconds = rd / MilliSecondNanoseconds
		tDur.MillisecondsNanosecs = MilliSecondNanoseconds * tDur.Milliseconds
		rd -= tDur.MillisecondsNanosecs
	}

	if rd >= MicroSecondNanoseconds {
		tDur.Microseconds = rd / MicroSecondNanoseconds
		tDur.MicrosecondsNanosecs = MicroSecondNanoseconds * tDur.Microseconds
		rd -= tDur.MicrosecondsNanosecs
	}

	tDur.Nanoseconds = rd

	return nil
}

// calcSummaryTimeElements - Calculates totals for Date, Time and
// sub-second nanoseconds.
//
// NOTE: (1) Before calling this method, ensure that tDur.StartTimeDateTz,
//           TimeDurationDto.EndTimeDateTz and TimeDurationDto.TimeDuration
//           are properly initialized.
//
//       (2) Before calling this method, ensure that the following methods are called
//           first, in sequence:
//             TimeDurationDto.calcYearsFromDuration
//             TimeDurationDto.calcMonthsFromDuration
//             TimeDurationDto.calcDateDaysWeeksFromDuration
//             TimeDurationDto.calcHoursMinSecs
//             TimeDurationDto.calcNanoseconds
//
func (tDur *TimeDurationDto) calcSummaryTimeElements() error {

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	tDur.TotDateNanoseconds = 0
	tDur.TotTimeNanoseconds = 0
	tDur.TotSubSecNanoseconds = 0

	tDur.TotDateNanoseconds = tDur.YearsNanosecs
	tDur.TotDateNanoseconds += tDur.MonthsNanosecs

	if tDur.DateDaysNanosecs == 0 {
		tDur.TotDateNanoseconds += tDur.WeeksNanosecs
		tDur.TotDateNanoseconds += tDur.WeekDaysNanosecs
	} else {
		tDur.TotDateNanoseconds += tDur.DateDaysNanosecs
	}

	tDur.TotSubSecNanoseconds = tDur.MillisecondsNanosecs
	tDur.TotSubSecNanoseconds += tDur.MicrosecondsNanosecs
	tDur.TotSubSecNanoseconds += tDur.Nanoseconds

	tDur.TotTimeNanoseconds = tDur.HoursNanosecs
	tDur.TotTimeNanoseconds += tDur.MinutesNanosecs
	tDur.TotTimeNanoseconds += tDur.SecondsNanosecs
	tDur.TotTimeNanoseconds += tDur.TotSubSecNanoseconds

	return nil
}

func (tDur *TimeDurationDto) preProcessDateFormatStr(dateTimeFmtStr string) string {

	if len(dateTimeFmtStr) == 0 {
		return FmtDateTimeYrMDayFmtStr
	}

	return dateTimeFmtStr
}

func (tDur *TimeDurationDto) preProcessTimeZoneLocation(timeZoneLocation string) string {

	if len(timeZoneLocation) == 0 {
		return TZones.UTC()
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		return TZones.Local()
	}

	return timeZoneLocation
}
