package datetime

import (
	"strings"
	"testing"
	"time"
)

func TestMilitaryDateTzDto_01(t *testing.T) {

	tstr := "12/02/2019 22:05:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	var testTime, expectedMilTime time.Time
	var err error
	var actualMilDateTimeGroup, expectedMilDateTimeGroup string
	var milDatTzDto MilitaryDateTzDto
	var expectedMilTimeLoc *time.Location

	testTime, err = time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
		return
	}

	expectedMilTimeLoc, err = time.LoadLocation(TZones.Military.Quebec())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Military.Quebec())\n" +
			"TZones.Military.Quebec()='%v'\n" +
			"Error='%v'\n", TZones.Military.Quebec(), err.Error())
	}

	expectedMilTime = testTime.In(expectedMilTimeLoc)

	milDatTzDto, err = MilitaryDateTzDto{}.New(testTime, "Q")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.NewStartEndTimes(testTime, \"Q\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedMilDateTimeGroup, err = DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)

	if err != nil {
		t.Errorf("Error returned by DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)\n" +
			"expectedMilTime='%v'\n" +
			"Error='%v'\n",
			expectedMilTime.Format(fmtStr) ,err.Error())
		return
	}

	actualMilDateTimeGroup, err = milDatTzDto.GetOpenDateTimeGroup()

	if err != nil {
		t.Errorf("Error returned by milDatTzDto.GetOpenDateTimeGroup()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if expectedMilDateTimeGroup != actualMilDateTimeGroup {
		t.Errorf("Error: Expected Military Date Time Group='%v'.\n" +
			"Actual Military Date Time Group='%v'\n" +
			"Military Time='%v'",
			expectedMilDateTimeGroup, actualMilDateTimeGroup, expectedMilTime.Format(fmtStr))
	}

}

func TestMilitaryDateTzDto_02(t *testing.T) {

	tstr := "12/02/2019 22:05:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	var testTime, expectedMilTime time.Time
	var err error
	var actualMilDateTimeGroup, expectedMilDateTimeGroup string
	var milDatTzDto MilitaryDateTzDto
	var expectedMilTimeLoc *time.Location

	testTime, err = time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
		return
	}

	expectedMilTimeLoc, err = time.LoadLocation(TZones.Military.Quebec())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Military.Quebec())\n" +
			"TZones.Military.Quebec()='%v'\n" +
			"Error='%v'\n", TZones.Military.Quebec(), err.Error())
		return
	}

	expectedMilTime = testTime.In(expectedMilTimeLoc)

	milDatTzDto, err = MilitaryDateTzDto{}.New(testTime, "Q")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.NewStartEndTimes(testTime, \"Q\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedMilDateTimeGroup, err = DtMgr{}.GetMilitaryCompactDateTimeGroup(expectedMilTime)

	if err != nil {
		t.Errorf("Error returned by DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)\n" +
			"expectedMilTime='%v'\n" +
			"Error='%v'\n",
			expectedMilTime.Format(fmtStr) ,err.Error())
		return
	}

	actualMilDateTimeGroup, err = milDatTzDto.GetCompactDateTimeGroup()

	if err != nil {
		t.Errorf("Error returned by milDatTzDto.GetOpenDateTimeGroup()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if expectedMilDateTimeGroup != actualMilDateTimeGroup {
		t.Errorf("Error: Expected Military Date Time Group='%v'.\n" +
			"Actual Military Date Time Group='%v'\n" +
			"Military Time='%v'",
			expectedMilDateTimeGroup, actualMilDateTimeGroup, expectedMilTime.Format(fmtStr))
	}
}

func TestMilitaryDateTzDto_New_01(t *testing.T) {

	tstr := "12/04/2019 03:12:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	var testTime, utcTime, expectedMilTime time.Time
	var err error
	var actualMilDateTimeGroup, expectedMilDateTimeGroup string
	var milDatTzDto MilitaryDateTzDto
	var expectedMilTimeLoc, utcLoc *time.Location

	testTime, err = time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
	}

	utcLoc, err = time.LoadLocation(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(time.LoadLocation(TZones.UTC()))\n" +
			"TZones.Military.Quebec()='%v'\n" +
			"Error='%v'\n", TZones.Military.Quebec(), err.Error())
	}

	utcTime = testTime.In(utcLoc)
	
	expectedMilTimeLoc, err = time.LoadLocation(TZones.Military.Zulu())
	
	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(time.LoadLocation(TZones.Military.Zulu())\n" +
			"TZones.Military.Zulu()='%v'\n" +
			"Error='%v'\n", TZones.Military.Zulu(), err.Error())
	}

	expectedMilTime = testTime.In(expectedMilTimeLoc)

	milDatTzDto, err = MilitaryDateTzDto{}.New(utcTime, "Z")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.NewStartEndTimes(testTime, \"Z\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedMilDateTimeGroup, err = DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)

	if err != nil {
		t.Errorf("Error returned by DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)\n" +
			"expectedMilTime='%v'\n" +
			"Error='%v'\n",
			expectedMilTime.Format(fmtStr) ,err.Error())
		return
	}

	actualMilDateTimeGroup, err = milDatTzDto.GetCompactDateTimeGroup()

	if err != nil {
		t.Errorf("Error returned by milDatTzDto.GetOpenDateTimeGroup()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if expectedMilDateTimeGroup != actualMilDateTimeGroup {
		t.Errorf("Error: Expected Military Date Time Group='%v'.\n" +
			"Actual Military Date Time Group='%v'\n" +
			"Military Time='%v'",
			expectedMilDateTimeGroup, actualMilDateTimeGroup, expectedMilTime.Format(fmtStr))
	}
	
	utcTimeStr := utcTime.Format(fmtStr)
	
	actualTimeStr := milDatTzDto.DateTime.Format(fmtStr)
	
	if utcTimeStr != actualTimeStr {
		t.Errorf("Error: Expected Military Civilian Date Time Format= '%v'.\n" +
			"Instead, Civilian Date Time Format='%v'\n", utcTimeStr, actualTimeStr)
	}
}

func TestMilitaryDateTzDto_NewNow_01(t *testing.T) {

	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	tNow := time.Now().Local()

	tNowStr := tNow.Format(fmtStr)

	tNowArray := strings.Split(tNowStr, " ")

	if len(tNowArray) != 4 {
		t.Errorf("Error: Expected length of tNowArray=='4'.\n" +
			"Instead, length of tNowArray='%v'\n", len(tNowArray))
		return
	}

	milTzDat := MilitaryTimeZoneData{}

	militaryTz, ok := milTzDat.UtcOffsetToMilitaryTimeZone(tNowArray[2])

	if !ok {
		t.Errorf("Error: militaryUTCToTzMap[tNowArray[2]] FAILED!\n" +
			"tNowArray[2]='%v'\n", tNowArray[2])
		return
	}

	tNow = time.Now().Local()

	milDatTzDto, err := MilitaryDateTzDto{}.NewNow(militaryTz)

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.NewNow(militaryTz)\n" +
			"militaryTz='%v'\n" +
			"Error='%v'\n", militaryTz, err.Error())
		return
	}

	tDuration := milDatTzDto.DateTime.Sub(tNow)

	if int64(tDuration) > (SecondNanoseconds * 3) {
		t.Errorf("Error: Duration from Local Now is greater than 3-seconds.\n" +
			"Duration='%v'\n", tDuration)
	}

}

