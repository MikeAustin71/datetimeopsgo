package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestCalendarUtility_GregorianDateToJulianDayNo_01(t *testing.T) {


	testDate := time.Date(
		2020,
		5,
		19,
		23,
		5,
		37,
		0,
		time.UTC)

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	calUtil := CalendarUtility{}

	expectedJDN := int64(2458989)

	gregorianDateUtc,
	julianDayNumber,
	err := calUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestCalendarUtility_GregorianDateToJulianDayNo_01 ")

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	if !gregorianDateUtc.Equal(testDate) {
		t.Errorf("Error: Expected gregorianDateUtc='%v'\n" +
			"Instead, gregorianDateUtc='%v'\n",
			testDate.Format(dateTimeFormat),
			gregorianDateUtc.Format(dateTimeFormat))
		return
	}

	if julianDayNumber != expectedJDN {
		t.Errorf("Error: Expected Julian Day Number='%v'\n" +
			"Instead, Julian Day Number='%v'\n",
			expectedJDN, julianDayNumber)
	}
}

func TestCalendarUtility_GregorianDateToJulianDayNo_02(t *testing.T) {

	testDate := time.Date(
		2013,
		1,
		1,
		0,
		30,
		0,
		0,
		time.UTC)

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	calUtil := CalendarUtility{}

	expectedJDN := int64(2456293)

	gregorianDateUtc,
	julianDayNumber,
	err := calUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestCalendarUtility_GregorianDateToJulianDayNo_02 ")

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	if !gregorianDateUtc.Equal(testDate) {
		t.Errorf("Error: Expected gregorianDateUtc='%v'\n" +
			"Instead, gregorianDateUtc='%v'\n",
			testDate.Format(dateTimeFormat),
			gregorianDateUtc.Format(dateTimeFormat))
		return
	}

	if julianDayNumber != expectedJDN {
		t.Errorf("Error: Expected Julian Day Number='%v'\n" +
			"Instead, Julian Day Number='%v'\n",
			expectedJDN, julianDayNumber)
	}
}

func TestCalendarUtility_GregorianDateToJulianDayNo_03(t *testing.T) {

	testDate := time.Date(
		2020,
		5,
		19,
		12,
		0,
		0,
		0,
		time.UTC)

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	calUtil := CalendarUtility{}

	expectedJDN := int64(2458989)

	gregorianDateUtc,
	julianDayNumber,
	err := calUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestCalendarUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	if !gregorianDateUtc.Equal(testDate) {
		t.Errorf("Error: Expected gregorianDateUtc='%v'\n" +
			"Instead, gregorianDateUtc='%v'\n",
			testDate.Format(dateTimeFormat),
			gregorianDateUtc.Format(dateTimeFormat))
		return
	}

	if julianDayNumber != expectedJDN {
		t.Errorf("Error: Expected Julian Day Number='%v'\n" +
			"Instead, Julian Day Number='%v'\n",
			expectedJDN, julianDayNumber)
	}
}

func TestCalendarUtility_GregorianDateToJulianDayNo_04(t *testing.T) {

	testDate := time.Date(
		2013,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC)

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	calUtil := CalendarUtility{}

	expectedJDN := int64(2456294)

	gregorianDateUtc,
	julianDayNumber,
	err := calUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestCalendarUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	if !gregorianDateUtc.Equal(testDate) {
		t.Errorf("Error: Expected gregorianDateUtc='%v'\n" +
			"Instead, gregorianDateUtc='%v'\n",
			testDate.Format(dateTimeFormat),
			gregorianDateUtc.Format(dateTimeFormat))
		return
	}

	if julianDayNumber != expectedJDN {
		t.Errorf("Error: Expected Julian Day Number='%v'\n" +
			"Instead, Julian Day Number='%v'\n",
			expectedJDN, julianDayNumber)
	}
}

func TestCalendarUtility_GregorianDateToJulianDayNo_05(t *testing.T) {

	baseDate := time.Date(
		2013,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC)

	var err error
	var centralTz *time.Location

	centralTz, err = time.LoadLocation(
		TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error: time.LoadLocation(TZones.America.Chicago())\n" +
			"returned an error.\nError='%v'\n",
			err.Error())
		return
	}

	testDate := baseDate.In(centralTz)

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	calUtil := CalendarUtility{}

	expectedJDN := int64(2456294)

	gregorianDateUtc,
	julianDayNumber,
	err := calUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestCalendarUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	if !gregorianDateUtc.Equal(baseDate) {
		t.Errorf("Error: Expected gregorianDateUtc='%v'\n" +
			"Instead, gregorianDateUtc='%v'\n",
			baseDate.Format(dateTimeFormat),
			gregorianDateUtc.Format(dateTimeFormat))
		return
	}

	if julianDayNumber != expectedJDN {
		t.Errorf("Error: Expected Julian Day Number='%v'\n" +
			"Instead, Julian Day Number='%v'\n",
			expectedJDN, julianDayNumber)
	}
}

