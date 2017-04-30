package main

import (
	"MikeAustin71/logopsgo/common"
	"fmt"
	"time"
)

func main() {

	// Get current time in the form of a string
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"

	t, err := time.Parse(fmtstr, tstr)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := common.GetDateTimeStr(t)

	fmt.Println(tstr, "=", t)
	fmt.Println("result=", result)

}
