package datetime

import (
	"errors"
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

	utcOffset = tStr[lenLeadOffsetStr : lenLeadOffsetStr+ 5]

	return utcOffset, tzAbbrv, err
}
