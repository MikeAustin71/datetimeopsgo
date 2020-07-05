package datetime

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type DTimeMechanics struct {
	lock *sync.Mutex
}

// AbsoluteTimeToTimeZoneDefConversion - Converts a given time to
// another time zone using the 'Absolute' conversion method.
// This means that the years, months, days, hours, minutes,
// seconds and subMicrosecondNanoseconds of the original 'dateTime' are not
// changed. That time value is simply assigned to another
// designated time zone. The target time zone is derived from
// input parameter 'timeZoneDefDto', an instance of type
// 'TimeZoneDefinition'.
//
// For example, assume that 'dateTime' represents 10:00AM in USA
// time zone 'Central Standard Time'.  Using the 'Absolute'
// conversion method, and converting this time value to the USA
// Eastern Standard Time Zone would result in a date time of
// 10:00AM EST or Eastern Standard Time. The time value of
// 10:00AM is not changed, it is simply assigned to another
// time zone.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// dateTime            time.Time
//                      - The date time to be converted.
//
// timeZoneDefDto TimeZoneDefinition
//                      - A properly initialized 'TimeZoneDefinition'
//                        object encapsulating the time zone to which
//                       'dateTime' will be converted.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// time.Time  - The date time converted to the time zone specified in
//              in input parameter 'timeZoneDefDto'.
//
// error      - If the method completes successfully this value is set
//              to 'nil'. If an error is encountered, the returned error
//              value encapsulates an appropriate error message.
//
func (dtMech *DTimeMechanics) AbsoluteTimeToTimeZoneDefConversion(
	dateTime time.Time,
	timeZoneDefDto TimeZoneDefinition) (time.Time, error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix := "DTimeMechanics.AbsoluteTimeToTimeZoneDefConversion() "

	if err := timeZoneDefDto.IsValid(); err != nil {
		return time.Time{},
			fmt.Errorf(ePrefix +
				"Input parameter 'timeZoneDefDto' is Invalid!\n" +
				"Error='%v'\n", err.Error())
	}

	return time.Date(dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		dateTime.Hour(),
		dateTime.Minute(),
		dateTime.Second(),
		dateTime.Nanosecond(),
		timeZoneDefDto.GetOriginalLocationPtr()), nil
}

// AbsoluteTimeToTimeZoneNameConversion - Converts a given time to
// another time zone using the 'Absolute' conversion method.
// This means that the years, months, days, hours, minutes,
// seconds and subMicrosecondNanoseconds of the original 'dateTime' are not
// changed. That time value is simply assigned to another
// designated time zone. The target time zone is derived from
// input parameter 'timeZoneDefDto', an instance of type
// 'TimeZoneDefinition'.
//
// For example, assume that 'dateTime' represents 10:00AM in USA
// time zone 'Central Standard Time'.  Using the 'Absolute'
// conversion method, and converting this time value to the USA
// Eastern Standard Time Zone would result in a date time of
// 10:00AM EST or Eastern Standard Time. The time value of
// 10:00AM is not changed, it is simply assigned to another
// time zone.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// dateTime            time.Time  - The date time to be converted.
//
// timeZoneName           string  - A string containing a valid IANA,
//                                  Military or "Local" time zone.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// time.Time  - The date time converted to the time zone specified in
//              in input parameter 'timeZoneDefDto'.
//
// error      - If the method completes successfully this value is set
//              to 'nil'. If an error is encountered, the returned error
//              value encapsulates an appropriate error message.
//
func (dtMech *DTimeMechanics) AbsoluteTimeToTimeZoneNameConversion(
	dateTime time.Time,
	timeZoneName string,
	ePrefix string) (time.Time, error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.AbsoluteTimeToTimeZoneNameConversion() "

	timeZoneName = strings.TrimRight(strings.TrimLeft(timeZoneName, " "), " ")

	if len(timeZoneName) == 0 {
		return time.Time{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeZoneName",
				inputParameterValue: "",
				errMsg:              "Error: 'timeZoneName' is an empty string!",
				err:                 nil,
			}
	}

	var err error
	tzSpec := TimeZoneSpecification{}
	tzMech := TimeZoneMechanics{}

	tzSpec,
	err = tzMech.GetTimeZoneFromName(
		dateTime,
		timeZoneName,
		TzConvertType.Absolute(),
		ePrefix)

	if err != nil {
		return time.Time{}, err
	}

	return tzSpec.referenceDateTime, nil
}

