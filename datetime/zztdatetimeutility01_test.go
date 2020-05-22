package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestDTimeUtility_GregorianDateToJulianDayNo_01(t *testing.T) {


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

	dtUtil := DTimeUtility{}

	expectedJDN := int64(2458989)

	gregorianDateUtc,
	julianDayNumber,
	err := dtUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestDTimeUtility_GregorianDateToJulianDayNo_01 ")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GregorianDateToJulianDayNo()\n" +
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

func TestDTimeUtility_GregorianDateToJulianDayNo_02(t *testing.T) {

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

	dtUtil := DTimeUtility{}

	expectedJDN := int64(2456293)

	gregorianDateUtc,
	julianDayNumber,
	err := dtUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestDTimeUtility_GregorianDateToJulianDayNo_02 ")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GregorianDateToJulianDayNo()\n" +
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

func TestDTimeUtility_GregorianDateToJulianDayNo_03(t *testing.T) {

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

	dtUtil := DTimeUtility{}

	expectedJDN := int64(2458989)

	gregorianDateUtc,
	julianDayNumber,
	err := dtUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestDTimeUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GregorianDateToJulianDayNo()\n" +
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

func TestDTimeUtility_GregorianDateToJulianDayNo_04(t *testing.T) {

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

	dtUtil := DTimeUtility{}

	expectedJDN := int64(2456294)

	gregorianDateUtc,
	julianDayNumber,
	err := dtUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestDTimeUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GregorianDateToJulianDayNo()\n" +
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

func TestDTimeUtility_GregorianDateToJulianDayNo_05(t *testing.T) {

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

	dtUtil := DTimeUtility{}

	expectedJDN := int64(2456294)

	gregorianDateUtc,
	julianDayNumber,
	err := dtUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestDTimeUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GregorianDateToJulianDayNo()\n" +
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

func TestDTimeUtility_GregorianDateToJulianDayNo_06(t *testing.T) {

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

	dtUtil := DTimeUtility{}

	expectedJDN := int64(2456294)

	gregorianDateUtc,
	julianDayNumber,
	err := dtUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestDTimeUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GregorianDateToJulianDayNo()\n" +
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

func TestDTimeUtility_GregorianDateToJulianDayNo_07(t *testing.T) {

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

	dtUtil := DTimeUtility{}

	// 1848-09-07 03:12:15.0000 = 2396277.63351
	expectedJDN := int64(2396277)

	gregorianDateUtc,
	julianDayNumber,
	err := dtUtil.GregorianDateToJulianDayNo(
		testDate,
		"TestDTimeUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GregorianDateToJulianDayNo()\n" +
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

func TestDTimeUtility_GregorianDateToJulianDate_01(t *testing.T) {

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

	dtUtil := DTimeUtility{}

	// 1848-09-07 03:12:15.0000 = 2396277.63351
	expectedJulianDateStr := "2396277.63351"
	digitsAfterDecimal := 5

	var gregorianDateUtc time.Time
	var julianDate float64

	gregorianDateUtc,
	julianDate,
	err = dtUtil.GregorianDateToJulianDate(
		testDate,
		digitsAfterDecimal,
		"TestDTimeUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GregorianDateToJulianDayNo()\n" +
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

func TestDTimeUtility_GregorianDateToJulianDate_02(t *testing.T) {

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

	dtUtil := DTimeUtility{}
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
	err = dtUtil.GregorianDateToJulianDate(
		testDate,
		digitsAfterDecimal,
		"TestDTimeUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GregorianDateToJulianDayNo()\n" +
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

func TestDTimeUtility_GregorianDateToJulianDate_03(t *testing.T) {

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

	dtUtil := DTimeUtility{}
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
	err = dtUtil.GregorianDateToJulianDate(
		testDate,
		digitsAfterDecimal,
		"TestDTimeUtility_GregorianDateToJulianDayNo_03 ")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GregorianDateToJulianDayNo()\n" +
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

