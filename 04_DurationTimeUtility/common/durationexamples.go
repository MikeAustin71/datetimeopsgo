package common

import (
	"errors"
	"fmt"
	"time"
)

// GetBasicDuration - Returns basic duration as a string
func GetBasicDuration() {
	t1str := "04/29/2017 19:54:30 -0500 CDT"
	t2str := "04/29/2017 20:56:32 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	durationUtility := DurationUtility{}
	t1, err := time.Parse(fmtstr, t1str)
	if err != nil {
		panic(errors.New("Time Parse1 Error:" + err.Error()))
	}

	t2, err := time.Parse(fmtstr, t2str)
	if err != nil {
		panic(errors.New("Time Parse2 Error:" + err.Error()))
	}

	duration, err := durationUtility.GetDuration(t1, t2)

	fmt.Println("Duration:", duration)
	// Duration: 1h2m2s
}

// GetElapsedTimeDuration - example of
// GetElapsedTime() in DateTime Utility
func GetElapsedTimeDuration() {
	tstr1 := "04/15/2017 19:54:30.123456489 -0500 CDT"
	tstr2 := "04/18/2017 09:21:16.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, tstr1)
	t2, _ := time.Parse(fmtstr, tstr2)
	durationUtility := DurationUtility{}

	ed, _ := durationUtility.GetElapsedTime(t1, t2)
	fmt.Println("Elapsed Time: ", ed.DurationStr)
	// "2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds"

	fmt.Println("")
	fmt.Println("Default Duration: ", ed.DefaultStr)
	// 61h26m46.864197832s

	fmt.Println("")
	fmt.Println("NanosecStr: ", ed.NanosecStr)
	// 2-Days 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds

}

func GetElapsedYears() {
	t1str := "02/15/2014 19:54:30 -0500 CDT"
	t2str := "04/30/2017 22:58:32 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	durationUtility := DurationUtility{}

	t1, err := time.Parse(fmtstr, t1str)
	if err != nil {
		panic(fmt.Errorf("Time Parse1 Error: %v", err))
	}

	t2, err := time.Parse(fmtstr, t2str)
	if err != nil {
		panic(fmt.Errorf("Time Parse2 Error: %v", err))
	}

	dur, err := durationUtility.GetDuration(t1, t2)
	if err != nil {
		panic(fmt.Errorf("Get Duration Failed: %v", err))
	}

	ed := durationUtility.GetDurationBreakDown(dur.StartDateTime, dur.TimeDuration)
	fmt.Println("Elapsed Time: ", ed.DurationStr)
	// Elapsed Time:  3-Years 74-Days 9-Hours 36-Minutes 26-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds

	fmt.Println("")
	fmt.Println("Default Duration: ", ed.DefaultStr)
	// Default Duration:  28083h4m2s

	fmt.Println("")
	fmt.Println("NanosecStr: ", ed.NanosecStr)
	// NanosecStr:  3-Years 74-Days 9-Hours 36-Minutes 26-Seconds 0-Nanoseconds

}

/*
	For the Gregorian calendar the average length of the calendar year
	(the mean year) across the complete leap cycle of 400 years is 365.2425 days.
	The Gregorian Average Year is therefore equivalent to 365 days, 5 hours,
	49 minutes and 12 seconds
	Sources:
	https://en.wikipedia.org/wiki/Year
	Source: https://en.wikipedia.org/wiki/Gregorian_calendar
*/
func CalcNanosecondsPerYear() {
	yearNow := StdYearNanoSeconds
	yearGregorianSecs := int64(31556952) * SecondNanoseconds
	// Source: https://en.wikipedia.org/wiki/Gregorian_calendar
	//Gregorian 365 days, 5 hours, 49 minutes and 12 seconds
	yearG2 := (DayNanoSeconds * int64(365)) +
		(HourNanoSeconds * int64(5)) +
		(MinuteNanoSeconds * int64(49)) +
		(SecondNanoseconds * int64(12))

	yearGregorian := int64(31556952000011380)

	AvgNanpSecsPerYear := int64(float64(365.2425) * float64(DayNanoSeconds))

	fmt.Println("                            Year Now: ", yearNow)
	fmt.Println("                      Gregorian Year: ", yearGregorian)
	fmt.Println("  Gregorian Avg in hours/minutes/sec: ", yearG2)
	fmt.Println("   Seconds In Average Gregorian Year: ", yearGregorianSecs)
	fmt.Println("Avg Gregorian 365.2425 days per year: ", AvgNanpSecsPerYear)

	/*
																Year Now:  31556952000000000
													Gregorian Year:  31556952000011380
			Gregorian Avg in hours/minutes/sec:  31556952000000000
			 Seconds In Average Gregorian Year:  31556952000000000
		Avg Gregorian 365.2425 days per year:  31556952000000000

	*/

}

