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
		t.Errorf("Error returned by LocationNameType(0).XParseString(textString, true)\n" +
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

	testTimeZoneClass := TzClass.OriginalTimeZone()

	actualValue := testTimeZoneClass.XValueInt()

	expectedValue := int(TzClass.OriginalTimeZone())

	if expectedValue != actualValue {
		t.Errorf("Error: Expected actualValue='%v'\n" +
			"Instead, actualValue='%v'\n", expectedValue, actualValue)
	}

}
