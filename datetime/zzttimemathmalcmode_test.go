package datetime

import "testing"

func TestTimeMathCalcMode_01 (t *testing.T) {

	testTimeMathCalcMode := TimeMathCalcMode(-1)

	isError := false

	if testTimeMathCalcMode < TCalcMode.LocalTimeZone() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testTimeMathCalcMode'= -1. However, NO ERROR WAS RETURNED!\n")
	}
}

func TestTimeMathCalcMode_02 (t *testing.T) {

	testTimeMathCalcMode := TimeMathCalcMode(57)

	isError := false

	if testTimeMathCalcMode > TCalcMode.UtcTimeZone() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testTimeMathCalcMode'= 57. However, NO ERROR WAS RETURNED!\n")
	}
}

func TestTimeMathCalcMode_03 (t *testing.T) {

	testTimeMathCalcMode := TCalcMode.LocalTimeZone()

	testString := testTimeMathCalcMode.String()

	if testString != "LocalTimeZone" {
		t.Errorf("Error: Expected String() to return 'LocalTimeZone'\n" +
			"because testTimeMathCalcMode='TCalcMode.LocalTimeZone()'\n" +
			"However, testTimeMathCalcMode='%v'\n", testString)
	}
}

func TestTimeMathCalcMode_04 (t *testing.T) {

	testTimeMathCalcMode := TCalcMode.UtcTimeZone()

	testString := testTimeMathCalcMode.String()

	if testString != "UtcTimeZone" {
		t.Errorf("Error: Expected String() to return 'UtcTimeZone'\n" +
			"because testTimeMathCalcMode='TCalcMode.UtcTimeZone()'\n" +
			"However, testTimeMathCalcMode='%v'\n", testString)
	}
}

