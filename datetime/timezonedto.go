package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)


// TimeZoneDto
// ===========
//
// TimeZoneDto is part of the date time operations library. The source code repository
// for this file is located at:
//
//     https://github.com/MikeAustin71/datetimeopsgo.git
//
// This source code file is located at:
//
//     MikeAustin71\datetimeopsgo\datetime\timezonedto.go
//
//
// Overview and General Usage
// ==========================
//
// TimeZoneDto is used to convert, store and transport time zone information.
// The user will use this Type to convert time.Time, date time values, between
// differing time zones.
//
// In addition to generating a date time converted to a time zone specified
// by the user, this Type automatically generates equivalent date time values
// for Time Zone Locations 'Local' and 'UTC'.
//
// If you are unfamiliar with the concept of a Time Zone Location, reference
// 'https://golang.org/pkg/time/'. The concept of Time Zone Location is used
// extensively by Type TimeZoneDto. Time Zone location must be designated as
// one of two values.
//
//      (1) the string 'Local' - signals the designation of the local time zone
//          location for the host computer.
// 
//      (2) IANA Time Zone Location -
//          See https://golang.org/pkg/time/#LoadLocation
//          and https://www.iana.org/time-zones to ensure that
//          the IANA Time Zone Database is properly configured
//          on your system. Note: IANA Time Zone Data base is
//          equivalent to 'tz database'.
//           Examples:
//            "America/New_York"
//            "America/Chicago"
//            "America/Denver"
//            "America/Los_Angeles"
//            "Pacific/Honolulu"
//
// Dependencies
// ============
// 
//   DateTzDto    - datetzdto.go
//   TimeZoneDef  - timezonedef.go
//
//
// TimeZoneDto - Time Zone Data Transfer Object Type and Methods
// =============================================================
//
type TimeZoneDto struct {
	Description string      // Unused - available for tagging, classification or
	//   labeling.
	TimeIn      DateTzDto   // Original input time value
	TimeOut     DateTzDto   // TimeOut - 'TimeIn' value converted to TimeOut
	TimeUTC     DateTzDto   // TimeUTC (Universal Coordinated Time aka 'Zulu') value
	//   equivalent to TimeIn
	TimeLocal   DateTzDto   // TimeIn value converted to the 'Local' Time Zone Location.
	//   'Local' is the Time Zone Location used by the host computer.
	DateTimeFmt string      // Date Time Format String. This format string is used to format
	//   Date Time text displays. The Default format string is:
	//   "2006-01-02 15:04:05.000000000 -0700 MST"
	lock         sync.Mutex // Mutex used to manage thread safe operations
}

// AddDate - Adds specified years, months and days values to the
// current time values maintained by this TimeZoneDto
//
// Input Parameters
// ================
// years    int    - Number of years to add to current TimeZoneDto instance
// months   int    - Number of months to add to current TimeZoneDto instance
// days     int    - Number of months to add to current TimeZoneDto instance
//
// Note:  The date input parameter may be either negative
//        or positive. Negative values will subtract time
//        from the current TimeZoneDto instance.
//
// Returns
// ======
// There only one return: An 'error' type.
//
// error - If errors are encountered, this method returns an error object.
//         Otherwise, the error value is 'nil'.
//
func (tzDto *TimeZoneDto) AddDate(years, months, days int) error {

	ePrefix := "TimeZoneDto.AddDate() "

	tZoneUtil := timeZoneDtoUtility{}

	err := tZoneUtil.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: This Time Zone Utility is INVALID!  Error='%v'", err.Error())
	}

	tzDto.TimeIn, err = tzDto.TimeIn.AddDate(years, months, days, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tzDto.TimeIn.AddDate(years, months, days).\n" +
			"TimeIn.DateTime='%v' years='%v  months='%v' days='%v'\nError='%v'\n",
			tzDto.TimeIn.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr), years, months, days, err.Error())
	}

	tzDto.TimeOut, err = tzDto.TimeOut.AddDate(years, months, days, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tzDto.TimeOut.AddDate(years, months, days).\n" +
			"tzDto.TimeOut='%v' years='%v' months='%v' days='%v'\nError='%v'\n",
			tzDto.TimeOut.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr), years, months, days, err.Error())
	}

	tzDto.TimeUTC, err = tzDto.TimeUTC.AddDate(years, months, days, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tzDto.TimeUTC.AddDate(years, months, days).\n" +
			"tzDto.TimeUTC='%v' years='%v' months='%v' days='%v'\n" +
			"Error='%v'\n",
			tzDto.TimeUTC.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr), years, months, days, err.Error())
	}

	tzDto.TimeLocal, err = tzDto.TimeLocal.AddDate(years, months, days, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tzDto.TimeLocal.AddDate(years, months, days).\n" +
			"tzDto.TimeLocal='%v' years='%v' months='%v' days='%v'\nError='%v'\n",
			tzDto.TimeLocal.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr), years, months, days, err.Error())
	}

	return nil
}

// addDateTime - Adds input time elements to the time
// value of the current TimeZoneDto instance.
//
// Input Parameters
// ================
// years        int  - Number of years added to current TimeZoneDto
// months       int  - Number of months added to current TimeZoneDto
// days         int  - Number of days added to current TimeZoneDto
// hours        int  - Number of hours added to current TimeZoneDto
// minutes      int  - Number of minutes added to current TimeZoneDto
// seconds      int  - Number of seconds added to current TimeZoneDto
// milliseconds int  - Number of milliseconds added to current TimeZoneDto
// microseconds int  - Number of microseconds added to current TimeZoneDto
// nanoseconds  int  - Number of nanoseconds added to current TimeZoneDto
//
// Note:  Date Time input parameters may be either negative or positive.
//        Negative values will subtract time from the current TimeZoneDto
//        instance.
//
// Returns
// =======
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'
//
func (tzDto *TimeZoneDto) AddDateTime(years, months, days, hours, minutes,
	seconds, milliseconds, microseconds, nanoseconds int) error {

	ePrefix := "TimeZoneDto.addDateTime() "

	tZoneUtil := timeZoneDtoUtility{}

	err := tZoneUtil.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"This current TimeZoneDto instance is INVALID! Error='%v'", err.Error())
	}

	tzDto.TimeIn, err = tzDto.TimeIn.AddDateTime(years, months, days, hours, minutes,
		seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tzDto.TimeIn.addDateTime(years, months, days, hours, minutes, seconds, "+
			"milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt). "+
			"Error='%v'", err.Error())
	}

	tzDto.TimeOut, err = tzDto.TimeOut.AddDateTime(years, months, days, hours, minutes,
		seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tzDto.TimeOut.addDateTime(years, months, days, hours, minutes, seconds, "+
			"milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt). "+
			"Error='%v'", err.Error())
	}

	tzDto.TimeUTC, err = tzDto.TimeUTC.AddDateTime(years, months, days, hours, minutes,
		seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tzDto.TimeUTC.addDateTime(years, months, days, hours, minutes, seconds, "+
			"milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt). "+
			"Error='%v'", err.Error())
	}

	tzDto.TimeLocal, err = tzDto.TimeLocal.AddDateTime(years, months, days, hours, minutes,
		seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tzDto.TimeLocal.AddDate(years, months, days, hours, minutes, seconds, "+
			"milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt). "+
			"Error='%v'", err.Error())
	}

	return nil
}

