package datetime

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"
)

type calendarDateTimeUtility struct {
	lock   sync.Mutex
}

// empty - Receives a pointer to a CalendarDateTime instance
// and proceeds to set the internal data elements to their
// zero values.
func (calDTimeUtil *calendarDateTimeUtility) empty(
	calDTime *CalendarDateTime,
	ePrefix string) error {

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "calendarDateTimeUtility.empty() "

	if calDTime == nil {
		return errors.New(ePrefix +
			"\nInput parameter 'calDTime' is a nil pointer!")
	}

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.year = 0
	calDTime.month = 0
	calDTime.dateDays = 0
	calDTime.hours = 0
	calDTime.minutes = 0
	calDTime.seconds = 0
	calDTime.milliseconds = 0
	calDTime.microseconds = 0
	calDTime.nanoseconds = 0
	calDTime.totSubSecNanoseconds = 0
	calDTime.totTimeNanoseconds = 0
	calDTime.dateTimeFmt = ""

	return nil
}
// preProcessDateFormatStr - Provides a standardized method
// for implementing a default date time format string.
//
func (calDTimeUtil *calendarDateTimeUtility) preProcessDateFormatStr(
	dateTimeFmtStr string) string {

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	dateTimeFmtStr = strings.TrimLeft(strings.TrimRight(dateTimeFmtStr, " "), " ")

	if len(dateTimeFmtStr) == 0 {
		lockDefaultDateTimeFormat.Lock()
		dateTimeFmtStr = DEFAULTDATETIMEFORMAT
		lockDefaultDateTimeFormat.Unlock()
	}

	return dateTimeFmtStr
}

// setCalDateTime - populates a CalendarDateTime instance.
//
func (calDTimeUtil *calendarDateTimeUtility) setCalDateTime(
	calDTime *CalendarDateTime,
	year int64,
	month,
	day,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	timeZoneLocation string,
	calendar CalendarSpec,
	dateTimeFmt    string,
	ePrefix string) error {

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.setZeroTimeDto() "

	calDTimeUtil2 := calendarDateTimeUtility{}

	err := calDTimeUtil2.empty(calDTime, ePrefix)

	if err != nil {
		return err
	}

	if month < 1 || month > 12 {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "month",
			inputParameterValue: strconv.Itoa(month) ,
			errMsg:              "'month' is INVALID!",
			err:                 nil,
		}
	}

	if day < 1 || day > 31 {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "day",
			inputParameterValue: strconv.Itoa(day) ,
			errMsg:              "'day' is INVALID!",
			err:                 nil,
		}
	}

	if hours < 0 || hours > 23 {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "hours",
			inputParameterValue: strconv.Itoa(hours) ,
			errMsg:              "'hours' is INVALID!",
			err:                 nil,
		}
	}

	if minutes < 0 || minutes > 59 {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "minutes",
			inputParameterValue: strconv.Itoa(minutes) ,
			errMsg:              "'minutes' is INVALID!",
			err:                 nil,
		}
	}

	if seconds < 0 || seconds > 59 {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "seconds",
			inputParameterValue: strconv.Itoa(seconds) ,
			errMsg:              "'seconds' is INVALID!",
			err:                 nil,
		}
	}

	if nanoseconds < 0 || nanoseconds > 999999999 {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "nanoseconds",
			inputParameterValue: strconv.Itoa(nanoseconds) ,
			errMsg:              "'nanoseconds' is INVALID!",
			err:                 nil,
		}
	}

	if !calendar.XIsValid()  {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "calendar",
			inputParameterValue: calendar.String() ,
			errMsg:              "'calendar' is INVALID!",
			err:                 nil,
		}
	}
	calMech := calendarMechanics{}

	var jDayNoDto JulianDayNoDto

	jDayNoDto, err = calMech.julianCalendarDateJulianDayNo(
		year,
		month,
		day,
		hours,
		minutes,
		seconds,
		nanoseconds,
		ePrefix)

	if err != nil {
		return err
	}

	timeZone := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromTimeZoneName(
		&timeZone,
		time.Now().UTC(),
		TimeZoneConversionType(0).Relative(),
		timeZoneLocation,
		ePrefix)

	if err != nil {
		return err
	}

	calDTime.year = year
	calDTime.dateTimeFmt = calDTimeUtil2.preProcessDateFormatStr(dateTimeFmt)
	calDTime.month = month
	calDTime.dateDays = day
	calDTime.hours = hours
	calDTime.minutes = minutes
	calDTime.seconds = seconds
	calDTime.timeZone = timeZone.CopyOut()
	calDTime.calendar = calendar
	calDTime.julianDayNumber = jDayNoDto

	return nil
}