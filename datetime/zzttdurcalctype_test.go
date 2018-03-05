package datetime

import (
	"testing"
)


func TestTDurCalcTypeString_001(t *testing.T) {

r := TDurCalcTypeSTDYEARMTH

	s := r.String()

	if "StdYearMthCalc" != s {
		t.Errorf("Expected TDurCalcTypeSTDYEARMTH string='%v'. Instead, string='%v' ", "StdYearMthCalc", s)
	}

}

func TestTDurCalcTypeString_002(t *testing.T) {

	r := TDurCalcTypeCUMWEEKS
	expectedStr := "CumMonthsCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcTypeCUMWEEKS string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}


func TestTDurCalcTypeString_003(t *testing.T) {

	r := TDurCalcTypeCUMWEEKS
	expectedStr := "CumWeeksCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcTypeCUMWEEKS string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeString_004(t *testing.T) {

	r := TDurCalcTypeCUMDAYS
	expectedStr := "CumDaysCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcTypeCUMDAYS string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeString_005(t *testing.T) {

	r := TDurCalcTypeCUMHOURS
	expectedStr := "CumHoursCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcTypeCUMHOURS string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeString_006(t *testing.T) {

	r := TDurCalcTypeCUMMINTUES
	expectedStr := "CumMinutesCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcTypeGregorianYrs string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeString_007(t *testing.T) {

	r := TDurCalcTypeGregorianYrs
	expectedStr := "GregorianYrsCalc"

	s := r.String()

	if expectedStr != s {
		t.Errorf("Expected TDurCalcTypeGregorianYrs string='%v'. Instead, string='%v' ", expectedStr, s)
	}

}

func TestTDurCalcTypeValue_001(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcTypeSTDYEARMTH

	i = int(r)

	if i != 0 {
		t.Errorf("Expected 'TDurCalcTypeSTDYEARMTH' value = 0. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_002(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcTypeCUMMONTHS

	i = int(r)

	if i != 1 {
		t.Errorf("Expected 'TDurCalcTypeCUMMONTHS' value = 1. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_003(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcTypeCUMWEEKS

	i = int(r)

	if i != 2 {
		t.Errorf("Expected 'TDurCalcTypeCUMWEEKS' value = 2. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_004(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcTypeCUMDAYS

	i = int(r)

	if i != 3 {
		t.Errorf("Expected 'TDurCalcTypeCUMDAYS' value = 3. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_005(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcTypeCUMHOURS

	i = int(r)

	if i != 4 {
		t.Errorf("Expected 'TDurCalcTypeCUMHOURS' value = 4. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_006(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcTypeCUMMINTUES

	i = int(r)

	if i != 5 {
		t.Errorf("Expected 'TDurCalcTypeCUMMINTUES' value = 5. Instead, got %v", i)
	}

}

func TestTDurCalcTypeValue_007(t *testing.T) {

	var r TDurCalcType

	var i int

	r = TDurCalcTypeGregorianYrs

	i = int(r)

	if i != 6 {
		t.Errorf("Expected 'TDurCalcTypeGregorianYrs' value = 6. Instead, got %v", i)
	}

}
