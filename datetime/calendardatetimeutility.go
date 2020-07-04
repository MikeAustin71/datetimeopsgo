package datetime

import (
	"errors"
	"fmt"
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
	calDTime.calendar = CalendarSpec(0).None()
	calDTime.yearNumberingMode = CalendarYearNumMode(0).None()
	calDTime.dateTimeFmt = ""

	return nil
}


// generateDateTimeStr - Converts input years, months, days, hours,
// minutes, seconds and nanoseconds to a formatted date time string
// the golang format string passed in input parameter 'dateFormatStr'.
//
func (calDTimeUtil *calendarDateTimeUtility) generateDateTimeStr(
	year int64,
	month,
	days,
	usDayOfWeekNumber,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	dateFormatStr,
	ePrefix string) (string, error) {

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "calendarDateTimeUtility.generateDateTimeStr() "

	var err error

	/*	replacementTokens := map[string]string{
			"!YearFourDigit!":"",
			"!YearTwoDigit!":"",
			"!YearOneDigit!":"",
			"!DayOfWeek!":"",
			"!DayOfWeekAbbrv!":"",
			"!MonthName!":"",
			"!MonthNameAbbrv!":"",
			"!MonthNumberTwoDigit!":"",
			"!MonthNumberOneDigit!":"",
			"!DateDayTwoDigit!":"",
			"!DateDayLeadUnderScore!":"",
			"!DateDayOneDigit!":"",
			"!HourTwentyFourTwoDigit!":"",
			"!HourTwelveTwoDigit!":"",
			"!HourTwelveOneDigit!":"",
			"!AMPMUpperCase!",
			"!AMPMLowerCase!",
			"!MinutesTwoDigit!":"",
			"!MinutesOneDigit!":"",
			"!SecondsTwoDigit!":"",
			"!SecondsOneDigit!":"",
			"!NanosecondsTrailingZeros!":"",
			"!NanosecondsNoTrailingZeros!":"",
			"!MillisecondsTrailingZeros!":"",
			"!MillisecondsNoTrailingZeros!":"",
			"!MicrosecondsTrailingZeros!":"",
			"!MicrosecondsNoTrailingZeros!":"",
			"!OffsetUTC!":"",
			"!TimeZone!":"",
		}
	*/

	replacementTokens := map[string]string{}

	dtMech := DTimeMechanics{}

	resultStr := dtMech.PreProcessDateFormatStr(dateFormatStr)

	calDtMech := calendarDateTimeMechanics{}

	resultStr, err = calDtMech.processDayOfWeek(
		resultStr,
		usDayOfWeekNumber,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr = calDtMech.processYears(
		resultStr,
		year,
		replacementTokens)

	resultStr, err = calDtMech.processMonths(
		resultStr,
		month,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processDateDay(
		resultStr,
		days,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processHours(
		resultStr,
		hours,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processMinutes(
		resultStr,
		minutes,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processSeconds(
		resultStr,
		seconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processNanoseconds(
		resultStr,
		nanoseconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processMicroseconds(
		resultStr,
		nanoseconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processMilliseconds(
		resultStr,
		nanoseconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processAmPm(
		resultStr,
		hours,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr = calDtMech.processOffset(
		resultStr,
		-999,
		-999,
		replacementTokens)

	resultStr = calDtMech.processTimeZone(
		resultStr,
		"UTC",
		replacementTokens)

	for key, value := range replacementTokens {

		resultStr = strings.Replace(resultStr,key,value,1)

	}

	return resultStr, err
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
	yearNumberMode CalendarYearNumMode,
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

	if !yearNumberMode.XIsValid()  {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "yearNumberMode",
			inputParameterValue: yearNumberMode.String() ,
			errMsg:              "'yearNumberMode' is INVALID!",
			err:                 nil,
		}
	}

	calMech := calendarMechanics{}

	var jDayNoDto JulianDayNoDto

	if calendar == CalendarSpec(0).Julian() {

		jDayNoDto, err = calMech.julianCalendarDateJulianDayNo(
			year,
			month,
			day,
			hours,
			minutes,
			seconds,
			nanoseconds,
			ePrefix)

	} else if calendar == CalendarSpec(0).Gregorian() {

		gregorianDateTime := time.Date(
			int(year),
			time.Month(month),
			day,
			hours,
			minutes,
			seconds,
			nanoseconds,
			time.UTC)

		_,
		jDayNoDto,
		err = calMech.gregorianDateToJulianDayNoTime(
			gregorianDateTime,
			ePrefix)

	} else {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Invalid Calendar Specification.\n" +
			"Calendar Specification='%v'\n",
			calendar.String())

		return err
	}


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

	dtMech := DTimeMechanics{}

	calDTime.year = year
	calDTime.dateTimeFmt =
		dtMech.PreProcessDateFormatStr(dateTimeFmt)
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