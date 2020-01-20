package main

import (
	"fmt"
	dt "github.com/MikeAustin71/datetimeopsgo/datetime"
	ex "github.com/MikeAustin71/datetimeopsgo/datetimeexamples"
	"strings"
	"time"
)

func main() {

	mainTest{}.mainTest062()

}

type mainTest struct {
	input  string
	output string
}

func (mt mainTest) mainTest062() {

	ePrefix := "mainTest062() TimeZoneDefinition"

	mt.mainPrintHdr(ePrefix , "-")

	utcOffset := "2020-01-19 04:21:18 +0700 +07"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	expectedOriginalTz := "Etc/GMT-7"
	expectedConvertibleTz := dt.TZones.Asia.Ho_Chi_Minh()

	utcOffsetTime, err := time.Parse(fmtStr, utcOffset)

	if err != nil {
		fmt.Printf("Received error from time parse utcOffset: %v\n",
			err.Error())
		return
	}

	var tzDef dt.TimeZoneDefinition

	tzDef, err = dt.TimeZoneDefinition{}.New(utcOffsetTime)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefinition{}.New(utcOffsetTime)\n" +
			"utcOffsetTime= '%v'\n" +
			"Error='%v'\n",
			utcOffsetTime.Format(fmtStr), err.Error())
		return
	}

	originalTzSpec := tzDef.GetOriginalTimeZone()
	
	ex.PrintOutDateTimeTimeZoneFields(originalTzSpec.GetReferenceDateTime(), "Original TZ Date Time")
	ex.PrintOutTimeZoneSpecFields(originalTzSpec, "Original Time Zone")


	convertibleTzSpec := tzDef.GetConvertibleTimeZone()
	
	ex.PrintOutDateTimeTimeZoneFields(convertibleTzSpec.GetReferenceDateTime(), "Convertible TZ Date Time")
	ex.PrintOutTimeZoneSpecFields(convertibleTzSpec, "Convertible Time Zone")

	
	
	

	actualOriginalTz := tzDef.GetOriginalTimeZoneName()

	actualConvertibleTz := tzDef.GetConvertibleTimeZoneName()


	if expectedOriginalTz != actualOriginalTz {
		fmt.Printf("Error: Expected actualOriginalTz='%v'.\n" +
			"Instead, actualOriginalTz='%v'\n",
			expectedOriginalTz, actualOriginalTz)
		return
	}

	if expectedConvertibleTz != actualConvertibleTz {
		fmt.Printf("Error: Expected actualConvertibleTz='%v'.\n" +
			"Instead, actualConvertibleTz='%v'\n",
			expectedConvertibleTz, actualOriginalTz)
		return
	}

	mt.mainPrintHdr("Successful Completion" , "=")

}

