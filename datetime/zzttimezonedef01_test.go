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
	zoneSignChar := "-"
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

	if zoneName != tzDef.GetOriginalZoneName() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneName()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneName()='%v'\n",
			zoneName, tzDef.GetOriginalZoneName())
	}

	if zoneOffset != tzDef.GetOriginalZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffset()='%v'\n", zoneOffset, tzDef.GetOriginalZoneOffsetTotalSeconds())
	}

	if zoneOffsetSecs != tzDef.GetOriginalZoneOffsetTotalSeconds() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetOriginalZoneOffsetTotalSeconds())
	}

	actOffsetSignChar,
	_,
	actOffsetHours,
	actOffsetMinutes,
	actOffsetSeconds := tzDef.GetOriginalOffsetElements()

	if zoneSignChar != actOffsetSignChar {
		t.Errorf("Error: Expected 'tzDef' Offset Sign Char='%v'.\n" +
			"Instead, 'tzDef' Offset Sign Char='%v'",
			zoneSignChar, actOffsetSignChar)
	}

	if offsetHours != actOffsetHours {
		t.Errorf("Error: Expected 'tzDef' Offset Hours='%v'.\n" +
			"Instead, 'tzDef' Offset Hours='%v'\n",
			offsetHours, actOffsetHours)
	}

	if offsetMinutes != actOffsetMinutes {
		t.Errorf("Error: Expected 'tzDef' Offset Minutes='%v'.\n" +
			"Instead, 'tzDef' Offset Minutes='%v'\n",
			offsetMinutes, actOffsetMinutes)
	}

	if offsetSeconds != actOffsetSeconds {
		t.Errorf("Error: Expected 'tzDef' Offset Seconds='%v'.\n" +
			"Instead, 'tzDef' Offset Seconds='%v'\n",
			offsetSeconds, actOffsetSeconds)
	}

	if locationName != tzDef.GetOriginalLocationName() {
		t.Errorf("Error: Expected tzDef.GerLocationName()='%v'.\n" +
			"Instead, tzDef.GetOriginalLocationName()='%v'\n",
			locationName, tzDef.GetOriginalLocationName())
	}

	if descStr != tzDef.GetOriginalTagDescription() {
		t.Errorf("Error: Expected tzDef.GetOriginalTagDescription()='%v'.\n" +
			"Instead, tzDef.GetOriginalTagDescription()='%v'\n",
			descStr, tzDef.GetOriginalTagDescription())
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
	zoneSignChar := "-"
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

	if zoneName != tzDef.GetOriginalZoneName() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneName()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneName()='%v'\n",
			zoneName, tzDef.GetOriginalZoneName())
	}

	if zoneOffset != tzDef.GetOriginalZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffset()='%v'\n",
			zoneOffset, tzDef.GetOriginalZoneOffset())
	}

	if utcOffset != tzDef.GetOriginalUtcOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalUtcOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalUtcOffset()='%v'\n",
			utcOffset, tzDef.GetOriginalUtcOffset())
	}

	if zoneOffsetSecs != tzDef.GetOriginalZoneOffsetTotalSeconds() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetOriginalZoneOffsetTotalSeconds())
	}

	if zoneSign != tzDef.GetOriginalOffsetSignValue() {
		t.Errorf("Error: Expected tzDef.GetOriginalOffsetSignValue()='%v'.\n" +
			"Instead, tzDef.GetOriginalOffsetSignValue()='%v'", zoneSign, tzDef.GetOriginalOffsetSignValue())
	}

	actOffsetSignChar,
	_,
	actOffsetHours,
	actOffsetMinutes,
	actOffsetSeconds := tzDef.GetOriginalOffsetElements()

	if zoneSignChar != actOffsetSignChar {
		t.Errorf("Error: Expected 'tzDef' Offset Sign Char='%v'\n" +
			"Instead, 'tzDef' Offset Sign Char='%v'\n",
			zoneSignChar, actOffsetSignChar)
	}

	if offsetHours != actOffsetHours {
		t.Errorf("Error: Expected 'tzDef' Offset Hours='%v'.\n" +
			"Instead, 'tzDef' Offset Hours='%v'\n",
			offsetHours, actOffsetHours )
	}

	if offsetMinutes != actOffsetMinutes {
		t.Errorf("Error: Expected 'tzDef' Offset Minutes='%v'.\n" +
			"Instead, 'tzDef' Offset Minutes='%v'\n",
			offsetMinutes, actOffsetMinutes)
	}

	if offsetSeconds != actOffsetSeconds {
		t.Errorf("Error: Expected 'tzDef' OffsetSeconds='%v'.\n" +
			"Instead, 'tzDef' Offset Seconds ='%v'\n",
			offsetSeconds, actOffsetSeconds)
	}

	if locationName != tzDef.GetOriginalLocationName() {
		t.Errorf("Error: Expected tzDef.GetOriginalLocationName()='%v'.\n" +
			"Instead, tzDef.GetOriginalLocationName()='%v'\n",
			locationName, tzDef.GetOriginalLocationName())
	}

	if descStr != tzDef.GetOriginalTagDescription() {
		t.Errorf("Error: Expected tzDef.Description='%v'.\n" +
			"Instead, tzDef.Description='%v'\n",
			descStr, tzDef.GetOriginalTagDescription())
	}

}

