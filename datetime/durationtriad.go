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

	You will find this source file, 'durationtriad.go' in the subdirectory:

			datetimeopsgo\datetime


	Dependencies
	============

	(1) 'timezonedto.go'
	(2)	'durationtimedto.go'



	Overview and Usage
	==================

	The principal component of this library is the DurationTriad. This
	type plus associated methods are used to manage, calculate and analyze
	time duration calculations.

	Usage generally involves providing start time and and end time as inputs. Thereafter
	the 'DurationTriad' calculates time duration and provides multiple formats
  to displaying the resulting time duration.

	When providing start times and end times, methods usually require standardized
	time zones or time zone locations.  If duration calculations are performed using
	start and end times with differing time zone locations, the calculation could
	produce unexpected results. are provided with are provided with


	Usage requires two operations:

	1. You must first initialize the DurationTriad type using one of the
		 four 'Set' methods shown below:
		 	a. SetStartTimeDuration() 		also NewStartTimeDuration()
		 	b. SetStartEndTimes()					also NewStartEndTimes()
		 	c. SetStartTimePlusTime()			also NewStartTimePlus()
		 	d. SetEndTimeMinusTimeDto()		also NewEndTimeMinusTimeDto()

	2. After the DurationTriad is initialized in step one above, you are free
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


// DurationTriad - holds elements of
// time duration
type DurationTriad struct {
	BaseTime     TimeDurationDto
	LocalTime    TimeDurationDto
	UTCTime			 TimeDurationDto
}

// CopyIn - Receives and incoming DurationTriad data
// structure and copies the values to the current DurationTriad
// data structure.
func (durT *DurationTriad) CopyIn(duIn DurationTriad) {
	durT.Empty()
	durT.BaseTime = duIn.BaseTime.CopyOut()
	durT.LocalTime = duIn.LocalTime.CopyOut()
	durT.UTCTime = duIn.UTCTime.CopyOut()

	return
}

// CopyOut - Returns a deep copy of the current
// DurationTriad data fields.
func (durT *DurationTriad) CopyOut() DurationTriad {
	duOut := DurationTriad{}
	duOut.BaseTime = durT.BaseTime.CopyOut()
	duOut.LocalTime = durT.LocalTime.CopyOut()
	duOut.UTCTime = durT.UTCTime.CopyOut()

	return duOut
}

// Equal - This method may be used to determine if two
// DurationTriad data structures are equivalent.
func (durT *DurationTriad) Equal(duIn DurationTriad) bool {

	if durT.BaseTime.Equal(duIn.BaseTime)  &&
		durT.LocalTime.Equal(duIn.LocalTime) &&
		durT.UTCTime.Equal(duIn.UTCTime) {

		return true
	}

	return false

}

// Empty - This method initializes
// all of the fields in this
// DurationTriad structure to their
// zero values.
func (durT *DurationTriad) Empty() {
	durT.BaseTime.Empty()
	durT.LocalTime.Empty()
	durT.UTCTime.Empty()
}

// NewStartEndTimes - Returns a New DurationTriad based on two input
// parameters, startDateTime and endDateTime.
//
// 	Usage:
//	======
//
// du, err := DurationTriad{}.NewStartEndTimes(startDateTime, endDateTime)
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
func (durT DurationTriad) NewStartEndTimes(startDateTime,
				endDateTime time.Time, timeZoneLocation, dateTimeFmtStr string ) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartTimeDuration() "

	du2 := DurationTriad{}

	err := du2.SetStartEndTimes(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix + "Error returned from du2.SetStartEndTimes(startDateTime, endDateTime).\nError='%v'", err)
	}

	return du2, nil

}

// NewStartTimeDuration - Returns a New DurationTriad based on 'startDateTime'
// and time.Duration input parameters.
//
// Usage:
//
// du, err := DurationTriad{}.NewStartTimeDuration(startDateTime, timeDto)
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
func (durT DurationTriad) NewStartTimeDuration(startDateTime time.Time, duration time.Duration,
						timeZoneLocation, dateTimeFmtStr string ) (DurationTriad, error) {


	ePrefix := "DurationTriad.NewStartTimeDuration() "

	du2 := DurationTriad{}

	err := du2.SetStartTimeDuration(startDateTime, duration, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix + "Error returned from du2.SetStartTimeDuration(startDateTime, duration).\nError='%v'", err)
	}

	return du2, nil
}


