package datetime

import (
	"time"
	"errors"
	"fmt"
	"strings"
)

// TimeDurationDto -
type TimeDurationDto struct {
	StartTimeDateTz      DateTzDto
	EndTimeDateTz        DateTzDto
	TimeDuration         time.Duration
	Years                int64
	YearsNanosecs        int64
	Months               int64
	MonthsNanosecs       int64
	Weeks                int64
	WeeksNanosecs        int64
	WeekDays             int64
	WeekDaysNanosecs     int64
	DateDays             int64
	DateDaysNanosecs     int64
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
	TotSubSecNanoseconds int64

}

// CopyIn - Receives a TimeDurationDto as an input parameters
// and proceeds to set all data fields of the current TimeDurationDto
// equal to the incoming TimeDurationDto.
//
// When this method completes, the current TimeDurationDto will
// equal in all respects to the incoming TimeDurationDto.
func (tDur *TimeDurationDto) CopyIn(t2Dur TimeDurationDto) {
	
	tDur.Empty()

	tDur.StartTimeDateTz 				= t2Dur.StartTimeDateTz.CopyOut()
	tDur.EndTimeDateTz 					=	t2Dur.EndTimeDateTz.CopyOut()
	tDur.TimeDuration     			= t2Dur.TimeDuration
	tDur.Years									= t2Dur.Years
	tDur.YearsNanosecs    			= t2Dur.YearsNanosecs
	tDur.Months           			= t2Dur.Months
	tDur.MonthsNanosecs   			= t2Dur.MonthsNanosecs
	tDur.Weeks            			= t2Dur.Weeks
	tDur.WeeksNanosecs    			= t2Dur.WeeksNanosecs
	tDur.WeekDays								= t2Dur.WeekDays
	tDur.WeekDaysNanosecs				= t2Dur.WeekDaysNanosecs
	tDur.DateDays								= t2Dur.DateDays
	tDur.DateDaysNanosecs				= t2Dur.DateDaysNanosecs
	tDur.Hours									= t2Dur.Hours
	tDur.HoursNanosecs					= t2Dur.HoursNanosecs
	tDur.Minutes								= t2Dur.Minutes
	tDur.MinutesNanosecs				= t2Dur.MinutesNanosecs
	tDur.Seconds								= t2Dur.Seconds
	tDur.SecondsNanosecs				= t2Dur.SecondsNanosecs
	tDur.Milliseconds						= t2Dur.Milliseconds
	tDur.MillisecondsNanosecs		= t2Dur.MillisecondsNanosecs
	tDur.Microseconds						= t2Dur.Microseconds
	tDur.MicrosecondsNanosecs 	= t2Dur.MicrosecondsNanosecs
	tDur.Nanoseconds						= t2Dur.MillisecondsNanosecs
	tDur.TotSubSecNanoseconds 	= t2Dur.TotSubSecNanoseconds

}

// CopyOut - Returns a deep copy of the current 
// TimeDurationDto instance.
func (tDur *TimeDurationDto) CopyOut() TimeDurationDto {

	t2Dur := TimeDurationDto{}
	
	t2Dur.StartTimeDateTz 			= tDur.StartTimeDateTz.CopyOut()
	t2Dur.EndTimeDateTz 				=	tDur.EndTimeDateTz.CopyOut()
	t2Dur.TimeDuration     			= tDur.TimeDuration
	t2Dur.Years									= tDur.Years
	t2Dur.YearsNanosecs    			= tDur.YearsNanosecs
	t2Dur.Months           			= tDur.Months
	t2Dur.MonthsNanosecs   			= tDur.MonthsNanosecs
	t2Dur.Weeks            			= tDur.Weeks
	t2Dur.WeeksNanosecs    			= tDur.WeeksNanosecs
	t2Dur.WeekDays							= tDur.WeekDays
	t2Dur.WeekDaysNanosecs			= tDur.WeekDaysNanosecs
	t2Dur.DateDays							= tDur.DateDays
	t2Dur.DateDaysNanosecs			= tDur.DateDaysNanosecs
	t2Dur.Hours									= tDur.Hours
	t2Dur.HoursNanosecs					= tDur.HoursNanosecs
	t2Dur.Minutes								= tDur.Minutes
	t2Dur.MinutesNanosecs				= tDur.MinutesNanosecs
	t2Dur.Seconds								= tDur.Seconds
	t2Dur.SecondsNanosecs				= tDur.SecondsNanosecs
	t2Dur.Milliseconds					= tDur.Milliseconds
	t2Dur.MillisecondsNanosecs	= tDur.MillisecondsNanosecs
	t2Dur.Microseconds					= tDur.Microseconds
	t2Dur.MicrosecondsNanosecs 	= tDur.MicrosecondsNanosecs
	t2Dur.Nanoseconds						= tDur.MillisecondsNanosecs
	t2Dur.TotSubSecNanoseconds 	= tDur.TotSubSecNanoseconds
	
	return t2Dur
}

