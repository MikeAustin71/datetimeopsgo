package main

import (
	// "MikeAustin71/datetimeopsgo/TimeZoneUtility/common"

	"time"
	"fmt"
)

/*
import (
	"MikeAustin71/datetimeopsgo/TimeZoneUtility/common"
	"fmt"
	"time"
)

*/

func main() {

	millisecond := int64(time.Millisecond)
	microsecond := int64(time.Microsecond)
	fmt.Println("     nanoseconds in second: ", int64(time.Second))
	fmt.Println("nanoseconds in millisecond: ", millisecond)
	fmt.Println("nanoseconds in microsecond: ", microsecond)

	fmt.Println()

}
