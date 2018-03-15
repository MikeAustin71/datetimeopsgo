package main

import (
	dt "../datetime"
	ex "../datetimeexamples"
	"time"
	"fmt"
)

func main() {

	mainTest014()

}

func mainTest014() {
	t1str := "2014-02-15 19:54:30.000000000 -0600 CST"
	t2str := "2017-04-30 22:58:32.000000000 -0500 CDT"
	t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)

	//t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t1OutStr := t1.Format(dt.FmtDateTimeYrMDayFmtStr)
	//t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)
	t2OutStr := t2.Format(dt.FmtDateTimeYrMDayFmtStr)
	t12Dur := t2.Sub(t1)

	timeDto := dt.TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}
	timeDto.NormalizeTimeElements()
	dur, err := dt.DurationTriad{}.NewEndTimeMinusTimeDto(t2, timeDto, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDto(t2, timeDto). " +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.DateTime.Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error- Expected Start Time %v. Instead, got %v.\n",
			t1OutStr, dur.BaseTime.StartTimeDateTz.DateTime.Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.DateTime.Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error- Expected End Time %v. Instead, got %v.\n",
			t2OutStr, dur.BaseTime.EndTimeDateTz.DateTime.Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	if t12Dur != dur.BaseTime.TimeDuration {
		fmt.Printf("Error- Expected Time Duration %v. Instead, got %v\n",
			t12Dur, dur.BaseTime.TimeDuration)
		return
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()


	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YrMthDay: %v. Instead, got %v\n", expected, outStr)
		return
	}

	fmt.Println("Successful Completion!")
}

func mainTest012() {
	// 101095442000000000
	t1str := "2014-02-15 19:54:30.000000000 -0600 CST"
	t2str := "2017-04-30 22:58:32.000000000 -0500 CDT"
	t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)
	eDur := t2.Sub(t1)

	timeDto := dt.TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	timeDto.ConvertToNegativeValues()

	dx1 := t2.AddDate(timeDto.Years, timeDto.Months, 0)

	dur := (int64(timeDto.Weeks * 7) + int64(timeDto.WeekDays)) * dt.DayNanoSeconds
	dur += int64(timeDto.Hours) * dt.HourNanoSeconds
	dur += int64(timeDto.Minutes) * dt.MinuteNanoSeconds
	dur += int64(timeDto.Seconds) * dt.SecondNanoseconds

	dx2 := dx1.Add(time.Duration(dur))

	fmt.Println("Expected Start Date Time: ", t1.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("  Actual Start Date Time: ", dx2.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("       Expected Duration: ", int64(eDur))
	fmt.Println("          ActualDuration:  101095442000000000")

}

func mainTest011() {

	t1str := "2014-02-15 19:54:30.000000000 -0600 CST"
	t2str := "2017-04-30 22:58:32.000000000 -0500 CDT"
	t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)

	tDto, err := dt.TimeDurationDto{}.NewStartEndTimesTzCalc(t1, t2, dt.TzIanaUsCentral,
					dt.TDurCalcTypeSTDYEARMTH, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesTzCalc() " +
			"Error='%v'", err.Error())
		return
	}

	ex.PrintTimeDurationDto(tDto)

}

func mainTest010() {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)


	timeDto := dt.TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}
	timeDto.NormalizeTimeElements()

	dur, err := dt.TimeDurationDto{}.NewEndTimeMinusTimeDto(t2, timeDto, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDto(t2, timeDto). Error='%v'\n", err.Error())
		return
	}

	fmt.Println("Expected Start Date Time: ", t1OutStr)
	fmt.Println("  Actual Start Date Time: ", dur.StartTimeDateTz.String())
	fmt.Println("-----------------------------------------")
	fmt.Println("  Expected End Date Time: ", t2OutStr)
	fmt.Println("    Actual End Date Time: ", dur.EndTimeDateTz.String())
	fmt.Println("-----------------------------------------")
	fmt.Println("       Expected Duration: ", t12Dur)
	fmt.Println("         Actual Duration: ", dur.TimeDuration.String())
}

