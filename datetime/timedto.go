package datetime

import (
	"time"
	"fmt"
	"errors"
	"strings"
)

// TimeDto - used for transmitting time elements.
// TimeDto data fields are designed to store one of two
// time components:
//		(1)	A specific point in time (date time).
//									or
//		(2) Incremental time which is useful in adding or subtracting
//					time values. Note that this structure does not track
//					time location or time zone. For a fully supported date time
//					structure, review the DateTzDto located in source file 'datetzdto.go'
//
type TimeDto struct {
	Years          				int		//	Number of Years
	Months         				int 	//	Number of Months
	Weeks               	int		//	Number of Weeks
	WeekDays            	int		//	Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
	DateDays            	int		//	Total Number of Days. Weeks x 7 plus WeekDays
	Hours               	int		//	Number of Hours.
	Minutes             	int		//	Number of Minutes
	Seconds             	int		//	Number of Seconds
	Milliseconds        	int		//	Number of Milliseconds
	Microseconds        	int		//	Number of Microseconds
	Nanoseconds         	int		//	Remaining Nanoseconds after Milliseconds & Microseconds
	TotSubSecNanoseconds	int		//	Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
															// 		plus remaining Nanoseconds
	TotTimeNanoseconds		int64	//	Total Number of equivalent Nanoseconds for Hours + Minutes
															//			+ Seconds + Milliseconds + Nanoseconds
}

// AddTimeDto - Adds time to the current TimeDto. The amount of time added
// is provided by the input parameter 't2Dto' of type TimeDto.
//
// Date time math uses timezone UTC.
//
//	Input Parameters
//	================
//
//	t2Dto						TimeDto	- The amount of time to be added to the current TimeDto
//															data fields.
//
//
func (tDto *TimeDto) AddTimeDto(t2Dto TimeDto) error {

	ePrefix := "TimeDto.AddTimeDto() "

	loc, err := time.LoadLocation(TzIanaUTC)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by time.LoadLocation(timeZoneLocation). TzIanaUTC='%v'  Error='%v'", TzIanaUTC, err.Error())
	}

	dTime := time.Date(tDto.Years,
		time.Month(tDto.Months),
		tDto.DateDays,
		tDto.Hours,
		tDto.Minutes,
		tDto.Seconds,
		tDto.TotSubSecNanoseconds,
		loc )

	d2Time := dTime.AddDate(t2Dto.Years, t2Dto.Months, 0)

	dur := int64(t2Dto.DateDays) * DayNanoSeconds
	dur += int64(t2Dto.Hours) * HourNanoSeconds
	dur += int64(t2Dto.Minutes) * MinuteNanoSeconds
	dur += int64(t2Dto.Seconds) * SecondNanoseconds
	dur += int64(t2Dto.Milliseconds) * MilliSecondNanoseconds
	dur += int64(t2Dto.Microseconds) * MicroSecondNanoseconds
	dur += int64(t2Dto.Nanoseconds)

	dateTime := d2Time.Add(time.Duration(dur))


	tDto.Empty()
	tDto.Years = dateTime.Year()
	tDto.Months = int(dateTime.Month())

	err = tDto.allocateWeeksAndDays(dateTime.Day())

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.allocateWeeksAndDays(dateTime.Day()). Error='%v'", err.Error())
	}

	tDto.Hours = dateTime.Hour()
	tDto.Minutes = dateTime.Minute()
	tDto.Seconds = dateTime.Second()

	err = tDto.allocateTotalNanoseconds(dateTime.Nanosecond())

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returne by tDto.allocateTotalNanoseconds(dateTime.Nanosecond()). Error='%v'", err.Error())
	}

	return nil
}


