package zztests

import (
	"github.com/MikeAustin71/datetimeopsgo/datetime"
	"testing"
	"time"
)

func TestTimeZoneDefDto_New_01(t *testing.T) {

	usPacificLoc, _ := time.LoadLocation(datetime.TZones.US.Pacific())

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

	tzDef, err := datetime.TimeZoneDefDto{}.New(tUsPacific)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(tUsPacific)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef.Description = descStr

	if zoneName != tzDef.ZoneName {
		t.Errorf("Error: Expected tzDef.ZoneName='%v'. Instead, tzDef.ZoneName='%v'", zoneName, tzDef.ZoneName)
	}

	if zoneOffset != tzDef.ZoneOffset {
		t.Errorf("Error: Expected tzDef.ZoneOffset='%v'. Instead, tzDef.ZoneOffset='%v'", zoneOffset, tzDef.ZoneOffset)
	}

	if zoneOffsetSecs != tzDef.ZoneOffsetSeconds {
		t.Errorf("Error: Expected tzDef.ZoneOffsetSeconds='%v'. Instead, tzDef.ZoneOffsetSeconds='%v'", zoneOffsetSecs, tzDef.ZoneOffsetSeconds)
	}

	if zoneSign != tzDef.ZoneSign {
		t.Errorf("Error: Expected tzDef.ZoneSign='%v'. Instead, tzDef.ZoneSign='%v'", zoneSign, tzDef.ZoneSign)
	}

	if offsetHours != tzDef.OffsetHours {
		t.Errorf("Error: Expected tzDef.OffsetHours='%v'. Instead, tzDef.OffsetHours='%v'", offsetHours, tzDef.OffsetHours)
	}

	if offsetMinutes != tzDef.OffsetMinutes {
		t.Errorf("Error: Expected tzDef.OffsetMinutes='%v'. Instead, tzDef.OffsetMinutes='%v'", offsetMinutes, tzDef.OffsetMinutes)
	}

	if offsetSeconds != tzDef.OffsetSeconds {
		t.Errorf("Error: Expected tzDef.OffsetSeconds='%v'. Instead, tzDef.OffsetSeconds='%v'", offsetSeconds, tzDef.OffsetSeconds)
	}

	if locationName != tzDef.LocationName {
		t.Errorf("Error: Expected tzDef.LocationName='%v'. Instead, tzDef.LocationName='%v'", locationName, tzDef.LocationName)
	}

	if descStr != tzDef.Description {
		t.Errorf("Error: Expected tzDef.Description='%v'. Instead, tzDef.Description='%v'", descStr, tzDef.Description)
	}

}

func TestTimeZoneDefDto_CopyOut_01(t *testing.T) {
	usPacificLoc, _ := time.LoadLocation(datetime.TZones.US.Pacific())

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "US/Pacific"
	descStr := "US-Pacific"

	tzDef0, err := datetime.TimeZoneDefDto{}.New(tUsPacific)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(tUsPacific)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef0.Description = descStr

	tzDef := tzDef0.CopyOut()

	if zoneName != tzDef.ZoneName {
		t.Errorf("Error: Expected tzDef.ZoneName='%v'. Instead, tzDef.ZoneName='%v'", zoneName, tzDef.ZoneName)
	}

	if zoneOffset != tzDef.ZoneOffset {
		t.Errorf("Error: Expected tzDef.ZoneOffset='%v'. Instead, tzDef.ZoneOffset='%v'", zoneOffset, tzDef.ZoneOffset)
	}

	if zoneOffsetSecs != tzDef.ZoneOffsetSeconds {
		t.Errorf("Error: Expected tzDef.ZoneOffsetSeconds='%v'. Instead, tzDef.ZoneOffsetSeconds='%v'", zoneOffsetSecs, tzDef.ZoneOffsetSeconds)
	}

	if zoneSign != tzDef.ZoneSign {
		t.Errorf("Error: Expected tzDef.ZoneSign='%v'. Instead, tzDef.ZoneSign='%v'", zoneSign, tzDef.ZoneSign)
	}

	if offsetHours != tzDef.OffsetHours {
		t.Errorf("Error: Expected tzDef.OffsetHours='%v'. Instead, tzDef.OffsetHours='%v'", offsetHours, tzDef.OffsetHours)
	}

	if offsetMinutes != tzDef.OffsetMinutes {
		t.Errorf("Error: Expected tzDef.OffsetMinutes='%v'. Instead, tzDef.OffsetMinutes='%v'", offsetMinutes, tzDef.OffsetMinutes)
	}

	if offsetSeconds != tzDef.OffsetSeconds {
		t.Errorf("Error: Expected tzDef.OffsetSeconds='%v'. Instead, tzDef.OffsetSeconds='%v'", offsetSeconds, tzDef.OffsetSeconds)
	}

	if locationName != tzDef.LocationName {
		t.Errorf("Error: Expected tzDef.LocationName='%v'. Instead, tzDef.LocationName='%v'", locationName, tzDef.LocationName)
	}

	if descStr != tzDef.Description {
		t.Errorf("Error: Expected tzDef.Description='%v'. Instead, tzDef.Description='%v'", descStr, tzDef.Description)
	}

}