func TestMilitaryDateTzDto_NewNowZulu_01(t *testing.T) {

	// fmtStr := "01/02/2006 15:04:05 -0700 MST"

	tNow := time.Now().UTC()

	milDatTzDto, err := MilitaryDateTzDto{}.NewNowZulu()

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.NewNow(militaryTz)\n" +
			"militaryTz='ZULU'\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tDuration := milDatTzDto.DateTime.Sub(tNow)

	if int64(tDuration) > (SecondNanoseconds * 3) {
		t.Errorf("Error: Duration from Local Now is greater than 3-seconds.\n" +
			"Duration='%v'\n", tDuration)
	}

}

func TestMilitaryDateTzDto_NewFromDateTzDto_01(t *testing.T) {

	tstr := "12/06/2019 03:12:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	testTime, err := time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr)\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
	}

	var dateTzDto DateTzDto

	dateTzDto, err = DateTzDto{}.NewDateTime(testTime, fmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewStartEndTimes(testTime, fmtStr)\n" +
			"testTime='%v'\n" +
			"Error='%v'\n", testTime.Format(fmtStr), err.Error())
	}

	var milTzDto MilitaryDateTzDto

	milTzDto, err = MilitaryDateTzDto{}.NewFromDateTzDto(dateTzDto)

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.NewFromDateTzDto(dateTzDto)\n" +
			"dateTzDto.DateTime='%v'\n" +
			"Error='%v'\n", dateTzDto.GetDateTimeValue().Format(fmtStr), err.Error())
		return
	}

	// 630pm on January 6th, 2012 in Fayetteville NC would read '061830RJAN12'
	// "12/06/2019 03:12:00 -0600 CST"
	expectedCompactDateGroup := "061512SDEC19"

	var actualCompactDateGroup string

	actualCompactDateGroup, err = milTzDto.GetCompactDateTimeGroup()

	if err != nil {
		t.Errorf("Error returned by milTzDto.GetCompactDateTimeGroup()\n" +
			"Error='%v'\n", err.Error())
	}

	if expectedCompactDateGroup != actualCompactDateGroup {
		t.Errorf("Error: Expected Compact Date Group='%v'.\n" +
			"Instead, Compact Date Group='%v'\n", expectedCompactDateGroup, actualCompactDateGroup)
	}

}

