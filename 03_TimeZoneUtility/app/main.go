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
	nowLocal := time.Now()
	Cst, _ := time.LoadLocation("America/Chicago")
	nowInCentralTime := nowLocal.In(Cst)

	if nowInCentralTime.Equal(nowLocal) {
		fmt.Println("Times are Equivalent")
	} else {
		fmt.Println("Times are NOT Equivalent")
	}



}
