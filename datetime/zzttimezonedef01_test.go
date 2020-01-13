package datetime

import (
	"testing"
	"time"
)

func TestTimeZoneDefDto_New_01(t *testing.T) {

	usPacificLoc, _ := time.LoadLocation(TZones.US.Pacific())

	tUsPacific :=
		time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "US/Pacific"
	descStr := "US-Pacific"

	tzDef, err := TimeZoneDefinition{}.New(tUsPacific)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefinition{}.New(tUsPacific)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef.SetTagDescription(descStr)

	if zoneName != tzDef.GetZoneName() {
		t.Errorf("Error: Expected tzDef.GetZoneName()='%v'.\n" +
			"Instead, tzDef.GetZoneName()='%v'\n",
			zoneName, tzDef.GetZoneName())
	}

	if zoneOffset != tzDef.GetZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetZoneOffset()='%v'\n", zoneOffset, tzDef.GetZoneOffsetSeconds())
	}

	if zoneOffsetSecs != tzDef.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetZoneOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetZoneOffsetSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetZoneOffsetSeconds())
	}

	if zoneSign != tzDef.GetZoneSign() {
		t.Errorf("Error: Expected tzDef.GetZoneSign()='%v'.\n" +
			"Instead, tzDef.GetZoneSign()='%v'",
			zoneSign, tzDef.GetZoneSign())
	}

	if offsetHours != tzDef.GetOffsetHours() {
		t.Errorf("Error: Expected tzDef.GetOffsetHours()='%v'.\n" +
			"Instead, tzDef.GetOffsetHours()='%v'\n",
			offsetHours, tzDef.GetOffsetHours())
	}

	if offsetMinutes != tzDef.GetOffsetMinutes() {
		t.Errorf("Error: Expected tzDef.GetOffsetMinutes()='%v'.\n" +
			"Instead, tzDef.OffsetMinutes='%v'\n",
			offsetMinutes, tzDef.GetOffsetMinutes())
	}

	if offsetSeconds != tzDef.GetOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.OffsetSeconds='%v'\n",
			offsetSeconds, tzDef.GetOffsetSeconds())
	}

	if locationName != tzDef.GetLocationName() {
		t.Errorf("Error: Expected tzDef.GerLocationName()='%v'.\n" +
			"Instead, tzDef.GetLocationName()='%v'\n",
			locationName, tzDef.GetLocationName())
	}

	if descStr != tzDef.GetTagDescription() {
		t.Errorf("Error: Expected tzDef.GetTagDescription()='%v'.\n" +
			"Instead, tzDef.GetTagDescription()='%v'\n",
			descStr, tzDef.GetTagDescription())
	}

}

func TestTimeZoneDefDto_CopyOut_01(t *testing.T) {
	usPacificLoc, _ := time.LoadLocation(TZones.US.Pacific())

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	utcOffset := "-0800"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "US/Pacific"
	descStr := "US-Pacific"

	tzDef0, err := TimeZoneDefinition{}.New(tUsPacific)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefinition{}.New(tUsPacific)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef0.SetTagDescription(descStr)

	tzDef := tzDef0.CopyOut()

	if zoneName != tzDef.GetZoneName() {
		t.Errorf("Error: Expected tzDef.GetZoneName()='%v'.\n" +
			"Instead, tzDef.GetZoneName()='%v'\n",
			zoneName, tzDef.GetZoneName())
	}

	if zoneOffset != tzDef.GetZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetZoneOffset()='%v'\n",
			zoneOffset, tzDef.GetZoneOffset())
	}

	if utcOffset != tzDef.GetUtcOffset() {
		t.Errorf("Error: Expected tzDef.GetUtcOffset()='%v'.\n" +
			"Instead, tzDef.GetUtcOffset()='%v'\n",
			utcOffset, tzDef.GetUtcOffset())
	}

	if zoneOffsetSecs != tzDef.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetZoneOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetZoneOffsetSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetZoneOffsetSeconds())
	}

	if zoneSign != tzDef.GetZoneSign() {
		t.Errorf("Error: Expected tzDef.GetZoneSign()='%v'.\n" +
			"Instead, tzDef.GetZoneSign()='%v'", zoneSign, tzDef.GetZoneSign())
	}

	if offsetHours != tzDef.GetOffsetHours() {
		t.Errorf("Error: Expected tzDef.GetOffsetHours()='%v'.\n" +
			"Instead, tzDef.GetOffsetHours() ='%v'\n",
			offsetHours, tzDef.GetOffsetHours() )
	}

	if offsetMinutes != tzDef.GetOffsetMinutes() {
		t.Errorf("Error: Expected tzDef.GetOffsetMinutes()='%v'.\n" +
			"Instead, tzDef.GetOffsetMinutes()='%v'\n",
			offsetMinutes, tzDef.GetOffsetMinutes())
	}

	if offsetSeconds != tzDef.GetOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetOffsetSeconds()='%v'\n",
			offsetSeconds, tzDef.GetOffsetSeconds())
	}

	if locationName != tzDef.GetLocationName() {
		t.Errorf("Error: Expected tzDef.GetLocationName()='%v'.\n" +
			"Instead, tzDef.GetLocationName()='%v'\n",
			locationName, tzDef.GetLocationName())
	}

	if descStr != tzDef.GetTagDescription() {
		t.Errorf("Error: Expected tzDef.Description='%v'.\n" +
			"Instead, tzDef.Description='%v'\n",
			descStr, tzDef.GetTagDescription())
	}

}