func TestMilitaryDateTzDto_GeoLocation_01(t *testing.T) {

	tstr := "12/04/2019 03:12:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	testTime, err := time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr)\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
	}

	milDatTzDto, err := MilitaryDateTzDto{}.New(testTime, "Bravo")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.NewStartEndTimes(testTime, \"B\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	milTzDat := MilitaryTimeZoneData{}

	geoLoc, ok := milTzDat.MilitaryTzToLocation("Bravo")

	if !ok {
		t.Error("Error: Military Time Zone To Location Look-up Failed!\n" +
			"Military Time Zone Text Name='Bravo'\n")
	}

	if geoLoc != milDatTzDto.GeoLocationDesc {
		t.Errorf("Error: Expected that Bravo Time Zone Geographic Location='%v'\n." +
			"Instead, the Bravo Geographic Location='%v'.\n", geoLoc, milDatTzDto.GeoLocationDesc)
	}
}

func TestMilitaryDateTzDto_UtcOffset_01(t *testing.T) {

	tstr := "12/04/2019 03:12:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	testTime, err := time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
		return
	}

	testTimeStr := testTime.Format(fmtStr)

	testTimeStrAry := strings.Split(testTimeStr, " ")

	expectedUtcOffset := testTimeStrAry[2]

	milDatTzDto, err := MilitaryDateTzDto{}.New(testTime, "Sierra")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.NewStartEndTimes(testTime, \"Sierra\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if expectedUtcOffset != milDatTzDto.UtcOffset {
		t.Errorf("Error: Expected Military Sierra Time Zone UTC Offset='%v'.\n" +
			"Instead, milDatTzDto.UtcOffset='%v'\n", expectedUtcOffset, milDatTzDto.UtcOffset)
	}

}

func TestMilitaryDateTzDto_EquivalentIanaTime_01(t *testing.T) {

	tstr := "12/04/2019 03:12:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	testTime, err := time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
		return
	}

	eastTzLoc, err:= time.LoadLocation(TZones.America.New_York())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(time.LoadLocation(TZones.America.New_York())\n" +
			"TZones.Military.Quebec()='%v'\n" +
			"Error='%v'\n", TZones.Military.Quebec(), err.Error())
		return
	}

	testTime = testTime.In(eastTzLoc)

	milDatTzDto, err := MilitaryDateTzDto{}.New(testTime, "Romeo")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.NewStartEndTimes(testTime, \"Romeo\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if milDatTzDto.EquivalentIanaTimeZone.GetOriginalLocationName() != TZones.Etc.GMTPlus05() {
		t.Errorf("Error: Expected Military Sierra Time Zone UTC Offset='%v'.\n" +
			"Instead, milDatTzDto.UtcOffset='%v'\n",
			milDatTzDto.EquivalentIanaTimeZone.GetOriginalLocationName(), TZones.Etc.GMTPlus05())
	}
}


func TestMilitaryDateTzDto_TimeZoneName01(t *testing.T) {

	tstr := "12/02/2019 22:05:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	var testTime time.Time
	var err error
	var milDatTzDto MilitaryDateTzDto

	testTime, err = time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
		return
	}

	milDatTzDto, err = MilitaryDateTzDto{}.New(testTime, "D")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.NewStartEndTimes(testTime, \"D\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if milDatTzDto.MilitaryTzTextName != "Delta" {
		t.Errorf("Error: Expected milDatTzDto.MilitaryTzTextName='Delta'.\n" +
			"Instead, milDatTzDto.MilitaryTzTextName='%v'\n", milDatTzDto.MilitaryTzTextName)
	}
}