// CopyOut - Creates a new TimeDto instance
// which precisely duplicates the current TimeDto
// instance, and returns it to the calling function.
func (tDto *TimeDto) CopyOut() TimeDto {

	t2Dto := TimeDto{}

	t2Dto.Years 								=  tDto.Years
	t2Dto.Months       					=  tDto.Months
	t2Dto.Weeks        					=  tDto.Weeks
	t2Dto.WeekDays 							=  tDto.WeekDays
	t2Dto.DateDays 							=  tDto.DateDays
	t2Dto.Hours        					=  tDto.Hours
	t2Dto.Minutes      					=  tDto.Minutes
	t2Dto.Seconds      					=  tDto.Seconds
	t2Dto.Milliseconds 					=  tDto.Milliseconds
	t2Dto.Microseconds 					=  tDto.Microseconds
	t2Dto.Nanoseconds  					=  tDto.Nanoseconds
	t2Dto.TotSubSecNanoseconds 	= tDto.TotSubSecNanoseconds
	t2Dto.TotTimeNanoseconds 		= tDto.TotTimeNanoseconds

	return t2Dto
}

// CopyIn - Receives a TimeDto input parameter, 'tDto2'
// and proceeds to copy all 'tDto2' data fields into
// the current TimeDto data fields. When this method
// completes, 'tDto' will be equivalent to 'tDto2'.
func (tDto *TimeDto) CopyIn(t2Dto TimeDto) {

	tDto.Empty()

	tDto.Years 								=  t2Dto.Years
	tDto.Months       				=  t2Dto.Months
	tDto.Weeks        				=  t2Dto.Weeks
	tDto.WeekDays 						=  t2Dto.WeekDays
	tDto.DateDays							=  t2Dto.DateDays
	tDto.Hours        				=  t2Dto.Hours
	tDto.Minutes      				=  t2Dto.Minutes
	tDto.Seconds      				=  t2Dto.Seconds
	tDto.Milliseconds 				=  t2Dto.Milliseconds
	tDto.Microseconds 				=  t2Dto.Microseconds
	tDto.Nanoseconds  				=  t2Dto.Nanoseconds
	tDto.TotSubSecNanoseconds =  t2Dto.TotSubSecNanoseconds
	tDto.TotTimeNanoseconds 	= 	t2Dto.TotTimeNanoseconds


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

	if tDto.TotSubSecNanoseconds < 0 {
		tDto.TotSubSecNanoseconds *= -1
	}

	if tDto.TotTimeNanoseconds < 0 {
		tDto.TotTimeNanoseconds *= -1
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
	tDto.TotSubSecNanoseconds *= -1
	tDto.TotTimeNanoseconds 	*= -1
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
	tDto.TotSubSecNanoseconds = 0
	tDto.TotTimeNanoseconds = 0
}

// Equal - Compares the data fields of input parameter TimeDto, 'tDto2',
// to the data fields of the current TimeDto, 'tDto'. If all data fields
// are equal, this method returns 'true'. Otherwise, the method returns
// false.
func (tDto *TimeDto) Equal(t2Dto TimeDto) bool {

	if tDto.Years								!=  t2Dto.Years 								||
		tDto.Months								!=  t2Dto.Months								||
		tDto.Weeks								!=  t2Dto.Weeks									||
		tDto.WeekDays 						!=  t2Dto.WeekDays 							||
		tDto.DateDays 						!=  t2Dto.DateDays 							||
		tDto.Hours								!=  t2Dto.Hours									||
		tDto.Minutes 							!=  t2Dto.Minutes								||
		tDto.Seconds    					!=  t2Dto.Seconds								||
		tDto.Milliseconds 				!=  t2Dto.Milliseconds 					||
		tDto.Microseconds 				!=  t2Dto.Microseconds					||
		tDto.Nanoseconds  				!=  t2Dto.Nanoseconds 					||
		tDto.TotSubSecNanoseconds !=	t2Dto.TotSubSecNanoseconds	||
		tDto.TotTimeNanoseconds 	!= t2Dto.TotTimeNanoseconds			{

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
//																	"Etc/UTC" = ZULU, GMT or UTC
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
func (tDto *TimeDto) GetDateTime(timeZoneLocation string) (time.Time, error) {
	ePrefix := "TimeDto.GetDateTime() "

	tzLoc := tDto.preProcessTimeZoneLocation(timeZoneLocation)

	loc, err := time.LoadLocation(tzLoc)

	if err != nil {
		return time.Time{}, fmt.Errorf(ePrefix +
			"Error returned from time.LoadLocation(timeZoneLocation). " +
			"timeZoneLocation='%v'  Error='%v'", timeZoneLocation, err.Error())
	}

	dTime := time.Date(tDto.Years,
		time.Month(tDto.Months),
		tDto.DateDays,
		tDto.Hours,
		tDto.Minutes,
		tDto.Seconds,
		tDto.TotSubSecNanoseconds,
		loc )

	return dTime, nil
}

// IsEmpty - Returns 'true' if all data fields in the current
// TimeDto instance are equal to zero or equal to their
// uninitialized values.
func (tDto *TimeDto) IsEmpty() bool {

	if tDto.Years 							== 0 &&
		tDto.Months								== 0 &&
		tDto.Weeks								== 0 &&
		tDto.WeekDays							== 0 &&
		tDto.DateDays							== 0 &&
		tDto.Hours								== 0 &&
		tDto.Minutes							== 0 &&
		tDto.Seconds							== 0 &&
		tDto.Milliseconds					== 0 &&
		tDto.Microseconds					== 0 &&
		tDto.Nanoseconds					== 0 &&
		tDto.TotSubSecNanoseconds == 0 &&
		tDto.TotTimeNanoseconds 	== 0	{
		return true
	}

	return false
}

// IsValidDateTime - Returns an error if the current tDto instance is invalid.
// Otherwise, if successful, this method returns 'nil'.
func (tDto *TimeDto) IsValidDateTime() error {

	ePrefix := "TimeDto.IsValidDateTime() "

	if tDto.Months < 1 || tDto.Months > 12 {
		return fmt.Errorf(ePrefix + "Error: Months value is INVALID! tDto.Months='%v'", tDto.Months)
	}

	if tDto.Weeks < 0 || tDto.Weeks > 4 {
		return fmt.Errorf(ePrefix + "Error: Weeks value is INVALID! tDto.Weeks='%v'", tDto.Weeks)
	}

	if tDto.WeekDays < 0 || tDto.WeekDays > 6 {
		return fmt.Errorf(ePrefix + "Error: WeekDays value is INVALID! tDto.WeekDays='%v'", tDto.WeekDays)
	}

	if tDto.DateDays < 0 || tDto.DateDays > 31 {
		return fmt.Errorf(ePrefix + "Error: Total WeekDays value is INVALID! tDto.DateDays='%v'", tDto.DateDays)
	}

	if tDto.Hours < 0 ||tDto.Hours > 24 {
		return fmt.Errorf(ePrefix + "Error: Hours value is INVALID! tDto.Hours='%v'", tDto.Hours)
	}

	if tDto.Minutes < 0 || tDto.Minutes > 59 {
		return fmt.Errorf(ePrefix + "Error: Minutes value is INVALID! tDto.Minutes='%v'", tDto.Minutes)
	}

	if tDto.Seconds < 0 || tDto.Seconds > 59 {
		return fmt.Errorf(ePrefix + "Error: Seconds value is INVALID! tDto.Seconds='%v'", tDto.Seconds)
	}

	if tDto.Milliseconds < 0 || tDto.Milliseconds > int(MilliSecondsPerSecond - 1) {
		return fmt.Errorf(ePrefix + "Error: Milliseconds value is INVALID! tDto.Milliseconds='%v'", tDto.Milliseconds)
	}

	if tDto.Microseconds < 0 || tDto.Microseconds > int(MicroSecondsPerMilliSecond - 1) {
		return fmt.Errorf(ePrefix + "Error: Microseconds value is INVALID! tDto.Microseconds='%v'", tDto.Microseconds)
	}

	if tDto.Nanoseconds < 0 || tDto.Nanoseconds > int(NanoSecondsPerMicroSecond - 1) {
		return fmt.Errorf(ePrefix + "Error: Nanoseconds value is INVALID! tDto.Nanoseconds='%v'", tDto.Nanoseconds)
	}

	if tDto.TotSubSecNanoseconds < 0 ||  tDto.TotSubSecNanoseconds > int(SecondNanoseconds - 1) {
		return fmt.Errorf(ePrefix + "Error: Total Nanoseconds value is INVALID! tDto.TotSubSecNanoseconds='%v'", tDto.TotSubSecNanoseconds)
	}

	if tDto.TotTimeNanoseconds < 0 {
		return fmt.Errorf(ePrefix + "Error: Total Time Nanoseconds value is INVALID! tDto.TotTimeNanoseconds='%v'", tDto.TotTimeNanoseconds)
	}

	return nil
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


	return t2Dto, nil
}

// NewTimeElements - Creates and returns a new TimeDto using basic
// time components as input parameters
func (tDto TimeDto) NewTimeElements(years, months, days, hours, minutes,
																			seconds, nanoseconds int) (TimeDto, error) {


	ePrefix := "TimeDto.NewTimeElements(...) "

	t2Dto := TimeDto{}

	err := t2Dto.SetTimeElements(years, months, 0, days, hours, minutes,
																	seconds, 0, 0,	nanoseconds)

	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix + "Error returned by t2Dto.SetTimeElements(...)  Error='%v'", err.Error())
	}


	return t2Dto, nil

}

// NewAddTimeDtos - Creates a new TimeDto which adds the values of the two TimeDto's
// passed as input parameters. Date time math is performed using time zone 'UTC'.
//
//	Input Parameters
//	================
//
//	t1Dto						TimeDto	- The value of this TimeDto will be added to the second
//															input parameter to create and return a summary TimeDto.
//
//	t2Dto						TimeDto	- The value of this TimeDto will be added to the first
//															input parameter to create and return a summary TimeDto.
//
//
func (tDto TimeDto) NewAddTimeDtos(t1Dto, t2Dto TimeDto) (TimeDto, error) {
	ePrefix := "TimeDto.NewAddTimeDtos(...) "

	tOutDto := t1Dto.CopyOut()

	err := tOutDto.AddTimeDto(t2Dto)

	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix + "Error returned by tOutDto.AddTimeDto(t2Dto, timeZoneLocation). Error='%v'", err.Error())
	}

	return tOutDto, nil
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

	err = t2Dto.IsValidDateTime()

	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix + "Error returned by t2Dto.IsValidDateTime()! Error='%v'", err.Error())
	}

	return t2Dto, nil
}

