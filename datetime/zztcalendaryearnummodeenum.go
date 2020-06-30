package datetime

import "testing"

func TestCalendarYearNumMode_01 (t *testing.T) {

	testCalYearNumMode := CalendarYearNumMode(-1)

	isError := false

	if testCalYearNumMode < CalYearMode.Astronomical() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testCalYearNumMode'= -1. However, NO ERROR WAS RETURNED!\n")
	}
}

func TestCalendarYearNumMode_02 (t *testing.T) {

	testCalYearNumMode := CalendarYearNumMode(5)

	isError := false

	if testCalYearNumMode > CalYearMode.CommonEra() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testCalYearNumMode'= +5.\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestCalendarYearNumMode_03 (t *testing.T) {

	testCalYearNumMode := CalYearMode.Astronomical()

	testString := testCalYearNumMode.String()

	if testString != "Astronomical" {
		t.Errorf("Error: Expected String() to return 'Astronomical'\n" +
			"because testCalYearNumMode='CalYearMode.Astronomical()'\n" +
			"However, testCalYearNumMode='%v'\n", testString)
	}
}

func TestCalendarYearNumMode_04 (t *testing.T) {

	testCalYearNumMode := CalYearMode.Astronomical()

	testString := testCalYearNumMode.String()

	if testString != "Astronomical" {
		t.Errorf("Error: Expected String() to return 'Astronomical'\n" +
			"because testCalYearNumMode='CalYearMode.Astronomical()'\n" +
			"However, testCalYearNumMode='%v'\n", testString)
	}
}

func TestCalendarYearNumMode_05 (t *testing.T) {

	testCalYearNumMode := CalYearMode.CommonEra()

	testString := testCalYearNumMode.String()

	if testString != "CommonEra" {
		t.Errorf("Error: Expected String() to return 'CommonEra'\n" +
			"because testCalYearNumMode='CalYearMode.CommonEra()'\n" +
			"However, testCalYearNumMode='%v'\n", testString)
	}
}

func TestCalendarYearNumModeString_001(t *testing.T) {

	r := CalendarYearNumMode(0).None()

	s := r.String()

	if "None" != s {
		t.Errorf("Expected CalendarYearNumMode(0).None() string='%v'.\n" +
			"Instead, string='%v' ",
			"None", s)
	}
}

func TestCalendarYearNumModeString_002(t *testing.T) {

	r := CalendarYearNumMode(0).Astronomical()

	s := r.String()

	if "Gregorian" != s {
		t.Errorf("Expected CalendarYearNumMode(0).Astronomical() string='%v'.\n" +
			"Instead, string='%v' ",
			"Astronomical", s)
	}
}

func TestCalendarYearNumModeString_003(t *testing.T) {

	r := CalendarYearNumMode(0).CommonEra()

	s := r.String()

	if "Julian" != s {
		t.Errorf("Expected CalendarYearNumMode(0).CommonEra() string='%v'.\n" +
			"Instead, string='%v' ",
			"CommonEra", s)
	}
}

