package datetime

import (
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

	tzDefDto, err := TimeZoneDefDto{}.NewFromTimeZoneName(TZones.America.New_York())

	if err != nil {
		t.Errorf("Error returned by dt.TimeZoneDefDto{}.NewFromTimeZoneName(" +
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
			dt1.Format(fmtStr), tzDefDto.GetLocationName() ,err.Error())
		return
	}

	dt2Str := dt2.Format(fmtStr)

	if expectedStr != dt2Str {
		t.Errorf("Error: Expected actual date time='%v'\n" +
			"Instead, actual date time='%v'\n", expectedStr, dt2Str)
	}
}