// AddDuration - Adds 'duration' to the time values maintained by the
// current TimeZoneDto.
//
// Input Parameters
// ================
//
// duration  time.Duration  - May be a positive or negative duration
//                            value which is added to the time value
//                            of the current TimeZoneDto instance.
//
// Note:   The time.Duration input parameter may be either negative
//         or positive. Negative values will subtract time from
//         the current TimeZoneDt instance.
//
// Returns
// =======
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'
//
func (tzDto *TimeZoneDto) AddDuration(duration time.Duration) error {

	ePrefix := "TimeZoneDto.AddDuration() "

	if duration == 0 {
		return nil
	}

	tZoneUtil := timeZoneDtoUtility{}

	err := tZoneUtil.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"This current TimeZoneDto instance is INVALID! Error='%v'", err.Error())
	}

	tzDto.TimeIn, err = tzDto.TimeIn.AddDuration(duration, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from tzDto.TimeIn.AddDuration(duration). "+
			"tzDto.TimeIn.DateTime='%v'  Error='%v'",
			tzDto.TimeIn.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	tzDto.TimeOut, err = tzDto.TimeOut.AddDuration(duration, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from tzDto.TimeOut.AddDuration(duration).\n"+
			"tzDto.TimeOut.DateTime='%v'\nError='%v'\n",
			tzDto.TimeOut.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	tzDto.TimeUTC, err = tzDto.TimeUTC.AddDuration(duration, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from tzDto.TimeUTC.AddDuration(duration). "+
			"tzDto.TimeUTC.DateTime='%v'  Error='%v'",
			tzDto.TimeUTC.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	tzDto.TimeLocal, err = tzDto.TimeLocal.AddDuration(duration, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned from tzDto.TimeLocal.AddDuration(duration).\n"+
			"tzDto.TimeLocal.DateTime='%v'\nError='%v'\n",
			tzDto.TimeLocal.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return nil
}

// AddMinusTimeDto - This method receives a TimeDto input parameter. It
// then proceeds to convert all time components to negative values and
// subtracts those time components from the time values of the current
// TimeZoneDto.
//
// Input Parameters:
// =================
//
// timeDto TimeDto - A TimeDto type containing time components (i.e.
//          years, months, weeks, days, hours, minutes,
//          seconds etc.) to be subtracted from the current
//          TimeZoneDto.
//
//         type TimeDto struct {
//          Years          int // Number of Years
//          Months         int // Number of Months
//          Weeks          int // Number of Weeks
//          WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours          int // Number of Hours.
//          Minutes        int // Number of Minutes
//          Seconds        int // Number of Seconds
//          Milliseconds   int // Number of Milliseconds
//          Microseconds   int // Number of Microseconds
//          Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//         }
//
// Returns
// =======
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'
//
func (tzDto *TimeZoneDto) AddMinusTimeDto(timeDto TimeDto) error {

	ePrefix := "TimeZoneDto.AddMinusTimeDto() "

	dateTzIn := tzDto.TimeIn.CopyOut()

	err := dateTzIn.AddMinusTimeDtoToThis(timeDto)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by dateTzIn.AddMinusTimeDtoToThis(timeDto) "+
			"Error='%v'", err.Error())
	}

	timeZoneLocation := tzDto.TimeOut.GetOriginalTzName()

	fmtStr := tzDto.TimeOut.GetDateTimeFmt()

	tz2Dto, err := TimeZoneDto{}.NewDateTz(dateTzIn, timeZoneLocation, fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by TimeZoneDto{}.NewDateTz(dateTzIn, timeZoneLocation, fmtStr) "+
			"Error='%v'", err.Error())
	}

	tzDto.CopyIn(tz2Dto)

	return nil
}

// AddPlusTimeDto - This method receives a TimeDto input parameter. It
// then proceeds to convert all time components to positive values and
// adds those time components to the time values of the current TimeZoneDto.
//
// Input Parameters:
// =================
//
// timeDto  TimeDto - A TimeDto type containing time components (i.e.
//          years, months, weeks, days, hours, minutes,
//          seconds etc.) to be added to the current
//          TimeZoneDto.
//
//          type TimeDto struct {
//           Years          int // Number of Years
//           Months         int // Number of Months
//           Weeks          int // Number of Weeks
//           WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//           DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//           Hours          int // Number of Hours.
//           Minutes        int // Number of Minutes
//           Seconds        int // Number of Seconds
//           Milliseconds   int // Number of Milliseconds
//           Microseconds   int // Number of Microseconds
//           Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//           TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                    //  plus remaining Nanoseconds
//          }
//
// Returns
// =======
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'
//
func (tzDto *TimeZoneDto) AddPlusTimeDto(timeDto TimeDto) error {

	ePrefix := "TimeZoneDto.AddPlusTimeDto() "

	dateTzIn := tzDto.TimeIn.CopyOut()

	err := dateTzIn.AddPlusTimeDtoToThis(timeDto)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by dateTzIn.AddPlusTimeDtoToThis(timeDto) "+
			"Error='%v'", err.Error())
	}

	timeZoneLocation := tzDto.TimeOut.GetOriginalTzName()

	fmtStr := tzDto.TimeOut.GetDateTimeFmt()

	tz2Dto, err := TimeZoneDto{}.NewDateTz(dateTzIn, timeZoneLocation, fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by TimeZoneDto{}.NewDateTz(dateTzIn, timeZoneLocation, fmtStr)\n"+
			"Error='%v'\n", err.Error())
	}

	tzDto.CopyIn(tz2Dto)

	return nil
}

// AddTime - Adds time elements to the time value of the current
// TimeZoneDto instance.
//
// Input Parameters:
// =================
//
// hours        - hours to be added to current TimeZoneDto
// minutes      - minutes to be added to current TimeZoneDto
// seconds      - seconds to be added to current TimeZoneDto
// milliseconds - milliseconds to be added to current TimeZoneDto
// microseconds - microseconds to be added to current TimeZoneDto
// nanoseconds  - nanoseconds to be added to current TimeZoneDto
//
// Note:   The time component input parameter may be either negative
//         or positive. Negative values will subtract time from
//         the current TimeZoneDto instance.
//
// Returns
// =======
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'.
//
func (tzDto *TimeZoneDto) AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds int) error {

	ePrefix := "TimeZoneDto.AddTime() "

	tZoneUtil := timeZoneDtoUtility{}

	err := tZoneUtil.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"This TimeZoneDto instance is INVALID! Error='%v'", err.Error())
	}

	tzDto.TimeIn, err =
		tzDto.TimeIn.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tzDto.TimeIn.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt) "+
			"Error='%v'", err.Error())
	}

	tzDto.TimeOut, err =
		tzDto.TimeOut.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tzDto.TimeOut.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt) "+
			"Error='%v'", err.Error())
	}

	tzDto.TimeUTC, err =
		tzDto.TimeUTC.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tzDto.TimeUTC.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt) "+
			"Error='%v'", err.Error())
	}

	tzDto.TimeLocal, err =
		tzDto.TimeLocal.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tzDto.TimeLocal.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds, tzDto.DateTimeFmt) "+
			"Error='%v'", err.Error())
	}

	return nil
}

// AddTimeDurationDto - Adds time duration as expressed by input type 'TimeDurationDto'
// to the time values maintained by the current TimeZoneDto.
//
// Input Parameters
// ================
//
// durDto  TimeDurationDto  - Contains the time duration value
//                            to be added to the current TimeZoneDto.
//
// Returns
// =======
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'
//
func (tzDto *TimeZoneDto) AddTimeDurationDto(durDto TimeDurationDto) error {

	ePrefix := "TimeZoneDto.AddTimeDurationDto() "

	err := tzDto.AddDuration(durDto.TimeDuration)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tzDto.AddDuration(durDto.TimeDuration) "+
			"Error='%v' ", err.Error())
	}

	return nil
}

