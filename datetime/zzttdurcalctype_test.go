package datetime

import (
	"testing"
)

func TestTDurCalcType_01 (t *testing.T) {

	testTimeDurCalcType := TDurCalcType(-1)

	isError := false

	if testTimeDurCalcType < TDurCalc.StdYearMth() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testTimeDurCalcType'= -1. However, NO ERROR WAS RETURNED!\n")
	}
}


func TestTDurCalcType_02 (t *testing.T) {

	testTimeDurCalcType := TDurCalcType(12)

	isError := false

	if testTimeDurCalcType > TDurCalc.GregorianYears() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testTimeDurCalcType'= +11. However, NO ERROR WAS RETURNED!\n")
	}
}

func TestTDurCalcType_03(t *testing.T) {

	testTimeDurCalcType := TDurCalc.CumHours()

	testString := testTimeDurCalcType.String()

	if testString != "CumHours" {
		t.Errorf("Error: Expected String() to return 'CumHours'\n" +
			"because testTimeDurCalcType='TDurCalc.CumHours()'\n" +
			"However, testTimeDurCalcType='%v'\n", testString)
	}
}

func TestTDurCalcType_04(t *testing.T) {

	textString := "CumMicroseconds"

	actualValue, err :=
		TDurCalcType(0).XParseString(textString, true)

	if err != nil {
		t.Errorf("Error returned by LocationNameType(0)." +
			"XParseString(textString, true)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if TDurCalc.CumMicroseconds() != actualValue {
		t.Errorf("Error: Expected actualValue=TDurCalc.CumMicroseconds()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTDurCalcType_05(t *testing.T) {

	textString := "cummicroseconds"

	actualValue, err :=
		TDurCalcType(0).XParseString(textString, false)

	if err != nil {
		t.Errorf("Error returned by LocationNameType(0).XParseString(textString, true)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if TDurCalc.CumMicroseconds() != actualValue {
		t.Errorf("Error: Expected actualValue=TDurCalc.CumMicroseconds()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTDurCalcType_06(t *testing.T) {

	testLocNameType := TDurCalc.CumNanoseconds()

	actualValue := testLocNameType.XValue()

	if TDurCalc.CumNanoseconds() != actualValue {
		t.Errorf("Error: Expected actualValue=TDurCalc.CumNanoseconds()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTDurCalcTypeString_001(t *testing.T) {

	r := TDurCalcType(0).StdYearMth()

	s := r.String()

	if "StdYearMth" != s {
		t.Errorf("Expected TDurCalcType(0).StdYearMth() string='%v'. Instead, string='%v' ",
			"StdYearMth", s)
	}

}

func TestTDurCalcTypeString_002(t *testing.T) {

	r := TDurCalcType(0).CumMonths()
	expectedStr := "CumMonths"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumMonths() string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestTDurCalcTypeString_003(t *testing.T) {

	r := TDurCalcType(0).CumWeeks()
	expectedStr := "CumWeeks"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumWeeks() string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestTDurCalcTypeString_004(t *testing.T) {

	r := TDurCalcType(0).CumDays()
	expectedStr := "CumDays"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumDays() string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestTDurCalcTypeString_005(t *testing.T) {

	r := TDurCalcType(0).CumHours()
	expectedStr := "CumHours"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumHours() string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestTDurCalcTypeString_006(t *testing.T) {

	r := TDurCalcType(0).CumMinutes()
	expectedStr := "CumMinutes"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumMinutes() string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestTDurCalcTypeString_007(t *testing.T) {

	r := TDurCalcType(0).CumSeconds()
	expectedStr := "CumSeconds"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumSeconds() string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestTDurCalcTypeString_008(t *testing.T) {

	r := TDurCalcType(0).GregorianYears()
	expectedStr := "GregorianYears"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).GregorianYears() string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestTDurCalcTypeString_009(t *testing.T) {

	r := TDurCalcType(0).CumMilliseconds()
	expectedStr := "CumMilliseconds"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumMilliseconds() string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestTDurCalcTypeString_010(t *testing.T) {

	r := TDurCalcType(0).CumMicroseconds()
	expectedStr := "CumMicroseconds"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumMicroseconds() string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestTDurCalcTypeString_011(t *testing.T) {

	r := TDurCalcType(0).CumNanoseconds()
	expectedStr := "CumNanoseconds"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. Instead, string='%v' ",
			expectedStr, s)
	}

}

func TestTDurCalcTypeString_012(t *testing.T) {

	expectedStr := "StdYearMth"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc." +
			"XParseString(expectedStr, true). expectedStr='%v' Error='%v' ",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_013(t *testing.T) {

	expectedStr := "CumMonths"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc." +
			"XParseString(expectedStr, true). expectedStr='%v' Error='%v' ",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_014(t *testing.T) {

	expectedStr := "CumWeeks"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc." +
			"XParseString(expectedStr, true). expectedStr='%v' Error='%v' ",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_015(t *testing.T) {

	expectedStr := "CumDays"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc." +
			"XParseString(expectedStr, true).\n" +
			"expectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_016(t *testing.T) {

	expectedStr := "CumHours"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc." +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds()\nstring='%v'.\n" +
			"Instead, string='%v'\n", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_017(t *testing.T) {

	expectedStr := "CumMinutes"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds()\nstring='%v'.\n" +
			"Instead, string='%v'\n", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_018(t *testing.T) {

	expectedStr := "CumSeconds"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc." +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'.\n" +
			"Instead, string='%v'\n", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_019(t *testing.T) {

	expectedStr := "CumMilliseconds"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_020(t *testing.T) {

	expectedStr := "CumMicroseconds"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_021(t *testing.T) {

	expectedStr := "CumNanoseconds"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_022(t *testing.T) {

	expectedStr := "GregorianYears"

	r, err :=  TDurCalc.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_023(t *testing.T) {

	expectedStr := "StdYearMth"
	testStr := "stdyearmth"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(testStr, true).\ntestStr='%v'\nError='%v'\n",
			testStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds()\nstring='%v'.\n" +
			"Instead, string='%v'\n", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_024(t *testing.T) {

	expectedStr := "CumMonths"
	testStr := "cummonths"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(testStr, true).\ntestStr='%v'\nError='%v'\n",
			testStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_025(t *testing.T) {

	expectedStr := "CumWeeks"
	testStr := "cumweeks"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(testStr, true).\ntestStr='%v'\nError='%v'\n",
			testStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'\n" +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_026(t *testing.T) {

	expectedStr := "CumDays"
	testStr := "cumdays"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc." +
			"XParseString(testStr, true).\n" +
			"testStr='%v'\nError='%v' ",
			testStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_027(t *testing.T) {

	expectedStr := "CumHours"
	testStr := "cumhourS"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc." +
			"XParseString(testStr, true).\n" +
			"testStr='%v'\n" +
			"Error='%v'\n",
			testStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'.\n" +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_028(t *testing.T) {

	expectedStr := "CumMinutes"
	testStr := "cumminUtes"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc." +
			"XParseString(testStr, true).\n" +
			"testStr='%v'\nError='%v'\n",
			testStr, err.Error())

		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_029(t *testing.T) {

	expectedStr := "CumSeconds"
	testStr := "cumseconds"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(testStr, true). testStr='%v'\nError='%v'\n",
			testStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_030(t *testing.T) {

	expectedStr := "CumMilliseconds"
	testStr := "cummiLLiseconds"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(testStr, true). testStr='%v'\n" +
			"Error='%v'\n",
			testStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_031(t *testing.T) {

	expectedStr := "CumMicroseconds"
	testStr := "cuMMicroseconds"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc." +
			"XParseString(testStr, true). testStr='%v' Error='%v' ",
			testStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_032(t *testing.T) {

	expectedStr := "CumNanoseconds"
	testStr := "cumnANoseconds"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(testStr, true). testStr='%v'\nError='%v'\n",
			testStr, err.Error())

		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_033(t *testing.T) {

	expectedStr := "GregorianYears"
	testStr := "grEGorianyears"

	r, err :=  TDurCalc.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by TDurCalc.\n" +
			"XParseString(testStr, true). testStr='%v'\nError='%v'\n",
			testStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TDurCalcType(0).CumNanoseconds() string='%v'. " +
			"Instead, string='%v' ", expectedStr, r.String())
	}

}

func TestTDurCalcTypeString_034(t *testing.T) {

	expectedStr := "XRAYxxx"

	_, err :=  TDurCalc.XParseString(expectedStr, false)

	if err == nil {
		t.Errorf("Error: Expected an 'error' return from " +
			"TDurCalc.XParseString(expectedStr, false). " +
			"expectedStr = '%v'. NO ERROR WAS RETURNED!!! ",expectedStr)
	}

}

func TestTDurCalcTypeString_035(t *testing.T) {

	expectedStr := "XRAYxxx"

	_, err :=  TDurCalc.XParseString(expectedStr, true)

	if err == nil {
		t.Errorf("Error: Expected an 'error' return from " +
			"TDurCalc.XParseString(expectedStr, false). " +
			"expectedStr = '%v'. NO ERROR WAS RETURNED!!! ",expectedStr)
	}

}


func TestTDurCalcTypeValue_001(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).None()

	i = int(r)

	if i != 0 {
		t.Errorf("Expected 'TDurCalcType(0).None()' value = 0. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_002(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumMonths()

	i = int(r)

	if i != 2 {
		t.Errorf("Expected 'TDurCalcType(0).CumMonths()' value = 1. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_003(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumWeeks()

	i = int(r)

	if i != 3 {
		t.Errorf("Expected 'TDurCalcType(0).CumWeeks()' value = 3. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_004(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumDays()

	i = int(r)

	if i != 4 {
		t.Errorf("Expected 'TDurCalcType(0).CumDays()' value = 4. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_005(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumHours()

	i = int(r)

	if i != 5 {
		t.Errorf("Expected 'TDurCalcType(0).CumHours()' value = 5. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_006(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumMinutes()

	i = int(r)

	if i != 6 {
		t.Errorf("Expected 'TDurCalcType(0).CumMinutes()' value = 6. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_007(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumSeconds()

	i = int(r)

	if i != 7 {
		t.Errorf("Expected 'TDurCalcType(0).CumSeconds()' value = 7. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_008(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumMilliseconds()

	i = int(r)

	if i != 8{
		t.Errorf("Expected 'TDurCalcType(0).CumMilliseconds()' value = 8. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_009(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumMicroseconds()

	i = int(r)

	if i != 9 {
		t.Errorf("Expected 'TDurCalcType(0).CumMicroseconds()' value = 9. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_010(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumNanoseconds()

	i = int(r)

	if i != 10 {
		t.Errorf("Expected 'TDurCalcType(0).CumNanoseconds()' value = 10. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_011(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).GregorianYears()

	i = int(r)

	if i != 11 {
		t.Errorf("Expected 'TDurCalcType(0).GregorianYears()' value = 11. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_012(t *testing.T) {

	testTDurCalcType := TDurCalcType(0).GregorianYears()

	actualIntValue := testTDurCalcType.XValueInt()

	expectedIntValue := int(TDurCalc.GregorianYears())

	if expectedIntValue != actualIntValue {
		t.Errorf("Error Expected actualIntValue='%v'.\n" +
			"Instead, actualIntValue='%v'",
			expectedIntValue, actualIntValue)
	}

}