// Empty - Resets all of the current TimeDurationDto
// data fields to their zero or uninitialized values.
func (tDur *TimeDurationDto) Empty() {
	tDur.StartTimeDateTz 			= DateTzDto{}
	tDur.EndTimeDateTz 				=	DateTzDto{}
	tDur.TimeDuration     		= time.Duration(0)
	tDur.Years								= 0
	tDur.YearsNanosecs    		= 0
	tDur.Months           		= 0
	tDur.MonthsNanosecs   		= 0
	tDur.Weeks            		= 0
	tDur.WeeksNanosecs    		= 0
	tDur.WeekDays							= 0
	tDur.WeekDaysNanosecs			= 0
	tDur.DateDays							= 0
	tDur.DateDaysNanosecs			= 0
	tDur.Hours								= 0
	tDur.HoursNanosecs				= 0
	tDur.Minutes							= 0
	tDur.MinutesNanosecs			= 0
	tDur.Seconds							= 0
	tDur.SecondsNanosecs			= 0
	tDur.Milliseconds					= 0
	tDur.MillisecondsNanosecs	= 0
	tDur.Microseconds					= 0
	tDur.MicrosecondsNanosecs = 0
	tDur.Nanoseconds					= 0
	tDur.TotSubSecNanoseconds = 0
}


// NewStartEndTimesTz - Creates and returns a new TimeDurationDto populated with 
// time duration data based on 'startDateTime' and 'endDateTime' input parameters.
// The user is required to specify a common Time Zone Location four use in converting
// date times to a common frame of reference to subsequent time duration calculations.
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
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
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
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesTz(startTime, endTime, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)
//
//		Note: 'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
// 							datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartEndTimesTz(startDateTime, endDateTime time.Time,
																timeZoneLocation, dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartEndTimesTz() "																	
	
	tDurDto := TimeDurationDto{}
	
	err := tDurDto.SetStartEndTimesTz(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)
	
	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix + "Error returned from " + 
			"SetStartEndTimesTz(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)." +
			"Error='%v'", err.Error())
	}
															
	return tDurDto, nil
}

// NewStartTimeDurationTz - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTime' and time duration. 'startDateTime' is converted to the specified
// 'timeZoneLocation' and the duration value is added to it in order to compute the
// ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to
// ending date time and the	actual starting date time is computed by
// subtracting duration.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
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
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationTz(startTime, duration, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)
//
//		Note: 'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
// 							datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartTimeDurationTz(startDateTime time.Time,
	duration time.Duration, timeZoneLocation, dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartTimeDurationTz() "

	t2Dur := TimeDurationDto{}

	err := t2Dur.SetStartTimeDurationTz(startDateTime, duration, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix + "Error returned by t2Dur.SetStartTimeDurationTz(...) Error='%v'", err.Error())
	}

	return t2Dur, nil

}

// NewStartTimePlusTimeDto - Creates and returns a new TimeDurationDto setting 
// the start date time, end date time and duration based on a starting date time
// and the time components contained in a TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	-   Starting date time. The ending date time will be computed
// 															by adding the time components of the 'plusTimeDto' to
// 															'startDateTime'.
//
// plusTimeDto		TimeDto 	- 	Time components (Years, months, weeks, days, hours etc.)
//															which will be added to 'startDateTime' to compute
//															time duration and ending date time.
//
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
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimePlusTimeDto(startTime, plusTimeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)
//
//		Note: 'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
// 							datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartTimePlusTimeDto(startDateTime time.Time,
								plusTimeDto TimeDto, timeZoneLocation, dateTimeFmtStr string)	(TimeDurationDto, error) {
									
	ePrefix := "TimeDurationDto.NewStartTimePlusTimeDto() "
	
	t2Dur := TimeDurationDto{}
	
	err := t2Dur.SetStartTimePlusTimeDto(startDateTime, plusTimeDto, timeZoneLocation, dateTimeFmtStr)
	
	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix + "Error returned by t2Dur.SetStartTimePlusTimeDto(...) Error='%v'", err.Error())
	}
	
	return t2Dur, nil
}

