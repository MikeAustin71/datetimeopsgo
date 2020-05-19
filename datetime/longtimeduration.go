package datetime

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

// LongTimeDuration - This type provides an alternative to time.Duration.
// time.Duration is limited in range to approximately 290-years. This type
// allows time duration computations which cover billions of years. This
// type is used exclusively in calculating time duration.
//
type LongTimeDuration struct {

	startDateTimeTz DateTzDto // Starting Date Time

	endDateTimeTz DateTzDto // Ending Date Time

	duration         *big.Int     // Time Duration expressed as a big.Int.
	                              // 'duration' is stored as a positive value

	sign             int          // Numerical sign of 'duration'. -1, 0, +1

	lock             *sync.Mutex  // Used for coordinating thread safe operations.
}

// New - Returns a new LongTimeDuration instance initialized to
// 'Zero' duration.
func (lngDur LongTimeDuration) New() (LongTimeDuration, error) {

	if lngDur.lock == nil {
		lngDur.lock = new(sync.Mutex)
	}

	lngDur.lock.Lock()

	defer lngDur.lock.Unlock()

	longDur := LongTimeDuration{}

	longDur.lock = new(sync.Mutex)

	ePrefix := "LongTimeDuration.New() "

	lngDurUtil := longTimeDurationUtility{}

	longDur2 := LongTimeDuration{}

	err := lngDurUtil.setZeroLongTermDuration(
		&longDur2,
		ePrefix)

	if err != nil {
		return LongTimeDuration{}, err
	}

	return longDur2, nil
}

