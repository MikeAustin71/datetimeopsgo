package datetime

import (
	"time"
	"fmt"
	"errors"
	"strings"
)

// TimeDto - used for transmitting
// time elements.
type TimeDto struct {
	Years          int // Number of Years
	Months         int // Number of Months
	Weeks          int // Number of Weeks
	WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
	DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
	Hours          int // Number of Hours.
	Minutes        int // Number of Minutes
	Seconds        int // Number of Seconds
	Milliseconds   int // Number of Milliseconds
	Microseconds   int // Number of Microseconds
	Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
	TotNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
													// 	plus remaining Nanoseconds
}

// CopyOut - Creates a new TimeDto instance
// which precisely duplicates the current TimeDto
// instance, and returns it to the calling function.
func (tDto *TimeDto) CopyOut() TimeDto {

	tDto2 := TimeDto{}

	tDto2.Years 					=  tDto.Years
	tDto2.Months       		=  tDto.Months
	tDto2.Weeks        		=  tDto.Weeks
	tDto2.WeekDays 				=  tDto.WeekDays
	tDto2.DateDays 				=  tDto.DateDays
	tDto2.Hours        		=  tDto.Hours
	tDto2.Minutes      		=  tDto.Minutes
	tDto2.Seconds      		=  tDto.Seconds
	tDto2.Milliseconds 		=  tDto.Milliseconds
	tDto2.Microseconds 		=  tDto.Microseconds
	tDto2.Nanoseconds  		=  tDto.Nanoseconds
	tDto2.TotNanoseconds 	= tDto.TotNanoseconds

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
	tDto.WeekDays 			=  tDto2.WeekDays
	tDto.DateDays				 =  tDto2.DateDays
	tDto.Hours        	=  tDto2.Hours
	tDto.Minutes      	=  tDto2.Minutes
	tDto.Seconds      	=  tDto2.Seconds
	tDto.Milliseconds 	=  tDto2.Milliseconds
	tDto.Microseconds 	=  tDto2.Microseconds
	tDto.Nanoseconds  	=  tDto2.Nanoseconds
	tDto.TotNanoseconds	=  tDto2.TotNanoseconds


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

	if tDto.WeekDays < 0 {
		tDto.WeekDays *= -1
	}

	if tDto.DateDays < 0 {
		tDto.DateDays *= -1
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

	if tDto.TotNanoseconds < 0 {
		tDto.TotNanoseconds *= -1
	}

}

// ConvertToNegativeValues - Multiplies time component
// values by -1
func (tDto *TimeDto) ConvertToNegativeValues() {
	tDto.ConvertToAbsoluteValues()
	tDto.Years 					*= -1
	tDto.Months 				*= -1
	tDto.Weeks 					*= -1
	tDto.WeekDays 			*= -1
	tDto.DateDays 			*= -1
	tDto.Hours 					*= -1
	tDto.Minutes 				*= -1
	tDto.Seconds 				*= -1
	tDto.Milliseconds 	*= -1
	tDto.Microseconds 	*= -1
	tDto.Nanoseconds 		*= -1
	tDto.TotNanoseconds	*= -1
}

// Empty - returns all TimeDto data fields to their
// uninitialized or zero state.
func (tDto *TimeDto) Empty() {
	tDto.Years 					= 0
	tDto.Months 				= 0
	tDto.Weeks 					= 0
	tDto.WeekDays 			= 0
	tDto.DateDays 			= 0
	tDto.Hours 					= 0
	tDto.Minutes 				= 0
	tDto.Seconds 				= 0
	tDto.Milliseconds 	= 0
	tDto.Microseconds 	= 0
	tDto.Nanoseconds 		= 0
	tDto.TotNanoseconds = 0
}

// Equal - Compares the data fields of input parameter TimeDto, 'tDto2',
// to the data fields of the current TimeDto, 'tDto'. If all data fields
// are equal, this method returns 'true'. Otherwise, the method returns
// false.
func (tDto *TimeDto) Equal(tDto2 TimeDto) bool {

	if tDto.Years					!=  tDto2.Years 					||
		tDto.Months					!=  tDto2.Months					||
		tDto.Weeks					!=  tDto2.Weeks						||
		tDto.WeekDays 			!=  tDto2.WeekDays 				||
		tDto.DateDays 			!=  tDto2.DateDays 				||
		tDto.Hours					!=  tDto2.Hours						||
		tDto.Minutes 				!=  tDto2.Minutes					||
		tDto.Seconds    		!=  tDto2.Seconds					||
		tDto.Milliseconds 	!=  tDto2.Milliseconds 		||
		tDto.Microseconds 	!=  tDto2.Microseconds		||
		tDto.Nanoseconds  	!=  tDto2.Nanoseconds 		||
		tDto.TotNanoseconds	!=	tDto2.TotNanoseconds	{

		return false
	}

	return true
}

// GetDateTime - Analyzes the current TimeDto instance and computes
// an equivalent date time (time.Time). The calling function must
// pass in a valid time zone location.
//
// Input Parameter
// ===============
//
// timeZoneLocation	string	- time zone location must be designated as one of two values.
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
//
func (tDto *TimeDto) GetDateTime(timeZoneLocation string) (time.Time, error) {
	ePrefix := "TimeDto.GetDateTime() "

	err := tDto.IsValid()

	if err != nil {
		return time.Time{}, fmt.Errorf("Error returned by this tDto.IsValid(). Error='%v'", err.Error())
	}

	if len(timeZoneLocation) == 0 {
		return time.Time{}, errors.New(ePrefix + "Error: Input parameter 'timeZoneLocation' is EMPTY!")
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		timeZoneLocation = "Local"
	}

	loc, err := time.LoadLocation(timeZoneLocation)

	if err != nil {
		return time.Time{}, fmt.Errorf(ePrefix +
			"Error returned from time.LoadLocation(timeZoneLocation). " +
			"timeZoneLocation='%v'  Error='%v'", timeZoneLocation, err.Error())
	}

	dTime := time.Date(int(tDto.Years),
		time.Month(int(tDto.Months)),
		int(tDto.DateDays),
		int(tDto.Hours),
		int(tDto.Minutes),
		int(tDto.Seconds),
		int(tDto.TotNanoseconds),
		loc )

	return dTime, nil
}

// IsValid - Returns an error if the current tDto instance is invalid.
// Otherwise, if successful, this method returns 'nil'.
func (tDto *TimeDto) IsValid() error {

	ePrefix := "TimeDto.IsValid() "

	if tDto.Months > 31 {
		return fmt.Errorf(ePrefix + "Error: Months value is INVALID! tDto.Months='%v'", tDto.Months)
	}

	if tDto.Weeks > 4 {
		return fmt.Errorf(ePrefix + "Error: Weeks value is INVALID! tDto.Weeks='%v'", tDto.Weeks)

	}

	if tDto.WeekDays > 6 {
		return fmt.Errorf(ePrefix + "Error: WeekDays value is INVALID! tDto.WeekDays='%v'", tDto.WeekDays)
	}

	if tDto.DateDays > 31 {
		return fmt.Errorf(ePrefix + "Error: Total WeekDays value is INVALID! tDto.DateDays='%v'", tDto.DateDays)
	}

	if tDto.Hours > 24 {
		return fmt.Errorf(ePrefix + "Error: Hours value is INVALID! tDto.Hours='%v'", tDto.Hours)
	}

	if tDto.Minutes > 59 {
		return fmt.Errorf(ePrefix + "Error: Minutes value is INVALID! tDto.Minutes='%v'", tDto.Minutes)
	}

	if tDto.Seconds > 59 {
		return fmt.Errorf(ePrefix + "Error: Seconds value is INVALID! tDto.Seconds='%v'", tDto.Seconds)
	}

	if tDto.Milliseconds > int(MilliSecondsPerSecond - 1) {
		return fmt.Errorf(ePrefix + "Error: Milliseconds value is INVALID! tDto.Milliseconds='%v'", tDto.Milliseconds)
	}

	if tDto.Microseconds > int(MicroSecondsPerMilliSecond - 1) {
		return fmt.Errorf(ePrefix + "Error: Microseconds value is INVALID! tDto.Microseconds='%v'", tDto.Microseconds)
	}

	if tDto.Nanoseconds > int(NanoSecondsPerMicroSecond - 1) {
		return fmt.Errorf(ePrefix + "Error: Nanoseconds value is INVALID! tDto.Nanoseconds='%v'", tDto.Nanoseconds)
	}

	if tDto.TotNanoseconds > int(SecondNanoseconds - 1) {
		return fmt.Errorf(ePrefix + "Error: Total Nanoseconds value is INVALID! tDto.TotNanoseconds='%v'", tDto.TotNanoseconds)
	}

	return nil
}


// IsZero - Returns 'true' if all data fields in the current
// TimeDto instance are equal to zero or their uninitialized
// values.
func (tDto *TimeDto) IsZero() bool {

	if tDto.Years 					== 0 &&
			tDto.Months					== 0 &&
			tDto.Weeks					== 0 &&
			tDto.WeekDays				== 0 &&
			tDto.DateDays				== 0 &&
			tDto.Hours					== 0 &&
			tDto.Minutes				== 0 &&
			tDto.Seconds				== 0 &&
			tDto.Milliseconds		== 0 &&
			tDto.Microseconds		== 0 &&
			tDto.Nanoseconds		== 0 &&
			tDto.TotNanoseconds	== 0 {
				return true
	}

	return false
}

// New - Returns a new TimeDto instance based on time element
// input parameters.
//
func (tDto TimeDto) New(years, months, weeks, days, hours, minutes,
				seconds, milliseconds, microseconds,
						nanoseconds int) (TimeDto, error) {

	ePrefix := "TimeDto.New(...) "

	t2Dto := TimeDto{}

	err := t2Dto.SetTimeElements(years, months, weeks, days, hours, minutes,
							seconds, milliseconds, microseconds,
										nanoseconds)

	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix + "Error returned by t2Dto.SetTimeElements(...)  Error='%v'", err.Error())
	}

	err = t2Dto.IsValid()

	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix + "Error returned by t2Dto.IsValid()! Error='%v'", err.Error())
	}

	return t2Dto, nil
}

