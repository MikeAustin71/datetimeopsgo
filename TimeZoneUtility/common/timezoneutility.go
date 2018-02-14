package common

import (
	"errors"
	"fmt"
	"time"
	"strings"
)

/*
	Overview - Location
	===================

  timezoneutility.go is part of the date time operations library. The source code repository
 	for this file is located at:

					https://github.com/MikeAustin71/datetimeopsgo.git



	Dependencies
	============

	None

 */

// NOTE: See https://golang.org/pkg/time/#LoadLocation
// and https://www.iana.org/time-zones to ensure that
// the IANA Time Zone Database is properly configured
// on your system. Note: IANA Time Zone Data base is
// equivalent to 'tz database'.
const (
	// TzUsEast - USA Eastern Time Zone
	// IANA database identifier
	TzUsEast = "America/New_York"

	// TzUsCentral - USA Central Time Zone
	// IANA database identifier
	TzUsCentral = "America/Chicago"

	// TzUsMountain - USA Mountain Time Zone
	// IANA database identifier
	TzUsMountain = "America/Denver"

	// TzUsPacific - USA Pacific Time Zone
	// IANA database identifier
	TzUsPacific  = "America/Los_Angeles"

	// TzUsHawaii - USA Hawaiian Time Zone
	// IANA database identifier
	TzUsHawaii = "Pacific/Honolulu"

	// tzUTC - UTC Time Zone IANA database
	// identifier
	TzUTC = "Zulu"

	NeutralDateFmt = "2006-01-02 15:04:05.000000000"
)

// DateTzDto - `Used to store and transfer date times.
// The descriptors contained is this structure are intended
// to define and identify a specific point in time.
//
// This Type is NOT used to define duration; that is, the
// difference or time span between two point in time. For
// these types of operations see:
// DurationTimeUtility/common/durationutil.go
//
// DateTzDto defines a specific point in time using
// a variety of descriptors including year, month, day
// hour, minute, second, millisecond, microsecond and
// and nanosecond. In addition this Type specifies a
// time.Time value as well as time zone location and
// time zone.
//
// If you are unfamiliar with the concept of a time
// zone location, consider the field TimeLoc and
// TimeLocName below:
//
// Time zone location must be designated as one of two values.
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
type DateTzDto struct {
	Year       			int							// Year Number
	Month      			int							// Month Number
	Day        			int							// Day Number
	Hour       			int							// Hour Number
	Minute     			int							// Minute Number
	Second     			int							// Second Number
	Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
	Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
	Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
	TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
	TimeZone   			string					// Time Zone associated with this Date Time. Example: "CDT" (abbreviation for Central Daylight Time)
	TimeZoneOffset	int							// TimeZoneOffset associated with this Date Time
	DateTime 				time.Time				// DateTime value for this DateTzDto Type
	TimeLoc    			*time.Location	// Time Location pointer associated with this DateTime value
	TimeLocName			string					// Time Location Name. Example: "America/Chicago"
}

// New - returns a new DateTzDto instance based on a time.Time ('dateTime')
// input parameter.
//
// Input Parameter
// ===============
//
// dateTime   time.Time - a date time value
//
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type defined as follows:
//
//	type DateTzDto struct {
//		Year       			int							// Year Number
//		Month      			int							// Month Number
//		Day        			int							// Day Number
//		Hour       			int							// Hour Number
//		Minute     			int							// Minute Number
//		Second     			int							// Second Number
//		Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//		Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//		Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																		// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//		TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//		TimeZone   			string					// Time Zone associated with this Date Time. Example: "CDT" (abbreviation for Central Daylight Time)
//		TimeZoneOffset	int							// TimeZoneOffset associated with this Date Time
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		TimeLoc    			*time.Location	// Time Location pointer associated with this DateTime value
//		TimeLocName			string					// Time Location Name. Example: "America/Chicago"
//	}
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
// Usage
// =====
//
// Example:
//			dtzDto, err := DateTzDto{}.New(dateTime)
//
func (dtz DateTzDto) New(dateTime time.Time)(DateTzDto, error) {
	ePrefix := "DateTzDto.New() "

	if dateTime.IsZero() {
		return DateTzDto{}, errors.New(ePrefix + "Error: Input parameter dateTime is Zero value!")
	}

	dtz2 := DateTzDto{}
	dtz2.Year  = dateTime.Year()
	dtz2.Month = int(dateTime.Month())
	dtz2.Day = dateTime.Day()
	dtz2.Hour = dateTime.Hour()
	dtz2.Minute = dateTime.Minute()
	dtz2.Second = dateTime.Second()
	dtz2.allocateNanoseconds(int64(dateTime.Nanosecond()))
	dtz2.DateTime = dateTime
	dtz2.TimeLoc = dateTime.Location()
	dtz2.TimeLocName = dtz2.TimeLoc.String()
	dtz2.TimeZone, dtz2.TimeZoneOffset = dateTime.Zone()

	return dtz2, nil
}

