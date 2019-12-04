package zztests

import (
	"fmt"
	dt "github.com/MikeAustin71/datetimeopsgo/datetime"
	"testing"
	"time"
)

func TestMilitaryDateTzDto_01(t *testing.T) {

	tstr := "12/02/2019 22:05:00 -0600 CST"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	var testTime, expectedMilTime time.Time
	var err error
	var actualMilDateTimeGroup, expectedMilDateTimeGroup string
	var milDatTzDto dt.MilitaryDateTzDto
	var expectedMilTimeLoc *time.Location

	testTime, err = time.Parse(fmtstr, tstr)

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
			expectedMilTime.Format(fmtstr) ,err.Error())
		return
	}

	actualMilDateTimeGroup, err = milDatTzDto.GetOpenDateTimeGroup()

	if err != nil {
		t.Errorf("Error returned by milDatTzDto.GetOpenDateTimeGroup()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if expectedMilDateTimeGroup != actualMilDateTimeGroup {
		fmt.Printf("Error: Expected Military Date Time Group='%v'.\n" +
			"Actual Military Date Time Group='%v'\n" +
			"Military Time='%v'",
			expectedMilDateTimeGroup, actualMilDateTimeGroup, expectedMilTime.Format(fmtstr))
	}


}


