package datetime

import (
	"time"
	"fmt"
)

// TimeDto - used for transmitting
// time elements.
type TimeDto struct {
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

// CopyOut - Creates a new TimeDto instance
// which precisely duplicates the current TimeDto
// instance, and returns it to the calling function.
func (tDto *TimeDto) CopyOut() TimeDto {

	tDto2 := TimeDto{}

	tDto2.Years 				=  tDto.Years
	tDto2.Months       	=  tDto.Months
	tDto2.Weeks        	=  tDto.Weeks
	tDto2.Days         	=  tDto.Days
	tDto2.Hours        	=  tDto.Hours
	tDto2.Minutes      	=  tDto.Minutes
	tDto2.Seconds      	=  tDto.Seconds
	tDto2.Milliseconds 	=  tDto.Milliseconds
	tDto2.Microseconds 	=  tDto.Microseconds
	tDto2.Nanoseconds  	=  tDto.Nanoseconds

	return tDto2
}

// CopyIn - Receives a TimeDto input parameter, 'tDto2'
// and proceeds to copy all 'tDto2' data fields into
// the current TimeDto data fields. When this method
// completes, 'tDto' will be equivalent to 'tDto2'.
func (tDto *TimeDto) CopyIn(tDto2 TimeDto) {

	tDto.Empty()

	tDto.Years 					=  tDto2.Years
	tDto.Months       	=  tDto2.Months
	tDto.Weeks        	=  tDto2.Weeks
	tDto.Days         	=  tDto2.Days
	tDto.Hours        	=  tDto2.Hours
	tDto.Minutes      	=  tDto2.Minutes
	tDto.Seconds      	=  tDto2.Seconds
	tDto.Milliseconds 	=  tDto2.Milliseconds
	tDto.Microseconds 	=  tDto2.Microseconds
	tDto.Nanoseconds  	=  tDto2.Nanoseconds


}


// ConvertToAbsoluteValues - Converts time components
// (Years, months, weeks days, hours, seconds, etc.)
// to absolute values.
func (tDto *TimeDto) ConvertToAbsoluteValues() {
	if tDto.Years < 0 {
		tDto.Years *= -1
	}

	if tDto.Months < 0 {
		tDto.Months *= -1
	}

	if tDto.Weeks < 0 {
		tDto.Weeks *= -1
	}

	if tDto.Days < 0 {
		tDto.Days *= -1
	}

	if tDto.Hours < 0 {
		tDto.Hours *= -1
	}

	if tDto.Minutes < 0 {
		tDto.Minutes *= -1
	}

	if tDto.Seconds < 0 {
		tDto.Seconds *= -1
	}

	if tDto.Milliseconds < 0 {
		tDto.Milliseconds *= -1
	}

	if tDto.Milliseconds < 0 {
		tDto.Milliseconds *= -1
	}

	if tDto.Microseconds < 0 {
		tDto.Microseconds *= -1
	}

	if tDto.Nanoseconds < 0 {
		tDto.Nanoseconds *= -1
	}

}

// ConvertToNegativeValues - Multiplies time component
// values by -1
func (tDto *TimeDto) ConvertToNegativeValues() {
	tDto.ConvertToAbsoluteValues()
	tDto.Years 				*= -1
	tDto.Months 			*= -1
	tDto.Weeks 				*= -1
	tDto.Days 				*= -1
	tDto.Hours 				*= -1
	tDto.Minutes 			*= -1
	tDto.Seconds 			*= -1
	tDto.Milliseconds *= -1
	tDto.Microseconds *= -1
	tDto.Nanoseconds 	*= -1
}

// Empty - returns all TimeDto data fields to their
// uninitialized or zero state.
func (tDto *TimeDto) Empty() {
	tDto.Years 					= 0
	tDto.Months 				= 0
	tDto.Weeks 					= 0
	tDto.Days 					= 0
	tDto.Hours 					= 0
	tDto.Minutes 				= 0
	tDto.Seconds 				= 0
	tDto.Milliseconds 	= 0
	tDto.Microseconds 	= 0
	tDto.Nanoseconds 		= 0
}

// Equal - Compares the data fields of input parameter TimeDto, 'tDto2',
// to the data fields of the current TimeDto, 'tDto'. If all data fields
// are equal, this method returns 'true'. Otherwise, the method returns
// false.
func (tDto *TimeDto) Equal(tDto2 TimeDto) bool {

	if tDto.Years		!=  tDto2.Years 						||
		tDto.Months			!=  tDto2.Months					||
		tDto.Weeks			!=  tDto2.Weeks						||
		tDto.Days 			!=  tDto2.Days						||
		tDto.Hours			!=  tDto2.Hours						||
		tDto.Minutes 		!=  tDto2.Minutes					||
		tDto.Seconds    	!=  tDto2.Seconds				||
		tDto.Milliseconds !=  tDto2.Milliseconds 	||
		tDto.Microseconds !=  tDto2.Microseconds	||
		tDto.Nanoseconds  !=  tDto2.Nanoseconds {

		return false
	}

	return true
}

// GetTotalDaysFromWeeksAndDays - Computes total number
// of days from tDto.Weeks plus tDto.Days
func (tDto *TimeDto) GetTotalDaysFromWeeksAndDays() int64 {

	totDays := tDto.Weeks * int64(7)
	totDays += tDto.Days

	return totDays

}

// GetTotalNanoSecs - Computes total number of nanoseconds
// from tDto.Milliseconds plus tDto.Microseconds plus
// tDto.Nanoseconds.
func (tDto *TimeDto) GetTotalNanoSecs() int64 {

	totNanoSecs := tDto.Milliseconds * int64(time.Millisecond)
	totNanoSecs += tDto.Microseconds * int64(time.Microsecond)
	totNanoSecs += tDto.Nanoseconds

	return totNanoSecs
}

func (tDto TimeDto) New(years, months, weeks, days, hours, minutes,
				seconds, milliseconds, microseconds, nanoseconds int) (TimeDto, error) {

	ePrefix := "TimeDto.New(...) "

	if years 				== 0 &&
		months 				== 0 &&
		weeks 				== 0 &&
		days 					== 0 &&
		hours 				== 0 &&
		minutes 			== 0 &&
		seconds 			== 0 &&
		milliseconds 	== 0 &&
		microseconds 	== 0 &&
		nanoseconds 	== 0 {

			return TimeDto{}, fmt.Errorf(ePrefix + "Error: All input parameters (years, months, weeks, days etc.) are ZERO Value!")
	}

	weeksAndDays := int64((weeks * 7) + days)

	totalNanoSecs := int64(milliseconds) * int64(time.Millisecond)
	totalNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totalNanoSecs += int64(nanoseconds)

	totalSeconds := int64(hours) * int64(3600)
	totalSeconds += int64(minutes) * int64(60)
	totalSeconds += int64(seconds)

	t2Dto := TimeDto{}

	t2Dto.Years = int64(years)
	t2Dto.Months = int64(months)
	t2Dto.allocateWeeksAndDays(weeksAndDays)
	t2Dto.allocateSeconds(totalSeconds)
	t2Dto.allocateTotalNanoseconds(totalNanoSecs)

	return t2Dto, nil
}

func (tDto *TimeDto) allocateWeeksAndDays(totalDays int64) {
	tDto.Weeks = 0
	tDto.Days = 0

	if totalDays >= 7 {
		tDto.Weeks = totalDays / 7
		totalDays -= tDto.Weeks * 7
	}

	tDto.Days = totalDays

}

func (tDto *TimeDto) allocateSeconds(totalSeconds int64) {
	tDto.Hours = 0
	tDto.Minutes = 0
	tDto.Seconds = 0

	if totalSeconds >= 3600 {
		tDto.Hours = totalSeconds / 3600
		totalSeconds -= tDto.Hours * 3600
	}

	if totalSeconds >= 60 {
		tDto.Minutes = totalSeconds / 60
		totalSeconds -= tDto.Minutes * 60
	}

	tDto.Seconds = totalSeconds
}

// allocateTotalNanoseconds - Allocates total nanoseconds to current
// TimeDto instance data fields: milliseconds, microseconds and
// nanoseconds.
func (tDto *TimeDto) allocateTotalNanoseconds(totalNanoSeconds int64) {

	tDto.Milliseconds = 0
	tDto.Microseconds = 0
	tDto.Nanoseconds = 0

	if totalNanoSeconds >= int64(time.Millisecond) {
		tDto.Milliseconds = totalNanoSeconds / int64(time.Millisecond)
		totalNanoSeconds -= tDto.Milliseconds * int64(time.Millisecond)
	}

	if totalNanoSeconds >= int64(time.Millisecond) {
		tDto.Microseconds = totalNanoSeconds / int64(time.Microsecond)
		totalNanoSeconds -= tDto.Microseconds * int64(time.Microsecond)
	}

	tDto.Nanoseconds = totalNanoSeconds

}

