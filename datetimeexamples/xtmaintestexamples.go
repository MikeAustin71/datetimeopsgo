package datetimeexamples

import (
	"fmt"
	dt "github.com/MikeAustin71/datetimeopsgo/datetime"
	"strings"
	"time"
)

type MainCodeExamples struct {
	Input   string
	Output  string
}

func (mc MainCodeExamples) mainCodeEx049() {

	ePrefix := "mainCodeEx049()"
	lineLen := 65
	mc.mainPrintHdr(ePrefix , "-")

	// t1EdtStr :=  "06/20/2019 09:58:32 -0400 EDT"
	// t1EstStr :=  "12/20/2019 09:58:32 -0500 EST"
	// t2CdtStr :=  "06/20/2019 09:58:32 -0500 CDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0600 CST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0600 MDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0700 MST"
	t2PdtStr :=  "06/20/2019 09:58:32 -0700 PDT"
	// t2PstStr :=  "12/20/2019 09:58:32 -0800 PST"

	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	timeStr := t2PdtStr
	t1, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1 Initial Date Time"},
		lineLen,
		fmtStr)

}

func (mc MainCodeExamples) mainCodeEx048() {

	ePrefix := "mainCodeEx048()"
	lineLen := 65
	mc.mainPrintHdr(ePrefix , "-")

	// t1EdtStr :=  "06/20/2019 09:58:32.000000000 -0400 EDT"
	t1EstStr :=  "12/20/2019 09:58:32.000000000 -0500 EST"
	// t2PdtStr :=  "06/20/2019 09:58:32.000000000 -0500 CDT"
	// t2PdtStr :=  "12/20/2019 09:58:32.000000000 -0600 CST"
	// t3PdtStr :=  "06/20/2019 09:58:32.000000000 -0600 MDT"
	// t3PdtStr :=  "12/20/2019 09:58:32.000000000 -0700 MST"
	// t4PdtStr :=  "06/20/2019 09:58:32.000000000 -0700 PDT"
	// t4PstStr :=  "12/20/2019 09:58:32.000000000 -0800 PST"
	// t5CestStr := "06/20/2019 09:58:32.000000000 +0200 CEST"
	// t7MSKStr := "06/20/2019 09:58:32.000000000 +0300 MSK"

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"
	timeStr := t1EstStr
	t1, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1 Initial Date Time"},
		lineLen,
		fmtStr)

	locNamePtr := t1.Location()
	locName := locNamePtr.String()
	lenLocName := len(locName)

	if lenLocName > 5 {
		_, err = time.LoadLocation(locName)

		if err != nil {
			fmt.Printf(ePrefix +
				"\nError returned by time.LoadLocation(locName)\n" +
				"locName='%v'\nError='%v'\n",
				locName, err.Error())
			return
		}

		fmt.Println("Location Loaded Successfully: ", locNamePtr.String())

	} else {

		offsetLeadLen := len("01/02/2006 15:04:05.000000000 ")

		t2AbbrvLookup := locName + timeStr[offsetLeadLen:offsetLeadLen+5]

		stdAbbrvs := dt.StdTZoneAbbreviations{}

		tZones, ok := stdAbbrvs.AbbrvOffsetToTimeZones(t2AbbrvLookup)

		if !ok {
			fmt.Printf(ePrefix +
				"\nError: Map TzAbbrvs to Time Zones Failed.\n" +
				"Lookup key='%v'\n", t2AbbrvLookup)
			return
		}

		fmt.Println("Abbreviation Lookup: ", t2AbbrvLookup)
		newTZone := ""

		if len(tZones) == 1 {
			newTZone = tZones[0]
		}
		/*
			if len(newTZone) == 0 {
				for i:= 0; i < len(tZones); i++ {
					if (strings.HasPrefix(tZones[i], locName) ||
						strings.HasSuffix(tZones[i], locName)) &&
						len(tZones[i]) > lenLocName {
						newTZone = tZones[i]
						break
					}
				}
			}

			if len(newTZone) == 0 {
				for i:= 0; i < len(tZones); i++ {
					if len(tZones[i]) <= lenLocName &&
						!strings.Contains(tZones[i], "-") {
						newTZone = tZones[i]
						break
					}
				}
			}
		*/

		if len(newTZone) == 0 {

			priorityList := []string {
				"UTC",
				"Etc/UTC",
				"Etc/GMT-0",
				"America/New_York",
				"America/Chicago",
				"America/Denver",
				"America/Los_Angeles",
				"Pacific/Honolulu",
				"America/Anchorage",
				"America/Adak",
				"America/Havana",
				"America/St_Johns",
				"America/Thule",
				"America",
				"EST5EDT",
				"CST6CDT",
				"MST7MDT",
				"PST8PDT",
				"US",
				"Europe/Paris",
				"Europe/London",
				"Europe/Dublin",
				"Europe/Rome",
				"Europe/Madrid",
				"Europe/Kiev",
				"Europe/Moscow",
				"Europe",
				"Asia/Shanghai",
				"Asia/Hong_Kong",
				"Asia/Seoul",
				"Asia/Tokyo",
				"Asia/Calcutta",
				"Asia/Karachi",
				"Asia/Manila",
				"Asia/Jerusalem",
				"Asia/Tel_Aviv",
				"Asia/Jakarta",
				"Asia/Makassar",
				"Asia",
				"Atlantic/Canary",
				"Atlantic",
				"Australia/Sydney",
				"Australia/Darwin",
				"Australia/Melbourne",
				"Australia/Adelaide",
				"Australia/Perth",
				"Australia",
				"Canada",
				"Pacific/Guam",
				"Pacific/Samoa",
				"Pacific",
				"Africa/Cairo",
				"Africa/Johannesburg",
				"Africa/Nairobi",
				"Africa/Lagos",
				"Africa",
				"Indian",
				"Etc",
				"Other",
				"Antarctica/McMurdo",
				"Antarctica",
			}

			for i:=0; i < len(priorityList) && len(newTZone)== 0 ; i++ {

				for j:=0; j < len(tZones); j++ {

					if strings.HasPrefix(tZones[j], priorityList[i]) {
						newTZone = tZones[j]
						break
					}
				}
			}
		}

		if len(newTZone) == 0 {
			newTZone = tZones[0]
		}

		locNamePtr , err = time.LoadLocation(newTZone)

		if err != nil {
			fmt.Printf(ePrefix +
				"\nError loading NewStartEndTimes Time Zone!\n" +
				"time.LoadLocation(newTZone)\n" +
				"newTZone='%v'\nError='%v'\n",
				newTZone, err.Error())
			return
		}

		t1 = t1.In(locNamePtr)
		locName = newTZone

		fmt.Printf("t1 converted to new time zone '%v'\n",
			locName)
	}

	PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1 Final Date Time"},
		lineLen,
		fmtStr)
}



func (mc MainCodeExamples) mainCodeEx047() {
	// CET is a valid time zone
	ePrefix := "mainCodeEx047()"
	lineLen := 65
	mc.mainPrintHdr(ePrefix , "-")

	// t1EdtStr :=  "06/20/2019 09:58:32.000000000 -0400 EDT"
	// t1EstStr :=  "12/20/2019 09:58:32.000000000 -0500 EST"
	// t2PdtStr :=  "06/20/2019 09:58:32.000000000 -0500 CDT"
	// t2PdtStr :=  "12/20/2019 09:58:32.000000000 -0600 CST"
	// t2PdtStr :=  "06/20/2019 09:58:32.000000000 -0600 MDT"
	// t2PdtStr :=  "12/20/2019 09:58:32.000000000 -0700 MST"
	t2PdtStr :=  "06/20/2019 09:58:32.000000000 -0700 PDT"
	// t2PstStr :=  "12/20/2019 09:58:32.000000000 -0800 PST"

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1Edt, err := time.Parse(fmtStr, t2PdtStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, t2PdtStr)\n" +
			"t2PdtStr='%v'\nError='%v'\n", t2PdtStr, err.Error())
		return
	}

	PrintOutDateTimeTimeZoneFields(
		t1Edt,
		[]string{"t1 EST Date Time"},
		lineLen,
		fmtStr)

	newTimeZoneName := "MST"

	newTimeZoneLocPtr, err := time.LoadLocation(newTimeZoneName)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(newTimeZoneName)\n" +
			"newTimeZoneName='%v'\nError='%v'\n",
			newTimeZoneName, err.Error())
		return
	}

	t2Cdt := t1Edt.In(newTimeZoneLocPtr)


	PrintOutDateTimeTimeZoneFields(
		t2Cdt,
		[]string{"t1 CDT Date Time"},
		lineLen,
		fmtStr)

}

func (mc MainCodeExamples) mainCodeEx046() {

	ePrefix := "mainCodeEx046()"
	lineLen := 65
	mc.mainPrintHdr(ePrefix , "-")

	tstr1 := "12/29/2019 17:54:30 -0800 PST"
	tstr2 := "12/29/2019 20:54:30 -0500 EST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	tzLeadLen := len("01/02/2006 15:04:05 -0700 ")

	tIn1, err := time.Parse(fmtStr, tstr1)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, tstr1)\n" +
			"testr='%v'\nError='%v'\n", tstr1, err.Error())
		return
	}

	tIn2, err := time.Parse(fmtStr, tstr2)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, tstr2)\n" +
			"testr='%v'\nError='%v'\n", tstr1, err.Error())
		return
	}

	PrintOutDateTimeTimeZoneFields(
		tIn1,
		[]string{"tIn1"},
		lineLen,
		fmtStr)

	PrintOutDateTimeTimeZoneFields(
		tIn2,
		[]string{"tIn2"},
		lineLen,
		fmtStr)

	tstr1TimeZone := tstr1[tzLeadLen:]

	fmt.Println("Time Zone tstr1: ", tstr1TimeZone)

	fmt.Println("Loading Time Zone: ", tstr1TimeZone)
	// CST6CDT
	tstr1TimeZoneLocPtr, err := time.LoadLocation(tstr1TimeZone)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(tstr1TimeZone).\n" +
			"tstr1TimeZone='%v'\nError='%v'\n", tstr1TimeZone ,err.Error())
		return
	}


	tstr2TimeZone := tstr2[tzLeadLen:]

	fmt.Println("Time Zone tstr2: ", tstr2TimeZone)
	fmt.Println("Loading Time Zone: ", tstr2TimeZone)
	// EST5EDT

	tstr2TimeZoneLocPtr, err := time.LoadLocation(tstr2TimeZone)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(tstr2TimeZone).\n" +
			"tstr2TimeZone='%v'\nError='%v'\n", tstr2TimeZone ,err.Error())
		return
	}

	t1ConvertedByT2 := tIn1.In(tstr2TimeZoneLocPtr)

	t2ConvertedByT1 := tIn2.In(tstr1TimeZoneLocPtr)

	PrintOutDateTimeTimeZoneFields(
		t1ConvertedByT2,
		[]string{"t1ConvertedByT2"},
		lineLen,
		fmtStr)

	PrintOutDateTimeTimeZoneFields(
		t2ConvertedByT1,
		[]string{"t2ConvertedByT1"},
		lineLen,
		fmtStr)

}

