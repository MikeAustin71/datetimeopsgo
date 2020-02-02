package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestDTimeMechanics_AddDateTimeByUtc_01(t *testing.T) {
	//t1str := "2020-03-07 03:00:00.000000000 -0500 EST"

	easternTzPtr, err := time.LoadLocation(TZones.America.New_York())

	if err != nil {
		t.Errorf("Error returned by " +
			"time.LoadLocation(dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1 := time.Date(
		2020,
		time.Month(3),
		7,
		3,
		0,
		0,
		0,
		easternTzPtr)

	expectedStr := "2020-03-08 04:00:00.000000000 -0400 EDT"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dtMech := DTimeMechanics{}

	t2 := dtMech.AddDateTimeByUtc(
		t1,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		0)

	t2OutStr := t2.Format(fmtStr)

	if expectedStr != t2OutStr {
		t.Errorf("Error: After addition of 1-Day,\n" +
			"expected result date time='%v'\n" +
			"Instead, result date time='%v'\n",
			expectedStr, t2OutStr)
	}
}

func TestDTimeMechanics_AddDateTime_01(t *testing.T) {
	//t1str := "2020-03-07 03:00:00.000000000 -0500 EST"

	easternTzPtr, err := time.LoadLocation(TZones.America.New_York())

	if err != nil {
		t.Errorf("Error returned by " +
			"time.LoadLocation(dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1 := time.Date(
		2020,
		time.Month(3),
		7,
		3,
		0,
		0,
		0,
		easternTzPtr)

	expectedStr := "2020-03-08 04:00:00.000000000 -0400 EDT"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dtMech := DTimeMechanics{}

	t2 := dtMech.AddDateTime(
		t1,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		0)

	t2OutStr := t2.Format(fmtStr)

	if expectedStr != t2OutStr {
		t.Errorf("Error: After addition of 1-Day,\n" +
			"expected result date time='%v'\n" +
			"Instead, result date time='%v'\n",
			expectedStr, t2OutStr)
	}
}


func TestDTimeMechanics_AbsoluteTimeToTimeZoneDefConversion_01(t *testing.T) {

	t1Str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	expectedStr := "2014-02-15 19:54:30.038175584 -0500 EST"

	dt1, err := time.Parse(fmtStr, t1Str)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, t1Str).\n" +
			"t1Str='%v'\n", t1Str)
		return
	}

	tzDefDto, err := TimeZoneDefinition{}.NewFromTimeZoneName(
		dt1,
		TZones.America.New_York(),
		TzConvertType.Relative())

	if err != nil {
		t.Errorf("Error returned by dt.TimeZoneDefinition{}.NewFromTimeZoneName(" +
			"dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dtMech := DTimeMechanics{}

	dt2, err := dtMech.AbsoluteTimeToTimeZoneDefConversion(dt1, tzDefDto)

	if err != nil {
		t.Errorf("Error returned by dtMech.AbsoluteTimeToTimeZoneDefConversion" +
			"(dt1, tzDefDto)\n" +
			"dt1='%v'\n" +
			"tzDefDto='%v'\n" +
			"Error='%v'\n",
			dt1.Format(fmtStr), tzDefDto.GetOriginalLocationName() ,err.Error())
		return
	}

	dt2Str := dt2.Format(fmtStr)

	if expectedStr != dt2Str {
		t.Errorf("Error: Expected actual date time='%v'\n" +
			"Instead, actual date time='%v'\n", expectedStr, dt2Str)
	}
}

func TestDTimeMechanics_GetTimeZoneFromDateTime_01 (t *testing.T) {

	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"
	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtMech := DTimeMechanics{}

	var tzSpec TimeZoneSpecification

	tzSpec,
		err = dtMech.GetTimeZoneFromDateTime(tPacificIn, "")

	if err != nil {
		t.Errorf("Error returned by dtMech.GetTimeZoneFromDateTime(tPacificIn, \"\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if TZones.America.Los_Angeles() != tzSpec.GetLocationName() {
		t.Errorf("Error: Expected Location Name='%v'\n" +
			"Instead, Location Name='%v'\n",
			TZones.America.Los_Angeles(), tzSpec.GetLocationName())
		return
	}

	if TZones.America.Los_Angeles() != tzSpec.GetTimeZoneName() {
		t.Errorf("Error: Expected Time Zone Name='%v'\n" +
			"Instead, Time Zone Name='%v'\n",
			TZones.America.Los_Angeles(), tzSpec.GetLocationName())
	}
}

func TestDTimeMechanics_GetTimeZoneFromDateTime_02 (t *testing.T) {

	pacificTime := "2017-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"
	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtMech := DTimeMechanics{}

	var tzSpec TimeZoneSpecification

	tzSpec,
		err = dtMech.GetTimeZoneFromDateTime(tPacificIn, "")

	if err != nil {
		t.Errorf("Error returned by dtMech.GetTimeZoneFromDateTime(tPacificIn, \"\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if TZones.America.Los_Angeles() != tzSpec.GetLocationName() {
		t.Errorf("Error: Expected Location Name='%v'\n" +
			"Instead, Location Name='%v'\n",
			TZones.America.Los_Angeles(), tzSpec.GetLocationName())
		return
	}

	if TZones.America.Los_Angeles() != tzSpec.GetTimeZoneName() {
		t.Errorf("Error: Expected Time Zone Name='%v'\n" +
			"Instead, Time Zone Name='%v'\n",
			TZones.America.Los_Angeles(), tzSpec.GetLocationName())
	}
}

func TestDTimeMechanics_GetTimeZoneFromDateTime_03 (t *testing.T) {

	pacificTime := "2017-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"
	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtMech := DTimeMechanics{}

	tPacificIn = time.Time{}

	_,
		err = dtMech.GetTimeZoneFromDateTime(tPacificIn, "")

	if err == nil {
		t.Error("Error: Expected an error return from dtMech.GetTimeZoneFromDateTime(tPacificIn, \"\")\n" +
			"because 'tPacificIn' has a ZERO value. However, NO ERROR WAS RETURNED!!\n")
	}
}

func TestDTimeMechanics_GetTimeZoneFromName_01(t *testing.T) {

	pacificTime := "2019-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtMech := DTimeMechanics{}

	var tzSpec TimeZoneSpecification

	tzSpec,
		err = dtMech.GetTimeZoneFromName(
			tPacificIn,
			TZones.America.New_York(),
			TzConvertType.Relative(),
			"")

	if err != nil {
		t.Errorf("Error returned by dtMech." +
			"GetTimeZoneFromName(tPacificIn,TZones.America.New_York()," +
			"TzConvertType.Relative(),\"\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var easternStdTimeZone *time.Location
	easternStdTimeZone,
	err = dtMech.LoadTzLocation(TZones.America.New_York(),"")

	if err != nil {
		t.Errorf("Error returned by DTimeUtility{}.LoadTzLocation(" +
			"TZones.America.New_York(),\"\")\n" +
			"Error='%v'\n", err.Error())
		return
	}
	expectedTime := tPacificIn.In(easternStdTimeZone).Format(fmtStr)

	timeActual := tzSpec.GetReferenceDateTime().Format(fmtStr)

	if expectedTime != timeActual {
		t.Errorf("Error: Expected Output Time= '%v'.\n" +
			"Instead, Output Time= '%v'\n",
			expectedTime, timeActual)
	}

}

func TestDTimeMechanics_GetTimeZoneFromName_02(t *testing.T) {

	pacificTime := "2019-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtMech := DTimeMechanics{}

	tPacificIn = time.Time{}

	_,
		err = dtMech.GetTimeZoneFromName(
			tPacificIn,
			TZones.America.New_York(),
			TzConvertType.Relative(),
			"")

	if err == nil {
		t.Error("Error: Expected an error return from dtMech.GetTimeZoneFromName(dateTime,...)\n" +
			"because 'dateTime' is a ZERO XValue! However, NO ERROR WAS RETURNED!\n")
	}
}

func TestDTimeMechanics_GetTimeZoneFromName_03(t *testing.T) {

	pacificTime := "2019-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtMech := DTimeMechanics{}
	timeZone := ""
	_,
		err = dtMech.GetTimeZoneFromName(
			tPacificIn,
			timeZone,
			TzConvertType.Relative(),
			"")

	if err == nil {
		t.Error("Error: Expected an error return from dtMech.GetTimeZoneFromName(timeZone,...)\n" +
			"because 'timeZone' is an EMPTY string! However, NO ERROR WAS RETURNED!\n")
	}
}

func TestDTimeMechanics_GetTimeZoneFromName_04(t *testing.T) {

	pacificTime := "2019-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtMech := DTimeMechanics{}
	_,
		err = dtMech.GetTimeZoneFromName(
			tPacificIn,
			TZones.America.New_York(),
			TzConvertType.None(),
			"")

	if err == nil {
		t.Error("Error: Expected an error return from dtMech.GetTimeZoneFromName(TimeZoneConversionType,...)\n" +
			"because 'TimeZoneConversionType' is equal to 'NONE'! However, NO ERROR WAS RETURNED!\n")
	}
}
