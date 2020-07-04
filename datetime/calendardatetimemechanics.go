package datetime

import (
	"fmt"
	"strings"
	"sync"
)

type calendarDateTimeMechanics struct {
	lock   sync.Mutex
}

// isGregorianLeapYear - Returns true if the year number is
// a leap year under the Gregorian Calendar.
//
// In the Gregorian calendar, three criteria must be taken
// into account to identify leap years:
//
// The year must be evenly divisible by 4;
//
// If the year can also be evenly divided by 100, it is not a leap year;
// unless...
//
// The year is also evenly divisible by 400. Then it is a leap year.
//
//
// According to these rules, the years 2000 and 2400 are leap years,
// while 1800, 1900, 2100, 2200, 2300, and 2500 are not leap years.
//
// Reference:
//   https://www.timeanddate.com/date/leapyear.html
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//
func (calDtMech *calendarDateTimeMechanics) isGregorianLeapYear(
	year int64) bool {

	var by4Remainder, by100Remainder int64

	by100Remainder = year % 100

	if by100Remainder == 0 {

		if year % 400 == 0 {
			return true
		}

		return false
	}

	by4Remainder = year % 4

	if by4Remainder == 0 {
		return true
	}

	return false
}

// isJulianLeapYear - Determines whether the input parameter
// 'year' is a leap year under the Julian Calendar.
//
// If 'year' is a Julian leap year, this method returns 'true'.
//
// Note: This method should NOT be used to determine leap years
// for the Revised Julian Calendar or the Goucher-Parker
// calendar.
//
// Reference:
//   https://en.wikipedia.org/wiki/Julian_calendar
//
func (calDtMech *calendarDateTimeMechanics) isJulianLeapYear(
	year int64) bool {

	remainder := year % 4

	if remainder == 0 {
		return true
	}

	return false
}

// isRevisedJulianLeapYear - Determines whether the input
// parameter 'year' is a leap year under the Revised Julian
// Calendar.
//
// The Revised Julian calendar has the same months and month
// lengths as the Julian calendar, but, in the Revised Julian
// calendar, years evenly divisible by 100 are not leap years,
// except that years with remainders of 200 or 600 when divided
// by 900 remain leap years, e.g. 2000 and 2400 as in the Gregorian
// Calendar.
//
// For additional information, reference:
//    https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
// Summary
//
// 1. Years evenly divisible by 4 are leap years unless they are
//    century years.
//
// 2. Years evenly divisible by 100 are not leap years unless when
//    divided by 900 those years have remainders of 200 or 600 in
//    which case they are leap years.
//
func (calDtMech *calendarDateTimeMechanics) isRevisedJulianLeapYear(
	year int64) bool {

	var by4Remainder, by100Remainder, by900Remainder int64

	by100Remainder = year % 100

	if by100Remainder == 0 {

		by900Remainder = year % 900

		if by900Remainder == 200 ||
			by900Remainder == 900 {
			return true
		}

		return false
	}

	by4Remainder = year % 4

	if by4Remainder == 0 {
		return true
	}

	return false
}

func (calDtMech *calendarDateTimeMechanics) isGoucherParkerLeapYear(
	year int64) bool {

	return false
}

// isMonthDayNoValid - Tests a month and day combination to
// determine if they are valid. If month and day are valid,
// this method returns true.
//
func (calDtMech *calendarDateTimeMechanics) isMonthDayNoValid(
	monthNo int,
	dayNo int,
	isLeapYear bool) bool {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	if dayNo < 1 {
		return false
	}

	// Month No : Standard Num Of Days In Month
	standardMthDays := map[int]int {
		1 : 31,
		2 : 28,
		3 : 31,
		4 : 30,
		5 : 31,
		6 : 30,
		7 : 31,
		8 : 31,
		9 : 30,
		10 : 31,
		11 : 30,
		12 : 31,
	}

	var ok bool
	var stdDays int

	stdDays, ok = standardMthDays[monthNo]

	if !ok {
		return false
	}

	if monthNo == 2 &&
		isLeapYear {
		stdDays++
	}

	if dayNo > stdDays {
		return false
	}

	return true
}