func (mc MainCodeExamples) mainCodeEx045() {

	ePrefix := "mainCodeEx045()"
	lineLen := 65
	mc.mainPrintHdr(ePrefix , "-")

	tstr1 := "06/29/2019 19:54:30 -0500 CDT"
	tstr2 := "06/29/2019 20:54:30 -0400 EDT"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	tzLeadLen := len("01/02/2006 15:04:05 -0700 ")

	tIn1, err := time.Parse(fmtStr, tstr1)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, tstr1)\n" +
			"testr='%v'\nError='%v'\n", tstr1, err.Error())
		return
	}

	tIn2, err := time.Parse(fmtStr, tstr2)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, tstr2)\n" +
			"testr='%v'\nError='%v'\n", tstr1, err.Error())
		return
	}

	PrintOutDateTimeTimeZoneFields(
		tIn1,
		[]string{"tIn1"},
		lineLen,
		fmtStr)

	PrintOutDateTimeTimeZoneFields(
		tIn2,
		[]string{"tIn2"},
		lineLen,
		fmtStr)

	tstr1TimeZone := tstr1[tzLeadLen:]

	fmt.Println("Time Zone tstr1: ", tstr1TimeZone)

	fmt.Println("Loading Time Zone: ", "CST6CDT")
	// CST6CDT
	tstr1TimeZoneLocPtr, err := time.LoadLocation("CST6CDT")

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(\"CST6CDT\").\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tstr2TimeZone := tstr2[tzLeadLen:]

	fmt.Println("Time Zone tstr2: ", tstr2TimeZone)
	fmt.Println("Loading Time Zone: ", "EST5EDT")
	// EST5EDT

	tstr2TimeZoneLocPtr, err := time.LoadLocation("EST5EDT")

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(\"EST5EDT\").\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1ConvertedByT2 := tIn1.In(tstr2TimeZoneLocPtr)

	t2ConvertedByT1 := tIn2.In(tstr1TimeZoneLocPtr)

	PrintOutDateTimeTimeZoneFields(
		t1ConvertedByT2,
		[]string{"t1ConvertedByT2"},
		lineLen,
		fmtStr)

	PrintOutDateTimeTimeZoneFields(
		t2ConvertedByT1,
		[]string{"t2ConvertedByT1"},
		lineLen,
		fmtStr)

}

func (mc MainCodeExamples) mainCodeEx044() {

	ePrefix := "mainCodeEx044()"
	lineLen := 65
	mc.mainPrintHdr(ePrefix , "-")

	tstr1 := "06/29/2019 19:54:30 -0500 CDT"
	tstr2 := "06/29/2019 20:54:30 -0400 EDT"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	offsetLeadLen := len("01/02/2006 15:04:05 ")

	tzLeadLen := len("01/02/2006 15:04:05 -0700 ")

	tIn1, err := time.Parse(fmtStr, tstr1)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, tstr1)\n" +
			"testr='%v'\nError='%v'\n", tstr1, err.Error())
		return
	}

	tIn2, err := time.Parse(fmtStr, tstr2)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, tstr2)\n" +
			"testr='%v'\nError='%v'\n", tstr1, err.Error())
		return
	}

	PrintOutDateTimeTimeZoneFields(
		tIn1,
		[]string{"tIn1"},
		lineLen,
		fmtStr)

	PrintOutDateTimeTimeZoneFields(
		tIn2,
		[]string{"tIn2"},
		lineLen,
		fmtStr)


	tIn1V2 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		tIn2.Location())

	PrintOutDateTimeTimeZoneFields(
		tIn1V2,
		[]string{"tIn1V2"},
		lineLen,
		fmtStr)

	tIn2V2 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		tIn1.Location())

	PrintOutDateTimeTimeZoneFields(
		tIn2V2,
		[]string{"tIn2V2"},
		lineLen,
		fmtStr)

	currTz := tstr2[tzLeadLen:]

	t2AbbrvLookup := currTz + tstr2[offsetLeadLen:offsetLeadLen+5]

	stdAbbrvs := dt.StdTZoneAbbreviations{}

	tZones, ok := stdAbbrvs.AbbrvOffsetToTimeZones(t2AbbrvLookup)

	if !ok {
		fmt.Printf(ePrefix +
			"\nError: Map TzAbbrvs to Time Zones Failed.\n" +
			"Lookup key='%v'\n", t2AbbrvLookup)
		return
	}

	newTZone := ""

	for i:= 0; i < len(tZones); i++ {
		if strings.HasPrefix(tZones[i], currTz) ||
			strings.HasSuffix(tZones[i], currTz){
			newTZone = tZones[i]
		}
	}

	if newTZone == "" {
		fmt.Println(ePrefix +
			"\nError: NewStartEndTimes Time Zone Look Up failed!")
		return
	}

	newTZonePtr, err := time.LoadLocation(newTZone)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(newTZone)\n" +
			"newTZone='%v'\nError='%v'\n", newTZone, err.Error())
		return
	}

	convertedTIn1 := tIn1.In(newTZonePtr)
	PrintOutDateTimeTimeZoneFields(
		convertedTIn1,
		[]string{"convertedTIn1"},
		lineLen,
		fmtStr)

}

func (mc MainCodeExamples) mainCodeEx043() {
	ePrefix := "mainCodeEx043()"
	lineLen := 65
	mc.mainPrintHdr(ePrefix , "-")

	tzName := "EST"
	tzLocPtr, err := time.LoadLocation(tzName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzName)\n" +
			"tzName='%v'\n" +
			"Error='%v'\n", tzName, err.Error())
		return
	}

	t1 := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		tzLocPtr)

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1"},
		lineLen,
		fmtStr)

	t1V2str := "06/30/2019 22:58:32.000000000 -0400 EDT"

	t1V2, err := time.Parse(fmtStr, t1V2str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str)\n" +
			"t1V2str='%v'\n" +
			"Error='%v'\n", t1V2str, err.Error())
		return
	}

	PrintOutDateTimeTimeZoneFields(
		t1V2,
		[]string{"t1V2"},
		lineLen,
		fmtStr)

}

func (mc MainCodeExamples) mainCodeEx042() {

	ePrefix := "mainCodeEx042()"

	mc.mainPrintHdr(ePrefix , "-")

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"

	tIn, _ := time.Parse(fmtStr, tstr)

	ianaPacificTzPtr, err := time.LoadLocation(ianaPacificTz)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(ianaPacificTz)\n" +
			"ianaPacificTz='%v'\n" +
			"Error='%v'\n", ianaPacificTz, err.Error())
		return
	}

	tOut := tIn.In(ianaPacificTzPtr)

	fmt.Println(" tIn: ", tIn.Format(fmtStr))
	fmt.Println("tOut: ", tOut.Format(fmtStr))

	dTzDtoIn, err := dt.DateTzDto{}.NewDateTime(tIn, fmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.DateTzDto{}.NewStartEndTimes(tIn, fmtStr)\n" +
			"tIn='%v'\n" +
			"Error='%v'\n", tIn.Format(fmtStr), err.Error())
		return
	}

	tzDefIn := dTzDtoIn.GetTimeZoneDef()

	PrintOutTimeZoneDefFields(tzDefIn)

	dTzDtoOut, err := dt.DateTzDto{}.NewTz(
		tIn,
		ianaPacificTz,
		dt.TzConvertType.Relative(),
		fmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.DateTzDto{}.NewTz(tIn,ianaPacificTz,...)\n" +
			"tIn='%v'\n" +
			"ianaPacificTz='%v'\n" +
			"Error='%v'\n",
			tIn.Format(fmtStr),
			ianaPacificTz,
			err.Error())
		return
	}

	tzDefOut := dTzDtoOut.GetTimeZoneDef()

	PrintOutTimeZoneDefFields(tzDefOut)

	tOut2 := tIn.In(tzDefOut.GetOriginalLocationPtr())

	fmt.Println()
	fmt.Println("Final tOut2: ", tOut2.Format(fmtStr))

}

func (mc MainCodeExamples) mainCodeEx041() {

	ePrefix := "mainCodeEx041()"

	mc.mainPrintHdr(ePrefix , "-")

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"

	tIn, _ := time.Parse(fmtStr, tstr)

	ianaPacificTzPtr, err := time.LoadLocation(ianaPacificTz)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(ianaPacificTz)\n" +
			"ianaPacificTz='%v'\n" +
			"Error='%v'\n", ianaPacificTz, err.Error())
		return
	}

	tOut := tIn.In(ianaPacificTzPtr)

	fmt.Println(" tIn: ", tIn.Format(fmtStr))
	fmt.Println("tOut: ", tOut.Format(fmtStr))

	tzDefIn, err := dt.TimeZoneDefinition{}.NewDateTime(tIn)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(tIn)\n" +
			"tIn='%v'\n" +
			"Error='%v'\n", tIn.Format(fmtStr), err.Error())
		return
	}

	PrintOutTimeZoneDefFields(tzDefIn)

	tzDefOut, err := dt.TimeZoneDefinition{}.NewDateTime(tOut)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(tOut)\n" +
			"tOut='%v'\n" +
			"Error='%v'\n", tIn.Format(fmtStr), err.Error())
		return
	}

	PrintOutTimeZoneDefFields(tzDefOut)

}

