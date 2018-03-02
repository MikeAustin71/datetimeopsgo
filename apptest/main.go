package main

import (
	dt "../datetime"
	ex "../datetimeexamples"
	"time"
	"fmt"

)

func main() {

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

func mainTest004() {
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

func mainTest003() {

	tDto, err := dt.TimeDto{}.New(0, 0, -8, 0, 0, 0, 0, 0, 0, 0 )

	if err != nil {
		fmt.Printf("Error returned from TimeDto{}.New(0, 0, -8, 0, 0, 0, 0, 0, 0, 0 ) Error='%v' \n", err.Error())
	}

	ex.PrintOutTimeDtoFields(tDto)

}

func mainTest002() {

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

func mainTest001() {
	t1str := "02/15/2014 19:54:30.158712300 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	dTzDto, err := dt.DateTzDto{}.NewDateTimeElements(2014, 2,15,19,54,30,158712300, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewDateTimeElements(2014, 2,15,19,54,30,158712300, TzUsCentral). Error='%v'\n", err.Error())
	}

	fmt.Println("t1OutStr: ", t1OutStr)
	ex.PrintOutDateTzDtoFields(dTzDto)

}