// NewFromDateTime - Creates and returns a new TimeDto based on
// a date time (time.Time) input parameter.
func (tDto TimeDto) NewFromDateTime(dateTime time.Time) (TimeDto, error) {
	
	ePrefix := "TimeDto.NewFromDateTime() "
	
	if dateTime.IsZero() {
		return TimeDto{}, errors.New(ePrefix + "Error: Input Parameter 'dateTime' has a ZERO Value!") 
	}
	
	t2Dto := TimeDto{}
	
	err := t2Dto.SetFromDateTime(dateTime)
	
	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix + "Error returned from t2Dto.SetFromDateTime(dateTime). Error='%v'", err.Error())
	}

	err = t2Dto.IsValid()

	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix + "Error returned by t2Dto.IsValid()! Error='%v'", err.Error())
	}

	return t2Dto, nil
}

// SetTimeElements - Sets the value of date fields for the current TimeDto instance
// based on time element input paramters.
//
func (tDto *TimeDto) SetTimeElements(years, months, weeks, days, hours, minutes,
seconds, milliseconds, microseconds,
nanoseconds int)  error {

	ePrefix := "TimeDto.SetTimeElements(...) "

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

		return fmt.Errorf(ePrefix + "Error: All input parameters (years, months, weeks, days etc.) are ZERO Value!")
	}

	if months > 12 {
		return fmt.Errorf(ePrefix + "Error: Input parameter 'months' is INVALID! months='%v'", months)
	}

	weeksAndDays := (weeks * 7) + days

	totalNanoSecs := milliseconds * int(time.Millisecond)
	totalNanoSecs += microseconds * int(time.Microsecond)
	totalNanoSecs += nanoseconds

	totalSeconds := hours * 3600
	totalSeconds += minutes * 60
	totalSeconds += seconds

	tDto.Empty()

	tDto.Years = years
	tDto.Months = months
	err := tDto.allocateWeeksAndDays(weeksAndDays)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.allocateWeeksAndDays(weeksAndDays). Error='%v'", err.Error())
	}

	err = tDto.allocateSeconds(totalSeconds)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.allocateSeconds(totalSeconds). Error='%v'", err.Error())
	}


	err = tDto.allocateTotalNanoseconds(totalNanoSecs)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.allocateTotalNanoseconds(totalNanoSecs). Error='%v'", err.Error())
	}

	err = tDto.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by t2Dto.IsValid()! Error='%v'", err.Error())
	}

	return nil
}