func (mc MainCodeExamples) mainCodeEx040() {
	/*
		--- FAIL: TestTimeZoneUtility_GetZoneOut_01 (0.00s)
		    zzttimezonedto01_test.go:528: Expected Zone Out='PDT'.
		        Instead, actual Zone Out='PST'
	*/

	ePrefix := "mainCodeEx040()"

	mc.mainPrintHdr(ePrefix , "-")

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtStr, tstr)
	tzu, _ := dt.TimeZoneDto{}.New(tIn, ianaPacificTz, fmtStr)

	fmt.Println("       tstr: ", tstr)
	inTimeStr := tzu.TimeIn.GetDateTimeValue().Format(fmtStr)
	fmt.Println(" tzu.TimeIn: ", inTimeStr)
	outTimeStr := tzu.TimeOut.GetDateTimeValue().Format(fmtStr)
	fmt.Println("tzu.TimeOut: ", outTimeStr)

	timeInDef := tzu.TimeIn.GetTimeZoneDef()

	PrintOutTimeZoneDefFields(timeInDef)

	timeOutDef := tzu.TimeOut.GetTimeZoneDef()

	PrintOutTimeZoneDefFields(timeOutDef)

	expectedZone := "PDT"

	actualZone := tzu.TimeOut.GetOriginalTzAbbreviation()

	if expectedZone != actualZone {
		fmt.Printf("Expected Zone Out='%v'.\n" +
			"Instead, actual Zone Out='%v'\n", expectedZone, actualZone)
	}

}

func (mc MainCodeExamples) mainCodeEx039() {

	ePrefix := "mainCodeEx039()"

	mc.mainPrintHdr(ePrefix , "-")

	tzName := "Australia/Adelaide"

	tzLocPtr, err := time.LoadLocation(tzName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzName)\n" +
			"tzName='%v'\n" +
			"Error='%v'\n", tzName, err.Error())
		return
	}

	t1 := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		tzLocPtr)

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	tzDef, err := dt.TimeZoneDefinition{}.NewDateTime(t1)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(t1)\n" +
			"t1='%v'\n" +
			"Error='%v'\n", t1.Format(fmtStr), err.Error())
		return
	}

	fmt.Println("t1 Date Time: ", t1.Format(fmtStr))

	PrintOutTimeZoneDefFields(tzDef)

	t2 := time.Date(
		2019,
		time.Month(12),
		30,
		22,
		58,
		32,
		0,
		tzLocPtr)

	tzDef, err = dt.TimeZoneDefinition{}.NewDateTime(t2)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(t2)\n" +
			"t2='%v'\n" +
			"Error='%v'\n",
			t2.Format(fmtStr),t1.Format(fmtStr))
		return
	}

	fmt.Println("t2 Date Time: ", t2.Format(fmtStr))

	PrintOutTimeZoneDefFields(tzDef)
}

func (mc MainCodeExamples) mainCodeEx038() {

	ePrefix := "mainCodeEx038()"
	lineLen := 65
	mc.mainPrintHdr(ePrefix , "-")

	t1str := "06/30/2019 22:58:32.000000000 -0400 EDT"
	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	tZoneName := dt.TZones.Other.EST05EDT()

	easternLocPtr, err := time.LoadLocation(tZoneName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.US.Central())\n" +
			"tZoneName='%v'\n" +
			"Error='%v'\n", tZoneName, err.Error())
		return
	}

	t1, err := time.Parse(fmtStr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str)\n" +
			"t1str='%v'\n" +
			"Error='%v'\n", t1str, err.Error())
		return
	}

	t2 := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		easternLocPtr)

	PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1 Date Time"},
		lineLen,
		fmtStr)

	PrintOutDateTimeTimeZoneFields(
		t2,
		[]string{"t2 Date Time"},
		lineLen,
		fmtStr)


}

func (mc MainCodeExamples) mainCodeEx037() {

	ePrefix := "mainCodeEx037()"

	mc.mainPrintHdr(ePrefix , "-")

	tzName := "EST"

	tzNameLocPtr, err := time.LoadLocation(tzName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzName).\n" +
			"tzName='%v'\n" +
			"Error='%v'\n", tzName, err.Error())
		return
	}

	t1 := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		tzNameLocPtr)

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	fmt.Println("t1 Date Time: ", t1.Format(fmtStr))

	tz2Name := dt.TZones.America.New_York()

	tz2NameLocPtr, err := time.LoadLocation(tz2Name)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzName).\n" +
			"tzName='%v'\n" +
			"Error='%v'\n", tzName, err.Error())
		return
	}

	t2 := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		tz2NameLocPtr)

	fmt.Println("t2 Date Time: ", t2.Format(fmtStr))
}

func (mc MainCodeExamples) mainCodeEx036() {

	ePrefix := "mainCodeEx036()"

	mc.mainPrintHdr(ePrefix , "-")

	t1str := "12/30/2019 22:58:32.000000000 -0400 EDT"
	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, err := time.Parse(fmtStr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtStr, t1str)\n" +
			"t1str='%v'\n" +
			"Error='%v'\n", t1str, err.Error())
		return
	}


	tzName := t1.Location().String()
	t1UTC := t1.In(time.UTC)
	fmt.Println("    t1 date Time String: ", t1str)
	fmt.Println("           t1 date time: ", t1.Format(fmtStr))
	fmt.Println("      t1 Time Zone Name: ", tzName)

	t1V2 := time.Date(
		t1.Year(),
		t1.Month(),
		t1.Day(),
		t1.Hour(),
		t1.Minute(),
		t1.Second(),
		t1.Nanosecond(),
		t1.Location())

	fmt.Println("  t1V2 Date Time String: ", t1V2.Format(fmtStr))
	fmt.Println("t1 UTC Date Time String: ", t1UTC.Format(fmtStr))
	fmt.Println()
	t1TzLoaded := "EST"
	_, err = time.LoadLocation(t1TzLoaded)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzName).\n" +
			"tzName='%v'\n" +
			"Error='%v'\n", t1TzLoaded, err.Error())
		return
	}


	t2str := "06/30/2019 22:58:32.000000000 -0500 CDT"

	t2, err := time.Parse(fmtStr, t2str)

	if err != nil {
		fmt.Printf("Error returned from time.Parse(fmtStr, t2str)\n" +
			"t2str='%v'\n" +
			"Error='%v'\n", t2str, err.Error())
		return
	}

	t3 := t2.In(t1.Location())
	tzName = t3.Location().String()

	fmt.Println("t2 date Time String: ", t2.Format(fmtStr))
	fmt.Println("       t3 date time: ", t3.Format(fmtStr))
	fmt.Println("  t3 Time Zone Name: ", tzName)
	fmt.Println()

}

func (mc MainCodeExamples) mainCodeEx035() {
	ePrefix := "mainCodeEx035()"

	mc.mainPrintHdr(ePrefix , "-")

	t1str := "06/30/2019 22:58:32.000000000 -0400 EDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	tZoneName := dt.TZones.US.Eastern()

	easternLocPtr, err := time.LoadLocation(tZoneName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.US.Central())\n" +
			"tZoneName='%v'\n" +
			"Error='%v'\n", tZoneName, err.Error())
		return
	}

	t1, err := time.Parse(fmtstr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str)\n" +
			"t1str='%v'\n" +
			"Error='%v'\n", t1str, err.Error())
		return
	}

	t1Dup := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		easternLocPtr)

	fmt.Printf("   t1 Time: %v\n", t1.Format(fmtstr))
	fmt.Printf("t1Dup Time: %v\n", t1Dup.Format(fmtstr))

	if t1.Equal(t1Dup) {
		fmt.Println("--- t1 is EQUAL to t1Dup ---")
	} else {
		fmt.Println("--- t1 is NOT equal to t1Dup ---")
	}

	fmt.Println()

	t1TzDefDtoDateTime, err :=
		dt.TimeZoneDefinition{}.NewFromTimeZoneName(
			t1,
			tZoneName,
			dt.TzConvertType.Relative())

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewFromTimeZoneName(tZoneName)\n" +
			"tZoneName='%v'\n" +
			"Error='%v'\n", tZoneName, err.Error())
	}

	mc.mainPrintHdr("       t1TzDefDtoDateTime Data" , "=")
	fmt.Println("          From t1")
	PrintOutTimeZoneDefFields(t1TzDefDtoDateTime)
	fmt.Println()

	t2TzDefDtoDateTime, err :=
		dt.TimeZoneDefinition{}.NewDateTime(t1)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(t1)\n" +
			"t1='%v'\n" +
			"Error='%v'\n", t1.Format(fmtstr), err.Error())
		return
	}

	mc.mainPrintHdr("       t2TzDefDtoDateTime Data" , "=")
	fmt.Println("          From t1")
	PrintOutTimeZoneDefFields(t2TzDefDtoDateTime)
	fmt.Println()


	t1TzDefDtoDateTime, err =
		dt.TimeZoneDefinition{}.NewFromTimeZoneName(
			t1,
			tZoneName,
			dt.TzConvertType.Relative())

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewFromTimeZoneName(tZoneName)\n" +
			"tZoneName='%v'\n" +
			"Error='%v'\n", tZoneName, err.Error())
	}

	t2TzDefDtoDateTime, err =
		dt.TimeZoneDefinition{}.NewDateTime(t1Dup)

	mc.mainPrintHdr("       t1TzDefDtoDateTime Data" , "*")
	fmt.Println("From: t1Dup")
	PrintOutTimeZoneDefFields(t1TzDefDtoDateTime)
	fmt.Println()

	mc.mainPrintHdr("       t2TzDefDtoDateTime Data" , "*")
	fmt.Println("From: t1Dup")
	PrintOutTimeZoneDefFields(t2TzDefDtoDateTime)
	fmt.Println()

}