// NewEndTimeMinusTimeDto - Returns a new DurationTriad based on two input parameters,
// 'endDateTime' and 'timeDto'. 'timeDto' is an instance of TimeDto which is
// subtracted from 'endDateTime' in order to calculate time duration. The user is
// required to provide Time Zone Location as an input parameter in order to ensure that
// time duration calculations are performed using equivalent time zones.
//
// Usage:
//
// du, err := DurationTriad{}.NewEndTimeMinusTimeDto(startDateTime, timeDto)
//
// Input Parameters:
// =================
//
// endDateTime		time.Time	- Ending date time. The TimeDto parameter will be subtracted
//														from this date time in order to compute the starting date time.
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
func (durT DurationTriad) NewEndTimeMinusTimeDto(endDateTime time.Time, minusTimeDto TimeDto,
														timeZoneLocation, dateTimeFmtStr string) (DurationTriad, error){

ePrefix := "DurationTriad.NewEndTimeMinusTimeDto() "

	du2 := DurationTriad{}

	err := du2.SetEndTimeMinusTimeDto(endDateTime, minusTimeDto, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix + "Error returned from du2.SetEndTimeMinusTimeDto(endDateTime, minusTimeDto).\nError='%v'", err)
	}

	return du2, nil
}

// NewStartTimePlusTime - Returns a New DurationTriad based on 'startDateTime'
// and DurationDto input parameters.
//
// Usage:
// du, err := DurationTriad{}.NewStartTimePlusTime(startDateTime, timeDto)
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting time
//
// plusTimeDto		  TimeDto - Provides time values which will be subtracted from
//														'startDateTime' in order to calculate duration.
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
//										TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
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
func (durT DurationTriad) NewStartTimePlusTime(startDateTime time.Time, plusTimeDto TimeDto,
														timeZoneLocation, dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartTimePlusTime() "

	du2 := DurationTriad{}

	err := du2.SetStartTimePlusTime(startDateTime, plusTimeDto, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix + "Error returned from du2.SetStartTimePlusTime(startDateTime, plusTimeDto).\nError='%v'", err)
	}

	return du2, nil
}

// SetStartTimeDuration - Receives a starting date time and
// a time duration. The method then calculates the ending
// date time, duration and populates the DurationTriad
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
func (durT *DurationTriad) SetStartTimeDuration(startDateTime time.Time,
				duration time.Duration, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartTimeDuration() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error - Start Time is Zero!")
	}

	tzLoc := durT.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: Input Parameter 'timeZoneLocation' INVALID. " +
			"timeZoneLocation='%v' tzLoc='%v' Error='%v'",
				timeZoneLocation, tzLoc, err.Error())
	}

	baseTime, err := TimeDurationDto{}.NewStartTimeDurationTzCalc(startDateTime, duration,
										tzLoc, TDurCalcTypeSTDYEARMTH, dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix +
		"baseTime calculation Error returned by " +
		"TimeDurationDto{}.NewStartTimeDurationTzCalc() Error='%v'",
			err.Error())

	}

	localTime, err := TimeDurationDto{}.NewStartTimeDurationTzCalc(startDateTime, duration,
										TzGoLocal, TDurCalcTypeSTDYEARMTH, dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix +
		"localTime calculation Error returned by " +
		"TimeDurationDto{}.NewStartTimeDurationTzCalc() Error='%v'",
			err.Error())

	}

	utcTime, err := TimeDurationDto{}.NewStartTimeDurationTzCalc(startDateTime, duration,
										TzGoLocal, TDurCalcTypeSTDYEARMTH, dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix +
		"utcTime calculation Error returned by " +
		"TimeDurationDto{}.NewStartTimeDurationTzCalc() Error='%v'",
			err.Error())

	}

	durT.Empty()
	durT.BaseTime = baseTime.CopyOut()
	durT.LocalTime = localTime.CopyOut()
	durT.UTCTime = utcTime.CopyOut()

	err = durT.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "ERROR: Duration Base Data is INVALID! Error='%v'", err.Error())
	}

	return nil
}