// SetFromDateTime - Sets the current TimeDto instance to new
// data field values based on input parameter 'dateTime' (time.Time)
func (tDto *TimeDto) SetFromDateTime(dateTime time.Time) error {
	ePrefix := "TimeDto.SetFromDateTime() "

	if dateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input Parameter 'dateTime' has a ZERO Value!")
	}

	tDto.Empty()
	
	tDto.Years = dateTime.Year()
	tDto.Months = int(dateTime.Month())
	err := tDto.allocateWeeksAndDays(dateTime.Day())

	if err != nil {
		return fmt.Errorf(ePrefix + "Error= '%v'", err.Error())
	}

	tDto.Hours = dateTime.Hour()
	tDto.Minutes = dateTime.Minute()
	tDto.Seconds = dateTime.Second()

	err = tDto.allocateTotalNanoseconds(dateTime.Nanosecond())

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by tDto.allocateTotalNanoseconds(int64(dateTime.Nanosecond())). " +
			"Error='%v'", err.Error())
	}

	err = tDto.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.IsValid()! Error='%v'", err.Error())
	}

	return nil
	
}

// allocateWeeksAndDays - This method receives a total number of
// days and allocates those days to Weeks and WeekDays. The result
// is stored in the Weeks and WeekDays data fields of the current
// TimeDto instance.
func (tDto *TimeDto) allocateWeeksAndDays(totalDays int) error {
	ePrefix := "TimeDto.allocateWeeksAndDays() "
	tDto.Weeks 		= 0
	tDto.WeekDays = 0
	tDto.DateDays = totalDays

	if totalDays >= 7 {
		tDto.Weeks = totalDays / 7
		totalDays -= tDto.Weeks * 7
	}

	if tDto.Weeks > 4 {
		return fmt.Errorf(ePrefix + "Error: The number of weeks is INVALID. weeks='%v'", tDto.Weeks)
	}

	tDto.WeekDays = totalDays

	if tDto.WeekDays > 31 {
		return fmt.Errorf(ePrefix + "Error: The number of days is INVALID. days='%v'", tDto.WeekDays)
	}

	return nil
}


