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