// NewDateTimeElements - creates a new DateTzDto object and populates the data fields based on
// input parameters.
//
// Input Parameters
// ================
//
// year 						int			- year number
// month						int			- month number 	1 - 12
// day							int			- day number   	1 - 31
// hour							int			- hour number  	0 - 24
// min							int			- minute number	0 - 59
// sec							int			- second number	0	-	59
// nsec							int			- nanosecond number 0 - 999999999
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
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type defined as follows:
//
//	type DateTzDto struct {
//		Year       			int							// Year Number
//		Month      			int							// Month Number
//		Day        			int							// Day Number
//		Hour       			int							// Hour Number
//		Minute     			int							// Minute Number
//		Second     			int							// Second Number
//		Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//		Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//		Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																		// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//		TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//		TimeZone   			string					// Time Zone associated with this Date Time. Example: "CDT" (abbreviation for Central Daylight Time)
//		TimeZoneOffset	int							// TimeZoneOffset associated with this Date Time
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		TimeLoc    			*time.Location	// Time Location pointer associated with this DateTime value
//		TimeLocName			string					// Time Location Name. Example: "America/Chicago"
//	}
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
// Usage
// =====
//
// Example:
//			dtzDto, err := DateTzDto{}.NewDateTimeElements(year, month, day, hour, min, sec, nanosecond , timeZoneLocation)
//
//
func (dtz DateTzDto) NewDateTimeElements(year, month, day, hour, min, sec, nanosecond int, timeZoneLocation string) (DateTzDto, error) {

	ePrefix := "DateTzDto.New() "

	dtz2 := DateTzDto{}

	if year < 0 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter year number is INVALID. 'year' must be greater than or equal to Zero. year='%v'", year)
	}

	if month < 1 || month > 12  {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter month number is INVALID. Correct range is 1-12. month='%v'", month)
	}


	if day < 1 || day > 31  {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter 'day' number is INVALID. Correct range is 1-31. day='%v'", day)
	}


	if hour < 0 || hour > 24 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter 'hour' number is INVALID. Correct range is 0-24. hour='%v'", hour)
	}

	if min < 0 || min > 59 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter minute number is INVALID. Correct range is 0 - 59. min='%v'", min)
	}

	if sec < 0 || sec > 59 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parmeter second number is INVALID. Correct range is 0 - 59. sec='%v'", sec)
	}


	maxNanoSecs := int(time.Second) - int(1)

	if nanosecond < 0 || nanosecond > maxNanoSecs {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter nanoseconds exceeds maximum limit and is INVLIAD. Correct range is 0 - %v. nanosecond='%v'", maxNanoSecs, nanosecond)
	}


	if len(timeZoneLocation) == 0 {
		return dtz2, errors.New(ePrefix + "Error: Input parameter 'timeZoneLocation' is an EMPTY STRING! 'timeZoneLocation' is required!")
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		timeZoneLocation = "Local"
	}

	loc, err := time.LoadLocation(timeZoneLocation)

	if err != nil {
		return dtz2, fmt.Errorf(ePrefix + "Error: Invalid time zone location! timeZoneLocation='%v'", timeZoneLocation)
	}

	dtz2.Year 			= year
	dtz2.Month			= month
	dtz2.Day 				= day
	dtz2.Hour 			= hour
	dtz2.Minute			= min
	dtz2.Second			= sec
	dtz2.TimeLoc 		= loc
	dtz2.DateTime = time.Date(year, time.Month(month), day, hour, min, sec, nanosecond, loc)
	dtz2.TimeZone, dtz2.TimeZoneOffset  = dtz2.DateTime.Zone()
	dtz2.TimeLocName = dtz2.TimeLoc.String()

	dtz2.allocateNanoseconds(int64(nanosecond))

	return dtz2, nil
}

// NewDateTime - creates a new DateTzDto object and populates the data fields based on
// input parameters.
//
// Input Parameters
// ================
//
// year 						int			- year number
// month						int			- month number 	1 - 12
// day							int			- day number   	1 - 31
// hour							int			- hour number  	0 - 24
// min							int			- minute number	0 - 59
// sec							int			- second number	0	-	59
// millisecond			int			- millisecond number 0 - 999
// microsecond			int			-	microsecond number 0 - 999
// nanosecond				int			- nanosecond number 0 - 999
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
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type defined as follows:
//
//	type DateTzDto struct {
//		Year       			int							// Year Number
//		Month      			int							// Month Number
//		Day        			int							// Day Number
//		Hour       			int							// Hour Number
//		Minute     			int							// Minute Number
//		Second     			int							// Second Number
//		Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//		Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//		Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																		// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//		TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//		TimeZone   			string					// Time Zone associated with this Date Time. Example: "CDT" (abbreviation for Central Daylight Time)
//		TimeZoneOffset	int							// TimeZoneOffset associated with this Date Time
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		TimeLoc    			*time.Location	// Time Location pointer associated with this DateTime value
//		TimeLocName			string					// Time Location Name. Example: "America/Chicago"
//	}
//
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
// Usage
// =====
//
// Example:
//			dtzDto, err := DateTzDto{}.New(year, month, day, hour, min, sec, nanosecond , timeZoneLocation)
//
//
func (dtz DateTzDto) NewDateTime(year, month, day, hour, minute, second,
					millisecond, microsecond, nanosecond int, timeZoneLocation string) (DateTzDto, error) {

	ePrefix := "DateTzDto.New() "

	dtz2 := DateTzDto{}

	var err error

	if year < 0 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter year number is INVALID. 'year' must be greater than or equal to Zero. year='%v'", year)
	}

	if month < 1 || month > 12  {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter month number is INVALID. Correct range is 1-12. month='%v'", month)
	}


	if day < 1 || day > 31  {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter 'day' number is INVALID. Correct range is 1-31. day='%v'", day)
	}


	if hour < 0 || hour > 24 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter 'hour' number is INVALID. Correct range is 0-24. hour='%v'", hour)
	}

	if minute < 0 || minute > 59 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter minute number is INVALID. Correct range is 0 - 59. min='%v'", minute)
	}

	if second < 0 || second > 59 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parmeter second number is INVALID. Correct range is 0 - 59. second='%v'", second)
	}

	if millisecond < 0 || millisecond > 999 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter millisecond is INVALID. Correct range is 0 - 999. millisecond='%v'", millisecond)
	}

	if microsecond < 0 || microsecond > 999 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter microsecond is INVALID. Correct range is 0 - 999,999. microsecond='%v'", microsecond)
	}

	if nanosecond < 0 || nanosecond > 999 {
		return dtz2, fmt.Errorf(ePrefix + "Error: Input parameter nanosecond is INVALID. Correct range is 0 - 999. nanosecond='%v'", nanosecond)
	}

	dtz2.TimeLoc, err = time.LoadLocation(timeZoneLocation)

	if err != nil {
		return dtz2, fmt.Errorf("Error returned from time.LoadLocation(timeZoneLocation). 'timeZoneLocation' is INVALID. timeZoneLocation='%v'  Error='%v'", timeZoneLocation, err.Error())
	}

	dtz2.TimeLocName = dtz2.TimeLoc.String()

	dtz2.TotalNanoSecs =  int64(millisecond) * int64(time.Millisecond)
	dtz2.Millisecond = millisecond
	dtz2.TotalNanoSecs += int64(microsecond) * int64(time.Microsecond)
	dtz2.Microsecond = microsecond
	dtz2.TotalNanoSecs += int64(nanosecond)
	dtz2.Nanosecond = nanosecond

	dtz2.DateTime = time.Date(year, time.Month(month),day, hour, minute, second, int(dtz2.TotalNanoSecs), dtz2.TimeLoc)
	dtz2.TimeZone, dtz2.TimeZoneOffset = dtz2.DateTime.Zone()
	dtz2.Year = dtz2.DateTime.Year()
	dtz2.Month = int(dtz2.DateTime.Month())
	dtz2.Hour = dtz2.DateTime.Hour()
	dtz2.Minute = dtz2.DateTime.Minute()
	dtz2.Second = dtz2.DateTime.Second()
	return dtz2, nil

}
// allocateNanoseconds - allocates total Nanoseconds to milliseconds, microseconds
// and nanoseconds.
func (dtz *DateTzDto) allocateNanoseconds(totNanoseconds int64) {

	if totNanoseconds == 0 {
		dtz.TotalNanoSecs = 0
		dtz.Millisecond = 0
		dtz.Microsecond = 0
		dtz.Nanosecond = 0
		return
	}

	r := int(totNanoseconds)

	dtz.Millisecond = r / int(time.Millisecond)

	r -= dtz.Millisecond * int(time.Millisecond)

	if r == 0 {
		return
	}

	dtz.Microsecond = r / int(time.Microsecond)

	r -= dtz.Microsecond * int(time.Microsecond)

	dtz.Nanosecond = r

	dtz.TotalNanoSecs = totNanoseconds

	return
}


