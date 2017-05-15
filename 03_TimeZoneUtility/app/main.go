package main

import (
	"fmt"
	"time"
)

/*
import (
	common "MikeAustin71/datetimeopsgo/03_TimeZoneUtility/common"
	"fmt"
	"time"
)

*/

func main() {
	now := time.Now()
	x, _ := time.LoadLocation(now.Location().String())
	y := now.Local().Location()

	fmt.Println("x", x)
	fmt.Println("y", y)

}