// ConvertTz - Converts 'tIn' Date Time from existing time zone to a 'targetTz'
// or target Time Zone. The results are stored and returned in a TimeZoneDto
// data structure.
//
// The input time and output time are equivalent times adjusted
// for different time zones.
//
// Input Parameters:
//
// tIn     time.Time  - initial time values
// targetTz  string   - time zone location must be designated as one of three
//                      types of time zones:
//
//               (1) the string 'Local' - signals the designation of the
//                   time zone location used by the host computer.
//
//               (2) IANA Time Zone Location -
//                   See https://golang.org/pkg/time/#LoadLocation
//                   and https://www.iana.org/time-zones to ensure that
//                   the IANA Time Zone Database is properly configured
//                   on your system. Note: IANA Time Zone Data base is
//                   equivalent to 'tz database'.
//                      Examples:
//                        "America/New_York"
//                        "America/Chicago"
//                        "America/Denver"
//                        "America/Los_Angeles"
//                        "Pacific/Honolulu"
//
//               (3) A Military Time Zone
//                   Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//        Examples:
//          "Alpha"   or "A"
//          "Bravo"   or "B"
//          "Charlie" or "C"
//          "Delta"   or "D"
//          "Zulu"    or "Z"
//
// dateTimeFmtStr string  - A date time format string which will be used
//                          to format and display 'dateTime'. Example:
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                          If 'dateTimeFmtStr' is submitted as an
//                          'empty string', a default date time format
//                          string will be applied. The default date time
//                          format string is:
//                          TZDtoDefaultDateTimeFormatStr =
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//
// There are two returns:
//             (1) A TimeZoneDto instance
//             (2) An error type
//
// (1) TimeZoneDto
//     If successful, this method creates a new TimeZoneDto,
//     populated with, TimeIn, TimeOut, TimeUTC and TimeLocal
//     date time values plus time zone information.
//
//     A TimeZoneDto structure is defined as follows:
//
//     type TimeZoneDto struct {
//      Description  string     // Unused - available for tagging, classification or
//                              //   labeling.
//      TimeIn       DateTzDto  // Original input time value
//      TimeOut      DateTzDto  // TimeOut - 'TimeIn' value converted to TimeOut
//      TimeUTC      DateTzDto  // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                              //   equivalent to TimeIn
//      TimeLocal    DateTzDto  // TimeIn value converted to the 'Local' Time Zone Location.
//                              //   'Local' is the Time Zone Location used by the host computer.
//      DateTimeFmt  string     // Date Time Format String. This format string is used to format
//                              //  Date Time text displays. The Default format string is:
//                              //   "2006-01-02 15:04:05.000000000 -0700 MST"
//     }
//
//
// (2) error - If errors are encountered, this method returns an error instance populated with
//             a valid 'error' message. If the method completes successfully the returned error
//             error type is set to 'nil'.
//
func (tzDto TimeZoneDto) ConvertTz(
	tIn time.Time,
	targetTz,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.ConvertTz() "
	var err error

	tzDtoOut := TimeZoneDto{}

	if isValidTz, _ := tzDto.IsValidTimeZone(targetTz); !isValidTz {
		return tzDtoOut,
		fmt.Errorf(ePrefix +
			"\nError: targetTz is INVALID!!\n" +
			"Input Time Zone == %v\n",  targetTz)
	}

	if tIn.IsZero() {
		return tzDtoOut,
		errors.New(ePrefix + "\nError: Input parameter time, 'tIn' is ZERO and INVALID\n")
	}

	var tzOutLocPtr *time.Location

	tzOutLocPtr, err = time.LoadLocation(targetTz)

	if err != nil {
		return tzDtoOut,
		fmt.Errorf(ePrefix +
			"\nError Loading Target IANA Time Zone\n" +
			"targetTz='%v'\n Errors: %v\n", targetTz, err.Error())
	}

	tzDtoOut.SetDateTimeFormatStr(dateTimeFmtStr)

	err = tzDtoOut.setTimeIn(tIn)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+
			"\nError returned by tzDtoOut.setTimeIn(tIn).\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzDtoOut.setTimeOut(tIn.In(tzOutLocPtr))

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+
			"\nError returned by tzDtoOut.setTimeOut(tIn.In(tzOutLocPtr)).\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzDtoOut.setUTCTime(tIn)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+
			"\nError returned by tzDtoOut.setUTCTime(tIn).\nError='%v'\n", err.Error())
	}

	err = tzDtoOut.setLocalTime(tIn)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+
			"\nError returned by tzDtoOut.SetLocalTime(tIn).\n" +
			"Error='%v'\n", err.Error())
	}

	return tzDtoOut, nil
}

// copyOut - Creates and returns a deep copy of the
// current TimeZoneDto instance.
//
// Input Parameters
// ================
// None
//
// Returns
// =======
// There is only one return: A TimeZoneDto instance
//
// A TimeZoneDto structure is defined as follows:
//
//     type TimeZoneDto struct {
//      Description  string     // Unused - available for tagging, classification or
//                              //   labeling.
//      TimeIn       DateTzDto  // Original input time value
//      TimeOut      DateTzDto  // TimeOut - 'TimeIn' value converted to TimeOut
//      TimeUTC      DateTzDto  // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                              //   equivalent to TimeIn
//      TimeLocal    DateTzDto  // TimeIn value converted to the 'Local' Time Zone Location.
//                              //   'Local' is the Time Zone Location used by the host computer.
//      DateTimeFmt  string     // Date Time Format String. This format string is used to format
//                              //  Date Time text displays. The Default format string is:
//                              //   "2006-01-02 15:04:05.000000000 -0700 MST"
//     }
//
func (tzDto *TimeZoneDto) CopyOut() TimeZoneDto {
	tzu2 := TimeZoneDto{}
	tzu2.Description = tzDto.Description
	tzu2.TimeIn = tzDto.TimeIn.CopyOut()
	tzu2.TimeOut = tzDto.TimeOut.CopyOut()
	tzu2.TimeUTC = tzDto.TimeUTC.CopyOut()
	tzu2.TimeLocal = tzDto.TimeLocal.CopyOut()
	tzu2.DateTimeFmt = tzDto.DateTimeFmt

	return tzu2
}

// CopyIn - Copies input parameter TimeZoneDto data fields
// into the current TimeZoneDto data fields.
// When the method completes, the current TimeZoneDto and
// the input parameter TimeZoneDto are equivalent.
//
// Input Parameters
// ================
//
// tzdto2 TimeZoneDto - A TimeZoneDto instance. The data
//           fields from this incoming TimeZoneDto
//           will be copied to the data fields
//           of the current TimeZoneDto.
//
// A TimeZoneDto structure is defined as follows:
//
//  type TimeZoneDto struct {
//   Description  string       // Unused - available for tagging, classification or
//                             //  labeling.
//   TimeIn       DateTzDto    // Original input time value
//   TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//   TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                             //   equivalent to TimeIn
//   TimeLocal   DateTzDto     // TimeIn value converted to the 'Local' Time Zone Location.
//                             //   'Local' is the Time Zone Location used by the host computer.
//   DateTimeFmt   string      // Date Time Format String. This format string is used to format
//                             //  Date Time text displays. The Default format string is:
//                             //   "2006-01-02 15:04:05.000000000 -0700 MST"
//  }
//
// Returns
// =======
//
//  None
//
func (tzDto *TimeZoneDto) CopyIn(tzdto2 TimeZoneDto) {

	tzDto.Empty()

	tzDto.Description = tzdto2.Description
	tzDto.TimeIn = tzdto2.TimeIn.CopyOut()
	tzDto.TimeOut = tzdto2.TimeOut.CopyOut()
	tzDto.TimeUTC = tzdto2.TimeUTC.CopyOut()
	tzDto.TimeLocal = tzdto2.TimeLocal.CopyOut()
	tzDto.DateTimeFmt = tzdto2.DateTimeFmt

}

// Equal - returns a boolean value indicating
// whether the current TimeZoneDto data structure
// is equivalent to the input parameter TimeZoneDto
// data structure.
//
// Input Parameter
// ===============
//
// tzdto2  TimeZoneDto - This input parameter TimeZoneDto
//                       is compared to the current TimeZoneDto
//                       to determine if they are equivalent.
//
// Return
// ======
//
// bool  - If the current TimeZoneDto is equivalent to the
//         input parameter TimeZoneDto, this method returns
//         'true'.
//
//         If the two TimeZoneDto's are NOT equivalent, this
//         method returns 'false'
//
func (tzDto *TimeZoneDto) Equal(tzdto2 TimeZoneDto) bool {

	if !tzDto.TimeIn.Equal(tzdto2.TimeIn) ||
		!tzDto.TimeOut.Equal(tzdto2.TimeOut) ||
		!tzDto.TimeUTC.Equal(tzdto2.TimeUTC) ||
		!tzDto.TimeLocal.Equal(tzdto2.TimeLocal) ||
		tzDto.Description != tzdto2.Description ||
		tzDto.DateTimeFmt != tzdto2.DateTimeFmt {

		return false
	}

	return true
}

// Empty - Clears or returns this
// TimeZoneDto to an uninitialized
// or 'Empty' state.
func (tzDto *TimeZoneDto) Empty() {
	tzDto.Description = ""
	tzDto.TimeIn = DateTzDto{}
	tzDto.TimeOut = DateTzDto{}
	tzDto.TimeUTC = DateTzDto{}
	tzDto.TimeLocal = DateTzDto{}
}

// IsValid - Analyzes the current TimeZoneDto
// instance and returns an error if the instance is INVALID.
//
func (tzDto *TimeZoneDto) IsValid() error {

	ePrefix := "TimeZoneDto.IsValid() "

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.isValidTimeZoneDto(tzDto, ePrefix)
}

