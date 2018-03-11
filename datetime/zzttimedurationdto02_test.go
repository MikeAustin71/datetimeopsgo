package datetime

import (
	"testing"
	"time"
)

func TestTimeDurationDto_DaylightSavings_01(t *testing.T) {

	// This test verifies duration over a daylight savings time threshold.

	locUSCentral, _ := time.LoadLocation(TzIanaUsCentral)

	t1USCentral := time.Date(2018, time.Month(3),10,18,0,0,0, locUSCentral)

	hoursDur := int64(24) * HourNanoSeconds

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationTzCalc(t1USCentral, time.Duration(hoursDur),
										TzIanaUsCentral,	TDurCalcTypeSTDYEARMTH, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dt.TimeDurationDto{}.New(t1USCentral, t2USCentral, fmtStr). " +
			"Error='%v'\n", err.Error())
	}

	outStr := t1Dur.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr)

	expectedStr := "2018-03-11 19:00:00.000000000 -0500 CDT"


	if expectedStr != outStr {
		t.Errorf("Error: Expected outStr='%v'.  Instead, outStr='%v'. ", expectedStr, outStr)
	}


/*

-- Gained an hour over Daylight savings threshold
Add Date Results - Cumulative Days
            Start Date Time:  2018-03-10 18:00:00.000000000 -0600 CST
      -- Duration = 24-Hours --
       Actual End Date Time:  2018-03-11 19:00:00.000000000 -0500 CDT
 */
}

