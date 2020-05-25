package datetime

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
)

// JulianDayNoDto - This type is used to transfer information
// on a Julian Day Number/Time.
//
// ------------------------------------------------------------------------
//
// Background
//
// The Julian Day Number (JDN) is the integer assigned to a
// whole solar day in the Julian day count starting from noon
// Universal time, with Julian day number 0 assigned to the
// day starting at noon on Monday, January 1, 4713 BC, proleptic
// Julian calendar (November 24, 4714 BC, in the proleptic
// Gregorian calendar), a date which preceded any dates in
// recorded history. For example, the Julian day number for
// the day starting at 12:00 UT (noon) on January 1, 2000, was
// 2 451 545.
//
// The Julian date (JD), or Julian Day Number/Time, of any instant
// is the Julian day number/ plus the fraction of a day since the
// preceding noon in Universal Time. Julian dates are expressed as
// a Julian day number with a decimal fraction added. For example,
// the Julian Date for 00:30:00.0 UT January 1, 2013, is
// 2 456 293.520 833.
//
// The Julian day number is based on the Julian Period proposed
// by Joseph Scaliger, a classical scholar, in 1583 (one year after
// the Gregorian calendar reform) as it is the product of three
// calendar cycles used with the Julian calendar.
//
//
// ------------------------------------------------------------------------
//
// Technical Considerations
//
// The Julian Day Number Time is a floating point number with an integer
// to the left of the decimal point representing the Julian Day Number
// and the fraction to the right of the decimal point representing time
// in hours minutes and seconds.
//
// Julian Day numbers start on day zero at noon. This means that Julian
// Day Number Times are valid for all dates on or after noon on Monday,
// January 1, 4713 BCE, in the proleptic Julian calendar or November 24,
// 4714 BCE, in the proleptic Gregorian calendar. Remember that the Golang
// 'time.Time' type uses Astronomical Year numbering with the Gregorian
// Calendar. In other words, the 'time.Time' type recognizes the year
// zero. Dates expressed in the 'Common Era' ('BCE' Before Common Era
// or 'CE' Common Era). Therefore a 'time.Time' year of '-4713' is equal
// to the year '4714 BCE'
//
// The 'JulianDayNoDto' type provides a Julian Day Number/Time as a float64.
// This version of the Julian Day Number/Time is accurate to within 1-second.
// In addition, the 'JulianDayNoDto' type also provides a Julian Day Number/Time
// stored as *big.Float type. This version of the Julian Day Number/Time is
// accurate to within one nanosecond.
//
//
//
// ------------------------------------------------------------------------
//
// Resources
//
// For more information on the Julian Day Number/Time see:
//   https://en.wikipedia.org/wiki/Julian_day
//
type JulianDayNoDto struct {
	julianDayNo             int64
	julianDayNoFraction     *big.Float
	julianDayNoNanoSecs     *big.Float
	totalNanoSeconds        int64
	hours                   int
	minutes                 int
	seconds                 int
	nanoseconds             int
	lock                    *sync.Mutex
}

// GetDayNoTimeNanosecs - Returns a *big.Float type representing
// the Julian Day Number/Time to an accuracy of nanoseconds.
//
// If the current instance of type JulianDayNoDto has been incorrectly
// initialized, this method will return the value of positive infinity.
//
func (jDNDto *JulianDayNoDto) GetDayNoTimeNanosecs() *big.Float {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	if jDNDto.julianDayNoNanoSecs == nil {
		return big.NewFloat(0.0).SetInf(true)
	}

	return big.NewFloat(0.0).Copy(jDNDto.julianDayNoNanoSecs)
}

// GetDayNoTimeSeconds - Returns a float64 value representing
// the Julian Day Number/Time to the nearest second.
//
// Typically a Julian Day Number/Time value can be represented
// by a float64 with 6-decimals to the right of the decimal
// place. However, if the day number is extremely large the
// value may exceed the limits of a float64 type. In that case,
// an error is returned.
//
func (jDNDto *JulianDayNoDto) GetDayNoTimeSeconds() (
	float64, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.GetDayNoTimeSeconds() "

	if jDNDto.julianDayNoNanoSecs == nil {
		return float64(0.0),
			errors.New(ePrefix + "\n" +
				"Error: This instance of JulianDayNoDto was " +
				"incorrectly initialized and is invalid.\n" +
				"'julianDayNoNanoSecs' is nil!")
	}

	var julianDayNoSecs float64
	var accuracy big.Accuracy

	bigJulianDayNoTime := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(6).
		Set(jDNDto.julianDayNoNanoSecs)

	julianDayNoSecs,
		accuracy =
		bigJulianDayNoTime.Float64()

		if accuracy != big.Exact {
			return float64(0.0),
				fmt.Errorf(ePrefix + "\n" +
					"Error: Julian Day Number/Time could exceeds the limits\n" +
					"of a float64 at 6-digits of precision.\n" +
					"Accuracy='%v'\n", accuracy)
		}

	return julianDayNoSecs, nil
}

// GetHours - Returns the internal data field
// 'hours' from the current instance of 'JulianDayNoDto'.
//
func (jDNDto *JulianDayNoDto) GetHours() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	return jDNDto.hours
}

// GetMinutes - Returns the internal data field
// 'minutes' from the current instance of 'JulianDayNoDto'.
//
func (jDNDto *JulianDayNoDto) GetMinutes() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	return jDNDto.minutes
}

// GetSeconds - Returns the internal data field
// 'seconds' from the current instance of 'JulianDayNoDto'.
//
func (jDNDto *JulianDayNoDto) GetSeconds() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	return jDNDto.seconds
}

// GetNanoseconds - Returns the internal data field
// 'Nanoseconds' from the current instance of 'JulianDayNoDto'.
//
func (jDNDto *JulianDayNoDto) GetNanoseconds() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	return jDNDto.nanoseconds
}

// New - Returns a new, populated instance of type
// JulianDayNoDto.
//
func (jDNDto JulianDayNoDto) New(
	julianDayNo int64,
	julianDayNoTimeFraction *big.Float) (
	JulianDayNoDto,
	error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.New() "

	julianDayNoDto := JulianDayNoDto{}

	jDNDtoUtil := julianDayNoDtoUtility{}

	err := jDNDtoUtil.setDto(
		&julianDayNoDto,
		julianDayNo,
		julianDayNoTimeFraction,
		ePrefix)

	return julianDayNoDto, err
}

// GetTimeFraction - Returns the fractional part of Julian Day
// Number/Time a type *big.Float. The integer portion of this
// this fractional number is always zero.
//
// If the current instance of type JulianDayNoDto has NOT been
// correctly initialized, this method will return positive
// infinity.
//
func (jDNDto *JulianDayNoDto) GetTimeFraction() *big.Float {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	if jDNDto.julianDayNoFraction == nil {
		return big.NewFloat(0.0).SetInf(true)
	}

	return big.NewFloat(0.0).Copy(jDNDto.julianDayNoFraction)

}