func GetTargetTimeFromMinusDuration() {
	tstr1 := "04/15/2017 19:54:30.123456489 +0000 UTC"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	du := DurationUtility{}

	t1, _ := time.Parse(fmtstr, tstr1)

	du.CalcTargetTimeFromMinusDuration(t1, TimesDto{Years: 1})

	tstr2 := du.StartDateTime.Format(fmtstr)

	fmt.Println("tstr1: ", t1.Format(fmtstr))
	fmt.Println("tstr2: ", tstr2)

	d1 := time.Duration((365 * 24)) * time.Hour

	t3 := t1.Add(-d1)

	fmt.Println("   t3: ", t3.Format(fmtstr))
	/*
		tstr1:  04/15/2017 19:54:30.123456489 +0000 UTC
		tstr2:  04/15/2016 19:54:30.123456489 +0000 UTC
			 t3:  04/15/2016 19:54:30.123456489 +0000 UTC
	*/

}

func GetTargetTimeFromMinusDuration2() {
	tstr1 := "04/15/2017 19:54:30.123456489 +0000 UTC"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	du := DurationUtility{}

	t1, _ := time.Parse(fmtstr, tstr1)
	fmt.Println()
	fmt.Println("     Common End Time: ", t1.Format(fmtstr))

	du.CalcTargetTimeFromMinusDuration(t1, TimesDto{Years: 1, Months: 2, Days: 3})

	fmt.Println("Calculated StartTime: ", du.StartDateTime.Format(fmtstr))

	tVerify := t1.AddDate(-1, -2, -3)

	fmt.Println("    Verify StartTime: ", tVerify.Format(fmtstr))

}

func GetTargetTimeFromPlusDuration() {
	tstr1 := "04/15/2017 19:54:30.123456489 +0000 UTC"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	du := DurationUtility{}

	t1, _ := time.Parse(fmtstr, tstr1)
	fmt.Println()
	fmt.Println(" Common Start Time: ", t1.Format(fmtstr))

	du.CalcTargetTimeFromPlusDuration(t1, TimesDto{Years: 1, Months: 2, Days: 3})

	fmt.Println("Calculated EndTime: ", du.EndDateTime.Format(fmtstr))

	tVerify := t1.AddDate(1, 2, 3)

	fmt.Println("    Verify EndTime: ", tVerify.Format(fmtstr))

}

func ExampleElapsedYearsBreakdown() error {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	durationUtility := DurationUtility{}
	t1, err := time.Parse(fmtstr, t1str)
	if err != nil {
		return fmt.Errorf("Time Parse1 Error:", err.Error())
	}

	t2, err := time.Parse(fmtstr, t2str)
	if err != nil {
		return fmt.Errorf("Time Parse2 Error:", err.Error())
	}

	dur, err := durationUtility.GetDuration(t1, t2)
	if err != nil {
		return fmt.Errorf("Get Duration Failed: ", err.Error())
	}

	//du := durationUtility.GetDurationBreakDown(dur.StartDateTime, dur.TimeDuration)

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	fmt.Println("     Expected Dur: ", expected)
	fmt.Println("     Duration Str: ", dur.DurationStr)
	fmt.Println("      Default Str: ", dur.DefaultStr)
	fmt.Println("         Days Str: ", dur.DaysStr)
	fmt.Println("        Hours Str: ", dur.HoursStr)
	fmt.Println("YearMthsWeeks Str: ", dur.YearsMthsWeeksStr)
	fmt.Println("    Cum Weeks Str: ", dur.CumWeeksStr)
	/* Output
		   Expected Dur:  3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds
	     Duration Str:  3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds
	      Default Str:  28082h4m2s
	         Days Str:  1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds
	        Hours Str:  28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds
	YearMthsWeeks Str:  3-Years 2-Months 2-Weeks 1-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds
	    Cum Weeks Str:  167-Weeks 1-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds

	*/

	return nil
}

