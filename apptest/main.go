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