func TestTimeZoneDefDto_CopyOut_02(t *testing.T) {
	americaLALoc, _ := time.LoadLocation(TZones.America.Los_Angeles())

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, americaLALoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	utcOffset := "-0800"
	zoneOffsetSecs := -28800
	zoneSignChar := "-"
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

	if zoneName != tzDef.GetOriginalZoneName() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneName()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneName()='%v'\n",
			zoneName, tzDef.GetOriginalZoneName())
	}

	if zoneOffset != tzDef.GetOriginalZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffset()='%v'\n",
			zoneOffset, tzDef.GetOriginalZoneOffset())
	}

	if utcOffset != tzDef.GetOriginalUtcOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalUtcOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalUtcOffset()='%v'\n",
			utcOffset, tzDef.GetOriginalUtcOffset())
	}

	if zoneOffsetSecs != tzDef.GetOriginalZoneOffsetTotalSeconds() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetOriginalZoneOffsetTotalSeconds())
	}

	actOffsetSignChar,
	_,
	actOffsetHours,
	actOffsetMinutes,
	actOffsetSeconds := tzDef.GetOriginalOffsetElements()

	if zoneSignChar != actOffsetSignChar {
		t.Errorf("Error: Expected 'tzDef' Zone Sign Char='%v'.\n" +
			"Instead, 'tzDef' Zone Sign Char='%v'\n",
			zoneSignChar, actOffsetSignChar)
	}

	if offsetHours != actOffsetHours {
		t.Errorf("Error: Expected 'tzDef' Offset Hours='%v'.\n" +
			"Instead, 'tzDef' Offset Hours='%v'\n",
			offsetHours, actOffsetHours)
	}

	if offsetMinutes != actOffsetMinutes {
		t.Errorf("Error: Expected 'tzDef' Offset Minutes='%v'.\n" +
			"Instead, 'tzDef' Offset Minutes='%v'\n",
			offsetMinutes, actOffsetMinutes)
	}

	if offsetSeconds != actOffsetSeconds {
		t.Errorf("Error: Expected 'tzDef' Offset Seconds='%v'.\n" +
			"Instead, 'tzDef' Offset Seconds='%v'\n",
			offsetSeconds, actOffsetSeconds)
	}

	if locationName != tzDef.GetOriginalLocationName() {
		t.Errorf("Error: Expected tzDef.GetOriginalLocationName()='%v'.\n" +
			"Instead, tzDef.GetOriginalLocationName()='%v'\n",
			locationName, tzDef.GetOriginalLocationName())
	}

	if descStr != tzDef.GetOriginalTagDescription() {
		t.Errorf("Error: Expected tzDef.GetOriginalTagDescription()='%v'.\n" +
			"Instead, tzDef.GetOriginalTagDescription()='%v'\n",
			descStr, tzDef.GetOriginalTagDescription())
	}

}