// CopyOut - returns a DateTzDto  instance
// which represents a deep copy of the current
// DateTzDto object.
func (dtz *DateTzDto) CopyOut() DateTzDto {
	dtz2 := DateTzDto{}

	dtz2.Year 					= dtz.Year
	dtz2.Month 					= dtz.Month
	dtz2.Day						= dtz.Day
	dtz2.Hour						= dtz.Hour
	dtz2.Minute					= dtz.Minute
	dtz2.Second					= dtz.Second
	dtz2.Millisecond		= dtz.Millisecond
	dtz2.Microsecond		= dtz.Microsecond
	dtz2.Nanosecond			= dtz.Nanosecond
	dtz2.TotalNanoSecs	= dtz.TotalNanoSecs

	if !dtz.DateTime.IsZero() {
		dtz2.DateTime = dtz.DateTime
		dtz2.TimeZone, dtz2.TimeZoneOffset = dtz2.DateTime.Zone()
		dtz2.TimeLoc = dtz2.DateTime.Location()
		dtz2.TimeLocName = dtz2.TimeLoc.String()
	} else {
		dtz2.TimeZone				= ""
		dtz2.TimeZoneOffset	= 0
		dtz2.DateTime				= time.Time{}
		dtz2.TimeLoc					= nil
		dtz2.TimeLocName			= ""
	}

	return dtz2
}

// Empty - sets all values of the current DateTzDto
// instance to their uninitialized or zero state.
func (dtz *DateTzDto) Empty() {

	dtz.Year 						= 0
	dtz.Month 					= 0
	dtz.Day							= 0
	dtz.Hour						= 0
	dtz.Minute					= 0
	dtz.Second					= 0
	dtz.Millisecond			= 0
	dtz.Microsecond			= 0
	dtz.Nanosecond			= 0
	dtz.TotalNanoSecs		= 0
	dtz.TimeZone				= ""
	dtz.TimeZoneOffset	= 0
	dtz.DateTime				= time.Time{}
	dtz.TimeLoc					= nil
	dtz.TimeLocName			= ""

	return
}

// TimeZoneUtility - Time Zone Data and Methods
type TimeZoneUtility struct {
	Description string
	TimeIn      time.Time
	TimeInLoc   *time.Location
	TimeOut     time.Time
	TimeOutLoc  *time.Location
	TimeUTC     time.Time
	TimeLocal		time.Time
}

// AddDate - Adds specified years, months and days values to the
// current time values maintained by this TimeZoneUtility
//
// Input Parameters
// ================
// years		int		- Number of years to add to current TimeZoneUtility instance
// months		int		- Number of months to add to current TimeZoneUtility instance
// days			int		- Number of months to add to current TimeZoneUtility instance
//
// Returns
// ======
// If successful, this method adds input date values to the current TimeZoneUtility.
//
// error	- If errors are encountered, this method returns an error object. Otherwise,
//					the error value is 'nil'.
//
func (tzu *TimeZoneUtility) AddDate(years, months, days int) error {

	ePrefix := "TimeZoneUtility.AddDate() "

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: This Time Zone Utility is INVALID!  Error='%v'", err.Error())
	}

	tzu.TimeIn = tzu.TimeIn.AddDate(years, months, days)
	tzu.TimeOut = tzu.TimeOut.AddDate(years, months, days)
	tzu.TimeUTC = tzu.TimeUTC.AddDate(years, months, days)
	tzu.TimeLocal = tzu.TimeLocal.AddDate(years, months, days)
	tzu.TimeInLoc = tzu.TimeIn.Location()
	tzu.TimeOutLoc = tzu.TimeOut.Location()

	return nil
}


// AddDateTime - Adds input time elements to the time
// value of the current TimeZoneUtility instance.
//
// Input Parameters
// ================
// years				int		- Number of years added to current TimeZoneUtility
// months				int		- Number of months added to current TimeZoneUtility
// days					int		- Number of days added to current TimeZoneUtility
// hours				int		- Number of hours added to current TimeZoneUtility
// minutes			int		- Number of minutes added to current TimeZoneUtility
// seconds			int		- Number of seconds added to current TimeZoneUtility
// milliseconds	int		- Number of milliseconds added to current TimeZoneUtility
// microseconds	int		- Number of microseconds added to current TimeZoneUtility
// nanoseconds	int		- Number of nanoseconds added to current TimeZoneUtility
//
// Note: 	Input parameters may be either negative or positive. Negative
// 				values will subtract time from the current TimeZoneUtility instance.
//
// Returns
// =======
//
// This method returns an error instance if errors are encountered. There
// are no other returns. If successful, the method updates
// the values of the current TimeZoneUtility instance.
//
func (tzu *TimeZoneUtility) AddDateTime(years, months, days, hours, minutes,
												seconds, milliseconds, microseconds, nanoseconds int) error {

	ePrefix := "TimeZoneUtility.AddDateTime() "

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "This TimeZoneUtility instance is INVALID! Error='%v'", err.Error())
	}

	err = tzu.AddDate(years, months, days)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tzu.AddDate(years, months, days). Error='%v'", err.Error())
	}

	err = tzu.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tzu.AddTime(...). Error='%v'", err.Error())
	}

	return nil
}

// AddDuration - Adds 'duration' to the time values maintained by the
// current TimeZoneUtility.
//
// Input Parameters
// ================
//
// duration		time.Duration		- May be a positive or negative duration
//															value which is added to the time value
//															of the current TimeZoneUtility instance.
//
func (tzu *TimeZoneUtility) AddDuration(duration time.Duration) error {

	ePrefix := "TimeZoneUtility.AddDuration() "

	if duration == 0 {
		return nil
	}

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "This TimeZoneUtility instance is INVALID! Error='%v'", err.Error())
	}

	tzu.TimeIn = tzu.TimeIn.Add(duration)
	tzu.TimeInLoc = tzu.TimeIn.Location()
	tzu.TimeOut = tzu.TimeOut.Add(duration)
	tzu.TimeOutLoc = tzu.TimeOut.Location()
	tzu.TimeUTC = tzu.TimeUTC.Add(duration)
	tzu.TimeLocal = tzu.TimeLocal.Add(duration)

	return nil
}

