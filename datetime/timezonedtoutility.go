package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type timeZoneDtoUtility struct {
	lock        sync.Mutex
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
func (tZoneUtil *timeZoneDtoUtility) addDateTime(
	tzDto *TimeZoneDto,
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.addDateTime() "

	if tzDto == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDto",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tzDto' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	tZoneUtil2 := timeZoneDtoUtility{}

	err := tZoneUtil2.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nThis current TimeZoneDto instance is INVALID!\n"+
			"Error='%v'\n", err.Error())
	}

	dtMech := DTimeMechanics{}

	newUtcDateTime := dtMech.AddDateTime(
		tzDto.TimeUTC.dateTimeValue,
		years,
		months,
		days,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds)

	tzDto2 := TimeZoneDto{}

	dateTimeFormat := tZoneUtil2.preProcessDateFormatStr(tzDto.DateTimeFmt)

	err = tZoneUtil2.setTimeInTzDef(
		&tzDto2,
		newUtcDateTime,
		TzConvertType.Relative(),
		tzDto.TimeIn.GetTimeZoneDef(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setTimeOutFromTimeZoneDef(
		&tzDto2,
		newUtcDateTime,
		TzConvertType.Relative(),
		tzDto.TimeOut.GetTimeZoneDef(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}


	err = tZoneUtil2.setUTCTime(
		&tzDto2,
		newUtcDateTime,
		TzConvertType.Absolute(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setLocalTime(
		&tzDto2,
		newUtcDateTime,
		TzConvertType.Relative(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	tzDto2.DateTimeFmt = dateTimeFormat

	tZoneUtil2.copyIn(tzDto, &tzDto2, ePrefix)

	return nil
}

// addDuration - Adds 'duration' to the time values maintained by the
// by input parameter 'tzDto' and instance of type 'TimeZoneDto'.
//
// Input Parameters
// ================
//
// duration  time.Duration
//         - May be a positive or negative duration
//           value which is added to the time value
//           of the current TimeZoneDto instance.
//
// Note:   The time.Duration input parameter may be either negative
//         or positive. Negative values will subtract time from
//         the current TimeZoneDt instance.
//
// Returns
// =======
//
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'
//
func (tZoneUtil *timeZoneDtoUtility) addDuration(
	tzDto *TimeZoneDto,
	duration time.Duration,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.addDuration() "

	if tzDto == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDto",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tzDto' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	timeZoneUtil2 := timeZoneDtoUtility{}

	err := timeZoneUtil2.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDto",
			inputParameterValue: "",
			errMsg:              fmt.Sprintf(
				"Input Parameter 'tzDto' is INVALID!\n" +
					"Validation Error='%v'", err.Error()),
			err:                 nil,
		}
	}

	dateTimeFormat := timeZoneUtil2.preProcessDateFormatStr(tzDto.DateTimeFmt)

	dTzUtil := dateTzDtoUtility{}

	var dateTzIn DateTzDto

	dateTzIn, err = dTzUtil.addDuration(
		&tzDto.TimeIn,
		duration,
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	tzDto2 := TimeZoneDto{}

	tZoneUtil2 := timeZoneDtoUtility{}

	err = tZoneUtil2.setTimeIn(
		&tzDto2,
		dateTzIn.dateTimeValue,
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setTimeOutFromTimeZoneDef(
		&tzDto2,
		dateTzIn.dateTimeValue,
		TzConvertType.Relative(),
		tzDto.TimeOut.GetTimeZoneDef(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setUTCTime(
		&tzDto2,
		dateTzIn.dateTimeValue,
		TzConvertType.Relative(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setLocalTime(
		&tzDto2,
		dateTzIn.dateTimeValue,
		TzConvertType.Relative(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	tzDto2.DateTimeFmt = dateTimeFormat

	tZoneUtil2.copyIn(tzDto, &tzDto2, ePrefix)

	return nil
}

// addMinusTimeDto - This method receives a TimeDto input parameter. It
// then proceeds to convert all time components to negative values and
// subtracts those time components from the time values of the current
// TimeZoneDto.
//
// Input Parameters:
// =================
//
// tzDto  *TimeZoneDto - The Time Zone Dto object from which
//                       'timeDto' parameter will be subtracted.
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
func (tZoneUtil *timeZoneDtoUtility) addMinusTimeDto(
	tzDto *TimeZoneDto,
	timeDto TimeDto,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.addMinusTimeDto() "

	if tzDto == nil {
		return &InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "tzDto",
				inputParameterValue: "",
				errMsg:              "Input Parameter 'tzDto' (TimeZoneDto) is a 'nil' pointer!",
				err:                 nil,
			}
	}

	tZoneUtil2 := timeZoneDtoUtility{}

	dateTimeFormat := tZoneUtil2.preProcessDateFormatStr(tzDto.DateTimeFmt)

	err := tZoneUtil2.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nInput Parameter 'tzDto' is INVALID!\n" +
			"Validation Error='%v'", err.Error())
	}

	err = timeDto.IsValid()

	if err != nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeDto",
			inputParameterValue: "",
			errMsg:              fmt.Sprintf(
				"Input prameter 'timeDto' is invalid!\n" +
					"%v\n", err.Error()),
			err:                 nil,
		}
	}

	dTzUtil := dateTzDtoUtility{}

	var dateTzIn DateTzDto
	tIn := tzDto.TimeIn

	dateTzIn, err = dTzUtil.addMinusTimeDto(
		&tIn,
		timeDto,
		ePrefix)

	if err != nil {
		return err
	}

	tzDto2 := TimeZoneDto{}

	err = tZoneUtil2.setTimeIn(
		&tzDto2,
		dateTzIn.dateTimeValue,
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setTimeOutFromTimeZoneDef(
		&tzDto2,
		dateTzIn.dateTimeValue,
		TzConvertType.Relative(),
		dateTzIn.GetTimeZoneDef(),
		tzDto.DateTimeFmt,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setUTCTime(
		&tzDto2,
		dateTzIn.dateTimeValue,
		TzConvertType.Relative(),
		tzDto.DateTimeFmt,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setLocalTime(
		&tzDto2,
		dateTzIn.dateTimeValue,
		TzConvertType.Relative(),
		tzDto.DateTimeFmt,
		ePrefix)

	if err != nil {
		return err
	}

	tZoneUtil2.copyIn(tzDto, &tzDto2, ePrefix)

	return nil
}

// addPlusTimeDto - This method receives a TimeDto input parameter. It
// then proceeds to convert all time components to positive values and
// adds those time components from the time values of the current
// TimeZoneDto.
//
// Input Parameters:
// =================
//
// tzDto  *TimeZoneDto - The Time Zone Dto object from which
//                       'timeDto' parameter will be subtracted.
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
func (tZoneUtil *timeZoneDtoUtility) addPlusTimeDto(
	tzDto *TimeZoneDto,
	timeDto TimeDto,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.addPlusTimeDto() "

	if tzDto == nil {
		return &InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "tzDto",
				inputParameterValue: "",
				errMsg:              "Input Parameter 'tzDto' (TimeZoneDto) is a 'nil' pointer!",
				err:                 nil,
			}
	}

	err := timeDto.IsValid()

	if err != nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeDto",
			inputParameterValue: "",
			errMsg:              fmt.Sprintf(
				"Input prameter 'timeDto' is invalid!\n" +
					"%v\n", err.Error()),
			err:                 nil,
		}
	}

	tZoneUtil2 := timeZoneDtoUtility{}

	dateTimeFormat := tZoneUtil2.preProcessDateFormatStr(tzDto.DateTimeFmt)

	err = tZoneUtil2.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nInput Parameter 'tzDto' is INVALID!\n" +
			"Validation Error='%v'", err.Error())
	}

	var dateTzIn DateTzDto

	tIn := tzDto.TimeIn.CopyOut()
	dTzUtil := dateTzDtoUtility{}

	dateTzIn, err = dTzUtil.addPlusTimeDto(
		&tIn,
		timeDto,
		ePrefix)

	if err != nil {
		return err
	}

	tzDto2 := TimeZoneDto{}

	err = tZoneUtil2.setTimeIn(
		&tzDto2,
		dateTzIn.dateTimeValue,
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setTimeOutFromTimeZoneDef(
		&tzDto2,
		dateTzIn.dateTimeValue,
		TzConvertType.Relative(),
		dateTzIn.GetTimeZoneDef(),
		tzDto.DateTimeFmt,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setUTCTime(
		&tzDto2,
		dateTzIn.dateTimeValue,
		TzConvertType.Relative(),
		tzDto.DateTimeFmt,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setLocalTime(
		&tzDto2,
		dateTzIn.dateTimeValue,
		TzConvertType.Relative(),
		tzDto.DateTimeFmt,
		ePrefix)

	if err != nil {
		return err
	}

	tZoneUtil2.copyIn(tzDto, &tzDto2, ePrefix)

	return nil
}

// Adds time to TimeZoneDto parameter, 'tzDto'.
//
func (tZoneUtil *timeZoneDtoUtility) addTime(
	tzDto *TimeZoneDto,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.addTime() "

	if tzDto == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDto",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tzDto' (TimeZoneDto) is a 'nil' pointer!",
			err:                 nil,
		}
	}

	tZoneUtil2 := timeZoneDtoUtility{}

	err := tZoneUtil2.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nInput Parameter 'TimeZoneDto' is INVALID!\n" +
			"Validation Error='%v'", err.Error())
	}

	dtMech := DTimeMechanics{}

	newDateTime := dtMech.AddDateTime(
		tzDto.TimeUTC.dateTimeValue,
		0,
		0,
		0,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds)

	tzDto2 := TimeZoneDto{}

	dateTimeFormat := tZoneUtil2.preProcessDateFormatStr(tzDto.DateTimeFmt)

	err = tZoneUtil2.setTimeInTzDef(
		&tzDto2,
		newDateTime,
		TzConvertType.Relative(),
		tzDto.TimeIn.GetTimeZoneDef(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}


	err = tZoneUtil2.setTimeOutFromTimeZoneDef(
		&tzDto2,
		newDateTime,
		TzConvertType.Relative(),
		tzDto.TimeOut.GetTimeZoneDef(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}


	err = tZoneUtil2.setUTCTime(
		&tzDto2,
		newDateTime,
		TzConvertType.Absolute(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	err = tZoneUtil2.setLocalTime(
		&tzDto2,
		newDateTime,
		TzConvertType.Relative(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	tzDto2.DateTimeFmt = dateTimeFormat

	tZoneUtil2.copyIn(tzDto, &tzDto2, ePrefix)

	return nil
}

// addTimeDurationDto - Adds time duration as expressed by input type 'TimeDurationDto'
// to the time values passed in 'tzDto'.
//
func (tZoneUtil *timeZoneDtoUtility) addTimeDurationDto(
	tzDto *TimeZoneDto,
	durDto TimeDurationDto,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.addTimeDurationDto() "

	if tzDto == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDto",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tzDto' is a 'nil' pointer!",
			err:                 nil,
		}
	}
	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.isValid(&durDto, ePrefix)

	if err != nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "durDto",
			inputParameterValue: "",
			errMsg:              fmt.Sprintf(
				"Input Parameter 'durDto' is INVALID!\n" +
					"Validation Error='%v'", err.Error()),
			err:                 nil,
		}
	}

	tZoneUtil2 := timeZoneDtoUtility{}

	return tZoneUtil2.addDuration(
		tzDto,
		durDto.TimeDuration,
		ePrefix)
}

// convertTz - Converts 'tIn' Date Time from existing time zone to a 'targetTz'
// or target Time Zone. The results are stored and returned in a new
// TimeZoneDto data structure. The current TimeZoneDto is NOT changed.
//
// The input time and output time are equivalent times adjusted for
// different time zones.
//
// Input Parameters:
//
// tIn     time.Time  - initial time values
//
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
func (tZoneUtil *timeZoneDtoUtility) convertTz(
	tzDto *TimeZoneDto,
	tIn time.Time,
	targetTimeZoneName,
	dateTimeFmtStr,
	ePrefix string) (TimeZoneDto, error) {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.convertTz() "

	if tzDto == nil {
		return TimeZoneDto{},
		&InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDto",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tzDto' (TimeZoneDto) is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tIn.IsZero() {
		return TimeZoneDto{},
		&InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tIn",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tIn' has a zero value!",
			err:                 nil,
		}
	}

	tZoneUtil2 := timeZoneDtoUtility{}

	isValidTimeZone,
	_,
	err :=
	tZoneUtil2.isValidTimeZoneName(targetTimeZoneName, ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	if !isValidTimeZone {
		return TimeZoneDto{},
			errors.New(ePrefix +
				"\nError: 'targetTimeZoneName' is an invalid time zone!\n")
	}

	tzDtoOut := TimeZoneDto{}

	dateTimeFmtStr = tZoneUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	tzDtoOut.DateTimeFmt = dateTimeFmtStr

	err = tZoneUtil2.setTimeIn(
		&tzDtoOut,
		tIn,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	err = tZoneUtil2.setTimeOutTz(
		&tzDtoOut,
		tIn,
		TzConvertType.Relative(),
		targetTimeZoneName,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	err = tZoneUtil2.setUTCTime(
		&tzDtoOut,
		tIn,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	err = tZoneUtil2.setLocalTime(
		&tzDtoOut,
		tIn,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	return tzDtoOut, nil
}

// copyIn - Copies input parameter TimeZoneDto data fields
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
func (tZoneUtil *timeZoneDtoUtility) copyIn(
	tzDto *TimeZoneDto,
	tzDto2 *TimeZoneDto,
	ePrefix string) {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.copyIn() "

	if tzDto == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	if tzDto2 == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tzDto2' is a 'nil' pointer!\n")
	}

	tZoneUtil2 := timeZoneDtoUtility{}

	tZoneUtil2.empty(tzDto, ePrefix)

	tzDto.Description = tzDto2.Description
	tzDto.TimeIn = tzDto2.TimeIn.CopyOut()
	tzDto.TimeOut = tzDto2.TimeOut.CopyOut()
	tzDto.TimeUTC = tzDto2.TimeUTC.CopyOut()
	tzDto.TimeLocal = tzDto2.TimeLocal.CopyOut()
	tzDto.DateTimeFmt = tzDto2.DateTimeFmt

	return
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
func (tZoneUtil *timeZoneDtoUtility) copyOut(
	tzDto *TimeZoneDto,
	ePrefix string) TimeZoneDto {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.copyOut() "

	if tzDto == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	tzu2 := TimeZoneDto{}

	tzu2.Description = tzDto.Description
	tzu2.TimeIn = tzDto.TimeIn.CopyOut()
	tzu2.TimeOut = tzDto.TimeOut.CopyOut()
	tzu2.TimeUTC = tzDto.TimeUTC.CopyOut()
	tzu2.TimeLocal = tzDto.TimeLocal.CopyOut()
	tzu2.DateTimeFmt = tzDto.DateTimeFmt

	return tzu2
}

// empty - Clears or returns the 'TimeZoneDto'
// object to an uninitialized or 'Empty' state.
//
func (tZoneUtil *timeZoneDtoUtility) empty(
	tzDto *TimeZoneDto,
	ePrefix string) {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.empty() "

	if tzDto == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	tzDto.Description = ""
	tzDto.TimeIn = DateTzDto{}
	tzDto.TimeOut = DateTzDto{}
	tzDto.TimeUTC = DateTzDto{}
	tzDto.TimeLocal = DateTzDto{}
	tzDto.DateTimeFmt = ""

	return
}

	// isValidTimeZoneDto - Analyzes the specified 'TimeZoneDto'
// instance and returns an error if the instance is INVALID.
//
func (tZoneUtil *timeZoneDtoUtility) isValidTimeZoneDto(
	tzDto *TimeZoneDto,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.isValidTimeZoneDto() "

	if tzDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	errorArray := make([]error, 0)

	if err := tzDto.TimeIn.IsValid(); err != nil {
		errorArray =
			append(errorArray, fmt.Errorf(ePrefix+"\ntzDto.TimeIn is INVALID!\nError='%v'\n", err.Error()))
	}

	if err := tzDto.TimeOut.IsValid(); err != nil {
		errorArray =
			append(errorArray, fmt.Errorf(ePrefix+"\nError: TimeOut is INVALID!\nError='%v'\n", err.Error()))
	}

	if err := tzDto.TimeUTC.IsValid(); err != nil {
		errorArray =
			append(errorArray, fmt.Errorf(ePrefix+"\nError: TimeUTC is INVALID!\nError='%v'\n", err.Error()))
	}

	if err := tzDto.TimeLocal.IsValid(); err != nil {
		errorArray =
			append(errorArray, fmt.Errorf(ePrefix+"Error:\nTimeLocal is INVALID!\nError='%v'\n", err.Error()))
	}

	dtUtil := DTimeUtility{}

	return dtUtil.ConsolidateErrors(errorArray)
}

func (tZoneUtil *timeZoneDtoUtility) isValidTimeZoneName(
	timeZoneName string,
	ePrefix string) (
				isValidTimeZone bool,
				timeZoneType TimeZoneType,
				err error) {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.isValidTimeZoneName() "

	isValidTimeZone = false
	timeZoneType = TzType.None()
	err = nil

	if len(timeZoneName) == 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "Input parameter 'timeZoneName' is an empty string!",
			err:                 nil,
		}

		return isValidTimeZone, timeZoneType, err
	}

	dateTime := time.Now().In(time.UTC)

	dtMech := DTimeMechanics{}
	var tzSpec TimeZoneSpecification

	tzSpec,
		err = dtMech.GetTimeZoneFromName(
		dateTime,
		timeZoneName,
		TzConvertType.Absolute(),
		ePrefix)

	if err != nil {
		err = nil
		return isValidTimeZone, timeZoneType, err
	}

	isValidTimeZone = true

	timeZoneType = tzSpec.GetTimeZoneType()

	return isValidTimeZone, timeZoneType, err
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
// (2) error - If errors are encountered, this method returns an error instance populated with
//        a valid 'error' message. If the method completes successfully the returned error
//       error type is set to 'nil'.
//
func (tZoneUtil *timeZoneDtoUtility) newAddDateTime(
	tzDtoIn *TimeZoneDto,
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFmtStr,
	ePrefix string) (TimeZoneDto, error) {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.newAddDateTime() "

	if tzDtoIn == nil {
		return TimeZoneDto{},
		&InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDtoIn",
			inputParameterValue: "",
			errMsg:              "Error: Input parameter 'tzDtoIn' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	tZoneDtoUtil2 := timeZoneDtoUtility{}

	err := tZoneDtoUtil2.isValidTimeZoneDto(tzDtoIn, ePrefix)

	if err != nil {
		return TimeZoneDto{},
			fmt.Errorf(ePrefix+
				"\nError: Input Parameter 'tzDtoIn' TimeZoneDto is INVALID!\n" +
				"Error='%v'\n",
				err.Error())
	}

	tzuOut := tZoneDtoUtil2.copyOut(tzDtoIn, ePrefix)

	tZoneDtoUtil2.setDateTimeFormat(
		&tzuOut,
		dateTimeFmtStr,
		ePrefix)

	err = tZoneDtoUtil2.addDateTime(
		&tzuOut,
		years,
		months,
		days,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds,
		ePrefix)

	return tzuOut, err
}

// newTzDto - Converts 'tIn' Date Time from existing time zone to a 'targetTz'
// or target Time Zone. The results are stored and returned in a TimeZoneDto
// data structure. TimeZoneDto.TimeIn stores the original 'tIn' Date Time,
// TimeZoneDto.TimeOut stores the converted date time using the target Time
// Zone.
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
//             (1) A New TimeZoneDto instance
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
func (tZoneUtil *timeZoneDtoUtility) newTzDto(
	tIn time.Time,
	targetTz,
	dateTimeFmtStr,
	ePrefix string) (TimeZoneDto, error) {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.newTzDto() "

	if tIn.IsZero() {
		return TimeZoneDto{},
						errors.New(ePrefix +
							"\nInput parameter 'tIn' has a value of 'zero'!\n")
	}

	tzDto2 := TimeZoneDto{}

	tZoneUtil2 := timeZoneDtoUtility{}

	dateTimeFmtStr = tZoneUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	tzDto2.DateTimeFmt = dateTimeFmtStr

	err := tZoneUtil2.setTimeIn(
		&tzDto2,
		tIn,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	err = tZoneUtil2.setTimeOutTz(
		&tzDto2,
		tIn,
		TzConvertType.Relative(),
		targetTz,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}


	err = tZoneUtil2.setUTCTime(
		&tzDto2,
		tIn,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	err = tZoneUtil2.setLocalTime(
		&tzDto2,
		tIn,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	return tzDto2, nil
}

// preProcessDateFormatStr - Provides a standardized method
// for implementing a default date time format string.
//
func (tZoneUtil *timeZoneDtoUtility) preProcessDateFormatStr(
	dateTimeFmtStr string) string {

	tZoneUtil.lock.Lock()
	defer tZoneUtil.lock.Unlock()

	dateTimeFmtStr = strings.TrimLeft(strings.TrimRight(dateTimeFmtStr, " "), " ")

	if len(dateTimeFmtStr) == 0 {
		lockDefaultDateTimeFormat.Lock()
		dateTimeFmtStr = DEFAULTDATETIMEFORMAT
		lockDefaultDateTimeFormat.Unlock()
	}

	return dateTimeFmtStr
}

// setDateTimeFormat - Sets the value of the TimeZoneDto.DateTimeFmt field.
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
//                          DEFAULTDATETIMEFORMAT =
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tZoneUtil *timeZoneDtoUtility) setDateTimeFormat(
	tzDto *TimeZoneDto,
	dateTimeFormatStr,
	ePrefix string) {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.setDateTimeFormat() "

	if tzDto == nil {
		panic("Input Parameter is a 'nil' pointer!")
	}

	if len(dateTimeFormatStr) == 0 {
		lockDefaultDateTimeFormat.Lock()
		dateTimeFormatStr = DEFAULTDATETIMEFORMAT
		lockDefaultDateTimeFormat.Unlock()
	}

	tzDto.DateTimeFmt = dateTimeFormatStr

	return
}

// setTimeIn - Assigns time and zone values to TimeZoneDto
// field 'TimeIn'.
//
func (tZoneUtil *timeZoneDtoUtility) setTimeIn(
	tzDto *TimeZoneDto,
	tIn time.Time,
	dateTimeFormatStr,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.setTimeIn() "

	if tzDto == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDto",
			inputParameterValue: "",
			errMsg:              "Input Parameter is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tIn.IsZero() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tIn",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tIn' has a value of 'zero'!",
			err:                 nil,
		}
	}

	tZoneUtil2 := timeZoneDtoUtility{}

	dateTimeFormatStr = tZoneUtil2.preProcessDateFormatStr(dateTimeFormatStr)

	tzDtoIn := DateTzDto{}

	dTzUtil := dateTzDtoUtility{}

	// Set tzDtoOut.TimeIn
	err := dTzUtil.setFromDateTime(
		&tzDtoIn,
		tIn,
		dateTimeFormatStr,
		ePrefix)

	if err != nil {
		return fmt.Errorf("tzDto.TimeIn configuration FAILED!\n" +
				"%v", err.Error())
	}

	tzDto.TimeIn = tzDtoIn.CopyOut()

	return nil
}

// setTimeInTzDef - Assigns time and zone values to TimeZoneDto
// field 'TimeIn'. Input parameter 'tIn' is first converted
// to the target time zone designated by 'targetTzSpec'.
//
func (tZoneUtil *timeZoneDtoUtility) setTimeInTzDef(
	tzDto *TimeZoneDto,
	tIn time.Time,
	timeConversionType TimeZoneConversionType,
	targetTzDef TimeZoneDefinition,
	dateTimeFormatStr,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.setTimeInTzDef() "

	if tzDto == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDto",
			inputParameterValue: "",
			errMsg:              "Input Parameter is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tIn.IsZero() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tIn",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tIn' has a value of 'zero'!",
			err:                 nil,
		}
	}

	err := targetTzDef.IsValid()

	if err != nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "targetTzDef",
			inputParameterValue: "",
			errMsg:              fmt.Sprintf(
				"Input Parameter 'targetTzDef' is INVALID!\n" +
					"Validation Error='%v'\n", err.Error()),
			err:                 nil,
		}
	}

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTzDef(
		&tzDto.TimeIn,
		tIn,
		timeConversionType,
		targetTzDef,
		dateTimeFormatStr,
		ePrefix)

	return err
}

// setTimeOutTz - Assigns date, time and time zone
// values to field 'TimeZoneDto.TimeOut' which is
// of type, 'DateTzDto'. The time zone conversion
// relies on the parameter 'tOutTimeZoneName' which
// must hold a valid time zone name.
//
// The parameter, 'tZoneConversionType', is an instance
// the type enumeration type TimeZoneConversionType.
// This parameter will determine how 'tOut' will be
// converted to the target time zone.
//
func (tZoneUtil *timeZoneDtoUtility) setTimeOutTz(
	tzDto *TimeZoneDto,
	tOut time.Time,
	tZoneConversionType TimeZoneConversionType,
	tOutTimeZoneName string,
	dateTimeFormat string,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.setTimeOutTz() "

	if tzDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	if tOut.IsZero() {
		return errors.New(ePrefix +
			"\nInput parameter 'tOut' has a value of 'zero'!\n")
	}

	if tZoneConversionType == TzConvertType.None() {
		return errors.New(ePrefix +
			"\nInput parameter 'tZoneConversionType' is INVALID.\n" +
			"tZoneConversionType= TzConvertType.None()\n")
	}

	if len(tOutTimeZoneName) == 0 {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tOutTimeZoneName",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tOutTimeZoneName' is an empty string!",
			err:                 nil,
		}
	}

	dTzUtil := dateTzDtoUtility{}

	dateTimeFormat = dTzUtil.preProcessDateFormatStr(dateTimeFormat)

	tzDtoOut := DateTzDto{}

	err := dTzUtil.setFromTimeTzName(
		&tzDtoOut, tOut,
		tZoneConversionType,
		tOutTimeZoneName,
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return fmt.Errorf( "tzDto.TimeOut configuration FAILED!\n" +
			"%v", err.Error())
	}

	tzDto.TimeOut = tzDtoOut.CopyOut()

	return nil
}


// setTimeOutTzSpec - Assigns date, time and time zone
// values to field 'tzDto.TimeOut' which is
// of type, 'DateTzDto'. The time zone conversion
// relies on the parameter 'tOutTimeZoneDef' which
// is of type, 'TimeZoneDefinition'.
//
// The parameter, 'tZoneConversionType', is an instance
// the type enumeration type TimeZoneConversionType.
// This parameter will determine how 'tOut' will be
// converted to the target time zone.
//
func (tZoneUtil *timeZoneDtoUtility) setTimeOutFromTimeZoneDef(
	tzDto *TimeZoneDto,
	tOut time.Time,
	timeConversionType TimeZoneConversionType,
	tOutTimeZoneDef TimeZoneDefinition,
	dateTimeFormat string,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.setTimeOutFromTimeZoneDef() "

	if tzDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	if tOut.IsZero() {
		return errors.New(ePrefix +
			"\nInput parameter 'tOut' has a value of 'zero'!\n")
	}

	if timeConversionType < TzConvertType.Absolute() ||
		timeConversionType > TzConvertType.Relative() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeConversionType",
			inputParameterValue: timeConversionType.String(),
			errMsg:              "Input parameter " +
				"'timeConversionType' MUST be either 'Absolute' or 'Relative'!",
			err:                 nil,
		}
	}

	err := tOutTimeZoneDef.IsValid()

	if err != nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tOutTimeZoneDef",
			inputParameterValue: "",
			errMsg:              fmt.Sprintf("Input parameter " +
				"'tOutTimeZoneDef' validation error!\n" +
				"'%v'",
			err.Error()),
			err:                 nil,
		}
	}

	dTzUtil := dateTzDtoUtility{}

	dateTimeFormat = dTzUtil.preProcessDateFormatStr(dateTimeFormat)

	dTz := DateTzDto{}

	err = dTzUtil.setFromTzDef(
		&dTz,
		tOut,
		timeConversionType,
		tOutTimeZoneDef,
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	tzDto.TimeOut = dTz.CopyOut()

	return nil
}


// setTimeOutTzSpec - Assigns date, time and time zone
// values to field 'TimeZoneDto.TimeOut' which is
// of type, 'DateTzDto'. The time zone conversion
// relies on the parameter 'tOutTimeZoneSpec' which
// is of type, 'TimeZoneSpecification'.
//
// The parameter, 'tZoneConversionType', is an instance
// the type enumeration type TimeZoneConversionType.
// This parameter will determine how 'tOut' will be
// converted to the target time zone.
//
func (tZoneUtil *timeZoneDtoUtility) setTimeOutTzSpec(
	tzDto *TimeZoneDto,
	tOut time.Time,
	tZoneConversionType TimeZoneConversionType,
	tOutTimeZoneSpec TimeZoneSpecification,
	dateTimeFormat string,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.setTimeOutTzSpec() "

	if tzDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	if tOut.IsZero() {
		return errors.New(ePrefix +
			"\nInput parameter 'tOut' has a value of 'zero'!\n")
	}

	if tZoneConversionType == TzConvertType.None() {
		return errors.New(ePrefix +
			"\nInput parameter 'tZoneConversionType' is INVALID.\n" +
			"tZoneConversionType= TzConvertType.None()\n")
	}

	err := tOutTimeZoneSpec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nInput Parameter 'tOutTimeZoneSpec' is Invalid!\n" +
			"Error='%v'\n", err.Error())
	}

	dTzUtil := dateTzDtoUtility{}

	dateTimeFormat = dTzUtil.preProcessDateFormatStr(dateTimeFormat)

	tzDtoOut := DateTzDto{}

	err = dTzUtil.setFromTzSpec(
		&tzDtoOut,
		tOut,
		tOutTimeZoneSpec,
		tZoneConversionType,
		dateTimeFormat,
		ePrefix)

	tzDto.TimeOut = tzDtoOut.CopyOut()

	return nil
}

