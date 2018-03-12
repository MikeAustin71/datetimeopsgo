package datetime

import (
	"testing"
)

func TestTimeDto_NormalizeTimeElements_01(t *testing.T) {

	t1Dto := TimeDto{}

	t1Dto.Years = 1955
	t1Dto.Months = 15
	t1Dto.DateDays = 32
	t1Dto.Hours = 48
	t1Dto.Minutes = 71
	t1Dto.Seconds = 125
	t1Dto.Milliseconds = 1001
	t1Dto.Microseconds = 1001
	t1Dto.Nanoseconds = 1001

	err := t1Dto.NormalizeTimeElements()

	if err != nil {
		t.Errorf("Error returned by t1Dto.NormalizeTimeElements(). Error='%v'", err.Error())
	}

}