// AddTime - Adds time elements to the time value of the current
// TimeZoneUtility instance.
//
// Input Parameters:
// =================
//
// hours				- hours to be added to current TimeZoneUtility
// minutes			- minutes to be added to current TimeZoneUtility
// seconds			- seconds to be added to current TimeZoneUtility
// milliseconds	- milliseconds to be added to current TimeZoneUtility
// microseconds	- microseconds to be added to current TimeZoneUtility
// nanoseconds	- nanoseconds to be added to current TimeZoneUtility
//
// Note: 	Negative time values may be entered to subtract time from the
// 				current TimeZoneUtility time values.
//
// Returns
// =======
//
// If successful this method updates the time value fields of the current TimeZoneUtility instance.
//
// error - 	If errors are encountered, the returned 'error' object is populated. Otherwise, 'error'
//					is set to 'nil'.
//
func (tzu *TimeZoneUtility) AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds int) error {

	ePrefix := "TimeZoneUtility.AddTime() "

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "This TimeZoneUtility instance is INVALID! Error='%v'", err.Error())
	}

	if hours < 0  {
		return fmt.Errorf(ePrefix + "Error: Input parameter 'hours' number is INVALID. Correct range equal to or greater than Zero. hours='%v'", hours)
	}


	var totNanoSecs  int64

	totNanoSecs = int64(time.Hour) * int64(hours)
	totNanoSecs += int64(time.Minute) * int64(minutes)
	totNanoSecs += int64(time.Second) * int64(seconds)
	totNanoSecs += int64(time.Millisecond) * int64(milliseconds)
	totNanoSecs += int64(time.Microsecond) * int64(microseconds)
	totNanoSecs += int64(nanoseconds)

	err = tzu.AddDuration(time.Duration(totNanoSecs))

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tzu.AddDuration(time.Duration(totNanoSecs)). Error='%v'", err.Error())
	}

	return nil
}

// ConvertTz - Convert Time from existing time zone to targetTZone.
// The results are stored in the TimeZoneUtility data structure.
// The input time and output time are equivalent times adjusted
// for different time zones.
//
// Input Parameters:
//
// tIn 				time.Time 	- initial time values
// targetTz 	string			- time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
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
// Output Values are returned in the tzu (TimeZoneUtility)
// data fields. tzu.TimeOut contains the correct time in the 'target' time
// zone.
//
func (tzu TimeZoneUtility) ConvertTz(tIn time.Time, targetTz string) (TimeZoneUtility, error) {

	ePrefix := "TimeZoneUtility.ConvertTz() "

	tzuOut := TimeZoneUtility{}

	if isValidTz, _, _ := tzu.IsValidTimeZone(targetTz); !isValidTz {
		return tzuOut, errors.New(fmt.Sprintf("%v Error: targetTz is INVALID!! Input Time Zone == %v", ePrefix, targetTz))
	}

	if tIn.IsZero() {
		return tzuOut, errors.New(ePrefix + "Error: Input parameter time, 'tIn' is zero and INVALID")
	}

	tzOut, err := time.LoadLocation(targetTz)

	if err != nil {
		return tzuOut, fmt.Errorf("%vError Loading Target IANA Time Zone 'targetTz', %v. Errors: %v ",ePrefix, targetTz, err.Error())
	}


	tzuOut.SetTimeIn(tIn)

	tzuOut.SetTimeOut(tIn.In(tzOut))

	tzuOut.SetUTCTime(tIn)

	err = tzuOut.SetLocalTime(tIn)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.SetLocalTime(tIn). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// CopyOut - Creates and returns a deep copy of the
// current TimeZoneUtility instance.
func (tzu *TimeZoneUtility) CopyOut() TimeZoneUtility {
	tzu2 := TimeZoneUtility{}
	tzu2.Description = tzu.Description
	tzu2.TimeIn = tzu.TimeIn
	tzu2.TimeInLoc = tzu.TimeInLoc
	tzu2.TimeOut = tzu.TimeOut
	tzu2.TimeOutLoc = tzu.TimeOutLoc
	tzu2.TimeUTC = tzu.TimeUTC
	tzu2.TimeLocal = tzu.TimeLocal

	return tzu2
}

// CopyToThis - Copies another TimeZoneUtility
// to the current TimeZoneUtility data fields.
func (tzu *TimeZoneUtility) CopyToThis(tzu2 TimeZoneUtility) {
	tzu.Empty()

	tzu.Description = tzu2.Description
	tzu.TimeIn = tzu2.TimeIn
	tzu.TimeInLoc = tzu2.TimeInLoc
	tzu.TimeOut = tzu2.TimeOut
	tzu.TimeOutLoc = tzu2.TimeOutLoc
	tzu.TimeUTC = tzu2.TimeUTC
	tzu.TimeLocal = tzu2.TimeLocal
}

// Equal - returns a boolean value indicating
// whether two TimeZoneUtility data structures
// are equivalent.
func (tzu *TimeZoneUtility) Equal(tzu2 TimeZoneUtility) bool {
	if tzu.TimeIn != tzu2.TimeIn ||
		tzu.TimeInLoc != tzu2.TimeInLoc ||
		tzu.TimeOut != tzu2.TimeOut ||
		tzu.TimeOutLoc != tzu2.TimeOutLoc ||
		tzu.TimeUTC != tzu2.TimeUTC  ||
		tzu.TimeLocal != tzu2.TimeLocal	 {

		return false
	}

	return true
}

// Empty - Clears or returns this
// TimeZoneUtility to an uninitialized
// state.
func (tzu *TimeZoneUtility) Empty() {
	tzu.Description = ""
	tzu.TimeIn = time.Time{}
	tzu.TimeInLoc = nil
	tzu.TimeOut = time.Time{}
	tzu.TimeOutLoc = nil
	tzu.TimeUTC = time.Time{}
	tzu.TimeLocal = time.Time{}
}

// GetLocationIn - Returns the time zone location for the
// TimeInLoc data field which is part of the current TimeZoneUtility
// structure.
func (tzu *TimeZoneUtility) GetLocationIn() (string, error) {
	ePrefix := "TimeZoneUtility.GetLocationIn() "

	if tzu.TimeIn.IsZero() {
		return "", errors.New(ePrefix + "Error: TimeIn is Zero and Uninitialized!")
	}

	return tzu.TimeInLoc.String(), nil
}

// Get LocationOut - - Returns the time zone location for the
// TimeInLoc data field which is part of the current TimeZoneUtility
// structure.
func (tzu *TimeZoneUtility) GetLocationOut() (string, error) {

	ePrefix := "TimeZoneUtility.GetLocationOut() "

	if tzu.TimeOut.IsZero() {
		return "", errors.New(ePrefix + "Error: TimeOut is Zero and Uninitialized!")
	}

	return tzu.TimeOutLoc.String(), nil
}

// GetTimeInDto - returns a DateTzDto instance representing the value
// of the TimeIn data field for the current TimeZoneUtility.
func (tzu *TimeZoneUtility) GetTimeInDto() (DateTzDto, error) {

	ePrefix := "TimeZoneUtility) GetTimeInDto() "

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "This TimeZoneUtiltiy instance is INVALID! Error='%v'", err.Error())
	}

	dtzDto, err := DateTzDto{}.New(tzu.TimeIn)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(tzu.TimeIn). tzu.TimeIn='%v', Error='%v'", tzu.TimeIn, err.Error())
	}

	return dtzDto, nil
}

