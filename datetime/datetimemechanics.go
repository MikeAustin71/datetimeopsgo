package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type dateTimeMechanics struct {
	lock sync.Mutex
}

// allocateSecondsToHrsMinSecs - Useful in calculating offset hours,
// minutes and seconds from UTC+0000. A total signed seconds value
// is passed as an input parameter. This method then breaks down
// hours, minutes and seconds as positive integer values. The sign
// of the hours, minutes and seconds is returned in the 'sign'
// parameter as +1 or -1.
//
func (dtMech *dateTimeMechanics) allocateSecondsToHrsMinSecs(
	signedTotalSeconds int) (hours, minutes, seconds, sign int) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	hours = 0
	minutes = 0
	seconds = 0
	sign = 1

	if signedTotalSeconds == 0 {
		return hours, minutes, seconds, sign
	}

	if signedTotalSeconds < 0 {
		sign = -1
	}

	remainingSeconds := signedTotalSeconds

	remainingSeconds *= sign

	hours = remainingSeconds / 3600

	remainingSeconds -= hours * 3600

	if remainingSeconds > 0 {
		minutes = remainingSeconds / 60
		remainingSeconds -= minutes * 60
	}

	seconds = remainingSeconds

	return hours, minutes, seconds, sign
}

// convertTzAbbreviationToTimeZone - receives an input parameter,
// 'tzAbbrvLookupKey' which is used to look up a time zone abbreviation
// and return an associated IANA Time Zone Name.
//
// The method uses the global variable, 'tzAbbrvToTimeZonePriorityList'
// to assign the IANA Time Zone in cases of multiple time zones
// associated with the Time Zone Abbreviation.
//
// The 'tzAbbrvLookupKey' is formatted the Time Zone Abbreviation
// followed by the UTC offsets as illustrated by the following
// examples:
//   "EDT-0400"
//   "EST-0500"
//   "CDT-0500"
//   "CST-0600"
//   "PDT-0700"
//   "PST-0800"
//
// The associated IANA Time Zone name is identified using the
// global variable 'mapTzAbbrvsToTimeZones' which is accessed
// through method StdTZoneAbbreviations{}.AbbrvOffsetToTimeZones().
//
// If an associated IANA Time Zone is not found the returned
// boolean value, 'isValidTzAbbreviation', is set to 'false'.
//
func (dtMech *dateTimeMechanics) convertTzAbbreviationToTimeZone(
	tzAbbrvLookupKey string,
	ePrefix string) (
	milTzLetter,
	milTzName,
	ianaTimeZoneName string,
	ianaLocationPtr *time.Location,
	err error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	milTzLetter = ""
	milTzName = ""
	ianaTimeZoneName = ""
	ianaLocationPtr = nil
	err = nil

	ePrefix += "dateTimeMechanics.convertTzAbbreviationToTimeZone() "

	if len(tzAbbrvLookupKey) == 0 {
		err = &InputParameterError{
			ePrefix:ePrefix,
			inputParameterName:tzAbbrvLookupKey,
			errMsg:"tzAbbrvLookKey is a zero length string!",
			err: nil}

		return milTzLetter,
			milTzName,
			ianaTimeZoneName,
			ianaLocationPtr,
			err

	}

	stdAbbrvs := StdTZoneAbbreviations{}

	tzones, ok := stdAbbrvs.AbbrvOffsetToTimeZones(tzAbbrvLookupKey)

	if !ok {
		ePrefix += "StdTZoneAbbreviations.AbbrvOffsetToTimeZones() "
		err = &TzAbbrvMapLookupError{
			ePrefix:  ePrefix,
			mapName:  "mapTzAbbrvsToTimeZones",
			lookUpId: tzAbbrvLookupKey,
			errMsg: "",
			err:      nil,
		}
		return milTzLetter,
			milTzName,
			ianaTimeZoneName,
			ianaLocationPtr,
			err
	}

	lenTZones := len(tzones)

	if lenTZones == 0 {
		err = &TzAbbrvMapLookupError{
			ePrefix:  ePrefix,
			mapName:  "mapTzAbbrvsToTimeZones",
			lookUpId: tzAbbrvLookupKey,
			errMsg: "Map returned a zero length time zones string array!",
			err:      nil,
		}
		return milTzLetter,
			milTzName,
			ianaTimeZoneName,
			ianaLocationPtr,
			err
	}

	var tzAbbrRef TimeZoneAbbreviationDto

	tzAbbrRef, ok = stdAbbrvs.AbbrvOffsetToTzReference(tzAbbrvLookupKey)

	if !ok {
		ePrefix += "StdTZoneAbbreviations.AbbrvOffsetToTzReference() "
		err = &TzAbbrvMapLookupError{
			ePrefix:  ePrefix,
			mapName:  "mapTzAbbreviationReference",
			lookUpId: tzAbbrvLookupKey,
			errMsg: "",
			err:      nil,
		}
		return milTzLetter,
			milTzName,
			ianaTimeZoneName,
			ianaLocationPtr,
			err
	}

	if tzAbbrRef.Location == "Military" {

		milTzLetter = tzAbbrRef.Abbrv
		milTzName = tzAbbrRef.AbbrvDescription

	}

	dtMech2 := dateTimeMechanics{}

//loadTzLocationPtr(
	if lenTZones == 1 {

		ianaLocationPtr, err = dtMech2.loadTzLocationPtr(tzones[0], ePrefix)

		if err != nil {
			milTzLetter = ""
			milTzName = ""
			ianaLocationPtr = nil

			return milTzLetter,
				milTzName,
				ianaTimeZoneName,
				ianaLocationPtr,
				err
		}

		ianaTimeZoneName = tzones[0]

		return milTzLetter,
			milTzName,
			ianaTimeZoneName,
			ianaLocationPtr,
			err
	}

	lockTzAbbrvToTimeZonePriorityList.Lock()
	defer lockTzAbbrvToTimeZonePriorityList.Unlock()

	for i := 0; i < lenTZones && ianaTimeZoneName == ""; i++ {

		for j := 0; j < lenTzAbbrvToTimeZonePriorityList; j++ {
			if strings.HasPrefix(tzones[i], tzAbbrvToTimeZonePriorityList[j]) {

				ianaTimeZoneName = tzones[i]
				break
			}
		}
	}

	if len(ianaTimeZoneName) == 0 {
		ianaTimeZoneName = tzones[0]
	}

	ianaLocationPtr, err = dtMech2.loadTzLocationPtr(ianaTimeZoneName, ePrefix)

	if err != nil {
		milTzLetter = ""
		milTzName = ""
		ianaTimeZoneName = ""
		ianaLocationPtr = nil
	}

	return milTzLetter,
		milTzName,
		ianaTimeZoneName,
		ianaLocationPtr,
		err
}