// NewEndTimeMinusTimeDto - Creates and returns a new TimeDurationDto setting 
// start date time, end date time and duration based on an ending date time
// and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// Input Parameters:
// =================
//
// endDateTime	time.Time	-   Ending date time. The starting date time will be computed
// 														by subtracting minusTimeDto from 'endDateTime'
//
// minusTimeDto	TimeDto 	- 	Time components (Years, months, weeks, days, hours etc.)
//														which will be subtracted from 'endDateTime' to compute
//														time duration and starting date time.
//
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
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewEndTimeMinusTimeDto(endTime, minusTimeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)
//
//		Note: 'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
// 							datetimeconstants.go
//
func (tDur TimeDurationDto) NewEndTimeMinusTimeDto(endDateTime time.Time,
								minusTimeDto TimeDto, timeZoneLocation, dateTimeFmtStr string)	(TimeDurationDto, error) {
									
	ePrefix := "TimeDurationDto.NewEndTimeMinusTimeDto() "
	
	t2Dur := TimeDurationDto{}
	
	err := t2Dur.SetEndTimeMinusTimeDto(endDateTime, minusTimeDto, timeZoneLocation, dateTimeFmtStr)
	
	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix + "Error returned by t2Dur.SetEndTimeMinusTimeDto(...) Error='%v'", err.Error())
	}
	
	return t2Dur, nil
}

// SetEndTimeMinusTimeDto - Sets start date time, end date time and duration
// based on an ending date time and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// Input Parameters:
// =================
//
// endDateTime	time.Time	-   Ending date time. The starting date time will be computed
// 														by subtracting minusTimeDto from 'endDateTime'
//
// minusTimeDto	TimeDto 	- 	Time components (Years, months, weeks, days, hours etc.)
//														which will be subtracted from 'endDateTime' to compute
//														time duration and starting date time.
//
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
func (tDur *TimeDurationDto) SetEndTimeMinusTimeDto(endDateTime time.Time,
	minusTimeDto TimeDto, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "TimeDurationDto.SetEndTimeMinusTimeDto() "

	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)
	dtFormat := tDur.preProcessDateFormatStr(dateTimeFmtStr)

	eDateTime, err := TimeZoneDto{}.New(endDateTime, tzLoc, dtFormat)

	if err!= nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(endDateTime, tzLoc, " +
			"dtFormat). Error='%v'", err.Error())
	}

	tDur.Empty()

	tDur.EndTimeDateTz = eDateTime.TimeOut.CopyOut()

	tDur.StartTimeDateTz, err = eDateTime.TimeOut.MinusTimeDto(minusTimeDto)

	tDur.TimeDuration = tDur.EndTimeDateTz.DateTime.Sub(tDur.StartTimeDateTz.DateTime)

	err = tDur.calcTimeDurationComponents()

	if err != nil {
		tDur.Empty()
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcTimeDurationComponents(). " +
			"Error='%v'", err.Error())
	}

	return nil
}

// SetStartEndTimesTz - Sets data field values for the current TimeDurationDto
// instance using a Start Date Time, End Date Time and a time zone specification.
// First, 'startDateTime' and 'endDateTime' are converted to the designate Time
// Zone Location. Next, 'startDateTime' is subtracted from 'endDateTime' to compute
// time duration.
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
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
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
func (tDur *TimeDurationDto) SetStartEndTimesTz(startDateTime,
endDateTime time.Time, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "TimeDurationDto.SetStartEndTimes() "
	
	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)
	dtFormat := tDur.preProcessDateFormatStr(dateTimeFmtStr)

	sTime, err := TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat)

	if err != nil {
		return fmt.Errorf(ePrefix + 
			"Error returned by TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat). " +
			"Error='%v'", err.Error())
	}

	eTime, err := TimeZoneDto{}.New(endDateTime, tzLoc, dtFormat)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(endDateTime, tzLoc, dtFormat). " +
			"Error='%v'", err.Error())
	}
	
	if eTime.TimeOut.DateTime.Before(sTime.TimeOut.DateTime) {
		s2 := sTime.CopyOut()
		sTime = eTime.CopyOut()
		eTime = s2.CopyOut()
	}

	tDur.Empty()
	tDur.StartTimeDateTz = sTime.TimeOut.CopyOut()
	tDur.EndTimeDateTz	= eTime.TimeOut.CopyOut()
	tDur.TimeDuration = tDur.EndTimeDateTz.DateTime.Sub(tDur.StartTimeDateTz.DateTime)

	err = tDur.calcTimeDurationComponents()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcTimeDurationComponents(). " +
				"Error='%v'", err.Error())
	}
	
	return nil
}