// NewFromDateTzDto - Creates and returns a new TimeDto instance based on
// a DateTzDto input parameter.
func (tDto TimeDto) NewFromDateTzDto(dTzDto DateTzDto) (TimeDto, error) {
	ePrefix := "TimeDto.NewFromDateTzDto() "

	tDto2 := TimeDto{}

	err := tDto2.SetFromDateTzDto(dTzDto)

	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix + "Error returned by tDto2.SetFromDateTzDto(dTzDto). Error='%v'", err.Error())
	}

	return tDto2, nil
}

// NormalizeTimeElements - Surveys the time elements of the current 
// TimeDto and normalizes time values. Example: Minutes > 59, Seconds > 59
// Hours > 23, Minures > 59, Seconds > 59 etc.
func (tDto *TimeDto) NormalizeTimeElements() error {

	ePrefix := "TimeDto.NormalizeTimeElements() "

	carry := 0

	sign := 1

	if tDto.Nanoseconds < 0 {
		sign = -1
		tDto.Nanoseconds *= sign
	}

	if tDto.Nanoseconds > 999 {
		carry = tDto.Nanoseconds / 1000
		tDto.Nanoseconds = tDto.Nanoseconds - (carry * 1000)
	}

	tDto.Nanoseconds *= sign

	sign = 1

	if tDto.Microseconds < 0 {
		sign = -1
		tDto.Microseconds *= sign
	}

	tDto.Microseconds += carry
	carry = 0

	if tDto.Microseconds > 999 {
		carry = tDto.Microseconds / 1000
		tDto.Microseconds = tDto.Microseconds - (carry * 1000)
	}

	tDto.Microseconds *= sign

	sign = 1

	if tDto.Milliseconds < 0 {
		sign = -1
		tDto.Milliseconds *= sign
	}

	tDto.Milliseconds += carry
	carry = 0

	if tDto.Milliseconds > 999 {
		carry = tDto.Milliseconds / 1000
		tDto.Milliseconds = tDto.Milliseconds - (carry * 1000)
	}

	tDto.Milliseconds *= sign


	sign = 1

	if tDto.Seconds < 0 {
		sign = -1
		tDto.Seconds *= sign
	}

	tDto.Seconds += carry
	carry = 0

	if tDto.Seconds > 59 {
		carry = tDto.Seconds / 60
		tDto.Seconds = tDto.Seconds - (carry * 60)
	}

	tDto.Seconds *= sign

	sign = 1

	if tDto.Minutes < 0 {
		sign = -1
		tDto.Minutes *= sign
	}

	tDto.Minutes += carry
	carry = 0

	if tDto.Minutes > 59 {
		carry = tDto.Minutes / 60
		tDto.Minutes = tDto.Minutes - (carry * 60)
	}

	tDto.Minutes *= sign

	sign = 1

	if tDto.Hours < 0 {
		sign = -1
		tDto.Hours *= sign
	}

	tDto.Hours += carry
	carry = 0

	if tDto.Hours > 23 {
		carry = tDto.Hours / 24
		tDto.Hours = tDto.Hours - (carry * 24)
	}

	tDto.Hours *= sign

	if tDto.DateDays == 0 {
		tDto.DateDays = (tDto.Weeks * 7 ) + tDto.WeekDays
	}

	sign = 1

	if tDto.DateDays < 0 {
		sign = -1
		tDto.DateDays *= sign
	}

	tDto.DateDays += carry

	tDto.DateDays *= sign

	sign = 1

	if tDto.Months < 0 {
		sign = -1
		tDto.Months *= sign
	}

	if tDto.Months > 12 {
		carry = tDto.Months / 12
		tDto.Months = tDto.Months - (carry * 12)
	}
	
	tDto.Months *= sign
	
	sign = 1

	if tDto.Years < 0 {
		sign = -1
		tDto.Years *= sign
	}

	tDto.Years += carry

	tDto.Years *= sign


	totSubNanoSecs :=  tDto.Nanoseconds
	totSubNanoSecs +=  int(int64(tDto.Microseconds) * MicroSecondNanoseconds)
	totSubNanoSecs +=  int(int64(tDto.Milliseconds) * MilliSecondNanoseconds)

	err := tDto.allocateTotalNanoseconds(totSubNanoSecs)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.allocateTotalNanoseconds(totSubNanoSecs) " +
			"totSubNanoSecs='%v' Error='%v",
			totSubNanoSecs, err.Error())
	}

	totSeconds := tDto.Hours * 3600
	totSeconds += tDto.Minutes * 60
	totSeconds += tDto.Seconds

	err = tDto.allocateSeconds(totSeconds)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.allocateSeconds(totSeconds) " +
			"totSeconds='%v' Error='%v",
			totSeconds, err.Error())
	}

	err = tDto.allocateWeeksAndDays(tDto.DateDays)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.allocateWeeksAndDays(tDto.DateDays) " +
			"tDto.DateDays='%v' Error='%v",
			tDto.DateDays, err.Error())
	}

	return nil
}