// GetTimeOutDto - returns a DateTzDto instance representing the value
// of the TimeOut data field for the current TimeZoneUtility.
func (tzu *TimeZoneUtility) GetTimeOutDto() (DateTzDto, error) {

	ePrefix := "TimeZoneUtility) GetTimeOutDto() "

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "This TimeZoneUtiltiy instance is INVALID! Error='%v'", err.Error())
	}

	dtzDto, err := DateTzDto{}.New(tzu.TimeOut)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(tzu.TimeOut). tzu.TimeOut='%v', Error='%v'", tzu.TimeOut, err.Error())
	}

	return dtzDto, nil
}

// GetTimeLocalDto - returns a DateTzDto instance representing the value
// of the TimeLocal data field for the current TimeZoneUtility.
func (tzu *TimeZoneUtility) GetTimeLocalDto() (DateTzDto, error) {

	ePrefix := "TimeZoneUtility) GetTimeLocalDto() "

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "This TimeZoneUtiltiy instance is INVALID! Error='%v'", err.Error())
	}

	dtzDto, err := DateTzDto{}.New(tzu.TimeLocal)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(tzu.TimeLocal). tzu.TimeLocal='%v', Error='%v'", tzu.TimeLocal, err.Error())
	}

	return dtzDto, nil
}

// GetTimeUtcDto - returns a DateTzDto instance representing the value
// of the TimeUTC data field for the current TimeZoneUtility.
func (tzu *TimeZoneUtility) GetTimeUtcDto() (DateTzDto, error) {

	ePrefix := "TimeZoneUtility) GetTimeLocalDto() "

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "This TimeZoneUtiltiy instance is INVALID! Error='%v'", err.Error())
	}

	dtzDto, err := DateTzDto{}.New(tzu.TimeUTC)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(tzu.TimeUTC). tzu.TimeUTC='%v', Error='%v'", tzu.TimeUTC, err.Error())
	}

	return dtzDto, nil
}


// GetZoneIn - Returns The Time Zone for the TimeIn
// data field contained in the current TimeZoneUtility
// structure.
func (tzu *TimeZoneUtility) GetZoneIn() (string, error) {

	ePrefix := "TimeZoneUtility.GetZoneIn() "

	if tzu.TimeOut.IsZero() {
		return "", errors.New(ePrefix + "Error: TimeOut is Zero and Uninitialized!")
	}

	tzZone, _ := tzu.TimeIn.Zone()

	return tzZone, nil

}

// GetZoneOut - Returns The Time Zone for the TimeOut
// data field contained in the current TimeZoneUtility
// structure.
func (tzu *TimeZoneUtility) GetZoneOut() (string, error) {

	ePrefix := "TimeZoneUtility.GetZoneOut() "

	if tzu.TimeOut.IsZero() {
		return "", errors.New(ePrefix + "Error: TimeOut is Zero and Uninitialized!")
	}

	tzZone, _ := tzu.TimeOut.Zone()

	return tzZone, nil

}

// IsTimeZoneUtilityValid - Analyzes the current TimeZoneUtility
// instance and returns an error if the instance is Invalid.
func (tzu *TimeZoneUtility) IsTimeZoneUtilityValid() error {

	ePrefix := "TimeZoneUtility.IsTimeZoneUtilityValid() "

	if tzu.TimeIn.IsZero() {
		return errors.New(ePrefix + "Error: TimeIn is Zero!")
	}

	if tzu.TimeOut.IsZero() {
		return errors.New(ePrefix + "Error: TimeOut is Zero!")
	}

	if tzu.TimeUTC.IsZero() {
		return errors.New(ePrefix + "Error: TimeUTC is Zero!")
	}

	if tzu.TimeLocal.IsZero() {
		return errors.New(ePrefix + "Error: TimeLocal is Zero!")
	}

	if tzu.TimeInLoc == nil {
		tzu.TimeInLoc = tzu.TimeIn.Location()
	}

	if tzu.TimeOutLoc == nil {
		tzu.TimeOutLoc = tzu.TimeOut.Location()
	}

	return nil

}

// IsValidTimeZone - Tests a Time Zone string and returns three boolean values
// designating whether the passed Time Zone string is:
// (1.) a valid time zone
// (2.) a valid iana time zone
// (3.) a valid Local time zone
func (tzu *TimeZoneUtility) IsValidTimeZone(tZone string) (isValidTz, isValidIanaTz, isValidLocalTz bool) {

	isValidTz = false

	isValidIanaTz = false

	isValidLocalTz = false

	if tZone == "" {
		return
	}

	if tZone == "Local" {
		isValidTz = true
		isValidLocalTz = true
		return
	}

	_, err := time.LoadLocation(tZone)

	if err != nil {
		return
	}

	isValidTz = true

	isValidIanaTz = true

	isValidLocalTz = false

	return

}


// New - Initializes and returns a new TimeZoneUtility object.
//
// Input Parameters
// ----------------
//
// tIn					 time.Time	- The input time object.
//
// tZoneOutLocation string	- The first input time value, 'tIn' will have its time zone
// 														changed to a new time zone location specified by this second
// 														parameter, 'tZoneOutLocation'. The new time associated with
// 														'tZoneOutLocation' is assigned to the TimeZoneUtility data
// 														field. The 'tZoneOutLocation' time zone location must be
// 														designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
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
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneUtility Type
//																(2) an Error type
//
// 	TimeZoneUtility - The two input parameters are used to populate and return
// 										a TimeZoneUtility structure:

//				type TimeZoneUtility struct {
//									Description string
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - 'tIn' value converted to TimeOut
//																							// 		based on 'timeZoneOutLocation' parameter
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value equivalent to TimeIn
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error object is populated
//					with an error message.
//
func (tzu TimeZoneUtility) New(tIn time.Time, timeZoneOutLocation string) (TimeZoneUtility, error) {

	tzuOut := TimeZoneUtility{}

	return tzuOut.ConvertTz(tIn, timeZoneOutLocation)
}