// SetStartTimeDurationTz - Sets start time, end time and duration for the
// current TimeDurationDto instance. 'startDateTime' is converted to the
// specified 'timeZoneLocation' and the duration value is added to it
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to
// ending date time and the	actual starting date time is computed by
// subtracting duration.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
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
//
func (tDur *TimeDurationDto) SetStartTimeDurationTz(startDateTime time.Time,
	duration time.Duration, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "TimeDurationDto.SetStartTimeDurationTz() "
	
	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)
	dtFormat := tDur.preProcessDateFormatStr(dateTimeFmtStr)

	xTime, err := TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat)

	if err!= nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(startDateTime, tzLoc, " +
			"dtFormat). Error='%v'", err.Error())
	}

	tDur.Empty()

	if duration < 0 {

		tDur.EndTimeDateTz = xTime.TimeOut.CopyOut()

		tDur.StartTimeDateTz, err = tDur.EndTimeDateTz.AddDuration(duration, dtFormat)

		if err != nil {
			tDur.Empty()
			return fmt.Errorf(ePrefix + "Error returned from tDur.EndTimeDateTz."+
				"AddDuration(duration, dtFormat) " +
				" Error='%v'", err.Error())

		}

		tDur.TimeDuration = duration * -1

	} else {

		tDur.StartTimeDateTz = xTime.TimeOut.CopyOut()

		tDur.EndTimeDateTz, err = tDur.StartTimeDateTz.AddDuration(duration, dtFormat)

		if err != nil {
			tDur.Empty()
			return fmt.Errorf(ePrefix + "Error returned from tDur.StartTimeDateTz."+
				"AddDuration(duration, dtFormat) " +
				" Error='%v'", err.Error())
		}

		tDur.TimeDuration = duration

	}

	err = tDur.calcTimeDurationComponents()

	if err != nil {
		tDur.Empty()
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcTimeDurationComponents(). " +
			"Error='%v'", err.Error())
	}

	return nil
}

// SetStartTimePlusTimeDto - Sets start date time, end date time and duration
// based on a starting date time and the time components contained in a TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	-   Starting date time. The ending date time will be computed
// 															by adding the time components of the 'plusTimeDto' to
// 															'startDateTime'.
//
// plusTimeDto		TimeDto 	- 	Time components (Years, months, weeks, days, hours etc.)
//															which will be added to 'startDateTime' to compute
//															time duration and ending date time.
//
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
func (tDur *TimeDurationDto) SetStartTimePlusTimeDto(startDateTime time.Time,
	plusTimeDto TimeDto, timeZoneLocation, dateTimeFmtStr string) error {
	
	ePrefix := "TimeDurationDto.SetStartTimePlusTimeDto() "

	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)
	dtFormat := tDur.preProcessDateFormatStr(dateTimeFmtStr)

	sDateTime, err := TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat)

	if err!= nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(startDateTime, tzLoc, " +
			"dtFormat). Error='%v'", err.Error())
	}

	tDur.Empty()
	
	tDur.StartTimeDateTz = sDateTime.TimeOut.CopyOut()
	
	tDur.EndTimeDateTz, err = sDateTime.TimeOut.PlusTimeDto(plusTimeDto)
	
	tDur.TimeDuration = tDur.EndTimeDateTz.DateTime.Sub(tDur.StartTimeDateTz.DateTime)

	err = tDur.calcTimeDurationComponents()

	if err != nil {
		tDur.Empty()
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcTimeDurationComponents(). " +
			"Error='%v'", err.Error())
	}

	return nil		
}
	