// NewFromDuration - Initializes and returns a new instance of LongTimeDuration.
//
func (lngDur LongTimeDuration) NewFromDuration(
	startDateTime time.Time,
	duration time.Duration) (LongTimeDuration, error) {

	if lngDur.lock == nil {
		lngDur.lock = new(sync.Mutex)
	}

	lngDur.lock.Lock()

	ePrefix := "LongTimeDuration.NewFromDuration() "

	defer lngDur.lock.Unlock()

	longDur := LongTimeDuration{}

	longDur.lock = new(sync.Mutex)

	if duration < 0 {
		longDur.sign = -1
		duration = duration * -1
	} else if duration == 0 {
		longDur.sign = 0
	} else {
		longDur.sign = 1
	}

	longDur.duration = big.NewInt(0).SetInt64(int64(duration))

	lngDurUtil := longTimeDurationUtility{}

	var err error
	var finalStartDateTime, finalEndDateTime time.Time

	if longDur.sign == -1 {

		finalStartDateTime,
			finalEndDateTime,
		err = lngDurUtil.getStartDateMinusDuration(
			&longDur,
			startDateTime,
			ePrefix)

	} else {

		finalStartDateTime,
			finalEndDateTime,
		err = lngDurUtil.getStartDatePlusDuration(
			&longDur,
			startDateTime,
			ePrefix)

	}

	if err != nil {
		return LongTimeDuration{}, err
	}

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTimeTzName(
		&longDur.startDateTimeTz,
		finalStartDateTime,
		TzConvertType.Relative(),
		startDateTime.Location().String(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return LongTimeDuration{}, err
	}

	err = dTzUtil.setFromTimeTzName(
		&longDur.endDateTimeTz,
		finalEndDateTime,
		TzConvertType.Relative(),
		startDateTime.Location().String(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return LongTimeDuration{}, err
	}

	return longDur, nil
}

// NewStartEndDateTimes - Computes long duration based on starting and ending
// date times.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date times.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     time.Time
//     - Starting date time
//
//
//  endDateTime       time.Time
//     - Ending date time
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  LongTimeDuration
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'LongTimeDuration' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  longDur, err := LongTimeDuration{}.NewStartEndDateTimes(
//                                     startDateTime,
//                                     endDateTime,
//                                     TZones.US.Central(),
//                                     FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (lngDur LongTimeDuration) NewStartEndDateTimes(
	startDateTime,
	endDateTime time.Time,
	timeZoneLocation string,
	dateTimeFmtStr string) (LongTimeDuration, error) {

	if lngDur.lock == nil {
		lngDur.lock = new(sync.Mutex)
	}

	lngDur.lock.Lock()

	defer lngDur.lock.Unlock()

	ePrefix := "LongTimeDuration.NewStartEndDateTimes() "

	longDur2 := LongTimeDuration{}

	lngDurUtil := longTimeDurationUtility{}

	err := lngDurUtil.setStartEndDateDuration(
		&longDur2,
		startDateTime,
		endDateTime,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return LongTimeDuration{}, err
	}

	return longDur2, nil
}


// AddDuration - Adds a time.Duration value to the total time duration
// maintained by the current instance of LongTimeDuration.
//
func (lngDur *LongTimeDuration) AddDuration(
	duration time.Duration) error {

	if lngDur.lock == nil {
		lngDur.lock = new(sync.Mutex)
	}

	lngDur.lock.Lock()

	defer lngDur.lock.Unlock()

	ePrefix := "LongTimeDuration.AddDuration() "

	lngDurUtil := longTimeDurationUtility{}

	return lngDurUtil.addDuration(
		lngDur,
		duration,
		ePrefix)
}

// IsValid - Analyzes the current instance of LongTimeDuration.
// If the time duration value is zero, the returned boolean value
// is set to true.
//
// If the current instance of LongTimeDuration is judged as 'invalid',
// an appropriate error message is returned.
//
func (lngDur *LongTimeDuration) IsValid() (isZero bool, err error) {

	if lngDur.lock == nil {
		lngDur.lock = new(sync.Mutex)
	}

	lngDur.lock.Lock()

	defer lngDur.lock.Unlock()

	ePrefix := "LongTimeDuration.IsValid() "

	lngDurUtil := longTimeDurationUtility{}

	isZero, err = lngDurUtil.isValid(
		lngDur,
		ePrefix)

	return isZero, err
}


// GenerateDatePlusDuration - Receives a starting date time and adds the
// duration value contained in the current LongTimeDuration instance.
// The date time generated by adding the current time duration value
// is returned as a type time.Time.
//
// This method will correctly process negative time duration values.
//
func (lngDur *LongTimeDuration) GenerateDatePlusDuration(
	baseDateTime time.Time) (finalStartDateTime, finalEndDateTime time.Time, err error) {

	if lngDur.lock == nil {
		lngDur.lock = new(sync.Mutex)
	}

	lngDur.lock.Lock()

	defer lngDur.lock.Unlock()

	ePrefix := "LongTimeDuration.GenerateDatePlusDuration() "

	if lngDur.duration == nil {
		return time.Time{},
		       time.Time{},
		       fmt.Errorf(ePrefix + "\n" +
			"Error: 'LongTimeDuration.duration' is nil and has NOT been initialized!\n")
	}

	lngDurUtil := longTimeDurationUtility{}

	if lngDur.sign == -1 {

		return lngDurUtil.getStartDateMinusDuration(
			lngDur,
			baseDateTime,
			ePrefix)
	}

	return lngDurUtil.getStartDatePlusDuration(
		lngDur,
		baseDateTime,
		ePrefix)
}

// GetLongDuration - Returns the time duration stored in this
// LongTimeDuration instance. Two values are returned. The
// first value, 'duration', is always returned as a positive
// value.
//
// Adding 'duration' to the starting date time will always
// yield the ending date time.
//
// The second value returned is the 'sign' value. If duration
// was originally configured as a negative number, 'sign' is
// set to '-1'. If 'duration' is a zero value, 'sign' will be
// set to zero ('0'). If 'duration' was originally submitted
// or subsequently computed, as a positive value, 'sign' will
// be set to set to plus one ('1').
//
func (lngDur *LongTimeDuration) GetLongDuration() (
	duration *big.Int, sign int) {

	if lngDur.lock == nil {
		lngDur.lock = new(sync.Mutex)
	}

	lngDur.lock.Lock()

	defer lngDur.lock.Unlock()


	bigZero := big.NewInt(0)


	duration = bigZero.Set(lngDur.duration)

	sign = lngDur.sign

	return duration, sign
}

// GetStartEndDatesTz - Returns two DateTzDto objects identifying
// the starting and ending date times for the duration represented
// by the current LongTimeDuration instance. The dates returned
// reflect the results of a mathematical calculation which adds
// a positive 
// the
//
func (lngDur *LongTimeDuration) GetStartEndDatesTz() (startDateTimeTz, endDateTimeTz DateTzDto){

	if lngDur.lock == nil {
		lngDur.lock = new(sync.Mutex)
	}

	lngDur.lock.Lock()

	defer lngDur.lock.Unlock()

	startDateTimeTz =	lngDur.startDateTimeTz.CopyOut()
	endDateTimeTz = lngDur.endDateTimeTz.CopyOut()

	return startDateTimeTz, endDateTimeTz
}

// GetStartEndDateTimes - Returns two time.Time values identifying
// the starting and ending date times for the duration represented
// by the current LongTimeDuration instance.
//
func (lngDur *LongTimeDuration) GetStartEndDateTimes() (startDateTime, endDateTime time.Time) {


	if lngDur.lock == nil {
		lngDur.lock = new(sync.Mutex)
	}

	lngDur.lock.Lock()

	defer lngDur.lock.Unlock()

	startDateTime = lngDur.startDateTimeTz.dateTimeValue
	endDateTime = lngDur.endDateTimeTz.dateTimeValue

	return startDateTime, endDateTime
}