// AddDateTime - Adds date and time values to an existing date time,
// 'baseDateTime'.
//
func (dtMech *DTimeMechanics) AddDateTime(
	baseDateTime time.Time,
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int) time.Time {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	var newDate time.Time

	if years > 0 ||
			months > 0 {

		newDate = baseDateTime.AddDate(years, months, 0)

	} else {

		newDate = baseDateTime

	}

	totNanoSecs := int64(days) * DayNanoSeconds
	totNanoSecs += int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	return newDate.Add(time.Duration(totNanoSecs))
}

// AddDateTimeByLocalTimeZone - Performs an addition operation
// by adding time components to a specified 'baseDateTime'. The
// addition operation is performed using the local time zone
// associated with input parameter 'baseDateTime'.
//
func (dtMech *DTimeMechanics) AddDateTimeByLocalTimeZone(
	baseDateTime time.Time,
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int) time.Time {


	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	var newDateTime time.Time

	if years != 0 ||
		months != 0 ||
		days != 0 {

		newDateTime = baseDateTime.AddDate(
			years,
			months,
			days)

	} else {

		newDateTime = baseDateTime

	}

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	return newDateTime.Add(time.Duration(totNanoSecs))
}

// AddDateTimeByUtc - Adds date and time values to an existing date time,
// 'baseDateTime'. 'baseDateTime' is first converted to equivalent
// UTC date time before the addition. Afterwards, the date time
// generated by the addition operation is converted back to the
// original time zone. This technique produces improved accuracy
// by eliminating anomalies caused by Daylight Savings Time.
//
func (dtMech *DTimeMechanics) AddDateTimeByUtc(
	baseDateTime time.Time,
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int) time.Time {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	var resultDateTime, newDateTime, baseDateTimeUtc time.Time

	baseDateTimeUtc = baseDateTime.In(time.UTC)

	if years != 0 ||
			months != 0 {

		newDateTime = baseDateTimeUtc.AddDate(years, months, 0)

	} else {

		newDateTime = baseDateTimeUtc

	}

	totNanoSecs := int64(days) * DayNanoSeconds
	totNanoSecs += int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	resultDateTime = newDateTime.Add(time.Duration(totNanoSecs))

	return resultDateTime.In(baseDateTime.Location())
}

// AddDurationByLocalTimeZone - Receives a date time input
// parameter ('baseDateTime') and proceeds to add input
// parameter 'timeDuration'. The local time zone is utilized
// in the addition operation.
//
func (dtMech *DTimeMechanics) AddDurationByLocalTimeZone(
	baseDateTime time.Time,
	timeDuration time.Duration) time.Time {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	return baseDateTime.Add(timeDuration)
}

// AddDurationByUtc - Receives a base date time and
// converts that date time to its UTC equivalent. Then,
// the time duration is added. After the addition
// operation the new date is converted back to the
// original time zone and returned.
//
func (dtMech *DTimeMechanics) AddDurationByUtc(
	baseDateTime time.Time,
	timeDuration time.Duration) time.Time {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	utcEquivalentDateTime := baseDateTime.In(time.UTC)

	utcTimePlusDuration := utcEquivalentDateTime.Add(timeDuration)

	return utcTimePlusDuration.In(baseDateTime.Location())
}

// AllocateSecondsToHrsMinSecs - Useful in calculating offset hours,
// minutes and seconds from UTC+0000. A total signed seconds value
// is passed as an input parameter. This method then breaks down
// hours, minutes and seconds as positive integer values. The sign
// of the hours, minutes and seconds is returned in the 'sign'
// parameter as +1 or -1.
//
func (dtMech *DTimeMechanics) AllocateSecondsToHrsMinSecs(
	signedTotalSeconds int) (hours, minutes, seconds, sign int) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

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