func mainTest009() {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto, err :=
				dt.TimeDurationDto{}.NewStartEndTimesTzCalc(t2, t1,
								dt.TzIanaUsCentral, dt.TDurCalcTypeSTDYEARMTH, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesTzCalc(). " +
			" Error='%v'\n", err.Error())
		return
	}

	fmt.Println("TimeDurationDto")
	ex.PrintTimeDurationDto(tDto)

	durT, err := dt.DurationTriad{}.NewStartEndTimes(t2, t1, dt.TzIanaUsCentral, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesTzCalc(). " +
			" Error='%v'\n", err.Error())
		return
	}

	fmt.Println("DurationTriad BaseTimeDto")
	ex.PrintTimeDurationDto(durT.BaseTime)
}

func mainTest008() {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := dt.DurationTriad{}

	du.SetStartEndTimes(t2, t1, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)

	expected := "0-Hours 0-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := du.BaseTime.GetYearMthDaysTimeAbbrv()

	fmt.Println("Expected: ", expected)
	fmt.Println("  Actual: ", dOut)
	fmt.Println("Start Time: ", du.BaseTime.StartTimeDateTz.String())
	fmt.Println("  End Time: ", du.BaseTime.EndTimeDateTz.String())

	/*
	if expected != dOut {
		fmt.Printf("Expected: %v. \nError - got %v\n", expected, dOut)
		return
	}
	*/

}

func mainTest007() {

	mthTest := int (time.Month(0))

	fmt.Println("===============================")
	fmt.Println("       Month Zero Test")
	fmt.Println("===============================")
	fmt.Println("int (time.Month(0))= ", mthTest)

	/* Result = Sending zero to time.Month(0) yields a
			a zero value. NOT GOOD! Best to use '1' in
			place of zero month number.

				===============================
							 Month Zero Test
				===============================
				int (time.Month(0))=  0

	 */

}

func mainTest006() {

	locUTC, _ := time.LoadLocation(dt.TzIanaUTC)

	fmt.Println()
	fmt.Println("2018-00")
	fmt.Println()

	tDateTime := time.Date(2018, 0, 0 ,0 ,0 ,0 ,0, locUTC)

	ex.PrintDateTime(tDateTime, dt.FmtDateTimeYrMDayFmtStr)
	/*		Result - Don't Use a Zero Month
					2018-00
					----------------------------------
										 Date Time
					----------------------------------
					Date Time:  2017-11-30 00:00:00.000000000 +0000 UCT
					The integer month is:  11
					The integer day is: 30
					The integer year is: 2017
					The integer hour is: 0
					The integer minute is: 0
					The integer second is: 0
					The integer nanosecond is 0
	*/

	fmt.Println()
	fmt.Println("2018-01")
	fmt.Println()

	t2 := time.Date(2018, 1, 0 ,0 ,0 ,0 ,0, locUTC)

	ex.PrintDateTime(t2, dt.FmtDateTimeYrMDayFmtStr)
	/* Result - Best Approach - Use 1 as month number instead of zero month number.
			Also - Use Zero Days. Convert days to duration and add the duration.
			2018-01

			----------------------------------
								 Date Time
			----------------------------------
			Date Time:  2017-12-31 00:00:00.000000000 +0000 UCT
			The integer month is:  12
			The integer day is: 31
			The integer year is: 2017
			The integer hour is: 0
			The integer minute is: 0
			The integer second is: 0
			The integer nanosecond is 0
	*/

	fmt.Println()
	fmt.Println("Add 1 Day")
	fmt.Println()

	dur := int64(24) * dt.HourNanoSeconds

	t3 := t2.Add(time.Duration(dur))

	ex.PrintDateTime(t3, dt.FmtDateTimeYrMDayFmtStr)
	/*
			Add 1 Day to	2017-12-31 00:00:00.000000000 +0000 UCT
			Gives desired result

			----------------------------------
								 Date Time
			----------------------------------
			Date Time:  2018-01-01 00:00:00.000000000 +0000 UCT
			The integer month is:  1
			The integer day is: 1
			The integer year is: 2018
			The integer hour is: 0
			The integer minute is: 0
			The integer second is: 0
			The integer nanosecond is 0
	 */

}

func mainTest005() {

	locUTC, _ := time.LoadLocation(dt.TzIanaUTC)

	tDateTime := time.Date(2018, 2, 0 ,0 ,0 ,0 ,0, locUTC)

	ex.PrintDateTime(tDateTime, dt.FmtDateTimeYrMDayFmtStr)

	fmt.Println()
	fmt.Println("Adding 3-days")
	fmt.Println()

	dur := int64(3) * dt.DayNanoSeconds
	t2 := tDateTime.Add(time.Duration(dur))

	ex.PrintDateTime(t2, dt.FmtDateTimeYrMDayFmtStr)

	expectedDt := time.Date(2018, 2, 3 ,0 ,0 ,0 ,0, locUTC)

	fmt.Println()
	fmt.Println("Complete Date 2018-02-03")
	fmt.Println()

	ex.PrintDateTime(expectedDt, dt.FmtDateTimeYrMDayFmtStr)

}