// NewAddDate - receives four parameters: a TimeZoneUtility 'tzuIn' and integer values for
// 'years', 'months' and 'days'.  The 'years', 'months' and 'days' values are added to the
// 'tzuIn' date time values and returned as a new TimeZoneUtility instance.
//
// Input Parameters
// ================
//
// years				int		- Number of years added to tzuIn value.
// months				int		- Number of months added to tzuIn value.
// days					int		- Number of days added to tzuIn value.
//
// Note: Negative date values may be used to subtract date values from the
// 			tzuIn value.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneUtility Type
//																(2) an Error type
//
//  TimeZoneUtility - 	The date input parameters are added to 'tzuIn to produce, populate and return
// 											a TimeZoneUtility structure defined as follows:
//
//				type TimeZoneUtility struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error object is populated
//					with an error message.
//
func (tzu TimeZoneUtility) NewAddDate(tzuIn TimeZoneUtility, years int, months int, days int) (TimeZoneUtility, error) {
	ePrefix := "TimeZoneUtility.NewAddDate()"

	err:= tzuIn.IsTimeZoneUtilityValid()

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error: Input parameter tzuIn (TimeZoneUtility) is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	if years == 0 && months == 0 && days == 0 {
		return tzuOut, nil
	}

	err = tzuOut.AddDate(years, months, days)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDate(years, months, days) years='%v' months='%v' days='%v'  Error='%v'",years, months, days, err.Error())
	}

	return tzuOut, nil
}

// NewAddDateTime - Receives a TimeZoneUtility input parameter, 'tzuIn'
// and returns a new TimeZoneUtility instance equal to 'tzuIn' plus the
// time value of the remaining time element input parameters.
//
// Input Parameters
// ================
// tzuIn				TimeZoneUtility - Base TimeZoneUtility object to
//																which time elements will be added.
// years				int		- Number of years added to 'tzuIn'
// months				int		- Number of months added to 'tzuIn'
// days					int		- Number of days added to 'tzuIn'
// hours				int		- Number of hours added to 'tzuIn'
// minutes			int		- Number of minutes added to 'tzuIn'
// seconds			int		- Number of seconds added to 'tzuIn'
// milliseconds	int		- Number of milliseconds added to 'tzuIn'
// microseconds	int		- Number of microseconds added to 'tzuIn'
// nanoseconds	int		- Number of nanoseconds added to 'tzuIn'
//
// Note: 	Input time element parameters may be either negative or positive.
// 				Negative values will subtract time from the returned TimeZoneUtility instance.
//
// Returns
// =======
//  There are two return values: 	(1) a TimeZoneUtility Type
//																(2) an Error type
//
// TimeZoneUtility - 	If successful, this method returns a valid,	populated TimeZoneUtility
// 										instance which is equal to the time value of 'tzuIn' plus the other
// 										input parameter date-time elements. The TimeZoneUtility structure
//										is defined as follows:
//
//				type TimeZoneUtility struct {
//									Description string
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error					 -  The method will return an 'error' object if errors
//										are encountered.
//
func (tzu TimeZoneUtility) NewAddDateTime(tzuIn TimeZoneUtility, years, months, days, hours, minutes,
seconds, milliseconds, microseconds, nanoseconds int) (TimeZoneUtility, error) {

	ePrefix := "TimeZoneUtility.NewAddDateTime() "

	err := tzuIn.IsTimeZoneUtilityValid()

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error: Input Parameter 'tzuIn' is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	err = tzuOut.AddDateTime(years, months, days, hours, minutes,
		seconds, milliseconds, microseconds, nanoseconds)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDateTime(...). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// NewAddDuration - receives two input parameters, a TimeZoneUtility 'tzuIn' and a
// time 'duration'. 'tzuIn' is adjusted for the specified 'duration' and the resulting
// new TimeZoneUtility is returned.
//
// Input Parameters
// ================
//
// tzuIn	TimeZoneUtility - The second parameter, 'duration', will be added
//													to this TimeZoneUtility.
//
// duration	time.Duration	- This duration value will be added to the
//													'tzuIn' input parameter to create, populate and
//													return a new updated TimeZoneUtility instance.
//
// Note: 	Input parameter 'duration' will accept both positive and negative values.
// 				Negative values will effectively subtract the duration from 'tzuIn' time
// 				values.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneUtility Type
//																(2) an Error type
//
//  TimeZoneUtility - 	The input parameter 'duration' is added to 'tzuIn to produce, populate and return
// 											a TimeZoneUtility structure:
//
//				type TimeZoneUtility struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error object is populated
//					with an error message.
//
func (tzu TimeZoneUtility) NewAddDuration(tzuIn TimeZoneUtility, duration time.Duration) (TimeZoneUtility, error) {
	ePrefix := "TimeZoneUtility.NewAddDuration() "

	err := tzuIn.IsTimeZoneUtilityValid()

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error: Input Parameter 'tzuIn' is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	err = tzuOut.AddDuration(duration)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDuration(duration). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// NewAddTime - returns a new TimeZoneUtility equivalent to the input TimeZoneUtility Plus time elements.
//
// Input Parameters:
// =================
//
// tzuIn TimeZoneUtility - 	The base TimeZoneUtility to which
//													time values will be added.
// hours				int				- Number of hours to be added to tzuIn
// minutes			int 			- Number of minutes to be added to tzuIn
// seconds			int 			- Number of seconds to be added to tzuIn
// milliseconds	int 			- Number of milliseconds to be added to tzuIn
// microseconds	int				- Number of microseconds to be added to tzuIn
// nanoseconds	int				- Number of nanoseconds to be added to tzuIn
//
// Note: Negative time values may be used to subtract time from 'tzuIn'.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneUtility Type
//																(2) an Error type
//
//  TimeZoneUtility - 	The time input parameters are added to 'tzuIn to produce, populate and return
// 											a TimeZoneUtility structure:
//
//				type TimeZoneUtility struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error object is populated
//					with an error message.
//
func (tzu TimeZoneUtility) NewAddTime(tzuIn TimeZoneUtility, hours, minutes, seconds, milliseconds, microseconds, nanoseconds int) (TimeZoneUtility, error) {

	ePrefix := "TimeZoneUtility.NewAddTime() "

	err := tzuIn.IsTimeZoneUtilityValid()

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error: Input Parameter 'tzuIn' is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	err = tzuOut.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf("Error returned by tzuOut.AddTime(...). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// NewTimeAddDate - returns a new TimeZoneUtility. The TimeZoneUtility is initialized
// with the 'tIn' time parameter. The 'TimeOut' data field will be set to the 'tIn'
// value and the time zone location specified by the second parameter, 'tZoneLocation'.
// The method will then add the remaining date element parameters to the new TimeZoneUtility
// instance and return it to the calling function.
//
// Input Parameters
// ================
// tIn			time.Time 			- Initial time value assigned to 'TimeIn' field
//														of the new TimeZoneUtility.
//
// tZoneOutLocation string	- The first input time value, 'tIn' will have its time zone
// 														changed to a new time zone location specified by this second
// 														parameter, 'tZoneOutLocation'. The new time associated with
// 														'tZoneOutLocation' is assigned to the TimeZoneUtility data
// 														field. The 'tZoneOutLocation' time zone location must be
// 														designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
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
// years				int		- Number of years added to initial TimeZoneUtility value.
// months				int		- Number of months added to initial TimeZoneUtility value.
// days					int		- Number of days added to initial TimeZoneUtility value.
//
// Note: Negative date values may be used to subtract date values from the
// 			initial TimeZoneUtility.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneUtility Type
//																(2) an Error type
//
//  TimeZoneUtility - 	The date input parameters are added to a TimeZoneUtility created from
//											input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneUtility
//											instance is then returned to the calling function. A TimeZoneUtility structure
//											is defined as follows:
//
//				type TimeZoneUtility struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error instance is populated
//					with an error message.
//
func (tzu TimeZoneUtility) NewTimeAddDate(tIn time.Time, tZoneOutLocation string, years,
																						months, days int) (TimeZoneUtility, error) {
	ePrefix := "TimeZoneUtility.NewTimeAddDate() "

	tzuOut, err := tzu.ConvertTz(tIn, tZoneOutLocation)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returne by tzu.ConvertTz(tIn, tZoneOutLocation). tIn='%v' tZoneOutLocation='%v'  Error='%v'", tIn, tZoneOutLocation, err.Error())
	}

	err = tzuOut.AddDate(years, months, days)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDate(years, months, days) years='%v' months='%v' days='%v' Error='%v'", years, months, days, err.Error())
	}

	return tzuOut, nil
}

// NewTimeAddDateTime - returns a new TimeZoneUtility. The TimeZoneUtility is initialized
// with the 'tIn' time parameter. The 'TimeOut' data field will be set to the 'tIn'
// value adjusted for the time zone location specified by the second parameter, 'tZoneLocation'.
// The method will then add the remaining date-time element parameters to the new TimeZoneUtility
// instance and return it to the calling function.
//
// Input Parameters
// ================
// tIn			time.Time 		- Initial time value assigned to 'TimeIn' field
//													of the new TimeZoneUtility.
//
// tZoneLocation string		- The first input time value, 'tIn' will have its time zone
// 													changed to a new time zone location specified by this second
// 													parameter, 'tZoneLocation'. This time zone location must be
// 													designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
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
// years				int		- Number of years added to initial TimeZoneUtility value.
// months				int		- Number of months added to initial TimeZoneUtility value.
// days					int		- Number of days added to initial TimeZoneUtility value.
// hours				int		- Number of hours to be added to initial TimeZoneUtility value.
// minutes			int		- Number of minutes to be added to initial TimeZoneUtility value.
// seconds			int 	- Number of seconds to be added to initial TimeZoneUtility value.
// milliseconds	int		- Number of milliseconds to be added to initial TimeZoneUtility value.
// microseconds	int		- Number of microseconds to be added to initial TimeZoneUtility value.
// nanoseconds	int 	- Number of nanoseconds to be added to initial TimeZoneUtility value.
//
// Note: Negative date-time values may be used to subtract date-time from the initial TimeZoneUtility.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneUtility Type
//																(2) an Error type
//
//  TimeZoneUtility - 	The date-time input parameters are added to a TimeZoneUtility created from
//											input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneUtility
//											instance is then returned to the calling function. A TimeZoneUtility structure
//											is defined as follows:
//
//				type TimeZoneUtility struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error instance is populated
//					with an error message.
//
func (tzu TimeZoneUtility) NewTimeAddDateTime(tIn time.Time, tZoneLocation string, years, months,
															days, hours, minutes, seconds, milliseconds, microseconds,
																	nanoseconds int) (TimeZoneUtility, error) {

	ePrefix := "TimeZoneUtility.NewTimeAddDateTime() "

	tzuOut, err := tzu.ConvertTz(tIn, tZoneLocation)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returne by tzu.ConvertTz(tIn, tZoneLocation). tIn='%v' tZoneLocation='%v'  Error='%v'", tIn, tZoneLocation, err.Error())
	}

	err = tzuOut.AddDateTime(years, months, days, hours, minutes, seconds, milliseconds,
														microseconds, nanoseconds)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDateTime(...)  Error='%v'", err.Error())
	}

	return tzuOut, nil
}