// getUtcOffsetTzAbbrvFromDateTime - Receives a time.Time, date
// time, input parameter and extracts and returns the
// 5-character UTC offset and the time zone abbreviation.
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
func (dtMech *dateTimeMechanics) getUtcOffsetTzAbbrvFromDateTime(
	dateTime time.Time,
	ePrefix string) (utcOffset, tzAbbrv string, err error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	utcOffset = ""
	tzAbbrv = ""
	err = nil

	ePrefix += "dateTimeMechanics.getUtcOffsetTzAbbrvFromDateTime() "

	if dateTime.IsZero() {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is ZERO!\n")

		return utcOffset, tzAbbrv, err
	}

	tStr := dateTime.Format("2006-01-02 15:04:05 -0700 MST")

	lenLeadTzAbbrvStr := len("2006-01-02 15:04:05 -0700 ")

	tzAbbrv = tStr[lenLeadTzAbbrvStr:]

	tzAbbrv = strings.TrimRight(tzAbbrv, " ")

	lenLeadOffsetStr := len("2006-01-02 15:04:05 ")

	utcOffset = tStr[lenLeadOffsetStr : lenLeadOffsetStr+5]

	return utcOffset, tzAbbrv, err
}

// loadTzLocationPtr - Provides a single method for calling
// time.LoadLocation(). This method may be altered in the future
// to load time zones from an internal file thus affording
// consistency in time zone definitions without relying on
// zoneinfo.zip databases residing on host computers.
//
// If successful, this method returns a *time.Location or
// location pointer to a valid time zone.
//
func (dtMech *dateTimeMechanics)loadTzLocationPtr(
	timeZoneName string,
	ePrefix string) (*time.Location, error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "dateTimeMechanics.loadTzLocationPtr() "

	if len(timeZoneName) == 0 {
		return nil,
			&TimeZoneError{
				ePrefix: ePrefix,
				errMsg:  "Input parameter 'timeZoneName' is a empty!",
				err:     nil,
			}
	}

	locPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		return nil,
			&TimeZoneError{
				ePrefix: ePrefix,
				errMsg:  fmt.Sprintf("Error returned by time.LoadLocation(timeZoneName)!\n" +
					"timeZoneName='%v'\nError='%v'\n", timeZoneName, err.Error()),
				err:     nil,
			}
	}

	return locPtr, nil
}