package datetime

import (
	"fmt"
	"testing"
	"time"
)

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