// ComputeDurationFromBaseTime - Computes time duration using
// using a Base Date Time and a time duration. The algorithm
// used to compute staring date time and ending date time is
// specified by input parameter 'timeMathCalcMode'.
//
func (dtMech *DTimeMechanics) ComputeDurationFromBaseTime(
	baseTime time.Time,
	duration time.Duration,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr string,
	ePrefix string) (
	newStartDateTime DateTzDto,
	newEndDateTime DateTzDto,
	err error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.ComputeDurationFromBaseTime() "

	newStartDateTime = DateTzDto{}

	newEndDateTime = DateTzDto{}

	err = nil

	if !timeMathCalcMode.XIsValid() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeMathCalcMode",
			inputParameterValue: timeMathCalcMode.String(),
			errMsg:              "Input Parameter 'timeMathCalcMode' is INVALID!",
			err:                 nil,
		}

		return newStartDateTime, newEndDateTime, err
	}

	var newBaseTime, startDateTime, endDateTime time.Time

	if timeMathCalcMode == TCalcMode.LocalTimeZone() {

		newBaseTime = baseTime

	} else if timeMathCalcMode == TCalcMode.UtcTimeZone() {

		newBaseTime = baseTime.In(time.UTC)

	} else {

		err = fmt.Errorf(ePrefix +
			"\nError: Input parameter 'timeMathCalcMode' is not equal to\n" +
			"'LocalTimeZone' or 'UtcTimeZone'\n")

		return newStartDateTime, newEndDateTime, err
	}

	if duration < 0 {

		endDateTime = newBaseTime

		startDateTime = endDateTime.Add(duration)

		duration = duration * -1

	} else {

		startDateTime = newBaseTime

		endDateTime = startDateTime.Add(duration)

	}

	tzMech := TimeZoneMechanics{}

	var tzSpec TimeZoneSpecification

	tzSpec, err = tzMech.GetTimeZoneFromName(
		baseTime,
		timeZoneLocation,
		TzConvertType.Relative(),
		ePrefix)

	if err != nil {
		return newStartDateTime, newEndDateTime, err
	}

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTzSpec(
		&newStartDateTime,
		startDateTime,
		tzSpec,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return newStartDateTime, newEndDateTime, err
	}

	err = dTzUtil.setFromTzSpec(
		&newEndDateTime,
		endDateTime,
		tzSpec,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		newStartDateTime = DateTzDto{}
		return newStartDateTime, newEndDateTime, err
	}

	return newStartDateTime, newEndDateTime, err
}

// ComputeDurationFromStartEndTimes - Computes time duration using
// the algorithm specified by input parameter
// 'timeMathCalcMode'.
//
func (dtMech *DTimeMechanics) ComputeDurationFromStartEndTimes(
	startTime,
	endTime time.Time,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr,
	ePrefix string) (
	duration time.Duration,
	newStartDateTime DateTzDto,
	newEndDateTime DateTzDto,
	err error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.ComputeDurationFromStartEndTimes() "

	duration = time.Duration(0)

	newStartDateTime = DateTzDto{}

	newEndDateTime = DateTzDto{}

	err = nil

	var tempStartTime, tempEndTime time.Time

	if endTime.Before(startTime) {

		tempStartTime = endTime
		tempEndTime = startTime

	} else {

		tempStartTime = startTime
		tempEndTime = endTime

	}

	if !timeMathCalcMode.XIsValid() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeMathCalcMode",
			inputParameterValue: timeMathCalcMode.String(),
			errMsg:              "Input Parameter 'timeMathCalcMode' is INVALID!",
			err:                 nil,
		}

		return duration, newStartDateTime, newEndDateTime, err
	}

	dtMech2 := DTimeMechanics{}

	if timeMathCalcMode == TCalcMode.LocalTimeZone() {

		return dtMech2.ComputeDurationByLocalTz(
			tempStartTime,
			tempEndTime,
			timeZoneLocation,
			dateTimeFmtStr,
			ePrefix)

	} else if timeMathCalcMode == TCalcMode.UtcTimeZone() {

		return dtMech2.ComputeDurationByUtc(
			tempStartTime,
			tempEndTime,
			timeZoneLocation,
			dateTimeFmtStr,
			ePrefix)

	}

	err = fmt.Errorf(ePrefix +
		"\nError: Input parameter 'timeMathCalcMode' is not equal to\n" +
		"'LocalTimeZone' or 'UtcTimeZone'\n")

	return duration, newStartDateTime, newEndDateTime, err
}

