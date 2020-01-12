package datetime

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type timeZoneDtoUtility struct {
	lock        sync.Mutex
}


// newTzDto - Converts 'tIn' Date Time from existing time zone to a 'targetTz'
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

	dtUtil := DTimeUtility{}

	_,
		_,
		_,
		_,
		_,
		err := dtUtil.GetTimeZoneFromName(targetTz, ePrefix)

	if err != nil {
		return TimeZoneDto{},
		fmt.Errorf("'targetTz' Time Zone is INVALID!\n" +
			"targetTz='%v'\n" +
			"%v", targetTz, err.Error())
	}

	tzDtoOut := TimeZoneDto{}

	dTzUtil := dateTzDtoUtility{}

	// Set tzDtoOut.TimeIn
	err = dTzUtil.setFromDateTime(
		&tzDtoOut.TimeIn,
		tIn,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{},
			fmt.Errorf("tzDtoOut.TimeIn configuration FAILED!\n" +
				"%v", err.Error())
	}

	// Set tzDtoOut.TimeOut
	err = dTzUtil.setFromTimeTz(
		&tzDtoOut.TimeOut,
		tIn,
		targetTz,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{},
			fmt.Errorf("tzDtoOut.TimeOut configuration FAILED!\n" +
				"%v", err.Error())
	}

	// Set tzDtoOut.TimeUTC
	err = dTzUtil.setFromTimeTz(
		&tzDtoOut.TimeUTC,
		tIn,
		TZones.UTC(),
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{},
			fmt.Errorf("tzDtoOut.TimeUTC configuration FAILED!\n" +
				"%v\n", err.Error())
	}
	// Set tzDtoOut.TimeLocal
	err = dTzUtil.setFromTimeTz(
		&tzDtoOut.TimeLocal,
		tIn,
		TZones.Local(),
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{},
			fmt.Errorf("tzDtoOut.TimeLocal configuration FAILED!\n" +
				"%v\n", err.Error())
	}

	return tzDtoOut, nil
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


// setTimeIn - Assigns time and zone values to TimeZoneDto
// field 'TimeIn'.
//
func (tZoneUtil *timeZoneDtoUtility) setTimeIn(
	tIn time.Time,
	tzDto *TimeZoneDto,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.setTimeIn() "

	if tzDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	if tIn.IsZero() {
		return errors.New(ePrefix +
			"\nInput parameter 'tIn' has a value of 'zero'!\n")
	}

	// tzDto.TimeIn, err = DateTzDto{}.New(tIn, tzDto.DateTimeFmt)

	tzDtoIn := DateTzDto{}

	dTzUtil := dateTzDtoUtility{}

	// Set tzDtoOut.TimeIn
	err := dTzUtil.setFromDateTime(
		&tzDtoIn,
		tIn,
		tzDto.DateTimeFmt,
		ePrefix)

	if err != nil {
		return fmt.Errorf("tzDto.TimeIn configuration FAILED!\n" +
				"%v", err.Error())
	}

	tzDto.TimeIn = tzDtoIn.CopyOut()

	return nil
}


// setTimeOut - Assigns time and zone values to field 'TimeOut'
func (tZoneUtil *timeZoneDtoUtility) setTimeOut(
	tzDto *TimeZoneDto,
	tOut time.Time,
	tOutTimeZoneName string,
	tZoneConversionType TimeZoneConversionType,
	dateTimeFormat string,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "timeZoneDtoUtility.setTimeOut() "

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

	// tzDto.TimeOut, err = DateTzDto{}.New(tOut, tzDto.DateTimeFmt)

	dTzUtil := dateTzDtoUtility{}

	tzDtoOut := DateTzDto{}

	err := dTzUtil.setFromTimeTz(
		&tzDtoOut, tOut,
		tOutTimeZoneName,
		tZoneConversionType,
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return fmt.Errorf( "tzDto.TimeOut configuration FAILED!\n" +
			"%v", err.Error())
	}

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

	ePrefix += "TimeZoneDto.setUTCTime() "

	if tzDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	if dateTime.IsZero() {
		return errors.New(ePrefix +
			"\nInput parameter 'dateTime' has a value of 'zero'!\n")
	}

	if tZoneConversionType == TzConvertType.None() {
		return errors.New(ePrefix +
			"\nInput parameter 'tZoneConversionType' is INVALID.\n" +
			"tZoneConversionType= TzConvertType.None()\n")
	}

	// tzDto.TimeUTC, err = DateTzDto{}.New(dateTime.UTC(), tzDto.DateTimeFmt)
	dTzUtil := dateTzDtoUtility{}

	tzDtoUtc := DateTzDto{}

	err := dTzUtil.setFromTimeTz(
						&tzDtoUtc,
						dateTime,
						TZones.UTC(),
						tZoneConversionType,
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
		return errors.New(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	if dateTime.IsZero() {
		return errors.New(ePrefix +
			"\nInput parameter 'dateTime' has a value of 'zero'!\n")
	}

	if tZoneConversionType == TzConvertType.None() {
		return errors.New(ePrefix +
			"\nInput parameter 'tZoneConversionType' is INVALID.\n" +
			"tZoneConversionType= TzConvertType.None()\n")
	}

	dTzUtil := dateTzDtoUtility{}

	tzDtoLocal := DateTzDto{}

	err := dTzUtil.setFromTimeTz(
		&tzDtoLocal,
		dateTime,
		TZones.Local(),
		tZoneConversionType,
		dateTimeFormat,
		ePrefix)

	if err != nil {
		return fmt.Errorf("tzDto.TimeLocal configuration FAILED!\n" +
			"%v", err.Error())
	}

	tzDto.TimeLocal = tzDtoLocal.CopyOut()

	return nil
}