func (mc MainCodeExamples) mainCodeEx034() {

	ePrefix := "mainCodeEx034()"
	title := fmt.Sprintf("       %v         ", ePrefix)
	ln := strings.Repeat("-", len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	// t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t2, _ := time.Parse(fmtstr, t2str)

	tzu1, err := dt.TimeZoneDto{}.New(t2, dt.TZones.US.Eastern(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.NewStartEndTimes(t1, TzUsEast).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tzu1.IsValid()

	if err != nil {
		fmt.Printf("tzu1 is INVALID!\n" +
			"%v\n", err.Error())
		return
	}

}

func (mc MainCodeExamples) mainCodeEx033() {

	ePrefix := "mainCodeEx033()"
	title := fmt.Sprintf("       %v         ", ePrefix)
	ln := strings.Repeat("-", len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	locUSCentral, err := time.LoadLocation(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	locTokyo, err := time.LoadLocation(dt.TZones.Asia.Tokyo())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
		return
	}

	t4USCentral := time.Date(
		2018,
		time.Month(3),
		06,
		20,
		02,
		18,
		792489279, locUSCentral)

	t4AsiaTokyo := t4USCentral.In(locTokyo)

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dTz, err := dt.DateTzDto{}.NewTz(
		t4AsiaTokyo,
		dt.TZones.US.Central(),
		dt.TzConvertType.Relative(),
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewTz(t4AsiaTokyo, TZones.US.Central(), fmtstr).\n" +
			"Error='%v'\n",
			err.Error())
		return
	}

	if !t4USCentral.Equal(dTz.GetDateTimeValue()) {
		fmt.Printf("Error: dTz Toyko to USA-Central\n" +
			"Expected DateTime='%v'.\n" +
			" Instead DateTime='%v'\n",
			t4USCentral.Format(fmtstr), dTz.GetDateTimeValue().Format(fmtstr))
		return
	}

	eTimeZoneDef, err := dt.TimeZoneDefinition{}.NewDateTime(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefinition{}.NewStartEndTimes(t4USCentral)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !eTimeZoneDef.Equal(dTz.GetTimeZoneDef()) {
		fmt.Printf("Expected dTz.GetTimeZoneDef().LocationName='%v'.\n"+
			"Instead, dTz.GetTimeZoneDef().LocationName='%v'\n",
			eTimeZoneDef.GetOriginalLocationName(), dTz.GetOriginalTzName())
		return
	}

	tDto, err := dt.TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral)\n"+
			"t4USCentral='%v'\nError='%v'\n",
			t4USCentral.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	dTzTimeComponents := dTz.GetTimeComponents()

	fmt.Println("tDto Comparison to dTz ")
	fmt.Printf("               Years: %v  -  %v\n", tDto.Years, dTzTimeComponents.Years )
	fmt.Printf("              Months: %v  -  %v\n", tDto.Months, dTzTimeComponents.Months )
	fmt.Printf("                Days: %v  -  %v\n", tDto.DateDays, dTzTimeComponents.DateDays )
	fmt.Printf("               Hours: %v  -  %v\n", tDto.Hours, dTzTimeComponents.Hours )
	fmt.Printf("             Minutes: %v  -  %v\n", tDto.Minutes, dTzTimeComponents.Minutes )
	fmt.Printf("             Seconds: %v  -  %v\n", tDto.Seconds, dTzTimeComponents.Seconds )
	fmt.Printf("        Milliseconds: %v  -  %v\n", tDto.Milliseconds, dTzTimeComponents.Milliseconds )
	fmt.Printf("        Microseconds: %v  -  %v\n", tDto.Microseconds, dTzTimeComponents.Microseconds )
	fmt.Printf("         Nanoseconds: %v  -  %v\n", tDto.Nanoseconds, dTzTimeComponents.Nanoseconds )
	fmt.Printf("TotSubSecNanoseconds: %v  -  %v\n", tDto.TotSubSecNanoseconds, dTzTimeComponents.TotSubSecNanoseconds )
	fmt.Println()
	fmt.Printf("                dTz Date Time: %v \n", dTz.GetDateTimeValue().Format(fmtstr))

	dTzTimeDtoDateTime, err := dTzTimeComponents.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by dTzTimeComponents.GetDateTime(dt.TZones.US.Central()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("dTz Time Components Date Time: %v\n", dTzTimeDtoDateTime.Format(fmtstr))

	tDtoDateTime, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by tDto.GetDateTime(dt.TZones.US.Central()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("             tDto USA Central: %v \n", tDtoDateTime.Format(fmtstr))
	fmt.Println()

	expectedDt, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from tDto.GetDateTime(TZones.US.Central()). "+
			"Error='%v'", err.Error())
		return
	}

	actualDt, err := dTzTimeComponents.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from dTz.GetTimeComponents().GetDateTime(TZones.US.Central()).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !tDto.Equal(dTz.GetTimeComponents()) {
		fmt.Printf("Expected dTz.Time (TimeDto) == '%v'\n" +
			"Instead, dTz.Time (TimeDto) == '%v'\n",
			expectedDt.Format(dt.FmtDateTimeYrMDayFmtStr), actualDt.Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	if dt.FmtDateTimeYrMDayFmtStr != dTz.GetDateTimeFmt() {
		fmt.Printf("Expected dTz.GetDateTimeFmt()='%v' Instead, dTz.GetDateTimeFmt()='%v' ",
			dt.FmtDateTimeYrMDayFmtStr, dTz.GetDateTimeFmt())
		return
	}

	title = fmt.Sprintf("       %v         ", "!!! Success !!!")
	ln = strings.Repeat("-", len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

}


func (mc MainCodeExamples) mainCodeEx032() {

	ePrefix := "mainCodeEx032()"
	title := fmt.Sprintf("       %v         ", ePrefix)
	ln := strings.Repeat("-", len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	t1Str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	expectedStr := "2014-02-15 19:54:30.038175584 -0500 EST"

	dt1, err := time.Parse(fmtStr, t1Str)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, t1Str).\n" +
			"t1Str='%v'\n", t1Str)
		return
	}

	dtMech := dt.DTimeMechanics{}

	dt2, err := dtMech.AbsoluteTimeToTimeZoneNameConversion(dt1, dt.TZones.America.New_York(),ePrefix)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by dtMech.AbsoluteTimeToTimeZoneDefConversion(dt1, dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dt2Str := dt2.Format(fmtStr)

	var result string

	if dt2Str != expectedStr {
		result = "*** FAILURE ***"
	} else {
		result = "!!! SUCCESS !!!"
	}

	title = fmt.Sprintf("       %v         ", result)
	ln = strings.Repeat("=", len(title))

	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	fmt.Printf("   Start Date Time: %v\n", t1Str)
	fmt.Printf("  Actual Date Time: %v\n", dt2Str)
	fmt.Printf("Expected Date Time: %v\n", expectedStr)
	fmt.Println()
}

func (mc MainCodeExamples) mainCodeEx031() {

	ePrefix := "mainCodeEx031()"
	title := fmt.Sprintf("       %v         ", ePrefix)
	ln := strings.Repeat("-", len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	t1Str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	expectedStr := "2014-02-15 19:54:30.038175584 -0500 EST"

	dt1, err := time.Parse(fmtStr, t1Str)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, t1Str).\n" +
			"t1Str='%v'\n", t1Str)
		return
	}

	tzDefDto, err := dt.TimeZoneDefinition{}.NewFromTimeZoneName(
		dt1,
		dt.TZones.America.New_York(),
		dt.TzConvertType.Relative())

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by dt.TimeZoneDefinition{}.NewFromTimeZoneName(dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dtMech := dt.DTimeMechanics{}

	dt2, err := dtMech.AbsoluteTimeToTimeZoneDefConversion(dt1, tzDefDto)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by dtMech.AbsoluteTimeToTimeZoneDefConversion(dt1, tzDefDto)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dt2Str := dt2.Format(fmtStr)

	var result string

	if dt2Str != expectedStr {
		result = "*** FAILURE ***"
	} else {
		result = "!!! SUCCESS !!!"
	}

	title = fmt.Sprintf("       %v         ", result)
	ln = strings.Repeat("=", len(title))

	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	fmt.Printf("   Start Date Time: %v\n", t1Str)
	fmt.Printf("  Actual Date Time: %v\n", dt2Str)
	fmt.Printf("Expected Date Time: %v\n", expectedStr)
	fmt.Println()
}

func (mc MainCodeExamples) mainCodeEx030() {

	fmt.Println("       mainCodeEx030()         ")
	fmt.Println("-----------------------------")
	fmt.Println()

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, err := time.Parse(fmtstr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str).\n" +
			"Error='%v'", err.Error())
		return
	}

	// tzLocName := "America/Chicago"
	tzLocName := "Local"

	tzLoc, err := time.LoadLocation(tzLocName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzLocName)\n" +
			"tzLocName='%v'\n" +
			"Error='%v'\n", tzLocName, err.Error())

		return
	}

	t2 := t1.In(tzLoc)

	dtz1, err := dt.DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewStartEndTimes(t2, fmtstr).\n" +
			"t2='%v'\n" +
			"Error='%v'\n", t2.Format(fmtstr), err.Error())
		return
	}

	tZoneDef := dtz1.GetTimeZoneDef()

	offsetSignChar,
	_,
	offsetHours,
	offsetMinutes,
	offsetSeconds := tZoneDef.GetOriginalOffsetElements()

	fmt.Println("------ Success!!! ------")
	fmt.Println()
	fmt.Printf("   Time Zone Name: %v\n", tzLocName)
	fmt.Println("  -- tZoneDef Values -- ")
	fmt.Println()
	fmt.Printf("         ZoneName: %v\n", tZoneDef.GetOriginalZoneName())
	fmt.Printf("ZoneOffsetSeconds: %v\n", tZoneDef.GetOriginalZoneOffsetTotalSeconds())
	fmt.Printf("         ZoneSign: %v\n", offsetSignChar)
	fmt.Printf("      OffsetHours: %v\n", offsetHours)
	fmt.Printf("    OffsetMinutes: %v\n", offsetMinutes)
	fmt.Printf("    OffsetSeconds: %v\n", offsetSeconds)
	fmt.Printf("       ZoneOffset: %v\n", tZoneDef.GetOriginalZoneOffset())
	fmt.Printf("       UTC Offset: %v\n", tZoneDef.GetOriginalUtcOffset())
	fmt.Printf("    Location Name: %v\n", tZoneDef.GetOriginalLocationName())
	fmt.Printf("        *Location: %v\n", tZoneDef.GetOriginalLocationPtr().String())
	fmt.Printf("      Description: %v\n", tZoneDef.GetOriginalTagDescription())
	fmt.Println()
	fmt.Println()
}
func (mc MainCodeExamples) mainCodeEx029() {

	fmt.Println("       mainCodeEx029()         ")
	fmt.Println("-----------------------------")
	fmt.Println()

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, err := time.Parse(fmtstr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str).\n" +
			"Error='%v'", err.Error())
		return
	}

	// tzLocName := "America/Chicago"
	tzLocName := "Asia/Kabul"

	tzLoc, err := time.LoadLocation(tzLocName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzLocName)\n" +
			"tzLocName='%v'\n" +
			"Error='%v'\n", tzLocName, err.Error())

		return
	}

	t2 := t1.In(tzLoc)

	dtz1, err := dt.DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewStartEndTimes(t2, fmtstr).\n" +
			"t2='%v'\n" +
			"Error='%v'\n", t2.Format(fmtstr), err.Error())
		return
	}

	tZoneDef := dtz1.GetTimeZoneDef()

	offsetSignChar,
	_,
	offsetHours,
	offsetMinutes,
	offsetSeconds := tZoneDef.GetOriginalOffsetElements()

	fmt.Println("------ Success!!! ------")
	fmt.Println()
	fmt.Printf("   Time Zone Name: %v\n", tzLocName)
	fmt.Println("  -- tZoneDef Values -- ")
	fmt.Println()
	fmt.Printf("         ZoneName: %v\n", tZoneDef.GetOriginalZoneName())
	fmt.Printf("ZoneOffsetSeconds: %v\n", tZoneDef.GetOriginalZoneOffsetTotalSeconds())
	fmt.Printf("         ZoneSign: %v\n", offsetSignChar)
	fmt.Printf("      OffsetHours: %v\n", offsetHours)
	fmt.Printf("    OffsetMinutes: %v\n", offsetMinutes)
	fmt.Printf("    OffsetSeconds: %v\n", offsetSeconds)
	fmt.Printf("       ZoneOffset: %v\n", tZoneDef.GetOriginalZoneOffset())
	fmt.Printf("       UTC Offset: %v\n", tZoneDef.GetOriginalUtcOffset())
	fmt.Printf("    Location Name: %v\n", tZoneDef.GetOriginalLocationName())
	fmt.Printf("        *Location: %v\n", tZoneDef.GetOriginalLocationPtr().String())
	fmt.Printf("      Description: %v\n", tZoneDef.GetOriginalTagDescription())
	fmt.Println()
	fmt.Println()
}

