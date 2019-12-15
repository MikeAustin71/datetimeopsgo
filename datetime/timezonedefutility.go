package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type timeZoneDefUtility struct {
	lock sync.Mutex
}


// allocateZoneOffsetSeconds - allocates a signed value of total offset seconds from
// UTC to the associated fields in the current TimeZoneDefDto instance.
func (tzDefUtil timeZoneDefUtility) allocateZoneOffsetSeconds(
  tzdef *TimeZoneDefDto,
	signedZoneOffsetSeconds int) {

	if signedZoneOffsetSeconds < 0 {
		tzdef.zoneSign = -1
	} else {
		tzdef.zoneSign = 1
	}

	tzdef.zoneOffsetSeconds = signedZoneOffsetSeconds

	signedZoneOffsetSeconds *= tzdef.zoneSign

	tzdef.offsetHours = 0
	tzdef.offsetMinutes = 0
	tzdef.offsetSeconds = 0

	if signedZoneOffsetSeconds == 0 {
		return
	}

	tzdef.offsetHours = signedZoneOffsetSeconds / 3600 // compute hours
	signedZoneOffsetSeconds -= tzdef.offsetHours * 3600

	tzdef.offsetMinutes = signedZoneOffsetSeconds / 60 // compute minutes
	signedZoneOffsetSeconds -= tzdef.offsetMinutes * 60

	tzdef.offsetSeconds = signedZoneOffsetSeconds

	return
}

// setZoneProfile - assembles and assigns the composite zone
// offset, zone names, zone abbreviation and UTC offsets.
//
// The TimeZoneDefDto.ZoneOffset field formatted in accordance
// with the following examples:
//      "-0600 CST"
//      "+0200 EET"
//
func (tzDefUtil timeZoneDefUtility) setZoneProfile(
	tzdef *TimeZoneDefDto) {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	tzdef.zoneOffset = ""

	// Generates an offset in the form of "+0330" or "-0330"
	if tzdef.zoneSign < 0 {
		tzdef.zoneOffset += "-"
	} else {
		tzdef.zoneOffset += "+"
	}

	tzdef.zoneOffset += fmt.Sprintf("%02d%02d", tzdef.offsetHours, tzdef.offsetMinutes)

	tzdef.utcOffset = tzdef.zoneOffset

	if tzdef.offsetSeconds > 0 {
		tzdef.zoneOffset += fmt.Sprintf("%02d", tzdef.offsetSeconds)
	}

	// Generates final ZoneOffset in the form
	// "-0500 CST" or "+0200 EET"
	tzdef.zoneOffset += " " + tzdef.zoneName

	return
}


// parseMilitaryTzNameAndLetter - Parses a text string which
// contains either a single letter military time zone designation
// or a multi-character military time zone text name.
//
// If successful, three string values are returned. The first
// is a string containing the valid Military Time Zone
// Letter designation. The second returned string contains
// the text name of the Military Time Zone. The third string,
//
// If an error is encountered, the return value, 'err' is populated
// with an appropriate error message. Otherwise, 'err' is set
// to 'nil' signaling no error was encountered.
//
func (tzDefUtil timeZoneDefUtility) parseMilitaryTzNameAndLetter(
	rawTz,
	ePrefix string) (milTzLetter, milTzName, equivalentIanaTimeZone string, err error) {

		tzDefUtil.lock.Lock()

		defer tzDefUtil.lock.Unlock()

	milTzLetter = ""
	milTzName = ""
	equivalentIanaTimeZone = ""
	err = nil

	ePrefix += "timeZoneDefUtility.parseMilitaryTzNameAndLetter() "

	lMilTz := len(rawTz)

	if lMilTz == 0 {
		err = errors.New(ePrefix +
			"\nError: Input Parameter 'rawTz' is EMPTY!\n")
		return milTzLetter, milTzName, equivalentIanaTimeZone, err
	}

	var ok bool
	milTzData := MilitaryTimeZoneData{}

	if lMilTz == 1 {

		milTzLetter = strings.ToUpper(rawTz)

		milTzName , ok =
			milTzData.MilTzLetterToTextName(milTzLetter)

		if !ok {
			err = fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'militaryTz' is INVALID!\n" +
				"'militaryTz' DOES NOT map to a valid Military Time Zone.\n" +
				"militaryTz='%v'", milTzLetter)
			return milTzLetter, milTzName, equivalentIanaTimeZone, err
		}

		equivalentIanaTimeZone, ok = milTzData.MilitaryTzToIanaTz(milTzName)

		if !ok {
			err = fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'rawTz' is INVALID!\n" +
				"'rawTz' DOES NOT map to a valid IANA Time Zone.\n" +
				"rawTz='%v'", milTzName)
			return milTzLetter, milTzName, equivalentIanaTimeZone, err
		}

	} else {
		// lMilTz > 1
		temp1 := rawTz[:1]
		temp2 := rawTz[1:]

		temp1 = strings.ToUpper(temp1)
		temp2 = strings.ToLower(temp2)

		milTzLetter = temp1
		milTzName = temp1 + temp2

		equivalentIanaTimeZone, ok = milTzData.MilitaryTzToIanaTz(milTzName)

		if !ok {
			err = fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'rawTz' is INVALID!\n" +
				"'rawTz' DOES NOT map to a valid IANA Time Zone.\n" +
				"Military Time Zone Letter='%v'\n" +
				"Military Time Zone Text Name='%v'", milTzLetter ,milTzName)
			milTzLetter = ""
			milTzName = ""
			equivalentIanaTimeZone = ""
			return milTzLetter, milTzName, equivalentIanaTimeZone, err
		}
	}

	err = nil

	return milTzLetter, milTzName, equivalentIanaTimeZone, err
}