// NewTimeAddDuration - receives a 'tIn' time.Time parameter and a 'tZoneLocation' parameter
// which are used to construct an initial TimeZoneUtility instance. The 'TimeOut'
// data field of this initial TimeZoneUtility will contain the value of 'tIn'
// converted to a different time zone specified by 'tZoneLocation'.
//
// The 'duration' parameter will be added to this initial TimeZoneUtility and
// an updated TimeZoneUtility instance will be returned to the calling function.
//
// Input Parameters
// ================
// tIn				time.Time 	- Initial time value assigned to 'TimeIn' field
//													of the new TimeZoneUtility.
//
// tZoneLocation string		- The first input time value, 'tIn' will have its time zone
// 													changed to a new time zone location specified by this second
// 													parameter, 'tZoneLocation'. This time zone location must be
// 													designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
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
// duration			time.Duration	- an int64 duration value which is added to the date time
//							value of the initial TimeZoneUtility created from 'tIn' and 'tZoneLocation'.
//
// Note: Negative duration values may be used to subtract time duration from the initial TimeZoneUtility
// 			 date time values.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneUtility Type
//																(2) an Error type
//
//  TimeZoneUtility - 	The duration input parameter is added to a TimeZoneUtility created from
//											input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneUtility
//											instance is then returned to the calling function. A TimeZoneUtility structure
//											is defined as follows:
//
//				type TimeZoneUtility struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error instance is populated
//					with an error message.
//
func (tzu TimeZoneUtility) NewTimeAddDuration(tIn time.Time, tZoneLocation string, duration time.Duration) (TimeZoneUtility, error) {
	ePrefix := "TimeZoneUtility.NewTimeAddDuration() "

	tzuOut, err := tzu.ConvertTz(tIn, tZoneLocation)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returne by tzu.ConvertTz(tIn, tZoneLocation). tIn='%v' tZoneLocation='%v'  Error='%v'", tIn, tZoneLocation, err.Error())
	}

	err = tzuOut.AddDuration(duration)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDuration(duration). duration='%v'  Error='%v'",duration, err.Error())
	}

	return tzuOut, nil
}