func (mt mainTest) mainTest061() {

	ePrefix := "mainTest061() TimeZoneDto.ConvertTz()"

	mt.mainPrintHdr(ePrefix , "-")

	utcTime := "2017-04-30 00:54:30 +0000 UTC"
	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	mountainTime := "2017-04-29 18:54:30 -0600 MDT"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	ianaCentralTz := "America/Chicago"
	ianaMountainTz := "America/Denver"
	tPacificIn, err := time.Parse(fmtstr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	tzu := dt.TimeZoneDto{}
	tzuCentral, err := tzu.ConvertTz(
		tPacificIn,
		ianaCentralTz,
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("#1 Error from ianaCentralTz TimeZoneDto.ConvertTz().\n" +
			"Error: %v\n", err.Error())
		return
	}

	centralTOut := tzuCentral.TimeOut.GetDateTimeValue().Format(fmtstr)

	if centralTime != centralTOut {
		fmt.Printf("Expected tzuCentral.TimeOut = '%v'.\n" +
			"Instead, tzuCentral.TimeOut = '%v'.\n",
			centralTime, centralTOut)
		return
	}

	tzuMountain, err := tzu.ConvertTz(
		tzuCentral.TimeOut.GetDateTimeValue(),
		ianaMountainTz, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("#2 Error from tzuMountain TimeZoneDto.ConvertTz().\n" +
			"Error: %v\n", err.Error())
		return
	}

	mountainTOut := tzuMountain.TimeOut.GetDateTimeValue().Format(fmtstr)

	if mountainTime != mountainTOut {
		fmt.Printf("Expected tzuMountain.TimeOut= '%v'.\n" +
			"Instead, tzuMountain.TimeOut= '%v'.\n",
			mountainTime, mountainTOut)
		return
	}

	tzuPacific, err := tzu.ConvertTz(
		tzuMountain.TimeOut.GetDateTimeValue(),
		ianaPacificTz, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error #3 from tzuMountain TimeZoneDto.ConvertTz().\n" +
			"Error: %v\n", err.Error())
		return
	}

	pacificTOut := tzuPacific.TimeOut.GetDateTimeValue().Format(fmtstr)

	if pacificTime != pacificTOut {

		fmt.Printf("Expected tzuPacific.TimeOut= '%v'.\n" +
			"Instead, tzuPacific.TimeOut= '%v'.\n",
			pacificTime, pacificTOut)
		return
	}

	exTOutLoc := "America/Los_Angeles"

	if exTOutLoc != tzuPacific.TimeOut.GetTimeZoneName() {
		fmt.Printf("Expected tzu.TimeOutLoc='%v'.\n" +
			"Instead tzu.TimeOutLoc='%v'\n" +
			"tzuPacific.TimeOut='%v'\n",
			exTOutLoc,
			tzuPacific.TimeOut.GetTimeZoneName(),
			tzuPacific.TimeOut.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	pacificUtcOut := tzuPacific.TimeUTC.GetDateTimeValue().Format(fmtstr)

	if utcTime != pacificUtcOut {
		fmt.Printf("Expected tzuPacific.TimeUTC= '%v'\n" +
			"Instead, tzuPacific.TimeUTC= '%v'\n",
			utcTime, pacificUtcOut)
		return
	}

	centralUtcOut := tzuCentral.TimeUTC.GetDateTimeValue().Format(fmtstr)

	if utcTime != centralUtcOut {
		fmt.Printf("Expected tzuCentral.TimeUTC= '%v'\n" +
			"Instead, tzuCentral.TimeUTC= '%v'\n",
			utcTime, pacificUtcOut)
		return
	}

	mountainUtcOut := tzuMountain.TimeUTC.GetDateTimeValue().Format(fmtstr)

	if utcTime != mountainUtcOut {
		fmt.Printf("Expected tzuMountain.TimeUTC= '%v'\n" +
			"Instead, tzuMountain.TimeUTC= '%v'\n",
			utcTime, pacificUtcOut)
		return
	}

	fmt.Println()
	mt.mainPrintHdr("Successful Completion" , "=")
}


func (mt mainTest) mainTest060() {

	ePrefix := "mainTest060()"

	mt.mainPrintHdr(ePrefix , "-")
	mt.mainPrintHdr("TimeZoneDto" , "-")

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtStr, t1str)
	t1OutStr := t1.Format(fmtStr)
	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtStr)

	if err != nil {
		fmt.Printf("Error returned from TimeZoneDto{}.New(t1, TzUsPacific ). Error='%v'", err.Error())
		return
	}

	fmt.Println("t1OutStr: ", t1OutStr)

	ex.PrintOutDateTimeTimeZoneFields(
		tzu1.TimeIn.GetDateTimeValue(), "Initial tzu1.TimeIn")

	t2, _ := time.Parse(fmtStr, t2str)
	t2OutStr := t2.Format(fmtStr)

	t12Dur := t2.Sub(t1)

	tdurDto, err := dt.TimeDurationDto{}.NewStartEndTimesCalcTz(
		t1,
		t2,
		dt.TDurCalcType(0).StdYearMth(),
		dt.TZones.US.Central(),
		fmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}.NewStartEndTimesCalcTz()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddTimeDurationDto(tdurDto)

	if err != nil {
		fmt.Printf("Error returned by tzu2.AddTimeDurationDto(tdurDto).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	tzu1OutStr := tzu1.TimeIn.GetDateTimeValue().Format(fmtStr)

	if t1OutStr != tzu1OutStr {
		fmt.Printf("Error: Expected Time1 TimeIn='%v'.\n" +
			"Instead Time1 TimeIn='%v'\n",
			t1OutStr, tzu1OutStr)
		return
	}

	tzu2OutStr := tzu2.TimeIn.GetDateTimeValue().Format(fmtStr)

	if t2OutStr != tzu2OutStr {
		fmt.Printf("Error: Expected after duration tzu2TimeIn='%v'.\n" +
			"Instead, tzu2TimeIn='%v'\n",
			t2OutStr, tzu2OutStr)
		return
	}

	actualDur := tzu2.TimeIn.Sub(tzu1.TimeIn)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeIn.Sub(tzu1.TimeIn)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeOut.Sub(tzu1.TimeOut)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeOut.Sub(tzu1.TimeOut)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
	}

	actualDur = tzu2.TimeUTC.Sub(tzu1.TimeUTC)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeUTC.Sub(tzu1.TimeUTC)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeLocal.Sub(tzu1.TimeLocal)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeLocal.Sub(tzu1.TimeLocal)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualTimeOutLoc := tzu1.TimeOut.GetTimeZoneName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu1.TimeOutLoc='%v'.\n" +
			"Instead, tzu1.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	actualTimeOutLoc = tzu2.TimeOut.GetTimeZoneName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu2.TimeOutLoc.String()='%v'.\n" +
			"Instead, tzu2.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	mt.mainPrintHdr("Successful Completion" , "=")
}

func (mt mainTest) mainTest059() {
	ePrefix := "mainTest059()"

	mt.mainPrintHdr(ePrefix , "-")
	mt.mainPrintHdr("ConvertTz" , "-")

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	// Invalid Target Iana Time Zone
	invalidTz := "XUZ Time Zone"
	tIn, err := time.Parse(fmtstr, tstr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtstr, tstr)\n" +
			"tstr='%v'\n" +
			"Error='%v'\n", tstr, err.Error())
		return
	}

	tzu := dt.TimeZoneDto{}

	_, err = tzu.ConvertTz(tIn, invalidTz, fmtstr)

	if err == nil {
		fmt.Printf("ConvertTz() failed to detect INVALID Target Time Zone.\n" +
			"err=='nil'\n" +
			"invalidTz='%v'\n", invalidTz)
	}

	return
}