// processAmPm - processes and returns correct AM/PM format
// for 12-Hour presentations.
//
func (calDtMech *calendarDateTimeMechanics) processAmPm(
	inputStr string,
	hourNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "calendarDateTimeMechanics.processAmPm() "

	if hourNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hourNumber' is LESS THAN ONE!\n" +
			"hourNumber='%v'\n", hourNumber)

		return resultStr, err
	}

	if hourNumber > 23 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hourNumber' is GREATER THAN 23!\n" +
			"hourNumber='%v'\n", hourNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, "PM") {

		amPmStr := "PM"

		if hourNumber < 1 && hourNumber > 0 {
			amPmStr = "AM"
		}

		tokenMap["!AMPMUpperCase!"] = amPmStr

		resultStr = strings.Replace(resultStr,
			"PM",
			"!AMPMUpperCase!",
			1)

	} else if strings.Contains(resultStr, "pm") {

		amPmStr := "pm"

		if hourNumber < 1 && hourNumber > 0 {
			amPmStr = "am"
		}

		tokenMap["!AMPMLowerCase!"] = amPmStr

		resultStr = strings.Replace(resultStr,
			"pm",
			"!AMPMLowerCase!",
			1)

	}

	return resultStr, err
}

// processDateDay - processes and returns correct date day format
//
func (calDtMech *calendarDateTimeMechanics) processDateDay(
	inputStr string,
	dateDayNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "calendarDateTimeMechanics.processDateDay() "

	if dateDayNumber < 1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'dateDayNumber' is LESS THAN ONE!\n" +
			"dateDayNumber='%v'\n", dateDayNumber)

		return resultStr, err
	}

	if dateDayNumber > 31 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'dateDayNumber' is GREATER THAN 31!\n" +
			"dateDayNumber='%v'\n", dateDayNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, "02") {

		dateDayStr := fmt.Sprintf("%02d",
			dateDayNumber)

		tokenMap["!DateDayTwoDigit!"] = dateDayStr

		resultStr = strings.Replace(resultStr,
			"02",
			"!DateDayTwoDigit!",
			1)

	} else if strings.Contains(resultStr, "_2") {

		var dateDayStr string

		if dateDayNumber < 10 {
			dateDayStr = fmt.Sprintf("_%d",
				dateDayNumber)
		} else {

			dateDayStr = fmt.Sprintf("%d",
				dateDayNumber)
		}

		tokenMap["!DateDayLeadUnderScore!"] = dateDayStr

		resultStr = strings.Replace(resultStr,
			"_2",
			"!DateDayLeadUnderScore!",
			1)

	} else if strings.Contains(resultStr, "2") {

		dateDayStr := fmt.Sprintf("%02d",
			dateDayNumber)

		tokenMap["!DateDayOneDigit!"] = dateDayStr

		resultStr = strings.Replace(resultStr,
			"2",
			"!DateDayOneDigit!",
			1)

	}

	return resultStr, err
}

// processDayOfWeek - processes and returns correct day of week format
func (calDtMech *calendarDateTimeMechanics) processDayOfWeek(
	inputStr string,
	dayOfWeekNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	daysOfWeek := map[int]string{
		0: "Sunday",
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
	}

	ePrefix += "calendarDateTimeMechanics.processDayOfWeek() "

	resultStr = inputStr

	if dayOfWeekNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'dayOfWeekNumber' is LESS THAN ZERO!\n" +
			"dayOfWeekNumber='%v'\n", dayOfWeekNumber)

		return resultStr, err
	}

	// Process Day Of Week
	if strings.Contains(resultStr, "Monday") {
		dayOfWeek := daysOfWeek[dayOfWeekNumber]
		tokenMap["!DayOfWeek!"] = dayOfWeek
		resultStr = strings.Replace(resultStr,
			"Monday",
			"!DayOfWeek!",
			1)

	} else if strings.Contains(resultStr, "Mon") {
		dayOfWeek := daysOfWeek[dayOfWeekNumber][0:3]
		tokenMap["!DayOfWeek!"] = dayOfWeek
		resultStr = strings.Replace(resultStr,
			"Mon",
			"!DayOfWeek!",
			1)
	}

	return resultStr, err
}