func ExampleCalcDurationElements() {
	t1Str := "04/30/2017 22:58:32.000000000 -0500 CDT"

	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1Str)
	t2 := t1.AddDate(4, 0, 2)

	td := t2.Sub(t1)

	du := DurationUtility{StartDateTime: t1, TimeDuration: td}
	du.CalcDurationElements()
	du.CalcDurationStrings()
	expected := "4-Years 0-Months 2-Days 0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	fmt.Println("   du.DurationStr: ", du.DurationStr)
	fmt.Println("         Expected: ", expected)
	fmt.Println()
	fmt.Println(" du.StartDateTime: ", du.StartDateTime.Format(fmtstr))
	fmt.Println("Verify Start Time: ", t1.Format(fmtstr))
	fmt.Println()
	fmt.Println("   du.EndDateTime: ", du.EndDateTime.Format(fmtstr))
	fmt.Println(" Verfity End Time: ", t2.Format(fmtstr))

	// Output
	// "4-Years 0-Months 2-Days 0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

}

func ExampleCalcMthDurationElements() {
	t1Str := "04/30/2017 22:58:32.000000000 -0500 CDT"

	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1Str)
	t2 := t1.AddDate(4, 3, 2)

	td := t2.Sub(t1)

	du := DurationUtility{StartDateTime: t1, TimeDuration: td}
	du.CalcDurationElements()
	du.CalcDurationStrings()
	expected := "4-Years 3-Months 2-Days 0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	fmt.Println("   du.DurationStr: ", du.DurationStr)
	fmt.Println("         Expected: ", expected)
	fmt.Println()
	fmt.Println(" du.StartDateTime: ", du.StartDateTime.Format(fmtstr))
	fmt.Println("Verify Start Time: ", t1.Format(fmtstr))
	fmt.Println()
	fmt.Println("   du.EndDateTime: ", du.EndDateTime.Format(fmtstr))
	fmt.Println(" Verfity End Time: ", t2.Format(fmtstr))
}

func ExampleElapsedYearsBreakdown2() {
	t1str := "02/15/2014 19:54:30.000000000 -0500 CDT"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected1 := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	durationUtility := DurationUtility{}
	t1, err := time.Parse(fmtstr, t1str)
	if err != nil {
		panic(fmt.Errorf("Time Parse1 Error:", err.Error()))
	}

	t2, err := time.Parse(fmtstr, t2str)
	if err != nil {
		panic(fmt.Errorf("Time Parse2 Error:", err.Error()))
	}

	dur, err := durationUtility.GetDuration(t1, t2)

	if err != nil {
		panic(fmt.Errorf("Get Duration Failed: ", err.Error()))
	}

	du := durationUtility.GetDurationBreakDown(dur.StartDateTime, dur.TimeDuration)

	if du.DurationStr != expected1 {
		panic(fmt.Errorf("Expected: %v, got:", expected1, du.DurationStr))
	}

	t3 := t1.AddDate(3, 2, 15).Add(time.Duration(int64((3 * HourNanoSeconds) + (4 * MinuteNanoSeconds) + (2 * SecondNanoseconds))))

	fmt.Println("Duration Start Date Time: ", du.StartDateTime.Format(fmtstr))
	fmt.Println("  Duration End Date Time: ", du.EndDateTime.Format(fmtstr))
	fmt.Println("      t3 Verify End Time: ", t3.Format(fmtstr))
	fmt.Println("          du.DurationStr: ", du.DurationStr)
	fmt.Println("    Expected DurationStr: ", expected1)
	fmt.Println("           Time Duration: ", du.TimeDuration)
	fmt.Println("     Time Duration int64: ", int64(du.TimeDuration))

	/* Output
		Duration Start Date Time:  02/15/2014 19:54:30.000000000 -0500 CDT
	    Duration End Date Time:  04/30/2017 22:58:32.000000000 -0500 CDT
	        t3 Verify End Time:  04/30/2017 22:58:32.000000000 -0500 CDT
	            du.DurationStr:  3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds
	      Expected DurationStr:  3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds
	             Time Duration:  28083h4m2s
	       Time Duration int64:  101099042000000000
	*/

}