// IsValidTimeZone - Tests a Time Zone Location string and
// returns two values:
//
// Input Parameters
// ============================================================================
//
// tZone   string  - The name of a valid time zone. This time zone must
//                   specify one of three types of time zones.
//
//                   (1) The 'Local' Time Zone.  Time Zone 'Local' specifies
//                       the time zone configured and applied on the host
//                       computer.
//
//                   (2) A IANA Time zone. This time must exist in the IANA
//                       database. The IANA database is widely recognized as
//                       a leading authority of global time zones.
//
//                       Reference:
//
//                          https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//                          https://en.wikipedia.org/wiki/Tz_database
//                          https://www.iana.org/time-zones
//
//                   (3) A single character or the full text name of a valid Military
//                       time zone. Military time zones are commonly used in aviation
//                       as well as at sea. They are also known as nautical or maritime
//                       time zones.
//                       Reference:
//                           https://en.wikipedia.org/wiki/List_of_military_time_zones
//                           http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                           https://www.timeanddate.com/time/zones/military
//
// Return Values
// ============================================================================
//
// isValidTz            bool - If true it signals the the time zone name contained in
//                             input parameter 'tZone' is a valid time zone.
//
// timeZoneType TimeZoneType - If return value 'isValidTz' is 'true', the 'timeZoneType'
//                             value will describe the valid time zone as one of three
//                             types: 'Local', 'IANA' or Military.
//
func (tzDto *TimeZoneDto) IsValidTimeZone(
	tZone string) (isValidTz bool, timeZoneType TimeZoneType) {

	isValidTz = false

	timeZoneType = TzType.None()

	if tZone == "" {
		return
	}

	ePrefix := "TimeZoneDto.IsValidTimeZone() "

	dTimeUtil := DTimeUtility{}

	var err error
	var tzSpec TimeZoneSpecification
	dateTime := time.Now().In(time.UTC)


	tzSpec,
		err = dTimeUtil.GetTimeZoneFromName(
			dateTime,
			tZone,
			TzConvertType.Absolute(),
			ePrefix)

	if err != nil {

		isValidTz = false

		timeZoneType = TzType.None()

		return isValidTz, timeZoneType
	}

	if tZone != tzSpec.GetLocationName() {

		isValidTz = false

		timeZoneType = TzType.None()

		return isValidTz, timeZoneType
	}

	isValidTz = true
	timeZoneType = tzSpec.GetTimeZoneType()

	return isValidTz, timeZoneType
}

// New - Initializes and returns a new TimeZoneDto object.
//
// Input Parameters
// ----------------
//
// tIn      time.Time   - The input time object.
//
// tZoneOutLocationName - string -
//                        The first input time value, 'tIn' will have its time zone
//                        changed to a new time zone location specified by this
//                        second parameter, 'tZoneOutLocationName'. The new time
//                        associated with 'tZoneOutLocationName' is assigned to
//                        the TimeZoneDto data field. The 'tZoneOutLocationName'
//                        time zone location must be designated as one of three
//                        types of time zones:
//
//                        (1) the string 'Local' - signals the designation of the
//                            time zone location used by the host computer.
//
//                        (2) IANA Time Zone Location -
//                           See https://golang.org/pkg/time/#LoadLocation
//                           and https://www.iana.org/time-zones to ensure that
//                           the IANA Time Zone Database is properly configured
//                           on your system. Note: IANA Time Zone Data base is
//                           equivalent to 'tz database'.
//                             Examples:
//                              "America/New_York"
//                              "America/Chicago"
//                              "America/Denver"
//                              "America/Los_Angeles"
//                              "Pacific/Honolulu"
//
//                        (3) A Military Time Zone
//                            Reference:
//                             https://en.wikipedia.org/wiki/List_of_military_time_zones
//                             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                             https://www.timeanddate.com/time/zones/military
//                             https://www.timeanddate.com/worldclock/timezone/alpha
//                             https://www.timeanddate.com/time/map/
//
//                            Examples:
//                              "Alpha"   or "A"
//                              "Bravo"   or "B"
//                              "Charlie" or "C"
//                              "Delta"   or "D"
//                              "Zulu"    or "Z"
//
// dateTimeFmtStr string  - A date time format string which will be used
//                to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                If 'dateTimeFmtStr' is submitted as an
//                'empty string', a default date time format
//                string will be applied. The default date time
//                format string is:
//                TZDtoDefaultDateTimeFormatStr =
//                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//  There are two return values:  (1) A TimeZoneDto Type
//                                (2) An Error type
//
//  (1) TimeZoneDto - The two input parameters are used to populate and return
//           a TimeZoneDto structure:
//
//  type TimeZoneDto struct {
//   Description  string     // Unused - available for tagging, classification or
//                           //  labeling.
//   TimeIn       DateTzDto  // Original input time value
//   TimeOut      DateTzDto  // TimeOut - 'TimeIn' value converted to TimeOut
//   TimeUTC      DateTzDto  // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                           //   equivalent to TimeIn
//   TimeLocal   DateTzDto   // TimeIn value converted to the 'Local' Time Zone Location.
//                           //   'Local' is the Time Zone Location used by the host computer.
//   DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                           //  Date Time text displays. The Default format string is:
//                           //   "2006-01-02 15:04:05.000000000 -0700 MST"
//  }
//
//
// (2) error - If the method completes successfully, the returned error instance is
//             set to nil. If errors are encountered, the returned error instance is populated
//             with an error message.
//
func (tzDto TimeZoneDto) New(
	tIn time.Time,
	timeZoneOutLocationName string,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.New() "
	tZoneUtil := timeZoneDtoUtility{}

	if tIn.IsZero() {
		return TimeZoneDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "tIn",
				inputParameterValue: "",
				errMsg:              "Input parameter 'tIn' has a ZERO value!",
				err:                 nil,
			}
	}

	if len(timeZoneOutLocationName) == 0 {
		return TimeZoneDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeZoneOutLocationName",
				inputParameterValue: "Input parameter 'timeZoneOutLocationName' is an EMPTY string!",
				errMsg:              "",
				err:                 nil,
			}
	}

	return tZoneUtil.newTzDto( tIn, timeZoneOutLocationName, dateTimeFmtStr, ePrefix)
}

// NewAddDate - receives four parameters: a TimeZoneDto 'tzuIn' and integer values for
// 'years', 'months' and 'days'.  The 'years', 'months' and 'days' values are added to the
// 'tzuIn' date time values and returned as a new TimeZoneDto instance.
//
// Input Parameters
// ================
//
// years     int  - Number of years added to tzuIn value.
// months    int  - Number of months added to tzuIn value.
// days      int  - Number of days added to tzuIn value.
//
// Note: Negative date values may be used to subtract date values from the
//    tzuIn value.
//
// dateTimeFmtStr string  - A date time format string which will be used
//                          to format and display 'dateTime'. Example:
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                          If 'dateTimeFmtStr' is submitted as an
//                          'empty string', a default date time format
//                          string will be applied. The default date time
//                          format string is:
//                          TZDtoDefaultDateTimeFormatStr =
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//  There are two return values:  (1) a TimeZoneDto Type
//                                (2) an Error type
//
//  (1) TimeZoneDto -  The date input parameters are added to 'tzuIn to produce,
//                     populate and return a TimeZoneDto structure defined as follows:
//
//    Description  string       // Unused - available for tagging, classification or
//                              //  labeling.
//    TimeIn       DateTzDto    // Original input time value
//
//   type TimeZoneDto struct {
//    TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//    TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                              //   equivalent to TimeIn
//    TimeLocal   DateTzDto     // TimeIn value converted to the 'Local' Time Zone Location.
//                              //   'Local' is the Time Zone Location used by the host computer.
//    DateTimeFmt   string      // Date Time Format String. This format string is used to format
//                              //  Date Time text displays. The Default format string is:
//                              //   "2006-01-02 15:04:05.000000000 -0700 MST"
//  }
//
//
// (2) error - If the method completes successfully, the returned error instance is
//       set to nil. If errors are encountered, the returned error object is
//        populated with an error message.
//
func (tzDto TimeZoneDto) NewAddDate(tzuIn TimeZoneDto, years, months, days int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewAddDate() "

	tZoneUtil := timeZoneDtoUtility{}

	err := tZoneUtil.isValidTimeZoneDto(&tzuIn, ePrefix)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+
			"\nError: Input parameter tzuIn (TimeZoneDto) is INVALID!\n" +
			"Error='%v'\n", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	tzuOut.SetDateTimeFormatStr(dateTimeFmtStr)

	err = tzuOut.AddDate(years, months, days)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+"Error returned by tzuOut.AddDate(years, months, days) years='%v' months='%v' days='%v'  Error='%v'", years, months, days, err.Error())
	}

	return tzuOut, nil
}