// processHours - processes and returns correct hour format
//
func (calDtMech *calendarDateTimeMechanics) processHours(
	inputStr string,
	hourNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "calendarDateTimeMechanics.processHours() "

	if hourNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hourNumber' is LESS THAN ONE!\n" +
			"hourNumber='%v'\n", hourNumber)

		return resultStr, err
	}

	if hourNumber > 23 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hourNumber' is GREATER THAN 23!\n" +
			"hourNumber='%v'\n", hourNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, "15") {

		hourStr := fmt.Sprintf("%02d", hourNumber)

		tokenMap["!HourTwentyFourTwoDigit!"] = hourStr

		resultStr = strings.Replace(resultStr,
			"15",
			"!HourTwentyFourTwoDigit!",
			1)

	} else if strings.Contains(resultStr, "03") {

		if hourNumber > 12 {
			hourNumber -= 12
		}

		hourStr := fmt.Sprintf("%02d", hourNumber)

		tokenMap["!HourTwelveTwoDigit!"] = hourStr

		resultStr = strings.Replace(resultStr,
			"03",
			"!HourTwelveTwoDigit!",
			1)

	} else if strings.Contains(resultStr, "3") {

		if hourNumber > 12 {
			hourNumber -= 12
		}

		hourStr := fmt.Sprintf("%02d", hourNumber)

		tokenMap["!HourTwelveOneDigit!"] = hourStr

		resultStr = strings.Replace(resultStr,
			"3",
			"!HourTwelveOneDigit!",
			1)
	}

	return resultStr, err
}


// processMicroseconds - processes and returns correct microseconds format
// Make certain this method is called after 'processNanoseconds()'.
//
func (calDtMech *calendarDateTimeMechanics) processMicroseconds(
	inputStr string,
	nanosecondsNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "calendarDateTimeMechanics.processMicroseconds() "

	if nanosecondsNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is LESS THAN ZERO!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	if nanosecondsNumber > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is GREATER THAN 999,999,999!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	microsecondsNumber := nanosecondsNumber / 1000

	if strings.Contains(resultStr, ".000000") {

		microsecondStr := fmt.Sprintf("%06d", microsecondsNumber)

		tokenMap["!MicrosecondsTrailingZeros!"] = microsecondStr

		resultStr = strings.Replace(resultStr,
			".000000",
			".!MicrosecondsTrailingZeros!",
			1)

	} else if strings.Contains(resultStr, ".999999") {

		microsecondStr := fmt.Sprintf("%d", microsecondsNumber)

		tokenMap["!MicrosecondsNoTrailingZeros!"] = microsecondStr

		resultStr = strings.Replace(resultStr,
			".999999",
			".!MicrosecondsNoTrailingZeros!",
			1)

	}

	return resultStr, err
}

// processMilliseconds - processes and returns correct milliseconds format
// Make certain this method is called after 'processNanoseconds()' and
// 'processMicroseconds()'.
//
func (calDtMech *calendarDateTimeMechanics) processMilliseconds(
	inputStr string,
	nanosecondsNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "calendarDateTimeMechanics.processMilliseconds() "

	if nanosecondsNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is LESS THAN ZERO!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	if nanosecondsNumber > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is GREATER THAN 999,999,999!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	millisecondsNumber := nanosecondsNumber / 1000000

	if strings.Contains(resultStr, ".000") {

		millisecondStr := fmt.Sprintf("%03d", millisecondsNumber)

		tokenMap["!MillisecondsTrailingZeros!"] = millisecondStr

		resultStr = strings.Replace(resultStr,
			".000",
			".!MillisecondsTrailingZeros!",
			1)

	} else if strings.Contains(resultStr, ".999") {

		millisecondStr := fmt.Sprintf("%d", millisecondsNumber)

		tokenMap["!MillisecondsNoTrailingZeros!"] = millisecondStr

		resultStr = strings.Replace(resultStr,
			".999",
			".!MillisecondsNoTrailingZeros!",
			1)

	}

	return resultStr, err
}

// processMinutes - processes and returns correct minute format
//
func (calDtMech *calendarDateTimeMechanics) processMinutes(
	inputStr string,
	minuteNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "calendarDateTimeMechanics.processMinutes() "

	if minuteNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minuteNumber' is LESS THAN ONE!\n" +
			"minuteNumber='%v'\n", minuteNumber)

		return resultStr, err
	}

	if minuteNumber > 59 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minuteNumber' is GREATER THAN 59!\n" +
			"minuteNumber='%v'\n", minuteNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, "04") {

		minuteStr := fmt.Sprintf("%02d", minuteNumber)

		tokenMap["!MinutesTwoDigit!"] = minuteStr

		resultStr = strings.Replace(resultStr,
			"04",
			"!MinutesTwoDigit!",
			1)

	} else if strings.Contains(resultStr, "4") {

		minuteStr := fmt.Sprintf("%d", minuteNumber)

		tokenMap["!MinutesOneDigit!"] = minuteStr

		resultStr = strings.Replace(resultStr,
			"4",
			"!MinutesOneDigit!",
			1)

	}

	return resultStr, err
}

