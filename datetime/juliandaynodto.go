package datetime

import (
	"errors"
	"fmt"
	"math"
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
	julianDayNo             *big.Int
	julianDayNoFraction     *big.Float
	julianDayNoNanoSecs     *big.Float
	julianDayNoSign         int
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

	result := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(200).
		Copy(jDNDto.julianDayNoNanoSecs)

	if jDNDto.julianDayNoSign == -1 {
		result = big.NewFloat(0).
		SetMode(big.ToNearestAway).
			SetPrec(200).
			Neg(result)
	}

	return result
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

	float64Result := 0.0

	if jDNDto.julianDayNoNanoSecs == nil {
		return float64Result,
			errors.New(ePrefix + "\n" +
				"Error: This instance of JulianDayNoDto was " +
				"incorrectly initialized and is invalid.\n" +
				"'julianDayNoNanoSecs' is nil!")
	}

	var julianDayNoSecs float64
	var accuracy big.Accuracy

	bigJulianDayNoTime := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(200).
		Set(jDNDto.julianDayNoNanoSecs)

	if jDNDto.julianDayNoSign == -1 {

		bigJulianDayNoTime = big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(200).
			Neg(bigJulianDayNoTime)
	}

	julianDayNoSecs,
		accuracy =
		bigJulianDayNoTime.Float64()

		if accuracy != big.Exact {
			return float64Result,
				fmt.Errorf(ePrefix + "\n" +
					"Error: Julian Day Number/Time could exceeds the limits\n" +
					"of a float64 at 6-digits of precision.\n" +
					"Accuracy='%v'\n", accuracy)
		}

		var roundFac float64

		roundFac = 0.0000005

		julianDayNoSecs += roundFac

		roundFac = 1000000.0

		julianDayNoSecs *= roundFac

	julianDayNoSecs = math.Floor(julianDayNoSecs)

	julianDayNoSecs /= roundFac

	if jDNDto.julianDayNoSign == -1 {

		roundFac = -1.0

		julianDayNoSecs *= roundFac
	}

	return julianDayNoSecs, nil
}

// GetHours - Returns the hours associated with this Julian
// Day Number Time instance. These hours are Gregorian Calendar
// Hours and therefore they may differ from Julian Day Number
// Time hours. 
//
// Remember that the Julian Day starts a noon, 12:00:00.000000.
// The Gregorian Calendar day starts at midnight 24:00:00.000000 or 
// 00:00:00.000000.
//
// Again this method returns Gregorian Calendar Hours.
//
func (jDNDto *JulianDayNoDto) GetHours() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	hoursInt := jDNDto.hours

		if jDNDto.julianDayNoSign == -1 {
			hoursInt *= -1
		}

	return hoursInt
}

// GetJulianTotalNanoSecondsInt64 - Returns the total nanoseconds
// associated with this Julian Day Time. The returned int64 value
// represents the total nanoseconds equaling the sum of the hours,
// minutes, seconds and nanoseconds encapsulated in this Julian Day
// Number/Time instance.
//
// Julian time represented by this total nanosecond value differs
// from Gregorian Calendar time because the Julian Day starts at
// noon (12:00:00.000000 12-hundred hours). Whereas the Gregorian
// calendar day starts at midnight (00:00:00.000000 Zero hours).
//
// This method returns the Julian time in total nanoseconds.
//
func (jDNDto *JulianDayNoDto) GetJulianTotalNanoSecondsInt64() int64 {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	totalNanoSeconds := jDNDto.totalNanoSeconds

	if jDNDto.julianDayNoSign == -1 {
		totalNanoSeconds *= -1
	}

	return totalNanoSeconds
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

	minutesInt := jDNDto.minutes

		if jDNDto.julianDayNoSign == -1 {
			minutesInt *= -1
		}

	return minutesInt
}


// GetJulianTotalNanoSecondsInt64 - Returns the total nanoseconds
// associated with this Julian Day Time. The returned int64 value
// represents the total nanoseconds equaling the sum of the hours,
// minutes, seconds and nanoseconds encapsulated in this Julian Day
// Number/Time instance as converted to a Gregorian Calendar day.
//
// Gregorian time represented by this total nanosecond value differs
// from Julian Day time because the Gregorian Day starts at midnight
// (00:00:00.000000 Zero hours). Whereas the Day starts at noon
// (12:00:00.000000 12-hundred hours).
//
// This method returns the Gregorian time in total nanoseconds.
//
func (jDNDto *JulianDayNoDto) GetGregorianTotalNanoSecondsInt64() int64 {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	result := jDNDto.totalNanoSeconds

	if result >= NoonNanoSeconds {
		result -= NoonNanoSeconds
	}

	if jDNDto.julianDayNoSign == -1 {
		result *= int64(-1)
	}

	return result
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

	secondsInt := jDNDto.seconds

	if jDNDto.julianDayNoSign == -1 {
			secondsInt *= -1
	}

	return secondsInt
}

// GetNanoseconds - Returns the internal data field
// 'Nanoseconds' from the current instance of 'JulianDayNoDto'.
//
func (jDNDto *JulianDayNoDto) GetNanoseconds() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	nanosecondsInt := jDNDto.nanoseconds

		if jDNDto.julianDayNoSign == -1 {
			nanosecondsInt *= -1
		}

	return nanosecondsInt
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