// ComputeDurationByLocalTz - Computes time duration by local
// time zone. The local time zone is specified by the input
// string parameter 'timeZoneLocation'.
//
func (dtMech *DTimeMechanics) ComputeDurationByLocalTz(
	startTime,
	endTime time.Time,
	timeZoneLocation,
	dateTimeFmtStr,
	ePrefix string) (
	                 duration time.Duration,
	                 newStartDateTime DateTzDto,
	                 newEndDateTime DateTzDto,
	                 err error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.ComputeDurationByLocalTz() "

	duration = time.Duration(0)

	newStartDateTime = DateTzDto{}

	newEndDateTime = DateTzDto{}

	err = nil

	var tempStartTime, tempEndTime time.Time

	if endTime.Before(startTime) {

		tempStartTime = endTime
		tempEndTime = startTime

	} else {

		tempStartTime = startTime
		tempEndTime = endTime

	}

	tzMech := TimeZoneMechanics{}

	var tzSpec TimeZoneSpecification

	tzSpec, err = tzMech.GetTimeZoneFromName(
		tempStartTime,
		timeZoneLocation,
		TzConvertType.Relative(),
		ePrefix)

	if err != nil {
		return duration, newStartDateTime, newEndDateTime, err
	}

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTzSpec(
		&newStartDateTime,
		tempStartTime,
		tzSpec,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return duration, newStartDateTime, newEndDateTime, err
	}

	err = dTzUtil.setFromTzSpec(
		&newEndDateTime,
		tempEndTime,
		tzSpec,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		newStartDateTime = DateTzDto{}
		return duration, newStartDateTime, newEndDateTime, err
	}

	duration =  newEndDateTime.dateTimeValue.
								Sub(newStartDateTime.dateTimeValue)

	return duration, newStartDateTime, newEndDateTime, err
}

// ComputeDurationByUtc - Computes time duration by first
// converting parameters 'startTime' and 'EndTime' to UTC
// Time Zone and then computing duration. Return values
// 'newStartDateTime' and 'newEndDateTime' are then computed
// by converting from UTC time to the local time zone designated
// by input string parameter, 'timeZoneLocation'.
//
func (dtMech *DTimeMechanics) ComputeDurationByUtc(
	startTime,
	endTime time.Time,
	timeZoneLocation,
	dateTimeFmtStr,
	ePrefix string) (duration time.Duration,
	                 newStartDateTime DateTzDto,
	                 newEndDateTime DateTzDto,
	                 err error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.ComputeDurationByUtc() "

	duration = time.Duration(0)

	newStartDateTime = DateTzDto{}

	newEndDateTime = DateTzDto{}

	err = nil

	var tempStartTime, tempEndTime time.Time

	if endTime.Before(startTime) {

		tempStartTime = endTime
		tempEndTime = startTime

	} else {

		tempStartTime = startTime
		tempEndTime = endTime

	}

	tzMech := TimeZoneMechanics{}

	var tzSpec TimeZoneSpecification

	tzSpec, err = tzMech.GetTimeZoneFromName(
		startTime,
		timeZoneLocation,
		TzConvertType.Relative(),
		ePrefix)

	if err != nil {
		return duration, newStartDateTime, newEndDateTime, err
	}

	startTimeUtc := tempStartTime.In(time.UTC)

	endTimeUtc := tempEndTime.In(time.UTC)

	duration = endTimeUtc.Sub(startTimeUtc)

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTzSpec(
		&newStartDateTime,
		startTimeUtc,
		tzSpec,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		duration = time.Duration(0)
		return duration, newStartDateTime, newEndDateTime, err
	}

	err = dTzUtil.setFromTzSpec(
		&newEndDateTime,
		endTimeUtc,
		tzSpec,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		duration = time.Duration(0)
		newStartDateTime = DateTzDto{}
		return duration, newStartDateTime, newEndDateTime, err
	}

	return duration, newStartDateTime, newEndDateTime, err
}

