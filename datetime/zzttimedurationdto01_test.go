package datetime

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"time"
	"fmt"
)

type timedurdtoTestSuite struct {

	suite.Suite
	locUSCentral	*time.Location
	locUSPacific	*time.Location
	locParis			*time.Location
	locCairo			*time.Location
	locMoscow			*time.Location
	locTokyo			*time.Location
	t1USCentral				time.Time
	t1USPacific				time.Time
	t1EuropeParis			time.Time
	t1AfricaCairo			time.Time
	t1EuropeMoscow		time.Time
	t1AsiaTokyo				time.Time
	t2USCentral				time.Time
	t2USPacific				time.Time
	t2EuropeParis			time.Time
	t2AfricaCairo			time.Time
	t2EuropeMoscow		time.Time
	t2AsiaTokyo				time.Time
	t3USCentral				time.Time
	t3USPacific				time.Time
	t3EuropeParis			time.Time
	t3AfricaCairo			time.Time
	t3EuropeMoscow		time.Time
	t3AsiaTokyo				time.Time
	t4USCentral				time.Time
	t4USPacific				time.Time
	t4EuropeParis			time.Time
	t4AfricaCairo			time.Time
	t4EuropeMoscow		time.Time
	t4AsiaTokyo				time.Time

	
	fmtStr string
}

func (suite *timedurdtoTestSuite) SetupSuite() {
	suite.fmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"


	suite.locUSCentral, _ = time.LoadLocation(TzIanaUsCentral)
	suite.locUSPacific, _ = time.LoadLocation(TzIanaUsPacific)
	suite.locParis, _ = time.LoadLocation(TzIanaEuropeParis)
	suite.locCairo, _ = time.LoadLocation(TzIanaAfricaCairo)
	suite.locMoscow, _ = time.LoadLocation(TzIanaEuropeMoscow)
	suite.locTokyo, _ =	time.LoadLocation(TzIanaAsiaTokyo)
}

func (suite *timedurdtoTestSuite) TearDownSuite() {
	//suite.DtFmt = FormatDateTimeUtility{}
}

func (suite *timedurdtoTestSuite) SetupTest() {

	//t1 = "1948-09-07 04:32:16.008185431 -0500 CDT"
	suite.t1USCentral = time.Date(1948, time.Month(9),7,4,32,16,8185431,suite.locUSCentral)
	suite.t1USPacific = suite.t1USCentral.In(suite.locUSPacific)
	suite.t1EuropeParis = suite.t1USPacific.In(suite.locParis)
	suite.t1AfricaCairo = suite.t1EuropeParis.In(suite.locCairo)
	suite.t1EuropeMoscow = suite.t1AfricaCairo.In(suite.locMoscow)
	suite.t1AsiaTokyo = suite.t1EuropeMoscow.In(suite.locTokyo)

	//t2 = "2014-02-15 19:54:30.987654321 -0600 CST"
	suite.t2USCentral = time.Date(2014, time.Month(2),15,19,54,30,987654321,suite.locUSCentral)
	suite.t2USPacific = suite.t2USCentral.In(suite.locUSPacific)
	suite.t2EuropeParis = suite.t2USPacific.In(suite.locParis)
	suite.t2AfricaCairo = suite.t2EuropeParis.In(suite.locCairo)
	suite.t2EuropeMoscow = suite.t2AfricaCairo.In(suite.locMoscow)
	suite.t2AsiaTokyo = suite.t2EuropeMoscow.In(suite.locTokyo)

	//t3 = "2017-04-30 22:58:32.628149653 -0500 CDT"
	suite.t3USCentral = time.Date(2017, time.Month(4),30,22,58,32,628149653,suite.locUSCentral)
	suite.t3USPacific = suite.t3USCentral.In(suite.locUSPacific)
	suite.t3EuropeParis = suite.t3USPacific.In(suite.locParis)
	suite.t3AfricaCairo = suite.t3EuropeParis.In(suite.locCairo)
	suite.t3EuropeMoscow = suite.t3AfricaCairo.In(suite.locMoscow)
	suite.t3AsiaTokyo = suite.t3EuropeMoscow.In(suite.locTokyo)

	//t4 = "2018-03-06 20:02:18.792489279 -0600 CST"
	suite.t4USCentral = time.Date(2018, time.Month(3),06,20,02,18,792489279,suite.locUSCentral)
	suite.t4USPacific = suite.t4USCentral.In(suite.locUSPacific)
	suite.t4EuropeParis = suite.t4USPacific.In(suite.locParis)
	suite.t4AfricaCairo = suite.t4EuropeParis.In(suite.locCairo)
	suite.t4EuropeMoscow = suite.t4AfricaCairo.In(suite.locMoscow)
	suite.t4AsiaTokyo = suite.t4EuropeMoscow.In(suite.locTokyo)

}

