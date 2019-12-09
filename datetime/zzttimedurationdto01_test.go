package datetime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type timedurdtoTestSuite struct {
	suite.Suite
	locUSCentral   *time.Location
	locUSPacific   *time.Location
	locParis       *time.Location
	locCairo       *time.Location
	locMoscow      *time.Location
	locTokyo       *time.Location
	t1USCentral    time.Time
	t1USPacific    time.Time
	t1EuropeParis  time.Time
	t1AfricaCairo  time.Time
	t1EuropeMoscow time.Time
	t1AsiaTokyo    time.Time
	t2USCentral    time.Time
	t2USPacific    time.Time
	t2EuropeParis  time.Time
	t2AfricaCairo  time.Time
	t2EuropeMoscow time.Time
	t2AsiaTokyo    time.Time
	t3USCentral    time.Time
	t3USPacific    time.Time
	t3EuropeParis  time.Time
	t3AfricaCairo  time.Time
	t3EuropeMoscow time.Time
	t3AsiaTokyo    time.Time
	t4USCentral    time.Time
	t4USPacific    time.Time
	t4EuropeParis  time.Time
	t4AfricaCairo  time.Time
	t4EuropeMoscow time.Time
	t4AsiaTokyo    time.Time
	t5USCentral    time.Time
	t5USPacific    time.Time
	t5EuropeParis  time.Time
	t5AfricaCairo  time.Time
	t5EuropeMoscow time.Time
	t5AsiaTokyo    time.Time

	fmtStr string
}

func (suite *timedurdtoTestSuite) SetupSuite() {
	suite.fmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"

	suite.locUSCentral, _ = time.LoadLocation(TZones.US.Central())
	suite.locUSPacific, _ = time.LoadLocation(TZones.US.Pacific())
	suite.locParis, _ = time.LoadLocation(TZones.Europe.Paris())
	suite.locCairo, _ = time.LoadLocation(TZones.Africa.Cairo())
	suite.locMoscow, _ = time.LoadLocation(TZones.Europe.Moscow())
	suite.locTokyo, _ = time.LoadLocation(TZones.Asia.Tokyo())
}

func (suite *timedurdtoTestSuite) TearDownSuite() {
	// suite.DtFmt = FormatDateTimeUtility{}
}

func (suite *timedurdtoTestSuite) SetupTest() {

	// t1 = "1948-09-07 04:32:16.008185431 -0500 CDT"
	suite.t1USCentral = time.Date(
		1948,
		time.Month(9),
		7,
		4,
		32,
		16,
		8185431,
		suite.locUSCentral)

	suite.t1USPacific = suite.t1USCentral.In(suite.locUSPacific)
	suite.t1EuropeParis = suite.t1USPacific.In(suite.locParis)
	suite.t1AfricaCairo = suite.t1EuropeParis.In(suite.locCairo)
	suite.t1EuropeMoscow = suite.t1AfricaCairo.In(suite.locMoscow)
	suite.t1AsiaTokyo = suite.t1EuropeMoscow.In(suite.locTokyo)

	// t2 = "2014-02-15 19:54:30.987654321 -0600 CST"
	suite.t2USCentral = time.Date(2014, time.Month(2), 15, 19, 54, 30, 987654321, suite.locUSCentral)
	suite.t2USPacific = suite.t2USCentral.In(suite.locUSPacific)
	suite.t2EuropeParis = suite.t2USPacific.In(suite.locParis)
	suite.t2AfricaCairo = suite.t2EuropeParis.In(suite.locCairo)
	suite.t2EuropeMoscow = suite.t2AfricaCairo.In(suite.locMoscow)
	suite.t2AsiaTokyo = suite.t2EuropeMoscow.In(suite.locTokyo)

	// t3 = "2017-04-30 22:58:32.628149653 -0500 CDT"
	suite.t3USCentral = time.Date(2017, time.Month(4), 30, 22, 58, 32, 628149653, suite.locUSCentral)
	suite.t3USPacific = suite.t3USCentral.In(suite.locUSPacific)
	suite.t3EuropeParis = suite.t3USPacific.In(suite.locParis)
	suite.t3AfricaCairo = suite.t3EuropeParis.In(suite.locCairo)
	suite.t3EuropeMoscow = suite.t3AfricaCairo.In(suite.locMoscow)
	suite.t3AsiaTokyo = suite.t3EuropeMoscow.In(suite.locTokyo)

	// t4 = "2018-03-06 20:02:18.792489279 -0600 CST"
	suite.t4USCentral = time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, suite.locUSCentral)
	suite.t4USPacific = suite.t4USCentral.In(suite.locUSPacific)
	suite.t4EuropeParis = suite.t4USPacific.In(suite.locParis)
	suite.t4AfricaCairo = suite.t4EuropeParis.In(suite.locCairo)
	suite.t4EuropeMoscow = suite.t4AfricaCairo.In(suite.locMoscow)
	suite.t4AsiaTokyo = suite.t4EuropeMoscow.In(suite.locTokyo)

	// t5 = "2018-07-04 15:09:05.458621349 -0500 CDT"
	suite.t5USCentral = time.Date(2018, time.Month(7), 04, 15, 9, 5, 458621349, suite.locUSCentral)
	suite.t5USPacific = suite.t5USCentral.In(suite.locUSPacific)
	suite.t5EuropeParis = suite.t5USPacific.In(suite.locParis)
	suite.t5AfricaCairo = suite.t5EuropeParis.In(suite.locCairo)
	suite.t5EuropeMoscow = suite.t5AfricaCairo.In(suite.locMoscow)
	suite.t5AsiaTokyo = suite.t5EuropeMoscow.In(suite.locTokyo)

}