func TestTimeZoneDefDto_Equal_01(t *testing.T) {

	usPacificLoc, _ := time.LoadLocation(TZones.US.Pacific())

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	utcOffset := "-0800"
	zoneOffsetSecs := -28800
	zoneSignChar := "-"
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

	if zoneName != tzDef.GetOriginalZoneName() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneName()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneName()='%v'\n",
			zoneName, tzDef.GetOriginalZoneName())
	}

	if zoneOffset != tzDef.GetOriginalZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffset()='%v'\n",
			zoneOffset, tzDef.GetOriginalZoneOffset())
	}

	if utcOffset != tzDef.GetOriginalUtcOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalUtcOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalUtcOffset()='%v'\n",
			utcOffset, tzDef.GetOriginalUtcOffset())
	}

	if zoneOffsetSecs != tzDef.GetOriginalZoneOffsetTotalSeconds() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetOriginalZoneOffsetTotalSeconds())
	}

	actOffsetSignChar,
	_,
	actOffsetHours,
	actOffsetMinutes,
	actOffsetSeconds := tzDef.GetOriginalOffsetElements()

	if zoneSignChar != actOffsetSignChar {
		t.Errorf("Error: Expected 'tzDef' Sign Char='%v'.\n" +
			"Instead, 'tzDef' Sign Char='%v'\n",
			zoneSignChar, actOffsetSignChar)
	}

	if offsetHours != actOffsetHours {
		t.Errorf("Error: Expected 'tzDef' Offset Hours='%v'.\n" +
			"Instead, 'tzDef' Offset Hours='%v'\n",
			offsetHours, actOffsetHours)
	}

	if offsetMinutes != actOffsetMinutes {
		t.Errorf("Error: Expected 'tzDef' Offset Minutes='%v'.\n" +
			"Instead, 'tzDef' Offset Minutes='%v'\n",
			offsetMinutes, actOffsetMinutes)
	}

	if offsetSeconds != actOffsetSeconds {
		t.Errorf("Error: Expected 'tzDef' Offset Seconds='%v'.\n" +
			"Instead, 'tzDef' Offset Seconds='%v'\n",
			offsetSeconds, actOffsetSeconds)
	}

	if locationName != tzDef.GetOriginalLocationName() {
		t.Errorf("Error: Expected tzDef.GetOriginalLocationName()='%v'.\n" +
			"Instead, tzDef.GetOriginalLocationName()='%v\n",
			locationName, tzDef.GetOriginalLocationName())
	}

	if descStr != tzDef.GetOriginalTagDescription() {
		t.Errorf("Error: Expected tzDef.GetOriginalTagDescription()='%v'.\n" +
			"Instead, tzDef.GetOriginalTagDescription()='%v'",
			descStr, tzDef.GetOriginalTagDescription())
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
	zoneSignChar := "-"
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

	if zoneName != tzDef.GetOriginalZoneName() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneName()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneName()='%v'\n",
			zoneName, tzDef.GetOriginalZoneName())
	}

	if zoneOffset != tzDef.GetOriginalZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffset()='%v'\n",
			zoneOffset, tzDef.GetOriginalZoneOffset())
	}

	if utcOffset != tzDef.GetOriginalUtcOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalUtcOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalUtcOffset()='%v'\n",
			zoneOffset, tzDef.GetOriginalUtcOffset())
	}

	if zoneOffsetSecs != tzDef.GetOriginalZoneOffsetTotalSeconds() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetOriginalZoneOffsetTotalSeconds())
	}

	actOffsetSignChar,
	_,
	actOffsetHours,
	actOffsetMinutes,
	actOffsetSeconds := tzDef.GetOriginalOffsetElements()

	if zoneSignChar != actOffsetSignChar {
		t.Errorf("Error: Expected 'tzDef' Offset Sign Char='%v'.\n" +
			"Instead, 'tzDef' Offset Sign Char='%v'\n",
			zoneSignChar, actOffsetSignChar)
	}

	if offsetHours != actOffsetHours {
		t.Errorf("Error: Expected 'tzDef' Offset Hours='%v'.\n" +
			"Instead, 'tzDef' Offset Hours='%v'\n",
			offsetHours, actOffsetHours)
	}

	if offsetMinutes != actOffsetMinutes {
		t.Errorf("Error: Expected 'tzDef' Offset Minutes='%v'.\n" +
			"Instead, 'tzDef' Offset Minutes='%v'\n",
			offsetMinutes, actOffsetMinutes)
	}

	if offsetSeconds != actOffsetSeconds {
		t.Errorf("Error: Expected 'tzDef' Offset Seconds='%v'.\n" +
			"Instead, 'tzDef' Offset Seconds='%v'\n",
			offsetSeconds, actOffsetSeconds)
	}

	if locationName != tzDef.GetOriginalLocationName() {
		t.Errorf("Error: Expected tzDef.GetOriginalLocationName()='%v'.\n" +
			"Instead, tzDef.GetOriginalLocationName()='%v'\n",
			locationName, tzDef.GetOriginalLocationName())
	}

	if descStr != tzDef.GetOriginalTagDescription() {
		t.Errorf("Error: Expected tzDef.GetOriginalTagDescription()='%v'.\n" +
			"Instead, tzDef.GetOriginalTagDescription()='%v'\n",
			descStr, tzDef.GetOriginalTagDescription())
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
	zoneSignChar := "-"
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

	if zoneName != tzDef.GetOriginalZoneName() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneName()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneName()='%v'\n", zoneName, tzDef.GetOriginalZoneName())
	}

	if zoneOffset != tzDef.GetOriginalZoneOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffset()='%v'\n",
			zoneOffset, tzDef.GetOriginalZoneOffset())
	}

	if utcOffset != tzDef.GetOriginalUtcOffset() {
		t.Errorf("Error: Expected tzDef.GetOriginalUtcOffset()='%v'.\n" +
			"Instead, tzDef.GetOriginalUtcOffset()='%v'\n",
			utcOffset, tzDef.GetOriginalUtcOffset())
	}

	if zoneOffsetSecs != tzDef.GetOriginalZoneOffsetTotalSeconds() {
		t.Errorf("Error: Expected tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'.\n" +
			"Instead, tzDef.GetOriginalZoneOffsetTotalSeconds()='%v'\n",
			zoneOffsetSecs, tzDef.GetOriginalZoneOffsetTotalSeconds())
	}

	actOffsetSignChar,
	_,
	actOffsetHours,
	actOffsetMinutes,
	actOffsetSeconds := tzDef.GetOriginalOffsetElements()

	if zoneSignChar != actOffsetSignChar {
		t.Errorf("Error: Expected 'tzDef' Offset Sign Char='%v'.\n" +
			"Instead, 'tzDef' Offset Sign Char='%v'\n",
			zoneSignChar, actOffsetSignChar)
	}

	if offsetHours != actOffsetHours {
		t.Errorf("Error: Expected 'tzDef' Offset Hours='%v'.\n" +
			"Instead, 'tzDef' Offset Hours='%v'\n",
			offsetHours, actOffsetHours)
	}

	if offsetMinutes != actOffsetMinutes {
		t.Errorf("Error: Expected 'tzDef' Offset Minutes='%v'.\n" +
			"Instead, 'tzDef' Offset Minutes='%v'\n",
			offsetMinutes, actOffsetMinutes)
	}

	if offsetSeconds != actOffsetSeconds {
		t.Errorf("Error: Expected 'tzDef' Offset Seconds='%v'.\n" +
			"Instead, 'tzDef' Offset Seconds='%v'\n",
			offsetSeconds, actOffsetSeconds)
	}

	if locationName != tzDef.GetOriginalLocationName() {
		t.Errorf("Error: Expected tzDef.GetOriginalLocationName()='%v'.\n" +
			"Instead, tzDef.GetOriginalLocationName()='%v'\n",
			locationName, tzDef.GetOriginalLocationName())
	}

	if descStr != tzDef.GetOriginalTagDescription() {
		t.Errorf("Error: Expected tzDef.GetOriginalTagDescription()='%v'.\n" +
			"Instead, tzDef.GetOriginalTagDescription()='%v'\n",
			descStr, tzDef.GetOriginalTagDescription())
	}

	tzDef0.originalTimeZone.locationPtr = nil

	if tzDef0.Equal(tzDef) {
		t.Error("Error: Expected tzDef0 to be NOT EQUAL to tzDef. IT WAS EQUAL!")
	}

}