func (mt mainTest) mainTest058() {
	ePrefix := "mainTest058()"

	mt.mainPrintHdr(ePrefix , "-")
	mt.mainPrintHdr("ConvertTzAbbreviationToTimeZone" , "-")

	tzAbbrv := "+10"
	testUtcOffset := "+1000"

	tzMech := dt.TimeZoneMechanics{}

	dateTime := time.Now().UTC()

	var staticTimeZone dt.TimeZoneSpecification
	var err error

	staticTimeZone,
	err = tzMech.ConvertUtcAbbrvToStaticTz(
		dateTime,
		dt.TzConvertType.Relative(),
		"Original Time Zone",
		testUtcOffset,
		ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}

	fmt.Println("staticTimeZone: ", staticTimeZone.GetLocationName())
	fmt.Println()
	ex.PrintOutDateTimeTimeZoneFields(staticTimeZone.GetReferenceDateTime(), "staticTimeZone")
	ex.PrintOutTimeZoneSpecFields(staticTimeZone, "staticTimeZone")

	var tzSpec dt.TimeZoneSpecification

	tzSpec,
	err =
	tzMech.ConvertTzAbbreviationToTimeZone(
		dateTime,
		dt.TimeZoneConversionType(0).Relative(),
		tzAbbrv+testUtcOffset,
		"",
		ePrefix)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by tzMech.ConvertTzAbbreviationToTimeZone()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(tzSpec.GetReferenceDateTime(), "tzSpec")
	ex.PrintOutTimeZoneSpecFields(tzSpec, "tzSpec")


	mt.mainPrintHdr("SUCCESS" , "!")

}

