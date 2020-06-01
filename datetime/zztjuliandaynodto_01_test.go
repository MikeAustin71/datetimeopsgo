package datetime

import (
	"testing"
	"time"
)

func TestJulianDayNoDto_New(t *testing.T) {
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



}