func TestCalendarUtility_GregorianDateToJulianDayNo_06(t *testing.T) {

	baseDate := time.Date(
		2013,
		1,
		1,
		12,
		0,
		0,
		1,
		time.UTC)

	var err error
	var centralTz *time.Location

	centralTz, err = time.LoadLocation(
		TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error: time.LoadLocation(TZones.America.Chicago())\n" +
			"returned an error.\nError='%v'\n",
			err.Error())
		return
	}

	testDate := baseDate.In(centralTz)

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	calUtil := CalendarUtility{}

	expectedJDN := int64(2456294)

	gregorianDateUtc,
	julianDayNumber,
	err := calUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestCalendarUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	if !gregorianDateUtc.Equal(baseDate) {
		t.Errorf("Error: Expected gregorianDateUtc='%v'\n" +
			"Instead, gregorianDateUtc='%v'\n",
			baseDate.Format(dateTimeFormat),
			gregorianDateUtc.Format(dateTimeFormat))
		return
	}

	if julianDayNumber != expectedJDN {
		t.Errorf("Error: Expected Julian Day Number='%v'\n" +
			"Instead, Julian Day Number='%v'\n",
			expectedJDN, julianDayNumber)
	}
}

func TestCalendarUtility_GregorianDateToJulianDayNo_07(t *testing.T) {

	testDate := time.Date(
		1848,
		9,
		7,
		3,
		12,
		15,
		0,
		time.UTC)

	var err error

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	calUtil := CalendarUtility{}

	// 1848-09-07 03:12:15.0000 = 2396277.63351
	expectedJDN := int64(2396277)

	gregorianDateUtc,
	julianDayNumber,
	err := calUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestCalendarUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	if !gregorianDateUtc.Equal(testDate) {
		t.Errorf("Error: Expected gregorianDateUtc='%v'\n" +
			"Instead, gregorianDateUtc='%v'\n",
			testDate.Format(dateTimeFormat),
			gregorianDateUtc.Format(dateTimeFormat))
		return
	}

	if julianDayNumber != expectedJDN {
		t.Errorf("Error: Expected Julian Day Number='%v'\n" +
			"Instead, Julian Day Number='%v'\n",
			expectedJDN, julianDayNumber)
	}
}

func TestCalendarUtility_GregorianDateToJulianDayNoTime_01(t *testing.T) {

	testDate := time.Date(
		1848,
		9,
		7,
		3,
		12,
		15,
		0,
		time.UTC)

	var err error

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	calUtil := CalendarUtility{}

	// 1848-09-07 03:12:15.0000 = 2396277.63351
	// Accurate to 5 decimal places
	expectedJulianDateStr := "2396277.63351"


	var julianDayNoDto JulianDayNoDto
	var gregorianDateUtc time.Time

	gregorianDateUtc,
	julianDayNoDto,
	err = calUtil.GregorianDateToJulianDayNoTime(
		testDate,
		"TestCalendarUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	if !gregorianDateUtc.Equal(testDate) {
		t.Errorf("Error: Expected gregorianDateUtc='%v'\n" +
			"Instead, gregorianDateUtc='%v'\n",
			testDate.Format(dateTimeFormat),
			gregorianDateUtc.Format(dateTimeFormat))
		return
	}

	julianDateStr, _, _ := julianDayNoDto.GetJulianDayNoTimeStr(5)

	if expectedJulianDateStr != julianDateStr {
		t.Errorf("\n" +
			"Error: Expected Julian Date= '%v'\n" +
			"Instead, Julian Date=        '%v'\n",
			expectedJulianDateStr, expectedJulianDateStr)
	}
}