// calcTimeDurationComponents - Is a summary routine which calls all
// of the subsidiary methods necessary to compute the duration time
// components (i.e. years, months, days, hours, minutes etc.).
func (tDur *TimeDurationDto) calcTimeDurationComponents() error {
	
ePrefix := "TimeDurationDto) calcTimeDurationComponents() "
	
	err := tDur.calcYearsFromDuration()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcYearsFromDuration(). Error='%v'", err.Error())
	}
	
	err = tDur.calcMonthsFromDuration()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcMonthsFromDuration(). Error='%v'", err.Error())
	}
	
	err = tDur.calcDateDaysWeeksFromDuration()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcDateDaysWeeksFromDuration(). Error='%v'", err.Error())
	}
	
	err = tDur.calcHoursMinSecs()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}
	
	err = tDur.calcAllNanoseconds()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcAllNanoseconds(). Error='%v'", err.Error())
	}
	
	return nil
}

// calcYearsFromDuration - Calculates number of years duration and nanoseconds
// represented by years duration using input parameters 'tDur.StartTimeDateTz' and
// 'tDur.EndTimeDateTz'.  Note: If the Time Zone Locations of 'tDur.StartTimeDateTz'
// and 'tDur.EndTimeDateTz' do NOT match, an error will be returned.
//
func (tDur *TimeDurationDto) calcYearsFromDuration() error {

	ePrefix := "TimeDurationDto.calcYearsFromDuration() "

	years := int64(0)
	yearNanosecs := int64(0)
	startTime := tDur.StartTimeDateTz.DateTime
	endTime := tDur.EndTimeDateTz.DateTime

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix + "Error: 'startTime' and 'endTime' Time Zone Location do NOT match! " +
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
// 'tDur.EndTimeDateTz.DateTime'. Note: Method TimeDurationDto.calcYearsFromDuration()
// must be called BEFORE calling this method.
func (tDur *TimeDurationDto) calcMonthsFromDuration() error {

	ePrefix := "TimeDurationDto.calcMonthsFromDuration() "

	startTime := tDur.StartTimeDateTz.DateTime
	endTime := tDur.EndTimeDateTz.DateTime

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix + "Error: 'startTime' and 'endTime' Time Zone Location do NOT match! " +
			"startTimeZoneLocation='%v'  endTimeZoneLocation='%v'",
			startTime.Location().String(), endTime.Location().String())
	}

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return errors.New(ePrefix + "Error: tDur.TimeDuration is ZERO!")
	}


	rd -= tDur.YearsNanosecs

	i := 0

	yearDateTime := startTime.Add(time.Duration(tDur.YearsNanosecs))

	mthDateTime := yearDateTime

	for mthDateTime.Before(endTime) {

		i++

		mthDateTime = yearDateTime.AddDate(0,i,0)

	}

	i -= 1

	if i > 0 {

		tDur.Months = int64(i)

		mthDateTime = yearDateTime.AddDate( 0, i, 0)

		tDur.MinutesNanosecs = int64(mthDateTime.Sub(yearDateTime))

	} else {
		tDur.Months = 0
		tDur.MonthsNanosecs = 0
	}

	return nil
}

// calcDateDaysWeeksFromDuration - Calculates the Days associated
// with the duration for this TimeDurationDto. Note method TimeDurationDto.calcMonthsFromDuration()
// MUST BE called before this method.
//
// Calculates 'tDur.DateDays', 'tDur.DateDaysNanosecs', 'tDur.Weeks', 'tDur.WeeksNanosecs', 
// 'tDur.WeekDays' and 'tDur.WeekDaysNanosecs'.
//
func (tDur *TimeDurationDto) calcDateDaysWeeksFromDuration() error {

	ePrefix := "TimeDurationDto.calcDateDaysFromDuration() "

	startTime := tDur.StartTimeDateTz.DateTime
	endTime := tDur.EndTimeDateTz.DateTime

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix + "Error: 'startTime' and 'endTime' Time Zone Location do NOT match! " +
			"startTimeZoneLocation='%v'  endTimeZoneLocation='%v'",
			startTime.Location().String(), endTime.Location().String())
	}

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return errors.New(ePrefix + "Error: tDur.TimeDuration is ZERO!")
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

		tDur.WeekDays = tDur.DateDays -  (tDur.Weeks * 7)
		tDur.WeekDaysNanosecs = tDur.WeekDays * DayNanoSeconds
		
	}

	return nil
}