// NewAddDateTime - Receives a TimeZoneDto input parameter, 'tzuIn'
// and returns a new TimeZoneDto instance equal to 'tzuIn' plus the
// time value of the remaining time element input parameters.
//
// Input Parameters
// ================
// tzdtoIn   TimeZoneDto - Base TimeZoneDto object to
//                which time elements will be added.
// years    int  - Number of years added to 'tzuIn'
// months    int  - Number of months added to 'tzuIn'
// days     int  - Number of days added to 'tzuIn'
// hours    int  - Number of hours added to 'tzuIn'
// minutes   int  - Number of minutes added to 'tzuIn'
// seconds   int  - Number of seconds added to 'tzuIn'
// milliseconds int  - Number of milliseconds added to 'tzuIn'
// microseconds int  - Number of microseconds added to 'tzuIn'
// nanoseconds int  - Number of nanoseconds added to 'tzuIn'
//
// Note:  Input time element parameters may be either negative or positive.
//     Negative values will subtract time from the returned TimeZoneDto instance.
//
// dateTimeFmtStr string  - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               TZDtoDefaultDateTimeFormatStr =
//                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//  There are two return values:  (1) a TimeZoneDto Type
//                (2) an Error type
//
// (1) TimeZoneDto -  If successful, this method returns a valid, populated TimeZoneDto
//           instance which is equal to the time value of 'tzuIn' plus the other
//           input parameter date-time elements. The TimeZoneDto structure
//          is defined as follows:
//
//  type TimeZoneDto struct {
//   Description  string     // Unused - available for tagging, classification or
//                  //  labeling.
//   TimeIn       DateTzDto    // Original input time value
//   TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//   TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                  //   equivalent to TimeIn
//   TimeLocal   DateTzDto    // TimeIn value converted to the 'Local' Time Zone Location.
//                  //   'Local' is the Time Zone Location used by the host computer.
//   DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                  //  Date Time text displays. The Default format string is:
//                  //   "2006-01-02 15:04:05.000000000 -0700 MST"
//  }
//
// (2) error - If errors are encountered, this method returns an error instance populated with
//        a valid 'error' message. If the method completes successfully the returned error
//       error type is set to 'nil'.
//
func (tzDto TimeZoneDto) NewAddDateTime(tzdtoIn TimeZoneDto, years, months, days, hours, minutes,
	seconds, milliseconds, microseconds, nanoseconds int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewAddDateTime() "

	tZoneUtil := timeZoneDtoUtility{}

	err := tZoneUtil.isValidTimeZoneDto(&tzdtoIn, ePrefix)

	if err != nil {
		return TimeZoneDto{},
			fmt.Errorf(ePrefix+"Error: Input Parameter 'tzdtoIn' is INVALID! Error='%v'",
				err.Error())
	}

	tzuOut := tzdtoIn.CopyOut()

	tzuOut.SetDateTimeFormatStr(dateTimeFmtStr)

	err = tzuOut.AddDateTime(years, months, days, hours, minutes,
		seconds, milliseconds, microseconds, nanoseconds)

	if err != nil {
		return TimeZoneDto{},
			fmt.Errorf(ePrefix+"Error returned by tzuOut.addDateTime(...). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// NewAddDuration - receives two input parameters, a TimeZoneDto 'tzuIn' and a
// time 'duration'. 'tzuIn' is adjusted for the specified 'duration' and the resulting
// new TimeZoneDto is returned.
//
// Input Parameters
// ================
//
// tzdtoIn TimeZoneDto  - The second parameter, 'duration', will be added
//             to this TimeZoneDto.
//
// duration time.Duration - This duration value will be added to the
//             'tzuIn' input parameter to create, populate and
//             return a new updated TimeZoneDto instance.
//
// dateTimeFmtStr string  - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               TZDtoDefaultDateTimeFormatStr =
//                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Note:  Input parameter 'duration' will accept both positive and negative values.
//     Negative values will effectively subtract the duration from 'tzuIn' time
//     values.
//
// Returns
// =======
//  There are two return values:  (1) a TimeZoneDto Type
//                (2) an Error type
//
//  (1) TimeZoneDto - The input parameter 'duration' is added to 'tzuIn to produce, populate and return
//           a TimeZoneDto structure:
//
//  type TimeZoneDto struct {
//   Description  string     // Unused - available for tagging, classification or
//                  //  labeling.
//   TimeIn       DateTzDto    // Original input time value
//   TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//   TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                  //   equivalent to TimeIn
//   TimeLocal   DateTzDto    // TimeIn value converted to the 'Local' Time Zone Location.
//                  //   'Local' is the Time Zone Location used by the host computer.
//   DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                  //  Date Time text displays. The Default format string is:
//                  //   "2006-01-02 15:04:05.000000000 -0700 MST"
//  }
//
// (2) error - If errors are encountered, this method returns an error instance populated with
//        a valid 'error' message. If the method completes successfully the returned error
//       error type is set to 'nil'.
//
func (tzDto TimeZoneDto) NewAddDuration(
	tzdtoIn TimeZoneDto,
	duration time.Duration,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewAddDuration() "

	tZoneUtil := timeZoneDtoUtility{}

	err := tZoneUtil.isValidTimeZoneDto(&tzdtoIn, ePrefix)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+"Error: Input Parameter 'tzdtoIn' is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzdtoIn.CopyOut()

	tzuOut.SetDateTimeFormatStr(dateTimeFmtStr)

	err = tzuOut.AddDuration(duration)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+"Error returned by tzuOut.AddDuration(duration). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// NewAddTime - returns a new TimeZoneDto equivalent to the input TimeZoneDto Plus time elements.
//
// Input Parameters:
// =================
//
// tzdtoIn TimeZoneDto   - The base TimeZoneDto to which
//              time values will be added.
// hours    int    - Number of hours to be added to tzuIn
// minutes   int    - Number of minutes to be added to tzuIn
// seconds   int    - Number of seconds to be added to tzuIn
// milliseconds int    - Number of milliseconds to be added to tzuIn
// microseconds int    - Number of microseconds to be added to tzuIn
// nanoseconds int    - Number of nanoseconds to be added to tzuIn
//
// Note: Negative time values may be used to subtract time from 'tzuIn'.
//
// dateTimeFmtStr string - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//             If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              TZDtoDefaultDateTimeFormatStr =
//                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//  There are two return values:  (1) a TimeZoneDto Type
//                (2) an Error type
//
//  TimeZoneDto -  The time input parameters are added to 'tzuIn to produce, populate and return
//            a TimeZoneDto structure:
//
//    type TimeZoneDto struct {
//         Description string     // Unused. Available for tagging and classification.
//         TimeIn      time.Time    // Original input time value
//         TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//         TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//                       //   based on a specific Time Zone Location.
//
//         TimeOutLoc  *time.Location // Time Zone Location associated with TimeOut
//         TimeUTC     time.Time    // TimeUTC (Universal Coordinated Time) value
//                           equivalent to TimeIn
//
//         TimeLocal  time.Time    // Equivalent to TimeIn value converted to the 'Local'
//                       // Time Zone Location. 'Local' is the Time Zone Location
//                       //  used by the host computer.
//    }
//
// error - If the method completes successfully, the returned error instance is
//     set to nil. If errors are encountered, the returned error object is populated
//     with an error message.
//
func (tzDto TimeZoneDto) NewAddTime(tzdtoIn TimeZoneDto, hours, minutes, seconds, milliseconds, microseconds,
	nanoseconds int, dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewAddTime() "

	tZoneUtil := timeZoneDtoUtility{}

	err := tZoneUtil.isValidTimeZoneDto(&tzdtoIn, ePrefix)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+"Error: Input Parameter 'tzdtoIn' is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzdtoIn.CopyOut()

	tzuOut.SetDateTimeFormatStr(dateTimeFmtStr)

	err = tzuOut.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf("Error returned by tzuOut.AddTime(...). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// NewDateTz - Receives a DateTzDto instance and converts the DateTime associated with this
// DateTzDto instance to a TimeZoneDto instance. The DateTzDto.DateTime value is converted to
// TimeZoneDto.TimeOut using the input parameter, 'tZoneOutLocation', and returned as part
// of the newly created TimeZoneDto instance.
//
// Input Parameters
// ================
//
// dateTzDto  DateTzDto - Input parameter from which dateTzDto.DateTime
//            will be extracted to form the TimeZoneDto.TimeIn
//            value for the returned TimeZoneDto instance.
//
//   A DateTzDto structure is defined as follows:
//    type DateTzDto struct {
//     Description   string    // Unused, available for classification, labeling or description
//     Year          int       // Year Number
//     Month         int       // Month Number
//     Day           int       // Day Number
//     Hour          int       // Hour Number
//     Minute        int       // Minute Number
//     Second        int       // Second Number
//     Millisecond   int       // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//     Microsecond   int       // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//     Nanosecond    int       // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//                             // Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//     TotalNanoSecs int64     // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//     DateTime      time.Time // DateTime value for this DateTzDto Type
//     DateTimeFmt   string    // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//     TimeZone      TimeZoneDefinition // Contains a detailed description of the Time Zone and Time Zone Location
//                                  //  associated with this date time.
//    }
//
//
// tZoneOutLocation string -  The Time Zone Location to which input parameter
//              'dateTzDto.DateTime' will be converted.
//
//               Time Zone Out Location must be designated as one
//               of two values:
//
//               (1) the string 'Local' - signals the designation of the
//                 time zone location used by the host computer.
//
//              (2) IANA Time Zone Location -
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system. Note: IANA Time Zone Data base is
//                  equivalent to 'tz database'.
//                  Examples:
//                   "America/New_York"
//                   "America/Chicago"
//                   "America/Denver"
//                   "America/Los_Angeles"
//                   "Pacific/Honolulu"
//
// dateTimeFmtStr string  - A date time format string which will be used
//                          to format and display 'dateTime'. Example:
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                          If 'dateTimeFmtStr' is submitted as an
//                          'empty string', a default date time format
//                          string will be applied. The default date time
//                          format string is:
//                           TZDtoDefaultDateTimeFormatStr =
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
// There are two returns:
//             (1) A TimeZoneDto instance
//             (2) An error type
//
// (1) TimeZoneDto
//    If successful, this method creates a new TimeZoneDto,
//    populated with, TimeIn, TimeOut, TimeUTC and TimeLocal
//    date time values plus time zone information.
//
//    A TimeZoneDto structure is defined as follows:
//
//    type TimeZoneDto struct {
//     Description  string       // Unused - available for tagging, classification or
//                               //  labeling.
//     TimeIn       DateTzDto    // Original input time value
//     TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//     TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                               //   equivalent to TimeIn
//     TimeLocal   DateTzDto     // TimeIn value converted to the 'Local' Time Zone Location.
//                               //   'Local' is the Time Zone Location used by the host computer.
//     DateTimeFmt   string      // Date Time Format String. This format string is used to format
//                               //  Date Time text displays. The Default format string is:
//                               //   "2006-01-02 15:04:05.000000000 -0700 MST"
//    }
//
//
// (2) error - If errors are encountered, this method returns an error instance populated with
//             a valid 'error' message. If the method completes successfully the returned error
//             error type is set to 'nil'.
//
func (tzDto TimeZoneDto) NewDateTz(dateTzDtoIn DateTzDto, tZoneOutLocationName, dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewDateTz() "

	if err:=dateTzDtoIn.IsValid(); err!=nil {
		return TimeZoneDto{},
			fmt.Errorf(ePrefix + "\nError: Input parameter 'dateTzDtoIn' is INVALID!\n" +
				"Error='%v'\n", err.Error())
	}

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut, err := tZoneUtil.newTzDto(dateTzDtoIn.GetDateTimeValue(), tZoneOutLocationName, dateTimeFmtStr, ePrefix)

	if err != nil {
		return TimeZoneDto{},
			fmt.Errorf("New Time Zone Dto creation Failed!\n" +
				"%v", err.Error())
	}

	return tzuOut, nil
}

// NewTimeAddDate - returns a new TimeZoneDto. The TimeZoneDto is initialized
// with the 'tIn' time parameter. The 'TimeOut' data field will be set to the 'tIn'
// value and the time zone location specified by the second parameter, 'tZoneLocation'.
// The method will then add the remaining date element parameters to the new TimeZoneDto
// instance and return it to the calling function.
//
// Input Parameters
// ================
// tIn   time.Time    - Initial time value assigned to 'TimeIn' field
//              of the new TimeZoneDto.
//
// tZoneOutLocation string - The first input time value, 'tIn' will have its time zone
//               changed to a new time zone location specified by this second
//               parameter, 'tZoneOutLocation'. The new time associated with
//               'tZoneOutLocation' is assigned to the TimeZoneDto data
//               field. The 'tZoneOutLocation' time zone location must be
//               designated as one of two values:
//
//               (1) the string 'Local' - signals the designation of the
//                 time zone location used by the host computer.
//
//              (2) IANA Time Zone Location -
//                 See https://golang.org/pkg/time/#LoadLocation
//                 and https://www.iana.org/time-zones to ensure that
//                 the IANA Time Zone Database is properly configured
//                 on your system. Note: IANA Time Zone Data base is
//                 equivalent to 'tz database'.
//                Examples:
//                 "America/New_York"
//                 "America/Chicago"
//                 "America/Denver"
//                 "America/Los_Angeles"
//                 "Pacific/Honolulu"
//
// years    int  - Number of years added to initial TimeZoneDto value.
// months    int  - Number of months added to initial TimeZoneDto value.
// days     int  - Number of days added to initial TimeZoneDto value.
//
// Note: Negative date values may be used to subtract date values from the
//    initial TimeZoneDto.
//
// dateTimeFmtStr string  - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               TZDtoDefaultDateTimeFormatStr =
//                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//  There are two return values:  (1) a TimeZoneDto Type
//                (2) an Error type
//
//  TimeZoneDto -  The date input parameters are added to a TimeZoneDto created from
//           input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//           instance is then returned to the calling function. A TimeZoneDto structure
//           is defined as follows:
//
//    type TimeZoneDto struct {
//     Description  string     // Unused - available for tagging, classification or
//                    //  labeling.
//     TimeIn       DateTzDto    // Original input time value
//     TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//     TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                    //   equivalent to TimeIn
//     TimeLocal   DateTzDto    // TimeIn value converted to the 'Local' Time Zone Location.
//                    //   'Local' is the Time Zone Location used by the host computer.
//     DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                    //  Date Time text displays. The Default format string is:
//                    //   "2006-01-02 15:04:05.000000000 -0700 MST"
//    }
//
// error - If the method completes successfully, the returned error instance is
//     set to nil. If errors are encountered, the returned error instance is populated
//     with an error message.
//
func (tzDto TimeZoneDto) NewTimeAddDate(
	tIn time.Time,
	tZoneOutLocationName string,
	years,
	months,
	days int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewTimeAddDate() "

	// tzuOut, err := tzDto.ConvertTz(tIn, tZoneOutLocationName, dateTimeFmtStr)

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut,
	err := tZoneUtil.newTzDto(
		tIn,
		tZoneOutLocationName,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{},
		fmt.Errorf("Creation of TimeZoneDto, 'tzuOut', FAILED!\n" +
			"%v", err.Error())
	}

	err = tzuOut.AddDate(years, months, days)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+"Error returned by tzuOut.AddDate(years, months, days) years='%v' months='%v' days='%v' Error='%v'", years, months, days, err.Error())
	}

	return tzuOut, nil
}

// NewTimeAddDateTime - returns a new TimeZoneDto. The TimeZoneDto is initialized
// with the 'tIn' time parameter. The 'TimeOut' data field will be set to the 'tIn'
// value adjusted for the time zone location specified by the second parameter, 'tZoneLocation'.
// The method will then add the remaining date-time element parameters to the new TimeZoneDto
// instance and return it to the calling function.
//
// Input Parameters
// ================
// tIn   time.Time   - Initial time value assigned to 'TimeIn' field
//             of the new TimeZoneDto.
//
// tZoneLocation string  - The first input time value, 'tIn' will have its time zone
//              changed to a new time zone location specified by this second
//              parameter, 'tZoneLocation'. This time zone location must be
//              designated as one of two values:
//
//               (1) the string 'Local' - signals the designation of the
//                 time zone location used by the host computer.
//
//              (2) IANA Time Zone Location -
//                 See https://golang.org/pkg/time/#LoadLocation
//                 and https://www.iana.org/time-zones to ensure that
//                 the IANA Time Zone Database is properly configured
//                 on your system. Note: IANA Time Zone Data base is
//                 equivalent to 'tz database'.
//                Examples:
//                 "America/New_York"
//                 "America/Chicago"
//                 "America/Denver"
//                 "America/Los_Angeles"
//                 "Pacific/Honolulu"
//
// years    int  - Number of years added to initial TimeZoneDto value.
// months    int  - Number of months added to initial TimeZoneDto value.
// days     int  - Number of days added to initial TimeZoneDto value.
// hours    int  - Number of hours to be added to initial TimeZoneDto value.
// minutes   int  - Number of minutes to be added to initial TimeZoneDto value.
// seconds   int  - Number of seconds to be added to initial TimeZoneDto value.
// milliseconds int  - Number of milliseconds to be added to initial TimeZoneDto value.
// microseconds int  - Number of microseconds to be added to initial TimeZoneDto value.
// nanoseconds int  - Number of nanoseconds to be added to initial TimeZoneDto value.
//
// Note: Negative date-time values may be used to subtract date-time from the initial TimeZoneDto.
//
// dateTimeFmtStr string  - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               TZDtoDefaultDateTimeFormatStr =
//                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//  There are two return values:  (1) a TimeZoneDto Type
//                (2) an Error type
//
//  (1) TimeZoneDto -  The date-time input parameters are added to a TimeZoneDto created from
//           input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//           instance is then returned to the calling function. A TimeZoneDto structure
//           is defined as follows:
//
//    type TimeZoneDto struct {
//     Description  string     // Unused - available for tagging, classification or
//                    //  labeling.
//     TimeIn       DateTzDto    // Original input time value
//     TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//     TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                    //   equivalent to TimeIn
//     TimeLocal   DateTzDto    // TimeIn value converted to the 'Local' Time Zone Location.
//                    //   'Local' is the Time Zone Location used by the host computer.
//     DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                    //  Date Time text displays. The Default format string is:
//                    //   "2006-01-02 15:04:05.000000000 -0700 MST"
//    }
//
// (2) error - If the method completes successfully, the returned error instance is
//       set to nil. If errors are encountered, the returned error instance is populated
//       with an error message.
//
func (tzDto TimeZoneDto) NewTimeAddDateTime(
	tIn time.Time,
	tZoneLocation string,
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewTimeAddDateTime() "

	// tzuOut, err := tzDto.ConvertTz(tIn, tZoneLocation, dateTimeFmtStr)

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut, err := tZoneUtil.newTzDto(
		tIn,
		tZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{},
		fmt.Errorf("Creation of TimeZoneDto, 'tzuOut', FAILED!\n" +
			"%v", err.Error())
	}

	err = tzuOut.AddDateTime(years, months, days, hours, minutes, seconds, milliseconds,
		microseconds, nanoseconds)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix+"Error returned by tzuOut.addDateTime(...)  Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// NewTimeAddDuration - receives a 'tIn' time.Time parameter and a 'tZoneLocation' parameter
// which are used to construct an initial TimeZoneDto instance. The 'TimeOut'
// data field of this initial TimeZoneDto will contain the value of 'tIn'
// converted to a different time zone specified by 'tZoneLocation'.
//
// The 'duration' parameter will be added to this initial TimeZoneDto and
// an updated TimeZoneDto instance will be returned to the calling function.
//
// Input Parameters
// ================
// tIn    time.Time  - Initial time value assigned to 'TimeIn' field
//             of the new TimeZoneDto.
//
// tZoneLocation string  - The first input time value, 'tIn' will have its time zone
//              changed to a new time zone location specified by this second
//              parameter, 'tZoneLocation'. This time zone location must be
//              designated as one of two values:
//
//               (1) the string 'Local' - signals the designation of the
//                 time zone location used by the host computer.
//
//              (2) IANA Time Zone Location -
//                 See https://golang.org/pkg/time/#LoadLocation
//                 and https://www.iana.org/time-zones to ensure that
//                 the IANA Time Zone Database is properly configured
//                 on your system. Note: IANA Time Zone Data base is
//                 equivalent to 'tz database'.
//                Examples:
//                 "America/New_York"
//                 "America/Chicago"
//                 "America/Denver"
//                 "America/Los_Angeles"
//                 "Pacific/Honolulu"
//
// duration  time.Duration - An int64 duration value which is added to the date time
//              value of the initial TimeZoneDto created from 'tIn' and 'tZoneLocation'.
//
//               Note: Negative duration values may be used to subtract time duration
//               from the initial TimeZoneDto date time values.
//
// dateTimeFmtStr string  - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               TZDtoDefaultDateTimeFormatStr =
//                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//  There are two return values:  (1) a TimeZoneDto Type
//                (2) an Error type
//
//  (1) TimeZoneDto -  The duration input parameter is added to a TimeZoneDto created from
//           input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//           instance is then returned to the calling function.
//
//   A TimeZoneDto structure is defined as follows:
//
//    type TimeZoneDto struct {
//     Description  string     // Unused - available for tagging, classification or
//                    //  labeling.
//     TimeIn       DateTzDto    // Original input time value
//     TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//     TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                    //   equivalent to TimeIn
//     TimeLocal   DateTzDto    // TimeIn value converted to the 'Local' Time Zone Location.
//                    //   'Local' is the Time Zone Location used by the host computer.
//     DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                    //  Date Time text displays. The Default format string is:
//                    //   "2006-01-02 15:04:05.000000000 -0700 MST"
//    }
//
//
// (2) error - If errors are encountered, this method returns an error instance populated with
//        a valid 'error' message. If the method completes successfully the returned error
//       error type is set to 'nil'.
//
func (tzDto TimeZoneDto) NewTimeAddDuration(
	tIn time.Time,
	tZoneLocation string,
	duration time.Duration,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewTimeAddDuration() "

	tZoneUtil := timeZoneDtoUtility{}

	// tzuOut, err := tzDto.ConvertTz(tIn, tZoneLocation, dateTimeFmtStr)

	tzuOut, err := tZoneUtil.newTzDto(tIn, tZoneLocation, dateTimeFmtStr, ePrefix)

	if err != nil {
		return TimeZoneDto{},
		fmt.Errorf("Creation of a New Time Zone Dto FAILED!\n" +
			"%v", err.Error())
	}

	err = tzuOut.AddDuration(duration)

	if err != nil {
		return TimeZoneDto{},
		fmt.Errorf(ePrefix+
			"\nError returned by tzuOut.AddDuration(duration).\n" +
			"duration='%v'\n" +
			"Error='%v'\n", duration, err.Error())
	}

	return tzuOut, nil
}

// NewTimeAddTime - receives a 'tIn' time.Time parameter and a 'tZoneLocation' parameter
// which are used to construct an initial TimeZoneDto instance. The 'TimeOut'
// data field of this initial TimeZoneDto will contain the value of 'tIn'
// converted to a different time zone specified by 'tZoneLocation'.
//
// The remaining time parameters will be added to this initial TimeZoneDto and
// the updated TimeZoneDto will be returned to the calling function.
//
// Input Parameters
// ================
// tIn    time.Time  - Initial time value assigned to 'TimeIn' field
//             of the new TimeZoneDto.
//
// tZoneLocation string  - The first input time value, 'tIn' will have its time zone
//              changed to a new time zone location specified by this second
//              parameter, 'tZoneLocation'. This time zone location must be
//              designated as one of two values:
//
//               (1) the string 'Local' - signals the designation of the
//                 time zone location used by the host computer.
//
//              (2) IANA Time Zone Location -
//                 See https://golang.org/pkg/time/#LoadLocation
//                 and https://www.iana.org/time-zones to ensure that
//                 the IANA Time Zone Database is properly configured
//                 on your system. Note: IANA Time Zone Data base is
//                 equivalent to 'tz database'.
//                Examples:
//                 "America/New_York"
//                 "America/Chicago"
//                 "America/Denver"
//                 "America/Los_Angeles"
//                 "Pacific/Honolulu"
//
// hours    int    - Number of hours to be added to initial TimeZoneDto
// minutes   int    - Number of minutes to be added to initial TimeZoneDto
// seconds   int    - Number of seconds to be added to initial TimeZoneDto
// milliseconds int    - Number of milliseconds to be added to initial TimeZoneDto
// microseconds int    - Number of microseconds to be added to initial TimeZoneDto
// nanoseconds int    - Number of nanoseconds to be added to initial TimeZoneDto
//
//             Note: Negative time values may be used to subtract time from
//                initial TimeZoneDto.
//
// dateTimeFmtStr string  - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               TZDtoDefaultDateTimeFormatStr =
//                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  Returns
// =======
//  There are two return values:  (1) a TimeZoneDto Type
//                (2) an Error type
//
//  (1) TimeZoneDto -  The time input parameters are added to a TimeZoneDto created from
//           input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//           instance is then returned to the calling function.
//    A TimeZoneDto structure is defined as follows:
//
//    type TimeZoneDto struct {
//     Description  string     // Unused - available for tagging, classification or
//                    //  labeling.
//     TimeIn       DateTzDto    // Original input time value
//     TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//     TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                    //   equivalent to TimeIn
//     TimeLocal   DateTzDto    // TimeIn value converted to the 'Local' Time Zone Location.
//                    //   'Local' is the Time Zone Location used by the host computer.
//     DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                    //  Date Time text displays. The Default format string is:
//                    //   "2006-01-02 15:04:05.000000000 -0700 MST"
//    }
//
//
//  (2) error - If errors are encountered, this method returns an error instance populated with
//        a valid 'error' message. If the method completes successfully the returned error
//       error type is set to 'nil'.
//
func (tzDto TimeZoneDto) NewTimeAddTime(
	tIn time.Time,
	tZoneLocationName string,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewTimeAddTime() "

	// tzuOut, err := tzDto.ConvertTz(tIn, tZoneLocationName, dateTimeFmtStr)

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut, err := tZoneUtil.newTzDto(
		tIn,
		tZoneLocationName,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{},
		fmt.Errorf("Creation of new TimeZoneDto, 'tzuOut', FAILED!\n" +
			"%v", err.Error())
	}

	err = tzuOut.AddTime(
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds)

	if err != nil {
		return TimeZoneDto{},
		fmt.Errorf(ePrefix+
			"\nError returned by tzuOut.AddTime(...)\nError='%v'\n", err.Error())
	}

	return tzuOut, nil
}

// ReclassifyTimeWithNewTz - Receives a valid time (time.Time) value and changes the existing time zone
// to that specified in the 'tZone' parameter. During this time reclassification operation, the time
// zone is changed but the time value remains unchanged.
// Input Parameters:
//
// tIn time.Time      - Initial time whose time zone will be changed to
//                      second input parameter, 'tZoneLocation'
//
// tZoneLocation string  - The first input time value, 'tIn' will have its time zone
//                         changed to a new time zone location specified by this second
//                         parameter, 'tZoneLocation'. This time zone location must be
//                         designated as one of two values:
//
//               (1) the string 'Local' - signals the designation of the
//                 time zone location used by the host computer.
//
//              (2) IANA Time Zone Location -
//                 See https://golang.org/pkg/time/#LoadLocation
//                 and https://www.iana.org/time-zones to ensure that
//                 the IANA Time Zone Database is properly configured
//                 on your system. Note: IANA Time Zone Data base is
//                 equivalent to 'tz database'.
//
//                 Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
func (tzDto *TimeZoneDto) ReclassifyTimeWithNewTz(
	tIn time.Time,
	timeConversionType TimeZoneConversionType,
	tZoneLocation string) (time.Time, error) {

	ePrefix := "TimeZoneDto.ReclassifyTimeWithNewTz() "

	if len(tZoneLocation) == 0 {
		return time.Time{}, errors.New(ePrefix + "Error: Time Zone Location, 'tZoneLocation', is an EMPTY string!")
	}

	if strings.ToLower(tZoneLocation) == "local" {
		tZoneLocation = "Local"
	}

	dtUtil := DTimeUtility{}

	tzSpec,
	err := dtUtil.GetTimeZoneFromName(
		tIn,
		tZoneLocation,
		timeConversionType,
		ePrefix)

	if err != nil {
		return time.Time{}, err
	}

	return tzSpec.referenceDateTime, nil
}

// SetDateTimeFormatStr - Sets the value of the TimeZoneDto.DateTimeFmt field.
//
// Input Parameter
// ===============
//
//
// dateTimeFmtStr string  - A date time format string which will be used
//                          to format and display 'dateTime'. Example:
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                          If 'dateTimeFmtStr' is submitted as an
//                          'empty string', a default date time format
//                          string will be applied. The default date time
//                          format string is:
//                          TZDtoDefaultDateTimeFormatStr =
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tzDto *TimeZoneDto) SetDateTimeFormatStr(dateTimeFmtStr string) {

	if len(dateTimeFmtStr) == 0 {
		dateTimeFmtStr = FmtDateTimeYrMDayFmtStr
	}

	tzDto.DateTimeFmt = dateTimeFmtStr

}

// Sub - Subtracts the input TimeZoneDto from the current TimeZoneDto
// and returns the duration. Duration is calculated as:
//       tzu.TimeLocal.Sub(tzu2.TimeLocal)
//
// The TimeLocal field is used to compute duration for this method.
//
func (tzDto *TimeZoneDto) Sub(tzu2 TimeZoneDto) (time.Duration, error) {

	ePrefix := "TimeZoneDto.Sub() "

	tZoneUtil := timeZoneDtoUtility{}

	err := tZoneUtil.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix +
			"\nError: Current TimeZoneDto (tzDto) is INVALID.\n" +
			"Error='%v'\n", err.Error())
	}

	err = tZoneUtil.isValidTimeZoneDto(&tzu2, ePrefix)

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix+"Error: Input Parameter 'tzu2' is INVALID! Error='%v'", err.Error())
	}

	return tzDto.TimeLocal.Sub(tzu2.TimeLocal), nil
}