func (mc MainCodeExamples) mainCodeEx028() {

	fmt.Println("       mainCodeEx028()         ")
	fmt.Println("-----------------------------")
	fmt.Println()

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, err := time.Parse(fmtstr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str).\n" +
			"Error='%v'", err.Error())
		return
	}

	tzLocName := "Asia/Ho_Chi_Minh"

	tzLoc, err := time.LoadLocation(tzLocName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzLocName)\n" +
			"tzLocName='%v'\n" +
			"Error='%v'\n", tzLocName, err.Error())

		return
	}

	t2 := t1.In(tzLoc)

	dtz1, err := dt.DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewStartEndTimes(t2, fmtstr).\n" +
			"t2='%v'\n" +
			"Error='%v'\n", t2.Format(fmtstr), err.Error())
		return
	}

	tZoneDef := dtz1.GetTimeZoneDef()

	offsetSignChar,
	_,
	offsetHours,
	offsetMinutes,
	offsetSeconds := tZoneDef.GetOriginalOffsetElements()

	fmt.Println("------ Success!!! ------")
	fmt.Println()
	fmt.Printf("   Time Zone Name: %v\n", tzLocName)
	fmt.Println(" --  tZoneDef Values   -- ")
	fmt.Println()
	fmt.Println(" -- Original Time Zone --")
	fmt.Printf("         ZoneName: %v\n", tZoneDef.GetOriginalZoneName())
	fmt.Printf("ZoneOffsetSeconds: %v\n", tZoneDef.GetOriginalZoneOffsetTotalSeconds())
	fmt.Printf("         ZoneSign: %v\n", offsetSignChar)
	fmt.Printf("      OffsetHours: %v\n", offsetHours)
	fmt.Printf("    OffsetMinutes: %v\n", offsetMinutes)
	fmt.Printf("    OffsetSeconds: %v\n", offsetSeconds)
	fmt.Printf("       ZoneOffset: %v\n", tZoneDef.GetOriginalZoneOffset())
	fmt.Printf("       UTC Offset: %v\n", tZoneDef.GetOriginalUtcOffset())
	fmt.Printf("    Location Name: %v\n", tZoneDef.GetOriginalLocationName())
	fmt.Printf("        *Location: %v\n", tZoneDef.GetOriginalLocationPtr().String())
	fmt.Printf("      Description: %v\n", tZoneDef.GetOriginalTagDescription())
}

func (mc MainCodeExamples) mainCodeEx027() {

	fmt.Println("       mainCodeEx027()         ")
	fmt.Println("-----------------------------")
	fmt.Println()

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	expectedOutDate := t1.Format(fmtstr)

	dtz1, err := dt.DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewStartEndTimes(t1, fmtstr).\n" +
			"Error='%v'", err.Error())
		return
	}

	if expectedOutDate != dtz1.String() {
		fmt.Printf("Error: Expected dtz1.String()='%v'.\n" +
			"Instead, dtz1.String()='%v'\n", expectedOutDate, dtz1.String())
	}

	t2 := t1.AddDate(5, 6, 12)

	dtz2, err := dtz1.AddDate(
		dt.TCalcMode.LocalTimeZone(),
		5,
		6,
		12,
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dtz1.AddDate(5, 6, 12, fmtstr).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedOutDate = t2.Format(fmtstr)

	if expectedOutDate != dtz2.String() {
		fmt.Printf("Error: Expected dtz2.String()='%v'.\n" +
			"Instead, dtz2.String()='%v'\n", expectedOutDate, dtz2.String())
		return
	}

	fmt.Println("------ Success!!! ------")

}

/*
	func (mc MainCodeExamples) mainCodeEx026() {

	fmt.Println("       mainCodeEx026()         ")
	fmt.Println("-----------------------------")
	fmt.Println()

	tstr := "12/06/2019 03:12:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	testTime, err := time.Parse(fmtStr, tstr)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtStr, tstr)\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
		return
	}

	var milTzDto dt.MilitaryDateTzDto
	var dateTzDto dt.DateTzDto

	milTzDto, err = dt.MilitaryDateTzDto{}.NewStartEndTimes(testTime, "Sierra")

	if err != nil {
		fmt.Printf("Error returned by MilitaryDateTzDto{}.NewStartEndTimes(testTime, \"Sierra\")\n" +
			"testTime='%v'\n" +
			"Error='%v'\n", testTime.Format(fmtStr), err.Error())
		return
	}

	dateTzDto, err = dt.DateTzDto{}.NewFromMilitaryDateTz(milTzDto, fmtStr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewFromMilitaryDateTz(milTzDto, fmtStr)\n" +
			"milTzDto.DateTime='%v'\n" +
			"Error='%v'\n", milTzDto.DateTime.Format(fmtStr), err.Error())
		return
	}

	fmt.Println("         Success!!           ")
	fmt.Println("-----------------------------")
	fmt.Println()
	fmt.Printf("Military DateTime: %v", dateTzDto.GetDateTimeValue().Format(fmtStr))

}
*/

func (mc MainCodeExamples) mainCodeEx025() {

	// expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	fmt.Println("       mainCodeEx025()     ")

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Eastern(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.NewStartEndTimes(t1, TzUsEast).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tDto := dt.TimeDto{Years: 3, Months: 2, DateDays: 15, Hours: 3, Minutes: 4, Seconds: 2}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddPlusTimeDto(tDto)

	if err != nil {
		fmt.Printf("Error returned by tzu2.AddPlusTimeDto(tDto). " +
			"Error='%v' ", err.Error())
		return
	}

	tzu2TimeInStr := tzu2.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t2OutStr != tzu2TimeInStr {
		fmt.Printf("Error: Expected tzu2.TimeIn='%v'.  Instead, tzu2.TimeIn='%v'. ", t2OutStr, tzu2TimeInStr)
		return
	}

	fmt.Println("       Successful Completion!     ")
}

func (mc MainCodeExamples) mainCodeEx024() {
	ePrefix := "mainTest.mainCodeEx024() "
	fmt.Println(ePrefix)
	tstr := "12/02/2019 22:05:00 -0600 CST"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	var testTime, expectedMilTime time.Time
	var err error
	var actualMilDateTimeGroup, expectedMilDateTimeGroup string
	var milDatTzDto dt.MilitaryDateTzDto
	var expectedMilTimeLoc *time.Location

	testTime, err = time.Parse(fmtstr, tstr)

	expectedMilTimeLoc, err = time.LoadLocation(dt.TZones.Military.Quebec())

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(dt.TZones.Military.Quebec())\n" +
			"dt.TZones.Military.Quebec()='%v'\n" +
			"Error='%v'\n", dt.TZones.Military.Quebec(), err.Error())
	}

	expectedMilTime = testTime.In(expectedMilTimeLoc)

	milDatTzDto, err = dt.MilitaryDateTzDto{}.New(testTime, "Q")

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by MilitaryDateTzDto{}.NewStartEndTimes(testTime, \"Q\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedMilDateTimeGroup, err = dt.DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)\n" +
			"expectedMilTime='%v'\n" +
			"Error='%v'\n",
			expectedMilTime.Format(fmtstr) ,err.Error())
		return
	}

	actualMilDateTimeGroup, err = milDatTzDto.GetOpenDateTimeGroup()

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by milDatTzDto.GetOpenDateTimeGroup()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if expectedMilDateTimeGroup != actualMilDateTimeGroup {
		fmt.Printf(ePrefix +
			"\nError: Expected Military Date Time Group='%v'.\n" +
			"Actual Military Date Time Group='%v'\n" +
			"Military Time='%v'",
			expectedMilDateTimeGroup, actualMilDateTimeGroup, expectedMilTime.Format(fmtstr))
	}

	fmt.Println("***** Success *****")
	fmt.Printf("Expected Military Date Time Group: %v\n", expectedMilDateTimeGroup)
	fmt.Printf("  Actual Military Date Time Group: %v\n", actualMilDateTimeGroup)
	fmt.Printf("              Original Start Time: %v\n", testTime.Format(fmtstr))
	fmt.Printf("      Military Time in Quebect Tz: %v\n", milDatTzDto.DateTime.Format(fmtstr))
}

func (mc MainCodeExamples) mainCodeEx023() {
	ePrefix := "mainTest.mainCodeEx023() "

	fmt.Println(ePrefix)
	tstr := "11/29/2017 19:54:30 -0600 CST"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"

	expected := "291954SNOV17"
	var testTime time.Time
	var err error
	var result string
	var milDatTzDto dt.MilitaryDateTzDto

	testTime, err = time.Parse(fmtstr, tstr)

	milDatTzDto, err = dt.MilitaryDateTzDto{}.New(testTime, "S")

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by MilitaryDateTzDto{}.NewStartEndTimes(testTime, \"S\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	result, err = milDatTzDto.GetCompactDateTimeGroup()

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by milDatTzDto.GetCompactDateTimeGroup()\n")
		return
	}

	if result != expected {
		fmt.Printf(ePrefix +
			"\nError: Expected result='%v'.\n" +
			"Instead, result='%v'.\n", expected, result)
		return
	}

	fmt.Println("  Result= ", result)
	fmt.Println("Expected= ", expected)

}