// GetDurationFromTimeComponents - Receives time components
// such as days, hours, minutes, seconds, milliseconds,
// microseconds and subMicrosecondNanoseconds and returns time duration
// (time.Duration)
//
func (dtMech *DTimeMechanics) GetDurationFromTimeComponents(
	days ,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int) time.Duration {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	totNanosecs := int64(days) * DayNanoSeconds
	totNanosecs += int64(hours) * HourNanoSeconds
	totNanosecs += int64(minutes) * MinuteNanoSeconds
	totNanosecs += int64(seconds) * SecondNanoseconds
	totNanosecs += int64(milliseconds) * MilliSecondNanoseconds
	totNanosecs += int64(microseconds) * MicroSecondNanoseconds
	totNanosecs += int64(nanoseconds)

	return time.Duration(totNanosecs)
}

// GetTimeZoneFromDateTime - Analyzes a date time object
// and returns a valid time zone in the form of a
// 'TimeZoneSpecification' instance.
//
// Because date time objects (time.Time) do not support
// Military Time Zones; therefore, Military Time Zones
// are never returned by this method.
//
func (dtMech *DTimeMechanics) GetTimeZoneFromDateTime(
	dateTime time.Time,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.GetTimeZoneFromName() "

	tzMech := TimeZoneMechanics{}

	return tzMech.GetTimeZoneFromDateTime(dateTime, ePrefix)
}

// GetTimeZoneFromName - Analyzes a time zone name passed
// through input parameter, 'timeZoneName'. If valid, the
// method populates and returns a 'TimeZoneSpecification'
// instance.
//
// This method will accept and successfully process one
// of three types of time zones:
//
//   (1) The time zone "Local", which Golang accepts as
//       the time zone currently configured on the host
//       computer.
//
//   (2) IANA Time Zone - A valid IANA Time Zone from the
//       IANA database.
//       See https://golang.org/pkg/time/#LoadLocation
//       and https://www.iana.org/time-zones to ensure that
//       the IANA Time Zone Database is properly configured
//       on your system.
//
//       IANA Time Zone Examples:
//         "America/New_York"
//         "America/Chicago"
//         "America/Denver"
//         "America/Los_Angeles"
//         "Pacific/Honolulu"
//         "Etc/UTC" = GMT or UTC
//
//    (3) A Military Time Zone
//        Reference:
//         https://en.wikipedia.org/wiki/List_of_military_time_zones
//         http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//         https://www.timeanddate.com/time/zones/military
//         https://www.timeanddate.com/worldclock/timezone/alpha
//         https://www.timeanddate.com/time/map/
//
//        Examples:
//          "Alpha"   or "A"
//          "Bravo"   or "B"
//          "Charlie" or "C"
//          "Delta"   or "D"
//          "Zulu"    or "Z"
//
// If the time zone "Zulu" is passed to this method, it will be
// classified as a Military Time Zone.
//
func (dtMech *DTimeMechanics) GetTimeZoneFromName(
	dateTime time.Time,
	timeZoneName string,
	timeConversionType TimeZoneConversionType,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.GetTimeZoneFromName() "

	tzSpec = TimeZoneSpecification{}
	err = nil

	if timeConversionType < TzConvertType.Absolute() ||
		timeConversionType > TzConvertType.Relative() {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeConversionType",
			inputParameterValue: timeConversionType.String(),
			errMsg:              "Input Parameter 'timeConversionType' " +
				"contains an invalid value!",
			err:                 nil,
		}
		return tzSpec, err
	}

	if len(timeZoneName) == 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "Input parameter is an EMPTY String!",
			err:                 nil,
		}

		return tzSpec, err
	}

	tzMech := TimeZoneMechanics{}

	return tzMech.GetTimeZoneFromName(
		dateTime,
		timeZoneName,
		timeConversionType,
		ePrefix)
}