// setUTCTime - Assigns UTC Time and zone values to TimeZoneDto fields
// 'TimeUTC' and 'TimeUTCZone'.
//
func (tZoneUtil *timeZoneDtoUtility) setUTCTime(
	tzDto *TimeZoneDto,
	dateTime time.Time,
	tZoneConversionType TimeZoneConversionType,
	dateTimeFormat string,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.setUTCTime() "

	if tzDto == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDto",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tzDto' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if dateTime.IsZero() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Input parameter 'dateTime' has a value of 'zero'!",
			err:                 nil,
		}
	}

	if tZoneConversionType < TzConvertType.Absolute() ||
		tZoneConversionType > TzConvertType.Relative() {

		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "",
			inputParameterValue: "tZoneConversionType",
			errMsg:              "Input Parameter 'tZoneConversionType' " +
				"MUST Be 'Absolute' or 'Relative'!",
			err:                 nil,
		}
	}

	// tzDto.TimeUTC, err = DateTzDto{}.New(dateTime.UTC(), tzDto.DateTimeFmt)
	dTzUtil := dateTzDtoUtility{}

	dateTimeFormat = dTzUtil.preProcessDateFormatStr(dateTimeFormat)

	tzDtoUtc := DateTzDto{}

	err := dTzUtil.setFromTimeTzName(
						&tzDtoUtc,
						dateTime,
						tZoneConversionType,
						TZones.UTC(),
						dateTimeFormat,
						ePrefix)

	if err != nil {
		return fmt.Errorf("tzDto.TimeUTC configuration FAILED!\n" +
			"%v", err.Error())
	}

	tzDto.TimeUTC = tzDtoUtc.CopyOut()

	return nil
}