// SetStartEndTimes - Calculate duration values and save the results in the DurationTriad
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
func (durT *DurationTriad) SetStartEndTimes(startDateTime,
					endDateTime time.Time, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartEndTimes() "

	if startDateTime.IsZero()  && endDateTime.IsZero() {
		return errors.New(ePrefix +
			"Error: Input parameters 'startDateTime' and 'endDateTime' are ZERO!")
	}


	tzLoc := durT.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Input paramenter 'timeZoneLocation' is INVALID. " +
			" time.LoadLocation(tzLoc). timeZoneLocation='%v', tzLoc='%v', Error='%v'",
				timeZoneLocation, tzLoc, err.Error())
	}

	baseTime, err := TimeDurationDto{}.NewStartEndTimesTzCalc(startDateTime,
						endDateTime, timeZoneLocation, TDurCalcTypeSTDYEARMTH, dateTimeFmtStr )

	if err != nil {
		return fmt.Errorf(ePrefix +
			"baseTime error returned by TimeDurationDto{}.NewStartEndTimesTzCalc(). " +
			"Error='%v' ", err.Error())
	}

	localTime, err := TimeDurationDto{}.NewStartEndTimesTzCalc(startDateTime,
		endDateTime, TzGoLocal, TDurCalcTypeSTDYEARMTH, dateTimeFmtStr )

	if err != nil {
		return fmt.Errorf(ePrefix +
			"localTime error returned by TimeDurationDto{}.NewStartEndTimesTzCalc(). " +
			"Error='%v' ", err.Error())
	}

	utcTime, err := TimeDurationDto{}.NewStartEndTimesTzCalc(startDateTime,
		endDateTime, TzIanaUTC, TDurCalcTypeSTDYEARMTH, dateTimeFmtStr )


	durT.Empty()
	durT.BaseTime = baseTime.CopyOut()
	durT.LocalTime = localTime.CopyOut()
	durT.UTCTime = utcTime.CopyOut()

	err = durT.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "ERROR: Duration Base Data is INVALID! Error='%v'", err.Error())
	}

	return nil
}