func TestCalendarYearNumModeXParseString_001(t *testing.T) {

	expectedStr := "None"

	r, err :=  CalYearMode.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by CalYearMode.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarYearNumMode(0).None() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarYearNumModeXParseString_002(t *testing.T) {

	expectedStr := "Astronomical"

	r, err :=  CalYearMode.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by CalYearMode.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalYearMode.Astronomical() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarYearNumModeXParseString_003(t *testing.T) {

	expectedStr := "CommonEra"

	r, err :=  CalYearMode.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by CalYearMode.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalYearMode.CommonEra() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarYearNumModeXParseString_004(t *testing.T) {

	expectedStr := "None"

	actualStr := "none"

	r, err :=  CalYearMode.XParseString(actualStr, false)

	if err != nil {
		t.Errorf("Error returned by CalYearMode.\n" +
			"XParseString(actualStr, false).\n" +
			"actualStr='%v'\nError='%v'\n",
			actualStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalYearMode.None() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarYearNumModeXParseString_005(t *testing.T) {

	expectedStr := "Astronomical"

	actualStr := "astronomical"

	r, err :=  CalYearMode.XParseString(actualStr, false)

	if err != nil {
		t.Errorf("Error returned by CalYearMode.\n" +
			"XParseString(actualStr, false).\n" +
			"actualStr='%v'\nError='%v'\n",
			actualStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected  CalYearMode.Astronomical() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarYearNumModeXParseString_006(t *testing.T) {

	expectedStr := "CommonEra"

	actualStr := "commonera"

	r, err :=  CalYearMode.XParseString(actualStr, false)

	if err != nil {
		t.Errorf("Error returned by CalYearMode.\n" +
			"XParseString(actualStr, false).\n" +
			"actualStr='%v'\nError='%v'\n",
			actualStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalYearMode.CommonEra() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarYearNumModeXParseString_007(t *testing.T) {

	expectedStr := "XRAYxxx"

	_, err :=  CalYearMode.XParseString(expectedStr, false)

	if err == nil {
		t.Errorf("Error: Expected an 'error' return from " +
			"CalYearMode.XParseString(expectedStr, false).\n" +
			"expectedStr = '%v'\n" +
			"However, NO ERROR WAS RETURNED!!!\n",expectedStr)
	}
}

func TestCalendarYearNumModeXParseString_008(t *testing.T) {

	expectedStr := "XRAYxxx"

	_, err :=  CalYearMode.XParseString(expectedStr, true)

	if err == nil {
		t.Errorf("Error: Expected an 'error' return from " +
			"CalYearMode.XParseString(expectedStr, true).\n" +
			"expectedStr = '%v'\n" +
			"However, NO ERROR WAS RETURNED!!!\n",expectedStr)
	}
}

func TestCalendarYearNumModeValue_001(t *testing.T) {

	expectedValue := 0

	var r CalendarYearNumMode

	var i int

	r = CalendarYearNumMode(0).None()

	i = int(r)

	if i != expectedValue {
		t.Errorf("Expected 'CalendarYearNumMode(0).None()' value = '%v'.\n" +
			"Instead, returned value was '%v'\n",
			expectedValue, i)
	}
}

func TestCalendarYearNumModeValue_002(t *testing.T) {

	expectedValue := 1

	var r CalendarYearNumMode

	var i int

	r = CalendarYearNumMode(0).Astronomical()

	i = int(r)

	if i != expectedValue {
		t.Errorf("Expected 'CalendarYearNumMode(0).Astronomical()' value = '%v'.\n" +
			"Instead, returned value was '%v'\n",
			expectedValue, i)
	}
}

func TestCalendarYearNumModeValue_003(t *testing.T) {

	expectedValue := 2

	var r CalendarYearNumMode

	var i int

	r = CalendarYearNumMode(0).CommonEra()

	i = int(r)

	if i != expectedValue {
		t.Errorf("Expected 'CalendarYearNumMode(0).CommonEra()' value = '%v'.\n" +
			"Instead, returned value was '%v'\n",
			expectedValue, i)
	}
}

func TestCalendarYearNumMode_XValueInt_001(t *testing.T) {

	testCalendarYearNumMode := CalYearMode.None()

	actualIntValue := testCalendarYearNumMode.XValueInt()

	expectedIntValue := int(CalYearMode.None())

	if expectedIntValue != actualIntValue {
		t.Errorf("Error Expected CalYearMode.None() actualIntValue='%v'.\n" +
			"Instead, actualIntValue='%v'",
			expectedIntValue, actualIntValue)
	}
}

func TestCalendarYearNumMode_XValueInt_002(t *testing.T) {

	testCalendarYearNumMode := CalendarYearNumMode(0).Astronomical()

	actualIntValue := testCalendarYearNumMode.XValueInt()

	expectedIntValue := int(CalYearMode.Astronomical())

	if expectedIntValue != actualIntValue {
		t.Errorf("Error Expected CalYearMode.Astronomical() actualIntValue='%v'.\n" +
			"Instead, actualIntValue='%v'",
			expectedIntValue, actualIntValue)
	}
}

func TestCalendarYearNumMode_XValueInt_003(t *testing.T) {

	testCalendarYearNumMode := CalendarYearNumMode(0).CommonEra()

	actualIntValue := testCalendarYearNumMode.XValueInt()

	expectedIntValue := int(CalYearMode.CommonEra())

	if expectedIntValue != actualIntValue {
		t.Errorf("Error Expected CalYearMode.CommonEra() actualIntValue='%v'.\n" +
			"Instead, actualIntValue='%v'",
			expectedIntValue, actualIntValue)
	}
}

func TestCalendarYearNumMode_XisValid_001(t *testing.T) {

	testCalendarYearNumMode := CalendarYearNumMode(0).Astronomical()

	if !testCalendarYearNumMode.XIsValid() {
		t.Error("Error: Expected testCalendarYearNumMode.XIsValid()=='true'.\n" +
			"Value is CalendarYearNumMode(0).Astronomical().\n" +
			"Instead, testCalendarYearNumMode.XIsValid()=='false'.\n")
	}
}

func TestCalendarYearNumMode_XisValid_002(t *testing.T) {

	testCalendarYearNumMode := CalendarYearNumMode(0).None()

	if testCalendarYearNumMode.XIsValid() {
		t.Error("Error: Expected testCalendarYearNumMode.XIsValid()=='false'.\n" +
			"Value is CalendarYearNumMode(0).None().\n" +
			"Instead, testCalendarYearNumMode.XIsValid()=='true'.\n")
	}
}

func TestCalendarYearNumMode_XisValid_003(t *testing.T) {

	testCalendarYearNumMode := CalendarYearNumMode(-7)

	if testCalendarYearNumMode.XIsValid() {
		t.Error("Error: Expected testCalendarYearNumMode.XIsValid()=='false'.\n" +
			"Value is CalendarYearNumMode(-7).\n" +
			"Instead, testCalendarYearNumMode.XIsValid()=='true'.\n")
	}
}

func TestCalendarYearNumMode_XisValid_004(t *testing.T) {

	testCalendarYearNumMode := CalendarYearNumMode(0).CommonEra()

	if !testCalendarYearNumMode.XIsValid() {
		t.Error("Error: Expected testCalendarYearNumMode.XIsValid()=='true'.\n" +
			"Value is CalendarYearNumMode(0).CommonEra().\n" +
			"Instead, testCalendarYearNumMode.XIsValid()=='false'.\n")
	}
}

func TestCalendarYearNumMode_XisValid_005(t *testing.T) {

	testCalendarYearNumMode := CalendarYearNumMode(0).Astronomical()

	if !testCalendarYearNumMode.XIsValid() {
		t.Error("Error: Expected testCalendarYearNumMode.XIsValid()=='true'.\n" +
			"Value is CalendarYearNumMode(0).Astronomical().\n" +
			"Instead, testCalendarYearNumMode.XIsValid()=='false'.\n")
	}
}

func TestCalendarYearNumMode_XisValid_006(t *testing.T) {

	testCalendarYearNumMode := CalendarYearNumMode(0).CommonEra()

	if !testCalendarYearNumMode.XIsValid() {
		t.Error("Error: Expected testCalendarYearNumMode.XIsValid()=='true'.\n" +
			"Value is CalendarYearNumMode(0).CommonEra().\n" +
			"Instead, testCalendarYearNumMode.XIsValid()=='false'.\n")
	}
}