func (mc MainCodeExamples) mainCodeEx022() {

	ePrefix := "mainTest.mainCodeEx022() "

	fmt.Println(ePrefix)

	t := time.Now().Local()

	militaryDateTime, err := dt.DtMgr{}.GetMilitaryCompactDateTimeGroup(t)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError retunred by dt.DtMgr{}.GetMilitaryCompactDateTimeGroup(t).\n" +
			"t='%v'\nError='%v'\n", t.Format(dt.FmtDateTimeYMDHMSTz),
			err.Error())
		return
	}

	fmt.Println("Standard Format: ", t.Format(dt.FmtDateTimeYMDHMSTz))
	fmt.Println("Military Format: ", militaryDateTime)

}

func (mc MainCodeExamples) mainCodeEx021() {
	tz := dt.TZones.Other.Factory()
	dtz, err := dt.DateTzDto{}.NewNowTz(tz, dt.FmtDateTimeYrMDayFmtStr )

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewNowTz(tz, "	+
			"dt.FmtDateTimeYrMDayFmtStr ).\n Error='%v' \n", err.Error())
		return
	}

	tzDef2, err := dt.TimeZoneDefinition{}.NewDateTime(dtz.GetDateTimeValue())
	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return
	}

	fmt.Println()
	fmt.Println("Testing tz", tz)
	fmt.Println("-------------------------------------")
	fmt.Println("    Zone Name: ", tzDef2.GetOriginalZoneName())
	fmt.Println("  Zone Offset: ", tzDef2.GetOriginalZoneOffset())
	fmt.Println("   UTC Offset: ", tzDef2.GetOriginalUtcOffset())
	fmt.Println("Location Name: ", tzDef2.GetOriginalLocationName())
	fmt.Println("    *Location: ", tzDef2.GetOriginalLocationPtr().String())
}


func (mc MainCodeExamples) mainCodeEx020() {

	tz1 := "Cuba"
	tz2 := "America/Havana"

	dtz, err := dt.DateTzDto{}.NewNowTz(tz1, dt.FmtDateTimeYrMDayFmtStr )

	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return
	}

	tzDef, err := dt.TimeZoneDefinition{}.NewDateTime(dtz.GetDateTimeValue())

	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return
	}

	dtz2, err := dt.DateTzDto{}.NewNowTz(tz2, dt.FmtDateTimeYrMDayFmtStr )

	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return
	}

	tzDef2, err := dt.TimeZoneDefinition{}.NewDateTime(dtz2.GetDateTimeValue())
	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return
	}

	fmt.Println("Testing tz1", tz1)
	fmt.Println("-------------------------------------")
	fmt.Println("    Zone Name: ", tzDef.GetOriginalZoneName())
	fmt.Println("       Offset: ", tzDef.GetOriginalZoneOffset())
	fmt.Println("Location Name: ", tzDef.GetOriginalLocationName())
	fmt.Println("    *Location: ", tzDef.GetOriginalLocationPtr().String())

	fmt.Println()
	fmt.Println("Testing tz2", tz2)
	fmt.Println("-------------------------------------")
	fmt.Println("    Zone Name: ", tzDef2.GetOriginalZoneName())
	fmt.Println("       Offset: ", tzDef2.GetOriginalZoneOffset())
	fmt.Println("Location Name: ", tzDef2.GetOriginalLocationName())
	fmt.Println("    *Location: ", tzDef2.GetOriginalLocationPtr().String())
}

func (mc MainCodeExamples) mainCodeEx019() {
	locUSCentral, err := time.LoadLocation(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	year := 2018
	month := 3
	day := 6
	hour := 20
	minute := 2
	second := 18
	nSecs := 792489279

	t1USCentral := time.Date(year, time.Month(month), day, hour, minute, second, nSecs, locUSCentral)

	minute = 3
	second = 20
	t2USCentral := time.Date(year, time.Month(month), day, hour, minute, second, nSecs, locUSCentral)

	tDur, err := dt.TimeDurationDto{}.NewStartEndTimesCalcTz(
		t1USCentral,
		t2USCentral,
		dt.TDurCalcType(0).StdYearMth(),
		dt.TZones.US.Central(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesCalcTz()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	str, _ := tDur.GetCumSecondsTimeStr()

	fmt.Println("Cumulative Seconds")
	fmt.Println(str)
	// 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds
}

func (mc MainCodeExamples) mainCodeEx018() {
	t1Dto := dt.TimeDto{}

	/*
		t1Dto.Years = 1955
		t1Dto.Months = 15
		t1Dto.DateDays = 32
		t1Dto.Hours = 48
		t1Dto.Minutes = 71
		t1Dto.Seconds = 125
		t1Dto.Milliseconds = 1001
		t1Dto.Microseconds = 1001
		t1Dto.Nanoseconds = 1001
	*/

	t1Dto.Years = 0
	t1Dto.Months = 0
	t1Dto.DateDays = 32
	t1Dto.Hours = 48
	t1Dto.Minutes = 71
	t1Dto.Seconds = 125
	t1Dto.Milliseconds = 0
	t1Dto.Microseconds = 0
	t1Dto.Nanoseconds = 0

	err := t1Dto.NormalizeTimeElements()
	if err != nil {
		fmt.Printf("Error returned by t1Dto.NormalizeTimeElements(). "+
			"Error='%v' ", err.Error())
		return
	}

	fmt.Println("After Normalize Time Elements ")
	PrintOutTimeDtoFields(t1Dto)

	_, err = t1Dto.NormalizeDays()

	if err != nil {
		fmt.Printf("Error returned by t1Dto.NormalizeDays(). "+
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("After ")
	PrintOutTimeDtoFields(t1Dto)
	dateTime, err := t1Dto.GetDateTime(dt.TZones.Other.UTC())

	if err != nil {
		fmt.Printf("Error returned by t1Dto.GetDateTime(dt.TZones.UTC()). Error='%v'\n",
			err.Error())
		return
	}

	fmt.Println("t1Dto.GetDateTime(): ", dateTime.Format(dt.FmtDateTimeYrMDayFmtStr))

}

func (mc MainCodeExamples) mainCodeEx017() {
	t1Dto := dt.TimeDto{}
	/*
		t1Dto.Years = 1955
		t1Dto.Months = 15
		t1Dto.DateDays = 32
		t1Dto.Hours = 48
		t1Dto.Minutes = 71
		t1Dto.Seconds = 125
		t1Dto.Milliseconds = 1001
		t1Dto.Microseconds = 1001
		t1Dto.Nanoseconds = 1001
	*/

	t1Dto.Years = 1955
	t1Dto.Months = 15
	t1Dto.DateDays = 32
	t1Dto.Hours = 48
	t1Dto.Minutes = 71
	t1Dto.Seconds = 125
	t1Dto.Milliseconds = 0
	t1Dto.Microseconds = 0
	t1Dto.Nanoseconds = 123456789

	fmt.Println("Original TimeDto Values")
	PrintOutTimeDtoFields(t1Dto)

	err := t1Dto.NormalizeTimeElements()

	if err != nil {
		fmt.Printf("Error returned by t1Dto.NormalizeTimeElements(). Error='%v' \n",
			err.Error())
		return
	}

	fmt.Println("After Normalize Time Elements")
	PrintOutTimeDtoFields(t1Dto)

	_, err = t1Dto.NormalizeDays()

	if err != nil {
		fmt.Printf("Error returned by t1Dto.NormalizeDays(). Error='%v' \n",
			err.Error())
		return
	}

	fmt.Println("After Normalize Days")
	PrintOutTimeDtoFields(t1Dto)

	utcLoc, _ := time.LoadLocation(dt.TZones.Other.UTC())

	tDate := time.Date(1956, 3, 34, 1, 13, 6, 2002001, utcLoc)

	fmt.Println("tDate: ", tDate.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Println("SUCCESSFUL COMPLETION!")
	/*

	   Original TimeDto Values
	   ========================================
	             TimeDto Printout
	   ========================================
	                      Years:  1955
	                     Months:  15
	                      Weeks:  0
	                   WeekDays:  0
	                   DateDays:  32
	                      Hours:  48
	                    Minutes:  71
	                    Seconds:  125
	               Milliseconds:  1001
	               Microseconds:  1001
	                Nanoseconds:  1001
	   Total SubSec Nanoseconds:  0
	     Total Time Nanoseconds:  0
	   ========================================
	   After Normalize Time Elements
	   ========================================
	             TimeDto Printout
	   ========================================
	                      Years:  1956
	                     Months:  3
	                      Weeks:  4
	                   WeekDays:  6
	                   DateDays:  34
	                      Hours:  1
	                    Minutes:  13
	                    Seconds:  6
	               Milliseconds:  2
	               Microseconds:  2
	                Nanoseconds:  1
	   Total SubSec Nanoseconds:  2002001
	     Total Time Nanoseconds:  4386002002001
	   ========================================
	   After Normalize Days
	   ========================================
	             TimeDto Printout
	   ========================================
	                      Years:  1956
	                     Months:  4
	                      Weeks:  0
	                   WeekDays:  3
	                   DateDays:  3
	                      Hours:  1
	                    Minutes:  13
	                    Seconds:  6
	               Milliseconds:  2
	               Microseconds:  2
	                Nanoseconds:  1
	   Total SubSec Nanoseconds:  2002001
	     Total Time Nanoseconds:  4386002002001
	   ========================================
	   tDate:  1956-04-03 01:13:06.002002001 +0000 UCT

	*/

}

func (mc MainCodeExamples) mainCodeEx016() {
	locUSCentral, err := time.LoadLocation(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	year := 2018
	month := 3
	day := 6
	hour := 20
	minute := 2
	second := 18
	nSecs := 792489279

	t4USCentral := time.Date(year, time.Month(month), day, hour, minute, second, nSecs, locUSCentral)

	t4Dto, err := dt.TimeDto{}.NewTimeComponents(year, month, 0, day, hour, minute,
		second, 0, 0, nSecs)

	if err != nil {
		fmt.Printf("Error returned by t4USCentral TimeDto{}.NewStartEndTimes(). Error='%v'\n", err.Error())
		return
	}

	t4TZoneDef, err := dt.TimeZoneDefinition{}.NewDateTime(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefinition{}.NewStartEndTimes(t4USCentral). Error='%v'", err.Error())
		return
	}

	locTokyo, err := time.LoadLocation(dt.TZones.Asia.Tokyo())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
		return
	}

	t5Tokyo := time.Date(2012, 9, 30, 11, 58, 48, 123456789, locTokyo)

	t5Dto, err := dt.TimeDto{}.NewTimeComponents(2012, 9, 0, 30, 11,
		58, 48, 0, 0, 123456789)

	if err != nil {
		fmt.Printf("Error returned by t5Tokyo TimeDto{}.NewStartEndTimes(). Error='%v'", err.Error())
		return
	}

	t5TZoneDef, err := dt.TimeZoneDefinition{}.NewDateTime(t5Tokyo)

	dTz1, err := dt.DateTzDto{}.NewDateTime(t5Tokyo, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewStartEndTimes(t4USCentral, FmtDateTimeYrMDayFmtStr)\n")
		return
	}

	timeComponents := dTz1.GetTimeComponents()

	if !t5Dto.Equal(timeComponents) {
		fmt.Print("Expected t5Dto == dTz1.timeComponents. It DID NOT!\n")

		fmt.Println("t5Dto")
		PrintOutTimeDtoFields(t5Dto)
		fmt.Println("\n\ndTz1.Time")
		PrintOutTimeDtoFields(timeComponents)
		return
	}

	if !t5TZoneDef.Equal(dTz1.GetTimeZoneDef()) {
		fmt.Print("Expected t5TZoneDef == dTz1.GetTimeZoneDef(). It DID NOT!")
	}

	err = dTz1.SetFromTimeDto(t4Dto, dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from dTz1.SetFromTimeDto(t4Dto, TZones.US.Central()). "+
			"Error='%v'\n", err.Error())
		return
	}

	if !t4USCentral.Equal(dTz1.GetDateTimeValue()) {
		fmt.Printf("Expected dTz1.DateTime='%v'.\n" +
			"Instead, dTz1.DateTime='%v'.\n",
			t4USCentral.Format(dt.FmtDateTimeYrMDayFmtStr),
			dTz1.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	if !t4Dto.Equal(timeComponents) {
		fmt.Print("Expected t4Dto TimeDto == dTz1.Time Time Dto. THEY ARE NOT EQUAL!\n")
		fmt.Println("t4Dto")

		PrintOutTimeDtoFields(t5Dto)
		fmt.Println("\n\ndTz1.Time")
		PrintOutTimeDtoFields(timeComponents)
		return
	}

	if !t4TZoneDef.Equal(dTz1.GetTimeZoneDef()) {
		fmt.Print("Expected t4TZoneDef TimeZoneDef == dTz1.GetTimeZoneDef() TimeZoneDef. " +
			"THEY ARE NOT EQUAL!\n")

		fmt.Println("t4TZoneDef")
		PrintOutTimeZoneDefFields(t4TZoneDef)
		fmt.Println("\n\ndTz1.GetTimeZoneDef()")
		PrintOutTimeZoneDefFields(dTz1.GetTimeZoneDef())

		return
	}

	if year != dTz1.GetTimeComponents().Years {
		fmt.Printf("Error: Expected Years='%v'.\n" +
			"Instead, Years='%v'\n", year, dTz1.GetTimeComponents().Years)
		return
	}

	if month != dTz1.GetTimeComponents().Months {
		fmt.Printf("Error: Expected Months='%v'.\n" +
			"Instead, Months='%v'\n", month, dTz1.GetTimeComponents().Months)
		return
	}

	if day != dTz1.GetTimeComponents().DateDays {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			day, dTz1.GetTimeComponents().DateDays)
		return
	}

	if hour != dTz1.GetTimeComponents().Hours {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			hour, dTz1.GetTimeComponents().Hours)
		return
	}

	if minute != dTz1.GetTimeComponents().Minutes {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			minute, dTz1.GetTimeComponents().Minutes)
		return
	}

	if second != dTz1.GetTimeComponents().Seconds {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			second, dTz1.GetTimeComponents().Seconds)
		return
	}

	if 792 != dTz1.GetTimeComponents().Milliseconds {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			792, dTz1.GetTimeComponents().Milliseconds)
		return
	}

	if 489 != dTz1.GetTimeComponents().Microseconds {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			489, dTz1.GetTimeComponents().Microseconds)
		return
	}

	if 279 != dTz1.GetTimeComponents().Nanoseconds {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			279, dTz1.GetTimeComponents().Nanoseconds)
		return
	}

	if nSecs != dTz1.GetTimeComponents().TotSubSecNanoseconds {
		fmt.Printf("Error: Expected dTz1.GetTimeComponents().TotSubSecNanoseconds='%v'. "+
			"Instead, dTz1.GetTimeComponents().TotSubSecNanoseconds='%v'\n", nSecs, dTz1.GetTimeComponents().TotSubSecNanoseconds)
		return
	}

	totTime := int64(hour) * int64(time.Hour)
	totTime += int64(minute) * int64(time.Minute)
	totTime += int64(second) * int64(time.Second)
	totTime += int64(nSecs)

	if totTime != dTz1.GetTimeComponents().TotTimeNanoseconds {
		fmt.Printf("Error: Expected tDto.TotTimeNanoseconds='%v'. "+
			"Instead, tDto.TotTimeNanoseconds='%v'\n", totTime, dTz1.GetTimeComponents().TotTimeNanoseconds)
		return
	}

	fmt.Println("SUCCESSFUL COMPLETION!!!")
}

