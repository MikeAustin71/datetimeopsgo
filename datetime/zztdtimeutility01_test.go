package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestDTimeUtility_AbsoluteTimeToTimeZoneDtoConversion_01(t *testing.T) {

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

	dtUtil := DTimeUtility{}

	dt2, err := dtUtil.AbsoluteTimeToTimeZoneDtoConversion(dt1, tzDefDto)

	if err != nil {
		t.Errorf("Error returned by dtUtil.AbsoluteTimeToTimeZoneDtoConversion" +
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

func TestDTimeUtility_GetTimeZoneFromDateTime_01 (t *testing.T) {

	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"
	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtUtil := DTimeUtility{}

	var tzSpec TimeZoneSpecification

	tzSpec,
		err = dtUtil.GetTimeZoneFromDateTime(tPacificIn, "")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GetTimeZoneFromDateTime(tPacificIn, \"\")\n" +
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

func TestDTimeUtility_GetTimeZoneFromDateTime_02 (t *testing.T) {

	pacificTime := "2017-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"
	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtUtil := DTimeUtility{}

	var tzSpec TimeZoneSpecification

	tzSpec,
		err = dtUtil.GetTimeZoneFromDateTime(tPacificIn, "")

	if err != nil {
		t.Errorf("Error returned by dtUtil.GetTimeZoneFromDateTime(tPacificIn, \"\")\n" +
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

func TestDTimeUtility_GetTimeZoneFromDateTime_03 (t *testing.T) {

	pacificTime := "2017-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"
	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtUtil := DTimeUtility{}

	tPacificIn = time.Time{}

	_,
		err = dtUtil.GetTimeZoneFromDateTime(tPacificIn, "")

	if err == nil {
		t.Error("Error: Expected an error return from dtUtil.GetTimeZoneFromDateTime(tPacificIn, \"\")\n" +
			"because 'tPacificIn' has a ZERO value. However, NO ERROR WAS RETURNED!!\n")
	}
}

func TestDTimeUtility_GetTimeZoneFromName_01(t *testing.T) {

	pacificTime := "2019-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtUtil := DTimeUtility{}

	var tzSpec TimeZoneSpecification

	tzSpec,
		err = dtUtil.GetTimeZoneFromName(
			tPacificIn,
			TZones.America.New_York(),
			TzConvertType.Relative(),
			"")

	if err != nil {
		t.Errorf("Error returned by dtUtil." +
			"GetTimeZoneFromName(tPacificIn,TZones.America.New_York()," +
			"TzConvertType.Relative(),\"\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var easternStdTimeZone *time.Location

	easternStdTimeZone,
	err = DTimeUtility{}.LoadTzLocation(TZones.America.New_York(),"")

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

func TestDTimeUtility_GetTimeZoneFromName_02(t *testing.T) {

	pacificTime := "2019-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtUtil := DTimeUtility{}

	tPacificIn = time.Time{}

	_,
		err = dtUtil.GetTimeZoneFromName(
			tPacificIn,
			TZones.America.New_York(),
			TzConvertType.Relative(),
			"")

	if err == nil {
		t.Error("Error: Expected an error return from dtUtil.GetTimeZoneFromName(dateTime,...)\n" +
			"because 'dateTime' is a ZERO Value! However, NO ERROR WAS RETURNED!\n")
	}
}

func TestDTimeUtility_GetTimeZoneFromName_03(t *testing.T) {

	pacificTime := "2019-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtUtil := DTimeUtility{}
	timeZone := ""
	_,
		err = dtUtil.GetTimeZoneFromName(
			tPacificIn,
			timeZone,
			TzConvertType.Relative(),
			"")

	if err == nil {
		t.Error("Error: Expected an error return from dtUtil.GetTimeZoneFromName(timeZone,...)\n" +
			"because 'timeZone' is an EMPTY string! However, NO ERROR WAS RETURNED!\n")
	}
}

func TestDTimeUtility_GetTimeZoneFromName_04(t *testing.T) {

	pacificTime := "2019-12-15 17:54:30 -0800 PST"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	tPacificIn, err := time.Parse(fmtStr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	dtUtil := DTimeUtility{}
	_,
		err = dtUtil.GetTimeZoneFromName(
			tPacificIn,
			TZones.America.New_York(),
			TzConvertType.None(),
			"")

	if err == nil {
		t.Error("Error: Expected an error return from dtUtil.GetTimeZoneFromName(TimeZoneConversionType,...)\n" +
			"because 'TimeZoneConversionType' is equal to 'NONE'! However, NO ERROR WAS RETURNED!\n")
	}
}
