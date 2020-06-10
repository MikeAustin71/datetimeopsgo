package main

import (
	"fmt"
	dt "github.com/MikeAustin71/datetimeopsgo/datetime"
	ex "github.com/MikeAustin71/datetimeopsgo/datetimeexamples"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func main() {

	mainTest{}.mainTest092()

}

type mainTest struct {
	input  string
	output string
}

func (mt mainTest) mainTest093() {
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html

	dateTimeFormat := dt.FmtDateTimeYrMDayFmtStr

	lineLen := 70

	ePrefix := "mainTest.mainTest092() "

	titles := []string{ePrefix,
		"Convert to Julian Day Numbers",
		"from Gregorian Calendar Date Time",
		ePrefix}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	lineSplitter := strings.Repeat("-", lineLen)

	testDate := time.Date(
		2013,
		1,
		1,
		0,
		30,
		0,
		0,
		time.UTC)


	var julianDayNoTimeFloat64 float64

	julianDayNoTimeFloat64 = 2456293.520833

	julianDayNoTimeFloat64Str := "2456293.520833"

	var err error

	fmt.Println(lineSplitter)
	fmt.Println("Converting Julian Date Number Time To Fractional Time")
	fmt.Println(lineSplitter)

	fmt.Printf("Test Date:                %v\n",
		testDate.Format(dateTimeFormat))

	fmt.Println(lineSplitter)

	decimalCnt1 := "         1         2         3         4         5         6         7"
	decimalCnt2 := "1230464890123456789012345678901234567890123456789012345678901234567890"

	spacer := strings.Repeat(" ",46)

	fmt.Println(spacer + decimalCnt1)
	fmt.Println(spacer + decimalCnt2)

	fmt.Printf("julianDayNoTimeFloat64 Number:  %.6f\n",
		julianDayNoTimeFloat64)

	fmt.Printf("julianDayNoTimeFloat64 String:  %v\n",
		julianDayNoTimeFloat64Str)


	var tDur dt.TimeDurationDto

	tDur, err = ex.CodeTimer()

	if err != nil {
		fmt.Printf("Error returned by ex.CodeTimer()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	julianDayNoTimeDto,
	err :=
		dt.JulianDayNoDto{}.NewFromFloat64(julianDayNoTimeFloat64)

	if err != nil {
		fmt.Printf("Error returned by dt.JulianDayNoDto{}.NewFromFloat64(julianDayNoTimeFloat64)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tDur.SetAutoEnd()

	if err != nil {
		fmt.Printf("Error returned by tDur.SetAutoEnd()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var bigJDNanoseconds *big.Float

	bigJDNanoseconds, err = julianDayNoTimeDto.GetDayNoTimeBigFloat()

	if err != nil {
		fmt.Printf("Error returned by julianDayNoTimeDto.GetDayNoTimeBigFloat()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("Julian Day Number Time NanSecs:     %80.70f\n",
		bigJDNanoseconds)

	julianDayNoTimeStr,
	strWidth,
	intWidth :=
		julianDayNoTimeDto.GetJulianDayNoTimeStr(20)

	if strWidth  < 0 {
		fmt.Printf("Error: julianDayNoTimeDto.GetJulianDayNoTimeStr(8) Failed!\n")
		return
	}

	spacer2 := strings.Repeat(" ",14 - intWidth)

	fmt.Printf("Julian Day Number Time NanSecs:" + spacer2 + "%v\n",
		julianDayNoTimeStr)

	var timeFraction *big.Float

	timeFraction, err = julianDayNoTimeDto.GetJulianTimeFraction()

	if err != nil {
		fmt.Printf("Error returned by julianDayNoTimeDto.GetJulianTimeFraction()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDayNoTimeSeconds float64

	julianDayNoTimeSeconds, err = julianDayNoTimeDto.GetDayNoTimeFloat64()

	if err != nil {
		fmt.Printf("Error returned by julianDayNoTimeDto.GetDayNoTimeFloat64()\n" +
			"Error='%v'\n", err.Error())
	}

	fmt.Println(lineSplitter)

	fmt.Printf("Time Fraction:                      %80.70f\n",
		timeFraction)

	fmt.Println(lineSplitter)

	fmt.Printf("Julian Day No Time Seconds:         %80.70f\n",
		julianDayNoTimeSeconds)

	timeMech := dt.TimeMechanics{}

	days,
	hours,
	minutes,
	seconds,
	_ := timeMech.ComputeFloat64TimeFracToGregorianSeconds(julianDayNoTimeSeconds)

	fmt.Printf("Days: %v  Hours: %v  Minutes: %v  Seconds %v\n",
		days, hours, minutes, seconds)

	fmt.Println(lineSplitter)

	expectedHours := testDate.Hour()

	if expectedHours >= 12 {
		expectedHours -= 12
	}

	fmt.Printf("Expected Hours:       %v\n",
		expectedHours)
	fmt.Printf("Actual Hours:         %v\n",
		julianDayNoTimeDto.GetGregorianHours())
	fmt.Println(lineSplitter)

	fmt.Printf("Expected Minutes:     %v\n",
		testDate.Minute())
	fmt.Printf("Actual Minutes:       %v\n",
		julianDayNoTimeDto.GetMinutes())
	fmt.Println(lineSplitter)

	fmt.Printf("Expected Seconds:     %v\n",
		testDate.Second())

	fmt.Printf("Actual Seconds:       %v\n",
		julianDayNoTimeDto.GetSeconds())
	fmt.Println(lineSplitter)

	fmt.Printf("Expected Nanoseconds: %v\n",
		testDate.Nanosecond())

	fmt.Printf("Actual Nanoseconds:   %v\n",
		julianDayNoTimeDto.GetNanoseconds())

	fmt.Println(lineSplitter)
	fmt.Printf("Code Execution Time:   %v\n",
		tDur.GetElapsedTimeStr())

	fmt.Println(lineSplitter)
	fmt.Println(lineSplitter)

}

func (mt mainTest) mainTest092() {
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html

	dateTimeFormat := dt.FmtDateTimeYrMDayFmtStr

	lineLen := 70

	ePrefix := "mainTest.mainTest092() "

	titles := []string{"Convert to Julian Day Numbers",
		"from Julian Calendar Date Time",
		ePrefix}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	lineSplitter := strings.Repeat("-", lineLen)
// October 15, 1582 13:00:00
	testDate := time.Date(
		1582,
		10,
		15,
		13,
		0,
		0,
		500000000,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1582,
			10,
			5,
			13,
			0,
			0,
			500000000,
			time.UTC)

	var err error

	fmt.Println(lineSplitter)
	fmt.Println("Converting Julian Day Number Time To Julian Calendar Date Time")
	fmt.Println(lineSplitter)

	fmt.Println(lineSplitter)

	decimalCnt1 := "         1         2         3         4         5         6         7"
	decimalCnt2 := "1230464890123456789012345678901234567890123456789012345678901234567890"

	spacer := strings.Repeat(" ",46)

	fmt.Println(spacer + decimalCnt1)
	fmt.Println(spacer + decimalCnt2)

	calUtil := dt.CalendarUtility{}

	var tDur dt.TimeDurationDto

	tDur, err = ex.CodeTimer()

	if err != nil {
		fmt.Printf("Error returned by ex.CodeTimer()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDayNoTimeDto dt.JulianDayNoDto

	_,
	julianDayNoTimeDto,
	err =
		calUtil.GregorianDateToJulianDayNoTime(
			testDate,
			ePrefix)

	if err != nil {
		fmt.Printf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	julianDateTime,
	err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		fmt.Printf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tDur.SetAutoEnd()

	if err != nil {
		fmt.Printf("Error returned by tDur.SetAutoEnd()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var bigJDNanoseconds *big.Float

	bigJDNanoseconds, err = julianDayNoTimeDto.GetDayNoTimeBigFloat()

	if err != nil {
		fmt.Printf("Error returned by julianDayNoTimeDto.GetDayNoTimeBigFloat()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("Julian Day Number Time NanSecs:     %80.70f\n",
		bigJDNanoseconds)

	julianDayNoTimeStr,
	strWidth,
	intWidth :=
		julianDayNoTimeDto.GetJulianDayNoTimeStr(20)

	if strWidth  < 0 {
		fmt.Printf("Error: julianDayNoTimeDto.GetJulianDayNoTimeStr(8) Failed!\n")
		return
	}

	spacer2 := strings.Repeat(" ",14 - intWidth)

	fmt.Printf("Julian Day Number Time NanSecs:" + spacer2 + "%v\n",
		julianDayNoTimeStr)

	var timeFraction *big.Float

	timeFraction, err = julianDayNoTimeDto.GetJulianTimeFraction()

	if err != nil {
		fmt.Printf("Error returned by julianDayNoTimeDto.GetJulianTimeFraction()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Println(lineSplitter)

	fmt.Printf("Time Fraction:                      %80.70f\n",
		timeFraction)

	fmt.Println(lineSplitter)

	fmt.Printf("Gregorian Calendar Test Date:      %v\n",
		testDate.Format(dateTimeFormat))

	fmt.Printf("Computed Julian Date Time:         %v\n",
		julianDateTime.Format(dateTimeFormat))

	fmt.Printf("Expected Juilain Date Time:        %v\n",
		expectedJulianDateTime.Format(dateTimeFormat))

	fmt.Println(lineSplitter)
	fmt.Printf("Code Execution Time:   %v\n",
		tDur.GetElapsedTimeStr())

	fmt.Println(lineSplitter)
	fmt.Println(lineSplitter)
}

func (mt mainTest) mainTest091() {
// https://keisan.casio.com/exec/system/1227779487#!
// https://www.aavso.org/jd-calculator
// http://numerical.recipes/julian.html
// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
// https://en.wikipedia.org/wiki/Julian_day

	// Julian Day Formula
	// https://www.howcast.com/videos/how-to-calculate-julian-dates
	// https://sciencing.com/calculate-age-lunar-years-5997325.html
	// https://stason.org/TULARC/society/calendars/2-15-1-Is-there-a-formula-for-calculating-the-Julian-day-nu.html
	// http://mathforum.org/library/drmath/view/51907.html


	dateTimeFormat := dt.FmtDateTimeYrMDayFmtStr

	lineLen := 70

	ePrefix := "mainTest.mainTest091() "

	titles := []string{"Convert Julian Day Numbers",
		"to Gregorian Calendar Date Times",
		ePrefix}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	lineSplitter := strings.Repeat("-", lineLen)

	//  1582  10     4.0  2299159.5         2299149.5
	expectedDateTime := time.Date(
		1582,
		10,
		4,
		0,
		0,
		0,
		0,
		time.UTC)


	var julianDateTimeInputValue float64
	var err error

	julianDateTimeInputStr := "2299149.500000"

	idxDot := strings.Index(julianDateTimeInputStr, ".")

	strSpacer := strings.Repeat(" ", 17 - idxDot)

	julianDateTimeInputValue,
	err = strconv.ParseFloat(julianDateTimeInputStr, 64)

	if err != nil {
		fmt.Printf("\n" +
			"Error returned by strconv.ParseFloat(julianDateTimeInputStr, 64).\n" +
			"julianDateTimeInputStr='%v'\n" +
			"Error='%v'\n",
			julianDateTimeInputStr, err.Error())
		return
	}


	var julianDayNoDtoDate, julianDayNoDtoFloat64 dt.JulianDayNoDto

	_, julianDayNoDtoDate, err = dt.JulianDayNoDto{}.NewFromGregorianDate(expectedDateTime)

	if err != nil {
		fmt.Printf("Error returned by JulianDayNoDto{}.NewFromGregorianDate().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	julianDayNoDtoFloat64, err = dt.JulianDayNoDto{}.NewFromFloat64(
		julianDateTimeInputValue)

	if err != nil {
		fmt.Printf("Error returned by JulianDayNoDto{}.NewFromFloat64().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var calcGregorianDateUtc time.Time

	calcGregorianDateUtc,
	julianDayNoDtoDate,
	err = dt.JulianDayNoDto{}.NewFromGregorianDate(expectedDateTime)

	if err != nil {
		fmt.Printf("\nError returned by JulianDayNoDto{}.NewFromGregorianDate(expectedDateTime).\n" +
			"expectedDateTime='%v'\n" +
			"Error='%v'\n",
			expectedDateTime.Format(dt.FmtDateTimeDMYNanoTz),
			err.Error())
		return
	}

	if ! calcGregorianDateUtc.Equal(expectedDateTime) {
		fmt.Printf("\nError:\n" +
			"Expected calcGregorianDateUtc= '%v'\n" +
			"Instead, calcGregorianDateUtc= '%v'\n",
			expectedDateTime.Format(dt.FmtDateTimeDMYNanoTz),
			calcGregorianDateUtc.Format(dt.FmtDateTimeDMYNanoTz))
		return
	}

	calUtil := dt.CalendarUtility{}
	var gregorianDateTimeFloat64,
			gregorianDateTimeBig time.Time

	gregorianDateTimeFloat64,
	err = calUtil.JulianDayNoTimeToGregorianCalendar(
		julianDayNoDtoFloat64,
		ePrefix)

	if err != nil {
		fmt.Printf("Conversion1:\n" +
			"%v\n", err.Error())
		return
	}

	gregorianDateTimeBig,
	err = calUtil.JulianDayNoTimeToGregorianCalendar(
		julianDayNoDtoDate,
		ePrefix)

	if err != nil {
		fmt.Printf("Conversion2:\n%v\n", err.Error())
		return
	}

	var calculatedJulianDayNoTimeFloat64,
	calculatedJulianDayNoTimeBigFloat *big.Float

	calculatedJulianDayNoTimeFloat64,
	err = julianDayNoDtoFloat64.GetDayNoTimeBigFloat()

	if err != nil {
		fmt.Printf("Error returned by julianDayNoDtoFloat64.GetDayNoTimeBigFloat().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	calculatedJulianDayNoTimeBigFloat,
	err = julianDayNoDtoDate.GetDayNoTimeBigFloat()

	fmt.Println(lineSplitter)
	fmt.Println("Converting Julian Date Number Time To Gregorian Date Time")
	fmt.Println(lineSplitter)
	fmt.Println(lineSplitter)
	decimalCnt1 := "         1         2         3         4         5         6         7"
	decimalCnt2 := "1230464890123456789012345678901234567890123456789012345678901234567890"

	decFmtStr := "%40.30f\n"

	spacer := strings.Repeat(" ",40)

	fmt.Println(spacer + decimalCnt1)
	fmt.Println(spacer + decimalCnt2)

	fmt.Printf("Input JDN/Time String:" + strSpacer + "%v\n",
		julianDateTimeInputStr)
	fmt.Printf("Input JDN/Time Value:         " + decFmtStr,
		julianDateTimeInputValue)

	fmt.Printf("Calculated JDN/Time float64:  " + decFmtStr,
		calculatedJulianDayNoTimeFloat64)

	fmt.Printf("Calculated JDN/Time MaxPrec:  " + decFmtStr,
			calculatedJulianDayNoTimeBigFloat)
	fmt.Println(lineSplitter)

	fmt.Printf("Test Date:                               %v\n",
		expectedDateTime.Format(dateTimeFormat))

	fmt.Printf("Calculated Gregorian Date Float64:       %v\n",
		gregorianDateTimeFloat64.Format(dateTimeFormat))

	fmt.Printf("Calculated Gregorian Date Big:           %v\n",
		gregorianDateTimeBig.Format(dateTimeFormat))

	fmt.Printf("Expected Gregorian Date:                 %v\n",
		expectedDateTime.Format(dateTimeFormat))
	fmt.Println(lineSplitter)
	fmt.Println(lineSplitter)

}

func (mt mainTest) mainTest090() {

	dateTimeFormat := dt.FmtDateTimeYrMDayFmtStr

	lineLen := 70

	titles := []string{"mainTest.mainTest090()",
		"Testbed for Computing Julian Day Numbers"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	lineSplitter := strings.Repeat("-", lineLen)

/*
	testDate := time.Date(
		2013,
		1,
		1,
		0,
		30,
		0,
		0,
		time.UTC)
	// Expected 2 456 293.520 833

	// Example: The Julian Date for 00:30:00.0 UT January 1, 2013, is 2 456 293.520 833

	expectedJulianDate := float64(2456293.520833)
*/

	testDate := time.Date(
		2013,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC)

	var expectedJulianDate float64

	expectedJulianDate = 2456294.000000

	calUtil := dt.CalendarUtility{}

	gregorianDateUtc,
	julianDayNoDto,
	err := calUtil.GregorianDateToJulianDayNoTime(
		testDate,
		"mainTest.mainTest090() ")

	if err != nil {
		fmt.Printf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	fmt.Println(lineSplitter)
	fmt.Println("Date Time Conversion to Julian Date")
	fmt.Println(lineSplitter)
	fmt.Printf("Test Date Time:          %v\n",
		testDate.Format(dateTimeFormat))
	fmt.Println(lineSplitter)
	fmt.Printf("gregorianDateUtc:        %v\n",
		gregorianDateUtc.Format(dateTimeFormat))

	var julianDayNoBigFloat *big.Float

	julianDayNoBigFloat, err = julianDayNoDto.GetDayNoTimeBigFloat()

	if err != nil {
		fmt.Printf("Error returned by julianDayNoDto.GetDayNoTimeBigFloat().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("Calculated Julian Date:  %40.30f\n",
		julianDayNoBigFloat)
	fmt.Printf("Expected Julian Date:    %20.8f\n",
		expectedJulianDate)

	fmt.Println(lineSplitter)
	fmt.Println(lineSplitter)



}

func (mt mainTest) mainTest089() {

	dateTimeFormat := dt.FmtDateTimeYrMDayFmtStr

	lineLen := 70

	titles := []string{"mainTest.mainTest089()",
		"Testbed for Computing Julian Day Numbers"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	lineSplitter := strings.Repeat("-", lineLen)


	testDate := time.Date(
		2013,
		1,
		1,
		0,
		30,
		0,
		0,
		time.UTC)
// Expected 2 456 293.520 833

// Example: The Julian Date for 00:30:00.0 UT January 1, 2013, is 2 456 293.520 833

	expectedJulianDayNo := int64(2456293)


/*
	testDate := time.Date(
		2020,
		5,
		19,
		12,
		0,
		0,
		0,
		time.UTC)

	// Expected 2458989.2829861 Now adjusted for beginning of Julian day at noon

	expectedJulianDayNo := int64(2458989)
*/

// http://www.csgnetwork.com/juliandaydate.html

	calUtil := dt.CalendarUtility{}

	gregorianDateUtc,
	julianDayNumber,
	err := calUtil.GregorianDateToJulianDayNo(
		testDate,
		"mainTest.mainTest089() ")

	if err != nil {
		fmt.Printf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	fmt.Println(lineSplitter)
	fmt.Println("Date Time Conversion to Julian Day No.")
	fmt.Println(lineSplitter)
	fmt.Printf("Test Date Time:          %v\n",
		testDate.Format(dateTimeFormat))
	fmt.Println(lineSplitter)
	fmt.Printf("gregorianDateUtc:        %v\n",
		gregorianDateUtc.Format(dateTimeFormat))
	fmt.Printf("Calculated Julian Day Number: %v\n",
		julianDayNumber)
	fmt.Printf("Expected Julian Day Number:   %v\n",
		expectedJulianDayNo)

	fmt.Println(lineSplitter)
	fmt.Println(lineSplitter)
}

func (mt mainTest) mainTest088() {

	// "0001-01-01 00:00:00.000000000 +0000 UTC"
	zeroTime := time.Time{}
	dateTimeFormat := dt.FmtDateTimeYrMDayFmtStr
		lineLen := 70

	titles := []string{"mainTest.mainTest088()",
		"Testing Starting and Ending Date",
		"Time Duration Calculation",
		"Starting Date Time Zero"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	lineSplitter := strings.Repeat("-", lineLen)
	fmt.Println(lineSplitter)
	fmt.Println("Zero and Zero Minus Times")
	fmt.Println(lineSplitter)
	fmt.Printf("Zero Time:               %v\n",
		zeroTime.Format(dateTimeFormat))

	t2ZeroMinus1Year := time.Date(
		zeroTime.Year() -1,
		1,
		1,
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("Zero Time -1Year         %v\n",
		t2ZeroMinus1Year.Format(dateTimeFormat))

	t500BCE := time.Date(
		zeroTime.Year() -500,
		1,
		1,
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("Zero Time -500Years      %v\n",
		t500BCE.Format(dateTimeFormat))
	fmt.Println(lineSplitter)

}


func (mt mainTest) mainTest087() {
	/*
	  "0001-01-01 00:00:00.000000000 +0000 UTC"
		zeroTime := time.Time{}
	*/
	lineLen := 70

	titles := []string{"mainTest.mainTest087()",
		"Testing Starting and Ending Date",
		"Time Duration Calculation",
		"Starting Date Time Zero"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	t1startDateTime := time.Time{}

	t2endDateTime := time.Date(
		1966,
		9,
		7,
		9,
		30,
		31,
		25647,
		time.UTC)

	longDur, err := dt.LongTimeDuration{}.NewStartEndDateTimes(
		t1startDateTime,
		t2endDateTime,
		dt.TZones.UTC(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by LongTimeDuration{}.NewStartEndDateTimes()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	lineSplitter := strings.Repeat("-", lineLen)

	fmt.Println(lineSplitter)
	fmt.Println("Expected Starting and Ending Date Times")
	fmt.Println(lineSplitter)
	fmt.Printf("Expected Start Date:     %v\n",
		t1startDateTime.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Printf("Expected End Date:       %v\n",
		t2endDateTime.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Printf("Expected Time Zone:      %v\n",
		t2endDateTime.Location().String())
	fmt.Println(lineSplitter)

	var finalStartDateTz, finalEndDateTz dt.DateTzDto

	finalStartDateTz,
	finalEndDateTz = longDur.GetStartEndDatesTz()

	finalStartDateTime := finalStartDateTz.GetDateTimeValue()

	fmt.Printf("Actual Start Date:       %v\n",
		finalStartDateTz.GetDateTimeText())
	fmt.Printf("Actual End Date:         %v\n",
		finalEndDateTz.GetDateTimeText())
	fmt.Printf("Actual Time Zone:        %v\n",
		finalStartDateTz.GetTimeZoneName())
	fmt.Printf("Final Start Time Zone:   %v\n",
		finalStartDateTime.Location().String())

}

func (mt mainTest) mainTest086() {
/*

 BaseTime Duration:
 5-Years 6-Months 12-Days 0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds

  Local Tz Times:
   Start Time: 2014-02-15 19:54:30.038175584 -0600 CST
     End Time: 2019-08-27 19:54:30.038175584 -0500 CDT

  UTC Tz Times:
   Start Time: 2014-02-15 19:54:30.038175584 -0600 CST
     End Time: 2019-08-27 20:54:30.038175584 -0500 CDT

*/
	lineLen := 70

	titles := []string{"mainTest.mainTest086()",
		"Testing Duration Since Time Zero"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	centralTz, err := time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned by LoadLocation(dt.TZones.America.Chicago())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1StartDate := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		centralTz)

	t2EndDate := time.Date(
		2019,
		8,
		27,
		19,
		54,
		30,
		38175584,
		centralTz)

	calcDuration := t2EndDate.Sub(t1StartDate)

	var lngDur dt.LongTimeDuration

	lngDur, err =
		dt.LongTimeDuration{}.NewFromDuration(
			t1StartDate,
			calcDuration)

	if err != nil {
		fmt.Printf("Error returne by LongTimeDuration.NewFromDuration()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var finalStartDate, finalEndDate time.Time

	finalStartDate,
	finalEndDate,
	err =	lngDur.GenerateDatePlusDuration(t1StartDate)

	if err != nil {
		fmt.Printf("Error returne by lngDur.GenerateDatePlusDuration(t1StartDate)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	lineSplitter := strings.Repeat("-", lineLen)

	fmt.Println(lineSplitter)
	fmt.Printf("       Base Data\n")
	fmt.Println(lineSplitter)
	fmt.Printf("Base Start Date:         %v\n",
		t1StartDate.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Printf("  Base End Date:         %v\n",
		t2EndDate.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Println(lineSplitter)
	fmt.Printf("  Long Duration Data\n")
	fmt.Println(lineSplitter)
	fmt.Printf("Calc Start Date:         %v\n",
		finalStartDate.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Printf("Calc End Date:           %v\n",
		finalEndDate.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println(lineSplitter)
	fmt.Println("Duration Comparison")
	fmt.Printf("Expected Duration:       %v\n",
		calcDuration)
	bigDur, _ := lngDur.GetLongDuration()

	reportedDur := time.Duration(bigDur.Int64())

	fmt.Printf("Reported Duration:       %v\n",
		reportedDur)
}

func (mt mainTest) mainTest085() {

	lineLen := 70

	titles := []string{"mainTest.mainTest084()",
		"Testing Duration Since Time Zero"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	t1 := time.Date(
		2014,
		3,
		5,
		9,
		35,
		21,
		53597,
		time.UTC)

// 0001-01-01 00:00:00.000000000 +0000 UTC
	zeroTime := time.Time{}

/*
	zeroTime := time.Date(
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		time.UTC)
*/


	fmt.Printf("Zero Date Time:          %v\n",
		zeroTime.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Printf("t1 Date Time:            %v\n",
		t1.Format(dt.FmtDateTimeYrMDayFmtStr))

	duration := t1.Sub(zeroTime)

	fmt.Printf("Duration:                %v\n",
		duration)

	t2 := zeroTime.Add(duration)

	fmt.Printf("t2 Date Time:            %v\n",
		t2.Format(dt.FmtDateTimeYrMDayFmtStr))

}

func (mt mainTest) mainTest084() {

	lineLen := 70

	titles := []string{"mainTest.mainTest084()",
		"Testing Change To Standard Time",
		"November 2, 2014",
		"Using TimeDto To Compute Duration"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	var err error
	var centralTz *time.Location

	centralTz, err = time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.America.Chicago())\n" +
			"Error='%v'\n", err.Error())
	}

	t1 := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		0,
		centralTz,
		)

	t1UTC := t1.In(time.UTC)

	t2UTC := t1.AddDate(3, 2, 15)

	nanoSecs := int64(0)

	nanoSecs += dt.HourNanoSeconds * 3
	nanoSecs += dt.MinuteNanoSeconds * 4
	nanoSecs += dt.SecondNanoseconds * 2

	t2UTC = t2UTC.Add(time.Duration(nanoSecs))

	calculatedUTCDuration := t2UTC.Sub(t1UTC)

	// 3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds
	// expected duration = "28082h4m2s"

	var tDto dt.TimeDto

	tDto, err = dt.TimeDto{}.NewFromDuration(
		calculatedUTCDuration)

	if err != nil {
		fmt.Printf("Error returned by TimeDto{}.NewFromDuration().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var actualDuration time.Duration

	actualDuration, err = tDto.GetDuration()

	if err != nil {
		fmt.Printf("Error returned by tDto.GetDuration().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	lineSplitter := strings.Repeat("-", lineLen)

	fmt.Printf("UTC Beginning and Ending Date Times")
	fmt.Println(lineSplitter)
	fmt.Printf("Start Time UTC:          %v\n", t1UTC)
	fmt.Printf("  End Time UTC:          %v\n", t1UTC)
	fmt.Println(lineSplitter)
	fmt.Println()
	fmt.Printf("TimeDto Duration Comparison")
	fmt.Println(lineSplitter)
	fmt.Printf("Expected Duration:       %v\n",
		calculatedUTCDuration)
	fmt.Printf(" TimeDto Duration:       %v\n",
		actualDuration)
	fmt.Println(lineSplitter)
	fmt.Println("TimeDto Contents")
	fmt.Println(lineSplitter)
	fmt.Printf("Years:                   %v\n",
		tDto.Years)
	fmt.Printf("Months:                  %v\n",
		tDto.Months)
	fmt.Printf("Weeks:                   %v\n",
		tDto.Weeks)
	fmt.Printf("Week Days:               %v\n",
		tDto.WeekDays)
	fmt.Printf("Date Days:               %v\n",
		tDto.DateDays)
	fmt.Printf("Hours:                   %v\n",
		tDto.Hours)
	fmt.Printf("Minutes:                 %v\n",
		tDto.Minutes)
	fmt.Printf("Seconds:                 %v\n",
		tDto.Seconds)
	fmt.Printf("Milliseconds:            %v\n",
		tDto.Milliseconds)
	fmt.Printf("Microseconds:            %v\n",
		tDto.Microseconds)
	fmt.Printf("Nanoseconds:             %v\n",
		tDto.Nanoseconds)
	fmt.Printf("TotSubSecNanoSeconds     %v\n",
		tDto.TotSubSecNanoseconds)
	fmt.Println(lineSplitter)
	fmt.Println(lineSplitter)
}


func (mt mainTest) mainTestBed02() {
	/* This a test vehicle for testing date addition.

			Daylight Savings Time Changed To Standard Time
		 on November 2, 2014
	  // FmtDateTimeYrMDayFmtStr - Year Month Day Date Format String
	  FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"

	*/

	lineLen := 70

	titles := []string{"mainTest.mainTestBed01()",
		"Change To Standard Time on November 2, 2014",
		"Test Bed For Subtracting Time Values"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	locationPtr, err := time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var endDateTimeUTCLocal, endDateTimeLocal,
			startDateTimeLocal, endDateTimeUTC,
			startDateTimeUTCLocal, startDateTimeUTC time.Time

	// Local Start Date Time
	// Base:      2014-11-02 00:00:00.000000000 -0500 CDT
	startDateTimeLocal = time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		locationPtr)

	startDateTimeUTC = startDateTimeLocal.In(time.UTC)
	startDateTimeUTCLocal = startDateTimeUTC.In(locationPtr)
/*
	hours := time.Duration(dt.HourNanoSeconds * -24)

	minutes := time.Duration(dt.MinuteNanoSeconds * -1)

	seconds := time.Duration(dt.SecondNanoseconds * -1)

	nanoSeconds := time.Duration(-1)

	duration := hours + minutes + seconds + nanoSeconds
*/

	endDateTimeLocal = startDateTimeLocal.AddDate(
		0,
		0,
		1)

	endDateTimeUTC = startDateTimeUTC.AddDate(
		0,
		0,
		1)


	endDateTimeUTCLocal = endDateTimeUTC.In(locationPtr)

	lineSplitter := strings.Repeat("-", lineLen)

	equalLineSplitter := strings.Repeat("=", lineLen)

	fmt.Println(equalLineSplitter)

	fmt.Printf(" Adding Date: " +
		 "1-Day\n")
	fmt.Printf("Base Date Local:             %v\n",
		startDateTimeLocal.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println(equalLineSplitter)

	fmt.Printf("Local Start Date:            %v\n",
		startDateTimeLocal.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Printf("Local End Date:              %v\n",
		endDateTimeLocal.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Println(equalLineSplitter)

	fmt.Println("UTC Conversion Data")
	fmt.Println(lineSplitter)

	fmt.Printf("Calculated UTC Start Date:   %v\n",
		startDateTimeUTC.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Printf("Initial UTC End Date:        %v\n",
		endDateTimeUTC.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Println(lineSplitter)

	fmt.Println("   Equivalent Local Date Times       ")
	fmt.Println("UTC Date Time Converted To Local Time")
	fmt.Println(lineSplitter)

	fmt.Printf("UTC -> Local Start Date:     %v\n",
		startDateTimeUTCLocal.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Printf("UTC -> End Date:             %v\n",
		endDateTimeUTCLocal.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Println(equalLineSplitter)

}

func (mt mainTest) mainTestBed01() {
	/* This a test vehicle for testing date subtraction.

			Daylight Savings Time Changed To Standard Time
		 on November 2, 2014
	  // FmtDateTimeYrMDayFmtStr - Year Month Day Date Format String
	  FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"

	*/

	lineLen := 70

	titles := []string{"mainTest.mainTestBed01()",
		"Change To Standard Time on November 2, 2014",
		"Test Bed For Subtracting Time Values"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	locationPtr, err := time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// Local End Date Time
	// Base:      2014-11-03 00:00:00.000000000 -0600 CST
	endDateTimeLocal := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		locationPtr)

	endDateTimeUTC := endDateTimeLocal.In(time.UTC)

	var endDateTimeUTCLocal, startDateTimeLocal,
	startDateTimeUTCLocal, startDateTimeUTC time.Time

	hours := time.Duration(dt.HourNanoSeconds * -24)

	minutes := time.Duration(dt.MinuteNanoSeconds * -1)

	seconds := time.Duration(dt.SecondNanoseconds * -1)

	var totalNano time.Duration

	totalNano = -1

	duration := hours + minutes + seconds + totalNano

	startDateTimeUTC = endDateTimeUTC.Add(duration)

	startDateTimeLocal = endDateTimeLocal.Add(duration)

	startDateTimeUTCLocal = startDateTimeUTC.In(locationPtr)

	endDateTimeUTCLocal = endDateTimeUTC.In(locationPtr)

	lineSplitter := strings.Repeat("-", lineLen)

	equalLineSplitter := strings.Repeat("=", lineLen)

	fmt.Println(equalLineSplitter)

	fmt.Printf(" Subtracting -24-hours, -1-minute, -1-second, 0-milliseconds," +
		 "0-microseconds, -1-nanoseconds\n")

	fmt.Println(equalLineSplitter)

	fmt.Printf("Calculated Local Start Date: %v\n",
		startDateTimeLocal.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Printf("Initial Local End Date:      %v\n",
		endDateTimeLocal.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Println(equalLineSplitter)
	fmt.Println("UTC Conversion Data")
	fmt.Println(lineSplitter)

	fmt.Printf("Calculated UTC Start Date:   %v\n",
		startDateTimeUTC.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Printf("Initial UTC End Date:        %v\n",
		endDateTimeUTC.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Println(lineSplitter)

	fmt.Println("   Equivalent Local Date Times       ")
	fmt.Println("UTC Date Time Converted To Local Time")
	fmt.Println(lineSplitter)

	fmt.Printf("UTC -> Local Start Date:     %v\n",
		startDateTimeUTCLocal.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Printf("UTC -> End Date:             %v\n",
		endDateTimeUTCLocal.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Println(equalLineSplitter)

}

func (mt mainTest) mainTest083() {
	/*
		Daylight Savings Time Changed To Standard Time
		 on November 2, 2014

		This will subtract 1-day from November 3, 2014
	                             Results - Output
	======================================================================
	                        mainTest.mainTest083()
	                    Testing Change To Standard Time
	                           November 2, 2014
	         Using TimeDurationDto and Time Math Calculation Modes
	                      LocalTimeZone & UtcTimeZone
	       24-hours Duration Added Over Daylight Savings Time Change
	======================================================================
	Initial Start Date:         2014-11-03 00:00:00.000000000 -0600 CST
	----------------------------------------------------------------------
	Expected Date Local:        2014-11-02 00:00:00.000000000 -0500 CDT
	Actual Date Local:          2014-11-02 00:00:00.000000000 -0500 CDT
	----------------------------------------------------------------------
	Expected Date UTC:          2014-11-02 01:00:00.000000000 -0500 CDT
	Actual Date UTC:            2014-11-02 01:00:00.000000000 -0500 CDT
	*/


	lineLen := 70

	titles := []string{"mainTest.mainTest083()",
		"Testing Change To Standard Time",
		"November 2, 2014",
		"Using TimeDurationDto and Time Math Calculation Modes",
		"LocalTimeZone & UtcTimeZone",
		"24-hours Duration Subtracted Over Daylight Savings Time Change"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	var err error
	var centralTz *time.Location

	centralTz, err = time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time." +
			"LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	initialStartDate := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		centralTz)

	initialStartDateUTC := initialStartDate.In(time.UTC)

	expectedEndDateLocalTz := initialStartDate.AddDate(
		0,
		0,
		-1)

	var negDuration time.Duration

	negDuration = time.Duration(int64(time.Hour) * int64(-24))

	expectedStartDateTimeUTC := initialStartDateUTC.Add(negDuration).In(centralTz)

	fmt.Printf("Initial Start Date:         %v\n",
		initialStartDate.Format(dt.FmtDateTimeYrMDayFmtStr))

	lineSplitter := strings.Repeat("-", lineLen)

	fmt.Println(lineSplitter)

	fmt.Printf("Expected Date Local:        %v\n",
		expectedEndDateLocalTz.Format(dt.FmtDateTimeYrMDayFmtStr))

	var tDurDto, tDurDto2 dt.TimeDurationDto

	tDurDto, err = dt.TimeDurationDto{}.NewStartTimeAddDate(
		initialStartDate,
		0,
		0,
		-1,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.America.Chicago(),
		dt.TCalcMode.LocalTimeZone(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by (Local) TimeDurationDto{}.NewStartTimeAddDate()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("Actual Date Local:          %v\n",
		tDurDto.GetThisStartDateTimeString())

	fmt.Println(lineSplitter)

	fmt.Printf("Expected Date UTC:          %v\n",
		expectedStartDateTimeUTC.Format(dt.FmtDateTimeYrMDayFmtStr))

	tDurDto2, err = dt.TimeDurationDto{}.NewStartTimeAddDate(
		initialStartDate,
		0,
		0,
		-1,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.America.Chicago(),
		dt.TCalcMode.UtcTimeZone(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by (Local) TimeDurationDto{}.NewStartTimeAddDate()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("Actual Date UTC:            %v\n",
		tDurDto2.GetThisStartDateTimeString())

}

func (mt mainTest) mainTest082() {


	/*
		Re-Test of mainTest073 using latest
		code updates.

		Variation on:
		TestDateTzDto_AddDate_01
		datetime\zztdatetzdto01_test.go

	 BaseTime Duration:
	 5-Years 6-Months 12-Days 0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds

	  Local Tz Times:
	   Start Time: 2014-02-15 19:54:30.038175584 -0600 CST
	     End Time: 2019-08-27 19:54:30.038175584 -0500 CDT

	  UTC Tz Times:
	   Start Time: 2014-02-15 19:54:30.038175584 -0600 CST
	     End Time: 2019-08-27 20:54:30.038175584 -0500 CDT
	 */

	lineLen := 70
	titles := []string{"mainTest.mainTest082()",
		"Testing TimeDurationDto",
		"Duration: \"5-Years, 6-Months 12-Days\""}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	locationPtr, err := time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-02-15 19:54:30.038175584 -0600 CST"
	startDateTime := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		locationPtr)

	// Local Tz End Date Time
	// 2019-08-27 19:54:30.038175584 -0500 CDT
	calcEndDateTimeLocal := time.Date(
		2019,
		8,
		27,
		19,
		54,
		30,
		38175584,
		locationPtr)

	// UTC Tz End Date Time
	// 2019-08-27 20:54:30.038175584 -0500 CDT
	calcEndDateTimeUTC := time.Date(
		2019,
		8,
		27,
		20,
		54,
		30,
		38175584,
		locationPtr)

	var startDateTimeTz, endDateTimeTz1, endDateTimeTz2 dt.DateTzDto

	startDateTimeTz, err = dt.DateTzDto{}.NewDateTime(
		startDateTime, fmtStr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewDateTime(startDateTime).\n" +
			"startDateTime='%v'" +
			"Error='%v'\n",
			startDateTime.Format(fmtStr), err.Error())
		return
	}

	endDateTimeTz1, err = startDateTimeTz.AddDate(
		dt.TCalcMode.LocalTimeZone(),
		5,
		6,
		12,
		fmtStr)

	if err != nil {
		fmt.Printf("Error returned by (endDateTimeTz1) " +
			"startDateTimeTz.AddDate().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	endDateTimeTz2, err = startDateTimeTz.AddDate(
		dt.TCalcMode.UtcTimeZone(),
		5,
		6,
		12,
		fmtStr)

	if err != nil {
		fmt.Printf("Error returned by (endDateTimeTz2) startDateTimeTz.AddDate().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	lineSplit := strings.Repeat("-", lineLen)

	fmt.Println(lineSplit)
	fmt.Printf("Starting Time Zone Local:   %v\n",
		startDateTime.Format(fmtStr))
	fmt.Println(lineSplit + "\n")

	fmt.Printf("Expected End Time Local:    %v\n",
		calcEndDateTimeLocal.Format(fmtStr))

	fmt.Printf("Actual End Time Local:      %v\n",
		endDateTimeTz1.GetDateTimeText())

	fmt.Println(lineSplit + "\n")

	fmt.Printf("Expected End Time UTC:      %v\n",
		calcEndDateTimeUTC.Format(fmtStr))

	fmt.Printf("Actual End Time UTC:        %v\n",
		endDateTimeTz2.GetDateTimeText())
	fmt.Println(lineSplit)
	fmt.Println(lineSplit)
}

func (mt mainTest) mainTest081() {
	/*
		Testing Change To Standard Time
		November 2, 2014

		TCalcMode.LocalTimeZone() and TCalcMode.UtcTimeZone()
		yield equal results even though 24-hours is added over
		a daylight savings time change.

	*/

	lineLen := 70

	titles := []string{"mainTest.mainTest081()",
		"Testing Change To Standard Time",
		"November 2, 2014",
		"Using TimeDurationDto and Time Math Calculation Modes",
	"LocalTimeZone & UtcTimeZone",
	"24-hours Duration Added Over Daylight Savings Time Change"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	var err error
	var centralTz *time.Location

	centralTz, err = time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time." +
			"LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	initialStartDate := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		centralTz)

	expectedDuration := time.Duration( dt.HourNanoSeconds * 24)

	localTzEndDate := initialStartDate.Add(expectedDuration)

	uTCTzEndDate :=initialStartDate.In(time.UTC)

	uTCTzEndDate = uTCTzEndDate.Add(expectedDuration)

	var tDur1, tDur2 dt.TimeDurationDto

	tDur1, err = dt.TimeDurationDto{}.NewStartTimeAddDuration(
		initialStartDate,
		expectedDuration,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.America.Chicago(),
		dt.TCalcMode.LocalTimeZone(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}." +
			"NewStartEndTimes(TCalcMode.LocalTimeZone()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	lineSplitter := strings.Repeat("-", lineLen)

	fmt.Printf(lineSplitter + "\n")

	fmt.Printf("Initial Start Date Time:    %v\n",
		initialStartDate.Format(dt.FmtDateTimeYrMDayFmtStr))

	actualStartDate := tDur1.GetThisStartDateTimeString()

	fmt.Printf("Actual Start Date Time #1:  %v\n",
		actualStartDate)

	fmt.Printf(lineSplitter + "\n")

	fmt.Printf("Expected Duration #1:       %v\n",
		expectedDuration)

	actualDuration := tDur1.GetThisTimeDuration()

	fmt.Printf("Actual Duration #1:         %v\n",
		actualDuration)

	fmt.Printf(lineSplitter + "\n")

	if expectedDuration != actualDuration {
		fmt.Printf("Error: Expected Local Time Duration = " +
			"25-hours.\n" +
			"Instead, actual duration = '%v'.", actualDuration.String())
		return
	}

	if actualStartDate != initialStartDate.Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error: Expected starting date time= '%v'.\n" +
			"Instead, actual starting date time= '%v'\n",
			initialStartDate.Format(dt.FmtDateTimeYrMDayFmtStr), actualStartDate)
		return
	}

	fmt.Printf("Expected End Date Time #1:  %v\n",
		localTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr))


	endDateTimeStr := tDur1.GetThisEndDateTimeString()

	fmt.Printf("Actual End Date Time #1:    %v\n",
		endDateTimeStr)

	fmt.Printf(lineSplitter + "\n")

	if endDateTimeStr != localTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error: Expected ending date time= '%v'.\n" +
			"Instead, actual starting date time= '%v'\n",
			localTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr), endDateTimeStr)
		return
	}

	tDur2, err = dt.TimeDurationDto{}.NewStartTimeAddDuration(
		initialStartDate,
		expectedDuration,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.America.Chicago(),
		dt.TCalcMode.UtcTimeZone(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}.NewStartEndTimes(TCalcMode.UtcTimeZone()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf(lineSplitter + "\n")

	fmt.Printf("Expected Duration #2:       %v\n",
		expectedDuration)


	actualDuration = tDur2.GetThisTimeDuration()

	fmt.Printf("Actual Duration #2:         %v\n",
		actualDuration)

	fmt.Printf(lineSplitter + "\n")

	if expectedDuration != actualDuration {
		fmt.Printf("Error: Expected UTC Time Duration = 25-hours.\n" +
			"Instead, actual duration = '%v'.", actualDuration.String())
		return
	}

	expectedEndDateTimeLocalTz := uTCTzEndDate.In(centralTz)

	fmt.Printf("Expected End Date Time #2:  %v\n",
		expectedEndDateTimeLocalTz.Format(dt.FmtDateTimeYrMDayFmtStr))



	actualEndDateTime := tDur1.GetThisEndDateTime()

	fmt.Printf("Actual End Date Time #2:    %v\n",
		actualEndDateTime.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Println()
	fmt.Printf(lineSplitter + "\n")

	actualEndDateTime = actualEndDateTime.In(time.UTC)
	endDateTimeStr = actualEndDateTime.Format(dt.FmtDateTimeYrMDayFmtStr)

	if endDateTimeStr != uTCTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error: Expected ending date time= '%v'.\n" +
			"Instead, actual ending date time= '%v'\n",
			uTCTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr), endDateTimeStr)
	}

	initialStartDateUTC := initialStartDate.In(time.UTC)
	fmt.Printf("Initial Start Date UTC:     %v\n",
		initialStartDateUTC.Format(dt.FmtDateTimeYrMDayFmtStr))

	expectedEndDateUTC := initialStartDateUTC.Add(expectedDuration)
	fmt.Printf("Expected End Date UTC:      %v\n",
		expectedEndDateUTC.Format(dt.FmtDateTimeYrMDayFmtStr))

	actualEndDateUTC := actualEndDateTime.In(time.UTC)
	fmt.Printf("Actual End Date UTC:        %v\n",
		actualEndDateUTC.Format(dt.FmtDateTimeYrMDayFmtStr))


	fmt.Printf(lineSplitter + "\n")

	fmt.Printf("Expected Duration:          %v\n",
		expectedDuration)

	actualDuration = expectedEndDateUTC.Sub(initialStartDateUTC)

	fmt.Printf("Actual Duration:            %v\n\n",
		actualDuration)

	convertedUTCStartDate := initialStartDateUTC.In(centralTz)

	fmt.Printf(lineSplitter + "\n")
	fmt.Printf("Converted UTC Start Date:   %v\n",
		convertedUTCStartDate.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Printf("Initial Start Date:         %v\n",
		initialStartDate.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Printf(lineSplitter + "\n\n")

	convertedUTCEndDate := expectedEndDateUTC.In(centralTz)

	fmt.Printf("Converted UTC End Date:     %v\n",
		convertedUTCEndDate.Format(dt.FmtDateTimeYrMDayFmtStr))
	actualEndDateTime = actualEndDateTime.In(centralTz)
	fmt.Printf("Original CTZ End Date:      %v\n",
		convertedUTCEndDate.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Printf(lineSplitter + "\n")
	fmt.Printf(lineSplitter + "\n\n")




}

func (mt mainTest) mainTest080() {
	/*
		Testing Change To Standard Time
		November 2, 2014

		TCalcMode.LocalTimeZone() and TCalcMode.UtcTimeZone()
		yield the same results because 25-hours duration is
	  added in both cases.

	*/

	lineLen := 70

	titles := []string{"mainTest.mainTest080()",
		"Testing Change To Standard Time",
		"November 2, 2014",
		"Using TimeDurationDto and Time Math Calculation Modes",
	"LocalTimeZone & UtcTimeZone"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	var err error
	var centralTz *time.Location

	centralTz, err = time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time." +
			"LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	initialStartDate := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		centralTz)

	localTzEndDate := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		centralTz)

	uTCTzEndDate := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		centralTz)

	var tDur1, tDur2 dt.TimeDurationDto

	expectedDuration := time.Duration( dt.HourNanoSeconds * 25)

	tDur1, err = dt.TimeDurationDto{}.NewStartTimeAddDuration(
		initialStartDate,
		expectedDuration,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.America.Chicago(),
		dt.TCalcMode.LocalTimeZone(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}." +
			"NewStartEndTimes(TCalcMode.LocalTimeZone()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	lineSplitter := strings.Repeat("-", lineLen)

	fmt.Printf(lineSplitter + "\n")

	fmt.Printf("Initial Start Date Time:    %v\n",
		initialStartDate.Format(dt.FmtDateTimeYrMDayFmtStr))

	actualStartDate := tDur1.GetThisStartDateTimeString()

	fmt.Printf("Actual Start Date Time #1:  %v\n",
		actualStartDate)

	fmt.Printf(lineSplitter + "\n")

	fmt.Printf("Expected Duration #1:       %v\n",
		expectedDuration)

	actualDuration := tDur1.GetThisTimeDuration()

	fmt.Printf("Actual Duration #1:         %v\n",
		actualDuration)

	fmt.Printf(lineSplitter + "\n")

	if expectedDuration != actualDuration {
		fmt.Printf("Error: Expected Local Time Duration = " +
			"25-hours.\n" +
			"Instead, actual duration = '%v'.", actualDuration.String())
		return
	}

	if actualStartDate != initialStartDate.Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error: Expected starting date time= '%v'.\n" +
			"Instead, actual starting date time= '%v'\n",
			initialStartDate.Format(dt.FmtDateTimeYrMDayFmtStr), actualStartDate)
		return
	}

	fmt.Printf("Expected End Date Time #1:  %v\n",
		localTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr))


	endDateTimeStr := tDur1.GetThisEndDateTimeString()

	fmt.Printf("Actual End Date Time #1:    %v\n",
		endDateTimeStr)

	fmt.Printf(lineSplitter + "\n")

	if endDateTimeStr != localTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error: Expected ending date time= '%v'.\n" +
			"Instead, actual starting date time= '%v'\n",
			localTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr), endDateTimeStr)
		return
	}

	tDur2, err = dt.TimeDurationDto{}.NewStartTimeAddDuration(
		initialStartDate,
		expectedDuration,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.America.Chicago(),
		dt.TCalcMode.UtcTimeZone(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}.NewStartEndTimes(TCalcMode.UtcTimeZone()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf(lineSplitter + "\n")

	fmt.Printf("Expected Duration #2:       %v\n",
		expectedDuration)


	actualDuration = tDur2.GetThisTimeDuration()

	fmt.Printf("Actual Duration #2:         %v\n",
		actualDuration)

	fmt.Printf(lineSplitter + "\n")

	if expectedDuration != actualDuration {
		fmt.Printf("Error: Expected UTC Time Duration = 25-hours.\n" +
			"Instead, actual duration = '%v'.", actualDuration.String())
		return
	}

	fmt.Printf("Expected End Date Time #2:  %v\n",
		uTCTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr))

	endDateTimeStr = tDur1.GetThisEndDateTimeString()

	actualEndDateTime := tDur1.GetThisEndDateTime()

	fmt.Printf("Actual End Date Time #2:    %v\n",
		endDateTimeStr)

	fmt.Println()
	fmt.Printf(lineSplitter + "\n")

	if endDateTimeStr != uTCTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error: Expected ending date time= '%v'.\n" +
			"Instead, actual starting date time= '%v'\n",
			uTCTzEndDate.Format(dt.FmtDateTimeYrMDayFmtStr), endDateTimeStr)
	}

	initialStartDateUTC := initialStartDate.In(time.UTC)
	fmt.Printf("Initial Start Date UTC:     %v\n",
		initialStartDateUTC.Format(dt.FmtDateTimeYrMDayFmtStr))

	expectedEndDateUTC := initialStartDateUTC.Add(expectedDuration)
	fmt.Printf("Expected End Date UTC:      %v\n",
		expectedEndDateUTC.Format(dt.FmtDateTimeYrMDayFmtStr))

	actualEndDateUTC := actualEndDateTime.In(time.UTC)
	fmt.Printf("Actual End Date UTC:        %v\n",
		actualEndDateUTC.Format(dt.FmtDateTimeYrMDayFmtStr))


	fmt.Printf(lineSplitter + "\n")

	fmt.Printf("Expected Duration:          %v\n",
		expectedDuration)

	actualDuration = expectedEndDateUTC.Sub(initialStartDateUTC)

	fmt.Printf("Actual Duration:            %v\n\n",
		actualDuration)

	convertedUTCStartDate := initialStartDateUTC.In(centralTz)

	fmt.Printf(lineSplitter + "\n")
	fmt.Printf("Converted UTC Start Date:   %v\n",
		convertedUTCStartDate.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Printf("Initial Start Date:         %v\n",
		initialStartDate.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Printf(lineSplitter + "\n\n")

	convertedUTCEndDate := expectedEndDateUTC.In(centralTz)
	fmt.Printf("Converted UTC End Date:     %v\n",
		convertedUTCEndDate.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Printf("Original CTZ End Date:      %v\n",
		endDateTimeStr)
	fmt.Printf(lineSplitter + "\n")
	fmt.Printf(lineSplitter + "\n\n")
}

func (mt mainTest) mainTest079() {
	lineLen := 70
	titles := []string{"mainTest.mainTest079()",
		"Testing Change To Standard Time",
		"November 2, 2014",
	"Using Time Math Calculation Mode = UtcTimeZone"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	var err error
	var centralTz *time.Location

	centralTz, err = time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1_1 := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("Initial Start Time:  %v\n",
		t1_1.Format(dt.FmtDateTimeYrMDayFmtStr))

	t1_2 := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("Initial End Time:    %v\n",
		t1_2.Format(dt.FmtDateTimeYrMDayFmtStr))

	var duration time.Duration
	var newStartDateTime, newEndDateTime dt.DateTzDto
/*
	duration,
	newStartDateTime,
	newEndDateTime,
	err =
	dtMech.ComputeDurationByUtc(
		t1_1,
		t1_2,
		dt.TZones.America.Chicago(),
		dt.FmtDateTimeYrMDayFmtStr,
		"mainTest078()")
*/

var tDur dt.TimeDurationDto

	tDur, err = dt.TimeDurationDto{}.NewStartEndTimes(
		t1_1,
		t1_2,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.America.Chicago(),
		dt.TCalcMode.UtcTimeZone(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimes().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	duration = tDur.GetThisTimeDuration()
	newStartDateTime = tDur.GetThisStartDateTimeTz()
	newEndDateTime = tDur.GetThisEndDateTimeTz()

	fmt.Printf("New Start Date Time: %v\n",
		newStartDateTime.GetDateTimeText())

	fmt.Printf("New End Date Time:   %v\n",
		newEndDateTime.GetDateTimeText())

	fmt.Printf("Time Duration:       %v\n",
		duration.String())

	initialStartDateTime := t1_1.In(time.UTC)

	fmt.Printf("Initial Start Date Time UTC: %v\n",
		initialStartDateTime.Format(dt.FmtDateTimeYrMDayFmtStr))

	initialEndDateTime := t1_2.In(time.UTC)

	fmt.Printf("Initial End Date Time UTC:   %v\n",
		initialEndDateTime.Format(dt.FmtDateTimeYrMDayFmtStr))

	utcDuration := initialEndDateTime.Sub(initialStartDateTime)

	fmt.Printf("UTC Duration:                %v\n",
		utcDuration.String())

}

func (mt mainTest) mainTest078() {
	lineLen := 70
	titles := []string{"mainTest.mainTest078()",
		"Testing Change To Standard Time",
		"November 2, 2014",
	"Using Time Math Calculation Mode = UtcTimeZone"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	var err error
	var centralTz *time.Location

	centralTz, err = time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1_1 := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("Initial Start Time:  %v\n",
		t1_1.Format(dt.FmtDateTimeYrMDayFmtStr))

	t1_2 := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("Initial End Time:    %v\n",
		t1_2.Format(dt.FmtDateTimeYrMDayFmtStr))

	dtMech := dt.DTimeMechanics{}

	var duration time.Duration
	var newStartDateTime, newEndDateTime dt.DateTzDto

	duration,
	newStartDateTime,
	newEndDateTime,
	err =
	dtMech.ComputeDurationByUtc(
		t1_1,
		t1_2,
		dt.TZones.America.Chicago(),
		dt.FmtDateTimeYrMDayFmtStr,
		"mainTest078()")

	if err != nil {
		fmt.Printf("Error returned by DTimeMechanics.ComputeDurationByUtc().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("New Start Date Time: %v\n",
		newStartDateTime.GetDateTimeText())

	fmt.Printf("New End Date Time:   %v\n",
		newEndDateTime.GetDateTimeText())

	fmt.Printf("Time Duration:       %v\n",
		duration.String())

	initialStartDateTime := t1_1.In(time.UTC)

	fmt.Printf("Initial Start Date Time UTC: %v\n",
		initialStartDateTime.Format(dt.FmtDateTimeYrMDayFmtStr))

	initialEndDateTime := t1_2.In(time.UTC)

	fmt.Printf("Initial End Date Time UTC:   %v\n",
		initialEndDateTime.Format(dt.FmtDateTimeYrMDayFmtStr))

	utcDuration := initialEndDateTime.Sub(initialStartDateTime)

	fmt.Printf("UTC Duration:                %v\n",
		utcDuration.String())

}

func (mt mainTest) mainTest077() {

	lineLen := 70
	titles := []string{"mainTest.mainTest077()",
		"Testing Change To Standard Time",
		"November 2, 2014"}


	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	centralTz, err := time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}


	t1_1 := time.Date(
		2014,
		10,
		31,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t1_1: '%v'\n", t1_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t1_2 := time.Date(
		2014,
		11,
		1,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t1_2: '%v'\n", t1_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t1Dur := t1_2.Sub(t1_1)

	fmt.Printf("t1Dur: '%v'\n", t1Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

	t2_1 := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t2_1: '%v'\n", t2_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t2_2 := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t2_2: '%v'\n", t2_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t2Dur := t2_2.Sub(t2_1)

	fmt.Printf("t2Dur: '%v'\n", t2Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

	t3_1 := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("t3_1: '%v'\n", t3_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t3_2 := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("t3_2: '%v'\n", t3_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t3Dur := t3_2.Sub(t3_1)

	fmt.Printf("t3Dur (UTC): '%v'\n", t3Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

}

func (mt mainTest) mainTest076() {

	lineLen := 70
	titles := []string{"mainTest.mainTest076()",
		"Testing Change to Daylight Savings Time",
		"March 9, 2014"}


	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	centralTz, err := time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1_1 := time.Date(
		2014,
		3,
		7,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t1_1: '%v'\n", t1_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t1_2 := time.Date(
		2014,
		3,
		8,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t1_2: '%v'\n", t1_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t1Dur := t1_2.Sub(t1_1)

	fmt.Printf("t1Dur: '%v'\n", t1Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

	t2_1 := time.Date(
		2014,
		3,
		9,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t2_1: '%v'\n", t2_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t2_2 := time.Date(
		2014,
		3,
		10,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t2_2: '%v'\n", t2_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t2Dur := t2_2.Sub(t2_1)

	fmt.Printf("t2Dur: '%v'\n", t2Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

	t3_1 := time.Date(
		2014,
		3,
		9,
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("t3_1: '%v'\n", t3_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t3_2 := time.Date(
		2014,
		3,
		10,
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("t3_2: '%v'\n", t3_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t3Dur := t3_2.Sub(t3_1)

	fmt.Printf("t3Dur (UTC): '%v'\n", t3Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

}


func (mt mainTest) mainTest075() {
	// Variation on:
	//   TestDurationTriad_NewStartEndDateTzDto_01
	// datetime\zztdurationtriad01_test.go

	lineLen := 70
	titles := []string{"mainTest.mainTest075()",
		"TestDurationTriad_NewStartEndDateTzDto_01",
		"DurationTriad and Time Duration Dto"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := dt.DateTzDto{}.NewDateTime(t1, dt.FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := dt.DateTzDto{}.NewDateTime(t2, dt.FmtDateTimeYrMDayFmtStr)

	dur, err := dt.DurationTriad{}.NewStartEndTimesTz(
		dateTz1,
		dateTz2,
		dt.TDurCalc.StdYearMth(),
		dateTz1.GetTimeZoneName(),
		dt.TCalcMode.LocalTimeZone(),
		dateTz1.GetDateTimeFmt())

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.GetThisStartDateTime().Format(fmtstr) {
		fmt.Printf("Error: Expected DurationTriad.StartTimeDateTz of %v.\n" +
			"Instead, got %v\n",
			t1OutStr, dur.BaseTime.GetThisStartDateTime().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.GetThisEndDateTime().Format(fmtstr) {
		fmt.Printf("Error: Expected DurationTriad.endDateTimeTz of %v.\n" +
			"Instead, got %v\n",
			t1OutStr, dur.BaseTime.GetThisEndDateTime().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != dur.BaseTime.GetThisTimeDuration() {
		fmt.Printf("Error: Expected DurationTriad.timeDuration of %v.\n" +
			"Instead, got %v\n",
			tOutDur, dur.BaseTime.GetThisTimeDuration())
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected :=
		"3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetYearsMthsWeeksTimeStr()

	expected =
		"3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		fmt.Printf("Error - Expected Default Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, _ = dur.BaseTime.GetCumDaysTimeStr()

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumHoursTimeStr()

	if err != nil {
		fmt.Printf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected Hours Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumWeeksDaysTimeStr()

	if err != nil {
		fmt.Printf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
	}

	loc, err := time.LoadLocation("Local")

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)
	startDateTimeTz := dur.LocalTime.GetThisStartDateTimeTz()
	if t1Local.Location().String() != startDateTimeTz.GetOriginalTzName() {
		fmt.Printf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'.\n",
			t1Local.Location().String(),
			startDateTimeTz.GetOriginalTzName())
	}

	if !t1Local.Equal(startDateTimeTz.GetDateTimeValue()) {
		fmt.Printf("Expected Local Start Time ='%v'.\n" +
			"Actual Local Start Time ='%v'.\n",
			t1Local.Format(dt.DEFAULTDATETIMEFORMAT),
			startDateTimeTz.GetDateTimeValue().Format(dt.DEFAULTDATETIMEFORMAT))
	}

	t2Local := dur.LocalTime.GetThisStartDateTime().Add(dur.LocalTime.GetThisTimeDuration())

	if !t2Local.Equal(dur.LocalTime.GetThisEndDateTime()) {
		fmt.Printf("Expected Local End Time='%v'.\n" +
			"Actual Local End Time='%v'.\n",
			t2Local.Format(dt.FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.GetThisEndDateTime().Format(dt.FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(dt.TZones.UTC())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(TZones.UTC()). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.GetThisStartDateTime()) {
		fmt.Printf("Expected UTC Start Time='%v'.\n" +
			"Actual UTC Start Time='%v'.\n",
			t1UTC.Format(dt.FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.GetThisStartDateTime().Format(dt.FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.GetThisStartDateTime().Add(dur.UTCTime.GetThisTimeDuration())

	if !t2UTC.Equal(dur.UTCTime.GetThisEndDateTime()) {
		fmt.Printf("Expected UTC End Time='%v'.\n" +
			"Actual UTC End Time='%v'.\n",
			t2UTC.Format(dt.FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.GetThisEndDateTime().Format(dt.FmtDateTimeYrMDayFmtStr))
	}

}

func (mt mainTest) mainTest074() {
	// Variation on:
	// TestDurationTriad_GetYearMthDaysTimeAbbrv
	// datetime\zztdurationtriad01_test.go
	lineLen := 70
	titles := []string{"mainTest.mainTest074()",
		"TestDurationTriad_GetYearMthDaysTimeAbbrv",
		"DurationTriad and Time Duration Dto"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := dt.DurationTriad{}

	err := du.SetStartEndTimes(
		t1,
		t2,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.US.Central(),
		dt.TCalcMode.LocalTimeZone(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by du.SetStartEndTimesTz(t2, t1, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr). Error='%v' ", err.Error())
	}

	fmt.Printf("Base Time Allocated Nanoseconds= '%v'\n",
		du.BaseTime.Nanoseconds)

	expected := "0-Hours 0-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := du.BaseTime.GetYearMthDaysTimeAbbrvStr()

	if expected != dOut {
		fmt.Printf("Expected result= '%v'.\n" +
			                " Instead result= '%v'.\n",
			                expected, dOut)
	}

}

func (mt mainTest) mainTest073() {
	// Variation on:
	// TestDateTzDto_AddDate_01
	// datetime\zztdatetzdto01_test.go
	lineLen := 70
	titles := []string{"mainTest.mainTest073()",
		"Testing AddDateTimeByUtc",
		"Adding \"5-Years, 6-Months 12-Days\""}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")


	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	locationPtr, err := time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-02-15 19:54:30.038175584 -0600 CST"
	t1 := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		locationPtr)

	t1Utc := t1.In(time.UTC)

	t1Result := t1.AddDate(5, 6, 12)

	t2Utc := t1Utc.AddDate(5, 6, 12)

	t2Result := t2Utc.In(locationPtr)

	titles = []string{"t1 Result",
		"Using Go Date Time Package Addition",
		"Adding 5-Years, 6-Months 12-Days"}

	ex.PrintOutDateTimeTimeZoneFields(
		t1Result,
		titles,
		lineLen,
		fmtStr)

	titles = []string{"t2 Result",
		"Using UTC Conversion",
		"Adding 5-Years, 6-Months 12-Days"}

	ex.PrintOutDateTimeTimeZoneFields(
		t2Result,
		titles,
		lineLen,
		fmtStr)

timeDto := dt.TimeDto{
	Years:                5,
	Months:               6,
	Weeks:                0,
	WeekDays:             0,
	DateDays:             12,
	Hours:                0,
	Minutes:              0,
	Seconds:              0,
	Milliseconds:         0,
	Microseconds:         0,
	Nanoseconds:          0,
	TotSubSecNanoseconds: 0,
	TotTimeNanoseconds:   0,
}

	dur, err := dt.DurationTriad{}.NewEndTimeMinusTimeDto(
		t2Result,
		timeDto,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.America.Chicago(),
		dt.TCalcMode.LocalTimeZone(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("\nCalculated Start Time: %v\n",
		dur.BaseTime.GetThisStartDateTime().Format(fmtStr))

	fmt.Printf("\n    Actual Start Time: %v\n\n",
		t1.Format(fmtStr))

	fmt.Printf("\nCalculated End Time: %v\n",
		dur.BaseTime.GetThisEndDateTime().Format(fmtStr))

	fmt.Printf("\n    Actual End Time: %v\n\n",
		t2Result.Format(fmtStr))

	fmt.Printf("\nCalculated UTC Start Time: %v\n",
		dur.UTCTime.GetThisStartDateTime().Format(fmtStr))

	fmt.Printf("\n    Actual UTC Start Time: %v\n\n",
		t1Utc.Format(fmtStr))

	fmt.Printf("\nCalculated UTC End Time: %v\n",
		dur.UTCTime.GetThisEndDateTime().Format(fmtStr))

	fmt.Printf("\n    Actual UTC End Time: %v\n\n",
		t2Utc.Format(fmtStr))

	fmt.Printf("\nCalculated BaseTime Duration: %v\n\n",
		dur.BaseTime.GetYearMthDaysTimeAbbrvStr())

	fmt.Printf("\n     Calculated UTC Duration: %v\n\n",
		dur.UTCTime.GetYearMthDaysTimeAbbrvStr())

	t2DurationDate := t1.Add(dur.UTCTime.GetThisTimeDuration())

	fmt.Printf("\n t1 + UTCTime Durtion: %v\n\n",
		t2DurationDate.Format(fmtStr))

	timeDur :=t1Result.Sub(t1)

	timeDurDto, err := dt.TimeDurationDto{}.NewStartTimeAddDuration(
	t1,
	timeDur,
	dt.TDurCalc.StdYearMth(),
	t1.Location().String(),
	dt.TCalcMode.LocalTimeZone(),
	fmtStr)

	if err != nil {
		fmt.Printf("Error retured by dt.TimeDurationDto{}.NewStartTimeAddDuration().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("\n                    t1 time: %v\n",
		t1.Format(fmtStr))

	fmt.Printf("\n             t1 result time: %v\n",
		t1Result.Format(fmtStr))
	
	fmt.Printf("\n     Calculated t1 Duration: %v\n\n",
		timeDurDto.GetYearMthDaysTimeAbbrvStr())

}

func (mt mainTest) mainTest072() {
	// TestDateTzDto_AddDate_01
	// datetime\zztdatetzdto01_test.go


	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	expectedOutDate := t1.Format(fmtstr)

	dtz1, err := dt.DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewStartEndTimes(t1, fmtstr).\n" +
			"Error='%v'", err.Error())
		return
	}

	if expectedOutDate != dtz1.String() {
		fmt.Printf("Error: Expected dtz1.String()='%v'. Instead, dtz1.String()='%v' ", expectedOutDate, dtz1.String())
		return
	}

	t2 := t1.AddDate(5, 6, 12)

	dtz2, err := dtz1.AddDate(
		dt.TCalcMode.LocalTimeZone(),
		5,
		6,
		12,
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dtz1.AddDate(5, 6, 12, fmtstr).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedOutDate = t2.Format(fmtstr)

	if expectedOutDate != dtz2.String() {
		fmt.Printf("Error: Expected dtz2.String()='%v'.\n" +
			"Instead, dtz2.String()='%v'\n", expectedOutDate, dtz2.String())
		return
	}

}

func (mt mainTest) mainTest071() {
	// TestDurationTriad_NewStartTimeMinusTime_01
	// \datetime\zztdurationtriad01_test.go
	//

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := dt.TimeDto{
		Years: 3, 
		Months: 2, 
		Weeks: 2, 
		WeekDays: 1, 
		Hours: 3, 
		Minutes: 4, 
		Seconds: 2}

	dur, err := dt.DurationTriad{}.NewEndTimeMinusTimeDto(
			t2,
			timeDto,
			dt.TDurCalc.StdYearMth(),
			dt.TZones.US.Central(),
			dt.TCalcMode.LocalTimeZone(),
			dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.GetThisStartDateTime().Format(fmtstr) {
		fmt.Printf("Error- Expected Start Time %v.\n" +
			"Instead, got %v.\n",
			t1OutStr, dur.BaseTime.GetThisStartDateTime().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.GetThisEndDateTime().Format(fmtstr) {
		fmt.Printf("Error- Expected End Time %v.\n" +
			"Instead, got %v.\n",
			t2OutStr, dur.BaseTime.GetThisEndDateTime().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.GetThisTimeDuration() {
		fmt.Printf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.GetThisTimeDuration())
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}

func (mt mainTest) mainTest070() {
	//t1str := "2020-03-07 03:00:00.000000000 -0500 EST"

	easternTzPtr, err := time.LoadLocation(dt.TZones.America.New_York())

	if err != nil {
		fmt.Printf("Error returned by " +
			"time.LoadLocation(dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1 := time.Date(
		2020,
		time.Month(3),
		7,
		3,
		0,
		0,
		0,
		easternTzPtr)

	expectedStr := "2020-03-08 04:00:00.000000000 -0400 EDT"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	funcName := "mainTest070()"

	lineLen := 70
	titles := []string{funcName,
		"Testing DateTimeOps DateTimeMechanics AddDateTimeByUtc",
		"Adding 1-Day To Date Time With IANA Time Zone",
	"Notice That Actual Matches Expected Result"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")


	titles = []string{"Time-1 (t1)",
		"Starting Date 1-day before Daylight Standard Time"}

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		titles,
		lineLen,
		fmtStr)

	dtMech := dt.DTimeMechanics{}

	t2 := dtMech.AddDateTimeByUtc(
		t1,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		0)

	titles = []string{"Time-2 (t2)",
		"Actual Date 1-day After Daylight Standard Time",
		"Includes added value of 1-Day"}

	ex.PrintOutDateTimeTimeZoneFields(
		t2,
		titles,
		lineLen,
		fmtStr)

	lineBreak := strings.Repeat("*", lineLen)
	fmt.Println()
	fmt.Println(lineBreak)
	fmt.Println("        Starting Date Time: ", t1.Format(fmtStr))
	fmt.Println("  Actual After 1-Day Added: ", t2.Format(fmtStr))
	fmt.Println("Expected After 1-Day Added: ", expectedStr)
	fmt.Println(lineBreak)
	fmt.Println()

	titles = []string{funcName, "Successful Completion!"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

}

func (mt mainTest) mainTest069() {
	// t1str := "2020-03-07 03:00:00.000000000 -0500 EST"
	// EST Time Zone Fails

	easternTzPtr, err := time.LoadLocation(dt.TZones.America.New_York())

	if err != nil {
		fmt.Printf("Error returned by " +
			"time.LoadLocation(dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1 := time.Date(
		2020,
		time.Month(3),
		7,
		3,
		0,
		0,
		0,
		easternTzPtr)

	expectedStr := "2020-03-08 04:00:00.000000000 -0400 EDT"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	funcName := "mainTest069()"

	lineLen := 70
	titles := []string{funcName,
		"Testing DateTimeOps DateTimeMechanics AddDateTimeByUtc",
		"Adding 1-Day",
		"Notice That Actual Matches Expected Result",
		"When Using IANA Convertible Time Zone"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	titles = []string{"Time-1 (t1)",
		"Starting Date 1-day before Daylight Standard Time"}

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		titles,
		lineLen,
		fmtStr)

	dtMech := dt.DTimeMechanics{}

	t2 := dtMech.AddDateTimeByUtc(
		t1,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		0)

	titles = []string{"Time-2 (t2)",
		"Actual Date 1-day After Daylight Standard Time",
		"Includes added value of 1-Day"}

	ex.PrintOutDateTimeTimeZoneFields(
		t2,
		titles,
		lineLen,
		fmtStr)

	lineBreak := strings.Repeat("*", lineLen)
	fmt.Println()
	fmt.Println(lineBreak)
	fmt.Println("        Starting Date Time: ", t1.Format(fmtStr))
	fmt.Println("  Actual After 1-Day Added: ", t2.Format(fmtStr))
	fmt.Println("Expected After 1-Day Added: ", expectedStr)
	fmt.Println(lineBreak)

	var funcResult string

	if t2.Format(fmtStr) != expectedStr {
		titles = []string{"FAILURE!","Expected Date Time Does NOT Match Actual Date Time!"}
		ex.PrintMainHeader(
			titles,
			lineLen,
			"",
			"*")

		funcResult = "@@ FAILURE @@"

	} else {
		titles = []string{"SUCCESS!","Expected Date Time DOES Match Actual Date Time!"}
		ex.PrintMainHeader(
			titles,
			lineLen,
			"",
			"*")

		funcResult = "Successful Completion!"
	}

	fmt.Println()

	titles = []string{funcName, funcResult}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

}

func (mt mainTest) mainTest068() {
	// 8 Mar 2020 - Daylight Saving Time Starts
	// When local standard time is about to reach
	// Sunday, 8 March 2020, 02:00:00 clocks are turned forward 1 hour to
	//Sunday, 8 March 2020, 03:00:00 local daylight time instead.

	// t1str := "2020-03-07 03:01:00.000000000 -0500 EST"


	easternTzPtr, err := time.LoadLocation(dt.TZones.America.New_York())

	if err != nil {
		fmt.Printf("Error returned by " +
			"time.LoadLocation(dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1 := time.Date(
		2020,
		time.Month(3),
		7,
		3,
		0,
		0,
		0,
		easternTzPtr)

	expectedStr := "2020-03-08 04:00:00.000000000 -0400 EDT"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	funcName := "mainTest068()"

	lineLen := 70
	titles := []string{funcName,
		"Testing Golang DateTime Package AddDate",
		"Adding 1-Day",
		"Notice that the Final Date is WRONG!"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	titles = []string{"Time-1 (t1)",
		"Starting Date 1-day before Daylight Standard Time"}

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		titles,
		lineLen,
		fmtStr)

	t2 := t1.AddDate(0, 0, 1)

	titles = []string{"Time-2 (t2)",
		"Actual Date 1-day After Daylight Standard Time",
	"Includes added value of 1-Day"}

	ex.PrintOutDateTimeTimeZoneFields(
		t2,
		titles,
		lineLen,
		fmtStr)

	lineBreak := strings.Repeat("*", lineLen)
	fmt.Println()
	fmt.Println(lineBreak)
	fmt.Println("        Starting Date Time: ", t1.Format(fmtStr))
	fmt.Println("  Actual After 1-Day Added: ", t2.Format(fmtStr))
	fmt.Println("Expected After 1-Day Added: ", expectedStr)
	fmt.Println(lineBreak)

	var funcResult string

	if t2.Format(fmtStr) != expectedStr {
		titles = []string{"Expected Date Time Does NOT Match Actual Date Time!"}
		ex.PrintMainHeader(
			titles,
			lineLen,
			"",
			"*")

		funcResult = "@@ FAILURE @@"

	} else {
		titles = []string{"Expected Date Time DOES Match Actual Date Time!"}
		ex.PrintMainHeader(
			titles,
			lineLen,
			"",
			"*")

		funcResult = "Successful Completion!"
	}

	fmt.Println()

	titles = []string{funcName, funcResult}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

}

func (mt mainTest) mainTest067() {

// 8 Mar 2020 - Daylight Saving Time Starts
// When local standard time is about to reach
// Sunday, 8 March 2020, 02:00:00 clocks are turned forward 1 hour to
//Sunday, 8 March 2020, 03:00:00 local daylight time instead.

		t1str := "2020-03-08 01:00:00.000000000 -0600 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	funcName := "mainTest067()"

	lineLen := 65
	titles := []string{funcName,
		"Edge Of Daylight Time",
	"Adding 2-Hours"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	lineBreak := strings.Repeat("-", lineLen)
	lineBreak2 := strings.Repeat("*", lineLen)
	t1, _ := time.Parse(fmtStr, t1str)

	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.NewStartEndTimes(t1, TzUsPacific).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var pacificPtr, utcPtr, localPtr *time.Location

	pacificPtr, err = time.LoadLocation(dt.TZones.US.Pacific())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.US.Pacific()).\n" +
			"dt.TZones.US.Pacific()='%v'\n" +
			"Error='%v'\n", dt.TZones.US.Pacific(), err.Error())
		return
	}

	utcPtr, err = time.LoadLocation(dt.TZones.UTC())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.UTC()).\n" +
			"dt.TZones.UTC()='%v'\n" +
			"Error='%v'\n", dt.TZones.UTC(), err.Error())
		return
	}

	localPtr, err = time.LoadLocation(dt.TZones.Local())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.Local()).\n" +
			"dt.TZones.UTC()='%v'\n" +
			"Error='%v'\n", dt.TZones.Local(), err.Error())
		return
	}

	t1In := t1
	t1Out := t1.In(pacificPtr)
	t1Utc := t1.In(utcPtr)
	t1Local := t1.In(localPtr)
	fmt.Println(lineBreak)
	fmt.Println(lineBreak)
	fmt.Println("   Calculated t1 In: ", t1In.Format(fmtStr))
	fmt.Println("  Calculated t1 Out: ", t1Out.Format(fmtStr) )
	fmt.Println("  Calculated t1 UTC: ", t1Utc.Format(fmtStr) )
	fmt.Println("Calculated t1 Local: ", t1Local.Format(fmtStr) )
	fmt.Println(lineBreak)
	fmt.Println()
	ex.PrintOutTimeZoneDtoFields(tzu1, "tzu1 - Before Time Addition")
	fmt.Println(lineBreak2)
	fmt.Println()

	err = tzu1.AddTime(
		2,
		0,
		0,
		0,
		0,
		0)

	if err != nil {
		fmt.Printf("Error returned by tzu1.AddTime(Add 2-hours)\n" +
			"Error='%v'\n", err.Error())
		return
	}
	fmt.Println(lineBreak2)
	ex.PrintOutTimeZoneDtoFields(tzu1, "tzu1 - After 2-Hour Addition")
	fmt.Println(lineBreak2)
	fmt.Println()
}

func (mt mainTest) mainTest066() {
	t1str := "2014-02-15 19:54:30.000000000 -0500 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	funcName := "mainTest066()"

	lineLen := 65
	titles := []string{funcName,
		"Time Zone Dto Comparison" }

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	lineBreak := strings.Repeat("-", lineLen)
	lineBreak2 := strings.Repeat("*", lineLen)

	t1, _ := time.Parse(fmtStr, t1str)
	t1OutStr := t1.Format(fmtStr)
	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.NewStartEndTimes(t1, TzUsPacific).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var pacificPtr, utcPtr, localPtr *time.Location

	pacificPtr, err = time.LoadLocation(dt.TZones.US.Pacific())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.US.Pacific()).\n" +
			"dt.TZones.US.Pacific()='%v'\n" +
			"Error='%v'\n", dt.TZones.US.Pacific(), err.Error())
		return
	}

	utcPtr, err = time.LoadLocation(dt.TZones.UTC())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.UTC()).\n" +
			"dt.TZones.UTC()='%v'\n" +
			"Error='%v'\n", dt.TZones.UTC(), err.Error())
		return
	}

	localPtr, err = time.LoadLocation(dt.TZones.Local())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.Local()).\n" +
			"dt.TZones.UTC()='%v'\n" +
			"Error='%v'\n", dt.TZones.Local(), err.Error())
		return
	}

	t1Out := t1.In(pacificPtr)
	t1Utc := t1.In(utcPtr)
	t1Local := t1.In(localPtr)

	fmt.Println(lineBreak)
	fmt.Println("   Expected t1 In: ", t1.Format(fmtStr))
	fmt.Println("  Expected t1 Out: ", t1Out.Format(fmtStr) )
	fmt.Println("  Expected t1 UTC: ", t1Utc.Format(fmtStr) )
	fmt.Println("Expected t1 Local: ", t1Local.Format(fmtStr) )
	fmt.Println(lineBreak)
	fmt.Println()
	ex.PrintOutTimeZoneDtoFields(tzu1, "tzu1")
	fmt.Println(lineBreak)
	fmt.Println()


	tzu1OutStrTIn := tzu1.TimeIn.GetDateTimeValue().Format(fmtStr)

	if t1OutStr != tzu1OutStrTIn {
		fmt.Printf("Error: Expected tzu1OutStrTIn='%v'.  Instead, tzu1OutStrTIn='%v'", t1OutStr, tzu1OutStrTIn)
	}

	t2 := t1.AddDate(3, 2, 15)
	// t2OutStr := t2.Format(fmtStr)

	tzu2, err := dt.TimeZoneDto{}.NewAddDate(tzu1, 3, 2, 15, fmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.NewAddDate(tzu1, 3, 2, 15, fmtStr)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t2Out := t2.In(pacificPtr)
	t2Utc := t2.In(utcPtr)
	t2Local := t2.In(localPtr)

	t3Utc := t1Utc.AddDate(3, 2, 15)

	fmt.Println(lineBreak2)
	fmt.Println("   Expected t2 In: ", t2.Format(fmtStr))
	fmt.Println("  Expected t2 Out: ", t2Out.Format(fmtStr) )
	fmt.Println("  Expected t2 UTC: ", t2Utc.Format(fmtStr) )
	fmt.Println("Expected t2 Local: ", t2Local.Format(fmtStr) )
	fmt.Println("      t3Utc Added: ", t3Utc.Format(fmtStr))
	fmt.Println(lineBreak)
	fmt.Println()
	ex.PrintOutTimeZoneDtoFields(tzu2, "tzu2")
	fmt.Println(lineBreak2)


	/*
		tzu2OutStrTIn := tzu2.TimeIn.GetDateTimeValue().Format(fmtStr)

		if t2OutStr != tzu2OutStrTIn {
			fmt.Printf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
			return
		}

		actualDuration, err := tzu2.Sub(tzu1)

		if err != nil {
			fmt.Printf("Error returned by tzu2.Sub(tzu1). Error='%v'", err.Error())
			return
		}

		utcPtr, err := time.LoadLocation(dt.TZones.UTC())

		if err != nil{
			fmt.Printf("Error return from time.LoadLocation(TZones.UTC()).\n" +
				"Error='%v'\n", err.Error())
			return
		}

		t1UTC := t1.In(utcPtr)

		t2UTC := t2.In(utcPtr)

		expectedDuration := t2UTC.Sub(t1UTC)

		if expectedDuration != actualDuration {
			fmt.Printf("Error: Expected Duration='%v'.\n" +
				"Instead, Actual Duration='%v'\n",
				expectedDuration, actualDuration)
		}

		*/

}

func (mt mainTest) mainTest065() {

	funcName := "mainTest065()"
	lineLen := 65
	titles := []string{funcName,
											"UTC Duration Comparison" }

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	t1str := "2014-02-14 19:54:30.000000000 -0600 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, err := time.Parse(fmtStr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtStr, t1str)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var utcPtr *time.Location

	utcPtr, err = time.LoadLocation(dt.TZones.UTC())

	if err != nil{
		fmt.Printf("Error return from time.LoadLocation(TZones.UTC()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	equalHeader := strings.Repeat("=", 65)
	fmt.Println(equalHeader)
	fmt.Println("   t1 Start Date Time: ", t1.Format(fmtStr))
	t1Utc := t1.In(utcPtr)
	fmt.Println("     t1 UTC Date Time: ", t1Utc.Format(fmtStr))
	fmt.Println(equalHeader)
	fmt.Println()

	t2 := t1.AddDate(3, 2, 15)
	fmt.Println(equalHeader)
	fmt.Println("   t2 Start Date Time: ", t2.Format(fmtStr))
	t2Utc := t2.In(utcPtr)
	fmt.Println("     t2 UTC Date Time: ", t2Utc.Format(fmtStr))
	fmt.Println(equalHeader)
	fmt.Println()

	t1_2BaseDuration := t2.Sub(t1)
	fmt.Println("  t1-t2 Base Duration: ", t1_2BaseDuration.String())

	t1_2UtcDuration := t2Utc.Sub(t1Utc)
	fmt.Println("   t1-t2 Utc Duration: ", t1_2UtcDuration.String())
	fmt.Println(equalHeader)
	fmt.Println()


	titles = []string{funcName,
		"Successful Completion" }

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

}

func (mt mainTest) mainTest064() {
	// TestDateTzDto_AddDate_01

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	expectedOutDate := t1.Format(fmtstr)

	dtz1, err := dt.DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewStartEndTimes(t1, fmtstr).\n" +
			"Error='%v'", err.Error())
		return
	}

	if expectedOutDate != dtz1.String() {
		fmt.Printf("Error: Expected dtz1.String()='%v'. Instead, dtz1.String()='%v' ", expectedOutDate, dtz1.String())
	}

	t2 := t1.AddDate(5, 6, 12)

	dtz2, err := dtz1.AddDate(
		dt.TCalcMode.LocalTimeZone(),
		5,
		6,
		12,
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dtz1.AddDate(5, 6, 12, fmtstr). Error='%v'", err.Error())
		return
	}

	expectedOutDate = t2.Format(fmtstr)

	if expectedOutDate != dtz2.String() {
		fmt.Printf("Error: Expected dtz2.String()='%v'. Instead, dtz2.String()='%v' ", expectedOutDate, dtz2.String())
	}

}

func (mt mainTest) mainTest063() {

	ePrefix := "mainTest063() TestTDurCalcTypeString_034"

	mt.mainPrintHdr(ePrefix , "-")

	expectedStr := "XRAYxxx"

	_, err :=  dt.TDurCalc.XParseString(expectedStr, false)

	if err == nil {
		fmt.Printf("Error: Expected an 'error' return from " +
			"TDurCalc.XParseString(expectedStr, false).\n" +
			"expectedStr='%v'.\n" +
			"NO ERROR WAS RETURNED!!! ",expectedStr)
		return
	}

	mt.mainPrintHdr("Successful Completion" , "=")
}

func (mt mainTest) mainTest062() {

	ePrefix := "mainTest062() TimeZoneDefinition"

	mt.mainPrintHdr(ePrefix , "-")

	lineLen := 65

	utcOffset := "2020-01-19 04:21:18 +0700 +07"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	expectedOriginalTz := "Etc/GMT-7"
	expectedConvertibleTz := dt.TZones.Asia.Ho_Chi_Minh()

	utcOffsetTime, err := time.Parse(fmtStr, utcOffset)

	if err != nil {
		fmt.Printf("Received error from time parse utcOffset: %v\n",
			err.Error())
		return
	}

	var tzDef dt.TimeZoneDefinition

	tzDef, err = dt.TimeZoneDefinition{}.NewDateTime(utcOffsetTime)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefinition{}.NewStartEndTimes(utcOffsetTime)\n" +
			"utcOffsetTime= '%v'\n" +
			"Error='%v'\n",
			utcOffsetTime.Format(fmtStr), err.Error())
		return
	}

	originalTzSpec := tzDef.GetOriginalTimeZone()
	
	ex.PrintOutDateTimeTimeZoneFields(
		originalTzSpec.GetReferenceDateTime(),
		[]string{"Original TZ Date Time"},
		lineLen,
		fmtStr)


	ex.PrintOutTimeZoneSpecFields(originalTzSpec, "Original Time Zone")


	convertibleTzSpec := tzDef.GetConvertibleTimeZone()
	
	ex.PrintOutDateTimeTimeZoneFields(
		convertibleTzSpec.GetReferenceDateTime(),
		[]string{"Convertible TZ Date Time"},
		lineLen,
		fmtStr)

	ex.PrintOutTimeZoneSpecFields(convertibleTzSpec, "Convertible Time Zone")

	actualOriginalTz := tzDef.GetOriginalTimeZoneName()

	actualConvertibleTz := tzDef.GetConvertibleTimeZoneName()


	if expectedOriginalTz != actualOriginalTz {
		fmt.Printf("Error: Expected actualOriginalTz='%v'.\n" +
			"Instead, actualOriginalTz='%v'\n",
			expectedOriginalTz, actualOriginalTz)
		return
	}

	if expectedConvertibleTz != actualConvertibleTz {
		fmt.Printf("Error: Expected actualConvertibleTz='%v'.\n" +
			"Instead, actualConvertibleTz='%v'\n",
			expectedConvertibleTz, actualOriginalTz)
		return
	}

	mt.mainPrintHdr("Successful Completion" , "=")

}

func (mt mainTest) mainTest061() {

	ePrefix := "mainTest061() TimeZoneDto.ConvertTz()"

	mt.mainPrintHdr(ePrefix , "-")

	utcTime := "2017-04-30 00:54:30 +0000 UTC"
	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	mountainTime := "2017-04-29 18:54:30 -0600 MDT"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	ianaCentralTz := "America/Chicago"
	ianaMountainTz := "America/Denver"
	tPacificIn, err := time.Parse(fmtstr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	tzu := dt.TimeZoneDto{}
	tzuCentral, err := tzu.ConvertTz(
		tPacificIn,
		ianaCentralTz,
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("#1 Error from ianaCentralTz TimeZoneDto.ConvertTz().\n" +
			"Error: %v\n", err.Error())
		return
	}

	centralTOut := tzuCentral.TimeOut.GetDateTimeValue().Format(fmtstr)

	if centralTime != centralTOut {
		fmt.Printf("Expected tzuCentral.TimeOut = '%v'.\n" +
			"Instead, tzuCentral.TimeOut = '%v'.\n",
			centralTime, centralTOut)
		return
	}

	tzuMountain, err := tzu.ConvertTz(
		tzuCentral.TimeOut.GetDateTimeValue(),
		ianaMountainTz, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("#2 Error from tzuMountain TimeZoneDto.ConvertTz().\n" +
			"Error: %v\n", err.Error())
		return
	}

	mountainTOut := tzuMountain.TimeOut.GetDateTimeValue().Format(fmtstr)

	if mountainTime != mountainTOut {
		fmt.Printf("Expected tzuMountain.TimeOut= '%v'.\n" +
			"Instead, tzuMountain.TimeOut= '%v'.\n",
			mountainTime, mountainTOut)
		return
	}

	tzuPacific, err := tzu.ConvertTz(
		tzuMountain.TimeOut.GetDateTimeValue(),
		ianaPacificTz, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error #3 from tzuMountain TimeZoneDto.ConvertTz().\n" +
			"Error: %v\n", err.Error())
		return
	}

	pacificTOut := tzuPacific.TimeOut.GetDateTimeValue().Format(fmtstr)

	if pacificTime != pacificTOut {

		fmt.Printf("Expected tzuPacific.TimeOut= '%v'.\n" +
			"Instead, tzuPacific.TimeOut= '%v'.\n",
			pacificTime, pacificTOut)
		return
	}

	exTOutLoc := "America/Los_Angeles"

	if exTOutLoc != tzuPacific.TimeOut.GetOriginalTzName() {
		fmt.Printf("Expected tzu.TimeOutLoc='%v'.\n" +
			"Instead tzu.TimeOutLoc='%v'\n" +
			"tzuPacific.TimeOut='%v'\n",
			exTOutLoc,
			tzuPacific.TimeOut.GetOriginalTzName(),
			tzuPacific.TimeOut.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	pacificUtcOut := tzuPacific.TimeUTC.GetDateTimeValue().Format(fmtstr)

	if utcTime != pacificUtcOut {
		fmt.Printf("Expected tzuPacific.TimeUTC= '%v'\n" +
			"Instead, tzuPacific.TimeUTC= '%v'\n",
			utcTime, pacificUtcOut)
		return
	}

	centralUtcOut := tzuCentral.TimeUTC.GetDateTimeValue().Format(fmtstr)

	if utcTime != centralUtcOut {
		fmt.Printf("Expected tzuCentral.TimeUTC= '%v'\n" +
			"Instead, tzuCentral.TimeUTC= '%v'\n",
			utcTime, pacificUtcOut)
		return
	}

	mountainUtcOut := tzuMountain.TimeUTC.GetDateTimeValue().Format(fmtstr)

	if utcTime != mountainUtcOut {
		fmt.Printf("Expected tzuMountain.TimeUTC= '%v'\n" +
			"Instead, tzuMountain.TimeUTC= '%v'\n",
			utcTime, pacificUtcOut)
		return
	}

	fmt.Println()
	mt.mainPrintHdr("Successful Completion" , "=")
}


func (mt mainTest) mainTest060() {

	ePrefix := "mainTest060()"
	lineLen := 65

	mt.mainPrintHdr(ePrefix , "-")
	mt.mainPrintHdr("TimeZoneDto" , "-")

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtStr, t1str)
	t1OutStr := t1.Format(fmtStr)
	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtStr)

	if err != nil {
		fmt.Printf("Error returned from TimeZoneDto{}.NewStartEndTimes(t1, TzUsPacific ). Error='%v'", err.Error())
		return
	}

	fmt.Println("t1OutStr: ", t1OutStr)

	ex.PrintOutDateTimeTimeZoneFields(
		tzu1.TimeIn.GetDateTimeValue(),
		[]string{"Initial tzu1.TimeIn"},
		lineLen,
		fmtStr)

	t2, _ := time.Parse(fmtStr, t2str)
	t2OutStr := t2.Format(fmtStr)

	t12Dur := t2.Sub(t1)

	tdurDto, err := dt.TimeDurationDto{}.NewStartEndTimes(
		t1,
		t2,
		dt.TDurCalcType(0).StdYearMth(),
		dt.TZones.US.Central(),
		dt.TCalcMode.LocalTimeZone(),
		fmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}.NewStartEndTimesCalcTz()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddTimeDurationDto(tdurDto)

	if err != nil {
		fmt.Printf("Error returned by tzu2.AddTimeDurationDto(tdurDto).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	tzu1OutStr := tzu1.TimeIn.GetDateTimeValue().Format(fmtStr)

	if t1OutStr != tzu1OutStr {
		fmt.Printf("Error: Expected Time1 TimeIn='%v'.\n" +
			"Instead Time1 TimeIn='%v'\n",
			t1OutStr, tzu1OutStr)
		return
	}

	tzu2OutStr := tzu2.TimeIn.GetDateTimeValue().Format(fmtStr)

	if t2OutStr != tzu2OutStr {
		fmt.Printf("Error: Expected after duration tzu2TimeIn='%v'.\n" +
			"Instead, tzu2TimeIn='%v'\n",
			t2OutStr, tzu2OutStr)
		return
	}

	actualDur := tzu2.TimeIn.Sub(tzu1.TimeIn)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeIn.Sub(tzu1.TimeIn)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeOut.Sub(tzu1.TimeOut)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeOut.Sub(tzu1.TimeOut)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
	}

	actualDur = tzu2.TimeUTC.Sub(tzu1.TimeUTC)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeUTC.Sub(tzu1.TimeUTC)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeLocal.Sub(tzu1.TimeLocal)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeLocal.Sub(tzu1.TimeLocal)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualTimeOutLoc := tzu1.TimeOut.GetOriginalTzName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu1.TimeOutLoc='%v'.\n" +
			"Instead, tzu1.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	actualTimeOutLoc = tzu2.TimeOut.GetOriginalTzName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu2.TimeOutLoc.String()='%v'.\n" +
			"Instead, tzu2.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	mt.mainPrintHdr("Successful Completion" , "=")
}

func (mt mainTest) mainTest059() {
	ePrefix := "mainTest059()"

	mt.mainPrintHdr(ePrefix , "-")
	mt.mainPrintHdr("ConvertTz" , "-")

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	// Invalid Target Iana Time Zone
	invalidTz := "XUZ Time Zone"
	tIn, err := time.Parse(fmtstr, tstr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtstr, tstr)\n" +
			"tstr='%v'\n" +
			"Error='%v'\n", tstr, err.Error())
		return
	}

	tzu := dt.TimeZoneDto{}

	_, err = tzu.ConvertTz(tIn, invalidTz, fmtstr)

	if err == nil {
		fmt.Printf("ConvertTz() failed to detect INVALID Target Time Zone.\n" +
			"err=='nil'\n" +
			"invalidTz='%v'\n", invalidTz)
	}

	return
}

func (mt mainTest) mainTest058() {
	ePrefix := "mainTest058()"
	lineLen := 65

	mt.mainPrintHdr(ePrefix , "-")
	mt.mainPrintHdr("ConvertTzAbbreviationToTimeZone" , "-")

	tzAbbrv := "+10"
	testUtcOffset := "+1000"

	tzMech := dt.TimeZoneMechanics{}

	dateTime := time.Now().UTC()

	var staticTimeZone dt.TimeZoneSpecification
	var err error

	staticTimeZone,
	err = tzMech.ConvertUtcAbbrvToStaticTz(
		dateTime,
		dt.TzConvertType.Relative(),
		"Original Time Zone",
		testUtcOffset,
		ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	fmt.Println("staticTimeZone: ", staticTimeZone.GetLocationName())
	fmt.Println()
	ex.PrintOutDateTimeTimeZoneFields(
		staticTimeZone.GetReferenceDateTime(),
		[]string{"staticTimeZone"},
		lineLen,
		fmtStr)

	ex.PrintOutTimeZoneSpecFields(staticTimeZone, "staticTimeZone")

	var tzSpec dt.TimeZoneSpecification

	tzSpec,
	err =
	tzMech.ConvertTzAbbreviationToTimeZone(
		dateTime,
		dt.TimeZoneConversionType(0).Relative(),
		tzAbbrv+testUtcOffset,
		"",
		ePrefix)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by tzMech.ConvertTzAbbreviationToTimeZone()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(
		tzSpec.GetReferenceDateTime(),
		[]string{"tzSpec"},
		lineLen,
		fmtStr)

	ex.PrintOutTimeZoneSpecFields(tzSpec, "tzSpec")


	mt.mainPrintHdr("SUCCESS" , "!")

}

func (mt mainTest) mainTest057() {

	ePrefix := "mainTest057()"
	lineLen := 65

	mt.mainPrintHdr(ePrefix , "-")

	dtMech := dt.DTimeMechanics{}

	locPtr, err := dtMech.LoadTzLocation(dt.TZones.Asia.Vladivostok(), ePrefix)

	if err != nil {
		fmt.Printf("Error='%v'\n", err.Error())
		return
	}

	dateTime := time.Date(
		2019,
		time.Month(6),
		15,
		11,
		23,
		0,
		0,
		locPtr)

	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	tzDef, err := dt.TimeZoneDefinition{}.NewFromTimeZoneName(
		dateTime, dt.TZones.Asia.Vladivostok(), dt.TzConvertType.Absolute())

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(dateTime, London Time)\n" +
			"dateTime='%v'\nError='%v'\n",
			dateTime.Format(fmtStr), err.Error())
		return
	}

	ex.PrintOutTimeZoneDefFields(tzDef)


	// t1EdtStr :=  "06/20/2019 09:58:32 -0400 EDT"
	// t1EstStr :=  "12/20/2019 09:58:32 -0500 EST"
	// t2CdtStr :=  "06/20/2019 09:58:32 -0500 CDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0600 CST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0600 MDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0700 MST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0700 PDT"
	// t2PstStr :=  "12/20/2019 09:58:32 -0800 PST"

	timeStr := dateTime.Format(fmtStr)

	t2, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(
		t2,
		[]string{"t2 Parse Result"},
		lineLen,
		fmtStr)

	tzDef, err = dt.TimeZoneDefinition{}.NewDateTime(t2)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(t1)\n" +
			"t2='%v'\nError='%v'\n", t2.Format(fmtStr), err.Error())
		return
	}

	ex.PrintOutTimeZoneDefFields(tzDef)

	mt.mainPrintHdr("SUCCESS" , "!")
}

func (mt mainTest) mainTest056() {

	ePrefix := "mainTest056()"

	mt.mainPrintHdr(ePrefix , "-")

	dtMech := dt.DTimeMechanics{}
	locUSCentral, err :=
		dtMech.LoadTzLocation(dt.TZones.US.Central(), ePrefix)

	if err != nil {
		fmt.Printf("Error='%v'\n", err.Error())
		return
	}

	locTokyo, err := dtMech.LoadTzLocation(dt.TZones.Asia.Tokyo(), ePrefix)

	if err != nil {
		fmt.Printf("Error='%v'", err.Error())
		return
	}

	t4USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)

	t4AsiaTokyo := t4USCentral.In(locTokyo)

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dTz, err := dt.DateTzDto{}.NewTz(
		t4AsiaTokyo,
		dt.TZones.US.Central(),
		dt.TzConvertType.Relative(),
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewTz(t4AsiaTokyo, TZones.US.Central(), fmtstr).\n" +
			"Error='%v'\n",
			err.Error())
		return
	}

	if !t4USCentral.Equal(dTz.GetDateTimeValue()) {
		fmt.Printf("Error: Expected DateTime='%v'. Instead DateTime='%v'",
			t4USCentral.Format(fmtstr), dTz.GetDateTimeValue().Format(fmtstr))
	}

	eTimeZoneDef, err := dt.TimeZoneDefinition{}.NewDateTime(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefinition{}.NewStartEndTimes(t4USCentral)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	areEqual := eTimeZoneDef.Equal(dTz.GetTimeZoneDef())

	if ! areEqual {
		fmt.Printf("Expected dTz.GetTimeZoneDef().LocationName='%v'.\n"+
			"Instead, dTz.GetTimeZoneDef().LocationName='%v'\n",
			eTimeZoneDef.GetOriginalLocationName(), dTz.GetOriginalTzName())
	}

	tDto, err := dt.TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral)\n"+
			"t4USCentral='%v'\nError='%v'\n",
			t4USCentral.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	expectedDt, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from tDto.GetDateTime(TZones.US.Central()). "+
			"Error='%v'", err.Error())
	}

	timeComponents := dTz.GetTimeComponents()

	actualDt, err := timeComponents.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from dTz.GetTimeComponents().GetDateTime(TZones.US.Central()).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !tDto.Equal(dTz.GetTimeComponents()) {
		fmt.Printf("Expected dTz.Time (TimeDto) == '%v' Instead, dTz.Time (TimeDto) == '%v'",
			expectedDt.Format(dt.FmtDateTimeYrMDayFmtStr),
			actualDt.Format(dt.FmtDateTimeYrMDayFmtStr))
	}

	if dt.FmtDateTimeYrMDayFmtStr != dTz.GetDateTimeFmt() {
		fmt.Printf("Expected dTz.GetDateTimeFmt()='%v' Instead, dTz.GetDateTimeFmt()='%v' ",
			dt.FmtDateTimeYrMDayFmtStr, dTz.GetDateTimeFmt())
	}

}

func (mt mainTest) mainTest055() {

	ePrefix := "mainTest055()"

	mt.mainPrintHdr(ePrefix , "-")

	// t1EdtStr :=  "06/20/2019 09:58:32 -0400 EDT"
	// t1EstStr :=  "12/20/2019 09:58:32 -0500 EST"
	// t2CdtStr :=  "06/20/2019 09:58:32 -0500 CDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0600 CST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0600 MDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0700 MST"
	t2PdtStr :=  "06/20/2019 09:58:32 -0700 PDT"
	// t2PstStr :=  "12/20/2019 09:58:32 -0800 PST"

	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	timeStr := t2PdtStr
	t1, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	tzDef, err := dt.TimeZoneDefinition{}.NewDateTime(t1)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(t1)\n" +
			"t1='%v'\nError='%v'\n", t1.Format(fmtStr), err.Error())
		return
	}

	ex.PrintOutTimeZoneDefFields(tzDef)

}



func (mt mainTest) mainTest054() {

	ePrefix := "mainTest054()"
	lineLen := 65
	mt.mainPrintHdr(ePrefix , "-")

	// t1EdtStr :=  "06/20/2019 09:58:32 -0400 EDT"
	// t1EstStr :=  "12/20/2019 09:58:32 -0500 EST"
	// t2CdtStr :=  "06/20/2019 09:58:32 -0500 CDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0600 CST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0600 MDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0700 MST"
	t2PdtStr :=  "06/20/2019 09:58:32 -0700 PDT"
	// t2PstStr :=  "12/20/2019 09:58:32 -0800 PST"

	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	timeStr := t2PdtStr
	t1, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1 Initial Date Time"},
		lineLen,
		fmtStr)

	t2 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		t1.Location())

	ex.PrintOutDateTimeTimeZoneFields(
		t2,
		[]string{"t2 Date Time"},
		lineLen,
		fmtStr)

	pacificTz := dt.TZones.America.Los_Angeles()

	pacificTzPtr, err := time.LoadLocation(pacificTz)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(pacificTz)\n" +
			"pacificTz='%v'\nError='%v'\n", pacificTz, err.Error())
		return
	}

	t3 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		pacificTzPtr)

	ex.PrintOutDateTimeTimeZoneFields(
		t3,
		[]string{"t3 Date Time"},
		lineLen,
		fmtStr)

}


func (mt mainTest) mainTest053() {

	ePrefix := "mainTest053()"

	mt.mainPrintHdr(ePrefix , "=")

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _ := dt.TimeZoneDto{}.New(tIn, ianaPacificTz, fmtstr)

	fmt.Println(" Time In: ", tzu.TimeIn.GetDateTimeValue().Format(fmtstr))
	fmt.Println("Time Out: ", tzu.TimeOut.GetDateTimeValue().Format(fmtstr))

	expectedZone := "PDT"

	actualZone := tzu.TimeOut.GetOriginalTzAbbreviation()

	if expectedZone != actualZone {
		fmt.Printf("Expected Zone Out='%v'.\n" +
			"Instead, actual Zone Out='%v'\n", expectedZone, actualZone)
		return
	}

	mt.mainPrintHdr("SUCCESS" , "!!!")


}


func (mt mainTest) mainTest052() {

	ePrefix := "mainTest052()"

	mt.mainPrintHdr(ePrefix , "=")

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned from TimeZoneDto{}.NewStartEndTimes(t1, TzUsPacific ).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tzu1.IsValid()

	if err != nil {
		fmt.Printf("'tzu1' is INVALID!\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("\nPassed IsValid() Check #1!\n")

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	t12Dur := t2.Sub(t1)

	tdurDto, err := dt.TimeDurationDto{}.NewStartEndTimes(
		t1,
		t2,
		dt.TDurCalcType(0).StdYearMth(),
		dt.TZones.US.Central(),
		dt.TCalcMode.LocalTimeZone(),
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}.NewStartEndTimesCalcTz()\n"+
			"Error='%v' ", err.Error())
		return
	}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddTimeDurationDto(tdurDto)
// Error expected here!
	if err != nil {
		fmt.Printf("Error returned by tzu2.AddTimeDurationDto(tdurDto).\n"+
			"Error='%v' ", err.Error())
		return
	}

	tzu1OutStr := tzu1.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t1OutStr != tzu1OutStr {
		fmt.Printf("Error: Expected Time1 TimeIn='%v'.\n" +
			"Instead Time1 TimeIn='%v'\n",
			t1OutStr, tzu1OutStr)
		return
	}

	tzu2OutStr := tzu2.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t2OutStr != tzu2OutStr {
		fmt.Printf("Error: Expected after duration tzu2TimeIn='%v'.\n" +
			"Instead, tzu2TimeIn='%v'\n",
			t2OutStr, tzu2OutStr)
		return
	}

	actualDur := tzu2.TimeIn.Sub(tzu1.TimeIn)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeIn.Sub(tzu1.TimeIn)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeOut.Sub(tzu1.TimeOut)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeOut.Sub(tzu1.TimeOut)='%v'.\n" +
			"Instead, duration='%v'\n", t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeUTC.Sub(tzu1.TimeUTC)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeUTC.Sub(tzu1.TimeUTC)='%v'.\n" +
			"Instead, duration='%v'\n", t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeLocal.Sub(tzu1.TimeLocal)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeLocal.Sub(tzu1.TimeLocal)='%v'.\n" +
			"Instead, duration='%v'\n", t12Dur, actualDur)
		return
	}

	actualTimeOutLoc := tzu1.TimeOut.GetOriginalTzName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu1.TimeOutLoc='%v'.\n" +
			"Instead, tzu1.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	actualTimeOutLoc = tzu2.TimeOut.GetOriginalTzName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu2.TimeOutLoc.String()='%v'.\n" +
			"Instead, tzu2.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	mt.mainPrintHdr("SUCCESS" , "!")

}

func (mt mainTest) mainTest051() {
	ePrefix := "mainTest051()"
	lineLen := 65
	mt.mainPrintHdr(ePrefix , "-")

	timeZone := dt.TZones.Local()
	// t2PdtStr :=  "06/20/2019 09:58:32.000000000 -0700 PDT"
	utcLoc, err := time.LoadLocation(timeZone)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.UTC())\n" +
			"Error='%v'\n", err.Error())
	}

	txUtc := time.Date(
		2019,
		time.Month(12),
		15,
		11,
		0,
		0,
		0,
		utcLoc)

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"
	timeStr := txUtc.Format(fmtStr)
	fmt.Println()
	fmt.Println("#1   timeStr: ", timeStr)
	fmt.Println("#1 time zone: ", timeZone)
	fmt.Println()

		t1, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1 Initial Date Time"},
		lineLen,
		fmtStr)

	_, err = time.LoadLocation(t1.Location().String())

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(t1.Location().String())\n" +
			"t1.Location().String()='%v'\n" +
			"Error='%v'\n", t1.Location().String(), err.Error())
	}

	tzMech := dt.TimeZoneMechanics{}

	var tzSpec dt.TimeZoneSpecification

	tzSpec,
	err = tzMech.GetConvertibleTimeZoneFromDateTime(
		t1,
		dt.TzConvertType.Absolute(),
		ePrefix,
		ePrefix)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by tzMech.GetConvertibleTimeZoneFromDateTime(t1, ePrefix)\n" +
			"t1='%v'\n" +
			"Error='%v'\n", t1.Format(fmtStr), err.Error())
		return
	}

	fmt.Println()
	fmt.Printf("        ianaTimeZoneName: %v\n", tzSpec.GetLocationName())
	fmt.Printf("         ianaLocationPtr: %v\n", tzSpec.GetLocationPointer().String())
	fmt.Printf("                 tzClass: %v\n\n", tzSpec.GetTimeZoneClass().String())

	fmt.Printf(ePrefix +
		"\nSuccess!\n" +
		"time zone pointer name = '%v'\n",t1.Location().String())

}

func (mt mainTest) mainTest050() {
/*
	 =========================================
	      Time Zone Fields From time.Time
	 =========================================
	    t1 Initial Date Time:  12/30/2019 09:00:00.000000000 +0700 +07
	               Zone Name:  +07
	     Abbreviation Lookup:  +07+0700
	               Zone Sign:  +
	       Zone Offset Hours:  7
	     Zone Offset Minutes:  0
	     Zone Offset Seconds:  0
	    Total Offset Seconds:  25200
	              UTC Offset:  UTC+0700
	        Location Pointer:  Asia/Ho_Chi_Minh
	           Location Name:  Asia/Ho_Chi_Minh
	 =========================================
 */
	ePrefix := "mainTest050()"
	lineLen := 65
	mt.mainPrintHdr(ePrefix , "-")

	// "Asia/Ho_Chi_Minh"
// timeZoneName := "Asia/Vladivostok"
timeZoneName := "Asia/Ho_Chi_Minh"

tzLocPtr, err := time.LoadLocation(timeZoneName)

if err != nil {
	fmt.Printf(ePrefix +
		"Error returned by time.LoadLocation(timeZoneName)\n" +
		"timeZoneName='%v'\n" +
		"Error='%v'\n", timeZoneName, err.Error())

	return

}

	t1 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		tzLocPtr)

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1 Initial Date Time"},
		lineLen,
		fmtStr)

}

// Prints Text Title Lines to the Console
func (mt mainTest) mainPrintHdr(textToPrint string, repeatStr string) {


	lenTextToPrint := len(textToPrint)

	lenExtra := 5

	lenBar := (lenExtra * 2) + lenTextToPrint

	blankMargin := strings.Repeat(" ", lenExtra)

	title := blankMargin + textToPrint + blankMargin
	bar := strings.Repeat(repeatStr, lenBar)

	fmt.Println(bar)
	fmt.Println(title)
	fmt.Println(bar)
	fmt.Println()

}