// SetTimeElements - Sets the value of date fields for the current TimeDto instance
// based on time element input parameters.
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

	t1Dto := TimeDto{}

	t1Dto.Years = years
	t1Dto.Months = months
	t1Dto.DateDays = (weeks * 7 ) + days

	err := t1Dto.allocateWeeksAndDays(t1Dto.DateDays)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by err := t1Dto.allocateWeeksAndDays(t1Dto.DateDays)")
	}

	t1Dto.Hours = hours
	t1Dto.Minutes = minutes
	t1Dto.Seconds = seconds
	t1Dto.Milliseconds = milliseconds
	t1Dto.Microseconds = microseconds
	t1Dto.Nanoseconds = nanoseconds

	t1Dto.NormalizeTimeElements()



	if err:=t1Dto.IsValidDateTime(); err !=nil {
		// This is an incremental TimeDto. Some
		// values may be negative.

		tDto.Empty()
		tDto.CopyIn(t1Dto)
		return nil

	}

	loc, err := time.LoadLocation(TzIanaUTC)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by time.LoadLocation(TzIanaUTC). Error='%v'", err.Error())
	}

	dateTime := time.Date(t1Dto.Years, time.Month(t1Dto.Months), t1Dto.DateDays,
						t1Dto.Hours, t1Dto.Minutes, t1Dto.Seconds, t1Dto.TotSubSecNanoseconds, loc)

	tDto.Empty()
	tDto.Years = dateTime.Year()
	tDto.Months = int(dateTime.Month())

	err = tDto.allocateWeeksAndDays(dateTime.Day())

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.allocateWeeksAndDays(dateTime.Day()). Error='%v'", err.Error())
	}

	tDto.Hours = dateTime.Hour()
	tDto.Minutes = dateTime.Minute()
	tDto.Seconds = dateTime.Second()

	err = tDto.allocateTotalNanoseconds(dateTime.Nanosecond())

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.allocateTotalNanoseconds(dateTime.Nanosecond()). Error='%v'", err.Error())
	}

	err = tDto.IsValidDateTime()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.IsValidDateTime()! Error='%v'", err.Error())
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
		return fmt.Errorf(ePrefix + "tDto.allocateWeeksAndDays(dTzDto.DateTime.Day()). Error= '%v'", err.Error())
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

	err = tDto.IsValidDateTime()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.IsValidDateTime()! Error='%v'", err.Error())
	}

	return nil
	
}

