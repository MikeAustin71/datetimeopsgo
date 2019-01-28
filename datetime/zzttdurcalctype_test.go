package datetime

import (
	"testing"
)

func TestTDurCalcTypeString_001(t *testing.T) {

	r := TDurCalcType(0).StdYearMth()

	s := r.String()

	if "StdYearMthCalc" != s {
		t.Errorf("Expected TDurCalcType(0).StdYearMth() string='%v'. Instead, string='%v' ", "StdYearMthCalc", s)
	}

}

func TestTDurCalcTypeString_002(t *testing.T) {

	r := TDurCalcType(0).CumMonths()
	expectedStr := "CumMonthsCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumWeeks() string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeString_003(t *testing.T) {

	r := TDurCalcType(0).CumWeeks()
	expectedStr := "CumWeeksCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumWeeks() string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeString_004(t *testing.T) {

	r := TDurCalcType(0).CumDays()
	expectedStr := "CumDaysCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumDays() string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeString_005(t *testing.T) {

	r := TDurCalcType(0).CumHours()
	expectedStr := "CumHoursCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).CumHours() string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeString_006(t *testing.T) {

	r := TDurCalcType(0).CumMinutes()
	expectedStr := "CumMinutesCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).GregorianYears() string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeString_007(t *testing.T) {

	r := TDurCalcType(0).CumSeconds()
	expectedStr := "CumSecondsCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).GregorianYears() string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeString_008(t *testing.T) {

	r := TDurCalcType(0).GregorianYears()
	expectedStr := "GregorianYrsCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcType(0).GregorianYears() string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeValue_001(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).StdYearMth()

	i = int(r)

	if i != 0 {
		t.Errorf("Expected 'TDurCalcType(0).StdYearMth()' value = 0. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_002(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumMonths()

	i = int(r)

	if i != 1 {
		t.Errorf("Expected 'TDurCalcType(0).CumMonths()' value = 1. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_003(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumWeeks()

	i = int(r)

	if i != 2 {
		t.Errorf("Expected 'TDurCalcType(0).CumWeeks()' value = 2. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_004(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumDays()

	i = int(r)

	if i != 3 {
		t.Errorf("Expected 'TDurCalcType(0).CumDays()' value = 3. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_005(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumHours()

	i = int(r)

	if i != 4 {
		t.Errorf("Expected 'TDurCalcType(0).CumHours()' value = 4. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_006(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumMinutes()

	i = int(r)

	if i != 5 {
		t.Errorf("Expected 'TDurCalcType(0).CumMinutes()' value = 5. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_007(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).CumSeconds()

	i = int(r)

	if i != 6 {
		t.Errorf("Expected 'TDurCalcType(0).CumMinutes()' value = 6. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_008(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcType(0).GregorianYears()

	i = int(r)

	if i != 7 {
		t.Errorf("Expected 'TDurCalcType(0).GregorianYears()' value = 7. Instead, got %v", i)
	}

}
