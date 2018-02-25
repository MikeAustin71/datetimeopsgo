package main

import (
	dt "../datetime"
	"time"
	"fmt"
)

func main() {

	//t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	//t1, _ := time.Parse(fmtstr, t1str)
	t1 := time.Now().Local()

	expectedOutDate := t1.Format(fmtstr)

	dtz1, err := dt.DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		fmt.Printf("Error Returned by dt.DateTzDto{}.New(t1, fmtstr). Error='%v'\n", err.Error())
	}

	fmt.Println("expectedOutDate: ", expectedOutDate)
	fmt.Println("  dtz1.DateTime: ", dtz1)
	dt.PrintOutDateTzDtoFields(dtz1)
}
