package datetimeexamples

import (
	dt "../datetime"
	"fmt"
	"time"
)

// ExampleTimeDuration001
func ExampleTimeDuration001() {
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	locUSCentral, _ := time.LoadLocation(dt.IanaTz.US.Central())

	t1USCentral := time.Date(2018, time.Month(3), 10, 18, 0, 0, 0, locUSCentral)
	//t1USCentral := time.Date(2018, time.Month(4),1,20,02,18,792489279, locUSCentral)

	//t2AddDate := t1USCentral.AddDate(0, 0, 1)

	hoursDur := int64(24) * dt.HourNanoSeconds

	t1Dur, err := dt.TimeDurationDto{}.NewStartTimeDurationCalcTz(t1USCentral, time.Duration(hoursDur),
		dt.TDurCalcType(0).StdYearMth(), dt.IanaTz.US.Central(), fmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.New(t1USCentral, t2USCentral, fmtStr). Error='%v'\n", err.Error())
	}

	fmt.Println("Add Date Results - Cumulative Days")
	fmt.Println("            Start Date Time: ", t1USCentral.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("      -- Duration = 24-Hours --")
	fmt.Println("       Actual End Date Time: ", t1Dur.EndTimeDateTz.DateTime.Format(dt.FmtDateTimeYrMDayFmtStr))
	//fmt.Println("             Add Date 1 Day: ", t2AddDate.Format(dt.FmtDateTimeYrMDayFmtStr))

}

// ExampleTimeDuration002
func ExampleTimeDuration002() {

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	locUSCentral, _ := time.LoadLocation(dt.IanaTz.US.Central())

	t1USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)
	//t1USCentral := time.Date(2018, time.Month(4),1,20,02,18,792489279, locUSCentral)

	t2USCentral := time.Date(2018, time.Month(7), 04, 15, 9, 5, 458621349, locUSCentral)
	//t2USCentral := time.Date(2018, time.Month(4),15,20,02,18,792489279, locUSCentral)

	t1Dur, err := dt.TimeDurationDto{}.NewStartEndTimesCalcTz(t1USCentral, t2USCentral,
		dt.TDurCalcType(0).CumDays(), dt.IanaTz.US.Central(), fmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.New(t1USCentral, t2USCentral, fmtStr). Error='%v'\n", err.Error())
	}

	tx1 := t1USCentral.AddDate(int(t1Dur.Years), int(t1Dur.Months), int(t1Dur.DateDays))

	dur := t1Dur.Hours * int64(time.Hour)
	dur += t1Dur.Minutes * int64(time.Minute)
	dur += t1Dur.Seconds * int64(time.Second)
	dur += t1Dur.Milliseconds * int64(time.Millisecond)
	dur += t1Dur.Microseconds * int64(time.Microsecond)
	dur += t1Dur.Nanoseconds

	tx2 := tx1.Add(time.Duration(dur))

	fmt.Println("Add Date Results - Cumulative Days")
	fmt.Println("            Start Date Time: ", t1USCentral.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("     Expected End Date Time: ", tx2.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("       Actual End Date Time: ", t1Dur.EndTimeDateTz.DateTime.Format(dt.FmtDateTimeYrMDayFmtStr))

	totDur := dur
	totDur += t1Dur.DateDays * int64(time.Hour) * int64(24)

	tx3 := t1USCentral.Add(time.Duration(totDur))
	fmt.Println("Acutal End Date by Duration: ", tx3.Format(dt.FmtDateTimeYrMDayFmtStr))

}

// ExampleTimeDuration003
func ExampleTimeDuration003() {
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	locUSCentral, _ := time.LoadLocation(dt.IanaTz.US.Central())

	t1USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)
	//t1USCentral := time.Date(2018, time.Month(4),1,20,02,18,792489279, locUSCentral)

	t2USCentral := time.Date(2018, time.Month(7), 04, 15, 9, 5, 458621349, locUSCentral)
	//t2USCentral := time.Date(2018, time.Month(4),15,20,02,18,792489279, locUSCentral)

	tDur, err := dt.TimeDurationDto{}.NewStartEndTimesCalcTz(t1USCentral, t2USCentral,
		dt.TDurCalcType(0).CumDays(), dt.IanaTz.US.Central(), fmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.New(t1USCentral, t2USCentral, fmtStr). Error='%v'\n", err.Error())
	}

	fmt.Println("Results Cumulative Days:")
	//fmt.Println(tDur.GetCumDaysTimeStr())
	fmt.Println(tDur.GetCumDaysTimeStr())

	fmt.Println("Time Duration Dto")
	PrintTimeDurationDto(tDur)

}

// PrintTimeDurationDto - Prints TimeDurationDto
// data fields.
func PrintTimeDurationDto(tDur dt.TimeDurationDto) {

	fmt.Println("     StartTimeDateTz: ", tDur.StartTimeDateTz.String())
	fmt.Println("       EndTimeDateTz: ", tDur.EndTimeDateTz.String())
	fmt.Println("        TimeDuration: ", int64(tDur.TimeDuration))
	fmt.Println("            CalcType: ", tDur.CalcType.String())
	fmt.Println("               Years: ", tDur.Years)
	fmt.Println("       YearsNanosecs: ", tDur.YearsNanosecs)
	fmt.Println("              Months: ", tDur.Months)
	fmt.Println("      MonthsNanosecs: ", tDur.MonthsNanosecs)
	fmt.Println("               Weeks: ", tDur.Weeks)
	fmt.Println("       WeeksNanosecs: ", tDur.WeeksNanosecs)
	fmt.Println("            WeekDays: ", tDur.WeekDays)
	fmt.Println("    WeekDaysNanosecs: ", tDur.WeekDaysNanosecs)
	fmt.Println("            DateDays: ", tDur.DateDays)
	fmt.Println("    DateDaysNanosecs: ", tDur.DateDaysNanosecs)
	fmt.Println("               Hours: ", tDur.Hours)
	fmt.Println("       HoursNanosecs: ", tDur.HoursNanosecs)
	fmt.Println("             Minutes: ", tDur.Minutes)
	fmt.Println("     MinutesNanosecs: ", tDur.MinutesNanosecs)
	fmt.Println("             Seconds: ", tDur.Seconds)
	fmt.Println("     SecondsNanosecs: ", tDur.SecondsNanosecs)
	fmt.Println("        Milliseconds: ", tDur.Milliseconds)
	fmt.Println("MillisecondsNanosecs: ", tDur.MillisecondsNanosecs)
	fmt.Println("        Microseconds: ", tDur.Microseconds)
	fmt.Println("MicrosecondsNanosecs: ", tDur.MicrosecondsNanosecs)
	fmt.Println("         Nanoseconds: ", tDur.Nanoseconds)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("TotSubSecNanoseconds: ", tDur.TotSubSecNanoseconds)
	fmt.Println("  TotDateNanoseconds: ", tDur.TotDateNanoseconds)
	fmt.Println("  TotTimeNanoseconds: ", tDur.TotTimeNanoseconds)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Check Total:")
	fmt.Println("   Date + Time Nanoseconds: ", tDur.TotTimeNanoseconds+tDur.TotDateNanoseconds)
	fmt.Println("Total Duration Nanoseconds: ", int64(tDur.TimeDuration))
	fmt.Println("-----------------------------------------------------")

}
