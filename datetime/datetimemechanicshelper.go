package datetime

import (
	"sync"
	"time"
)

type dateTimeMechanicsHelper struct {
	input  string
	output string
	lock   sync.Mutex
}

// gregorianDateToJulianDayNo - Converts a Gregorian calendar
// date to Julian Day Number (JDN).
//
// This algorithm is valid for all (possibly proleptic) Gregorian
// calendar dates after November 23, −4713. Divisions are integer
// divisions, fractional parts are ignored.[72]
//
// JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075
//
// Reference Wikipedia:
//   https://en.wikipedia.org/wiki/Julian_day
//
// Note that the input parameter 'gregorianDateTime' will first be
// converted to Universal Coordinated Time (UTC).
//
//  Example: The Julian Date for 00:30:00.0 UT January 1, 2013, is 2 456 293.520 833
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  gregorianDateTime  time.Time
//     - This date time value will be converted to Universal
//       Coordinated Time (UTC) before conversion to a Julian
//       Day Number (JDN).
//
//
//  ePrefix            string
//     - A string containing the names of the calling functions
//       which invoked this method. The last character in this
//       string should be a blank space.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  gregorianDateUtc   time.Time
//     - The input parameter 'gregorianDateTime' converted to Universal
//       Coordinated Time (UTC). This is the date time used to compute
//       the Julian Day Number (JDN)
//
//
//  julianDayNo        int64
//     - An int64 value representing the Julian Day Number (JDN) equivalent
//       to input parameter 'gregorianDateTime' converted to Universal
//       Coordinated Time (UTC).
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered this error Type will encapsulate
//       an error message.
//
//
func (dtMechHelper *dateTimeMechanicsHelper) gregorianDateToJulianDayNo(
	gregorianDateTime time.Time,
	ePrefix string) (
	gregorianDateUtc time.Time,
	julianDayNo int64,
	err error) {

	dtMechHelper.lock.Lock()

	defer dtMechHelper.lock.Unlock()

	ePrefix += "dtMechHelper.dateTimeMechanicsHelper() "

	gregorianDateUtc = gregorianDateTime.UTC()
	julianDayNo = 0
	err =  nil

	// This algorithm is valid for all (possibly proleptic) Gregorian
	// calendar dates after November 23, −4713.
	limitDateTime := time.Date(
		-4713,
		11,
		24,
		0,
		0,
		0,
		0,
		time.UTC)

	if gregorianDateUtc.Before(limitDateTime) {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "gregorianDateTime",
			inputParameterValue: "",
			errMsg:              "Error: 'gregorianDateTime' is invalid. Valid date times must\n" +
				"occur after November 23, −4713 UTC as expressed in Astronomical\n" +
				"year numbering.",
			err:                 nil,
		}

		return gregorianDateUtc, julianDayNo, err
	}

	Year := int64(gregorianDateUtc.Year())
	Month := int64(gregorianDateUtc.Month())
	Day := int64(gregorianDateUtc.Day())

	/*
		JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075
	*/

	julianDayNo =
		(int64(1461) * (Year + int64(4800) +
		(Month - int64(14))/int64(12)))/int64(4) +
		(int64(367) * (Month - int64(2) -
		int64(12) * ((Month - int64(14))/int64(12))))/int64(12) -
		(int64(3) * ((Year + int64(4900) +
		(Month - int64(14))/int64(12))/int64(100)))/int64(4) +
		Day - int64(32075)

	gregorianTimeNanoSecs := int64(gregorianDateUtc.Hour()) * HourNanoSeconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Minute()) * MinuteNanoSeconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Second()) * SecondNanoseconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Nanosecond())

	if gregorianTimeNanoSecs < NoonNanoSeconds {
		julianDayNo -= 1
	}

	return gregorianDateUtc, julianDayNo, err

}
