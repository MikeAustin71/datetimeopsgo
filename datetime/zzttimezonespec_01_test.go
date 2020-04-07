package datetime

import (
	"testing"
	"time"
)

func TestTimeZoneSpecification_01(t *testing.T) {

	dateTime := time.Date(
		2020,
		1,
		22,
		21,
		43,
		0,
		0,
		time.UTC)

	tzSpec := TimeZoneSpecification{}

	err := tzSpec.SetTimeZone(
		dateTime,
		"Z",
		"Zulu",
		"Test Zone",
		"Test Tag",
		TzClass.OriginalTimeZone(),
		"")

	if err != nil {
		t.Errorf("Error returned by tzSpec.SetTimeZone().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualLocNameType := tzSpec.GetLocationNameType()

	if LocNameType.ConvertibleTimeZone() != actualLocNameType {
		t.Errorf("Error: Expected actualLocNameType=" +
			"'LocNameType.ConvertibleTimeZone()'\n" +
			"Instead actualLocNameType='%v'\n", actualLocNameType.String())
		return
	}

	actualTimeZoneCategory := tzSpec.GetTimeZoneCategory()

	if TzCat.TextName() != actualTimeZoneCategory {
		t.Errorf("Error: Expected actualTimeZoneCategory=" +
			"'TzCat.TextName()'\n" +
			"Instead actualTimeZoneCategory='%v'\n", actualTimeZoneCategory.String())
		return
	}

	actualTzClass := tzSpec.GetTimeZoneClass()

	if TzClass.OriginalTimeZone() != actualTzClass {
		t.Errorf("Error: Expected actualTzClass='TzClass.OriginalTimeZone()'\n" +
			"Instead actualTzClass='%v'\n", actualTzClass.String())
		return
	}

	actualTzType := tzSpec.GetTimeZoneType()

	if TzType.Military() != actualTzType {
		t.Errorf("Error: Expected actualTzType='TzType.Military()'\n" +
			"Instead actualTzType='%v'\n", actualTzType.String())
		return
	}

	actualTzUtcOffsetStat := tzSpec.GetTimeZoneUtcOffsetStatus()

	if TzUtcStatus.Static() != actualTzUtcOffsetStat {
		t.Errorf("Error: Expected actualTzUtcOffsetStat='TzUtcStatus.Static()'\n" +
			"Instead actualTzUtcOffsetStat='%v'\n", actualTzUtcOffsetStat.String())
		return
	}

	actualDateTime := tzSpec.GetReferenceDateTime()

	if !dateTime.Equal(actualDateTime) {
		t.Errorf("Error: Expected actualDateTime='%v'\n" +
			"Instead actualDateTime='%v'\n",
			dateTime.Format(FmtDateTimeTzNanoYMD),
			actualDateTime.Format(FmtDateTimeTzNanoYMD))
	}
}

func TestTimeZoneSpecification_02 (t *testing.T) {

	locPtr, err := time.LoadLocation(TZones.Europe.London())

	dateTime := time.Date(
		2020,
		1,
		22,
		21,
		43,
		0,
		0,
		locPtr)

	zoneLabel := "Test Zone London"

	tagDesc := "Test Tag London"

	tzSpec := TimeZoneSpecification{}

	err = tzSpec.SetTimeZone(
		dateTime,
		"",
		"",
		zoneLabel,
		tagDesc,
		TzClass.OriginalTimeZone(),
		"")

	if err != nil {
		t.Errorf("Error returned by tzSpec.SetTimeZone().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualZoneLabel := tzSpec.GetZoneLabel()

	if zoneLabel != actualZoneLabel {
		t.Errorf("Error: Expected actualZoneLabel='%v'\n" +
			"Instead actualZoneLabel='%v'\n",
			zoneLabel, actualZoneLabel)
		return
	}

	actualTagDesc := tzSpec.GetTagDescription()

	if tagDesc != actualTagDesc {
		t.Errorf("Error: Expected actualTagDesc='%v'\n" +
			"Instead actualTagDesc='%v'\n",
			tagDesc, actualTagDesc)
		return
	}

	actualLocNameType := tzSpec.GetLocationNameType()

	if LocNameType.ConvertibleTimeZone() != actualLocNameType {
		t.Errorf("Error: Expected actualLocNameType=" +
			"'LocNameType.ConvertibleTimeZone()'\n" +
			"Instead actualLocNameType='%v'\n", actualLocNameType.String())
		return
	}

	actualTimeZoneCategory := tzSpec.GetTimeZoneCategory()

	if TzCat.TextName() != actualTimeZoneCategory {
		t.Errorf("Error: Expected actualTimeZoneCategory=" +
			"'TzCat.TextName()'\n" +
			"Instead actualTimeZoneCategory='%v'\n", actualTimeZoneCategory.String())
		return
	}

	actualTzClass := tzSpec.GetTimeZoneClass()

	if TzClass.OriginalTimeZone() != actualTzClass {
		t.Errorf("Error: Expected actualTzClass='TzClass.OriginalTimeZone()'\n" +
			"Instead actualTzClass='%v'\n", actualTzClass.String())
		return
	}

	actualTzType := tzSpec.GetTimeZoneType()

	if TzType.Iana() != actualTzType {
		t.Errorf("Error: Expected actualTzType='TzType.Iana()'\n" +
			"Instead actualTzType='%v'\n", actualTzType.String())
		return
	}

	actualTzUtcOffsetStat := tzSpec.GetTimeZoneUtcOffsetStatus()

	if TzUtcStatus.Variable() != actualTzUtcOffsetStat {
		t.Errorf("Error: Expected actualTzUtcOffsetStat='TzUtcStatus.Variable()'\n" +
			"Instead actualTzUtcOffsetStat='%v'\n", actualTzUtcOffsetStat.String())
		return
	}

	actualDateTime := tzSpec.GetReferenceDateTime()

	if !dateTime.Equal(actualDateTime) {
		t.Errorf("Error: Expected actualDateTime='%v'\n" +
			"Instead actualDateTime='%v'\n",
			dateTime.Format(FmtDateTimeTzNanoYMD),
			actualDateTime.Format(FmtDateTimeTzNanoYMD))
	}
}


func TestTimeZoneSpecification_03 (t *testing.T) {

	locPtr, err := time.LoadLocation(TZones.Asia.Vladivostok())

	dateTime := time.Date(
		2020,
		1,
		22,
		21,
		43,
		0,
		0,
		locPtr)

	zoneLabel := "Test Zone Vladivostok"

	tagDesc := "Test Tag Vladivostok"

	tzSpec := TimeZoneSpecification{}

	err = tzSpec.SetTimeZone(
		dateTime,
		"",
		"",
		zoneLabel,
		tagDesc,
		TzClass.OriginalTimeZone(),
		"")

	if err != nil {
		t.Errorf("Error returned by tzSpec.SetTimeZone().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualZoneLabel := tzSpec.GetZoneLabel()

	if zoneLabel != actualZoneLabel {
		t.Errorf("Error: Expected actualZoneLabel='%v'\n" +
			"Instead actualZoneLabel='%v'\n",
			zoneLabel, actualZoneLabel)
		return
	}

	actualTagDesc := tzSpec.GetTagDescription()

	if tagDesc != actualTagDesc {
		t.Errorf("Error: Expected actualTagDesc='%v'\n" +
			"Instead actualTagDesc='%v'\n",
			tagDesc, actualTagDesc)
		return
	}

	actualLocNameType := tzSpec.GetLocationNameType()

	if LocNameType.ConvertibleTimeZone() != actualLocNameType {
		t.Errorf("Error: Expected actualLocNameType=" +
			"'LocNameType.ConvertibleTimeZone()'\n" +
			"Instead actualLocNameType='%v'\n", actualLocNameType.String())
		return
	}

	actualTimeZoneCategory := tzSpec.GetTimeZoneCategory()

	if TzCat.TextName() != actualTimeZoneCategory {
		t.Errorf("Error: Expected actualTimeZoneCategory=" +
			"'TzCat.TextName()'\n" +
			"Instead actualTimeZoneCategory='%v'\n", actualTimeZoneCategory.String())
		return
	}

	actualTzClass := tzSpec.GetTimeZoneClass()

	if TzClass.OriginalTimeZone() != actualTzClass {
		t.Errorf("Error: Expected actualTzClass='TzClass.OriginalTimeZone()'\n" +
			"Instead actualTzClass='%v'\n", actualTzClass.String())
		return
	}

	actualTzType := tzSpec.GetTimeZoneType()

	if TzType.Iana() != actualTzType {
		t.Errorf("Error: Expected actualTzType='TzType.Iana()'\n" +
			"Instead actualTzType='%v'\n", actualTzType.String())
		return
	}

	actualTzUtcOffsetStat := tzSpec.GetTimeZoneUtcOffsetStatus()

	if TzUtcStatus.Static() != actualTzUtcOffsetStat {
		t.Errorf("Error: Expected actualTzUtcOffsetStat='TzUtcStatus.Static()'\n" +
			"Instead actualTzUtcOffsetStat='%v'\n", actualTzUtcOffsetStat.String())
		return
	}

	actualDateTime := tzSpec.GetReferenceDateTime()

	if !dateTime.Equal(actualDateTime) {
		t.Errorf("Error: Expected actualDateTime='%v'\n" +
			"Instead actualDateTime='%v'\n",
			dateTime.Format(FmtDateTimeTzNanoYMD),
			actualDateTime.Format(FmtDateTimeTzNanoYMD))
	}
}

func TestTimeZoneSpecification_04 (t *testing.T) {

	dtStr := "12/30/2019 09:00:00.000000000 +0700 +07"

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	dateTime, err := time.Parse(fmtStr, dtStr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, dtStr)\n" +
			"fmtStr='%v'\n" +
			"dtStr='%v'\n" +
			"Error='%v'\n", fmtStr, dtStr, err.Error())
		return
	}

	zoneLabel := "Test Zone +07"

	tagDesc := "Test Tag +07"

	tzSpec := TimeZoneSpecification{}

	err = tzSpec.SetTimeZone(
		dateTime,
		"",
		"",
		zoneLabel,
		tagDesc,
		TzClass.OriginalTimeZone(),
		"")

	if err != nil {
		t.Errorf("Error returned by tzSpec.SetTimeZone().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualZoneLabel := tzSpec.GetZoneLabel()

	if zoneLabel != actualZoneLabel {
		t.Errorf("Error: Expected actualZoneLabel='%v'\n" +
			"Instead actualZoneLabel='%v'\n",
			zoneLabel, actualZoneLabel)
		return
	}

	actualTagDesc := tzSpec.GetTagDescription()

	if tagDesc != actualTagDesc {
		t.Errorf("Error: Expected actualTagDesc='%v'\n" +
			"Instead actualTagDesc='%v'\n",
			tagDesc, actualTagDesc)
		return
	}

	actualLocNameType := tzSpec.GetLocationNameType()

	if LocNameType.NonConvertibleTimeZone() != actualLocNameType {
		t.Errorf("Error: Expected actualLocNameType=" +
			"'LocNameType.NonConvertibleTimeZone()'\n" +
			"Instead actualLocNameType='%v'\n", actualLocNameType.String())
		return
	}

	actualTimeZoneCategory := tzSpec.GetTimeZoneCategory()

	if TzCat.UtcOffset() != actualTimeZoneCategory {
		t.Errorf("Error: Expected actualTimeZoneCategory=" +
			"'TzCat.UtcOffset()'\n" +
			"Instead actualTimeZoneCategory='%v'\n", actualTimeZoneCategory.String())
		return
	}

	actualTzClass := tzSpec.GetTimeZoneClass()

	if TzClass.OriginalTimeZone() != actualTzClass {
		t.Errorf("Error: Expected actualTzClass='TzClass.OriginalTimeZone()'\n" +
			"Instead actualTzClass='%v'\n", actualTzClass.String())
		return
	}

	actualTzType := tzSpec.GetTimeZoneType()

	if TzType.Iana() != actualTzType {
		t.Errorf("Error: Expected actualTzType='TzType.Iana()'\n" +
			"Instead actualTzType='%v'\n", actualTzType.String())
		return
	}

	actualTzUtcOffsetStat := tzSpec.GetTimeZoneUtcOffsetStatus()

	if TzUtcStatus.Static() != actualTzUtcOffsetStat {
		t.Errorf("Error: Expected actualTzUtcOffsetStat='TzUtcStatus.Static()'\n" +
			"Instead actualTzUtcOffsetStat='%v'\n", actualTzUtcOffsetStat.String())
		return
	}

	actualDateTime := tzSpec.GetReferenceDateTime()

	if !dateTime.Equal(actualDateTime) {
		t.Errorf("Error: Expected actualDateTime='%v'\n" +
			"Instead actualDateTime='%v'\n",
			dateTime.Format(FmtDateTimeTzNanoYMD),
			actualDateTime.Format(FmtDateTimeTzNanoYMD))
	}
}