// GetUtcOffsetTzAbbrvFromDateTime - Receives a time.Time,
// date time, input parameter and extracts and returns the
// 5-character UTC offset and the time zone abbreviation.
//
// UTC Offsets are returned in the format illustrated by the
// following examples:
//
//   +1030
//   -0500
//   +1100
//   -1100
//
// The time zone abbreviation,'tzAbbrv', is formatted as
// shown in the following example ('CST').
//
// Example:
//  Time String:  2019-12-26 00:56:15 -0600 CST
//  Returned UTC Offset:  '-0600'
//  Returned Time Zone Abbreviation: 'CST'
//
func (dtMech *DTimeMechanics) GetUtcOffsetTzAbbrvFromDateTime(
	dateTime time.Time,
	ePrefix string) (
	utcOffset,
	tzAbbrv string,
	err error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	utcOffset = ""
	tzAbbrv = ""
	err = nil

	ePrefix += "DTimeMechanics.GetUtcOffsetTzAbbrvFromDateTime() "

	tStr := dateTime.Format("2006-01-02 15:04:05 -0700 MST")

	lenLeadTzAbbrvStr := len("2006-01-02 15:04:05 -0700 ")

	tzAbbrv = tStr[lenLeadTzAbbrvStr:]

	tzAbbrv = strings.TrimRight(tzAbbrv, " ")

	lenLeadOffsetStr := len("2006-01-02 15:04:05 ")

	utcOffset = tStr[lenLeadOffsetStr : lenLeadOffsetStr+5]

	return utcOffset, tzAbbrv, err
}