func TestTimeZoneDefDto_CopyOut_02(t *testing.T) {
	americaLALoc, _ := time.LoadLocation(TZones.America.Los_Angeles())

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, americaLALoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	utcOffset := "-0800"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "America/Los_Angeles"
	descStr := "America-Los_Angeles"

	tzDef0, err := TimeZoneDefinition{}.New(tUsPacific)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefinition{}.New(tUsPacific)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef0.originalTimeZone.tagDescription = descStr

	tzDef := tzDef0.CopyOut()

	if zoneName != tzDef.GetZoneName() {
		t.Errorf("Error: Expected tzDef.GetZoneName()='%v'.\n" +
			"Instead, tzDef.GetZoneName()='%v'\n",
			zoneName, tzDef.GetZoneName())
	}

	if zoneOffset != tzDef.GetZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetZoneOffset()='%v'\n",
			zoneOffset, tzDef.GetZoneOffset())
	}

	if utcOffset != tzDef.GetUtcOffset() {
		t.Errorf("Error: Expected tzDef.GetUtcOffset()='%v'.\n" +
			"Instead, tzDef.GetUtcOffset()='%v'\n",
			utcOffset, tzDef.GetUtcOffset())
	}

	if zoneOffsetSecs != tzDef.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetZoneOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetZoneOffsetSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetZoneOffsetSeconds())
	}

	if zoneSign != tzDef.GetZoneSign() {
		t.Errorf("Error: Expected tzDef.GetZoneSign()='%v'.\n" +
			"Instead, tzDef.GetZoneSign()='%v'\n",
			zoneSign, tzDef.GetZoneSign())
	}

	if offsetHours != tzDef.GetOffsetHours() {
		t.Errorf("Error: Expected tzDef.GetOffsetHours()='%v'.\n" +
			"Instead, tzDef.GetOffsetHours()='%v'\n",
			offsetHours, tzDef.GetOffsetHours())
	}

	if offsetMinutes != tzDef.GetOffsetMinutes() {
		t.Errorf("Error: Expected tzDef.GetOffsetMinutes()='%v'.\n" +
			"Instead, tzDef.GetOffsetMinutes()='%v'\n",
			offsetMinutes, tzDef.GetOffsetMinutes())
	}

	if offsetSeconds != tzDef.GetOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetOffsetSeconds()='%v'\n",
			offsetSeconds, tzDef.GetOffsetSeconds())
	}

	if locationName != tzDef.GetLocationName() {
		t.Errorf("Error: Expected tzDef.GetLocationName()='%v'.\n" +
			"Instead, tzDef.GetLocationName()='%v'\n",
			locationName, tzDef.GetLocationName())
	}

	if descStr != tzDef.GetTagDescription() {
		t.Errorf("Error: Expected tzDef.GetTagDescription()='%v'.\n" +
			"Instead, tzDef.GetTagDescription()='%v'\n",
			descStr, tzDef.GetTagDescription())
	}

}