// SetFromDateTzDto - Sets the data field values of the current TimeDto
// instance based on a DateTzDto input parameter.
func (tDto *TimeDto) SetFromDateTzDto(dTzDto DateTzDto) error {

	ePrefix := "TimeDto.SetFromDateTzDto() "

	if dTzDto.IsEmpty() {
		return errors.New(ePrefix + "Error: Input parameter 'dTzDto' (DateTzDto) is EMPTY!")
	}

	if err:= dTzDto.IsValid(); err!=nil {
		return fmt.Errorf(ePrefix + "Error: Input parameter 'dTzDto' (DateTzDto) is INVALID! Error='%v'", err.Error())
	}

	tDto.Empty()

	tDto.Years = dTzDto.DateTime.Year()
	tDto.Months = int(dTzDto.DateTime.Month())
	err := tDto.allocateWeeksAndDays(dTzDto.DateTime.Day())

	if err != nil {
		return fmt.Errorf(ePrefix + "tDto.allocateWeeksAndDays(dTzDto.DateTime.Day()). Error= '%v'", err.Error())
	}

	tDto.Hours = dTzDto.DateTime.Hour()
	tDto.Minutes = dTzDto.DateTime.Minute()
	tDto.Seconds = dTzDto.DateTime.Second()

	if err := tDto.allocateTotalNanoseconds(dTzDto.DateTime.Nanosecond()); err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by tDto.allocateTotalNanoseconds(int64(dateTime.Nanosecond())). " +
			"Error='%v'", err.Error())
	}

	if err := tDto.IsValidDateTime(); err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDto.IsValidDateTime()! Error='%v'", err.Error())
	}

	return nil
}