func (suite *timedurdtoTestSuite) TearDownTest() {

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_New_01() {

	t1Dur, err := TimeDurationDto{}.New(suite.t2USCentral, suite.t3USCentral, suite.fmtStr)

	assert.Nil(suite.T(), err, "Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMth" == t1Dur.CalcType.String(), "Error: CalcType String NOT EQUAL to StdYearMth!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur, "Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur), "Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

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

	assert.Nil(suite.T(), err, "Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMth" == t1Dur.CalcType.String(), "Error: CalcType String NOT EQUAL to StdYearMth!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur, "Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur), "Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewAutoEnd_01() {

	t1Dur, err := TimeDurationDto{}.NewAutoEnd(suite.t1AsiaTokyo, TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewAutoEnd() :")

	assert.Equal(suite.T(), TZones.US.Central(), t1Dur.StartTimeDateTz.TimeZone.LocationName, "Expected Start Time Zone NOT EQUAL To Actual Start Time Zone!")

	assert.Equal(suite.T(), TZones.US.Central(), t1Dur.EndTimeDateTz.TimeZone.LocationName, "Expected Start Time Zone NOT EQUAL To Actual Start Time Zone!")

	assert.True(suite.T(), suite.t1USCentral.Equal(t1Dur.StartTimeDateTz.GetDateTimeValue()), "Error: Expected Starting Date Time NOT EQUAL to t1Dur.StartTimeDateTz!")

	cLoc, _ := time.LoadLocation(TZones.US.Central())

	checkTime := time.Now().In(cLoc)

	checkDur := checkTime.Sub(t1Dur.EndTimeDateTz.GetDateTimeValue())

	testMax := time.Duration(int64(2) * int64(time.Second))

	s := fmt.Sprintf("Error: checkDur > testMax Nanoseconds!. checkDur='%v'  testMax='%v'", checkDur, testMax)
	assert.True(suite.T(), testMax > checkDur, s)
}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewAutoStart_01() {

	t1Dur, err := TimeDurationDto{}.NewAutoStart(TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewAutoStart() :")

	assert.Equal(suite.T(), TZones.US.Central(), t1Dur.StartTimeDateTz.TimeZone.LocationName, "Expected Start Time Zone NOT EQUAL To Actual Start Time Zone!")

	assert.Equal(suite.T(), TZones.US.Central(), t1Dur.EndTimeDateTz.TimeZone.LocationName, "Expected Start Time Zone NOT EQUAL To Actual Start Time Zone!")

	assert.True(suite.T(),
		suite.fmtStr == t1Dur.StartTimeDateTz.GetDateTimeFmt(),
		"Error: Expected suite.fmtSTr to EQUAL t1Dur.StartTimeDateTz.DateTimeFmt. THEY ARE NOT EQUAL!")

	err = t1Dur.SetAutoEnd()

	assert.Nil(suite.T(), err, "Error t1Dur.SetAutoEnd() :")

	checkDur := t1Dur.TimeDuration
	testMax := time.Duration(int64(2) * int64(time.Second))

	s := fmt.Sprintf("Error: checkDur > testMax Nanoseconds!. checkDur='%v'  testMax='%v'", checkDur, testMax)
	assert.True(suite.T(), testMax > checkDur, s)

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesTz_01() {

	// In this test, t2 is submitted as a Tokyo Time Zone and t3 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified.
	t1Dur, err := TimeDurationDto{}.NewStartEndTimesTz(suite.t2AsiaTokyo, suite.t3AfricaCairo, TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMth" == t1Dur.CalcType.String(), "Error: CalcType String NOT EQUAL to StdYearMth!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur, "Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur), "Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesTzCalc_01() {

	// In this test, t2 is submitted as a Tokyo Time Zone and t3 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Also,
	// The calculation type is specified as "Standard".
	t1Dur, err := TimeDurationDto{}.NewStartEndTimesCalcTz(suite.t2AsiaTokyo, suite.t3AfricaCairo,
		TDurCalcType(0).StdYearMth(), TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMth" == t1Dur.CalcType.String(), "Error: CalcType String NOT EQUAL to StdYearMth!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur, "Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur), "Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")
}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesCalc_01() {

	t1Dur, err := TimeDurationDto{}.NewStartEndTimesCalc(suite.t2USCentral, suite.t3USCentral, TDurCalcType(0).StdYearMth(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMth" == t1Dur.CalcType.String(), "Error: CalcType String NOT EQUAL to StdYearMth!")

	expectedDateNanosecs := int64(tx1.Sub(suite.t2USCentral))

	assert.Equal(suite.T(), expectedDateNanosecs, t1Dur.TotDateNanoseconds, "Expected Date Nanoseconds DOES NOT EQUAL t1Dur.TotDateNanoseconds!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur, "Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur), "Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesDateDto_01() {

	dTzStart, err := DateTzDto{}.New(suite.t2USCentral, suite.fmtStr)
	assert.Nil(suite.T(), err, "Error DateTzDto{}.New(suite.t2AsiaTokyo): ")

	dTzEnd, err := DateTzDto{}.New(suite.t3AfricaCairo, suite.fmtStr)
	assert.Nil(suite.T(), err, "Error DateTzDto{}.New(suite.t2AsiaTokyo): ")

	t1Dur, err := TimeDurationDto{}.NewStartEndTimesDateDto(dTzStart, dTzEnd, suite.fmtStr)

	assert.Nil(suite.T(), err,
		"Error TimeDurationDto{}.NewStartEndTimesDateDto(dTzStart, dTzEnd, suite.fmtStr): " )

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMth" == t1Dur.CalcType.String(), "Error: CalcType String NOT EQUAL to StdYearMth!")

	expectedDateNanosecs := int64(tx1.Sub(suite.t2USCentral))

	assert.Equal(suite.T(), expectedDateNanosecs, t1Dur.TotDateNanoseconds, "Expected Date Nanoseconds DOES NOT EQUAL t1Dur.TotDateNanoseconds!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur, "Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur), "Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesDateDtoTzCalc_01() {

	// In this test, t2 is submitted as a Tokyo Time Zone and t3 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Also,
	// The calculation type is specified as "Standard".

	dTzStart, err := DateTzDto{}.New(suite.t2AsiaTokyo, suite.fmtStr)
	assert.Nil(suite.T(), err, "Error DateTzDto{}.New(suite.t2AsiaTokyo): ")

	dTzEnd, err := DateTzDto{}.New(suite.t3AfricaCairo, suite.fmtStr)

	t1Dur, err := TimeDurationDto{}.NewStartEndTimesDateTzDtoCalcTz(dTzStart, dTzEnd,
		TDurCalcType(0).StdYearMth(), TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMth" == t1Dur.CalcType.String(), "Error: CalcType String NOT EQUAL to StdYearMth!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur, "Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur), "Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartEndTimesDateDtoCalc_01() {

	// In this test, t2 is submitted as a Tokyo Time Zone and t3 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Also,
	// The calculation type is specified as "Standard".

	dTzStart, err := DateTzDto{}.New(suite.t2USCentral, suite.fmtStr)
	assert.Nil(suite.T(), err, "Error DateTzDto{}.New(suite.t2AsiaTokyo): ")

	dTzEnd, err := DateTzDto{}.New(suite.t3AfricaCairo, suite.fmtStr)

	t1Dur, err := TimeDurationDto{}.NewStartEndTimesDateDtoCalc(dTzStart, dTzEnd, TDurCalcType(0).StdYearMth(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error:")

	expectedTimeDur := suite.t3USCentral.Sub(suite.t2USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t2USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "StdYearMth" == t1Dur.CalcType.String(), "Error: CalcType String NOT EQUAL to StdYearMth!")

	expectedTimeDur = suite.t3USCentral.Sub(suite.t2USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur, "Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur), "Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDuration_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	t1Dur, err := TimeDurationDto{}.NewStartTimeDuration(suite.t1USCentral, actualTimeDuration, suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartTimeDurationTz:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr))
	assert.True(suite.T(), t1Dur.EndTimeDateTz.GetDateTimeValue().Equal(suite.t4USCentral), s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), tx2.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds, "Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds, "Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration, "Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationTz_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationTz(suite.t1USCentral, actualTimeDuration, TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartTimeDurationTz:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr))
	assert.True(suite.T(), t1Dur.EndTimeDateTz.GetDateTimeValue().Equal(suite.t4USCentral), s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), tx2.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds, "Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds, "Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration, "Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationTzCalc_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationCalcTz(suite.t1USCentral, actualTimeDuration,
		TDurCalcType(0).StdYearMth(), TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartTimeDurationTz:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr))
	assert.True(suite.T(), t1Dur.EndTimeDateTz.GetDateTimeValue().Equal(suite.t4USCentral), s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), tx2.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds, "Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds, "Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration, "Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationCalc_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationCalc(suite.t1USCentral, actualTimeDuration, TDurCalcType(0).StdYearMth(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartTimeDurationTz:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr))
	assert.True(suite.T(), t1Dur.EndTimeDateTz.GetDateTimeValue().Equal(suite.t4USCentral), s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), tx2.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds, "Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds, "Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration, "Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationDateDto_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	dTz, err := DateTzDto{}.New(suite.t1USCentral, suite.fmtStr)

	assert.Nil(suite.T(), err, "Error  DateTzDto{}.New(suite.t1USCentral, suite.fmtStr):")

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationDateDto(dTz, actualTimeDuration, suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartTimeDurationDateDto:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr))
	assert.True(suite.T(), t1Dur.EndTimeDateTz.GetDateTimeValue().Equal(suite.t4USCentral), s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), tx2.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds, "Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds, "Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration, "Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationDateDtoTz_01() {

	dTz, err := DateTzDto{}.New(suite.t1USCentral, suite.fmtStr)

	assert.Nil(suite.T(), err, "Error DateTzDto{}.New(suite.t4USCentral, suite.fmtStr):")

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationDateDtoTz(dTz, actualTimeDuration, TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartTimeDurationDateDtoTz:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr))
	assert.True(suite.T(), t1Dur.EndTimeDateTz.GetDateTimeValue().Equal(suite.t4USCentral), s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), tx2.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds, "Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds, "Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration, "Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")
}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationDateDtoTzCalc_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	dTz, err := DateTzDto{}.New(suite.t1USCentral, suite.fmtStr)

	assert.Nil(suite.T(), err, "Error DateTzDto{}.New(suite.t1USCentral, suite.fmtStr):")

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationDateDtoTzCalc(dTz, actualTimeDuration, TZones.US.Central(), TDurCalcType(0).StdYearMth(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartTimeDurationTz:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr))
	assert.True(suite.T(), t1Dur.EndTimeDateTz.GetDateTimeValue().Equal(suite.t4USCentral), s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), tx2.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds, "Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds, "Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration, "Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeDurationDateDtoCalc_01() {

	actualTimeDuration := suite.t4USCentral.Sub(suite.t1USCentral)

	dTz, err := DateTzDto{}.New(suite.t1USCentral, suite.fmtStr)

	assert.Nil(suite.T(), err, "Error DateTzDto{}.New(suite.t1USCentral, suite.fmtStr):")

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationDateDtoCalc(dTz, actualTimeDuration, TDurCalcType(0).StdYearMth(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartTimeDurationTz:")

	s := fmt.Sprintf("Error: Expected EndDateTime NOT EQUAL to t1Dur.EndDateTime! t1Dur.EndTime='%v' t4USCentral='%v' ", t1Dur.EndTimeDateTz.String(), suite.t4USCentral.Format(suite.fmtStr))
	assert.True(suite.T(), t1Dur.EndTimeDateTz.GetDateTimeValue().Equal(suite.t4USCentral), s)

	tx1 := suite.t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), tx2.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated EndDateTime (tx2) NOT EQUAL to t1Dur.EndTimeDateTz!")

	assert.True(suite.T(), dur == t1Dur.TotTimeNanoseconds, "Error: Summary Time NOT EQUAL to t1Dur.TotTimeNanoseconds!")

	duration := tx1.Sub(suite.t1USCentral)

	assert.True(suite.T(), int64(duration) == t1Dur.TotDateNanoseconds, "Error: Calculated Date Duration NOT EQUAL to t1Dur.TotDateNanoseconds!")

	assert.True(suite.T(), actualTimeDuration == t1Dur.TimeDuration, "Error: Actual Duration DOES NOT EQUAL t1Dur.TimeDuration!")
}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimePlusTimeDto_01() {

	year := 69
	month := 5
	day := 27
	hour := 15
	minute := 30
	second := 2
	millisecond := 784
	microsecond := 303
	nanosecond := 848

	tDto, err := TimeDto{}.New(year, month, 0, day, hour, minute, second,
		millisecond, microsecond, nanosecond)

	assert.Nil(suite.T(), err, "Error TimeDto{}.New(year, month, ...):")

	tDur, err := TimeDurationDto{}.NewStartTimePlusTimeDto(suite.t1USCentral, tDto, suite.fmtStr)

	assert.Nil(suite.T(), err, "Error TimeDurationDto{}.NewStartTimePlusTimeDto(...):")

	assert.True(suite.T(), suite.t4USCentral.Equal(tDur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected EndDateTime (suite.t4USCentral) NOT EQUAL to t1Dur.EndTimeDateTz!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_NewStartTimeMinusTimeDto_01() {

	year := 69
	month := 5
	day := 27
	hour := 15
	minute := 30
	second := 2
	millisecond := 784
	microsecond := 303
	nanosecond := 848

	tDto, err := TimeDto{}.New(year, month, 0, day, hour, minute, second,
		millisecond, microsecond, nanosecond)

	assert.Nil(suite.T(), err, "Error TimeDto{}.New(year, month, ...):")

	tDur, err := TimeDurationDto{}.NewEndTimeMinusTimeDto(suite.t4USCentral, tDto, suite.fmtStr)

	assert.Nil(suite.T(), err, "Error TimeDurationDto{}.NewStartTimePlusTimeDto(...):")

	assert.True(suite.T(), suite.t4USCentral.Equal(tDur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected EndDateTime (suite.t4USCentral) NOT EQUAL to t1Dur.EndTimeDateTz!")

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_ReCalcEndDateTimeToNow_01() {

	tDur, err := TimeDurationDto{}.New(suite.t1AsiaTokyo, suite.t2AsiaTokyo, suite.fmtStr)

	assert.Nil(suite.T(), err, "Error TimeDurationDto{}.New(...):")

	err = tDur.ReCalcEndDateTimeToNow()

	assert.Nil(suite.T(), err, "Error TimeDurationDto{}.ReCalcEndDateTimeToNow():")

	s := fmt.Sprintf("Expected StartDateTime Time Zone='%v'. "+
		"Instead StartDateTime TimeZone='%v'",
		TZones.Asia.Tokyo(), tDur.StartTimeDateTz.TimeZone.LocationName)

	assert.Equal(suite.T(), TZones.Asia.Tokyo(), tDur.StartTimeDateTz.TimeZone.LocationName, s)

	s = fmt.Sprintf("Expected EndDateTime Time Zone='%v'. "+
		"Instead EndDateTime TimeZone='%v'",
		TZones.Asia.Tokyo(), tDur.EndTimeDateTz.TimeZone.LocationName)

	assert.Equal(suite.T(), TZones.Asia.Tokyo(), tDur.EndTimeDateTz.TimeZone.LocationName, s)

	assert.True(suite.T(), suite.t1AsiaTokyo.Equal(tDur.StartTimeDateTz.GetDateTimeValue()),
		"Error: Expected StartDateTime (suite.t1AsiaTokyo) NOT EQUAL to t1Dur.StartTimeDateTz!")

	testMax := time.Duration(int64(2) * int64(time.Second))

	tokyoNow := time.Now().In(suite.t1AsiaTokyo.Location())

	actualDur := tokyoNow.Sub(tDur.EndTimeDateTz.GetDateTimeValue())

	s = fmt.Sprintf("Error: Expected actual duration since Now to be less than '%v'. "+
		"Instead, actual duration ='%v'.", testMax, actualDur)

	assert.True(suite.T(), actualDur < testMax, s)

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_TestCumMonths_01() {

	// In this test, t4 is submitted as a Tokyo Time Zone and t5 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Note,
	// The calculation type is specified as Cumulative Months.
	t1Dur, err := TimeDurationDto{}.NewStartEndTimesCalcTz(suite.t4AsiaTokyo, suite.t5AfricaCairo,
		TDurCalcType(0).CumMonths(), TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error:")

	expectedTimeDur := suite.t5USCentral.Sub(suite.t4USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t4USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()), "Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "CumMonths" == t1Dur.CalcType.String(),
		"Error: CalcType String NOT EQUAL to CumMonths!")

	expectedTimeDur = suite.t5USCentral.Sub(suite.t4USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur, "Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur), "Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

	expectedOutStr := "3-Months 27-Days 19-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds"
	actualOutStr, err := t1Dur.GetCumMonthsDaysTimeStr()

	assert.Nil(suite.T(), err, "Error t1Dur.GetCumMonthsDaysTimeStr():")

	s := fmt.Sprintf("Expected OutStr='%v'. Instead OutStr='%v'", expectedOutStr, actualOutStr)
	assert.Equal(suite.T(), expectedOutStr, actualOutStr, s)

	expectedTime := int64(3)
	actualTime := t1Dur.Months
	s = fmt.Sprintf("Expected Cumulative Months='%v'. Instead Cumulative Months='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(3)
	actualTime = t1Dur.Weeks

	s = fmt.Sprintf("Expected Weeks='%v'. Instead Weeks='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(6)
	actualTime = t1Dur.WeekDays

	s = fmt.Sprintf("Expected WeekDays='%v'. Instead WeekDays='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(27)
	actualTime = t1Dur.DateDays

	s = fmt.Sprintf("Expected Days='%v'. Instead Days='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(19)
	actualTime = t1Dur.Hours

	s = fmt.Sprintf("Expected Hours='%v'. Instead Hours='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(6)
	actualTime = t1Dur.Minutes

	s = fmt.Sprintf("Expected Minutes='%v'. Instead Minutes='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(46)
	actualTime = t1Dur.Seconds

	s = fmt.Sprintf("Expected Seconds='%v'. Instead Seconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(666)
	actualTime = t1Dur.Milliseconds

	s = fmt.Sprintf("Expected Milliseconds='%v'. Instead Milliseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(132)
	actualTime = t1Dur.Microseconds

	s = fmt.Sprintf("Expected Microseconds='%v'. Instead Microseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(70)
	actualTime = t1Dur.Nanoseconds

	s = fmt.Sprintf("Expected Nanoseconds='%v'. Instead Nanoseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(10346806666132070)
	actualTime = int64(t1Dur.TimeDuration)

	s = fmt.Sprintf("Expected Time Duration='%v'. Instead Time Duration='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	/*
		Results Cumulative Months:
		3-Months 27-Days 19-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds
		Time Duration Dto
				 StartTimeDateTz:  2018-03-06 20:02:18.792489279 -0600 CST
					 EndTimeDateTz:  2018-07-04 15:09:05.458621349 -0500 CDT
						TimeDuration:  10346806666132070
								CalcType:  CumMonthsCalc
									 Years:  0
					 YearsNanosecs:  0
									Months:  3
					MonthsNanosecs:  7945200000000000
									 Weeks:  3
					 WeeksNanosecs:  1814400000000000
								WeekDays:  6
				WeekDaysNanosecs:  518400000000000
								DateDays:  27
				DateDaysNanosecs:  2332800000000000
									 Hours:  19
					 HoursNanosecs:  68400000000000
								 Minutes:  6
				 MinutesNanosecs:  360000000000
								 Seconds:  46
				 SecondsNanosecs:  46000000000
						Milliseconds:  666
		MillisecondsNanosecs:  666000000
						Microseconds:  132
		MicrosecondsNanosecs:  132000
						 Nanoseconds:  70
		-----------------------------------------------------
		TotSubSecNanoseconds:  666132070
			TotDateNanoseconds:  10278000000000000
			TotTimeNanoseconds:  68806666132070
		-----------------------------------------------------
		Check Total:
			 Date + Time Nanoseconds:  10346806666132070
		Total Duration Nanoseconds:  10346806666132070
		-----------------------------------------------------
	*/

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_TestCumMonths_02() {

	// In this test, t4 is submitted as a Tokyo Time Zone and t5 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Note,
	// The calculation type is specified as Cumulative Months.
	t1Dur, err := TimeDurationDto{}.NewStartEndTimesTz(suite.t4AsiaTokyo, suite.t5AfricaCairo, TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartEndTimesTz :")

	err = t1Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumMonths())

	assert.Nil(suite.T(), err, "Error ReCalcTimeDurationAllocation:")

	expectedTimeDur := suite.t5USCentral.Sub(suite.t4USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t4USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	expectedEndDate := tx1.Add(time.Duration(dur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()),
		"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	expectedEndDate = tx1.Add(time.Duration(t1Dur.TotTimeNanoseconds))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()),
		"Error: Tot Time Duration + End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "CumMonths" == t1Dur.CalcType.String(),
		"Error: CalcType String NOT EQUAL to CumMonths!")

	expectedTimeDur = suite.t5USCentral.Sub(suite.t4USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,
		"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),
		"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

	expectedOutStr := "3-Months 27-Days 19-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds"
	actualOutStr, err := t1Dur.GetCumMonthsDaysTimeStr()

	assert.Nil(suite.T(), err, "Error t1Dur.GetCumMonthsDaysTimeStr():")

	s := fmt.Sprintf("Expected OutStr='%v'. Instead OutStr='%v'", expectedOutStr, actualOutStr)
	assert.Equal(suite.T(), expectedOutStr, actualOutStr, s)

	expectedTime := int64(3)
	actualTime := t1Dur.Months
	s = fmt.Sprintf("Expected Cumulative Months='%v'. Instead Cumulative Months='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(3)
	actualTime = t1Dur.Weeks

	s = fmt.Sprintf("Expected Weeks='%v'. Instead Weeks='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(6)
	actualTime = t1Dur.WeekDays

	s = fmt.Sprintf("Expected WeekDays='%v'. Instead WeekDays='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(27)
	actualTime = t1Dur.DateDays

	s = fmt.Sprintf("Expected Days='%v'. Instead Days='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(19)
	actualTime = t1Dur.Hours

	s = fmt.Sprintf("Expected Hours='%v'. Instead Hours='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(6)
	actualTime = t1Dur.Minutes

	s = fmt.Sprintf("Expected Minutes='%v'. Instead Minutes='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(46)
	actualTime = t1Dur.Seconds

	s = fmt.Sprintf("Expected Seconds='%v'. Instead Seconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(666)
	actualTime = t1Dur.Milliseconds

	s = fmt.Sprintf("Expected Milliseconds='%v'. Instead Milliseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(132)
	actualTime = t1Dur.Microseconds

	s = fmt.Sprintf("Expected Microseconds='%v'. Instead Microseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(70)
	actualTime = t1Dur.Nanoseconds

	s = fmt.Sprintf("Expected Nanoseconds='%v'. Instead Nanoseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(10346806666132070)
	actualTime = int64(t1Dur.TimeDuration)

	s = fmt.Sprintf("Expected Time Duration='%v'. Instead Time Duration='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_TestCumWeeks_01() {

	// In this test, t4 is submitted as a Tokyo Time Zone and t5 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Note,
	// the calculation type is specified as Cumulative Months.
	t1Dur, err := TimeDurationDto{}.NewStartEndTimesCalcTz(suite.t4AsiaTokyo, suite.t5AfricaCairo,
		TDurCalcType(0).CumWeeks(), TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error:")

	expectedTimeDur := suite.t5USCentral.Sub(suite.t4USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t4USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), 0)

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	timeDaysDur := dur + (WeekNanoSeconds * t1Dur.Weeks) + (DayNanoSeconds * t1Dur.WeekDays)

	expectedEndDate := tx1.Add(time.Duration(timeDaysDur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()),
		"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "CumWeeks" == t1Dur.CalcType.String(),
		"Error: CalcType String NOT EQUAL to CumWeeks!")

	expectedTimeDur = suite.t5USCentral.Sub(suite.t4USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,
		"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.WeeksNanosecs + t1Dur.WeekDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),
		"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

	expectedOutStr :=
		"17-Weeks 0-WeekDays 18-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds"

	actualOutStr, err := t1Dur.GetCumWeeksDaysTimeStr()

	assert.Nil(suite.T(), err, "Error t1Dur.GetCumWeeksDaysTimeStr():")

	s := fmt.Sprintf("Expected OutStr='%v'. Instead OutStr='%v'", expectedOutStr, actualOutStr)
	assert.Equal(suite.T(), expectedOutStr, actualOutStr, s)

	expectedTime := int64(0)
	actualTime := t1Dur.Months
	s = fmt.Sprintf("Expected Months='%v'. Instead Months='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(17)
	actualTime = t1Dur.Weeks

	s = fmt.Sprintf("Expected Cumulative Weeks='%v'. Instead Cumulative Weeks='%v'",
		expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(0)
	actualTime = t1Dur.WeekDays

	s = fmt.Sprintf("Expected WeekDays='%v'. Instead WeekDays='%v'",
		expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(0)
	actualTime = t1Dur.DateDays

	s = fmt.Sprintf("Expected Days='%v'. Instead Days='%v'",
		expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(18)
	actualTime = t1Dur.Hours

	s = fmt.Sprintf("Expected Hours='%v'. Instead Hours='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(6)
	actualTime = t1Dur.Minutes

	s = fmt.Sprintf("Expected Minutes='%v'. Instead Minutes='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(46)
	actualTime = t1Dur.Seconds

	s = fmt.Sprintf("Expected Seconds='%v'. Instead Seconds='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(666)
	actualTime = t1Dur.Milliseconds

	s = fmt.Sprintf("Expected Milliseconds='%v'. Instead Milliseconds='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(132)
	actualTime = t1Dur.Microseconds

	s = fmt.Sprintf("Expected Microseconds='%v'. Instead Microseconds='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(70)
	actualTime = t1Dur.Nanoseconds

	s = fmt.Sprintf("Expected Nanoseconds='%v'. Instead Nanoseconds='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(10346806666132070)
	actualTime = int64(t1Dur.TimeDuration)

	s = fmt.Sprintf("Expected Time Duration='%v'. Instead Time Duration='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	/*
		$ go run main.go
		Results Cumulative Weeks:
		17-Weeks 0-WeekDays 18-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds
		Time Duration Dto
				 StartTimeDateTz:  2018-03-06 20:02:18.792489279 -0600 CST
					 EndTimeDateTz:  2018-07-04 15:09:05.458621349 -0500 CDT
						TimeDuration:  10346806666132070
								CalcType:  CumWeeksCalc
									 Years:  0
					 YearsNanosecs:  0
									Months:  0
					MonthsNanosecs:  0
									 Weeks:  17
					 WeeksNanosecs:  10281600000000000
								WeekDays:  0
				WeekDaysNanosecs:  0
								DateDays:  0
				DateDaysNanosecs:  0
									 Hours:  18
					 HoursNanosecs:  64800000000000
								 Minutes:  6
				 MinutesNanosecs:  360000000000
								 Seconds:  46
				 SecondsNanosecs:  46000000000
						Milliseconds:  666
		MillisecondsNanosecs:  666000000
						Microseconds:  132
		MicrosecondsNanosecs:  132000
						 Nanoseconds:  70
		-----------------------------------------------------
		TotSubSecNanoseconds:  666132070
			TotDateNanoseconds:  10281600000000000
			TotTimeNanoseconds:  65206666132070
		-----------------------------------------------------
		Check Total:
			 Date + Time Nanoseconds:  10346806666132070
		Total Duration Nanoseconds:  10346806666132070
		-----------------------------------------------------
	*/

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_TestCumWeeks_02() {

	// In this test, t4 is submitted as a Tokyo Time Zone and t5 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Note,
	// the calculation type is specified as Cumulative Months.
	t1Dur, err := TimeDurationDto{}.NewStartEndTimesTz(suite.t4AsiaTokyo, suite.t5AfricaCairo, TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartEndTimesTz:")

	err = t1Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumWeeks())

	assert.Nil(suite.T(), err, "Error ReCalcTimeDurationAllocation:")

	expectedTimeDur := suite.t5USCentral.Sub(suite.t4USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,
		"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,
		"Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t4USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), 0)

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	timeDaysDur := dur + (WeekNanoSeconds * t1Dur.Weeks) + (DayNanoSeconds * t1Dur.WeekDays)

	expectedEndDate := tx1.Add(time.Duration(timeDaysDur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()),
		"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "CumWeeks" == t1Dur.CalcType.String(),
		"Error: CalcType String NOT EQUAL to CumWeeks!")

	expectedTimeDur = suite.t5USCentral.Sub(suite.t4USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,
		"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.WeeksNanosecs + t1Dur.WeekDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),
		"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

	expectedOutStr :=
		"17-Weeks 0-WeekDays 18-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds"

	actualOutStr, err := t1Dur.GetCumWeeksDaysTimeStr()

	assert.Nil(suite.T(), err, "Error t1Dur.GetCumWeeksDaysTimeStr():")

	s := fmt.Sprintf("Expected OutStr='%v'. Instead OutStr='%v'", expectedOutStr, actualOutStr)
	assert.Equal(suite.T(), expectedOutStr, actualOutStr, s)

	expectedTime := int64(0)
	actualTime := t1Dur.Months
	s = fmt.Sprintf("Expected Months='%v'. Instead Months='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(17)
	actualTime = t1Dur.Weeks

	s = fmt.Sprintf("Expected Cumulative Weeks='%v'. Instead Cumulative Weeks='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(0)
	actualTime = t1Dur.WeekDays

	s = fmt.Sprintf("Expected WeekDays='%v'. Instead WeekDays='%v'", expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(0)
	actualTime = t1Dur.DateDays

	s = fmt.Sprintf("Expected Days='%v'. Instead Days='%v'", expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(18)
	actualTime = t1Dur.Hours

	s = fmt.Sprintf("Expected Hours='%v'. Instead Hours='%v'", expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(6)
	actualTime = t1Dur.Minutes

	s = fmt.Sprintf("Expected Minutes='%v'. Instead Minutes='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(46)
	actualTime = t1Dur.Seconds

	s = fmt.Sprintf("Expected Seconds='%v'. Instead Seconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(666)
	actualTime = t1Dur.Milliseconds

	s = fmt.Sprintf("Expected Milliseconds='%v'. Instead Milliseconds='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(132)
	actualTime = t1Dur.Microseconds

	s = fmt.Sprintf("Expected Microseconds='%v'. Instead Microseconds='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(70)
	actualTime = t1Dur.Nanoseconds

	s = fmt.Sprintf("Expected Nanoseconds='%v'. Instead Nanoseconds='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(10346806666132070)
	actualTime = int64(t1Dur.TimeDuration)

	s = fmt.Sprintf("Expected Time Duration='%v'. Instead Time Duration='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_TestCumDays_01() {

	// In this test, t4 is submitted as a Tokyo Time Zone and t5 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Note,
	// the calculation type is specified as Cumulative Months.
	t1Dur, err := TimeDurationDto{}.NewStartEndTimesCalcTz(suite.t4AsiaTokyo, suite.t5AfricaCairo,
		TDurCalcType(0).CumDays(), TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartEndTimesCalcTz:")

	expectedTimeDur := suite.t5USCentral.Sub(suite.t4USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,
		"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration, "Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t4USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), 0)

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds
	dayTimeDur := dur + t1Dur.DateDays*DayNanoSeconds

	expectedEndDate := tx1.Add(time.Duration(dayTimeDur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()),
		"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "CumDays" == t1Dur.CalcType.String(),
		"Error: CalcType String NOT EQUAL to CumDays!")

	expectedTimeDur = suite.t5USCentral.Sub(suite.t4USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,
		"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),
		"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

	expectedOutStr :=
		"119-Days 18-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds"

	actualOutStr, err := t1Dur.GetCumDaysTimeStr()

	assert.Nil(suite.T(), err, "Error t1Dur.GetCumDaysTimeStr():")

	s := fmt.Sprintf("Expected OutStr='%v'. Instead OutStr='%v'", expectedOutStr, actualOutStr)
	assert.Equal(suite.T(), expectedOutStr, actualOutStr, s)

	expectedTime := int64(0)
	actualTime := t1Dur.Years
	s = fmt.Sprintf("Expected Years='%v'. Instead Years='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(0)
	actualTime = t1Dur.Months
	s = fmt.Sprintf("Expected Months='%v'. Instead Months='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(0)
	actualTime = t1Dur.Weeks

	s = fmt.Sprintf("Expected Cumulative Weeks='%v'. Instead Cumulative Weeks='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(0)
	actualTime = t1Dur.WeekDays

	s = fmt.Sprintf("Expected WeekDays='%v'. Instead WeekDays='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(119)
	actualTime = t1Dur.DateDays

	s = fmt.Sprintf("Expected Days='%v'. Instead Days='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(18)
	actualTime = t1Dur.Hours

	s = fmt.Sprintf("Expected Hours='%v'. Instead Hours='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(6)
	actualTime = t1Dur.Minutes

	s = fmt.Sprintf("Expected Minutes='%v'. Instead Minutes='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(46)
	actualTime = t1Dur.Seconds

	s = fmt.Sprintf("Expected Seconds='%v'. Instead Seconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(666)
	actualTime = t1Dur.Milliseconds

	s = fmt.Sprintf("Expected Milliseconds='%v'. Instead Milliseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(132)
	actualTime = t1Dur.Microseconds

	s = fmt.Sprintf("Expected Microseconds='%v'. Instead Microseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(70)
	actualTime = t1Dur.Nanoseconds

	s = fmt.Sprintf("Expected Nanoseconds='%v'. Instead Nanoseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(10346806666132070)
	actualTime = int64(t1Dur.TimeDuration)

	s = fmt.Sprintf("Expected Time Duration='%v'. Instead Time Duration='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	/*
		Results Cumulative Days:
		119-Days 18-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds
		Time Duration Dto
		     StartTimeDateTz:  2018-03-06 20:02:18.792489279 -0600 CST
		       EndTimeDateTz:  2018-07-04 15:09:05.458621349 -0500 CDT
		        TimeDuration:  10346806666132070
		            CalcType:  CumDaysCalc
		               Years:  0
		       YearsNanosecs:  0
		              Months:  0
		      MonthsNanosecs:  0
		               Weeks:  0
		       WeeksNanosecs:  0
		            WeekDays:  0
		    WeekDaysNanosecs:  0
		            DateDays:  119
		    DateDaysNanosecs:  10281600000000000
		               Hours:  18
		       HoursNanosecs:  64800000000000
		             Minutes:  6
		     MinutesNanosecs:  360000000000
		             Seconds:  46
		     SecondsNanosecs:  46000000000
		        Milliseconds:  666
		MillisecondsNanosecs:  666000000
		        Microseconds:  132
		MicrosecondsNanosecs:  132000
		         Nanoseconds:  70
		-----------------------------------------------------
		TotSubSecNanoseconds:  666132070
		  TotDateNanoseconds:  10281600000000000
		  TotTimeNanoseconds:  65206666132070
		-----------------------------------------------------
		Check Total:
		   Date + Time Nanoseconds:  10346806666132070
		Total Duration Nanoseconds:  10346806666132070
		-----------------------------------------------------
	*/

}

func (suite *timedurdtoTestSuite) TestTimeDurationDto_TestCumDays_02() {

	// In this test, t4 is submitted as a Tokyo Time Zone and t5 is submitted as a Cairo
	// Time Zone. However, a standard timezone of US Central is specified. Note,
	// the calculation type is specified as Cumulative Months.
	t1Dur, err := TimeDurationDto{}.NewStartEndTimesTz(suite.t4AsiaTokyo, suite.t5AfricaCairo, TZones.US.Central(), suite.fmtStr)

	assert.Nil(suite.T(), err, "Error NewStartEndTimesTz:")

	err = t1Dur.ReCalcTimeDurationAllocation(TDurCalcType(0).CumDays())

	assert.Nil(suite.T(), err, "Error ReCalcTimeDurationAllocation:")

	expectedTimeDur := suite.t5USCentral.Sub(suite.t4USCentral)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,
		"Expected Time Duration NOT EQUAL To Actual Time Duration!")

	expectedTimeDur = time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, t1Dur.TimeDuration,
		"Expected Time Duration DID NOT EQUAL Date + Time Duration !")

	tx1 := suite.t4USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), 0)

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	timeDaysDur := dur + t1Dur.DateDays*DayNanoSeconds

	expectedEndDate := tx1.Add(time.Duration(timeDaysDur))

	assert.True(suite.T(), expectedEndDate.Equal(t1Dur.EndTimeDateTz.GetDateTimeValue()),
		"Error: Expected Calculated End Date NOT EQUAL to t1Dur.EndDate!")

	assert.True(suite.T(), "CumDays" == t1Dur.CalcType.String(),
		"Error: CalcType String NOT EQUAL to CumDays!")

	expectedTimeDur = suite.t5USCentral.Sub(suite.t4USCentral)

	calculatedTimeDur := time.Duration(t1Dur.TotDateNanoseconds + t1Dur.TotTimeNanoseconds)

	assert.Equal(suite.T(), expectedTimeDur, calculatedTimeDur,
		"Subtracted Time Duration DID NOT EQUAL Date + Time Duration !")

	dur = t1Dur.YearsNanosecs + t1Dur.MonthsNanosecs + t1Dur.DateDaysNanosecs +
		t1Dur.HoursNanosecs + t1Dur.MinutesNanosecs + t1Dur.SecondsNanosecs +
		t1Dur.MillisecondsNanosecs + t1Dur.MicrosecondsNanosecs +
		t1Dur.Nanoseconds

	assert.Equal(suite.T(), expectedTimeDur, time.Duration(dur),
		"Expected Subtracted Duration DID NOT EQUAL Sum of All Component Nanoseconds!")

	expectedOutStr :=
		"119-Days 18-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds"

	actualOutStr, err := t1Dur.GetCumDaysTimeStr()

	assert.Nil(suite.T(), err, "Error t1Dur.GetCumDaysTimeStr():")

	s := fmt.Sprintf("Expected OutStr='%v'. Instead OutStr='%v'", expectedOutStr, actualOutStr)
	assert.Equal(suite.T(), expectedOutStr, actualOutStr, s)

	expectedTime := int64(0)
	actualTime := t1Dur.Years
	s = fmt.Sprintf("Expected Years='%v'. Instead Years='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(0)
	actualTime = t1Dur.Months
	s = fmt.Sprintf("Expected Months='%v'. Instead Months='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(0)
	actualTime = t1Dur.Weeks

	s = fmt.Sprintf("Expected Cumulative Weeks='%v'. Instead Cumulative Weeks='%v'",
		expectedTime, actualTime)

	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(0)
	actualTime = t1Dur.WeekDays

	s = fmt.Sprintf("Expected WeekDays='%v'. Instead WeekDays='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(119)
	actualTime = t1Dur.DateDays

	s = fmt.Sprintf("Expected Days='%v'. Instead Days='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(18)
	actualTime = t1Dur.Hours

	s = fmt.Sprintf("Expected Hours='%v'. Instead Hours='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(6)
	actualTime = t1Dur.Minutes

	s = fmt.Sprintf("Expected Minutes='%v'. Instead Minutes='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(46)
	actualTime = t1Dur.Seconds

	s = fmt.Sprintf("Expected Seconds='%v'. Instead Seconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(666)
	actualTime = t1Dur.Milliseconds

	s = fmt.Sprintf("Expected Milliseconds='%v'. Instead Milliseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(132)
	actualTime = t1Dur.Microseconds

	s = fmt.Sprintf("Expected Microseconds='%v'. Instead Microseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(70)
	actualTime = t1Dur.Nanoseconds

	s = fmt.Sprintf("Expected Nanoseconds='%v'. Instead Nanoseconds='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

	expectedTime = int64(10346806666132070)
	actualTime = int64(t1Dur.TimeDuration)

	s = fmt.Sprintf("Expected Time Duration='%v'. Instead Time Duration='%v'", expectedTime, actualTime)
	assert.Equal(suite.T(), expectedTime, actualTime, s)

}

func TestTimeDuroTestSuite(t *testing.T) {
	tests := new(timedurdtoTestSuite)
	suite.Run(t, tests)
}