func TestTimeZoneDefDto_Equal_01(t *testing.T) {

	usPacificLoc, _ := time.LoadLocation(TZones.US.Pacific())

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	utcOffset := "-0800"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "US/Pacific"
	descStr := "US-Pacific"

	tzDef0, err := TimeZoneDefinition{}.New(tUsPacific)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefinition{}.New(tUsPacific).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef0.SetTagDescription(descStr)

	tzDef := tzDef0.CopyOut()

	if zoneName != tzDef.GetZoneName() {
		t.Errorf("Error: Expected tzDef.GetZoneName()='%v'.\n" +
			"Instead, tzDef.GetZoneName()='%v'\n",
			zoneName, tzDef.GetZoneName())
	}

	if zoneOffset != tzDef.GetZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetZoneOffset()='%v'\n",
			zoneOffset, tzDef.GetZoneOffset())
	}

	if utcOffset != tzDef.GetUtcOffset() {
		t.Errorf("Error: Expected tzDef.GetUtcOffset()='%v'.\n" +
			"Instead, tzDef.GetUtcOffset()='%v'\n",
			utcOffset, tzDef.GetUtcOffset())
	}

	if zoneOffsetSecs != tzDef.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetZoneOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetZoneOffsetSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetZoneOffsetSeconds())
	}

	if zoneSign != tzDef.GetZoneSign() {
		t.Errorf("Error: Expected tzDef.GetZoneSign()='%v'.\n" +
			"Instead, tzDef.GetZoneSign()='%v'\n",
			zoneSign, tzDef.GetZoneSign())
	}

	if offsetHours != tzDef.GetOffsetHours() {
		t.Errorf("Error: Expected tzDef.GetOffsetHours()='%v'.\n" +
			"Instead, tzDef.GetOffsetHours()='%v'\n",
			offsetHours, tzDef.GetOffsetHours())
	}

	if offsetMinutes != tzDef.GetOffsetMinutes() {
		t.Errorf("Error: Expected tzDef.GetOffsetMinutes()='%v'.\n" +
			"Instead, tzDef.GetOffsetMinutes()='%v'\n",
			offsetMinutes, tzDef.GetOffsetMinutes())
	}

	if offsetSeconds != tzDef.GetOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetOffsetSeconds()='%v'\n",
			offsetSeconds, tzDef.GetOffsetSeconds())
	}

	if locationName != tzDef.GetLocationName() {
		t.Errorf("Error: Expected tzDef.GetLocationName()='%v'.\n" +
			"Instead, tzDef.GetLocationName()='%v\n",
			locationName, tzDef.GetLocationName())
	}

	if descStr != tzDef.GetTagDescription() {
		t.Errorf("Error: Expected tzDef.GetTagDescription()='%v'.\n" +
			"Instead, tzDef.GetTagDescription()='%v'",
			descStr, tzDef.GetTagDescription())
	}

	if !tzDef0.Equal(tzDef) {
		t.Error("Error: Expected tzDef0 to be EQUAL to tzDef. IT WAS NOT!")
	}

}

func TestTimeZoneDefDto_Equal_02(t *testing.T) {

	usPacificLoc, _ := time.LoadLocation(TZones.US.Pacific())

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	utcOffset := "-0800"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "US/Pacific"
	descStr := "US-Pacific"

	tzDef0, err := TimeZoneDefinition{}.New(tUsPacific)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefinition{}.New(tUsPacific).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef0.SetTagDescription(descStr)

	tzDef := tzDef0.CopyOut()

	if zoneName != tzDef.GetZoneName() {
		t.Errorf("Error: Expected tzDef.GetZoneName()='%v'.\n" +
			"Instead, tzDef.GetZoneName()='%v'\n",
			zoneName, tzDef.GetZoneName())
	}

	if zoneOffset != tzDef.GetZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetZoneOffset()='%v'\n",
			zoneOffset, tzDef.GetZoneOffset())
	}

	if utcOffset != tzDef.GetUtcOffset() {
		t.Errorf("Error: Expected tzDef.GetUtcOffset()='%v'.\n" +
			"Instead, tzDef.GetUtcOffset()='%v'\n",
			zoneOffset, tzDef.GetUtcOffset())
	}

	if zoneOffsetSecs != tzDef.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetZoneOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetZoneOffsetSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetZoneOffsetSeconds())
	}

	if zoneSign != tzDef.GetZoneSign() {
		t.Errorf("Error: Expected tzDef.GetZoneSign()='%v'.\n" +
			"Instead, tzDef.GetZoneSign()='%v'\n",
			zoneSign, tzDef.GetZoneSign())
	}

	if offsetHours != tzDef.GetOffsetHours() {
		t.Errorf("Error: Expected tzDef.GetOffsetHours()='%v'.\n" +
			"Instead, tzDef.GetOffsetHours()='%v'\n",
			offsetHours, tzDef.GetOffsetHours())
	}

	if offsetMinutes != tzDef.GetOffsetMinutes() {
		t.Errorf("Error: Expected tzDef.GetOffsetMinutes()='%v'.\n" +
			"Instead, tzDef.GetOffsetMinutes()='%v'\n",
			offsetMinutes, tzDef.GetOffsetMinutes())
	}

	if offsetSeconds != tzDef.GetOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetOffsetSeconds()='%v'\n",
			offsetSeconds, tzDef.GetOffsetSeconds())
	}

	if locationName != tzDef.GetLocationName() {
		t.Errorf("Error: Expected tzDef.GetLocationName()='%v'.\n" +
			"Instead, tzDef.GetLocationName()='%v'\n",
			locationName, tzDef.GetLocationName())
	}

	if descStr != tzDef.GetTagDescription() {
		t.Errorf("Error: Expected tzDef.GetTagDescription()='%v'.\n" +
			"Instead, tzDef.GetTagDescription()='%v'\n",
			descStr, tzDef.GetTagDescription())
	}

	tzDef0.originalTimeZone.locationPtr = nil

	if tzDef0.Equal(tzDef) {
		t.Error("Error: Expected tzDef0 to be NOT EQUAL to tzDef. IT WAS EQUAL!")
	}

}

