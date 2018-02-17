package main

import (
	// "MikeAustin71/datetimeopsgo/TimeZoneDto/common"

	"time"
	"fmt"
	"MikeAustin71/datetimeopsgo/TimeZoneDto/common"
)

/*
import (
	"MikeAustin71/datetimeopsgo/TimeZoneDto/common"
	"fmt"
	"time"
)

*/

func main() {

	// fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	//neutralFmtStr := "2006-01-02 15:04:05.000000000"
	usPacificLoc, _ :=time.LoadLocation(common.TzIanaUsPacific)

	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)

	tzDef1, err := common.TimeZoneDefDto{}.New(tUsPacific)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefDto{}.New(tUsPacific). Error='%v'", err.Error())
		return
	}

	tzDef2 := tzDef1.CopyOut()


	tzDef3 := tzDef2.CopyOut()

	tzDef3.Location = nil

	fmt.Printf("First tzDef1")
	fmt.Printf("        Location Name: %v\n", tzDef1.LocationName)
	fmt.Printf("Pointer Location Name: %x\n", tzDef1.LocationName)
	fmt.Printf("     Location Pointer: %x\n", tzDef1.Location)


	fmt.Printf("Second tzDef2")
	fmt.Printf("        Location Name: %v\n", tzDef2.LocationName)
	fmt.Printf("Pointer Location Name: %x\n", tzDef2.LocationName)
	fmt.Printf("     Location Pointer: %x\n", tzDef2.Location)


	fmt.Printf("third tzDef3")
	fmt.Printf("        Location Name: %v\n", tzDef3.LocationName)
	fmt.Printf("Pointer Location Name: %x\n", tzDef3.LocationName)
	fmt.Printf("     Location Pointer: %x\n", tzDef3.Location)
	fmt.Printf("    LocationPtrString: %v\n", tzDef3.Location.String())

}