// setLocalTime - Assigns Local Time to TimeZoneDto field
// 'TimeLocal'.
//
func (tZoneUtil *timeZoneDtoUtility) setLocalTime(
	tzDto *TimeZoneDto,
	dateTime time.Time,
	tZoneConversionType TimeZoneConversionType,
	dateTimeFormat string,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.SetLocalTime() "


	if tzDto == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzDto",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tzDto' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if dateTime.IsZero() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Input parameter 'dateTime' has a value of 'zero'!",
			err:                 nil,
		}
	}

	if tZoneConversionType < TzConvertType.Absolute() ||
		tZoneConversionType > TzConvertType.Relative() {

		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "",
			inputParameterValue: "tZoneConversionType",
			errMsg:              "Input Parameter 'tZoneConversionType' " +
				"MUST Be 'Absolute' or 'Relative'!",
			err:                 nil,
		}
	}

	dTzUtil := dateTzDtoUtility{}

	dateTimeFormat = dTzUtil.preProcessDateFormatStr(dateTimeFormat)

	tzDtoLocal := DateTzDto{}

	err := dTzUtil.setFromTimeTzName(
		&tzDtoLocal,
		dateTime,
		tZoneConversionType,
		TZones.Local(),
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return err
	}

	tzDto.TimeLocal = tzDtoLocal.CopyOut()

	return nil
}

