package datetime

import (
	"testing"
	"time"
)

func TestJulianDayNoDto_New_01(t *testing.T) {
	// https://keisan.casio.com/exec/system/1227779487#!
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html
	// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
	// https://en.wikipedia.org/wiki/Julian_day

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

func TestJulianDayNoDto_New_02(t *testing.T) {
	// https://keisan.casio.com/exec/system/1227779487#!
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html
	// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
	// https://en.wikipedia.org/wiki/Julian_day

	// last day before Gregorian reform
	testDate := time.Date(
		1582,
		10,
		4,
		0,
		0,
		0,
		0,
		time.UTC)

	// to 9-digits
	expectedJulianDayNoTimeStr :=
	"2299149.500000000"

	digitsToRightOfDecimal := 9

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
	_	:= julianDayNoDto.GetJulianDayNoTimeStr(digitsToRightOfDecimal)

	if expectedJulianDayNoTimeStr != julianDayNoTimeStr {
		t.Errorf("Error:\n" +
			"Expected Julian Day Number/Time= '%v'\n" +
			"Instead, Juilan Day Number/Time= '%v'\n",
			expectedJulianDayNoTimeStr, julianDayNoTimeStr)
	}

}

func TestJulianDayNoDto_New_03(t *testing.T) {
	// https://keisan.casio.com/exec/system/1227779487#!
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html
	// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
	// https://en.wikipedia.org/wiki/Julian_day

	// first day of Gregorian reform
	testDate := time.Date(
		1582,
		10,
		15,
		0,
		0,
		0,
		0,
		time.UTC)

	// to 9-digits
	expectedJulianDayNoTimeStr :=
	"2299160.50000000000000000000"

	digitsToRightOfDecimal := 20

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
	_	:= julianDayNoDto.GetJulianDayNoTimeStr(digitsToRightOfDecimal)

	if expectedJulianDayNoTimeStr != julianDayNoTimeStr {
		t.Errorf("Error:\n" +
			"Expected Julian Day Number/Time= '%v'\n" +
			"Instead, Juilan Day Number/Time= '%v'\n",
			expectedJulianDayNoTimeStr, julianDayNoTimeStr)
	}

}

func TestJulianDayNoDto_New_04(t *testing.T) {
	// https://keisan.casio.com/exec/system/1227779487#!
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html
	// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
	// https://en.wikipedia.org/wiki/Julian_day

	// first day of Gregorian reform
	testDate := time.Date(
		1979,
		12,
		31,
		12,
		0,
		0,
		0,
		time.UTC)

	// to 9-digits
	expectedJulianDayNoTimeStr :=
	"2444239.00000000000000000000"

	digitsToRightOfDecimal := 20

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
	_	:= julianDayNoDto.GetJulianDayNoTimeStr(digitsToRightOfDecimal)

	if expectedJulianDayNoTimeStr != julianDayNoTimeStr {
		t.Errorf("Error:\n" +
			"Expected Julian Day Number/Time= '%v'\n" +
			"Instead, Juilan Day Number/Time= '%v'\n",
			expectedJulianDayNoTimeStr, julianDayNoTimeStr)
	}
}

func TestJulianDayNoDto_New_05(t *testing.T) {
	// https://keisan.casio.com/exec/system/1227779487#!
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html
	// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
	// https://en.wikipedia.org/wiki/Julian_day

	// first day of Gregorian reform
	testDate := time.Date(
		1980,
		1,
		1,
		0,
		0,
		0,
		0,
		time.UTC)

	// to 9-digits
	expectedJulianDayNoTimeStr :=
	"2444239.50000000000000000000"

	digitsToRightOfDecimal := 20

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
	_	:= julianDayNoDto.GetJulianDayNoTimeStr(digitsToRightOfDecimal)

	if expectedJulianDayNoTimeStr != julianDayNoTimeStr {
		t.Errorf("Error:\n" +
			"Expected Julian Day Number/Time= '%v'\n" +
			"Instead, Juilan Day Number/Time= '%v'\n",
			expectedJulianDayNoTimeStr, julianDayNoTimeStr)
	}
}

func TestJulianDayNoDto_New_06(t *testing.T) {
	// https://keisan.casio.com/exec/system/1227779487#!
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html
	// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
	// https://en.wikipedia.org/wiki/Julian_day

	// first day of Gregorian reform
	testDate := time.Date(
		1980,
		1,
		1,
		0,
		0,
		0,
		0,
		time.UTC)

	// to 9-digits
	expectedJulianDayNoTimeStr :=
	"2444239.50000000000000000000"

	digitsToRightOfDecimal := 20

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
	_	:= julianDayNoDto.GetJulianDayNoTimeStr(digitsToRightOfDecimal)

	if expectedJulianDayNoTimeStr != julianDayNoTimeStr {
		t.Errorf("Error:\n" +
			"Expected Julian Day Number/Time= '%v'\n" +
			"Instead, Juilan Day Number/Time= '%v'\n",
			expectedJulianDayNoTimeStr, julianDayNoTimeStr)
	}
}

func TestJulianDayNoDto_New_07(t *testing.T) {
	// https://keisan.casio.com/exec/system/1227779487#!
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html
	// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
	// https://en.wikipedia.org/wiki/Julian_day

	// first day of Gregorian reform
	testDate := time.Date(
		1980,
		1,
		1,
		12,
		0,
		0,
		0,
		time.UTC)

	// to 9-digits
	expectedJulianDayNoTimeStr :=
	"2444240.00000000000000000000"

	digitsToRightOfDecimal := 20

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
	_	:= julianDayNoDto.GetJulianDayNoTimeStr(digitsToRightOfDecimal)

	if expectedJulianDayNoTimeStr != julianDayNoTimeStr {
		t.Errorf("Error:\n" +
			"Expected Julian Day Number/Time= '%v'\n" +
			"Instead, Juilan Day Number/Time= '%v'\n",
			expectedJulianDayNoTimeStr, julianDayNoTimeStr)
	}
}

func TestJulianDayNoDto_New_08(t *testing.T) {
	// https://keisan.casio.com/exec/system/1227779487#!
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html
	// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
	// https://en.wikipedia.org/wiki/Julian_day

	// first day of Gregorian reform
	testDate := time.Date(
		1980,
		1,
		2,
		0,
		0,
		0,
		0,
		time.UTC)

	// to 9-digits
	expectedJulianDayNoTimeStr :=
	"2444240.50000000000000000000"

	digitsToRightOfDecimal := 20

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
	_	:= julianDayNoDto.GetJulianDayNoTimeStr(digitsToRightOfDecimal)

	if expectedJulianDayNoTimeStr != julianDayNoTimeStr {
		t.Errorf("Error:\n" +
			"Expected Julian Day Number/Time= '%v'\n" +
			"Instead, Juilan Day Number/Time= '%v'\n",
			expectedJulianDayNoTimeStr, julianDayNoTimeStr)
	}
}


func TestJulianDayNoDto_New_10(t *testing.T) {
	// https://keisan.casio.com/exec/system/1227779487#!
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html
	// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
	// https://en.wikipedia.org/wiki/Julian_day

	testDate := time.Date(
		-4713,
		11,
		25,
		0,
		0,
		0,
		0,
		time.UTC)

	// to 9-digits
	expectedJulianDayNoTimeStr :=
		"0.500000000"

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

func TestJulianDayNoDto_New_09(t *testing.T) {
	// https://keisan.casio.com/exec/system/1227779487#!
	// https://www.aavso.org/jd-calculator
	// http://numerical.recipes/julian.html
	// https://quasar.as.utexas.edu/BillInfo/JulianDateCalc.html
	// https://en.wikipedia.org/wiki/Julian_day

	testDate := time.Date(
		-4713,
		11,
		23,
		12,
		0,
		0,
		0,
		time.UTC)

	_,
	_,
	err := JulianDayNoDto{}.NewFromGregorianDate(testDate)

	if err == nil {
		t.Error("Expected an error return from JulianDayNoDto{}.NewFromGregorianDate(testDate).\n" +
			"However, NO ERROR WAS RETURNED!!!")
		return
	}

}