// SetEndTimeMinusTimeDto - Calculate duration values based on an Ending Date Time and
// a TimeDto structure consisting of time values (Years, Months, weeks, days, hours,
// minutes etc.). The time values in the 'timeDto' parameter are subtracted
// from 'endDateTime'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time.
//
// As a result. true values for StartTimeDateTz, EndTimeDateTz and
// TimeDuration are stored in the DurationTriad data structure.
//
// Input Parameters
// ================
//
// endDateTime		time.Time - The ending date time value from which TimeDto
//														parameter 'minusTimeDto' will be subtracted
//														in order to compute the Starting Date Time.
//
// minusTimeDto			TimeDto - An instance of TimeDto containing time values,
// 														(Years, Months, weeks, days, hours, minutes etc.)
//														which will be subtracted from input parameter
// 														'endDateTime' in order to compute the Starting
//														Date Time and Time Duration.
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
func (durT *DurationTriad) SetEndTimeMinusTimeDto(endDateTime time.Time,
		minusTimeDto TimeDto, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetEndTimeMinusTimeDto() "


	tzLoc := durT.preProcessTimeZoneLocation(timeZoneLocation)

	targetLoc, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix 	+
			"Error: TimeZoneLocation input parameter is INVALID! " +
			"timeZoneLocation='%v' tzLoc='%v'  Error='%v'",
				timeZoneLocation, tzLoc, err.Error())
	}

	targetEndDateTime := endDateTime.In(targetLoc)

	baseTime, err := TimeDurationDto{}.NewEndTimeMinusTimeDto(targetEndDateTime,minusTimeDto, dateTimeFmtStr )

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error computing baseTime. TimeDurationDto{}.NewEndTimeMinusTimeDto(). " +
			"targetEndDateTime='%v' Error='%v'",
				targetEndDateTime, err.Error())
	}


	targetLoc, err = time.LoadLocation(TzGoLocal)

	if err != nil {
		return fmt.Errorf(ePrefix 	+
			"Error: Local TimeZoneLocation input parameter is INVALID! " +
			"timeZoneLocation='%v' Error='%v'",
			TzGoLocal, err.Error())
	}

	targetEndDateTime = endDateTime.In(targetLoc)

	localTime, err := TimeDurationDto{}.NewEndTimeMinusTimeDto(targetEndDateTime,minusTimeDto, dateTimeFmtStr )

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error computing localTime. TimeDurationDto{}.NewEndTimeMinusTimeDto(). " +
			"targetEndDateTime='%v' Error='%v'",
			targetEndDateTime, err.Error())
	}


	targetLoc, err = time.LoadLocation(TzIanaUTC)

	if err != nil {
		return fmt.Errorf(ePrefix 	+
			"Error: UTC TimeZoneLocation input parameter is INVALID! " +
			"timeZoneLocation='%v' Error='%v'",
			TzIanaUTC, err.Error())
	}

	targetEndDateTime = endDateTime.In(targetLoc)

	utcTime, err := TimeDurationDto{}.NewEndTimeMinusTimeDto(targetEndDateTime,minusTimeDto, dateTimeFmtStr )

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error computing localTime. TimeDurationDto{}.NewEndTimeMinusTimeDto(). " +
			"targetEndDateTime='%v' Error='%v'",
			targetEndDateTime, err.Error())
	}


	durT.Empty()
	durT.BaseTime = baseTime.CopyOut()
	durT.LocalTime = localTime.CopyOut()
	durT.UTCTime = utcTime.CopyOut()

	err = durT.IsValid()

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
// then stored in the DurationTriad data structure.
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
func (durT *DurationTriad) SetStartTimePlusTime(startDateTime time.Time, plusTimeDto TimeDto,
	timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartTimePlusTime() "

	tzLoc := durT.preProcessTimeZoneLocation(timeZoneLocation)

	targetLoc, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: TimeZoneLocation is INVALID! " +
			"timeZoneLocation='%v'  tzLoc='%v'  Error='%v'",
				timeZoneLocation, tzLoc, err.Error())
	}

	targetStartDateTime := startDateTime.In(targetLoc)

	baseTime, err := TimeDurationDto{}.NewStartTimePlusTimeDto(targetStartDateTime,
											plusTimeDto, dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "baseTime calculation error returned by " +
			"TimeDurationDto{}.NewStartTimePlusTimeDto(). Error=%v'",
				err.Error())
	}

	targetLoc, err = time.LoadLocation(TzGoLocal)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: Local TimeZoneLocation is INVALID! " +
			"timeZoneLocation='%v' Error='%v'",
				TzGoLocal, err.Error())
	}

	targetStartDateTime = startDateTime.In(targetLoc)

	localTime, err := TimeDurationDto{}.NewStartTimePlusTimeDto(targetStartDateTime,
											plusTimeDto, dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "localTime calculation error returned by " +
			"TimeDurationDto{}.NewStartTimePlusTimeDto(). Error=%v'",
				err.Error())
	}

	targetLoc, err = time.LoadLocation(TzIanaUTC)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: UTC TimeZoneLocation is INVALID! " +
			"timeZoneLocation='%v' Error='%v'",
				TzGoLocal, err.Error())
	}

	targetStartDateTime = startDateTime.In(targetLoc)

	utcTime, err := TimeDurationDto{}.NewStartTimePlusTimeDto(targetStartDateTime,
											plusTimeDto, dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "utcTime calculation error returned by " +
			"TimeDurationDto{}.NewStartTimePlusTimeDto(). Error=%v'",
				err.Error())
	}


	durT.Empty()
	durT.BaseTime = baseTime.CopyOut()
	durT.LocalTime = localTime.CopyOut()
	durT.UTCTime = utcTime.CopyOut()

	err = durT.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: Duration Base Data INVALID! Error='%v'", err.Error())
	}

	return nil
}

// IsValid - Validates the current DurationTriad instance.
//
func (durT *DurationTriad) IsValid() error {

	ePrefix := "DurationTriad.IsValid() "

	err := durT.BaseTime.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "INVALID durT.BaseTime. Error='%v'", err.Error())
	}

	err = durT.LocalTime.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "INVALID durT.LocalTime. Error='%v'", err.Error())
	}

	err = durT.UTCTime.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "INVALID durT.UTCTime. Error='%v'", err.Error())
	}


	return nil
}


func (durT *DurationTriad) preProcessDateFormatStr(dateTimeFmtStr string) string {

	if len(dateTimeFmtStr) == 0 {
		return FmtDateTimeYrMDayFmtStr
	}

	return dateTimeFmtStr
}


func (durT *DurationTriad) preProcessTimeZoneLocation(timeZoneLocation string) string {

	if len(timeZoneLocation) == 0 {
		return TzIanaUTC
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		return "Local"
	}

	return timeZoneLocation
}