func TestTimeZoneDefDto_CopyOut_02(t *testing.T) {
	americaLALoc, _ := time.LoadLocation(datetime.TZones.America.Los_Angeles())

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, americaLALoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "America/Los_Angeles"
	descStr := "America-Los_Angeles"

	tzDef0, err := datetime.TimeZoneDefDto{}.New(tUsPacific)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(tUsPacific)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef0.Description = descStr

	tzDef := tzDef0.CopyOut()

	if zoneName != tzDef.ZoneName {
		t.Errorf("Error: Expected tzDef.ZoneName='%v'. Instead, tzDef.ZoneName='%v'", zoneName, tzDef.ZoneName)
	}

	if zoneOffset != tzDef.ZoneOffset {
		t.Errorf("Error: Expected tzDef.ZoneOffset='%v'. Instead, tzDef.ZoneOffset='%v'", zoneOffset, tzDef.ZoneOffset)
	}

	if zoneOffsetSecs != tzDef.ZoneOffsetSeconds {
		t.Errorf("Error: Expected tzDef.ZoneOffsetSeconds='%v'. Instead, tzDef.ZoneOffsetSeconds='%v'", zoneOffsetSecs, tzDef.ZoneOffsetSeconds)
	}

	if zoneSign != tzDef.ZoneSign {
		t.Errorf("Error: Expected tzDef.ZoneSign='%v'. Instead, tzDef.ZoneSign='%v'", zoneSign, tzDef.ZoneSign)
	}

	if offsetHours != tzDef.OffsetHours {
		t.Errorf("Error: Expected tzDef.OffsetHours='%v'. Instead, tzDef.OffsetHours='%v'", offsetHours, tzDef.OffsetHours)
	}

	if offsetMinutes != tzDef.OffsetMinutes {
		t.Errorf("Error: Expected tzDef.OffsetMinutes='%v'. Instead, tzDef.OffsetMinutes='%v'", offsetMinutes, tzDef.OffsetMinutes)
	}

	if offsetSeconds != tzDef.OffsetSeconds {
		t.Errorf("Error: Expected tzDef.OffsetSeconds='%v'. Instead, tzDef.OffsetSeconds='%v'", offsetSeconds, tzDef.OffsetSeconds)
	}

	if locationName != tzDef.LocationName {
		t.Errorf("Error: Expected tzDef.LocationName='%v'. Instead, tzDef.LocationName='%v'", locationName, tzDef.LocationName)
	}

	if descStr != tzDef.Description {
		t.Errorf("Error: Expected tzDef.Description='%v'. Instead, tzDef.Description='%v'", descStr, tzDef.Description)
	}

}