func (mt mainTest) mainTest057() {

	ePrefix := "mainTest057()"

	mt.mainPrintHdr(ePrefix , "-")

	dtUtil := dt.DTimeUtility{}

	locPtr, err := dtUtil.LoadTzLocation(dt.TZones.Asia.Vladivostok(), ePrefix)

	if err != nil {
		fmt.Printf("Error='%v'\n", err.Error())
		return
	}

	dateTime := time.Date(
		2019,
		time.Month(6),
		15,
		11,
		23,
		0,
		0,
		locPtr)

	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	tzDef, err := dt.TimeZoneDefinition{}.NewFromTimeZoneName(
		dateTime, dt.TZones.Asia.Vladivostok(), dt.TzConvertType.Absolute())

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.New(dateTime, London Time)\n" +
			"dateTime='%v'\nError='%v'\n",
			dateTime.Format(fmtStr), err.Error())
		return
	}

	ex.PrintOutTimeZoneDefFields(tzDef)


	// t1EdtStr :=  "06/20/2019 09:58:32 -0400 EDT"
	// t1EstStr :=  "12/20/2019 09:58:32 -0500 EST"
	// t2CdtStr :=  "06/20/2019 09:58:32 -0500 CDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0600 CST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0600 MDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0700 MST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0700 PDT"
	// t2PstStr :=  "12/20/2019 09:58:32 -0800 PST"

	timeStr := dateTime.Format(fmtStr)

	t2, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(t2, "t2 Parse Result")

	tzDef, err = dt.TimeZoneDefinition{}.New(t2)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.New(t1)\n" +
			"t2='%v'\nError='%v'\n", t2.Format(fmtStr), err.Error())
		return
	}

	ex.PrintOutTimeZoneDefFields(tzDef)

	mt.mainPrintHdr("SUCCESS" , "!")
}

func (mt mainTest) mainTest056() {

	ePrefix := "mainTest056()"

	mt.mainPrintHdr(ePrefix , "-")

	dtUtil := dt.DTimeUtility{}
	locUSCentral, err :=
		dtUtil.LoadTzLocation(dt.TZones.US.Central(), ePrefix)

	if err != nil {
		fmt.Printf("Error='%v'\n", err.Error())
		return
	}

	locTokyo, err := dtUtil.LoadTzLocation(dt.TZones.Asia.Tokyo(), ePrefix)

	if err != nil {
		fmt.Printf("Error='%v'", err.Error())
		return
	}

	t4USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)

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
		fmt.Printf("Error: Expected DateTime='%v'. Instead DateTime='%v'",
			t4USCentral.Format(fmtstr), dTz.GetDateTimeValue().Format(fmtstr))
	}

	eTimeZoneDef, err := dt.TimeZoneDefinition{}.New(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefinition{}.New(t4USCentral)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	areEqual := eTimeZoneDef.Equal(dTz.GetTimeZoneDef())

	if ! areEqual {
		fmt.Printf("Expected dTz.GetTimeZoneDef().LocationName='%v'.\n"+
			"Instead, dTz.GetTimeZoneDef().LocationName='%v'\n",
			eTimeZoneDef.GetOriginalLocationName(), dTz.GetTimeZoneName())
	}

	tDto, err := dt.TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral)\n"+
			"t4USCentral='%v'\nError='%v'\n",
			t4USCentral.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	expectedDt, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from tDto.GetDateTime(TZones.US.Central()). "+
			"Error='%v'", err.Error())
	}

	timeComponents := dTz.GetTimeComponents()

	actualDt, err := timeComponents.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from dTz.GetTimeComponents().GetDateTime(TZones.US.Central()).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !tDto.Equal(dTz.GetTimeComponents()) {
		fmt.Printf("Expected dTz.Time (TimeDto) == '%v' Instead, dTz.Time (TimeDto) == '%v'",
			expectedDt.Format(dt.FmtDateTimeYrMDayFmtStr),
			actualDt.Format(dt.FmtDateTimeYrMDayFmtStr))
	}

	if dt.FmtDateTimeYrMDayFmtStr != dTz.GetDateTimeFmt() {
		fmt.Printf("Expected dTz.GetDateTimeFmt()='%v' Instead, dTz.GetDateTimeFmt()='%v' ",
			dt.FmtDateTimeYrMDayFmtStr, dTz.GetDateTimeFmt())
	}

}