func TestTimeZoneDefDto_Equal_03(t *testing.T) {

	americaLALoc, _ := time.LoadLocation(TZones.America.Los_Angeles())

	tAmericaLA := time.Date(2014, 2, 15, 19, 54, 30, 38175584, americaLALoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	utcOffset := "-0800"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "America/Los_Angeles"
	descStr := "America-Los_Angeles"

	tzDef0, err := TimeZoneDefinition{}.New(tAmericaLA)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefinition{}.New(tAmericaLA).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef0.SetTagDescription(descStr)

	tzDef := tzDef0.CopyOut()

	if zoneName != tzDef.GetZoneName() {
		t.Errorf("Error: Expected tzDef.GetZoneName()='%v'.\n" +
			"Instead, tzDef.GetZoneName()='%v'\n", zoneName, tzDef.GetZoneName())
	}

	if zoneOffset != tzDef.GetZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetZoneOffset()='%v'\n",
			zoneOffset, tzDef.GetZoneOffset())
	}

	if utcOffset != tzDef.GetUtcOffset() {
		t.Errorf("Error: Expected tzDef.GetUtcOffset()='%v'.\n" +
			"Instead, tzDef.GetUtcOffset()='%v'\n",
			utcOffset, tzDef.GetUtcOffset())
	}

	if zoneOffsetSecs != tzDef.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetZoneOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetZoneOffsetSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetZoneOffsetSeconds())
	}

	if zoneSign != tzDef.GetZoneSign() {
		t.Errorf("Error: Expected tzDef.GetZoneSign()='%v'.\n" +
			"Instead, tzDef.GetZoneSign()='%v'\n", zoneSign, tzDef.GetZoneSign())
	}

	if offsetHours != tzDef.GetOffsetHours() {
		t.Errorf("Error: Expected tzDef.GetOffsetHours()='%v'.\n" +
			"Instead, tzDef.GetOffsetHours()='%v'\n",
			offsetHours, tzDef.GetOffsetHours())
	}

	if offsetMinutes != tzDef.GetOffsetMinutes() {
		t.Errorf("Error: Expected tzDef.GetOffsetMinutes()='%v'.\n" +
			"Instead, tzDef.GetOffsetMinutes()='%v'\n",
			offsetMinutes, tzDef.GetOffsetMinutes())
	}

	if offsetSeconds != tzDef.GetOffsetSeconds() {
		t.Errorf("Error: Expected tzDef.GetOffsetSeconds()='%v'.\n" +
			"Instead, tzDef.GetOffsetSeconds()='%v'\n",
			offsetSeconds, tzDef.GetOffsetSeconds())
	}

	if locationName != tzDef.GetLocationName() {
		t.Errorf("Error: Expected tzDef.GetLocationName()='%v'.\n" +
			"Instead, tzDef.GetLocationName()='%v'\n",
			locationName, tzDef.GetLocationName())
	}

	if descStr != tzDef.GetTagDescription() {
		t.Errorf("Error: Expected tzDef.GetTagDescription()='%v'.\n" +
			"Instead, tzDef.GetTagDescription()='%v'\n",
			descStr, tzDef.GetTagDescription())
	}

	tzDef0.originalTimeZone.locationPtr = nil

	if tzDef0.Equal(tzDef) {
		t.Error("Error: Expected tzDef0 to be NOT EQUAL to tzDef. IT WAS EQUAL!")
	}

}