// allocateWeeksAndDays - This method receives a total number of
// days and allocates those days to Weeks and WeekDays. The result
// is stored in the Weeks and WeekDays data fields of the current
// TimeDto instance.
func (tDto *TimeDto) allocateWeeksAndDays(totalDays int) error {

	sign := 1

	if totalDays < 0 {
		sign = -1
		totalDays *= -1
	}

	tDto.Weeks 		= 0
	tDto.WeekDays = 0
	tDto.DateDays = totalDays

	if totalDays >= 7 {
		tDto.Weeks = totalDays / 7
		totalDays -= tDto.Weeks * 7
	}

	tDto.WeekDays = totalDays

	if sign == -1 {
		tDto.Weeks 		*= sign
		tDto.WeekDays *= sign
		tDto.DateDays *= sign
	}


	return nil
}


// allocateSeconds - Receives totalSeconds and proceeds to
// allocate Hours, Minutes and Seconds. The result is stored
// the Hours, Minutes and Seconds data fields of the current
// TimeDto instance.
func (tDto *TimeDto) allocateSeconds(totalSeconds int) error {

	sign := 1

	if totalSeconds < 0 {
		sign = -1
		totalSeconds *= -1
	}

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

	if sign == -1 {

		tDto.Hours 		*= sign
		tDto.Minutes 	*= sign
		tDto.Seconds 	*= sign

	}


	return nil
}

