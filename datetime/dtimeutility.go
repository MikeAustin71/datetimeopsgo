package datetime

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type DTimeUtility struct {
	lock sync.Mutex
}

// ConsolidateErrors - Receives an array of errors and converts them
// to a single error which is returned to the caller. Multiple errors
// are separated by a new line character.
//
// If the length of the error array is zero, this method returns nil.
//
func (dtUtil *DTimeUtility) ConsolidateErrors(errs []error) error {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	lErrs := len(errs)

	if lErrs == 0 {
		return nil
	}

	errStr := ""

	for i := 0; i < lErrs; i++ {

		if errs[i] == nil {
			continue
		}

		tempStr := fmt.Sprintf("%v", errs[i].Error())

		tempStr = strings.TrimLeft(strings.TrimRight(tempStr, " "), " ")

		strLen := len(tempStr)

		for strings.HasSuffix(tempStr,"\n") &&
			strLen > 1 {

			tempStr = tempStr[0:strLen-1]
			strLen--
		}

		if i == (lErrs - 1) {
			errStr += fmt.Sprintf("%v", tempStr)
		} else if i == 0 {
			errStr = fmt.Sprintf("\n%v\n\n", tempStr)
		} else {
			errStr += fmt.Sprintf("%v\n\n", tempStr)
		}
	}

	return fmt.Errorf("%v", errStr)
}

// GetUtcOffsetTzAbbrvFromDateTime - Receives a time.Time, date
// time, input parameter and extracts and returns the
// 5-character UTC offset and UTC offset.
//
// UTC Offsets are returned in the format illustrated by the
// following examples:
//   +1030
//   -0500
//   +1100
//   -1100
//
// Example:
//  Time String:  2019-12-26 00:56:15 -0600 CST
//  Returned UTC Offset:  '-0600'
//  Returned Time Zone Abbreviation: 'CST'
//
func (dtUtil *DTimeUtility) GetUtcOffsetTzAbbrvFromDateTime(
	dateTime time.Time,
	ePrefix string) (utcOffset, tzAbbrv string, err error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.GetUtcOffsetTzAbbrvFromDateTime() "

	dtMech := DTimeMechanics{}

	return dtMech.GetUtcOffsetTzAbbrvFromDateTime(dateTime, ePrefix)
}

// PreProcessTimeZoneLocation - Scans a time zone location
// name string and attempts to correct errors. If input
// parameter 'timeZoneLocation' is an empty string, this
// method set 'timeZoneLocation' to 'UTC'.
//
func (dtUtil *DTimeUtility) PreProcessTimeZoneLocation(
	timeZoneLocation string) string {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	timeZoneLocation =
		strings.TrimLeft(strings.TrimRight(timeZoneLocation, " "), " ")

	if len(timeZoneLocation) == 0 {
		return timeZoneLocation
	}

	testZone := strings.ToLower(timeZoneLocation)

	if testZone == "utc" {

		timeZoneLocation = TZones.UTC()

	} else if testZone == "uct" {

		timeZoneLocation = TZones.UCT()

	} else if testZone == "gmt" {

		timeZoneLocation = TZones.UTC()

	} else if testZone == "local" {

		return TZones.Local()
	}

	return timeZoneLocation

}

// ParseMilitaryTzNameAndLetter - Parses a text string which
// contains either a single letter military time zone designation
// or a multi-character time zone text name.
//
// If successful, three populated strings are returned. The first
// is the valid Military Time Zone Letter designation. The second
// returned string contains the text name of the Military Time
// Zone. The third string contains the name of the equivalent
// IANA Time Zone. This is required because Golang does not
// currently support Military Time Zones.
//
// In addition to the three strings, a successful method completion
// will also return the equivalent IANA Time Zone Location pointer
// (*time.Location).
//
// If an error is encountered, the return value, 'err' is populated
// with an appropriate error message. Otherwise, 'err' is set
// equal to 'nil' signaling no error was encountered.
//
func (dtUtil *DTimeUtility) ParseMilitaryTzNameAndLetter(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	timeZoneName string,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.ParseMilitaryTzNameAndLetter() "

	tzSpec = TimeZoneSpecification{}
	err = nil

	if dateTime.IsZero() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Input parameter 'dateTime' " +
				"has a ZERO value!",
			err:                 nil,
		}
		return tzSpec, err
	}

	if timeConversionType != TzConvertType.Relative() &&
		timeConversionType != TzConvertType.Absolute() {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeConversionType",
			inputParameterValue: timeConversionType.String(),
			errMsg:              "Input parameter 'timeConversionType' value is Invalid!",
			err:                 nil,
		}

		return tzSpec, err
	}

	timeZoneName =
		strings.TrimLeft(strings.TrimLeft(timeZoneName, " "), " ")

	lMilTz := len(timeZoneName)

	if lMilTz == 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "Error: Input Parameter 'timeZoneName' is empty string!",
			err:                 nil,
		}

		return tzSpec, err
	}

	tzMech := TimeZoneMechanics{}

	return tzMech.ParseMilitaryTzNameAndLetter(
		dateTime,
		timeConversionType,
		timeZoneName,
		ePrefix)
}