func mainTest004() {
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	locUSCentral, _ := time.LoadLocation(dt.TzIanaUsCentral)
	locUSPacific, _ := time.LoadLocation(dt.TzIanaUsPacific)
	locParis, _ := time.LoadLocation(dt.TzIanaEuropeParis)
	locCairo, _ := time.LoadLocation(dt.TzIanaAfricaCairo)
	locMoscow, _ := time.LoadLocation(dt.TzIanaEuropeMoscow)
	locTokyo, _ :=	time.LoadLocation(dt.TzIanaAsiaTokyo)

	t1USCentral := time.Date(1948, time.Month(9),7,4,32,16,8185431,locUSCentral)
	t1USPacific := t1USCentral.In(locUSPacific)
	t1EuropeParis := t1USPacific.In(locParis)
	t1AfricaCairo := t1EuropeParis.In(locCairo)
	t1EuropeMoscow := t1AfricaCairo.In(locMoscow)
	t1AsiaTokyo := t1EuropeMoscow.In(locTokyo)
	t1bUSCentral := t1AsiaTokyo.In(locUSCentral)

	fmt.Println("t1USCentral: ", t1USCentral.Format(fmtStr))
	fmt.Println("t1USPacific: ", t1USPacific.Format(fmtStr))
	fmt.Println("t1EuropeParis: ", t1EuropeParis.Format(fmtStr))
	fmt.Println("t1AfricaCairo: ", t1AfricaCairo.Format(fmtStr))
	fmt.Println("t1EuropeMoscow: ", t1EuropeMoscow.Format(fmtStr))
	fmt.Println("t1AsiaTokyo: ", t1AsiaTokyo.Format(fmtStr))
	fmt.Println("t1bUSCentral: ", t1bUSCentral.Format(fmtStr))

}

func mainTest003() {
	loc, _ := time.LoadLocation(dt.TzIanaUsCentral)

	t1 := time.Date(2014, time.Month(15), 67, 19, 54, 30, 158712300, loc)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	addYear  	:= 	0
	addMonth 	:= 	15
	addDay 		:=	64
	addHours	:=  0
	addMinutes := 0
	addSeconds := 0
	addMilliSeconds := 0
	addMicroSeconds := 0
	addNanoSeconds := 0

	var totDuration int64

	t2 := t1.AddDate(addYear,addMonth, addDay)

	totDuration = int64(addHours) * int64(time.Hour)
	totDuration += int64(addMinutes) * int64(time.Minute)
	totDuration += int64(addSeconds) * int64(time.Second)
	totDuration += int64(addMilliSeconds) * int64(time.Millisecond)
	totDuration += int64(addMicroSeconds) * int64(time.Microsecond)
	totDuration += int64(addNanoSeconds)

	t3 := t2.Add(time.Duration(totDuration))

	fmt.Println("t1: ", t1.Format(fmtstr))
	fmt.Println("t2: ", t2.Format(fmtstr))
	fmt.Println("t2: ", t3.Format(fmtstr))

}

func mainTest002() {

	tDto, err := dt.TimeDto{}.New(0, 0, -8, 0, 0, 0, 0, 0, 0, 0 )

	if err != nil {
		fmt.Printf("Error returned from TimeDto{}.New(0, 0, -8, 0, 0, 0, 0, 0, 0, 0 ) Error='%v' \n", err.Error())
	}

	ex.PrintOutTimeDtoFields(tDto)

}

func mainTest001() {

	loc, _ := time.LoadLocation(dt.TzIanaUsCentral)
	t1 := time.Date(2014, time.Month(2), 15, 19, 54, 30, 158712300, loc)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	tDto, err := dt.TimeDto{}.New(2014, 2, 0, 15, 19, 54, 30, 0, 0, 158712300)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDto{}.New(year, month, ...). Error=%v \n", err.Error())
	}

	t2, err := tDto.GetDateTime(dt.TzIanaUsCentral)

	fmt.Println("t1: ", t1.Format(fmtstr))
	fmt.Println("t2: ", t2.Format(fmtstr))

}