func (mt mainTest) mainTest055() {

	ePrefix := "mainTest055()"

	mt.mainPrintHdr(ePrefix , "-")

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

	tzDef, err := dt.TimeZoneDefinition{}.New(t1)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.New(t1)\n" +
			"t1='%v'\nError='%v'\n", t1.Format(fmtStr), err.Error())
		return
	}

	ex.PrintOutTimeZoneDefFields(tzDef)

}



func (mt mainTest) mainTest054() {

	ePrefix := "mainTest054()"

	mt.mainPrintHdr(ePrefix , "-")

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

	ex.PrintOutDateTimeTimeZoneFields(t1, "t1 Initial Date Time")

	t2 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		t1.Location())

	ex.PrintOutDateTimeTimeZoneFields(t2, "t2 Date Time")

	pacificTz := dt.TZones.America.Los_Angeles()

	pacificTzPtr, err := time.LoadLocation(pacificTz)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(pacificTz)\n" +
			"pacificTz='%v'\nError='%v'\n", pacificTz, err.Error())
		return
	}

	t3 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		pacificTzPtr)

	ex.PrintOutDateTimeTimeZoneFields(t3, "t3 Date Time")

}


func (mt mainTest) mainTest053() {

	ePrefix := "mainTest053()"

	mt.mainPrintHdr(ePrefix , "=")

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _ := dt.TimeZoneDto{}.New(tIn, ianaPacificTz, fmtstr)

	fmt.Println(" Time In: ", tzu.TimeIn.GetDateTimeValue().Format(fmtstr))
	fmt.Println("Time Out: ", tzu.TimeOut.GetDateTimeValue().Format(fmtstr))

	expectedZone := "PDT"

	actualZone := tzu.TimeOut.GetTimeZoneAbbreviation()

	if expectedZone != actualZone {
		fmt.Printf("Expected Zone Out='%v'.\n" +
			"Instead, actual Zone Out='%v'\n", expectedZone, actualZone)
		return
	}

	mt.mainPrintHdr("SUCCESS" , "!!!")


}


func (mt mainTest) mainTest052() {

	ePrefix := "mainTest052()"

	mt.mainPrintHdr(ePrefix , "=")

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned from TimeZoneDto{}.New(t1, TzUsPacific ).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tzu1.IsValid()

	if err != nil {
		fmt.Printf("'tzu1' is INVALID!\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("\nPassed IsValid() Check #1!\n")

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	t12Dur := t2.Sub(t1)

	tdurDto, err := dt.TimeDurationDto{}.NewStartEndTimesCalcTz(
		t1,
		t2,
		dt.TDurCalcType(0).StdYearMth(),
		dt.TZones.US.Central(),
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}.NewStartEndTimesCalcTz()\n"+
			"Error='%v' ", err.Error())
		return
	}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddTimeDurationDto(tdurDto)
// Error expected here!
	if err != nil {
		fmt.Printf("Error returned by tzu2.AddTimeDurationDto(tdurDto).\n"+
			"Error='%v' ", err.Error())
		return
	}

	tzu1OutStr := tzu1.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t1OutStr != tzu1OutStr {
		fmt.Printf("Error: Expected Time1 TimeIn='%v'.\n" +
			"Instead Time1 TimeIn='%v'\n",
			t1OutStr, tzu1OutStr)
		return
	}

	tzu2OutStr := tzu2.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t2OutStr != tzu2OutStr {
		fmt.Printf("Error: Expected after duration tzu2TimeIn='%v'.\n" +
			"Instead, tzu2TimeIn='%v'\n",
			t2OutStr, tzu2OutStr)
		return
	}

	actualDur := tzu2.TimeIn.Sub(tzu1.TimeIn)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeIn.Sub(tzu1.TimeIn)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeOut.Sub(tzu1.TimeOut)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeOut.Sub(tzu1.TimeOut)='%v'.\n" +
			"Instead, duration='%v'\n", t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeUTC.Sub(tzu1.TimeUTC)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeUTC.Sub(tzu1.TimeUTC)='%v'.\n" +
			"Instead, duration='%v'\n", t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeLocal.Sub(tzu1.TimeLocal)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeLocal.Sub(tzu1.TimeLocal)='%v'.\n" +
			"Instead, duration='%v'\n", t12Dur, actualDur)
		return
	}

	actualTimeOutLoc := tzu1.TimeOut.GetTimeZoneName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu1.TimeOutLoc='%v'.\n" +
			"Instead, tzu1.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	actualTimeOutLoc = tzu2.TimeOut.GetTimeZoneName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu2.TimeOutLoc.String()='%v'.\n" +
			"Instead, tzu2.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	mt.mainPrintHdr("SUCCESS" , "!")

}

