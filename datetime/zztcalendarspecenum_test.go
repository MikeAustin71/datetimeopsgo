package datetime

import "testing"

func TestCalendarSpec_01 (t *testing.T) {

	testCalSpecification := CalendarSpec(-1)

	isError := false

	if testCalSpecification < CalSpec.Gregorian() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testCalSpecification'= -1. However, NO ERROR WAS RETURNED!\n")
	}
}

func TestCalendarSpec_02 (t *testing.T) {

	testCalendarSpecification := CalendarSpec(5)

	isError := false

	if testCalendarSpecification > CalSpec.GoucherParker() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testCalendarSpecification'= +5. However, NO ERROR WAS RETURNED!\n")
	}
}

func TestCalendarSpec_03 (t *testing.T) {

	testCalendarSpecification := CalSpec.Gregorian()

	testString := testCalendarSpecification.String()

	if testString != "Gregorian" {
		t.Errorf("Error: Expected String() to return 'Gregorian'\n" +
			"because testCalendarSpecification='CalSpec.Gregorian()'\n" +
			"However, testCalendarSpecification='%v'\n", testString)
	}
}

func TestCalendarSpec_04 (t *testing.T) {

	testCalendarSpecification := CalSpec.Julian()

	testString := testCalendarSpecification.String()

	if testString != "Julian" {
		t.Errorf("Error: Expected String() to return 'Julian'\n" +
			"because testCalendarSpecification='CalSpec.Julian()'\n" +
			"However, testCalendarSpecification='%v'\n", testString)
	}
}

func TestCalendarSpec_05 (t *testing.T) {

	testCalendarSpecification := CalSpec.RevisedJulian()

	testString := testCalendarSpecification.String()

	if testString != "RevisedJulian" {
		t.Errorf("Error: Expected String() to return 'RevisedJulian'\n" +
			"because testCalendarSpecification='CalSpec.RevisedJulian()'\n" +
			"However, testCalendarSpecification='%v'\n", testString)
	}
}

func TestCalendarSpec_06 (t *testing.T) {

	testCalendarSpecification := CalSpec.GoucherParker()

	testString := testCalendarSpecification.String()

	if testString != "GoucherParker" {
		t.Errorf("Error: Expected String() to return 'GoucherParker'\n" +
			"because testCalendarSpecification='CalSpec.GoucherParker()'\n" +
			"However, testCalendarSpecification='%v'\n", testString)
	}
}

func TestCalendarSpecString_001(t *testing.T) {

	r := CalendarSpec(0).None()

	s := r.String()

	if "None" != s {
		t.Errorf("Expected CalendarSpec(0).None() string='%v'.\n" +
			"Instead, string='%v' ",
			"None", s)
	}
}

func TestCalendarSpecString_002(t *testing.T) {

	r := CalendarSpec(0).Gregorian()

	s := r.String()

	if "Gregorian" != s {
		t.Errorf("Expected CalendarSpec(0).Gregorian() string='%v'.\n" +
			"Instead, string='%v' ",
			"Gregorian", s)
	}
}

func TestCalendarSpecString_003(t *testing.T) {

	r := CalendarSpec(0).Julian()

	s := r.String()

	if "Julian" != s {
		t.Errorf("Expected CalendarSpec(0).Julian() string='%v'.\n" +
			"Instead, string='%v' ",
			"Julian", s)
	}
}

func TestCalendarSpecString_004(t *testing.T) {

	r := CalendarSpec(0).RevisedJulian()

	s := r.String()

	if "RevisedJulian" != s {
		t.Errorf("Expected CalendarSpec(0).RevisedJulian() string='%v'.\n" +
			"Instead, string='%v' ",
			"RevisedJulian", s)
	}
}

func TestCalendarSpecString_005(t *testing.T) {

	r := CalendarSpec(0).GoucherParker()

	s := r.String()

	if "GoucherParker" != s {
		t.Errorf("Expected CalendarSpec(0).GoucherParker() string='%v'.\n" +
			"Instead, string='%v' ",
			"GoucherParker", s)
	}
}

