package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type dateTzDtoUtility struct {
	input  string
	output string
	lock   sync.Mutex
}

// addDateTime - Adds date time components to the date time value of the
// current DateTzDto instance. The updated date time value is returned to
// the calling function as a new DateTzDto instance.
//
func (dTzUtil *dateTzDtoUtility) addDateTime(
	dTz *DateTzDto,
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFormatStr,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addDateTime() "

	if dTz == nil {
		return DateTzDto{},
			errors.New(ePrefix +
				"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	newDate := dTz.dateTimeValue.AddDate(years, months, 0)

	totNanoSecs := int64(days) * DayNanoSeconds
	totNanoSecs += int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	newDateTime := newDate.Add(time.Duration(totNanoSecs))

	dTzUtil2 := dateTzDtoUtility{}

	dTz2 := DateTzDto{}

	err := dTzUtil2.setFromDateTime(&dTz2, newDateTime, dateTimeFormatStr, ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dTz2, nil
}

// addDuration - Adds Duration to the DateTime Value of the input
// parameter 'dTz' (DateTzDto) and returns a new DateTzDto instance
// with the updated Date Time value.
//
func (dTzUtil *dateTzDtoUtility) addDuration(
	dTz *DateTzDto,
	duration time.Duration,
	dateTimeFmtStr,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addDuration() "

	if dTz == nil {
		return DateTzDto{},
			errors.New(ePrefix +
				"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	newDateTime := dTz.dateTimeValue.Add(duration)

	dTzUtil2 := dateTzDtoUtility{}

	dtz2 := DateTzDto{}

	err := dTzUtil2.setFromDateTime(&dtz2, newDateTime, dateTimeFmtStr, ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// addMinusTimeDto - Creates and returns a new DateTzDto by
// subtracting a TimeDto from the value of the input
// parameter 'dTz' (DateTzDto) instance.
//
func (dTzUtil *dateTzDtoUtility) addMinusTimeDto(
	dTz *DateTzDto,
	minusTimeDto TimeDto,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addMinusTimeDto() "

	if dTz == nil {
		return DateTzDto{},
			errors.New(ePrefix +
				"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	tDto := minusTimeDto.CopyOut()

	err := tDto.NormalizeTimeElements()

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"\nError returned by tDto.NormalizeTimeElements().\n"+
				"Error='%v'\n", err.Error())
	}

	_, err = tDto.NormalizeDays()

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"\nError returned by tDto.NormalizeDays().\n"+
				"\nError='%v'\n", err.Error())
	}

	tDto.ConvertToNegativeValues()

	dt1 := dTz.dateTimeValue.AddDate(tDto.Years,
		tDto.Months,
		0)

	totNanosecs := int64(tDto.DateDays) * DayNanoSeconds
	totNanosecs += int64(tDto.Hours) * HourNanoSeconds
	totNanosecs += int64(tDto.Minutes) * MinuteNanoSeconds
	totNanosecs += int64(tDto.Seconds) * SecondNanoseconds
	totNanosecs += int64(tDto.Milliseconds) * MilliSecondNanoseconds
	totNanosecs += int64(tDto.Microseconds) * MicroSecondNanoseconds
	totNanosecs += int64(tDto.Nanoseconds)

	dt2 := dt1.Add(time.Duration(totNanosecs))

	dtz2 := DateTzDto{}
	dTzUtil2 := dateTzDtoUtility{}

	err = dTzUtil2.setFromDateTime(&dtz2,
		dt2,
		dTz.dateTimeFmt,
		ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// addPlusTimeDto - Creates and returns a new DateTzDto by adding a TimeDto
// to the value of theDateTzDto instance passed as an input parameter.
//
// The value of the input parameter DateTzDto instance is not be altered.
//
func (dTzUtil *dateTzDtoUtility) addPlusTimeDto(
	dTz *DateTzDto,
	plusTimeDto TimeDto,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addPlusTimeDto() "

	if dTz == nil {
		return DateTzDto{},
			errors.New(ePrefix +
				"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	tDto := plusTimeDto.CopyOut()

	err := tDto.NormalizeTimeElements()

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"\nError returned by tDto.NormalizeTimeElements().\n"+
				"\nError='%v'\n", err.Error())
	}

	_, err = tDto.NormalizeDays()

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"\nError returned by tDto.NormalizeDays().\n"+
				"\nError='%v'\n", err.Error())
	}

	tDto.ConvertToAbsoluteValues()

	dt1 := dTz.dateTimeValue.AddDate(tDto.Years,
		tDto.Months,
		0)

	incrementalDur := int64(tDto.DateDays) * DayNanoSeconds
	incrementalDur += int64(tDto.Hours) * HourNanoSeconds
	incrementalDur += int64(tDto.Minutes) * MinuteNanoSeconds
	incrementalDur += int64(tDto.Seconds) * SecondNanoseconds
	incrementalDur += int64(tDto.Milliseconds) * MilliSecondNanoseconds
	incrementalDur += int64(tDto.Microseconds) * MicroSecondNanoseconds
	incrementalDur += int64(tDto.Nanoseconds)

	dt2 := dt1.Add(time.Duration(incrementalDur))

	dTz2 := DateTzDto{}

	dTzUtil2 := dateTzDtoUtility{}

	err = dTzUtil2.setFromDateTime(&dTz2, dt2, dTz.dateTimeFmt, ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dTz2, nil
}

// addTime - Adds input parameter time components (hours, minutes, seconds etc.)
// to the date time value of the input parameter 'dTz' (DateTzDto). The resulting
// updated date time value is returned to the calling function in the form of a
// new DateTzDto instance.
//
// The value of the input parameter 'dTz' (DateTzDto) instance is NOT altered.
//
func (dTzUtil *dateTzDtoUtility) addTime(
	dTz *DateTzDto,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFormatStr,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addTime() "

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	newDateTime := dTz.dateTimeValue.Add(time.Duration(totNanoSecs))

	dTzUtil2 := dateTzDtoUtility{}

	dtz2 := DateTzDto{}

	err := dTzUtil2.setFromDateTime(&dtz2, newDateTime, dateTimeFormatStr, ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// copyIn - Receives two parameters which are pointers
// to types DateTzDto. The method then copies all of
// the data field values from 'incomingDtz' into
// 'baseDtz'.
//
func (dTzUtil *dateTzDtoUtility) copyIn(
	baseDtz,
	incomingDtz *DateTzDto) {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	dTzUtil2 := dateTzDtoUtility{}

	if baseDtz == nil ||
		incomingDtz == nil {
		return
	}

	dTzUtil2.empty(baseDtz)

	baseDtz.tagDescription = incomingDtz.tagDescription
	baseDtz.timeComponents = incomingDtz.timeComponents.CopyOut()
	baseDtz.dateTimeFmt = incomingDtz.dateTimeFmt

	if incomingDtz.dateTimeValue.IsZero() {
		baseDtz.timeZone = TimeZoneDefDto{}
		baseDtz.dateTimeValue = time.Time{}
	} else {
		baseDtz.dateTimeValue = incomingDtz.dateTimeValue
		baseDtz.timeZone = incomingDtz.timeZone.CopyOut()
	}

}

// copyOut - Returns a deep copy of input parameter
// 'dTz' which is a pointer to a type 'DateTzDto'.
func (dTzUtil *dateTzDtoUtility) copyOut(
	dTz *DateTzDto) DateTzDto {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	dtz2 := DateTzDto{}

	if dTz == nil {
		return dtz2
	}

	dtz2.tagDescription = dTz.tagDescription
	dtz2.timeComponents = dTz.timeComponents.CopyOut()
	dtz2.dateTimeFmt = dTz.dateTimeFmt

	if dTz.dateTimeValue.IsZero() {
		dtz2.timeZone = TimeZoneDefDto{}
		dtz2.dateTimeValue = time.Time{}
	} else {
		dtz2.dateTimeValue = dTz.dateTimeValue
		dtz2.timeZone = dTz.timeZone.CopyOut()
	}

	return dtz2
}

// empty - Receives a pointer to a type 'DateTzDto' and
// proceeds to set all internal member variables to their
// 'zero' or uninitialized values.
//
func (dTzUtil *dateTzDtoUtility) empty(dTz *DateTzDto) {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	if dTz == nil {
		return
	}

	dTz.tagDescription = ""
	dTz.timeComponents.Empty()
	dTz.timeZone = TimeZoneDefDto{}
	dTz.dateTimeValue = time.Time{}
	dTz.dateTimeFmt = ""

	return
}

// isEmptyDateTzDto - Analyzes an instanceof DateTzDto to
// determine if all data fields are uninitialized or zero
// values.
//
func (dTzUtil *dateTzDtoUtility) isEmptyDateTzDto(
	dTz *DateTzDto) bool {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	if dTz == nil {
		return true
	}

	if dTz.tagDescription == "" &&
		dTz.timeComponents.IsEmpty() &&
		dTz.dateTimeValue.IsZero() &&
		dTz.dateTimeFmt == "" &&
		dTz.timeZone.IsEmpty() {

		return true
	}

	return false
}

// isValidDateTzDto - Analyzes an instance of 'DateTzDto' to
// determine if is value. If the instance evaluates as invalid,
// an error is returned.
//
func (dTzUtil *dateTzDtoUtility) isValidDateTzDto(
	dTz *DateTzDto,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.isValidDateTzDto() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	dTzUtil2 := dateTzDtoUtility{}

	if dTzUtil2.isEmptyDateTzDto(dTz) {
		return errors.New(ePrefix +
			"\nThis 'DateTzDto' instance is EMPTY!\n")
	}

	if dTz.dateTimeValue.IsZero() {
		return errors.New(ePrefix +
			"\nError: DateTzDto.DateTime is ZERO!\n")
	}

	if dTz.timeZone.IsEmpty() {
		return errors.New(ePrefix +
			"\nError: DateTzDto.TimeZone is EMPTY!\n")
	}

	if err := dTz.timeComponents.IsValid(); err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: dTz.timeComponents is INVALID.\n"+
			"Error='%v'\n", err.Error())
	}


	/* TODO - Fix IsValidFromDateTime()
	if !dTz.timeZone.IsValidFromDateTime(dTz.dateTimeValue) {
		return errors.New(ePrefix +
			"\nError: dTz.TimeZone is INVALID!\n")
	}
*/

	return nil
}

// preProcessDateFormatStr - Provides a standardized method
// for implementing a default date time format string.
//
func (dTzUtil *dateTzDtoUtility) preProcessDateFormatStr(
	dateTimeFmtStr string) string {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	dateTimeFmtStr = strings.TrimLeft(strings.TrimRight(dateTimeFmtStr, " "), " ")

	if len(dateTimeFmtStr) == 0 {
		return FmtDateTimeYrMDayFmtStr
	}

	return dateTimeFmtStr
}

// setFromDateTime - Sets the values for DateTzDto fields encapsulated
// in input parameter 'dTz'. The field values will be set
// based on an input parameter 'dateTime' (Type time.time) and a
// date time format string.
func (dTzUtil *dateTzDtoUtility) setFromDateTime(
	dTz *DateTzDto,
	dateTime time.Time,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromDateTime() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input Parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dateTime.IsZero() {
		return errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is ZERO!\n")
	}

	dTzUtil2 := dateTzDtoUtility{}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	tDto := TimeDto{}

	tDtoUtil := timeDtoUtility{}

	err := tDtoUtil.setFromDateTime(&tDto, dateTime, ePrefix)

	if err != nil {
		return err
	}

	timeZone := TimeZoneDefDto{}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromDateTime(&timeZone, dateTime, ePrefix)

	if err != nil {
		return err
	}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = dateTime
	dTz.timeComponents = tDto.CopyOut()
	dTz.timeZone = timeZone.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}

// SetFromDateTimeComponents - Sets the values of the Date Time fields
// for the current DateTzDto instance based on time components
// and a Time Zone Location.
//
// Note that this variation of time elements breaks time down by
// hour, minute, second, millisecond, microsecond and nanosecond.
//
func (dTzUtil *dateTzDtoUtility) setFromDateTimeComponents(
	dTz *DateTzDto,
	year,
	month,
	day,
	hour,
	minute,
	second,
	millisecond,
	microsecond,
	nanosecond int,
	timeZoneName,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromDateTimeComponents() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	tDto, err := TimeDto{}.New(year, month, 0, day, hour, minute,
		second, millisecond, microsecond, nanosecond)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by TimeDto{}.New(year, month,...).  "+
			"Error='%v'", err.Error())
	}

	dTzUtil2 := dateTzDtoUtility{}

	fmtStr := dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	dtUtil := DTimeUtility{}

	tzl := dtUtil.PreProcessTimeZoneLocation(timeZoneName)

	_, err = time.LoadLocation(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by time.LoadLocation(tzl).\nINVALID 'timeZoneName'!\n"+
			"tzl='%v'\ntimeZoneName='%v'\nError='%v'\n",
			tzl, timeZoneName, err.Error())
	}

	var dt time.Time

	dt, err = tDto.GetDateTime(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tDto.GetDateTime(tzl).\n"+
			"\ntimeZoneName='%v'\ntzl='%v'\nError='%v'\n",
			timeZoneName, tzl, err.Error())
	}

	timeZone := TimeZoneDefDto{}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromTimeZoneName(
		&timeZone,
		dt,
		timeZoneName,
		TzConvertType.Absolute(),
		ePrefix)

	if err != nil {
		return err
	}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = dt
	dTz.timeZone = timeZone.CopyOut()
	dTz.timeComponents = tDto.CopyOut()
	dTz.dateTimeFmt = fmtStr

	return nil
}

// setFromDateTimeElements - Sets the values of input parameter
// 'dTz' (type DateTzDto). 'dTz' data fields are set based on
// input parameters consisting of date time elements,
// a time zone location and a date time format string.
//
// Date Time elements include year, month, day, hour, minute,
// second and nanosecond.
//
func (dTzUtil *dateTzDtoUtility) setFromDateTimeElements(
	dTz *DateTzDto,
	year,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	timeZoneName,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromDateTimeElements() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	tDto, err := TimeDto{}.New(year, month, 0, day, hour, minute, second,
		0, 0, nanosecond)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned from TimeDto{}.New(year, month, ...).\n"+
			"Error='%v'\n", err.Error())
	}

	dTzUtil2 := dateTzDtoUtility{}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	dtUtil := DTimeUtility{}

	timeZoneName = dtUtil.PreProcessTimeZoneLocation(timeZoneName)

	_, err = time.LoadLocation(timeZoneName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by time.LoadLocation(tzl).\n"+
			"INVALID 'timeZoneName'!\n"+
			"tzl='%v'\ntimeZoneName='%v'\n"+
			"Error='%v'\n",
			timeZoneName, timeZoneName, err.Error())
	}

	var dt time.Time

	dt, err = tDto.GetDateTime(timeZoneName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tDto.GetDateTime(tzl).\n"+
			"\ntimeZoneName='%v'\ntzl='%v'\n"+
			"Error='%v'\n",
			timeZoneName, timeZoneName, err.Error())
	}

	timeZone := TimeZoneDefDto{}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromTimeZoneName(
		&timeZone,
		dt,
		timeZoneName,
		TzConvertType.Absolute(),
		ePrefix)

	if err != nil {
		return err
	}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = dt
	dTz.timeZone = timeZone.CopyOut()
	dTz.timeComponents = tDto.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}

