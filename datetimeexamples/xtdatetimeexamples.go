package datetimeexamples


import (
	 dt "../datetime"
	"fmt"
	"time"
)

// ExamplesDateTimeTimeStampEverything
func ExamplesDateTimeTimeStampEverything(){
	du := dt.DtMgr{}

	fmt.Println("------------------------")
	fmt.Println("  Testing Time Stamps")
	fmt.Println("------------------------")
	fmt.Println("   Time Stamp Everything Fmt: ", du.GetTimeStampEverything())
	fmt.Println("Time Stamp AbbrvNano Sec Fmt: ", du.GetTimeStampYMDAbbrvDowNano())
}

// ExampleDateTimeGetCurrentTimeAsInts - Breaks down time
// to constituent elements as integers.
func ExampleDateTimeGetCurrentTimeAsInts() {
	// Get current time
	t := time.Now().Local()
	var i int64
	i = int64(t.Month())
	fmt.Println("The integer month is: ", i)
	i = int64(t.Day())
	fmt.Println("The integer day is:", i)
	i = int64(t.Year())
	fmt.Println("The integer year is:", i)
	i = int64(t.Hour())
	fmt.Println("The integer hour is:", i)
	i = int64(t.Minute())
	fmt.Println("The integer minute is:", i)
	i = int64(t.Second())
	fmt.Println("The integer second is:", i)
	i = int64(t.Nanosecond())
	fmt.Println("The integer nanosecond is", i)
}

// PrintDateTime
func PrintDateTime(dt time.Time, fmtStr string) {

	var i int64
	fmt.Println("----------------------------------")
	fmt.Println("           Date Time")
	fmt.Println("----------------------------------")
	fmt.Println("Date Time: ", dt.Format(fmtStr))
	i = int64(dt.Month())
	fmt.Println("The integer month is: ", i)
	i = int64(dt.Day())
	fmt.Println("The integer day is:", i)
	i = int64(dt.Year())
	fmt.Println("The integer year is:", i)
	i = int64(dt.Hour())
	fmt.Println("The integer hour is:", i)
	i = int64(dt.Minute())
	fmt.Println("The integer minute is:", i)
	i = int64(dt.Second())
	fmt.Println("The integer second is:", i)
	i = int64(dt.Nanosecond())
	fmt.Println("The integer nanosecond is", i)

}

// ExamplesDateTimeGetEverythingFormat - provides a sample of the
// 'GetEverything Date Time Format!
func ExamplesDateTimeGetEverythingFormat() {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	testTime, _ := time.Parse(fmtstr, tstr)
	dateT := dt.DtMgr{}
	str := dateT.GetDateTimeEverything(testTime)
	fmt.Println("Everything Format: ", str)
	// Saturday April 29, 2017 19:54:30.123456489 -0500 CDT

	testTime2, _ := time.Parse(dt.FmtDateTimeTzNano, tstr)

	str2 := dateT.GetDateTimeEverything(testTime2)

	fmt.Println("Time Zone Second Format: ", str2)

}

// ExamplesDateTimeGetCurrentTimeAsString - Get current time in the form of a string
func ExamplesDateTimeGetCurrentTimeAsString() {
	tstr := "04/29/2017 19:54:30 -0500 CDT"

	dateT	:= dt.DtMgr{}
	t, err := time.Parse(dt.FmtDateTimeTzNano, tstr)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := dateT.GetDateTimeStr(t)

	fmt.Println(tstr, "=", t)
	fmt.Println("result=", result)
}