// NewTimeAddTime - receives a 'tIn' time.Time parameter and a 'tZoneLocation' parameter
// which are used to construct an initial TimeZoneUtility instance. The 'TimeOut'
// data field of this initial TimeZoneUtility will contain the value of 'tIn'
// converted to a different time zone specified by 'tZoneLocation'.
//
// The remaining time parameters will be added to this initial TimeZoneUtility and
// the updated TimeZoneUtility will be returned to the calling function.
//
// Input Parameters
// ================
// tIn				time.Time 	- Initial time value assigned to 'TimeIn' field
//													of the new TimeZoneUtility.
//
// tZoneLocation string		- The first input time value, 'tIn' will have its time zone
// 													changed to a new time zone location specified by this second
// 													parameter, 'tZoneLocation'. This time zone location must be
// 													designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
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
// hours				int				- Number of hours to be added to initial TimeZoneUtility
// minutes			int 			- Number of minutes to be added to initial TimeZoneUtility
// seconds			int 			- Number of seconds to be added to initial TimeZoneUtility
// milliseconds	int 			- Number of milliseconds to be added to initial TimeZoneUtility
// microseconds	int				- Number of microseconds to be added to initial TimeZoneUtility
// nanoseconds	int				- Number of nanoseconds to be added to initial TimeZoneUtility
//
// Note: Negative time values may be used to subtract time from initial TimeZoneUtility.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneUtility Type
//																(2) an Error type
//
//  TimeZoneUtility - 	The time input parameters are added to a TimeZoneUtility created from
//											input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneUtility
//											instance is then returned to the calling function. A TimeZoneUtility structure
//											is defined as follows:
//
//				type TimeZoneUtility struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error instance is populated
//					with an error message.
//
func (tzu TimeZoneUtility) NewTimeAddTime(tIn time.Time, tZoneLocation string, hours, minutes, seconds, milliseconds, microseconds, nanoseconds int) (TimeZoneUtility, error) {

ePrefix := "TimeZoneUtility.NewTimeAddTime() "

	tzuOut, err := tzu.ConvertTz(tIn, tZoneLocation)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returne by tzu.ConvertTz(tIn, tZoneLocation). tIn='%v' tZoneLocation='%v'  Error='%v'", tIn, tZoneLocation, err.Error())
	}

	err = tzuOut.AddTime(hours, minutes, seconds, milliseconds,
		microseconds, nanoseconds)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddTime(...)  Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// ReclassifyTimeWithNewTz - Receives a valid time (time.Time) value and changes the existing time zone
// to that specified in the 'tZone' parameter. During this time reclassification operation, the time
// zone is changed but the time value remains unchanged.
// Input Parameters:
//
// tIn time.Time 					- initial time whose time zone will be changed to
//													second input parameter, 'tZoneLocation'
//
// tZoneLocation string		- The first input time value, 'tIn' will have its time zone
// 													changed to a new time zone location specified by this second
// 													parameter, 'tZoneLocation'. This time zone location must be
// 													designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
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
func (tzu *TimeZoneUtility) ReclassifyTimeWithNewTz(tIn time.Time, tZoneLocation string) (time.Time, error) {
	ePrefix := "TimeZoneUtility.ReclassifyTimeWithNewTz() "

	strTime := tzu.TimeWithoutTimeZone(tIn)

	if len(tZoneLocation) == 0 {
		return time.Time{}, errors.New(ePrefix + "Error: Time Zone Location, 'tZoneLocation', is an EMPTY string!")
	}

	if strings.ToLower(tZoneLocation) == "local" {
		tZoneLocation = "Local"
	}

	isValidTz, _, _ := tzu.IsValidTimeZone(tZoneLocation)

	if !isValidTz {
		return time.Time{}, fmt.Errorf(ePrefix + "Error: Input Time Zone Location is INVALID! tZoneLocation='%v'", tZoneLocation)
	}

	tzNew, err := time.LoadLocation(tZoneLocation)

	if err != nil {
		return time.Time{}, fmt.Errorf(ePrefix + "Error returned by time.Location('%v') - Error: %v", tZoneLocation, err.Error())
	}

	tOut, err := time.ParseInLocation(NeutralDateFmt, strTime, tzNew)

	if err != nil {
		return tOut, fmt.Errorf(ePrefix + "Error returned by time.Parse - Error: %v", err.Error())
	}

	return tOut, nil
}

// SetTimeIn - Assigns value to field 'TimeIn'
func (tzu *TimeZoneUtility) SetTimeIn(tIn time.Time) {
	tzu.TimeIn = tIn
	tzu.TimeInLoc = tIn.Location()
}

// SetTimeOut - Assigns value to field 'TimeOut'
func (tzu *TimeZoneUtility) SetTimeOut(tOut time.Time) {
	tzu.TimeOut = tOut
	tzu.TimeOutLoc = tOut.Location()
}

// SetUTCTime - Assigns UTC Time to field 'TimeUTC'
func (tzu *TimeZoneUtility) SetUTCTime(t time.Time) {

	tzu.TimeUTC = t.UTC()
}

// SetLocalTime - Assigns Local Time to field 'TimeLocal'
func (tzu *TimeZoneUtility) SetLocalTime(t time.Time) error {
	ePrefix := "TimeZoneUtility.SetLocalTime() "

	tzLocal, err := time.LoadLocation("Local")

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by time.LoadLocation(\"Local\") Error='%v'", err.Error())
	}

	tzu.TimeLocal = t.In(tzLocal)

	return nil
}

// Sub - Subtracts the input TimeZoneUtility from the current TimeZoneUtility
// and returns the duration. Duration is calculated as:
// 						tzu.TimeLocal.Sub(tzu2.TimeLocal)
//
// The TimeLocal field is used to compute duration for this method.
//
func (tzu *TimeZoneUtility) Sub(tzu2 TimeZoneUtility) (time.Duration, error) {

	ePrefix := "TimeZoneUtility.Sub() "

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix + "Error: Current TimeZoneUtility (tzu) is INVALID. Error='%v'", err.Error())
	}

	err = tzu2.IsTimeZoneUtilityValid()

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix + "Error: Input Parameter 'tzu2' is INVALID! Error='%v'", err.Error())
	}

	return tzu.TimeLocal.Sub(tzu2.TimeLocal), nil
}

// TimeWithoutTimeZone - Returns a Time String containing
// NO time zone. - When the returned string is converted to
// time.Time - in defaults to a UTC time zone.
func (tzu *TimeZoneUtility) TimeWithoutTimeZone(t time.Time) string {
	return t.Format(NeutralDateFmt)
}