func (mc MainCodeExamples) mainCodeEx015() {

	// t1str :="2017-04-30 22:58:32.515539300 -0500 CDT"
	// t1, err := time.Parse(FmtDateTimeYrMDayFmtStr, t1str)

	dTzDto, err := dt.DateTzDto{}.NewDateTimeElements(2017, 04, 30, 22, 58, 32, 515539300, dt.TZones.US.Central(), dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned from DateTzDto{}.NewDateTimeElements(year, month, day,...). "+
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Println("dTzDto.DateTime: ", dTzDto.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))

	tDto, err := dt.TimeDto{}.NewFromDateTime(dTzDto.GetDateTimeValue())

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDto{}.NewFromDateTime(dTzDto.DateTime)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dt2, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by tDto.GetDateTime(TZones.US.Central()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !dt2.Equal(dTzDto.GetDateTimeValue()) {
		fmt.Printf("Error: Expected dTzDto.DateTime='%v'.\n" +
			"It did NOT! dTzDto.DateTime='%v'\n",
			dt2.Format(dt.FmtDateTimeYrMDayFmtStr),
			dTzDto.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	fmt.Println("Success!")
}

func (mc MainCodeExamples) mainCodeEx014() {
	t1str := "2014-02-15 19:54:30.000000000 -0600 CST"
	t2str := "2017-04-30 22:58:32.000000000 -0500 CDT"
	t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)

	//t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t1OutStr := t1.Format(dt.FmtDateTimeYrMDayFmtStr)
	//t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)
	t2OutStr := t2.Format(dt.FmtDateTimeYrMDayFmtStr)
	t12Dur := t2.Sub(t1)

	timeDto := dt.TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err := timeDto.NormalizeTimeElements()

	if err != nil {
		fmt.Printf("Error returned by timeDto.NormalizeTimeElements(). "+
			"Error='%v' ", err.Error())
		return
	}

	dur, err := dt.DurationTriad{}.NewEndTimeMinusTimeDtoTz(
		dt.TCalcMode.LocalTimeZone(),
		t2,
		timeDto,
		dt.TZones.US.Central(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto). "+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.GetTypeStartDateTime().Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error- Expected Start Time %v.\n" +
			"Instead, got %v.\n",
			t1OutStr, dur.BaseTime.GetTypeStartDateTime().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error- Expected End Time %v. Instead, got %v.\n",
			t2OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	if t12Dur != dur.BaseTime.TimeDuration {
		fmt.Printf("Error- Expected Time Duration %v. Instead, got %v\n",
			t12Dur, dur.BaseTime.TimeDuration)
		return
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YrMthDay: %v. Instead, got %v\n", expected, outStr)
		return
	}

	fmt.Println("Successful Completion!")
}

func (mc MainCodeExamples) mainCodeEx012() {
	// 101095442000000000
	t1str := "2014-02-15 19:54:30.000000000 -0600 CST"
	t2str := "2017-04-30 22:58:32.000000000 -0500 CDT"
	t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)
	eDur := t2.Sub(t1)

	timeDto := dt.TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	timeDto.ConvertToNegativeValues()

	dx1 := t2.AddDate(timeDto.Years, timeDto.Months, 0)

	dur := (int64(timeDto.Weeks*7) + int64(timeDto.WeekDays)) * dt.DayNanoSeconds
	dur += int64(timeDto.Hours) * dt.HourNanoSeconds
	dur += int64(timeDto.Minutes) * dt.MinuteNanoSeconds
	dur += int64(timeDto.Seconds) * dt.SecondNanoseconds

	dx2 := dx1.Add(time.Duration(dur))

	fmt.Println("Expected Start Date Time: ", t1.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("  Actual Start Date Time: ", dx2.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("       Expected Duration: ", int64(eDur))
	fmt.Println("          ActualDuration:  101095442000000000")

}

func (mc MainCodeExamples) mainCodeEx011() {

	t1str := "2014-02-15 19:54:30.000000000 -0600 CST"
	t2str := "2017-04-30 22:58:32.000000000 -0500 CDT"
	t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)

	tDto, err := dt.TimeDurationDto{}.NewStartEndTimesCalcTz(t1, t2,
		dt.TDurCalcType(0).StdYearMth(), dt.TZones.US.Central(), dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesCalcTz() "+
			"Error='%v'", err.Error())
		return
	}

	PrintTimeDurationDto(tDto)

}

func (mc MainCodeExamples) mainCodeEx010() {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := dt.TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err := timeDto.NormalizeTimeElements()

	if err != nil {
		fmt.Printf("Error returned by timeDto.NormalizeTimeElements(). "+
			"Error='%v' ", err.Error())
		return
	}

	dur, err := dt.TimeDurationDto{}.NewEndTimeMinusTimeDto(
		dt.TCalcMode.LocalTimeZone(),
		t2,
		timeDto,
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto). Error='%v'\n", err.Error())
		return
	}

	startTimeTz := dur.GetTypeStartDateTimeTz()

	fmt.Println("Expected Start Date Time: ", t1OutStr)
	fmt.Println("  Actual Start Date Time: ", startTimeTz.String())
	fmt.Println("-----------------------------------------")
	fmt.Println("  Expected End Date Time: ", t2OutStr)
	fmt.Println("    Actual End Date Time: ", dur.EndTimeDateTz.String())
	fmt.Println("-----------------------------------------")
	fmt.Println("       Expected Duration: ", t12Dur)
	fmt.Println("         Actual Duration: ", dur.TimeDuration.String())
}

func (mc MainCodeExamples) mainCodeEx009() {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto, err :=
		dt.TimeDurationDto{}.NewStartEndTimesCalcTz(t2, t1,
			dt.TDurCalcType(0).StdYearMth(), dt.TZones.US.Central(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesCalcTz(). "+
			" Error='%v'\n", err.Error())
		return
	}

	fmt.Println("TimeDurationDto")
	PrintTimeDurationDto(tDto)

	durT, err := dt.DurationTriad{}.NewStartEndTimesTz(t2, t1, dt.TZones.US.Central(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesCalcTz(). "+
			" Error='%v'\n", err.Error())
		return
	}

	fmt.Println("DurationTriad BaseTimeDto")
	PrintTimeDurationDto(durT.BaseTime)
}

func (mc MainCodeExamples) mainCodeEx008() {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := dt.DurationTriad{}

	err := du.SetStartEndTimesTz(t2, t1, dt.TZones.US.Central(), dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by du.SetStartEndTimesTz(t2, t1, dt.TZones.US.Central(), "+
			"dt.FmtDateTimeYrMDayFmtStr). Error='%v' ", err.Error())
		return
	}

	expected := "0-Hours 0-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := du.BaseTime.GetYearMthDaysTimeAbbrvStr()

	startTimeTz := du.BaseTime.GetTypeStartDateTimeTz()

	fmt.Println("  Expected: ", expected)
	fmt.Println("    Actual: ", dOut)
	fmt.Println("Start Time: ", startTimeTz.String())
	fmt.Println("  End Time: ", du.BaseTime.EndTimeDateTz.String())

}

func (mc MainCodeExamples) mainCodeEx007() {

	mthTest := int(time.Month(0))

	fmt.Println("===============================")
	fmt.Println("       Month Zero Test")
	fmt.Println("===============================")
	fmt.Println("int (time.Month(0))= ", mthTest)

	/* Result = Sending zero to time.Month(0) yields a
	a zero value. NOT GOOD! Best to use '1' in
	place of zero month number.

		===============================
					 Month Zero Test
		===============================
		int (time.Month(0))=  0

	*/

}

func (mc MainCodeExamples) mainCodeEx006() {

	locUTC, _ := time.LoadLocation(dt.TZones.Other.UTC())

	fmt.Println()
	fmt.Println("2018-00")
	fmt.Println()

	tDateTime := time.Date(2018, 0, 0, 0, 0, 0, 0, locUTC)

	PrintDateTime(tDateTime, dt.FmtDateTimeYrMDayFmtStr)
	/*		Result - Don't Use a Zero Month
			2018-00
			----------------------------------
								 Date Time
			----------------------------------
			Date Time:  2017-11-30 00:00:00.000000000 +0000 UCT
			The integer month is:  11
			The integer day is: 30
			The integer year is: 2017
			The integer hour is: 0
			The integer minute is: 0
			The integer second is: 0
			The integer nanosecond is 0
	*/

	fmt.Println()
	fmt.Println("2018-01")
	fmt.Println()

	t2 := time.Date(2018, 1, 0, 0, 0, 0, 0, locUTC)

	PrintDateTime(t2, dt.FmtDateTimeYrMDayFmtStr)
	/* Result - Best Approach - Use 1 as month number instead of zero month number.
	Also - Use Zero Days. Convert days to duration and add the duration.
	2018-01

	----------------------------------
						 Date Time
	----------------------------------
	Date Time:  2017-12-31 00:00:00.000000000 +0000 UCT
	The integer month is:  12
	The integer day is: 31
	The integer year is: 2017
	The integer hour is: 0
	The integer minute is: 0
	The integer second is: 0
	The integer nanosecond is 0
	*/

	fmt.Println()
	fmt.Println("Add 1 Day")
	fmt.Println()

	dur := int64(24) * dt.HourNanoSeconds

	t3 := t2.Add(time.Duration(dur))

	PrintDateTime(t3, dt.FmtDateTimeYrMDayFmtStr)
	/*
		Add 1 Day to	2017-12-31 00:00:00.000000000 +0000 UCT
		Gives desired result

		----------------------------------
							 Date Time
		----------------------------------
		Date Time:  2018-01-01 00:00:00.000000000 +0000 UCT
		The integer month is:  1
		The integer day is: 1
		The integer year is: 2018
		The integer hour is: 0
		The integer minute is: 0
		The integer second is: 0
		The integer nanosecond is 0
	*/

}

func (mc MainCodeExamples) mainCodeEx005() {

	locUTC, _ := time.LoadLocation(dt.TZones.Other.UTC())

	tDateTime := time.Date(2018, 2, 0, 0, 0, 0, 0, locUTC)

	PrintDateTime(tDateTime, dt.FmtDateTimeYrMDayFmtStr)

	fmt.Println()
	fmt.Println("Adding 3-days")
	fmt.Println()

	dur := int64(3) * dt.DayNanoSeconds
	t2 := tDateTime.Add(time.Duration(dur))

	PrintDateTime(t2, dt.FmtDateTimeYrMDayFmtStr)

	expectedDt := time.Date(2018, 2, 3, 0, 0, 0, 0, locUTC)

	fmt.Println()
	fmt.Println("Complete Date 2018-02-03")
	fmt.Println()

	PrintDateTime(expectedDt, dt.FmtDateTimeYrMDayFmtStr)

}

func (mc MainCodeExamples) mainCodeEx004() {
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	locUSCentral, _ := time.LoadLocation(dt.TZones.US.Central())
	locUSPacific, _ := time.LoadLocation(dt.TZones.US.Pacific())
	locParis, _ := time.LoadLocation(dt.TZones.Europe.Paris())
	locCairo, _ := time.LoadLocation(dt.TZones.Africa.Cairo())
	locMoscow, _ := time.LoadLocation(dt.TZones.Europe.Moscow())
	locTokyo, _ := time.LoadLocation(dt.TZones.Asia.Tokyo())

	t1USCentral := time.Date(1948, time.Month(9), 7, 4, 32, 16, 8185431, locUSCentral)
	t1USPacific := t1USCentral.In(locUSPacific)
	t1EuropeParis := t1USPacific.In(locParis)
	t1AfricaCairo := t1EuropeParis.In(locCairo)
	t1EuropeMoscow := t1AfricaCairo.In(locMoscow)
	t1AsiaTokyo := t1EuropeMoscow.In(locTokyo)
	t1bUSCentral := t1AsiaTokyo.In(locUSCentral)

	fmt.Println("t1USCentral: ", t1USCentral.Format(fmtStr))
	fmt.Println("t1USPacific: ", t1USPacific.Format(fmtStr))
	fmt.Println("t1EuropeParis: ", t1EuropeParis.Format(fmtStr))
	fmt.Println("t1AfricaCairo: ", t1AfricaCairo.Format(fmtStr))
	fmt.Println("t1EuropeMoscow: ", t1EuropeMoscow.Format(fmtStr))
	fmt.Println("t1AsiaTokyo: ", t1AsiaTokyo.Format(fmtStr))
	fmt.Println("t1bUSCentral: ", t1bUSCentral.Format(fmtStr))

}

func (mc MainCodeExamples) mainCodeEx003() {
	loc, _ := time.LoadLocation(dt.TZones.US.Central())

	t1 := time.Date(2014, time.Month(15), 67, 19, 54, 30, 158712300, loc)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	addYear := 0
	addMonth := 15
	addDay := 64
	addHours := 0
	addMinutes := 0
	addSeconds := 0
	addMilliSeconds := 0
	addMicroSeconds := 0
	addNanoSeconds := 0

	var totDuration int64

	t2 := t1.AddDate(addYear, addMonth, addDay)

	totDuration = int64(addHours) * int64(time.Hour)
	totDuration += int64(addMinutes) * int64(time.Minute)
	totDuration += int64(addSeconds) * int64(time.Second)
	totDuration += int64(addMilliSeconds) * int64(time.Millisecond)
	totDuration += int64(addMicroSeconds) * int64(time.Microsecond)
	totDuration += int64(addNanoSeconds)

	t3 := t2.Add(time.Duration(totDuration))

	fmt.Println("t1: ", t1.Format(fmtstr))
	fmt.Println("t2: ", t2.Format(fmtstr))
	fmt.Println("t2: ", t3.Format(fmtstr))

}

func (mc MainCodeExamples) mainCodeEx002() {

	tDto, err := dt.TimeDto{}.NewTimeComponents(0, 0, -8, 0, 0, 0, 0, 0, 0, 0)

	if err != nil {
		fmt.Printf("Error returned from TimeDto{}.NewStartEndTimes(0, 0, -8, 0, 0, 0, 0, 0, 0, 0 ) Error='%v' \n", err.Error())
	}

	PrintOutTimeDtoFields(tDto)

}

func (mc MainCodeExamples) mainCodeEx001() {

	loc, _ := time.LoadLocation(dt.TZones.US.Central())
	t1 := time.Date(2014, time.Month(2), 15, 19, 54, 30, 158712300, loc)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	tDto, err := dt.TimeDto{}.NewTimeComponents(2014, 2, 0, 15, 19, 54, 30, 0, 0, 158712300)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDto{}.NewStartEndTimes(year, month, ...). Error=%v \n", err.Error())
		return
	}

	t2, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by tDto.GetDateTime(dt.TZones.US.Central()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Println("t1: ", t1.Format(fmtstr))
	fmt.Println("t2: ", t2.Format(fmtstr))

}

func (mc MainCodeExamples) mainPrintHdr(textToPrint string, repeatStr string) {


	lenTextToPrint := len(textToPrint)

	lenExtra := 5

	lenBar := (lenExtra * 2) + lenTextToPrint

	blankMargin := strings.Repeat(" ", lenExtra)

	title := blankMargin + textToPrint + blankMargin
	bar := strings.Repeat(repeatStr, lenBar)

	fmt.Println(bar)
	fmt.Println(title)
	fmt.Println(bar)
	fmt.Println()

}


