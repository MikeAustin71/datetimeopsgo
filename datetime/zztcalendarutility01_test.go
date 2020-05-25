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
	expectedJulianDateStr := "2396277.63351"
	digitsAfterDecimal := 5

	var gregorianDateUtc time.Time
	var julianDate float64

	gregorianDateUtc,
	julianDate,
	err = calUtil.GregorianDateToJulianDayNoTime(
		testDate,
		digitsAfterDecimal,
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

	julianDateStr := fmt.Sprintf("%.5f",
		julianDate)

	if expectedJulianDateStr != julianDateStr {
		t.Errorf("Error: Expected Julian Date='%v'\n" +
			"Instead, Julian Date='%v'\n",
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
	var julianDate float64

	gregorianDateUtc,
	julianDate,
	err = calUtil.GregorianDateToJulianDayNoTime(
		testDate,
		digitsAfterDecimal,
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

	julianDateStr := fmt.Sprintf("%.5f",
		julianDate)

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
	var julianDate float64

	gregorianDateUtc,
	julianDate,
	err = calUtil.GregorianDateToJulianDayNoTime(
		testDate,
		digitsAfterDecimal,
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

	julianDateStr := fmt.Sprintf("%.5f",
		julianDate)

	if expectedJulianDateStr != julianDateStr {
		t.Errorf("Error: Expected Julian Date='%v'\n" +
			"Instead, Julian Date='%v'\n",
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

	digitsAfterDecimal := 6

	calUtil := CalendarUtility{}

	gregorianDateTime,
	err := calUtil.JulianDayNoTimeToGregorianCalendar(
		julianDateTime,
		digitsAfterDecimal,
		"")

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

	var julianDateTime float64

	julianDateTime = 0.000000

	digitsAfterDecimal := 6

	calUtil := CalendarUtility{}

	gregorianDateTime,
	err := calUtil.JulianDayNoTimeToGregorianCalendar(
		julianDateTime,
		digitsAfterDecimal,
		"")

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
