package main

import (
	dt "../datetime"
	ex "../datetimeexamples"
	"time"
	"fmt"
)

func main() {

	t0str := "2017-04-30 22:58:32.515539300 -0500 CDT"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t0, err := time.Parse(fmtstr, t0str)

	if err != nil {
		fmt.Printf("Error retruned from time.Parse(fmtstr, t0str). t0str='%v'  Error='%v'\n", t0str, err.Error())
		return
	}

	tDto, err := dt.TimeDto{}.New(2017, 04, 0, 30, 22, 58,32,0,0, 515539300)

	if err != nil {
		fmt.Printf("Error returned from TimeDto{}.New(...)  Error='%v'\n", err.Error())
		return
	}

	fmt.Println("Original t0str: ", t0str)
	fmt.Println("Original t0: ", t0.Format(fmtstr))
	ex.PrintOutTimeDtoFields(tDto)
}