// allocateSeconds - Receives totalSeconds and proceeds to
// allocate Hours, Minutes and Seconds. The result is stored
// the Hours, Minutes and Seconds data fields of the current
// TimeDto instance.
func (tDto *TimeDto) allocateSeconds(totalSeconds int) error {

	ePrefix := "TimeDto.allocateSeconds() "

	tDto.Hours = 0
	tDto.Minutes = 0
	tDto.Seconds = 0

	if totalSeconds >= 3600 {
		tDto.Hours = totalSeconds / 3600
		totalSeconds -= tDto.Hours * 3600
	}

	if tDto.Hours > 24 {
		return fmt.Errorf(ePrefix + "Error: The number of Hours is INVALID. Hours='%v'", tDto.Hours)
	}


	if totalSeconds >= 60 {
		tDto.Minutes = totalSeconds / 60
		totalSeconds -= tDto.Minutes * 60
	}

	if tDto.Minutes > 59 {
		return fmt.Errorf(ePrefix + "Error: The number of Minutes is INVALID. Minutes='%v'", tDto.Minutes)
	}

	tDto.Seconds = totalSeconds

	if tDto.Seconds > 59 {
		return fmt.Errorf(ePrefix + "Error: The number of Seconds is INVALID. Secondes='%v'", tDto.Seconds)
	}

	return nil
}

// allocateTotalNanoseconds - Allocates total nanoseconds to current
// TimeDto instance data fields: milliseconds, microseconds and
// nanoseconds.
func (tDto *TimeDto) allocateTotalNanoseconds(totalNanoSeconds int) error {
	ePrefix := "TimeDto.allocateTotalNanoseconds() "

	tDto.Milliseconds = 0
	tDto.Microseconds = 0
	tDto.Nanoseconds = 0
	tDto.TotNanoseconds = totalNanoSeconds

	if totalNanoSeconds >= int(time.Millisecond) {
		tDto.Milliseconds = totalNanoSeconds / int(time.Millisecond)
		totalNanoSeconds -= tDto.Milliseconds * int(time.Millisecond)
	}

	if tDto.Milliseconds > int(MilliSecondsPerSecond - 1) {
		return fmt.Errorf(ePrefix + "Error: Milliseconds is INVALID. Milliseconds='%v'", tDto.Milliseconds)
	}

	if totalNanoSeconds >= int(time.Microsecond) {
		tDto.Microseconds = totalNanoSeconds / int(time.Microsecond)
		totalNanoSeconds -= tDto.Microseconds * int(time.Microsecond)
	}

	if tDto.Microseconds > int(MicroSecondsPerMilliSecond - 1) {
		return fmt.Errorf(ePrefix + "Error: Microseconds is INVALID. Microseconds='%v'", tDto.Microseconds)
	}

	tDto.Nanoseconds = totalNanoSeconds

	if tDto.Nanoseconds > int(NanoSecondsPerMicroSecond - 1) {
		return fmt.Errorf(ePrefix + "Error: Nanoseconds is INVALID. Nanoseconds= %v", tDto.Nanoseconds)
	}

	return nil
}