func TestTimeZoneDefDto_Equal_01(t *testing.T) {

	usPacificLoc, _ := time.LoadLocation(datetime.TZones.US.Pacific())

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "US/Pacific"
	descStr := "US-Pacific"

	tzDef0, err := datetime.TimeZoneDefDto{}.New(tUsPacific)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(tUsPacific).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef0.Description = descStr

	tzDef := tzDef0.CopyOut()

	if zoneName != tzDef.ZoneName {
		t.Errorf("Error: Expected tzDef.ZoneName='%v'. Instead, tzDef.ZoneName='%v'", zoneName, tzDef.ZoneName)
	}

	if zoneOffset != tzDef.ZoneOffset {
		t.Errorf("Error: Expected tzDef.ZoneOffset='%v'. Instead, tzDef.ZoneOffset='%v'", zoneOffset, tzDef.ZoneOffset)
	}

	if zoneOffsetSecs != tzDef.ZoneOffsetSeconds {
		t.Errorf("Error: Expected tzDef.ZoneOffsetSeconds='%v'. Instead, tzDef.ZoneOffsetSeconds='%v'", zoneOffsetSecs, tzDef.ZoneOffsetSeconds)
	}

	if zoneSign != tzDef.ZoneSign {
		t.Errorf("Error: Expected tzDef.ZoneSign='%v'. Instead, tzDef.ZoneSign='%v'", zoneSign, tzDef.ZoneSign)
	}

	if offsetHours != tzDef.OffsetHours {
		t.Errorf("Error: Expected tzDef.OffsetHours='%v'. Instead, tzDef.OffsetHours='%v'", offsetHours, tzDef.OffsetHours)
	}

	if offsetMinutes != tzDef.OffsetMinutes {
		t.Errorf("Error: Expected tzDef.OffsetMinutes='%v'. Instead, tzDef.OffsetMinutes='%v'", offsetMinutes, tzDef.OffsetMinutes)
	}

	if offsetSeconds != tzDef.OffsetSeconds {
		t.Errorf("Error: Expected tzDef.OffsetSeconds='%v'. Instead, tzDef.OffsetSeconds='%v'", offsetSeconds, tzDef.OffsetSeconds)
	}

	if locationName != tzDef.LocationName {
		t.Errorf("Error: Expected tzDef.LocationName='%v'. Instead, tzDef.LocationName='%v'", locationName, tzDef.LocationName)
	}

	if descStr != tzDef.Description {
		t.Errorf("Error: Expected tzDef.Description='%v'. Instead, tzDef.Description='%v'", descStr, tzDef.Description)
	}

	if !tzDef0.Equal(tzDef) {
		t.Error("Error: Expected tzDef0 to be EQUAL to tzDef. IT WAS NOT!")
	}

}

func TestTimeZoneDefDto_Equal_02(t *testing.T) {

	usPacificLoc, _ := time.LoadLocation(datetime.TZones.US.Pacific())

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "US/Pacific"
	descStr := "US-Pacific"

	tzDef0, err := datetime.TimeZoneDefDto{}.New(tUsPacific)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(tUsPacific).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef0.Description = descStr

	tzDef := tzDef0.CopyOut()

	if zoneName != tzDef.ZoneName {
		t.Errorf("Error: Expected tzDef.ZoneName='%v'. Instead, tzDef.ZoneName='%v'", zoneName, tzDef.ZoneName)
	}

	if zoneOffset != tzDef.ZoneOffset {
		t.Errorf("Error: Expected tzDef.ZoneOffset='%v'. Instead, tzDef.ZoneOffset='%v'", zoneOffset, tzDef.ZoneOffset)
	}

	if zoneOffsetSecs != tzDef.ZoneOffsetSeconds {
		t.Errorf("Error: Expected tzDef.ZoneOffsetSeconds='%v'. Instead, tzDef.ZoneOffsetSeconds='%v'", zoneOffsetSecs, tzDef.ZoneOffsetSeconds)
	}

	if zoneSign != tzDef.ZoneSign {
		t.Errorf("Error: Expected tzDef.ZoneSign='%v'. Instead, tzDef.ZoneSign='%v'", zoneSign, tzDef.ZoneSign)
	}

	if offsetHours != tzDef.OffsetHours {
		t.Errorf("Error: Expected tzDef.OffsetHours='%v'. Instead, tzDef.OffsetHours='%v'", offsetHours, tzDef.OffsetHours)
	}

	if offsetMinutes != tzDef.OffsetMinutes {
		t.Errorf("Error: Expected tzDef.OffsetMinutes='%v'. Instead, tzDef.OffsetMinutes='%v'", offsetMinutes, tzDef.OffsetMinutes)
	}

	if offsetSeconds != tzDef.OffsetSeconds {
		t.Errorf("Error: Expected tzDef.OffsetSeconds='%v'. Instead, tzDef.OffsetSeconds='%v'", offsetSeconds, tzDef.OffsetSeconds)
	}

	if locationName != tzDef.LocationName {
		t.Errorf("Error: Expected tzDef.LocationName='%v'. Instead, tzDef.LocationName='%v'", locationName, tzDef.LocationName)
	}

	if descStr != tzDef.Description {
		t.Errorf("Error: Expected tzDef.Description='%v'. Instead, tzDef.Description='%v'", descStr, tzDef.Description)
	}

	tzDef0.Location = nil

	if tzDef0.Equal(tzDef) {
		t.Error("Error: Expected tzDef0 to be NOT EQUAL to tzDef. IT WAS EQUAL!")
	}

}