// calcHoursMinSecs - Calculates Hours, Minute, and 
// Seconds of duration using startTime, tDur.StartTimeDateTz, 
// and endTime, tDur.EndTimeDateTz.DateTime.
//
// NOTE: 	Method TimeDurationDto.calcDateDaysFromDuration()
//				MUST BE CALLED BEFORE this method.
func (tDur *TimeDurationDto) calcHoursMinSecs() error {
	
	ePrefix := "TimeDurationDto.calcHoursMinSecs() "

	startTime := tDur.StartTimeDateTz.DateTime
	endTime := tDur.EndTimeDateTz.DateTime

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix + "Error: 'startTime' and 'endTime' Time Zone Location do NOT match! " +
			"startTimeZoneLocation='%v'  endTimeZoneLocation='%v'",
			startTime.Location().String(), endTime.Location().String())
	}

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return errors.New(ePrefix + "Error: tDur.TimeDuration is ZERO!")
	}

	rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs +
					tDur.DateDaysNanosecs

	tDur.Hours 						= 0
	tDur.HoursNanosecs 		= 0
	tDur.Minutes 					= 0
	tDur.MinutesNanosecs 	= 0
	tDur.Seconds 					= 0 
	tDur.SecondsNanosecs 	= 0

	if rd >= HourNanoSeconds {
		tDur.Hours = rd / HourNanoSeconds
		tDur.HoursNanosecs = HourNanoSeconds * tDur.Hours
		rd -= tDur.HoursNanosecs
	}

	if rd >= MinuteNanoSeconds {
		tDur.Minutes = rd / MinuteNanoSeconds
		tDur.MinutesNanosecs = MinuteNanoSeconds * tDur.Minutes
		rd -=tDur.MinutesNanosecs
	}

	if rd >= SecondNanoseconds {
		tDur.Seconds = rd / SecondNanoseconds
		tDur.SecondsNanosecs = SecondNanoseconds * tDur.Seconds
		rd -= tDur.SecondsNanosecs
	}

	return nil
}


// calcAllNanoseconds - Calculates 'tDur.Milliseconds', 'tDur.MillisecondsNanosecs', 
// 'tDur.Microseconds', 'tDur.MicrosecondsNanosecs', 'tDur.Nanoseconds' and 
// 'tDur.TotSubSecNanoseconds'.
//
// NOTE: 	Method TimeDurationDto.calcHoursMinSecs() MUST BE CALLED BEFORE calling
// 				this method.
//
func (tDur *TimeDurationDto) calcAllNanoseconds() error {

	ePrefix := "TimeDurationDto.calcAllNanoseconds() "

	startTime := tDur.StartTimeDateTz.DateTime
	endTime := tDur.EndTimeDateTz.DateTime

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix + "Error: 'startTime' and 'endTime' Time Zone Location do NOT match! " +
			"startTimeZoneLocation='%v'  endTimeZoneLocation='%v'",
			startTime.Location().String(), endTime.Location().String())
	}

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return errors.New(ePrefix + "Error: tDur.TimeDuration is ZERO!")
	}

	rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs +
					tDur.DateDaysNanosecs + tDur.HoursNanosecs +
					tDur.MinutesNanosecs + tDur.SecondsNanosecs
					
	tDur.Milliseconds 					= 0
	tDur.MillisecondsNanosecs 	= 0
	tDur.Microseconds						= 0
	tDur.MicrosecondsNanosecs		= 0
	tDur.Nanoseconds 						= 0
	tDur.TotSubSecNanoseconds		= rd
	
	if rd >= MilliSecondNanoseconds {
		tDur.Milliseconds = rd / MilliSecondNanoseconds
		tDur.MillisecondsNanosecs = MilliSecondNanoseconds * tDur.Milliseconds
		rd -= tDur.MicrosecondsNanosecs
	}

	if rd >= MicroSecondNanoseconds {
		tDur.Microseconds = rd / MicroSecondNanoseconds
		tDur.MicrosecondsNanosecs = MicroSecondNanoseconds * tDur.Microseconds
		rd -= tDur.MicrosecondsNanosecs
	}

	tDur.Nanoseconds = rd
	
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
		return TzIanaUTC
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		return "Local"
	}

	return timeZoneLocation
}
