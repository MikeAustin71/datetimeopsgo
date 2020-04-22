package datetimeexamples

import (
	"fmt"
	dt "github.com/MikeAustin71/datetimeopsgo/datetime"
	"time"
)

// ExampleTimeDto001
func ExampleTimeDto001() {

	t0Dto, err := dt.TimeDto{}.NewTimeComponents(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		fmt.Printf("Error returned by TimeDto{}.NewStartEndTimes(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'\n", err.Error())
		return
	}

	fmt.Println("====================================")
	fmt.Println("         Original TimeDto")
	fmt.Println("====================================")
	PrintOutTimeDtoFields(t0Dto)

	t2Dto, err := dt.TimeDto{}.NewTimeComponents(0, 14, 0, 0, 0, 0, 0, 0, 0, 0)

	err = t0Dto.AddTimeDto(t2Dto)

	if err != nil {
		fmt.Printf("Error returned by t0Dto.AddTimeDto(t2Dto). Error='%v'", err.Error())
		return
	}

	fmt.Println("Added Months= ", 14)

	fmt.Println("====================================")
	fmt.Println("         Final TimeDto")
	fmt.Println("====================================")
	PrintOutTimeDtoFields(t0Dto)

}

// ExampleTimeDto002
func ExampleTimeDto002() {

	tDto, err := dt.TimeDto{}.NewTimeComponents(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		fmt.Printf("Error: dt.TimeDto{}.NewStartEndTimes() Error='%v'", err.Error())
	}

	fmt.Println("--------- TimeDto --------")
	PrintOutTimeDtoFields(tDto)

}

// ExampleTimeDto003
func ExampleTimeDto003() {

	locUTC, _ := time.LoadLocation(dt.TZones.UTC())

	realT0 := time.Time{}.In(locUTC).AddDate(-1, 0, 0)

	fmt.Println("realT0: ", realT0.Format(dt.FmtDateTimeYrMDayFmtStr))

	tDto := dt.TimeDto{}
	tDto.Years = 0
	tDto.Months = 0
	tDto.Weeks = -56

	fmt.Println("====================================")
	fmt.Println("         Original TimeDto")
	fmt.Println("====================================")
	PrintOutTimeDtoFields(tDto)

	err := tDto.NormalizeTimeElements()

	if err != nil {
		panic(fmt.Errorf("ExampleTimeDto003() Error returned by tDto.NormalizeTimeElements(). "+
			"Error='%v' ", err.Error()))
	}

	fmt.Println("------------------------------------")
	fmt.Println("        Normalized TimeDto")
	fmt.Println("------------------------------------")
	PrintOutTimeDtoFields(tDto)

	startDate, _ := tDto.GetDateTime(dt.TZones.UTC())

	fmt.Println("Calculated Start Date: ", startDate.Format(dt.FmtDateTimeYrMDayFmtStr))

	/*
		====================================
		         Original TimeDto
		====================================
		========================================
		          TimeDto Printout
		========================================
		                   Years:  0
		                  Months:  0
		                   Weeks:  -8
		                WeekDays:  0
		                DateDays:  0
		                   Hours:  0
		                 Minutes:  0
		                 Seconds:  0
		            Milliseconds:  0
		            Microseconds:  0
		             Nanoseconds:  0
		Total SubSec Nanoseconds:  0
		  Total Time Nanoseconds:  0
		========================================
		------------------------------------
		        Normalized TimeDto
		------------------------------------
		========================================
		          TimeDto Printout
		========================================
		                   Years:  -1
		                  Months:  11
		                   Weeks:  4
		                WeekDays:  2
		                DateDays:  30
		                   Hours:  0
		                 Minutes:  0
		                 Seconds:  0
		            Milliseconds:  0
		            Microseconds:  0
		             Nanoseconds:  0
		Total SubSec Nanoseconds:  0
		  Total Time Nanoseconds:  0
		========================================
	*/

}

