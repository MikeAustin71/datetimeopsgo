package common

import (
	"testing"
	"time"
)


func TestDateTzDto_NewDateTimeElements_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.123000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	dTzDto, err := DateTzDto{}.NewDateTimeElements(2014, 2,15,19,54,30,123, TzUsPacific)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(2014, 2,15,19,54,30,123, TzUsPacific). Error='%v'", err.Error())
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'",t1OutStr, actualTimeStr)
	}

}