// processMonths - processes and returns correct month
// formatting.
func (calDtMech *calendarDateTimeMechanics) processMonths(
	inputStr string,
	monthNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "calendarDateTimeMechanics.processAmPm() "

	monthsOfYear := map[int] string {
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "May",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December",
	}

	monthStr := monthsOfYear[monthNumber]

	if monthStr == "" {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'monthNumber' is INVALID!\n" +
			"monthNumber='%v'\n", monthNumber)
		return resultStr, err
	}

	// Process Months
	if strings.Contains(inputStr, "January") {

		tokenMap["!MonthName!"] = monthStr

		resultStr = strings.Replace(resultStr,
			"January",
			"!MonthName!",
			1)

	} else if strings.Contains(inputStr, "Jan") {

		monthStr = monthStr[0:3]

		tokenMap["!MonthNameAbbrv!"] = monthStr

		resultStr = strings.Replace(resultStr,
			"January",
			"!MonthNameAbbrv!",
			1)

	} else if strings.Contains(inputStr, "01") {

		monthStr = fmt.Sprintf("%02d", monthNumber)

		tokenMap["!MonthNumberTwoDigit!"] = monthStr

		resultStr = strings.Replace(resultStr,
			"01",
			"!MonthNumberTwoDigit!",
			1)

	} else if strings.Contains(inputStr, "1") {

		lenInputStr := len(inputStr)
		lastStrIdx := lenInputStr - 1

		for i:=0; i < lenInputStr; i++ {

			if inputStr[i] == '1' {

				if i == lastStrIdx {

					monthStr = fmt.Sprintf("%d", monthNumber)
					tokenMap["!MonthNumberOneDigit!"] = monthStr

					resultStr = resultStr[0:i] + "!MonthNumberOneDigit!"

					break
				} else {

					if inputStr[i+1] == '5' {
						continue

					} else {

						monthStr = fmt.Sprintf("%d", monthNumber)

						tokenMap["!MonthNumberOneDigit!"] = monthStr

						resultStr = resultStr[0:i+1] +
							"!MonthNumberOneDigit!" +
							resultStr[i+2:]
						break
					}
				}
			}
		}
	}

	return resultStr, err
}

// processNanoseconds - processes and returns correct nanoseconds format
// Make certain to call this method before calling 'processMicroseconds()'
// and 'processMilliseconds()'.
//
func (calDtMech *calendarDateTimeMechanics) processNanoseconds(
	inputStr string,
	nanosecondsNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	resultStr = inputStr

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	ePrefix += "calendarDateTimeMechanics.processAmPm() "

	if nanosecondsNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is LESS THAN ZERO!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	if nanosecondsNumber > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is GREATER THAN 999,999,999!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, ".000000000") {

		nanosecondStr := fmt.Sprintf("%09d", nanosecondsNumber)

		tokenMap["!NanosecondsTrailingZeros!"] = nanosecondStr

		resultStr = strings.Replace(resultStr,
			".000000000",
			".!NanosecondsTrailingZeros!",
			1)

	} else if strings.Contains(resultStr, ".999999999") {

		nanosecondStr := fmt.Sprintf("%d", nanosecondsNumber)

		tokenMap["!NanosecondsNoTrailingZeros!"] = nanosecondStr

		resultStr = strings.Replace(resultStr,
			".999999999",
			".!NanosecondsNoTrailingZeros!",
			1)

	}

	return resultStr, err
}

