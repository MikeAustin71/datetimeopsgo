package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type DTimeUtility struct {
	lock      sync.Mutex
}

// AbsoluteTimeToTimeZoneDtoConversion - Converts a given time to
// another time zone using the 'Absolute' conversion method.
// This means that the years, months, days, hours, minutes,
// seconds and nanoseconds of the original 'dateTime' are not
// changed. That time value is simply assigned to another
// designated time zone. The target time zone is derived from
// input parameter 'timeZoneDefDto', an instance of type
// 'TimeZoneDefDto'.
//
// For example, assume that 'dateTime' represents 10:00AM in USA
// time zone 'Central Standard Time'.  Using the 'Absolute'
// conversion method, and converting this time value to the USA
// Eastern Standard Time Zone would result in a date time of
// 10:00AM EST or Eastern Standard Time. The time value of
// 10:00AM is not changed, it is simply assigned to another
// time zone.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// dateTime            time.Time  - The date time to be converted.
//
// timeZoneDefDto TimeZoneDefDto  - A properly initialized 'TimeZoneDto'
//                                  encapsulating the time zone to which
//                                  'dateTime' will be converted.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// time.Time  - The date time converted to the time zone specified in
//              in input parameter 'timeZoneDefDto'.
//
// error      - If the method completes successfully this value is set
//              to 'nil'. If an error is encountered, the returned error
//              value encapsulates an appropriate error message.
//
func (dtUtil *DTimeUtility) AbsoluteTimeToTimeZoneDtoConversion(
	dateTime time.Time,
	timeZoneDefDto TimeZoneDefDto) (time.Time, error) {

		dtUtil.lock.Lock()

		defer dtUtil.lock.Unlock()

	ePrefix := "DTimeUtility.AbsoluteTimeToTimeZoneDtoConversion() "

	if dateTime.IsZero() {
		return time.Time{},
			errors.New(ePrefix +
				"\nError: Input parameter 'dateTime' is zero!")
	}

	if !timeZoneDefDto.IsValid() {
		return time.Time{},
		 errors.New(ePrefix +
		 	"Input parameter 'timeZoneDefDto' is Invalid!\n")
	}

	zoneOffset := timeZoneDefDto.GetZoneOffset()

	// FmtDateTimeTzNanoYMD = "2006-01-02 15:04:05.000000000 -0700 MST"
	lenTimeElement := len("2006-01-02 15:04:05.000000000 ")
	dtStr := dateTime.Format(FmtDateTimeTzNanoYMD)

	dtStr = dtStr[:lenTimeElement] + zoneOffset

	resultDateTime, err := time.Parse(FmtDateTimeTzNanoYMD, dtStr)

	if err != nil {
		return time.Time{},
			fmt.Errorf(ePrefix +
				"\nError returned by time.Parse(FmtDateTimeTzNanoYMD, dtStr)\n" +
				"FmtDateTimeTzNanoYMD='%v'\ndtStr='%v'\nError='%v'\n",
				FmtDateTimeTzNanoYMD, dtStr, err.Error())
	}

	return resultDateTime, nil
}

// AbsoluteTimeToTimeZoneDtoConversion - Converts a given time to
// another time zone using the 'Absolute' conversion method.
// This means that the years, months, days, hours, minutes,
// seconds and nanoseconds of the original 'dateTime' are not
// changed. That time value is simply assigned to another
// designated time zone. The target time zone is derived from
// input parameter 'timeZoneDefDto', an instance of type
// 'TimeZoneDefDto'.
//
// For example, assume that 'dateTime' represents 10:00AM in USA
// time zone 'Central Standard Time'.  Using the 'Absolute'
// conversion method, and converting this time value to the USA
// Eastern Standard Time Zone would result in a date time of
// 10:00AM EST or Eastern Standard Time. The time value of
// 10:00AM is not changed, it is simply assigned to another
// time zone.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// dateTime            time.Time  - The date time to be converted.
//
// timeZoneName           string  - A string containing a valid IANA,
//                                  Military or "Local" time zone.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// time.Time  - The date time converted to the time zone specified in
//              in input parameter 'timeZoneDefDto'.
//
// error      - If the method completes successfully this value is set
//              to 'nil'. If an error is encountered, the returned error
//              value encapsulates an appropriate error message.
//
func (dtUtil *DTimeUtility) AbsoluteTimeToTimeZoneNameConversion(
	dateTime time.Time,
	timeZoneName string) (time.Time, error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix := "DTimeUtility.AbsoluteTimeToTimeZoneNameConversion() "

	if dateTime.IsZero() {
		return time.Time{},
			errors.New(ePrefix +
				"\nError: Input parameter 'dateTime' is zero!")
	}

	timeZoneName = strings.TrimLeft(strings.TrimRight(timeZoneName, " "), " ")

	if len(timeZoneName) == 0  {
		return time.Time{},
		 errors.New(ePrefix +
		 	"Input parameter 'timeZoneName' is an empty string!\n")
	}

	timeZoneDefDto := TimeZoneDefDto{}

	tzDefUtil := timeZoneDefUtility{}

	err := tzDefUtil.setFromTimeZoneName(&timeZoneDefDto, timeZoneName, ePrefix)

	if err != nil {
		return time.Time{}, err
	}

	zoneOffset := timeZoneDefDto.GetZoneOffset()

	// FmtDateTimeTzNanoYMD = "2006-01-02 15:04:05.000000000 -0700 MST"
	lenTimeElement := len("2006-01-02 15:04:05.000000000 ")
	dtStr := dateTime.Format(FmtDateTimeTzNanoYMD)

	dtStr = dtStr[:lenTimeElement] + zoneOffset

	resultDateTime, err := time.Parse(FmtDateTimeTzNanoYMD, dtStr)

	if err != nil {
		return time.Time{},
			fmt.Errorf(ePrefix +
				"\nError returned by time.Parse(FmtDateTimeTzNanoYMD, dtStr)\n" +
				"FmtDateTimeTzNanoYMD='%v'\ndtStr='%v'\nError='%v'\n",
				FmtDateTimeTzNanoYMD, dtStr, err.Error())
	}

	return resultDateTime, nil
}
