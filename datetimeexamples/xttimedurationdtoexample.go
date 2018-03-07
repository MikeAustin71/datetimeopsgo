package datetimeexamples
import(
	dt "../datetime"
	"fmt"
)

// PrintTimeDurationDto - Prints TimeDurationDto
// data fields.
func PrintTimeDurationDto(tDur dt.TimeDurationDto){

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
	fmt.Println("   Date + Time Nanoseconds: ", tDur.TotTimeNanoseconds + tDur.TotDateNanoseconds)
	fmt.Println("Total Duration Nanoseconds: ", int64(tDur.TimeDuration))
	fmt.Println("-----------------------------------------------------")



}