package datetimeexamples

import (
	"fmt"
	dt "github.com/MikeAustin71/datetimeopsgo/datetime"
	"time"
)

// ExampleSetStartEndTimes
func ExampleSetStartEndTimes() {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	dur := dt.DurationTriad{}

	err := dur.SetStartEndTimes(
		t1,
		t2,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.US.Central(),
		dt.TCalcMode.LocalTimeZone(),
		fmtstr)

	if err != nil {
		panic(fmt.Errorf("ExampleSetStartEndTimes() Error returned by "+
			"dur.SetStartEndTimes(). Error='%v' ", err.Error()))
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	fmt.Println("        Expected: ", expected)
	fmt.Println(" YrsMthsDaysTime: ", outStr)

	expected = "28082h4m2s"
	fmt.Println()
	outStr = dur.BaseTime.GetDefaultDurationStr()

	fmt.Println("Expected Default: ", expected)
	fmt.Println("Default Duration: ", outStr)

	fmt.Println()

	expected = "1170-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	outStr, err = dur.BaseTime.GetCumDaysTimeStr()

	if err != nil {
		fmt.Printf("Error returned by dur.BaseTime.GetCumDaysTimeStr(). "+
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Println("     Expected WeekDays: ", expected)
	fmt.Println("     WeekDays Duration: ", outStr)

	fmt.Println()
	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	outStr, err = dur.BaseTime.GetCumHoursTimeStr()

	if err != nil {
		fmt.Printf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Println("    Expected Hours: ", expected)
	fmt.Println("    Hours Duration: ", outStr)

	fmt.Println()
	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	outStr = dur.BaseTime.GetYrMthWkDayHrMinSecNanosecsStr()

	fmt.Println("Expected YrMthWeekDayTime : ", expected)
	fmt.Println("          YrMthWeekDayTime: ", outStr)

	fmt.Println()

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	outStr, err = dur.BaseTime.GetCumWeeksDaysTimeStr()

	if err != nil {
		fmt.Printf("Error returned by dur.BaseTime.GetCumWeeksDaysTimeStr(). "+
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Println("Expected Weeks WeekDays Time: ", expected)
	fmt.Println("     Weeks WeekDays Duration: ", outStr)

}

// Example_NewStartTimeDuration_01
func ExampleNewStartTimeDuration01() {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t1Utc := t1.UTC()
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t2Utc := t2.UTC()
	t12Dur := t2.Sub(t1)
	t12UTCDur := t2Utc.Sub(t1Utc)

	dur, err := dt.DurationTriad{}.NewStartTimeDuration(
		t1,
		t12Dur,
		dt.TDurCalc.StdYearMth(),
		dt.TZones.US.Central(),
		dt.TCalcMode.LocalTimeZone(),
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewStartTimeTzDuration(t1, t12Dur). Error='%v'\n", err.Error())
		return
	}

	// dur.SetStartTimeDurationTz(t1, t12Dur)

	if t1OutStr != dur.BaseTime.GetThisStartDateTime().Format(fmtstr) {
		fmt.Printf("Error- Expected Start Time %v.\n " +
			"Instead, got %v.\n", t1OutStr,
			dur.BaseTime.GetThisStartDateTime().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.GetThisEndDateTime().Format(fmtstr) {
		fmt.Printf("Error- Expected End Time %v.\n" +
			"Instead, got %v.\n",
			t2OutStr, dur.BaseTime.GetThisEndDateTime().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.GetThisTimeDuration() {
		fmt.Printf("Error- Expected Time Duration %v.\n" +
			"Instead, got %v\n", t12Dur, dur.BaseTime.GetThisTimeDuration())
	}

	if t12Dur != t12UTCDur {
		fmt.Printf("Time In Duration different from Time UTC Duration.\n" +
			"t12Dur='%v'  t12UTCDur='%v'", t12Dur, t12UTCDur)
	}

	fmt.Println("Time  In Duration: ", t12Dur)
	fmt.Println("Time UTC Duration: ", t12UTCDur)

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	tDatePlus := t1.AddDate(3, 2, 15)
	tDatePlusDur := tDatePlus.Sub(t1)
	ns := int64(0)

	ns += 3 * dt.HourNanoSeconds
	ns += 4 * dt.MinuteNanoSeconds
	ns += 2 * dt.SecondNanoseconds

	t3 := tDatePlus.Add(time.Duration(ns))
	t3Output := t3.Format(fmtstr)
	fmt.Println("Expected Output Date: ", t2OutStr)
	fmt.Println("Computed Output Date: ", t3Output)

	outStr = dur.BaseTime.GetYearMthDaysTimeStr()

	ans := dur.BaseTime.YearsNanosecs

	fmt.Println("  Actual Counted Year NanoSeconds: ", ans)
	t3Years := t1.AddDate(3, 0, 0)
	t3YearsDur := t3Years.Sub(t1)
	fmt.Println("        Computed Year NanoSeconds: ", int64(t3YearsDur))
	ans += dur.BaseTime.MonthsNanosecs
	ans += dur.BaseTime.DateDaysNanosecs
	fmt.Println()
	fmt.Println("Actual Counted YearMonthDay Nanoseconds: ", ans)
	fmt.Println("      Computed YearMonthDay NanoSeconds: ", int64(tDatePlusDur))
	fmt.Println("---------------------------------------------")
	t3YearsMonths := t3Years.AddDate(0, 2, 0)
	t3MonthsDur := t3YearsMonths.Sub(t3Years)
	fmt.Println("  Actual Counted Month NanoSeconds: ", dur.BaseTime.MonthsNanosecs)
	fmt.Println("        Computed Month NanoSeconds: ", int64(t3MonthsDur))
	fmt.Println("           Info 1-Hour NanoSeconds: ", dt.HourNanoSeconds)
}