// LoadTzLocation - Provides a single method for calling
// time.LoadLocation(). This method may be altered in the future
// to load time zones from an internal file thus affording
// consistency in time zone definitions without relying on
// zoneinfo.zip databases residing on host computers.
//
// If successful, this method returns a *time.Location or
// location pointer to a valid time zone.
//
func (dtMech *DTimeMechanics) LoadTzLocation(
	timeZoneName string,
	ePrefix string) (*time.Location, error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.LoadTzLocation() "

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

// ConvertTimeToNewTimeZoneName - Receives a valid time (time.Time)
// value and changes the existing time zone to that specified
// in parameter 'tZoneLocationName'.
//
// Input Parameters
// ================
//
//   dateTime time.Time
//          - Initial time whose time zone will be changed to
//            target time zone input parameter, 'tZoneLocationName'
//
//
//   timeConversionType TimeZoneConversionType
//          - This parameter determines the algorithm that will
//            be used to convert parameter 'dateTime' to the time
//            zone specified by parameter 'timeZoneName'.
//
//            TimeZoneConversionType is an enumeration type which
//            be used to convert parameter 'dateTime' to the time
//            must be set to one of two values:
//            This parameter determines the algorithm that will
//               TimeZoneConversionType(0).Absolute()
//               TimeZoneConversionType(0).Relative()
//            Note: You can also use the global variable
//            'TzConvertType' for easier access:
//               TzConvertType.Absolute()
//               TzConvertType.Relative()
//
//            Absolute Time Conversion - Identifies the 'Absolute' time
//            to time zone conversion algorithm. This algorithm provides
//            that a time value in time zone 'X' will be converted to the
//            same time value in time zone 'Y'.
//
//            For example, assume the time 10:00AM is associated with time
//            zone USA Central Standard time and that this time is to be
//            converted to USA Eastern Standard time. Applying the 'Absolute'
//            algorithm would convert ths time to 10:00AM Eastern Standard
//            time.  In this case the hours, minutes and seconds have not been
//            altered. 10:00AM in USA Central Standard Time has simply been
//            reclassified as 10:00AM in USA Eastern Standard Time.
//
//            Relative Time Conversion - Identifies the 'Relative' time to time
//            zone conversion algorithm. This algorithm provides that times in
//            time zone 'X' will be converted to their equivalent time in time
//            zone 'Y'.
//
//            For example, assume the time 10:00AM is associated with time zone
//            USA Central Standard time and that this time is to be converted to
//            USA Eastern Standard time. Applying the 'Relative' algorithm would
//            convert ths time to 11:00AM Eastern Standard time. In this case the
//            hours, minutes and seconds have been changed to reflect an equivalent
//            time in the USA Eastern Standard Time Zone.
//
// tZoneLocationName string
//          - The first input time value, 'tIn' will have its time zone
//            changed to a new time zone location specified by this second
//            parameter, 'tZoneLocation'. This time zone location must be
//            designated as one of three types of time zones:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                configured for the host computer executing this code.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                   Examples:
//                     "America/New_York"
//                     "America/Chicago"
//                     "America/Denver"
//                     "America/Los_Angeles"
//                     "Pacific/Honolulu"
//
//            (3) A valid Military Time Zone
//                Military time zones are commonly used in
//                aviation as well as at sea. They are also
//                known as nautical or maritime time zones.
//                Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//
//             Note:
//                 The source file 'timezonedata.go' contains over 600 constant
//                 time zone declarations covering all IANA and Military Time
//                 Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TZones'.
//
func (dtMech *DTimeMechanics) ConvertTimeToNewTimeZoneName(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	tZoneLocationName string) (time.Time, error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix := "DTimeMechanics.ConvertTimeToNewTimeZoneName() "

	tzMech := TimeZoneMechanics{}

	tzSpec,
	err := tzMech.GetTimeZoneFromName(
		dateTime,
		tZoneLocationName,
		timeConversionType,
		ePrefix)

	if err != nil {
		return time.Time{}, err
	}

	return tzSpec.GetReferenceDateTime(), nil
}

// RelativeTimeToTimeNameZoneConversion - Converts a time value
// to its equivalent time in another time zone specified by input
// parameter string, 'timeZoneName'.
//
// The 'timeZoneName' string must specify one of three types of
// time zones:
//
//   (1) The string 'Local' - selects the local time zone
//                            location for the host computer.
//
//   (2) IANA Time Zone Location -
//      See https://golang.org/pkg/time/#LoadLocation
//      and https://www.iana.org/time-zones to ensure that
//      the IANA Time Zone Database is properly configured
//      on your system. Note: IANA Time Zone Data base is
//      equivalent to 'tz database'.
//     Examples:
//      "America/New_York"
//      "America/Chicago"
//      "America/Denver"
//      "America/Los_Angeles"
//      "Pacific/Honolulu"
//      "Etc/UTC" = GMT or UTC
//
//    (3) A Military Time Zone
//        Reference:
//         https://en.wikipedia.org/wiki/List_of_military_time_zones
//         http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//         https://www.timeanddate.com/time/zones/military
//         https://www.timeanddate.com/worldclock/timezone/alpha
//         https://www.timeanddate.com/time/map/
//
//        Examples:
//          "Alpha"   or "A"
//          "Bravo"   or "B"
//          "Charlie" or "C"
//          "Delta"   or "D"
//          "Zulu"    or "Z"
//
func (dtMech *DTimeMechanics) RelativeTimeToTimeNameZoneConversion(
	dateTime time.Time,
	timeZoneName string,
	ePrefix string) (time.Time, error) {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.RelativeTimeToTimeNameZoneConversion() "

	timeZoneName = strings.TrimRight(strings.TrimLeft(timeZoneName, " "), " ")

	if len(timeZoneName) == 0 {
		return time.Time{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeZoneName",
				inputParameterValue: "",
				errMsg:              "Error: 'timeZoneName' is an empty string!",
				err:                 nil,
			}
	}

	var err error
	tzSpec := TimeZoneSpecification{}
	tzMech := TimeZoneMechanics{}

	tzSpec,
	err = tzMech.GetTimeZoneFromName(
		dateTime,
		timeZoneName,
		TzConvertType.Relative(),
		ePrefix)

	if err != nil {
		return time.Time{}, err
	}

	return tzSpec.referenceDateTime, nil
}

// PreProcessDateFormatStr - If parameter, 'dateTimeFmtStr'
// is determined, a default value is substituted.
//
func (dtMech *DTimeMechanics) PreProcessDateFormatStr(
	dateTimeFmtStr string) string {

	if dtMech.lock == nil {
		dtMech.lock = new(sync.Mutex)
	}

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()


	dateTimeFmtStr = strings.TrimLeft(strings.TrimRight(dateTimeFmtStr, " "), " ")

	if len(dateTimeFmtStr) == 0 {
		lockDefaultDateTimeFormat.Lock()
		dateTimeFmtStr = DEFAULTDATETIMEFORMAT
		lockDefaultDateTimeFormat.Unlock()
	}

	return dateTimeFmtStr
}