func (mt mainTest) mainTest051() {
	ePrefix := "mainTest051()"

	mt.mainPrintHdr(ePrefix , "-")

	timeZone := dt.TZones.Local()
	// t2PdtStr :=  "06/20/2019 09:58:32.000000000 -0700 PDT"
	utcLoc, err := time.LoadLocation(timeZone)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.UTC())\n" +
			"Error='%v'\n", err.Error())
	}

	txUtc := time.Date(
		2019,
		time.Month(12),
		15,
		11,
		0,
		0,
		0,
		utcLoc)

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"
	timeStr := txUtc.Format(fmtStr)
	fmt.Println()
	fmt.Println("#1   timeStr: ", timeStr)
	fmt.Println("#1 time zone: ", timeZone)
	fmt.Println()

		t1, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(t1, "t1 Initial Date Time")

	_, err = time.LoadLocation(t1.Location().String())

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(t1.Location().String())\n" +
			"t1.Location().String()='%v'\n" +
			"Error='%v'\n", t1.Location().String(), err.Error())
	}

	dtUtil := dt.DTimeUtility{}

	var tzSpec dt.TimeZoneSpecification

	tzSpec,
	err = dtUtil.GetConvertibleTimeZoneFromDateTime(
		t1,
		dt.TzConvertType.Absolute(),
		ePrefix)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by dtUtil.GetConvertibleTimeZoneFromDateTime(t1, ePrefix)\n" +
			"t1='%v'\n" +
			"Error='%v'\n", t1.Format(fmtStr), err.Error())
		return
	}

	fmt.Println()
	fmt.Printf("        ianaTimeZoneName: %v\n", tzSpec.GetLocationName())
	fmt.Printf("         ianaLocationPtr: %v\n", tzSpec.GetLocationPointer().String())
	fmt.Printf("                 tzClass: %v\n\n", tzSpec.GetTimeZoneClass().String())

	fmt.Printf(ePrefix +
		"\nSuccess!\n" +
		"time zone pointer name = '%v'\n",t1.Location().String())

}

func (mt mainTest) mainTest050() {
/*
	 =========================================
	      Time Zone Fields From time.Time
	 =========================================
	    t1 Initial Date Time:  12/30/2019 09:00:00.000000000 +0700 +07
	               Zone Name:  +07
	     Abbreviation Lookup:  +07+0700
	               Zone Sign:  +
	       Zone Offset Hours:  7
	     Zone Offset Minutes:  0
	     Zone Offset Seconds:  0
	    Total Offset Seconds:  25200
	              UTC Offset:  UTC+0700
	        Location Pointer:  Asia/Ho_Chi_Minh
	           Location Name:  Asia/Ho_Chi_Minh
	 =========================================
 */
	ePrefix := "mainTest050()"

	mt.mainPrintHdr(ePrefix , "-")

	// "Asia/Ho_Chi_Minh"
// timeZoneName := "Asia/Vladivostok"
timeZoneName := "Asia/Ho_Chi_Minh"

tzLocPtr, err := time.LoadLocation(timeZoneName)

if err != nil {
	fmt.Printf(ePrefix +
		"Error returned by time.LoadLocation(timeZoneName)\n" +
		"timeZoneName='%v'\n" +
		"Error='%v'\n", timeZoneName, err.Error())

	return

}

	t1 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		tzLocPtr)

	ex.PrintOutDateTimeTimeZoneFields(t1, "t1 Initial Date Time")

}

// Prints Text Title Lines to the Console
func (mt mainTest) mainPrintHdr(textToPrint string, repeatStr string) {


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