// setFromTimeTz - Sets the time values for input parameter 'dTz'
// (type DateTzDto). The new values will be  based on input parameters
// 'dateTime', 'timeZoneLocation' and a date time format string,
// 'dateTimeFmtStr'.
//
func (dTzUtil *dateTzDtoUtility) setFromTimeTz(
	dTz *DateTzDto,
	dateTime time.Time,
	timeZoneLocationName string,
	timeZoneConversionType TimeZoneConversionType,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromTimeTz() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dateTime.IsZero() {
		return errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is ZERO and INVALID!\n")
	}

	tZoneDefDto := TimeZoneDefDto{}

	tzDefUtil := timeZoneDefUtility{}

	var err error

	err = tzDefUtil.setFromTimeZoneName(
		&tZoneDefDto, dateTime,
		timeZoneLocationName,
		timeZoneConversionType,
		ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: Invalid Time Zone Location Name!\n"+
			"timeZoneLocationName='%v'\nError='%v'\n",
			timeZoneLocationName, err.Error())
	}

	var targetDateTime time.Time

	if timeZoneConversionType == TzConvertType.Absolute() {
		// FmtDateTimeTzNanoYMD = "2006-01-02 15:04:05.000000000 -0700 MST"
		dtUtil := DTimeUtility{}
		targetDateTime, err = dtUtil.AbsoluteTimeToTimeZoneDtoConversion(dateTime, tZoneDefDto)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by dtUtil.AbsoluteTimeToTimeZoneDtoConversion(dateTime,tZoneDefDto)\n"+
				"dateTime='%v'\ntZoneDefDto='%v'\nError='%v'\n",
				dateTime.Format(FmtDateTimeTzNanoYMD), tZoneDefDto.GetLocationName(), err.Error())
		}
	} else {
		// Must be TzConvertType.Relative() or TzConvertType.None()
		// This the default.
		targetDateTime = dateTime.In(tZoneDefDto.GetLocationPtr())
	}

	var tDto TimeDto

	tDtoUtil := timeDtoUtility{}

	err =tDtoUtil.setFromDateTime(&tDto, targetDateTime, ePrefix)

	if err != nil {
		return  err
	}

	dTzUtil2 := dateTzDtoUtility{}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = targetDateTime
	dTz.timeZone = tZoneDefDto
	dTz.timeComponents = tDto
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}