func TestTimeMathCalcMode_05 (t *testing.T) {

	textString := "LocalTimeZone"

	actualValue, err :=
		TimeMathCalcMode(0).XParseString(textString, true)

	if err != nil {
		t.Errorf("Error returned by TimeMathCalcMode(0)." +
			"XParseString(textString, true)\n" +
			"textString='%v'\n" +
			"Error='%v'\n", textString, err.Error())
		return
	}

	if TCalcMode.LocalTimeZone() != actualValue {
		t.Errorf("Error: Expected actualValue=TCalcMode.LocalTimeZone()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTimeMathCalcMode_06 (t *testing.T) {

	textString := "UtcTimeZone"

	actualValue, err :=
		TimeMathCalcMode(0).XParseString(textString, true)

	if err != nil {
		t.Errorf("Error returned by TimeMathCalcMode(0)." +
			"XParseString(textString, true)\n" +
			"textString='%v'\n" +
			"Error='%v'\n", textString, err.Error())
		return
	}

	if TCalcMode.UtcTimeZone() != actualValue {
		t.Errorf("Error: Expected actualValue=TCalcMode.UtcTimeZone()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTimeMathCalcMode_07 (t *testing.T) {

	textString := "localtimezone"

	actualValue, err :=
		TimeMathCalcMode(0).XParseString(textString, false)

	if err != nil {
		t.Errorf("Error returned by TimeMathCalcMode(0)." +
			"XParseString(textString, true)\n" +
			"textString='%v'\n" +
			"Error='%v'\n", textString, err.Error())
		return
	}

	if TCalcMode.LocalTimeZone() != actualValue {
		t.Errorf("Error: Expected actualValue=TCalcMode.LocalTimeZone()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTimeMathCalcMode_08 (t *testing.T) {

	textString := "utctimezone"

	actualValue, err :=
		TimeMathCalcMode(0).XParseString(textString, false)

	if err != nil {
		t.Errorf("Error returned by TimeMathCalcMode(0)." +
			"XParseString(textString, true)\n" +
			"textString='%v'\n" +
			"Error='%v'\n", textString, err.Error())
		return
	}

	if TCalcMode.UtcTimeZone() != actualValue {
		t.Errorf("Error: Expected actualValue=TCalcMode.UtcTimeZone()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTimeMathCalcMode_09 (t *testing.T) {

	expectedStr := "LocalTimeZone"

	r, err :=  TCalcMode.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TCalcMode." +
			"XParseString(expectedStr, true).\n" +
			"expectedStr='%v'\n" +
			"Error='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TCalcMode.LocalTimeZone() string='%v'.\n" +
			"Instead, string='%v'\n", expectedStr, r.String())
	}

}

func TestTimeMathCalcMode_10 (t *testing.T) {

	expectedStr := "UtcTimeZone"

	r, err :=  TCalcMode.XParseString(expectedStr, true)

	if err != nil {
		t.Errorf("Error returned by TCalcMode." +
			"XParseString(expectedStr, true).\n" +
			"expectedStr='%v'\n" +
			"Error='%v'\n",
			expectedStr, err.Error())
		return
	}

	if expectedStr != r.String() {
		t.Errorf("Expected TCalcMode.UtcTimeZone() string='%v'.\n" +
			"Instead, string='%v'\n", expectedStr, r.String())
	}

}

func TestTimeMathCalcMode_11 (t *testing.T) {

	expectedStr := "UtcTimeXXX"

	_, err :=  TCalcMode.XParseString(expectedStr, true)

	if err == nil {
		t.Errorf("Error: Expected an 'error' return from " +
			"TCalcMode.XParseString(expectedStr, false).\n" +
			"expectedStr = '%v'. NO ERROR WAS RETURNED!!! ",
			expectedStr)
	}

}

func TestTimeMathCalcMode_12 (t *testing.T) {

	expectedStr := "localTimeXXX"

	_, err :=  TCalcMode.XParseString(expectedStr, false)

	if err == nil {
		t.Errorf("Error: Expected an 'error' return from " +
			"TCalcMode.XParseString(expectedStr, false).\n" +
			"expectedStr = '%v'. NO ERROR WAS RETURNED!!! ",
			expectedStr)
	}

}

func TestTimeMathCalcModeTypeValue_01 (t *testing.T) {

	var r TimeMathCalcMode

	var i int

	r = TimeMathCalcMode(0).None()

	i = int(r)

	if i != 0 {
		t.Errorf("Expected 'TimeMathCalcMode(0).None()' value = 0.\n" +
			"Instead, got %v", i)
	}

}

func TestTimeMathCalcModeTypeValue_02 (t *testing.T) {

	var r TimeMathCalcMode

	var i int

	r = TimeMathCalcMode(0).LocalTimeZone()

	i = int(r)

	if i != 1 {
		t.Errorf("Expected 'TimeMathCalcMode(0).LocalTimeZone()' value = 1.\n" +
			"Instead, got %v", i)
	}

}

func TestTimeMathCalcModeTypeValue_03 (t *testing.T) {

	var r TimeMathCalcMode

	var i int

	r = TimeMathCalcMode(0).LocalTimeZone()

	i = r.XValueInt()

	if i != 1 {
		t.Errorf("Expected 'TimeMathCalcMode(0).LocalTimeZone()' value = 1.\n" +
			"Instead, got %v", i)
	}

}
func TestTimeMathCalcModeTypeValue_04 (t *testing.T) {

	var r TimeMathCalcMode

	var i int

	r = TimeMathCalcMode(0).UtcTimeZone()

	i = int(r)

	if i != 2 {
		t.Errorf("Expected 'TimeMathCalcMode(0).UtcTimeZone()' value = 2.\n" +
			"Instead, got %v", i)
	}

}

func TestTimeMathCalcModeTypeValue_05 (t *testing.T) {

	var r TimeMathCalcMode

	var i int

	r = TimeMathCalcMode(0).UtcTimeZone()

	i = r.XValueInt()

	if i != 2 {
		t.Errorf("Expected 'TimeMathCalcMode(0).UtcTimeZone()' value = 2.\n" +
			"Instead, got %v", i)
	}

}

func TestTimeMathCalcMode_XIsValid_01 (t *testing.T) {

	timeMathCalcMode := TCalcMode.LocalTimeZone()

	if !timeMathCalcMode.XIsValid() {
		t.Error("Error: Expected timeMathCalcMode.XIsValid()=='true'.\n" +
			"Value is TCalcMode.LocalTimeZone().\n" +
			"Instead, timeMathCalcMode.XIsValid()=='false'.\n")
	}

}

func TestTimeMathCalcMode_XIsValid_02 (t *testing.T) {

	timeMathCalcMode := TCalcMode.UtcTimeZone()

	if !timeMathCalcMode.XIsValid() {
		t.Error("Error: Expected timeMathCalcMode.XIsValid()=='true'.\n" +
			"Value is TCalcMode.UtcTimeZone().\n" +
			"Instead, timeMathCalcMode.XIsValid()=='false'.\n")
	}

}

func TestTimeMathCalcMode_XIsValid_03 (t *testing.T) {

	timeMathCalcMode := TCalcMode.None()

	if timeMathCalcMode.XIsValid() {
		t.Error("Error: Expected timeMathCalcMode.XIsValid()=='false'.\n" +
			"Value is TCalcMode.None().\n" +
			"Instead, timeMathCalcMode.XIsValid()=='true'.\n")
	}

}

func TestTimeMathCalcMode_XIsValid_04 (t *testing.T) {

	timeMathCalcMode := TimeMathCalcMode(-99)

	if timeMathCalcMode.XIsValid() {
		t.Error("Error: Expected timeMathCalcMode.XIsValid()=='false'.\n" +
			"Value is TimeMathCalcMode(-99).\n" +
			"Instead, timeMathCalcMode.XIsValid()=='true'.\n")
	}

}

func TestTimeMathCalcMode_XIsValid_05 (t *testing.T) {

	timeMathCalcMode := TimeMathCalcMode(105)

	if timeMathCalcMode.XIsValid() {
		t.Error("Error: Expected timeMathCalcMode.XIsValid()=='false'.\n" +
			"Value is TimeMathCalcMode(105).\n" +
			"Instead, timeMathCalcMode.XIsValid()=='true'.\n")
	}

}

func TestTimeMathCalcMode_XIsValid_06 (t *testing.T) {

	timeMathCalcMode := TCalcMode.UtcTimeZone()

	if !timeMathCalcMode.XIsValid() {
		t.Error("Error: Expected timeMathCalcMode.XIsValid()=='true'.\n" +
			"Value is TCalcMode.UtcTimeZone().\n" +
			"Instead, timeMathCalcMode.XIsValid()=='false'.\n")
	}

	intValue := timeMathCalcMode.XValueInt()

	if intValue != 2 {
		t.Errorf("Error: Expected TCalcMode.UtcTimeZone() " +
			"to yield an integer value of '2'.\n" +
			"Instead, integer value = '%v'.\n" +
			"Also, timeMathCalcMode.XIsValid()='%v'",
			intValue, timeMathCalcMode.XIsValid())
	}

}

func TestTimeMathCalcMode_XIsValid_07 (t *testing.T) {

	timeMathCalcMode := TCalcMode.LocalTimeZone()

	if !timeMathCalcMode.XIsValid() {
		t.Error("Error: Expected timeMathCalcMode.XIsValid()=='true'.\n" +
			"Value is TCalcMode.LocalTimeZone().\n" +
			"Instead, timeMathCalcMode.XIsValid()=='false'.\n")
	}

	intValue := timeMathCalcMode.XValueInt()

	if intValue != 1 {
		t.Errorf("Error: Expected TCalcMode.LocalTimeZone() " +
			"to yield an integer value of '1'.\n" +
			"Instead, integer value = '%v'.\n" +
			"Also, timeMathCalcMode.XIsValid()='%v'",
			intValue, timeMathCalcMode.XIsValid())
	}

}