func (suite *timedurdtoTestSuite) TearDownTest() {

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_New_01() {

	t1Dur, err := TimeDurationDto{}.New(suite.t2USCentral, suite.t3USCentral, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays) )

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMthCalc" == t1Dur.CalcType.String(),"Error: CalcType String NOT EQUAL to Std Calc!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
					t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
						t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
							t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

	/*
	3-Years 2-Months 15-Days 3-Hours 4-Minutes 1-Seconds 640-Milliseconds 495-Microseconds 332-Nanoseconds
	Caclulated End Date:  2017-04-30 22:58:32.628149653 -0500 CDT
		Expected End Date:  2017-04-30 22:58:32.628149653 -0500 CDT
	Time Duration Dto
			 StartTimeDateTz:  2014-02-15 19:54:30.987654321 -0600 CST
				 EndTimeDateTz:  2017-04-30 22:58:32.628149653 -0500 CDT
					TimeDuration:  101095441640495332
							CalcType:  StdYearMthCalc
								 Years:  3
				 YearsNanosecs:  94694400000000000
								Months:  2
				MonthsNanosecs:  5094000000000000
								 Weeks:  2
				 WeeksNanosecs:  1209600000000000
							WeekDays:  1
			WeekDaysNanosecs:  86400000000000
							DateDays:  15
			DateDaysNanosecs:  1296000000000000
								 Hours:  3
				 HoursNanosecs:  10800000000000
							 Minutes:  4
			 MinutesNanosecs:  240000000000
							 Seconds:  1
			 SecondsNanosecs:  1000000000
					Milliseconds:  640
	MillisecondsNanosecs:  640000000
					Microseconds:  495
	MicrosecondsNanosecs:  495000
					 Nanoseconds:  332
	-----------------------------------------------------
	TotSubSecNanoseconds:  640495332
		TotDateNanoseconds:  101084400000000000
		TotTimeNanoseconds:  11041640495332
	-----------------------------------------------------
	Check Total:
		 Date + Time Nanoseconds:  101095441640495332
	Total Duration Nanoseconds:  101095441640495332
	-----------------------------------------------------

	 */


}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_New_02() {

	// In this test, t3 is submitted with a Moscow time zone. It should be automatically
	// converted to US Central time matching the first time component (start time).
	// Results should be same as using suite.t3USCentral
	t1Dur, err := TimeDurationDto{}.New(suite.t2USCentral, suite.t3EuropeMoscow, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays) )

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMthCalc" == t1Dur.CalcType.String(),"Error: CalcType String NOT EQUAL to Std Calc!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
					t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
						t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
							t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesTz_01() {

	// In this test, t2 is submitted as a Tokyo Time Zone and t3 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified.
	t1Dur, err := TimeDurationDto{}.NewStartEndTimesTz(suite.t2AsiaTokyo, suite.t3AfricaCairo, TzIanaUsCentral,  suite.fmtStr)

	assert.Nil(suite.T(),err,"Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays) )

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMthCalc" == t1Dur.CalcType.String(),"Error: CalcType String NOT EQUAL to Std Calc!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
					t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
						t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
							t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesTzCalc_01() {

	// In this test, t2 is submitted as a Tokyo Time Zone and t3 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Also,
	// The calculation type is specified as "Standard".
	t1Dur, err := TimeDurationDto{}.NewStartEndTimesTzCalc(suite.t2AsiaTokyo, suite.t3AfricaCairo, TzIanaUsCentral, TDurCalcTypeSTDYEARMTH, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays) )

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMthCalc" == t1Dur.CalcType.String(),"Error: CalcType String NOT EQUAL to Std Calc!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
					t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
						t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
							t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")
}


func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesCalc_01() {

	t1Dur, err := TimeDurationDto{}.NewStartEndTimesCalc(suite.t2USCentral, suite.t3USCentral, TDurCalcTypeSTDYEARMTH, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays) )

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMthCalc" == t1Dur.CalcType.String(),"Error: CalcType String NOT EQUAL to Std Calc!")

	expectedDateNanosecs := int64(tx1.Sub(suite.t2USCentral))

	assert.Equal(suite.T(), expectedDateNanosecs, t1Dur.TotDateNanoseconds ,"Expected Date Nanoseconds DOES NOT EQUAL t1Dur.TotDateNanoseconds!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesDateDto_01() {

	dTzStart, err := DateTzDto{}.New(suite.t2USCentral, suite.fmtStr)
	assert.Nil(suite.T(),err,"Error DateTzDto{}.New(suite.t2AsiaTokyo): ")

	dTzEnd, err := DateTzDto{}.New(suite.t3AfricaCairo, suite.fmtStr)
	assert.Nil(suite.T(),err,"Error DateTzDto{}.New(suite.t2AsiaTokyo): ")

	t1Dur, err := TimeDurationDto{}.NewStartEndTimesDateDto(dTzStart, dTzEnd, suite.fmtStr)

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays) )

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMthCalc" == t1Dur.CalcType.String(),"Error: CalcType String NOT EQUAL to Std Calc!")

	expectedDateNanosecs := int64(tx1.Sub(suite.t2USCentral))

	assert.Equal(suite.T(), expectedDateNanosecs, t1Dur.TotDateNanoseconds ,"Expected Date Nanoseconds DOES NOT EQUAL t1Dur.TotDateNanoseconds!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")


}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesDateDtoTzCalc_01() {

	// In this test, t2 is submitted as a Tokyo Time Zone and t3 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Also,
	// The calculation type is specified as "Standard".

	dTzStart, err := DateTzDto{}.New(suite.t2AsiaTokyo, suite.fmtStr)
	assert.Nil(suite.T(),err,"Error DateTzDto{}.New(suite.t2AsiaTokyo): ")

	dTzEnd, err := DateTzDto{}.New(suite.t3AfricaCairo, suite.fmtStr)

	t1Dur, err := TimeDurationDto{}.NewStartEndTimesDateDtoTzCalc(dTzStart, dTzEnd, TzIanaUsCentral, TDurCalcTypeSTDYEARMTH, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays) )

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMthCalc" == t1Dur.CalcType.String(),"Error: CalcType String NOT EQUAL to Std Calc!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
					t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
						t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
							t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesDateDtoCalc_01() {

	// In this test, t2 is submitted as a Tokyo Time Zone and t3 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Also,
	// The calculation type is specified as "Standard".

	dTzStart, err := DateTzDto{}.New(suite.t2USCentral, suite.fmtStr)
	assert.Nil(suite.T(),err,"Error DateTzDto{}.New(suite.t2AsiaTokyo): ")

	dTzEnd, err := DateTzDto{}.New(suite.t3AfricaCairo, suite.fmtStr)

	t1Dur, err := TimeDurationDto{}.NewStartEndTimesDateDtoCalc(dTzStart, dTzEnd, TDurCalcTypeSTDYEARMTH, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,"Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays) )

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(),expectedEndDate.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMthCalc" == t1Dur.CalcType.String(),"Error: CalcType String NOT EQUAL to Std Calc!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
					t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
						t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
							t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDuration_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	t1Dur, err := TimeDurationDto{}.NewStartTimeDuration(suite.t1USCentral, actualTimeDuration, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error NewStartTimeDuration:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr) )
	assert.True(suite.T(),t1Dur.EndTimeDateTz.DateTime.Equal(suite.t4USCentral),s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),tx2.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds ,"Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds ,"Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration ,"Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationTz_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationTz(suite.t1USCentral, actualTimeDuration, TzIanaUsCentral, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error NewStartTimeDuration:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr) )
	assert.True(suite.T(),t1Dur.EndTimeDateTz.DateTime.Equal(suite.t4USCentral),s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),tx2.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds ,"Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds ,"Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration ,"Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationTzCalc_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationTzCalc(suite.t1USCentral, actualTimeDuration, TzIanaUsCentral, TDurCalcTypeSTDYEARMTH, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error NewStartTimeDuration:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr) )
	assert.True(suite.T(),t1Dur.EndTimeDateTz.DateTime.Equal(suite.t4USCentral),s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),tx2.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds ,"Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds ,"Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration ,"Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationCalc_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationCalc(suite.t1USCentral, actualTimeDuration, TDurCalcTypeSTDYEARMTH, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error NewStartTimeDuration:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr) )
	assert.True(suite.T(),t1Dur.EndTimeDateTz.DateTime.Equal(suite.t4USCentral),s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),tx2.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds ,"Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds ,"Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration ,"Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationDateDto_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	dTz, err := DateTzDto{}.New(suite.t1USCentral, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error  DateTzDto{}.New(suite.t1USCentral, suite.fmtStr):")

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationDateDto(dTz, actualTimeDuration, suite.fmtStr)

	assert.Nil(suite.T(),err,"Error NewStartTimeDurationDateDto:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr) )
	assert.True(suite.T(),t1Dur.EndTimeDateTz.DateTime.Equal(suite.t4USCentral),s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(),tx2.Equal(t1Dur.EndTimeDateTz.DateTime),"Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds ,"Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds ,"Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration ,"Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}


func TestTimeDuroTestSuite(t *testing.T) {
	tests := new(timedurdtoTestSuite)
	suite.Run(t, tests)
}