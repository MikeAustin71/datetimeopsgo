package datetime

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"
)

type calendarDateTimeUtility struct {
	lock   sync.Mutex
}


// copyOut - Returns a deep copy of input parameter
// 'calDTime' which is a pointer to a type 'CalendarDateTime'.
//
func (calDTimeUtil *calendarDateTimeUtility) copyOut(
	calDTime *CalendarDateTime,
	ePrefix string) ( newCalDTime CalendarDateTime, err error) {

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "calendarDateTimeUtility.copyOut() "

	newCalDTime = CalendarDateTime{}

	if calDTime == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'calDTime' is a nil pointer!")

		return newCalDTime, err
	}

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	newCalDTime.dateTime = calDTime.dateTime.CopyOut()
	newCalDTime.julianDayNumber = calDTime.julianDayNumber.CopyOut()
	newCalDTime.usDayOfWeekNo = calDTime.usDayOfWeekNo
	newCalDTime.timeZone = calDTime.timeZone.CopyOut()
	newCalDTime.calendar = calDTime.calendar
	newCalDTime.yearNumberingMode = calDTime.yearNumberingMode
	newCalDTime.dateTimeFmt = calDTime.dateTimeFmt

	return newCalDTime, err
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

	calDTime.dateTime.Empty()
	calDTime.julianDayNumber = JulianDayNoDto{}.NewZero()
	var err error

	calDTime.timeZone, err = TimeZoneDefinition{}.New()

	if err != nil {
		calDTime.timeZone = TimeZoneDefinition{}
	}

	calDTime.usDayOfWeekNo = UsWeekDayNo.None()
	calDTime.calendar = CalendarSpec(0).None()
	calDTime.yearNumberingMode = CalendarYearNumMode(0).None()
	calDTime.dateTimeFmt = ""

	return nil
}


// generateDateTimeStr - Converts input years, months, days, hours,
// minutes, seconds and subMicrosecondNanoseconds to a formatted date time string
// the golang format string passed in input parameter 'dateFormatStr'.
//
func (calDTimeUtil *calendarDateTimeUtility) generateDateTimeStr(
	year int64,
	month,
	days int,
	usDayOfWeekNumber UsDayOfWeekNo,
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
	hour,
	minute,
	second,
	nanosecond int,
	timeZoneLocation string,
	calendar CalendarSpec,
	dateTimeFmt    string,
	ePrefix string) error {

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "calendarDateTimeUtility.setCalDateTime() "

	calDTimeUtil2 := calendarDateTimeUtility{}

	err := calDTimeUtil2.empty(calDTime, ePrefix)

	if err != nil {
		return err
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

	timeZone := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}

	baseDateTime := time.Now().UTC()

	err = tzDefUtil.setFromTimeZoneName(
		&timeZone,
		baseDateTime,
		TimeZoneConversionType(0).Relative(),
		timeZoneLocation,
		ePrefix)

	if err != nil {
		return err
	}

	isLeapYear := false

	switch calendar {

	case CalSpec.Gregorian():

		calGregMech := calendarGregorianMechanics{}

		isLeapYear = calGregMech.isLeapYear(year)

	default:
		return fmt.Errorf(ePrefix + "\n" +
			"Error: Only Gregorian Calendar is currently supported!\n" +
			"calendar='%s'\n", calendar.String())
	}

	calUtil := CalendarUtility{}

	err = calUtil.IsValidDateTimeComponents(
		isLeapYear,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		ePrefix)

	if err != nil {
		return err
	}

	var jDayNoDto JulianDayNoDto

	yearNumberingMode := CalYearMode.None()

	calMech := calendarMechanics{}

	if calendar == CalendarSpec(0).Julian() {

		jDayNoDto, err = calMech.julianCalendarDateJulianDayNo(
			year,
			month,
			day,
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

		if err != nil {
			return err
		}

		yearNumberingMode = CalYearMode.Astronomical()

	} else if calendar == CalendarSpec(0).Gregorian() {

		var julianDayNo int64
		var julianDayNoTimeFraction *big.Float

		calGregUtil := CalendarGregorianUtility{}

		julianDayNo,
		_,
		julianDayNoTimeFraction,
		err =	calGregUtil.GetJDN(
			year,
			month,
			day,
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

		jDNDtoUtil := julianDayNoDtoUtility{}

		err = jDNDtoUtil.setDto(
			&jDayNoDto,
			julianDayNo,
			julianDayNoTimeFraction,
			ePrefix)

		yearNumberingMode = CalYearMode.Astronomical()

	} else if calendar == CalendarSpec(0).RevisedGoucherParker() {


		jDayNoDto, err = calMech.revisedGoucherParkerToJulianDayNo(
			year,
			month,
			day,
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

		if err != nil {
			return err
		}

		yearNumberingMode = CalYearMode.Astronomical()

	} else if calendar == CalendarSpec(0).RevisedJulian() {


		yearNumberingMode = CalYearMode.Astronomical()

	} else {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Invalid Calendar Specification.\n" +
			"Calendar Specification='%v'\n",
			calendar.String())

		return err
	}

	dtMech := DTimeMechanics{}

	var usDayOfWeekNo UsDayOfWeekNo

	usDayOfWeekNo, err =
		calMech.usDayOfWeekNumber(jDayNoDto, ePrefix)

	if err != nil {
		return err
	}

	calDTime.dateTime.isLeapYear = isLeapYear
	calDTime.dateTime.year = year
	calDTime.dateTime.month = month
	calDTime.dateTime.day = day
	calDTime.dateTime.hour = hour
	calDTime.dateTime.minute = minute
	calDTime.dateTime.second = second
	calDTime.dateTime.nanosecond = nanosecond
	calDTime.dateTime.isThisInstanceValid = true
	calDTime.lock = new(sync.Mutex)
	calDTime.julianDayNumber = jDayNoDto.CopyOut()
	calDTime.usDayOfWeekNo = usDayOfWeekNo
	calDTime.timeZone = timeZone.CopyOut()
	calDTime.calendar = calendar
	calDTime.yearNumberingMode = yearNumberingMode
	calDTime.dateTimeFmt =
		dtMech.PreProcessDateFormatStr(dateTimeFmt)

	return nil
}