func TestCalendarSpecXParseString_001(t *testing.T) {

	expectedStr := "None"

	r, err :=  CalSpec.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by CalSpec.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarSpec(0).None() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarSpecXParseString_002(t *testing.T) {

	expectedStr := "Gregorian"

	r, err :=  CalSpec.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by CalSpec.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarSpec(0).Gregorian() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarSpecXParseString_003(t *testing.T) {

	expectedStr := "Julian"

	r, err :=  CalSpec.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by CalSpec.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarSpec(0).Julian() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarSpecXParseString_004(t *testing.T) {

	expectedStr := "RevisedJulian"

	r, err :=  CalSpec.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by CalSpec.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarSpec(0).RevisedJulian() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}

}

func TestCalendarSpecXParseString_005(t *testing.T) {

	expectedStr := "GoucherParker"

	r, err :=  CalSpec.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by CalSpec.\n" +
			"XParseString(expectedStr, true).\nexpectedStr='%v'\nError='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarSpec(0).RevisedJulian() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarSpecXParseString_006(t *testing.T) {

	expectedStr := "None"

	actualStr := "none"

	r, err :=  CalSpec.XParseString(actualStr, false)

	if err != nil {
		t.Errorf("Error returned by CalSpec.\n" +
			"XParseString(actualStr, false).\n" +
			"actualStr='%v'\nError='%v'\n",
			actualStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarSpec(0).None() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarSpecXParseString_007(t *testing.T) {

	expectedStr := "Gregorian"

	actualStr := "gregorian"

	r, err :=  CalSpec.XParseString(actualStr, false)

	if err != nil {
		t.Errorf("Error returned by CalSpec.\n" +
			"XParseString(actualStr, false).\n" +
			"actualStr='%v'\nError='%v'\n",
			actualStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarSpec(0).Gregorian() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarSpecXParseString_008(t *testing.T) {

	expectedStr := "Julian"

	actualStr := "julian"

	r, err :=  CalSpec.XParseString(actualStr, false)

	if err != nil {
		t.Errorf("Error returned by CalSpec.\n" +
			"XParseString(actualStr, false).\n" +
			"actualStr='%v'\nError='%v'\n",
			actualStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarSpec(0).Julian() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarSpecXParseString_009(t *testing.T) {

	expectedStr := "RevisedJulian"

	actualStr := "revisedjulian"

	r, err :=  CalSpec.XParseString(actualStr, false)

	if err != nil {
		t.Errorf("Error returned by CalSpec.\n" +
			"XParseString(actualStr, false).\n" +
			"actualStr='%v'\nError='%v'\n",
			actualStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarSpec(0).RevisedJulian() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}

}

func TestCalendarSpecXParseString_010(t *testing.T) {

	expectedStr := "GoucherParker"

	actualStr := "goucherparker"

	r, err :=  CalSpec.XParseString(actualStr, false)

	if err != nil {
		t.Errorf("Error returned by CalSpec.\n" +
			"XParseString(actualStr, false).\n" +
			"actualStr='%v'\nError='%v'\n",
			actualStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected CalendarSpec(0).RevisedJulian() string='%v'.\n" +
			"Instead, string='%v'\n",
			expectedStr, r.String())
	}
}

func TestCalendarSpecXParseString_011(t *testing.T) {

	expectedStr := "XRAYxxx"

	_, err :=  CalSpec.XParseString(expectedStr, false)

	if err == nil {
		t.Errorf("Error: Expected an 'error' return from " +
			"CalSpec.XParseString(expectedStr, false).\n" +
			"expectedStr = '%v'\n" +
			"However, NO ERROR WAS RETURNED!!!\n",expectedStr)
	}
}

func TestCalendarSpecXParseString_012(t *testing.T) {

	expectedStr := "XRAYxxx"

	_, err :=  CalSpec.XParseString(expectedStr, true)

	if err == nil {
		t.Errorf("Error: Expected an 'error' return from " +
			"CalSpec.XParseString(expectedStr, true).\n" +
			"expectedStr = '%v'\n" +
			"However, NO ERROR WAS RETURNED!!!\n",expectedStr)
	}
}

func TestCalendarSpecValue_001(t *testing.T) {

	expectedValue := 0

	var r CalendarSpec

	var i int

	r = CalendarSpec(0).None()

	i = int(r)

	if i != expectedValue {
		t.Errorf("Expected 'CalendarSpec(0).None()' value = '%v'.\n" +
			"Instead, returned value was '%v'\n",
			expectedValue, i)
	}
}

func TestCalendarSpecValue_002(t *testing.T) {

	expectedValue := 1

	var r CalendarSpec

	var i int

	r = CalendarSpec(0).Gregorian()

	i = int(r)

	if i != expectedValue {
		t.Errorf("Expected 'CalendarSpec(0).Gregorian()' value = '%v'.\n" +
			"Instead, returned value was '%v'\n",
			expectedValue, i)
	}
}

func TestCalendarSpecValue_003(t *testing.T) {

	expectedValue := 2

	var r CalendarSpec

	var i int

	r = CalendarSpec(0).Julian()

	i = int(r)

	if i != expectedValue {
		t.Errorf("Expected 'CalendarSpec(0).Julian()' value = '%v'.\n" +
			"Instead, returned value was '%v'\n",
			expectedValue, i)
	}
}

func TestCalendarSpecValue_004(t *testing.T) {

	expectedValue := 3

	var r CalendarSpec

	var i int

	r = CalendarSpec(0).RevisedJulian()

	i = int(r)

	if i != expectedValue {
		t.Errorf("Expected 'CalendarSpec(0).RevisedJulian()' value = '%v'.\n" +
			"Instead, returned value was '%v'\n",
			expectedValue, i)
	}
}

func TestCalendarSpecValue_005(t *testing.T) {

	expectedValue := 4

	var r CalendarSpec

	var i int

	r = CalendarSpec(0).GoucherParker()

	i = int(r)

	if i != expectedValue {
		t.Errorf("Expected 'CalendarSpec(0).GoucherParker()' value = '%v'.\n" +
			"Instead, returned value was '%v'\n",
			expectedValue, i)
	}
}

func TestCalendarSpec_XValueInt_001(t *testing.T) {

	testCalendarSpec := CalendarSpec(0).None()

	actualIntValue := testCalendarSpec.XValueInt()

	expectedIntValue := int(CalSpec.None())

	if expectedIntValue != actualIntValue {
		t.Errorf("Error Expected CalSpec.None() actualIntValue='%v'.\n" +
			"Instead, actualIntValue='%v'",
			expectedIntValue, actualIntValue)
	}
}

func TestCalendarSpec_XValueInt_002(t *testing.T) {

	testCalendarSpec := CalendarSpec(0).Gregorian()

	actualIntValue := testCalendarSpec.XValueInt()

	expectedIntValue := int(CalSpec.Gregorian())

	if expectedIntValue != actualIntValue {
		t.Errorf("Error Expected CalSpec.Gregorian() actualIntValue='%v'.\n" +
			"Instead, actualIntValue='%v'",
			expectedIntValue, actualIntValue)
	}
}

func TestCalendarSpec_XValueInt_003(t *testing.T) {

	testCalendarSpec := CalendarSpec(0).Julian()

	actualIntValue := testCalendarSpec.XValueInt()

	expectedIntValue := int(CalSpec.Julian())

	if expectedIntValue != actualIntValue {
		t.Errorf("Error Expected CalSpec.Julian() actualIntValue='%v'.\n" +
			"Instead, actualIntValue='%v'",
			expectedIntValue, actualIntValue)
	}
}

func TestCalendarSpec_XValueInt_004(t *testing.T) {

	testCalendarSpec := CalendarSpec(0).RevisedJulian()

	actualIntValue := testCalendarSpec.XValueInt()

	expectedIntValue := int(CalSpec.RevisedJulian())

	if expectedIntValue != actualIntValue {
		t.Errorf("Error Expected CalSpec.RevisedJulian() actualIntValue='%v'.\n" +
			"Instead, actualIntValue='%v'",
			expectedIntValue, actualIntValue)
	}
}

func TestCalendarSpec_XValueInt_005(t *testing.T) {

	testCalendarSpec := CalendarSpec(0).GoucherParker()

	actualIntValue := testCalendarSpec.XValueInt()

	expectedIntValue := int(CalSpec.GoucherParker())

	if expectedIntValue != actualIntValue {
		t.Errorf("Error Expected CalSpec.GoucherParker() actualIntValue='%v'.\n" +
			"Instead, actualIntValue='%v'",
			expectedIntValue, actualIntValue)
	}
}

func TestCalendarSpec_XisValid_001(t *testing.T) {

	testCalendarSpec := CalendarSpec(0).GoucherParker()

	if !testCalendarSpec.XIsValid() {
		t.Error("Error: Expected testCalendarSpec.XIsValid()=='true'.\n" +
			"Value is CalendarSpec(0).GoucherParker().\n" +
			"Instead, testCalendarSpec.XIsValid()=='false'.\n")
	}
}

func TestCalendarSpec_XisValid_002(t *testing.T) {

	testCalendarSpec := CalendarSpec(0).None()

	if testCalendarSpec.XIsValid() {
		t.Error("Error: Expected testCalendarSpec.XIsValid()=='false'.\n" +
			"Value is CalendarSpec(0).None().\n" +
			"Instead, testCalendarSpec.XIsValid()=='true'.\n")
	}
}
func TestCalendarSpec_XisValid_003(t *testing.T) {

	testCalendarSpec := CalendarSpec(-7)

	if testCalendarSpec.XIsValid() {
		t.Error("Error: Expected testCalendarSpec.XIsValid()=='false'.\n" +
			"Value is CalendarSpec(-7).\n" +
			"Instead, testCalendarSpec.XIsValid()=='true'.\n")
	}
}

func TestCalendarSpec_XisValid_004(t *testing.T) {

	testCalendarSpec := CalendarSpec(0).RevisedJulian()

	if !testCalendarSpec.XIsValid() {
		t.Error("Error: Expected testCalendarSpec.XIsValid()=='true'.\n" +
			"Value is CalendarSpec(0).RevisedJulian().\n" +
			"Instead, testCalendarSpec.XIsValid()=='false'.\n")
	}
}

func TestCalendarSpec_XisValid_005(t *testing.T) {

	testCalendarSpec := CalendarSpec(0).Julian()

	if !testCalendarSpec.XIsValid() {
		t.Error("Error: Expected testCalendarSpec.XIsValid()=='true'.\n" +
			"Value is CalendarSpec(0).Julian().\n" +
			"Instead, testCalendarSpec.XIsValid()=='false'.\n")
	}
}

func TestCalendarSpec_XisValid_006(t *testing.T) {

	testCalendarSpec := CalendarSpec(0).Gregorian()

	if !testCalendarSpec.XIsValid() {
		t.Error("Error: Expected testCalendarSpec.XIsValid()=='true'.\n" +
			"Value is CalendarSpec(0).Gregorian().\n" +
			"Instead, testCalendarSpec.XIsValid()=='false'.\n")
	}
}