func TestCalendarUtility_GregorianDateToJulianDayNoTime_02(t *testing.T) {

	testDate := time.Date(
		1776,
		7,
		4,
		12,
		1,
		0,
		0,
		time.UTC)

	var err error

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	calUtil := CalendarUtility{}
	var expectedJulianDate float64
	// 1776-07-04 12:01:0.0000 = 2369916.00069
	expectedJulianDate = 2369916.00069
	expectedJulianDateStr := fmt.Sprintf("%.5f",
		expectedJulianDate)

	digitsAfterDecimal := 5

	var gregorianDateUtc time.Time
	var julianDayNoDto JulianDayNoDto

	gregorianDateUtc,
		julianDayNoDto,
	err = calUtil.GregorianDateToJulianDayNoTime(
		testDate,
		"TestCalendarUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	if !gregorianDateUtc.Equal(testDate) {
		t.Errorf("Error: Expected gregorianDateUtc='%v'\n" +
			"Instead, gregorianDateUtc='%v'\n",
			testDate.Format(dateTimeFormat),
			gregorianDateUtc.Format(dateTimeFormat))
		return
	}

	julianDateStr,_ , _ := julianDayNoDto.GetJulianDayNoTimeStr(digitsAfterDecimal)

	if expectedJulianDateStr != julianDateStr {
		t.Errorf("Error: Expected Julian Date='%v'\n" +
			"Instead, Julian Date='%v'\n",
			expectedJulianDateStr, expectedJulianDateStr)
	}
}

func TestCalendarUtility_GregorianDateToJulianDayNoTime_03(t *testing.T) {

	testDate := time.Date(
		1776,
		7,
		4,
		11,
		59,
		59,
		0,
		time.UTC)

	var err error

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	calUtil := CalendarUtility{}
	var expectedJulianDate float64
	// 1776-07-04 11:59:59.0000 = 2369915.99999
	expectedJulianDate = 2369915.99999
	expectedJulianDateStr := fmt.Sprintf("%.5f",
		expectedJulianDate)

	digitsAfterDecimal := 5

	var gregorianDateUtc time.Time
	var julianDayNoDto JulianDayNoDto

	gregorianDateUtc,
		julianDayNoDto,
	err = calUtil.GregorianDateToJulianDayNoTime(
		testDate,
		"TestCalendarUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNo()\n" +
			"testDate='%v'\n" +
			"Error='%v'\n",
			testDate.Format(dateTimeFormat),
			err.Error())
		return
	}

	if !gregorianDateUtc.Equal(testDate) {
		t.Errorf("Error: Expected gregorianDateUtc='%v'\n" +
			"Instead, gregorianDateUtc='%v'\n",
			testDate.Format(dateTimeFormat),
			gregorianDateUtc.Format(dateTimeFormat))
		return
	}

	julianDateStr, _, _ := julianDayNoDto.GetJulianDayNoTimeStr(digitsAfterDecimal)

	if expectedJulianDateStr != julianDateStr {
		t.Errorf("Error: \n" +
			"Expected Julian Date= '%v'\n" +
			"Instead, Julian Date= '%v'\n",
			expectedJulianDateStr, expectedJulianDateStr)
	}
}

