package datetime

import (
	"testing"
	"time"
)

func TestJulianDayNoDto_New_01(t *testing.T) {
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html

	testDate := time.Date(
		-4713,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)

	// to 9-digits
	expectedJulianDayNoTimeStr :=
	"0.000000000"

	gregorianDateTimeUtc,
	julianDayNoDto,
	err := JulianDayNoDto{}.NewFromGregorianDate(testDate)

	if err != nil {
		t.Errorf("Error returned by JulianDayNoDto{}.NewFromGregorianDate(testDate)\n" +
			"testDate='%v'\nError='%v'\n",
			testDate.Format(FmtDateTimeYrMDayFmtStr),
			err.Error())
		return
	}

	if !testDate.Equal(gregorianDateTimeUtc) {
		t.Errorf("\n" +
			"Error: Expected gregorianDateTimeUtc= '%v'.\n" +
			"Instead, gregorianDateTimeUtc=        '%v'\n",
			testDate.Format(FmtDateTimeYrMDayFmtStr),
			gregorianDateTimeUtc.Format(FmtDateTimeYrMDayFmtStr))
		return
	}

	julianDayNoTimeStr,
	_,
	_	:= julianDayNoDto.GetJulianDayNoTimeStr(9)

	if expectedJulianDayNoTimeStr != julianDayNoTimeStr {
		t.Errorf("Error:\n" +
			"Expected Julian Day Number/Time= '%v'\n" +
			"Instead, Juilan Day Number/Time= '%v'\n",
			expectedJulianDayNoTimeStr, julianDayNoTimeStr)
	}

}