// processOffset - processes and returns offset hours and
// offset minutes format. If the 'offSetHours' value or the
// 'offSetMinutes' value is invalid, the offset will be deleted
// from the formatted date/time string.
//
func (calDtMech *calendarDateTimeMechanics) processOffset(
	inputStr string,
	offSetHours int,
	offSetMinutes int,
	tokenMap map[string]string) (resultStr string) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr

	isValid := true

	if offSetHours < -23 || offSetHours > 23 {
		isValid = false
	}

	if offSetMinutes < -59 || offSetMinutes > 59 {
		isValid = false
	}

	var offsetFmtStr string
	var keys []string


	if !isValid {

		offsetFmtStr = ""

		keys = []string {
			" Z0700",
			"Z0700",
			" Z07:00",
			"Z07:00",
			" -0700",
			"-0700",
			" -07:00",
			"-07:00",
			" -07",
			"-07",
		}

		lenKeys := len(keys)

		for i:=0; i < lenKeys; i++ {

			if strings.Contains(resultStr, keys[i]){

				tokenMap["!OffsetUTC!"] = offsetFmtStr

				resultStr = strings.Replace(resultStr,
					keys[i],
					"!OffsetUTC!",
					1)
				break
			}

		}

		return resultStr
	}

	keys = []string {
		"Z0700",
		"Z07:00",
		"-0700",
		"-07:00",
		"-07",
	}

	numberSign := 1

	if offSetHours < 0 {
		numberSign = -1
		offSetHours *= -1
	}

	if offSetMinutes < 0 {
		offSetMinutes *= -1
	}

	zPrefix := "Z"

	leadPrefix := ""

	if numberSign == -1 {
		zPrefix = "Z-"
		leadPrefix = "-"
	}

	values := []string {
		fmt.Sprintf(zPrefix + "%02d%02d",
			offSetHours,
			offSetMinutes),
		fmt.Sprintf(zPrefix + "%02d:%02d",
			offSetHours,
			offSetMinutes),
		fmt.Sprintf(leadPrefix + "%02d%02d",
			offSetHours,
			offSetMinutes),
		fmt.Sprintf(leadPrefix + "%02d:%02d",
			offSetHours,
			offSetMinutes),
		fmt.Sprintf(leadPrefix + "%02d",
			offSetHours),
	}

	lenKeys := len(keys)

	for i:=0; i < lenKeys; i++ {

		if strings.Contains(resultStr, keys[i]){

			tokenMap["!OffsetUTC!"] = values[i]

			resultStr = strings.Replace(resultStr,
				keys[i],
				"!OffsetUTC!",
				1)
			break
		}
	}

	return resultStr
}

// processSeconds - processes and returns correct second format
//
func (calDtMech *calendarDateTimeMechanics) processSeconds(
	inputStr string,
	secondNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr
	err = nil

	ePrefix += "calendarDateTimeMechanics.processSeconds() "

	if secondNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'secondNumber' is LESS THAN ONE!\n" +
			"secondNumber='%v'\n", secondNumber)

		return resultStr, err
	}

	if secondNumber > 59 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'secondNumber' is GREATER THAN 59!\n" +
			"secondNumber='%v'\n", secondNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, "05") {

		secondStr := fmt.Sprintf("%02d", secondNumber)

		tokenMap["!SecondsTwoDigit!"] = secondStr

		resultStr = strings.Replace(resultStr,
			"05",
			"!SecondsTwoDigit!",
			1)

	} else if strings.Contains(resultStr, "5") {

		secondStr := fmt.Sprintf("%d", secondNumber)

		tokenMap["!SecondsOneDigit!"] = secondStr

		resultStr = strings.Replace(resultStr,
			"5",
			"!SecondsOneDigit!",
			1)

	}

	return resultStr, err
}

// processYears - Returns the formatted year in target date time string
//
func (calDtMech *calendarDateTimeMechanics) processYears(
	inputStr string,
	year int64,
	tokenMap map[string]string) (resultStr string) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr

	// Process Years
	if strings.Contains(inputStr, "2006") {

		tokenMap["!YearFourDigit!"] =
			fmt.Sprintf("%d", year)

		resultStr = strings.Replace(resultStr,
			"2006",
			"!YearFourDigit!",
			1)

	} else if strings.Contains(inputStr, "06") {

		yearStr := fmt.Sprintf("%02d",
			year % 100)

		tokenMap["!YearTwoDigit!"] =
			yearStr

		resultStr = strings.Replace(resultStr,
			"06",
			"!YearTwoDigit!",
			1)

	} else if strings.Contains(inputStr, "6") {

		yearStr := fmt.Sprintf("%d",
			year % 100)

		tokenMap["!YearOneDigit!"] =
			yearStr

		resultStr = strings.Replace(resultStr,
			"6",
			"!YearOneDigit!",
			1)
	}

	return resultStr
}

// processTimeZone - Returns the formatted time zone in
// the target date time string.
//
func (calDtMech *calendarDateTimeMechanics) processTimeZone(
	inputStr string,
	timeZoneAbbrv string,
	tokenMap map[string]string) (resultStr string) {

	calDtMech.lock.Lock()

	defer calDtMech.lock.Unlock()

	resultStr = inputStr

	// Process Time Zones
	if strings.Contains(inputStr, "MST") {

		tokenMap["!TimeZone!"] = timeZoneAbbrv

		resultStr = strings.Replace(resultStr,
			"MST",
			"!TimeZone!",
			1)

	}

	return resultStr
}