// setFromTimeDto - Receives data from a TimeDto input parameter
// and sets all data fields of the current DateTzDto instance
// accordingly. When the method completes, the values of the
// current DateTzDto will equal the values of the input parameter
// TimeDto instance.
//
func (dTzUtil *dateTzDtoUtility) setFromTimeDto(
	dTz *DateTzDto,
	tDto TimeDto,
	timeZoneLocation,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	ePrefix += "dateTzDtoUtility.setFromTimeDto() "

	if tDto.IsEmpty() {

		return fmt.Errorf(ePrefix + "\nError: Input parameter 'tDto' date time elements equal ZERO!\n")
	}

	t2Dto := tDto.CopyOut()

	var err error

	err = t2Dto.NormalizeTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by t2Dto.NormalizeTimeElements().\nError='%v'\n",
			err.Error())
	}

	t2Dto.ConvertToAbsoluteValues()

	if err = t2Dto.IsValid(); err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: Input Parameter tDto (TimeDto) is INVALID.\nError='%v'\n",
			err.Error())
	}

	dtUtil := DTimeUtility{}

	timeZoneLocation = dtUtil.PreProcessTimeZoneLocation(timeZoneLocation)

	_, err = time.LoadLocation(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by time.LoadLocation(timeZoneLocation).\n"+
			"timeZoneLocation='%v'\nError='%v'\n", timeZoneLocation, err.Error())
	}

	dTzUtil2 := dateTzDtoUtility{}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	var dateTime time.Time

	dateTime, err = tDto.GetDateTime(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tDto.GetDateTime(timeZoneLocation).\n"+
			"timeZoneLocation='%v'\nError='%v'\n",
			timeZoneLocation, err.Error())
	}

	timeZoneDef := TimeZoneDefDto{}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromDateTime(&timeZoneDef, dateTime, ePrefix)

	if err != nil {
		return err
	}

	dTzUtil2.empty(dTz)
	dTz.dateTimeValue = dateTime
	dTz.timeZone = timeZoneDef.CopyOut()
	dTz.timeComponents = t2Dto.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}