// allocateTotalNanoseconds - Allocates total nanoseconds to current
// TimeDto instance data fields: milliseconds, microseconds and
// nanoseconds.
//
// In addition, this method calculates TimeDto.TotTimeNanoseconds which
// is the sum of hours, minutes, seconds, milliseconds, microseconds and
// nanoseconds. Before calling this method, TimeDto Hours, Minutes and
// Seconds must be properly initialized.
//
func (tDto *TimeDto) allocateTotalNanoseconds(totalNanoSeconds int) error {

	sign := 1

	if totalNanoSeconds < 0 {
		sign = -1
		totalNanoSeconds *= -1
	}

	tDto.Milliseconds = 0
	tDto.Microseconds = 0
	tDto.Nanoseconds = 0
	tDto.TotSubSecNanoseconds = totalNanoSeconds
	tDto.TotTimeNanoseconds = int64(totalNanoSeconds)

	if totalNanoSeconds >= int(time.Millisecond) {
		tDto.Milliseconds = totalNanoSeconds / int(time.Millisecond)
		totalNanoSeconds -= tDto.Milliseconds * int(time.Millisecond)
	}


	if totalNanoSeconds >= int(time.Microsecond) {
		tDto.Microseconds = totalNanoSeconds / int(time.Microsecond)
		totalNanoSeconds -= tDto.Microseconds * int(time.Microsecond)
	}

	tDto.Nanoseconds = totalNanoSeconds

	// calculate total time nanoseconds
	tDto.TotTimeNanoseconds += int64(time.Hour) * int64(tDto.Hours)
	tDto.TotTimeNanoseconds += int64(time.Minute) * int64(tDto.Minutes)
	tDto.TotTimeNanoseconds += int64(time.Second) * int64(tDto.Seconds)

	if sign == -1 {

		tDto.Milliseconds 				*= sign
		tDto.Microseconds 				*= sign
		tDto.Nanoseconds 					*= sign
		tDto.TotSubSecNanoseconds *= sign
		tDto.TotTimeNanoseconds		*= int64(sign)
	}


	return nil
}

func (tDto *TimeDto) preProcessTimeZoneLocation(timeZoneLocation string) string {

	if len(timeZoneLocation) == 0 {
		return TzIanaUTC
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		return "Local"
	}

	return timeZoneLocation
}