// ExampleTimeDto004
func ExampleTimeDto004() {

	year := 69
	month := 5
	day := 27
	hour := 15
	minute := 30
	second := 2
	millisecond := 784
	microsecond := 303
	nanosecond := 848

	tDto := dt.TimeDto{}
	tDto.Years = year
	tDto.Months = month
	tDto.DateDays = day
	tDto.Hours = hour
	tDto.Minutes = minute
	tDto.Seconds = second
	tDto.Milliseconds = millisecond
	tDto.Microseconds = microsecond
	tDto.Nanoseconds = nanosecond

	fmt.Println("====================================")
	fmt.Println("         Original TimeDto")
	fmt.Println("====================================")
	PrintOutTimeDtoFields(tDto)

	err := tDto.NormalizeTimeElements()

	if err != nil {
		panic(fmt.Errorf("ExampleTimeDto004() Error returned by "+
			"tDto.NormalizeTimeElements(). Error='%v' ", err.Error()))
	}

	fmt.Println("------------------------------------")
	fmt.Println("        Normalized TimeDto")
	fmt.Println("------------------------------------")
	PrintOutTimeDtoFields(tDto)

}

// ExampleTimeDto005
func ExampleTimeDto005() {
	year := 69
	month := 5
	day := 27
	hour := 15
	minute := 30
	second := 2
	millisecond := 784
	microsecond := 303
	nanosecond := 848

	tDto, err := dt.TimeDto{}.NewTimeComponents(year, month, 0, day, hour, minute, second,
		millisecond, microsecond, nanosecond)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDto{}.NewStartEndTimes(). Error='%v' \n", err.Error())
		return
	}

	locUSCentral, _ := time.LoadLocation(dt.TZones.US.Central())

	t1USCentral := time.Date(1948, time.Month(9), 7, 4, 32, 16, 8185431, locUSCentral)
	tDur, err := dt.TimeDurationDto{}.NewStartTimePlusTimeDto(
		t1USCentral,
		tDto,
		dt.TDurCalc.StdYearMth(),
		t1USCentral.Location().String(),
		dt.TCalcMode.LocalTimeZone(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}.NewStartTimePlusTimeDto(). "+
			"t1USCentral='%v'  Error:='%v'\n",
			t1USCentral.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	t4USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)

	if !t4USCentral.Equal(tDur.GetThisEndDateTime()) {
		fmt.Printf("Error: expected EndDateTime='%v'.\n" +
			"Instead, EndDateTime='%v'\n"+
			t4USCentral.Format(dt.FmtDateTimeYrMDayFmtStr),
			tDur.GetThisEndDateTime().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

}

// ExampleTimeDto006
func ExampleTimeDto006() {

	t1Dto := dt.TimeDto{}

	t1Dto.Years = -1
	t1Dto.Months = 15
	t1Dto.DateDays = 0
	t1Dto.Hours = 0
	t1Dto.Minutes = 0
	t1Dto.Seconds = 0
	t1Dto.Milliseconds = 0
	t1Dto.Microseconds = 0
	t1Dto.Nanoseconds = 0

	fmt.Println("=================================================")
	fmt.Println("            Original TimeDto")
	fmt.Println("=================================================")
	PrintOutTimeDtoFields(t1Dto)

	err := t1Dto.NormalizeTimeElements()

	if err != nil {
		fmt.Printf("Error returned by t1Dto.NormalizeTimeElements(). Error='%v'\n", err.Error())
		return
	}

	fmt.Println()
	fmt.Println("-------------------------------------------------")
	fmt.Println("             Normalized TimeDto")
	fmt.Println("-------------------------------------------------")
	PrintOutTimeDtoFields(t1Dto)

}

// ExampleTimeDto007
func ExampleTimeDto007() {

	t0Dto, err := dt.TimeDto{}.NewTimeComponents(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		fmt.Printf("Error returned by TimeDto{}.NewStartEndTimes(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'\n", err.Error())
		return
	}

	fmt.Println("Original t0Dto - TimeDto - Data Fields")
	PrintOutTimeDtoFields(t0Dto)
	fmt.Println("-----------------------------------------")
	fmt.Println()

	t2Dto, err := dt.TimeDto{}.NewTimeComponents(0, 36, 0, 0, 0, 0, 0, 0, 0, 0)

	if err != nil {
		fmt.Printf("Error returned by TimeDto{}.NewStartEndTimes(0, 0, 0, 1, 0, 0,0,0,0,0). Error='%v'\n", err.Error())
		return
	}

	err = t0Dto.AddTimeDto(t2Dto)
	if err != nil {
		fmt.Printf("Error returned by t0Dto.AddTimeDto(t2Dto, dt.TZones.US.Central()). Error='%v'\n", err.Error())
		return
	}

	fmt.Println("Final t0Dto - TimeDto - Data Fields")
	PrintOutTimeDtoFields(t0Dto)

}