func TestCalendarUtility_JulianDayNoTimeToGregorianDateTime_01(t *testing.T) {

	expectedDateTime := time.Date(
		2013,
		1,
		1,
		0,
		30,
		0,
		0,
		time.UTC)

	var julianDateTime float64

	julianDateTime = 2456293.520833

	julianDayNoDto, err := JulianDayNoDto{}.NewFromFloat64(julianDateTime)

	if err != nil {
		t.Errorf("Error returned by JulianDayNoDto{}.NewFromFloat64(julianDateTime)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	calUtil := CalendarUtility{}
	var gregorianDateTime time.Time

	gregorianDateTime,
	err = calUtil.JulianDayNoTimeToGregorianCalendar(
		julianDayNoDto,
		"TestCalendarUtility_JulianDayNoTimeToGregorianDateTime_01")

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToGregorianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedDateTime.Equal(gregorianDateTime) {
		t.Errorf("Error: Expected " +
			"gregorianDateTime='%v'\n" +
			"Instead, gregorianDateTime='%v'\n",
			expectedDateTime.Format(FmtDateTimeYrMDayFmtStr),
			gregorianDateTime.Format(FmtDateTimeYrMDayFmtStr))
	}
}

func TestCalendarUtility_JulianDayNoTimeToGregorianDateTime_02(t *testing.T) {

	expectedDateTime := time.Date(
		-4713,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)

	var julianDayNoTimeFloat64 float64

	julianDayNoTimeFloat64 = 0.000000

	digitsAfterDecimal := 6

	expectedJulianDayNoStr := fmt.Sprintf("%.6F",
		julianDayNoTimeFloat64)

	var julianDayNoDto JulianDayNoDto
	var err error

	julianDayNoDto, err =
		JulianDayNoDto{}.NewFromFloat64(julianDayNoTimeFloat64)

	if err != nil {
		t.Errorf("Error returned by JulianDayNoDto{}.NewFromFloat64(julianDayNoTimeFloat64)")
		return
	}

	actualJulianDayNoTimeStr, _, _ :=
		julianDayNoDto.GetJulianDayNoTimeStr(digitsAfterDecimal)

	if expectedJulianDayNoStr != actualJulianDayNoTimeStr {
		t.Errorf("Error:\n" +
			"Expected Julian Day No String='%v'\n" +
			"Instead, Julian Day No String='%v'\n",
			expectedJulianDayNoStr, actualJulianDayNoTimeStr)
		return
	}

	calUtil := CalendarUtility{}

	gregorianDateTime,
	err := calUtil.JulianDayNoTimeToGregorianCalendar(
		julianDayNoDto,
		"TestCalendarUtility_JulianDayNoTimeToGregorianDateTime_02")

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToGregorianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedDateTime.Equal(gregorianDateTime) {
		t.Errorf("Error:\n" +
			"Expected gregorianDateTime='%v'\n" +
			"Instead, gregorianDateTime='%v'\n",
			expectedDateTime.Format(FmtDateTimeYrMDayFmtStr),
			gregorianDateTime.Format(FmtDateTimeYrMDayFmtStr))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_01(t *testing.T) {

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_01() "

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
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			testDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_02(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_02() "

	testDate := time.Date(
		-500,
		2,
		28,
		13,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-500,
			3,
			5,
			13,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			testDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_03(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_03() "

	gregorianDate := time.Date(
		-500,
		3,
		1,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-500,
			3,
			6,
			10,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_04(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_04() "

	gregorianDate := time.Date(
		-300,
		2,
		27,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-300,
			3,
			3,
			10,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_05(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_05() "

	gregorianDate := time.Date(
		-300,
		2,
		28,
		15,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-300,
			3,
			4,
			15,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_06(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_06() "

	gregorianDate := time.Date(
		-300,
		3,
		1,
		15,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-300,
			3,
			5,
			15,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_07(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_07() "

	gregorianDate := time.Date(
		-200,
		2,
		27,
		15,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-200,
			3,
			2,
			15,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_08(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_08() "

	gregorianDate := time.Date(
		-200,
		2,
		28,
		15,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-200,
			3,
			3,
			15,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_09(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_09() "

	gregorianDate := time.Date(
		-200,
		3,
		1,
		15,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-200,
			3,
			4,
			15,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_10(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_10() "

	gregorianDate := time.Date(
		-100,
		2,
		27,
		15,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-100,
			3,
			1,
			15,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_11(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_11() "

	gregorianDate := time.Date(
		-100,
		2,
		28,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-100,
			3,
			2,
			10,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_12(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_12() "

	gregorianDate := time.Date(
		-100,
		3,
		1,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			-100,
			3,
			3,
			10,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_13(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_13() "

	gregorianDate := time.Date(
		100,
		2,
		27,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			100,
			2,
			29,
			10,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}
	var julianDayNoTimeDto JulianDayNoDto

	_,
	julianDayNoTimeDto,
	err =
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_14a(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_14a() "

	gregorianDate := time.Date(
		100,
		2,
		28,
		14,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			100,
			3,
			1,
			14,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_14b(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_14b() "

	gregorianDate := time.Date(
		100,
		2,
		28,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			100,
			3,
			2,
			10,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_15(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_15() "

	gregorianDate := time.Date(
		100,
		3,
		1,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			100,
			3,
			2,
			10,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_16(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_16() "

	gregorianDate := time.Date(
		100,
		2,
		28,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			100,
			3,
			1,
			10,
			0,
			0,
			0,
			time.UTC)

	var err error
	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_17(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_17() "

	gregorianDate := time.Date(
		100,
		3,
		1,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			100,
			3,
			2,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_18(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_18() "

	gregorianDate := time.Date(
		200,
		2,
		27,
		16,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			200,
			2,
			28,
			16,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_19(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_19() "

	gregorianDate := time.Date(
		200,
		2,
		28,
		16,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			200,
			2,
			29,
			16,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_20(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_20() "

	gregorianDate := time.Date(
		200,
		3,
		1,
		16,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			200,
			3,
			1,
			16,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_21(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_21() "

	gregorianDate := time.Date(
		300,
		2,
		28,
		16,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			300,
			2,
			28,
			16,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_22(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_22() "

	gregorianDate := time.Date(
		300,
		3,
		1,
		16,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			300,
			2,
			29,
			16,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_23(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_23() "

	gregorianDate := time.Date(
		300,
		3,
		2,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			300,
			3,
			1,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_24(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_24() "

	gregorianDate := time.Date(
		500,
		3,
		1,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			500,
			2,
			28,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_25(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_25() "

	gregorianDate := time.Date(
		500,
		3,
		2,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			500,
			2,
			29,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_26(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_26() "

	gregorianDate := time.Date(
		500,
		3,
		3,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			500,
			3,
			1,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_27(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_27() "

	gregorianDate := time.Date(
		600,
		3,
		2,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			600,
			2,
			28,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_28(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_28() "

	gregorianDate := time.Date(
		600,
		3,
		3,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			600,
			2,
			29,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_29(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_29() "

	gregorianDate := time.Date(
		600,
		3,
		4,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			600,
			3,
			1,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_30(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_30() "

	gregorianDate := time.Date(
		700,
		3,
		3,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			700,
			2,
			28,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_31(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_31() "

	gregorianDate := time.Date(
		700,
		3,
		4,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			700,
			2,
			29,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_32(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_32() "

	gregorianDate := time.Date(
		700,
		3,
		5,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			700,
			3,
			1,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_33(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_33() "

	gregorianDate := time.Date(
		900,
		3,
		4,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			900,
			2,
			28,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_34(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_34() "

	gregorianDate := time.Date(
		900,
		3,
		5,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			900,
			2,
			29,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_35(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_35() "

	gregorianDate := time.Date(
		900,
		3,
		6,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			900,
			3,
			1,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_36(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_36() "

	gregorianDate := time.Date(
		1000,
		3,
		5,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1000,
			2,
			28,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_37(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_37() "

	gregorianDate := time.Date(
		1000,
		3,
		6,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1000,
			2,
			29,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_38(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_38() "

	gregorianDate := time.Date(
		1000,
		3,
		7,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1000,
			3,
			1,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_39(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_39() "

	gregorianDate := time.Date(
		1100,
		3,
		6,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1100,
			2,
			28,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_40(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_40() "

	gregorianDate := time.Date(
		1100,
		3,
		7,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1100,
			2,
			29,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_41(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_41() "

	gregorianDate := time.Date(
		1100,
		3,
		8,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1100,
			3,
			1,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_42(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_42() "

	gregorianDate := time.Date(
		1300,
		3,
		7,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1300,
			2,
			28,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_43(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_43() "

	gregorianDate := time.Date(
		1300,
		3,
		8,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1300,
			2,
			29,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_44(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_44() "

	gregorianDate := time.Date(
		1300,
		3,
		9,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1300,
			3,
			1,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_45(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_45() "

	gregorianDate := time.Date(
		1400,
		3,
		8,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1400,
			2,
			28,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_46(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_46() "

	gregorianDate := time.Date(
		1400,
		3,
		9,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1400,
			2,
			29,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_47(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_47() "

	gregorianDate := time.Date(
		1400,
		3,
		10,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1400,
			3,
			1,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_48(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_48() "

	gregorianDate := time.Date(
		1500,
		3,
		9,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1500,
			2,
			28,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_49(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_49() "

	gregorianDate := time.Date(
		1500,
		3,
		10,
		10,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1500,
			2,
			29,
			10,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

func TestCalendarUtility_JulianDayNoTimeToJulianDateTime_50(t *testing.T) {

	// https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars

	ePrefix := "TestCalendarUtility_JulianDayNoTimeToJulianDateTime_50() "

	gregorianDate := time.Date(
		1500,
		3,
		11,
		15,
		0,
		0,
		0,
		time.UTC)

	expectedJulianDateTime :=
		time.Date(
			1500,
			3,
			1,
			15,
			0,
			0,
			0,
			time.UTC)

	calUtil := CalendarUtility{}

	_,
	julianDayNoTimeDto,
	err :=
		calUtil.GregorianDateToJulianDayNoTime(
			gregorianDate,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.GregorianDateToJulianDayNoTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var julianDateTime time.Time

	dateTimeFormat := FmtDateTimeYrMDayFmtStr

	julianDateTime,
		err =
		calUtil.JulianDayNoTimeToJulianCalendar(
			julianDayNoTimeDto,
			ePrefix)

	if err != nil {
		t.Errorf("Error returned by calUtil.JulianDayNoTimeToJulianCalendar()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !expectedJulianDateTime.Equal(julianDateTime) {
		t.Errorf("Error:\n" +
			"Expected Julian Date Time: %v\n" +
			"Instead, Julian Date Time: %v\n",
			expectedJulianDateTime.Format(dateTimeFormat),
			julianDateTime.Format(dateTimeFormat))
	}
}

