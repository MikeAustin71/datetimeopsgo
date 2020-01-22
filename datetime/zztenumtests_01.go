package datetime

import "testing"


func TestLocationNameType_01 (t *testing.T) {

	testLocNameType := LocationNameType(-1)

	isError := false

	if testLocNameType < LocNameType.None() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testLocNameType'= -1. However, NO ERROR WAS RETURNED!\n")
	}
}

func TestLocationNameType_02 (t *testing.T) {

	testLocNameType := LocationNameType(3)

	isError := false

	if testLocNameType > LocNameType.ConvertibleTimeZone() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testLocNameType'= +3. However, NO ERROR WAS RETURNED!\n")
	}
}

func TestLocationNameType_03(t *testing.T) {

	testLocNameType := LocNameType.NonConvertibleTimeZone()

	testString := testLocNameType.String()

	if testString != "NonConvertibleTimeZone" {
		t.Errorf("Error: Expected String() to return 'NonConvertibleTimeZone'\n" +
			"because testLocNameType='LocNameType.NonConvertibleTimeZone()'\n" +
			"However, testLocNameType='%v'\n", testString)
	}
}

func TestLocationNameType_04(t *testing.T) {

	textString := "ConvertibleTimeZone"

	actualValue, err :=
		LocationNameType(0).XParseString(textString, true)

	if err != nil {
		t.Errorf("Error returned by LocationNameType(0).XParseString(textString, true)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if LocNameType.ConvertibleTimeZone() != actualValue {
		t.Errorf("Error: Expected actualValue=LocNameType.ConvertibleTimeZone()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestLocationNameType_05(t *testing.T) {

	textString := "convertibletimezone"

	actualValue, err :=
		LocationNameType(0).XParseString(textString, false)

	if err != nil {
		t.Errorf("Error returned by LocationNameType(0).XParseString(textString, false)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if LocNameType.ConvertibleTimeZone() != actualValue {
		t.Errorf("Error: Expected actualValue=LocNameType.ConvertibleTimeZone()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestLocationNameType_06(t *testing.T) {

	testLocNameType := LocNameType.NonConvertibleTimeZone()

	actualValue := testLocNameType.XValue()

	if LocNameType.NonConvertibleTimeZone() != actualValue {
		t.Errorf("Error: Expected actualValue=LocNameType.NonConvertibleTimeZone()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestLocationNameType_07(t *testing.T) {

	testLocNameType := LocationNameType(-5)

	testString := testLocNameType.String()

	if testString != "" {
		t.Errorf("Error: Expected String() to return an Empty String\n" +
			"because testLocNameType='LocationNameType(-5)'\n" +
			"However, testLocNameType='%v'\n", testString)
	}

}

func TestLocationNameType_08(t *testing.T) {

	textString := "XRayzyxwVuTS"

	_, err :=
		LocationNameType(0).XParseString(textString, true)

	if err == nil {
		t.Error("Expected an error return from LocationNameType(0).XParseString(textString, true)\n" +
			"because 'textString' is an INVALID string value.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}
}

func TestLocationNameType_09(t *testing.T) {

	textString := "zvfuyhtsdbcrgq"

	_, err :=
		LocationNameType(0).XParseString(textString, false)

	if err == nil {
		t.Error("Expected an error return from LocationNameType(0).XParseString(textString, false)\n" +
			"because textString is INVALID!" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}
}

func TestLocationNameType_10(t *testing.T) {

	testLocNameType := LocNameType.NonConvertibleTimeZone()

	actualValue := testLocNameType.XValueInt()

	expectedValue := int(LocNameType.NonConvertibleTimeZone())

	if expectedValue != actualValue {
		t.Errorf("Error: Expected actualValue int='%v'\n" +
			"Instead, actualValue='%v'\n", expectedValue, actualValue)
	}
}

func TestTimeZoneCategory_01 (t *testing.T) {

	testTimeZoneCat := TimeZoneCategory(-1)

	isError := false

	if testTimeZoneCat < TzCat.None() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testTimeZoneCat'= -1. However, NO ERROR WAS RETURNED!\n")
	}
}

func TestTimeZoneCategory_02(t *testing.T) {
	testTimeZoneCat := TimeZoneCategory(3)

	isError := false

	if testTimeZoneCat > TzCat.UtcOffset() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testTimeZoneCat'= +3. However, NO ERROR WAS RETURNED!\n")
	}
}

func TestTimeZoneCategory_03(t *testing.T) {
	testTimeZoneCat := TzCat.UtcOffset()
	
	testString := testTimeZoneCat.String()
	
	if testString != "UtcOffset" {
		t.Errorf("Error: Expected String() to return 'UtcOffset'\n" +
			"because testTimeZoneCat='TzCat.UtcOffset()'\n" +
			"However, testTimeZoneCat='%v'\n", testString)
	}
}

func TestTimeZoneCategory_04(t *testing.T) {
	
	textString := "TextName"
		
	actualValue, err := 
		TimeZoneCategory(0).XParseString(textString, true)
	
	if err != nil {
		t.Errorf("Error returned by TimeZoneCategory(0).XParseString(textString, true)\n" +
			"Error='%v'\n", err.Error())
		return
	}
	
	if TzCat.TextName() != actualValue {
		t.Errorf("Error: Expected actualValue=TzCat.TextName()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
	
}

func TestTimeZoneCategory_05(t *testing.T) {
	
	textString := "textname"
		
	actualValue, err := 
		TimeZoneCategory(0).XParseString(textString, false)
	
	if err != nil {
		t.Errorf("Error returned by TimeZoneCategory(0).XParseString(textString, true)\n" +
			"Error='%v'\n", err.Error())
		return
	}
	
	if TzCat.TextName() != actualValue {
		t.Errorf("Error: Expected actualValue=TzCat.TextName()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
	
}

func TestTimeZoneCategory_06(t *testing.T) {

	testTimeZoneCat := TzCat.UtcOffset()

	actualValue := testTimeZoneCat.XValue()

	if TzCat.UtcOffset() != actualValue {
		t.Errorf("Error: Expected actualValue=TzCat.UtcOffset()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTimeZoneCategory_07(t *testing.T) {

	testTimeZoneCategory := TimeZoneCategory(-5)

	testString := testTimeZoneCategory.String()

	if testString != "" {
		t.Errorf("Error: Expected testTimeZoneCategory.String() to return an Empty String\n" +
			"because testTimeZoneCategory='TimeZoneCategory(-5)'\n" +
			"However, testTimeZoneCategory='%v'\n", testString)
	}

}

func TestTimeZoneCategory_08(t *testing.T) {

	textString := "XRayzyxwVuTS"

	_, err :=
		TimeZoneCategory(0).XParseString(textString, true)

	if err == nil {
		t.Error("Expected an error return from\n" +
			"TimeZoneCategory(0).XParseString(textString, true)\n" +
			"because 'textString' is an INVALID string value.\n" +
			"However, NO ERROR WAS RETURNED\n")
		return
	}
}

func TestTimeZoneCategory_09(t *testing.T) {

	textString := "zvfuyhtsdbcrgq"

	_, err :=
		TimeZoneCategory(0).XParseString(textString, false)

	if err == nil {
		t.Error("Expected an error return from\n" +
			"TimeZoneCategory(0).XParseString(textString, false)\n" +
			"because textString is INVALID!" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}
}

func TestTimeZoneCategory_10(t *testing.T) {

	testTimeZoneCat := TzCat.UtcOffset()

	actualValue := testTimeZoneCat.XValueInt()

	expectedValue := int(TzCat.UtcOffset())

	if expectedValue != actualValue {
		t.Errorf("Error: Expected actualValue='%v'\n" +
			"Instead, actualValue='%v'\n", expectedValue, actualValue)
	}
}

func TestTimeZoneClass_01(t *testing.T) {

	testTimeZoneClass := TimeZoneClass(-1)

	isError := false

	if testTimeZoneClass < TzClass.None() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testTimeZoneClass'= -1. However, NO ERROR WAS RETURNED!\n")
	}

}

func TestTimeZoneClass_02(t *testing.T) {

	testTimeZoneClass := TimeZoneClass(3)

	isError := false

	if testTimeZoneClass > TzClass.OriginalTimeZone() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testTimeZoneClass'= +3. However, NO ERROR WAS RETURNED!\n")
	}

}

func TestTimeZoneClass_03(t *testing.T) {
	
	testTimeZoneClass := TzClass.AlternateTimeZone()

	testString := testTimeZoneClass.String()

	if testString != "AlternateTimeZone" {
		t.Errorf("Error: Expected String() to return 'AlternateTimeZone'\n" +
			"because testTimeZoneClass='TzCat.AlternateTimeZone()'\n" +
			"However, testZoneTimeClass='%v'\n", testString)
	}
}

func TestTimeZoneClass_04(t *testing.T) {

	textString := "OriginalTimeZone"

	actualValue, err :=
		TimeZoneClass(0).XParseString(textString, true)

	if err != nil {
		t.Errorf("Error returned by TimeZoneClass(0).XParseString(textString, true)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if TzClass.OriginalTimeZone() != actualValue {
		t.Errorf("Error: Expected actualValue=TzClass.OriginalTimeZone()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}

}

func TestTimeZoneClass_05(t *testing.T) {

	textString := "alternatetimezone"

	actualValue, err :=
		TimeZoneClass(0).XParseString(textString, false)

	if err != nil {
		t.Errorf("Error returned by TimeZoneClass(0).XParseString(textString, true)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if TzClass.AlternateTimeZone() != actualValue {
		t.Errorf("Error: Expected actualValue=TzClass.AlternateTimeZone()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}

}

func TestTimeZoneClass_06(t *testing.T) {

	testTimeZoneClass := TzClass.OriginalTimeZone()

	actualValue := testTimeZoneClass.XValue()

	if TzClass.OriginalTimeZone() != actualValue {
		t.Errorf("Error: Expected actualValue=TzCat.UtcOffset()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}

}

func TestTimeZoneClass_07(t *testing.T) {

	testTimeZoneClass := TimeZoneClass(-5)

	testString := testTimeZoneClass.String()

	if testString != "" {
		t.Errorf("Error: Expected String() to return an Empty String\n" +
			"because testTimeZoneClass='TimeZoneClass(-5)'\n" +
			"However, testTimeZoneClass='%v'\n", testString)
	}

}

func TestTimeZoneClass_08(t *testing.T) {

	textString := "XRayzyxwVuTS"

	_, err :=
		TimeZoneClass(0).XParseString(textString, true)

	if err == nil {
		t.Error("Expected an error return from TimeZoneClass(0).XParseString(textString, true)\n" +
			"because 'textString' is an INVALID string value.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}
}

func TestTimeZoneClass_09(t *testing.T) {

	textString := "zvfuyhtsdbcrgq"

	_, err :=
		TimeZoneClass(0).XParseString(textString, false)

	if err == nil {
		t.Error("Expected an error return from TimeZoneClass(0).XParseString(textString, false)\n" +
			"because textString is INVALID!" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}
}

func TestTimeZoneClass_10(t *testing.T) {

	testTimeZoneClass := TzClass.OriginalTimeZone()

	actualValue := testTimeZoneClass.XValueInt()

	expectedValue := int(TzClass.OriginalTimeZone())

	if expectedValue != actualValue {
		t.Errorf("Error: Expected actualValue='%v'\n" +
			"Instead, actualValue='%v'\n", expectedValue, actualValue)
	}

}

func TestTimeZoneConversionType_01(t *testing.T) {

	testTimeZoneConversionType := TimeZoneConversionType(-1)

	isError := false

	if testTimeZoneConversionType < TimeZoneConversionType(0).None() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testTimeZoneConversionType'= -1.\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestTimeZoneConversionType_02(t *testing.T) {

	testTimeZoneConversionType := TimeZoneConversionType(3)

	isError := false

	if testTimeZoneConversionType > TimeZoneConversionType(0).Relative() {
		isError = true
	}

	if !isError {
		t.Error("ERROR: Expected isError='true' because\n" +
			"'testTimeZoneConversionType'= +3.\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestTimeZoneConversionType_03(t *testing.T) {

	testTimeZoneConversionType := TzConvertType.Relative()

	testString := testTimeZoneConversionType.String()

	if testString != "Relative" {
		t.Errorf("Error: Expected String() to return 'Relative'\n" +
			"because testTimeZoneConversionType='TzConvertType.Relative()'\n" +
			"However, testTimeZoneConversionType='%v'\n", testString)
	}
}

func TestTimeZoneConversionType_04(t *testing.T) {

	testTimeZoneConversionType := TimeZoneConversionType(0).Absolute()

	testString := testTimeZoneConversionType.String()

	if testString != "Absolute" {
		t.Errorf("Error: Expected String() to return 'Absolute'\n" +
			"because testTimeZoneConversionType='TimeZoneConversionType(0).Absolute()'\n" +
			"However, testTimeZoneConversionType='%v'\n", testString)
	}
}

func TestTimeZoneConversionType_05(t *testing.T) {

	textString := "Absolute"

	actualValue, err :=
		TimeZoneConversionType(0).XParseString(textString, true)

	if err != nil {
		t.Errorf("Error returned by TimeZoneConversionType(0).XParseString(textString, true)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if TimeZoneConversionType(0).Absolute() != actualValue {
		t.Errorf("Error: Expected actualValue=TimeZoneConversionType(0).Absolute()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTimeZoneConversionType_06(t *testing.T) {

	textString := "Relative"

	actualValue, err :=
		TimeZoneConversionType(0).XParseString(textString, true)

	if err != nil {
		t.Errorf("Error returned by TimeZoneConversionType(0).XParseString(textString, true)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if TimeZoneConversionType(0).Relative() != actualValue {
		t.Errorf("Error: Expected actualValue=TimeZoneConversionType(0).Relative()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTimeZoneConversionType_07(t *testing.T) {

	textString := "absolute"

	actualValue, err :=
		TimeZoneConversionType(0).XParseString(textString, false)

	if err != nil {
		t.Errorf("Error returned by TimeZoneConversionType(0).XParseString(textString, false)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if TimeZoneConversionType(0).Absolute() != actualValue {
		t.Errorf("Error: Expected actualValue=TimeZoneConversionType(0).Absolute()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTimeZoneConversionType_08(t *testing.T) {

	textString := "relative"

	actualValue, err :=
		TimeZoneConversionType(0).XParseString(textString, true)

	if err != nil {
		t.Errorf("Error returned by TimeZoneConversionType(0).XParseString(textString, true)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if TimeZoneConversionType(0).Relative() != actualValue {
		t.Errorf("Error: Expected actualValue=TimeZoneConversionType(0).Relative()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTimeZoneConversionType_09(t *testing.T) {

	testTimeZoneConversionType := TzConvertType.Absolute()

	actualValue := testTimeZoneConversionType.XValue()

	if TzConvertType.Absolute() != actualValue {
		t.Errorf("Error: Expected actualValue=TzConvertType.Absolute()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}

}

func TestTimeZoneConversionType_10(t *testing.T) {

	testTimeConversionType := TzConvertType.Relative()

	actualValue := testTimeConversionType.XValue()

	if TzConvertType.Relative() != actualValue {
		t.Errorf("Error: Expected actualValue=TzConvertType.Relative()\n" +
			"Instead, actualValue='%v'\n", actualValue.String())
	}
}

func TestTimeZoneConversionType_11(t *testing.T) {

	testTimeZoneConversionType := TimeZoneConversionType(-5)

	actualValue := testTimeZoneConversionType.String()

	if actualValue != "" {
		t.Errorf("Error: Expected String() to return an Empty String\n" +
			"because testTimeZoneConversionType='TimeZoneConversionType(-5)'\n" +
			"However, testTimeZoneConversionType='%v'\n", actualValue)
	}

}

func TestTimeZoneConversionType_12(t *testing.T) {

	textString := "XRayzyxwVuTS"

	_, err :=
		TimeZoneConversionType(0).XParseString(textString, true)

	if err == nil {
		t.Error("Expected an error return from TimeZoneConversionType(0).XParseString(textString, true)\n" +
			"because 'textString' is an INVALID string value.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}
}

func TestTimeZoneConversionType_13(t *testing.T) {

	textString := "xrayzyxwvuts"

	_, err :=
		TimeZoneConversionType(0).XParseString(textString, false)

	if err == nil {
		t.Error("Expected an error return from TimeZoneConversionType(0).XParseString(textString, false)\n" +
			"because 'textString' is an INVALID string value.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}
}

func TestTimeZoneConversionType_14(t *testing.T) {

	textString := "Xr"

	_, err :=
		TimeZoneConversionType(0).XParseString(textString, true)

	if err == nil {
		t.Error("Expected an error return from TimeZoneConversionType(0).XParseString(textString, true)\n" +
			"because 'textString' consists of only 2-characters.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}
}

func TestTimeZoneConversionType_15(t *testing.T) {

	textString := "xr"

	_, err :=
		TimeZoneConversionType(0).XParseString(textString, false)

	if err == nil {
		t.Error("Expected an error return from TimeZoneConversionType(0).XParseString(textString, false)\n" +
			"because 'textString' consists of only two lower case characters.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}
}
