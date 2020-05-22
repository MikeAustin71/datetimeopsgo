package datetime

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type DTimeUtility struct {
	lock *sync.Mutex
}

// ConsolidateErrors - Receives an array of errors and converts them
// to a single error which is returned to the caller. Multiple errors
// are separated by a new line character.
//
// If the length of the error array is zero, this method returns nil.
//
func (dtUtil *DTimeUtility) ConsolidateErrors(errs []error) error {

	if dtUtil.lock == nil {
		dtUtil.lock = new(sync.Mutex)
	}

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

// Compares two date times to determine if the
// Years, Months, Days, Hours, Minutes, Seconds
// and Nanoseconds are equivalent. This method
// ignores time zones.
func (dtUtil *DTimeUtility) EqualDateTimeComponents(
	dateTime1 time.Time,
	dateTime2 time.Time) bool {

	if dtUtil.lock == nil {
		dtUtil.lock = new(sync.Mutex)
	}

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	if dateTime1.IsZero() &&
		dateTime2.IsZero() {
		return true
	}

	if dateTime1.Year() == dateTime2.Year() &&
			dateTime1.Year() == dateTime2.Year() &&
			dateTime1.Month() == dateTime2.Month() &&
			dateTime1.Day() == dateTime2.Day() &&
			dateTime1.Hour() == dateTime2.Hour() &&
			dateTime1.Minute() == dateTime2.Minute() &&
			dateTime1.Second() == dateTime2.Second() &&
			dateTime1.Nanosecond() == dateTime2.Nanosecond() {
		return true
	}

	return false
}

// ConvertAstronomicalToGregorianBce - Used to convert negative
// year Astronomical date time values to their correct Gregorian
// Year or BCE equivalent by subtracting one year.
//
// Under the Gregorian calendar system there is NO year zero. The
// year 1BCE (Before Common Era) is followed by the year 1CE
// (Common Era).
//
// Under the Astronomical year numbering system, the year zero does
// exist and is used in computations.
//
// Gregorian Year Zero:
// See Wikipedia https://en.wikipedia.org/wiki/Year_zero :
//
//  "The year zero does not exist in the Anno Domini (AD) system
//  commonly used to number years in the Gregorian calendar and
//  in its predecessor, the Julian calendar. In this system, the
//  year 1 BC is followed by AD 1. However, there is a year zero
//  in astronomical year numbering (where it coincides with the
//  Julian year 1 BC) and in ISO 8601:2004 (where it coincides
//  with the Gregorian year 1 BC), as well as in all Buddhist
//  and Hindu calendars."
//
// For information on the "Common Era" calendar notation see:
//   https://en.wikipedia.org/wiki/Common_Era
//
// For Astronomical year numbering see:
//   https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
func (dtUtil *DTimeUtility) ConvertAstronomicalToGregorianBce(
	dateTime time.Time) time.Time {

	if dtUtil.lock == nil {
		dtUtil.lock = new(sync.Mutex)
	}

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	var newDateTime time.Time

	if dateTime.Year() < 0 {
		newDateTime = time.Date(
			dateTime.Year() - 1,
			dateTime.Month(),
			dateTime.Day(),
			dateTime.Hour(),
			dateTime.Minute(),
			dateTime.Second(),
			dateTime.Nanosecond(),
			dateTime.Location())
	} else {
		newDateTime = dateTime
	}

	return newDateTime
}

// ConvertGregorianToAstronomicalBce - Used to convert negative year
// Gregorian Date Time values to their correct Astronomical Year or
// BCE equivalent by adding one year.
//
// In computing durations which extend before the common era
// ('BCE'), this type uses the Gregorian calendar system. In
// other words, under this system, there is NO year zero. The
// year 1BCE is followed by the year 1CE.
//
// See Wikipedia https://en.wikipedia.org/wiki/Year_zero :
//
//  "The year zero does not exist in the Anno Domini (AD) system
//  commonly used to number years in the Gregorian calendar and
//  in its predecessor, the Julian calendar. In this system, the
//  year 1 BC is followed by AD 1. However, there is a year zero
//  in astronomical year numbering (where it coincides with the
//  Julian year 1 BC) and in ISO 8601:2004 (where it coincides
//  with the Gregorian year 1 BC), as well as in all Buddhist
//  and Hindu calendars."
//
// For information on the "Common Era" calendar notation see:
//   https://en.wikipedia.org/wiki/Common_Era
//
func (dtUtil *DTimeUtility) ConvertGregorianToAstronomicalBce(
	dateTime time.Time) time.Time {

	if dtUtil.lock == nil {
		dtUtil.lock = new(sync.Mutex)
	}

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	var newDateTime time.Time

	if dateTime.Year() < 0 {
		newDateTime = time.Date(
			dateTime.Year() - 1,
			dateTime.Month(),
			dateTime.Day(),
			dateTime.Hour(),
			dateTime.Minute(),
			dateTime.Second(),
			dateTime.Nanosecond(),
			dateTime.Location())
	} else {
		newDateTime = dateTime
	}

	return newDateTime
}