// TimeWithoutTimeZone - Returns a Time String containing
// NO time zone. - When the returned string is converted to
// time.Time, it will default to a UTC time zone.
func (tzDto *TimeZoneDto) TimeWithoutTimeZone(t time.Time) string {
	return t.Format(FmtDateTimeNeutralDateFmt)
}

// setTimeIn - Assigns time and zone values to field 'TimeIn'
func (tzDto *TimeZoneDto) setTimeIn(tIn time.Time) error {

	ePrefix := "TimeZoneDto.SetTimeIn() "
	var err error

	tzDto.TimeIn, err = DateTzDto{}.New(tIn, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by DateTzDto{}.New(tIn,tzDto.DateTimeFmt)\n" +
			"tIn='%v'\nError='%v'\n",
			tIn.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return nil
}

// setTimeOut - Assigns time and zone values to field 'TimeOut'
func (tzDto *TimeZoneDto) setTimeOut(tOut time.Time) error {

	ePrefix := "TimeZoneDto.setTimeOut()"

	var err error

	tzDto.TimeOut, err = DateTzDto{}.New(tOut, tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error retrned by DateTzDto{}.New(tOut,tzDto.DateTimeFmt). tOut='%v'  Error='%v'", tOut.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return nil

}

// setUTCTime - Assigns UTC Time and zone values to fields 'TimeUTC' and 'TimeUTCZone'
func (tzDto *TimeZoneDto) setUTCTime(t time.Time) error {

	ePrefix := "TimeZoneDto.setTimeOut()"

	var err error

	tzDto.TimeUTC, err = DateTzDto{}.New(t.UTC(), tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error retrned by DateTzDto{}.New(t.UTC(),tzDto.DateTimeFmt) tUTC='%v'  Error='%v'", t.UTC(), err.Error())
	}

	return nil

}

// setLocalTime - Assigns Local Time to field 'TimeLocal'
func (tzDto *TimeZoneDto) setLocalTime(t time.Time) error {
	ePrefix := "TimeZoneDto.SetLocalTime() "

	tzLocal, err := time.LoadLocation("Local")

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by time.LoadLocation(\"Local\") Error='%v'", err.Error())
	}

	tzDto.TimeLocal, err = DateTzDto{}.New(t.In(tzLocal), tzDto.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error retrned by DateTzDto{}.New(t.In(tzLocal),tzDto.DateTimeFmt). t.In(tzLocal)='%v'  Error='%v'", t.In(tzLocal).Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return nil
}
