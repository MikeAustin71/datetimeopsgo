package main

import (
	// "MikeAustin71/datetimeopsgo/TimeZoneUtility/common"

	"time"
	"fmt"
	"MikeAustin71/datetimeopsgo/TimeZoneUtility/common"
)

/*
import (
	"MikeAustin71/datetimeopsgo/TimeZoneUtility/common"
	"fmt"
	"time"
)

*/

func main() {

	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)


	dtz2, err := common.DateTzDto{}.New(t1)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.New(t1). Error='%v'", err.Error())
		return
	}


	fmt.Println("       dtz2 DateTime: ", dtz2.DateTime.Format(fmtstr))
	fmt.Println("t1 Date Time Out Str: ", t1OutStr)

	fmt.Println()

}