func TestTimeZoneDefDto_Equal_03(t *testing.T) {

	americaLALoc, _ := time.LoadLocation(datetime.TZones.America.Los_Angeles())

	tamericaLA := time.Date(2014, 2, 15, 19, 54, 30, 38175584, americaLALoc)
	zoneName := "PST"
	zoneOffset := "-0800 PST"
	zoneOffsetSecs := -28800
	zoneSign := -1
	offsetHours := 8
	offsetMinutes := 0
	offsetSeconds := 0
	locationName := "America/Los_Angeles"
	descStr := "America-Los_Angeles"

	tzDef0, err := datetime.TimeZoneDefDto{}.New(tamericaLA)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(tamericaLA).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzDef0.Description = descStr

	tzDef := tzDef0.CopyOut()

	if zoneName != tzDef.ZoneName {
		t.Errorf("Error: Expected tzDef.ZoneName='%v'.\n" +
			"Instead, tzDef.ZoneName='%v'\n", zoneName, tzDef.ZoneName)
	}

	if zoneOffset != tzDef.ZoneOffset {
		t.Errorf("Error: Expected tzDef.ZoneOffset='%v'.\n" +
			"Instead, tzDef.ZoneOffset='%v'\n", zoneOffset, tzDef.ZoneOffset)
	}

	if zoneOffsetSecs != tzDef.ZoneOffsetSeconds {
		t.Errorf("Error: Expected tzDef.ZoneOffsetSeconds='%v'.\n" +
			"Instead, tzDef.ZoneOffsetSeconds='%v'\n", zoneOffsetSecs, tzDef.ZoneOffsetSeconds)
	}

	if zoneSign != tzDef.ZoneSign {
		t.Errorf("Error: Expected tzDef.ZoneSign='%v'.\n" +
			"Instead, tzDef.ZoneSign='%v'\n", zoneSign, tzDef.ZoneSign)
	}

	if offsetHours != tzDef.OffsetHours {
		t.Errorf("Error: Expected tzDef.OffsetHours='%v'.\n" +
			"Instead, tzDef.OffsetHours='%v'\n", offsetHours, tzDef.OffsetHours)
	}

	if offsetMinutes != tzDef.OffsetMinutes {
		t.Errorf("Error: Expected tzDef.OffsetMinutes='%v'.\n" +
			"Instead, tzDef.OffsetMinutes='%v'\n", offsetMinutes, tzDef.OffsetMinutes)
	}

	if offsetSeconds != tzDef.OffsetSeconds {
		t.Errorf("Error: Expected tzDef.OffsetSeconds='%v'.\n" +
			"Instead, tzDef.OffsetSeconds='%v'\n", offsetSeconds, tzDef.OffsetSeconds)
	}

	if locationName != tzDef.LocationName {
		t.Errorf("Error: Expected tzDef.LocationName='%v'.\n" +
			"Instead, tzDef.LocationName='%v'\n", locationName, tzDef.LocationName)
	}

	if descStr != tzDef.Description {
		t.Errorf("Error: Expected tzDef.Description='%v'.\n" +
			"Instead, tzDef.Description='%v'\n", descStr, tzDef.Description)
	}

	tzDef0.Location = nil

	if tzDef0.Equal(tzDef) {
		t.Error("Error: Expected tzDef0 to be NOT EQUAL to tzDef. IT WAS EQUAL!")
	}

}
