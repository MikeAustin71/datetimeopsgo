package zztests

import (
	dt "github.com/MikeAustin71/datetimeopsgo/datetime"
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
	var milDatTzDto dt.MilitaryDateTzDto
	var expectedMilTimeLoc *time.Location

	testTime, err = time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
		return
	}

	expectedMilTimeLoc, err = time.LoadLocation(dt.TZones.Military.Quebec())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(dt.TZones.Military.Quebec())\n" +
			"dt.TZones.Military.Quebec()='%v'\n" +
			"Error='%v'\n", dt.TZones.Military.Quebec(), err.Error())
	}

	expectedMilTime = testTime.In(expectedMilTimeLoc)

	milDatTzDto, err = dt.MilitaryDateTzDto{}.New(testTime, "Q")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.New(testTime, \"Q\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedMilDateTimeGroup, err = dt.DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)

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
	var milDatTzDto dt.MilitaryDateTzDto
	var expectedMilTimeLoc *time.Location

	testTime, err = time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
		return
	}

	expectedMilTimeLoc, err = time.LoadLocation(dt.TZones.Military.Quebec())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(dt.TZones.Military.Quebec())\n" +
			"dt.TZones.Military.Quebec()='%v'\n" +
			"Error='%v'\n", dt.TZones.Military.Quebec(), err.Error())
		return
	}

	expectedMilTime = testTime.In(expectedMilTimeLoc)

	milDatTzDto, err = dt.MilitaryDateTzDto{}.New(testTime, "Q")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.New(testTime, \"Q\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedMilDateTimeGroup, err = dt.DtMgr{}.GetMilitaryCompactDateTimeGroup(expectedMilTime)

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
	var milDatTzDto dt.MilitaryDateTzDto
	var expectedMilTimeLoc, utcLoc *time.Location

	testTime, err = time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
	}

	utcLoc, err = time.LoadLocation(dt.TZones.UTC())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(time.LoadLocation(dt.TZones.UTC()))\n" +
			"dt.TZones.Military.Quebec()='%v'\n" +
			"Error='%v'\n", dt.TZones.Military.Quebec(), err.Error())
	}

	utcTime = testTime.In(utcLoc)
	
	expectedMilTimeLoc, err = time.LoadLocation(dt.TZones.Military.Zulu())
	
	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(time.LoadLocation(dt.TZones.Military.Zulu())\n" +
			"dt.TZones.Military.Zulu()='%v'\n" +
			"Error='%v'\n", dt.TZones.Military.Zulu(), err.Error())
	}

	expectedMilTime = testTime.In(expectedMilTimeLoc)

	milDatTzDto, err = dt.MilitaryDateTzDto{}.New(utcTime, "Z")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.New(testTime, \"Z\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedMilDateTimeGroup, err = dt.DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)

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

func TestMilitaryDateTzDto_GeoLocation_01(t *testing.T) {

	tstr := "12/04/2019 03:12:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	testTime, err := time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
	}

	milDatTzDto, err := dt.MilitaryDateTzDto{}.New(testTime, "Bravo")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.New(testTime, \"B\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	geoLoc := dt.MilitaryTzLocationMap["Bravo"]

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

	milDatTzDto, err := dt.MilitaryDateTzDto{}.New(testTime, "Sierra")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.New(testTime, \"Sierra\")\n" +
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

	eastTzLoc, err:= time.LoadLocation(dt.TZones.America.New_York())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(time.LoadLocation(dt.TZones.America.New_York())\n" +
			"dt.TZones.Military.Quebec()='%v'\n" +
			"Error='%v'\n", dt.TZones.Military.Quebec(), err.Error())
		return
	}

	testTime = testTime.In(eastTzLoc)

	milDatTzDto, err := dt.MilitaryDateTzDto{}.New(testTime, "Romeo")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.New(testTime, \"Romeo\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if milDatTzDto.EquivalentIanaTimeZone.LocationName != dt.TZones.Etc.GMTPlus05() {
		t.Errorf("Error: Expected Military Sierra Time Zone UTC Offset='%v'.\n" +
			"Instead, milDatTzDto.UtcOffset='%v'\n",
			milDatTzDto.EquivalentIanaTimeZone.LocationName, dt.TZones.Etc.GMTPlus05())
	}
}


func TestMilitaryDateTzDto_TimeZoneName01(t *testing.T) {

	tstr := "12/02/2019 22:05:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	var testTime time.Time
	var err error
	var milDatTzDto dt.MilitaryDateTzDto

	testTime, err = time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
		return
	}

	milDatTzDto, err = dt.MilitaryDateTzDto{}.New(testTime, "D")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.New(testTime, \"D\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if milDatTzDto.MilitaryTzTextName != "Delta" {
		t.Errorf("Error: Expected milDatTzDto.MilitaryTzTextName='Delta'.\n" +
			"Instead, milDatTzDto.MilitaryTzTextName='%v'\n", milDatTzDto.MilitaryTzTextName)
	}
}

