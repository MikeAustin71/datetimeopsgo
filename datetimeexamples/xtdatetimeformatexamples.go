package datetimeexamples

import (
	dt "../datetime"
	"errors"
	"fmt"
	"../testlibs"
	"time"
)


// WriteAllFormatsToFile - This method will write all generated
// formats to a text file in this directory structure. Be advised,
// currently, over 6-million formats are generated. The size of
// the text file on disk is approximately 270-megabytes.
func WriteAllFormatsToFile() {
	startTime := time.Now()

	du := dt.DurationUtility{}

	dtf := dt.FormatDateTimeUtility{}

	dtf.CreateAllFormatsInMemory()

	endTimeGetFormats := time.Now()
	du.SetStartEndTimes(startTime, endTimeGetFormats, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFmtOpts, _ := du.GetYearMthDaysTime()
	fmt.Println("Elapsed Time For Format Map Creation: ", etFmtOpts.DisplayStr)
	fmt.Println()

	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		panic(errors.New("CreateAllFormatsInMemory Completed, but no formats were created! FormatMap length == 0"))
	}

	outputFile := "./format-files/datetimeformats.txt"

	writeDto, err := dtf.WriteAllFormatsInMemoryToFile(outputFile)

	if err != nil {
		panic(err)
	}

	endTime := time.Now()

	du.SetStartEndTimes(endTimeGetFormats, endTime, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFileWrite, _ := du.GetYearMthDaysTime()

	d2 := dt.DurationUtility{}

	d2.SetStartEndTimes(startTime, endTime, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	et, _ := d2.GetYearMthDaysTime()
	nu := testlibs.NumStrUtility{}
	fmt.Println("Formats File Write Operation Completed to file: ", outputFile)
	fmt.Println("Number Date Time formats Generated: ", nu.DLimInt(writeDto.NumberOfFormatsGenerated, ','))
	fmt.Println("Number of Map Keys Generated: ", nu.DLimInt(writeDto.NumberOfFormatMapKeysGenerated, ','))
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Elapsed Run Time For Write File Operations: ", etFileWrite.DisplayStr)
	fmt.Println("Elapsed Run Time For All Operations: ", et.DisplayStr)

}

// WriteFormatStatsToFile - This method writes data to a text file.
// The text file is small, currently about 3-kilobytes in size.
// The data output to the text file describes the size of the
// slices contained in dtf.FormatMap
func WriteFormatStatsToFile() {

	dtf := dt.FormatDateTimeUtility{}

	startTime := time.Now()

	dtf.CreateAllFormatsInMemory()

	endTimeCreateFormats := time.Now()
	du := dt.DurationUtility{}
	du.SetStartEndTimes(startTime, endTimeCreateFormats, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFmtOpts, _ := du.GetYearMthDaysTime()
	nu := testlibs.NumStrUtility{}

	fmt.Println("                 Format Map Creation Stats                          ")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Elapsed Time For Format Map Creation: ", etFmtOpts.DisplayStr)
	fmt.Println("Number of Formats Created: ", nu.DLimInt(dtf.NumOfFormatsGenerated, ','))
	fmt.Println("Number of Format Map Keys Created: ", nu.DLimInt(len(dtf.FormatMap), ','))
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println()

	outputDto, err := dtf.WriteFormatStatsToFile("./formats/fmtStats.txt")

	if err != nil {
		fmt.Println("*** ERROR ***")
		fmt.Println("Output File: ", outputDto.OutputPathFileName)
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println("           Format Statistics Successfully Written to File!          ")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("File Write Operation Completed to file: ", outputDto.OutputPathFileName)
	fmt.Println("Date Time formats Generated: ", nu.DLimInt(outputDto.NumberOfFormatsGenerated, ','))
	fmt.Println("Number of Map Keys Generated: ", nu.DLimInt(outputDto.NumberOfFormatMapKeysGenerated, ','))
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Elapsed Run Time For Write File Operations: ", outputDto.ElapsedTimeForFileWriteOps)
	fmt.Println()
}

// HammerSampleDateTimes
func HammerSampleDateTimes() {
	startTime := time.Now()

	dtf := dt.FormatDateTimeUtility{}

	dtf.CreateAllFormatsInMemory()

	endTimeGetFormats := time.Now()
	du := dt.DurationUtility{}

	du.SetStartEndTimes(startTime, endTimeGetFormats, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFmtOpts, _ := du.GetYearMthDaysTime()
	fmt.Println("********************************************************")
	fmt.Println("Elapsed Time For Format Creation: ", etFmtOpts.DisplayStr)

	lFmts := testlibs.NumStrUtility{}.DLimInt(dtf.NumOfFormatsGenerated, ',')
	fmt.Println("     Number of Formats Generated: ", lFmts)
	fmt.Println("    Number of Map Keys Generated: ", len(dtf.FormatMap))
	fmt.Println("********************************************************")
	fmt.Println()

	dateTimes := getZDateTimeSamples()
	hammerTestsPerCycle := len(dateTimes)
	hammerCycles := 20
	hammerStartTime := time.Now()
	totalTests := 0

	for i := 0; i < hammerCycles; i++ {
		for _, dtTime := range dateTimes {
			totalTests++
			TestParseDateTime(dtf, dtTime, "")
		}

	}

	hammerEndTime := time.Now()
	du.SetStartEndTimes(hammerStartTime, hammerEndTime, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etParseOpts, _ := du.GetYearMthDaysTime()

	fmt.Println("********************************************************")
	fmt.Println("Elapsed Time For Parse Tests: ", etParseOpts.DisplayStr)
	fmt.Println("Number of cycles: ", hammerCycles)
	fmt.Println("Number of Tests per cycle: ", hammerTestsPerCycle)
	fmt.Println("Total Tests: ", totalTests)
	fmt.Println("********************************************************")

}

// TestParseSampleDateTimes
func TestParseSampleDateTimes() {
	startTime := time.Now()

	du := dt.DurationUtility{}

	dtf := dt.FormatDateTimeUtility{}

	dtf.CreateAllFormatsInMemory()

	endTimeGetFormats := time.Now()

	du.SetStartEndTimes(startTime, endTimeGetFormats, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFmtOpts, _ := du.GetYearMthDaysTime()
	fmt.Println("********************************************************")
	fmt.Println("Elapsed Time For Format Creation: ", etFmtOpts.DisplayStr)

	lFmts := testlibs.NumStrUtility{}.DLimInt(dtf.NumOfFormatsGenerated, ',')
	fmt.Println("     Number of Formats Generated: ", lFmts)
	fmt.Println("    Number of Map Keys Generated: ", len(dtf.FormatMap))
	fmt.Println("********************************************************")
	fmt.Println()

	dateTimes := getDateTimeSamples()

	for _, dtTime := range dateTimes {
		TestParseDateTime(dtf, dtTime, "")
	}

}

// TestParseDateTime - For running this parse method, be sure that formats
// are loaded in memory in field DurationUtility.FormatMap.
func TestParseDateTime(dtf dt.FormatDateTimeUtility, dateTimeStr string, probableDateTimeFormat string) {

	startTimeParse := time.Now()

	_, err := dtf.ParseDateTimeString(dateTimeStr, probableDateTimeFormat)

	endTimeParse := time.Now()
	fmt.Println()
	du := dt.DurationUtility{}
	du.SetStartEndTimes(startTimeParse, endTimeParse, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFmtOpts, _ := du.GetYearMthDaysTime()
	fmt.Println("Elapsed Time For Time Parse: ", etFmtOpts.DisplayStr)

	if err != nil {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Printf("Failure attempting to format date time: %v/n", dateTimeStr)
		fmt.Printf("Formatted date time string: %v/n", dtf.FormattedDateTimeStringIn)
		fmt.Println("Time Parse Failed - Error: ", err.Error())
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		return
	}

	dtf.OriginalDateTimeStringIn = dateTimeStr

	printSuccessfulTimeParseResults(dtf)

}

// TestParseDateTimeCreateFormatsInMemory
func TestParseDateTimeCreateFormatsInMemory(dateTimeStr string, probableDateTimeFormat string) {

	startTime := time.Now()

	du := dt.DurationUtility{}

	dtf := dt.FormatDateTimeUtility{}

	dtf.CreateAllFormatsInMemory()

	endTimeGetFormats := time.Now()

	du.SetStartEndTimes(startTime, endTimeGetFormats, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFmtOpts, _ := du.GetYearMthDaysTime()
	fmt.Println("Elapsed Time For Format Creation: ", etFmtOpts.DisplayStr)

	lFmts := testlibs.NumStrUtility{}.DLimInt(dtf.NumOfFormatsGenerated, ',')
	fmt.Println("Number of Formats Generated: ", lFmts)
	fmt.Println()

	startTimeParse := time.Now()

	_, err := dtf.ParseDateTimeString(dateTimeStr, probableDateTimeFormat)

	endTimeParse := time.Now()
	fmt.Println()
	du.SetStartEndTimes(startTimeParse, endTimeParse, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFmtOpts, _ = du.GetYearMthDaysTime()
	fmt.Println("Elapsed Time For Time Parse: ", etFmtOpts.DisplayStr)
	fmt.Println("Actual Duration Value: ", etFmtOpts.TimeDuration)

	du.SetStartEndTimes(startTime, endTimeParse, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFmtOpts, _ = du.GetYearMthDaysTime()
	fmt.Println("Total Elapsed Time For All Operations: ", etFmtOpts.DisplayStr)
	fmt.Println()

	if err != nil {
		printTimeParseErrorResults(dtf, err)
		return
	}

	dtf.OriginalDateTimeStringIn = dateTimeStr

	printSuccessfulTimeParseResults(dtf)
}

// TestParseDateTimeFromFile
func TestParseDateTimeFromFile(dateTimeStr string, probableDateTimeFormat string) {

	startTime := time.Now()

	du := dt.DurationUtility{}

	dtf := dt.FormatDateTimeUtility{}

	drDto, err := dtf.LoadAllFormatsFromFileIntoMemory("./format-files/datetimeformats.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("Elapsed Time For File Read and Format Creation: ", drDto.ElapsedTimeForFileReadOps)

	nu := testlibs.NumStrUtility{}

	fmt.Println("Number of Formats Generated: ", nu.DLimInt(drDto.NumberOfFormatsGenerated, ','))
	fmt.Println("   Number of Keys Generated: ", nu.DLimInt(drDto.NumberOfFormatMapKeysGenerated, ','))

	startTimeParse := time.Now()

	_, err = dtf.ParseDateTimeString(dateTimeStr, probableDateTimeFormat)

	endTimeParse := time.Now()
	fmt.Println()
	du.SetStartEndTimes(startTimeParse, endTimeParse, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFmtOpts, _ := du.GetYearMthDaysTime()

	fmt.Println("Elapsed Time For Time Parse: ", etFmtOpts.DisplayStr)
	fmt.Println("Actual Duration Value: ", etFmtOpts.TimeDuration)

	du.SetStartEndTimes(startTime, endTimeParse, dt.TzIanaUsCentral, dt.FmtDateTimeYrMDayFmtStr)
	etFmtOpts, _ = du.GetYearMthDaysTime()
	fmt.Println("Total Elapsed Time For All Operations: ", etFmtOpts.DisplayStr)
	fmt.Println()

	if err != nil {
		printTimeParseErrorResults(dtf, err)
		return
	}

	dtf.OriginalDateTimeStringIn = dateTimeStr

	printSuccessfulTimeParseResults(dtf)

}

// TestLoadandWriteFileAllFormats
func TestLoadandWriteFileAllFormats() {
	dtf := dt.FormatDateTimeUtility{}
	fmtFile := "./format-files/TestRead.txt"
	dtoR, err := dtf.LoadAllFormatsFromFileIntoMemory(fmtFile)

	if err != nil {
		panic(err)
	}

	nu := testlibs.NumStrUtility{}
	dtFmt := "2006-01-02 15:04:05.000000000"
	fmt.Println("Results of LoadAllFormatsFromFileIntoMemory")
	fmt.Println("-------------------------------------------")
	fmt.Println("  Target Read File: ", dtoR.PathFileName)
	fmt.Println("   Read Start Time: ", dtoR.FileReadStartTime.Format(dtFmt))
	fmt.Println("     Read End Time: ", dtoR.FileReadEndTime.Format(dtFmt))
	fmt.Println("      Elapsed Time: ", dtoR.ElapsedTimeForFileReadOps)
	fmt.Println("Number of Map Keys: ", nu.DLimInt(dtoR.NumberOfFormatMapKeysGenerated, ','))
	fmt.Println(" Number of Formats: ", nu.DLimInt(dtoR.NumberOfFormatsGenerated, ','))

	outputFile := "./format-files/TestOutput.txt"
	dtoW, err := dtf.WriteAllFormatsInMemoryToFile(outputFile)

	if err != nil {
		panic(err)
	}

	fmt.Println("Results of WriteAllFormatsInMemoryToFile")
	fmt.Println("----------------------------------------")
	fmt.Println("Target Write File: ", dtoW.OutputPathFileName)
	fmt.Println(" Write Start Time: ", dtoW.FileWriteStartTime.Format(dtFmt))
	fmt.Println("   Write End Time: ", dtoW.FileWriteEndTime.Format(dtFmt))
	fmt.Println("Number of Map Keys: ", nu.DLimInt(dtoW.NumberOfFormatMapKeysGenerated, ','))
	fmt.Println(" Number of Formats: ", nu.DLimInt(dtoW.NumberOfFormatsGenerated, ','))

}

// TestSingleDigitFormats
func TestSingleDigitFormats() {

	dateTimes := getDateTimeSamples()

	tFmts := make([]string, 0, 20)

	tFmts = append(tFmts, "2006-1-2")
	tFmts = append(tFmts, "06-1-2")
	tFmts = append(tFmts, "2006-1-2 15:4:5")
	tFmts = append(tFmts, "2006-1-2 15:04:5")
	tFmts = append(tFmts, "2006-1-2 15:04:05")
	tFmts = append(tFmts, "2006-1-2 15:04")
	tFmts = append(tFmts, "2006-1-2 15:4")
	tFmts = append(tFmts, "06-1-2 15:4:5")
	tFmts = append(tFmts, "2006-1-2 3:4:5 PM")
	tFmts = append(tFmts, "06-1-2 3:4:5 pm")
	tFmts = append(tFmts, "2006-1-2 3:4:5PM")
	tFmts = append(tFmts, "06-1-2 3:4:5pm")
	tFmts = append(tFmts, "2006-1-2 03:04:05PM")
	tFmts = append(tFmts, "06-1-2 03:04:05pm")
	tFmts = append(tFmts, "2006-1-2 3:4:5 P.M.")
	tFmts = append(tFmts, "06-1-2 3:4:5 p.m.")
	tFmts = append(tFmts, "2006-1-2 3:4:5P.M.")
	tFmts = append(tFmts, "06-1-2 3:4:5p.m.")
	tFmts = append(tFmts, "2006-1-2 3:4 P.M.")
	tFmts = append(tFmts, "06-1-2 3:4 p.m.")
	tFmts = append(tFmts, "2006-1-2 3:4P.M.")
	tFmts = append(tFmts, "06-1-2 3:4p.m.")

	tFmts = append(tFmts, "2006-1-2 3:04PM")
	tFmts = append(tFmts, "06-1-2 3:04pm")
	tFmts = append(tFmts, "2006-1-2 3:04P.M.")
	tFmts = append(tFmts, "06-1-2 3:04p.m.")

	tFmts = append(tFmts, "2006-1-2 3:04 P.M.")
	tFmts = append(tFmts, "06-1-2 3:04 p.m.")
	tFmts = append(tFmts, "2006-1-2 3:04P.M.")
	tFmts = append(tFmts, "06-1-2 3:04p.m.")
	tFmts = append(tFmts, "2016-11-26 16:26 CST -0600")
	tFmts = append(tFmts, "2017-6-2 00:33:21 CDT -0500")
	tFmts = append(tFmts, "2016-2-5 6:02 CDT -0600")
	tFmts = append(tFmts, "June 12th, 2016 4:26 PM")

	fmtDateTimeEverything := "Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	var isSuccess bool

	for _, tDateTimeStr := range dateTimes {

		isSuccess = false

		for _, xFmt := range tFmts {

			t, err := time.Parse(xFmt, tDateTimeStr)

			if err == nil {
				fmt.Println("Success = Input: ", tDateTimeStr, " Format: ", xFmt, " Output: ", t.Format(fmtDateTimeEverything))
				isSuccess = true
			}
		}

		if !isSuccess {
			fmt.Println("Failure - Could Not Locatate Format for Time String: ", tDateTimeStr)
		}

	}

	return
}

func getYDateTimeSamples() [][][]string {
	d := make([][][]string, 0)
	// FmtDateTimeEverything := "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
	d = append(d, [][]string{{"Saturday 11/12/2016 4:26 PM", "Saturday November 12, 2016 16:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"7-6-16 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"November 12, 2016", "Saturday November 12, 2016 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"November 12, 11:26pm -0600 CST 2016", "Saturday November 12, 2016 23:26:00.000000000 -0600 CST"}})
	d = append(d, [][]string{{"November 12, 2016 11:6pm +0000 UTC", "Saturday November 12, 2016 23:06:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"November 12, 2016 1:6pm +0000 UTC", "Saturday November 12, 2016 13:06:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"November 12, 2016 1:06pm -0500 EST", "Saturday November 12, 2016 13:06:00.000000000 -0500 EST"}})
	d = append(d, [][]string{{"5/31/2017 23:2:17 -0700 PDT", "Wednesday May 31, 2017 23:02:17.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"2016-11-12 23:26:00 +0000 UTC", "Saturday November 12, 2016 23:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"2016-11-12 23:26:00Z", "Saturday November 12, 2016 23:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"2017-6-12 11:26 p.m. Z", "Monday June 12, 2017 23:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"2017-11-26 16:26 -0600", "Sunday November 26, 2017 16:26:00.000000000 -0600 CST"}})
	d = append(d, [][]string{{"2017-6-5 17:16 +0100 BST", "Monday June 5, 2017 17:16:00.000000000 +0100 BST"}})
	d = append(d, [][]string{{"11/12/16 4:26 PM", "Saturday November 12, 2016 16:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 4:4 P.M.", "Saturday November 12, 2016 16:04:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 4:4:4.012 AM", "Saturday November 12, 2016 04:04:04.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/1/16 4:4:04.012 A.M.", "Tuesday November 1, 2016 04:04:04.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"Monday June 5, 2017 17:24:46.064223400 -0500 CDT", "Monday June 5, 2017 17:24:46.064223400 -0500 CDT"}})
	d = append(d, [][]string{{"6-5-2017 17:30:17 -0700 PDT", "Monday June 5, 2017 17:30:17.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"11/12/16 4:04:0.012 PM", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/2/16 04:04:0.012 PM", "Wednesday November 2, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 4:04:00.012 PM", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 04:4:0.012 PM", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 04:04:00.012 AM", "Saturday November 12, 2016 04:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 04:04:00.012 A.M.", "Saturday November 12, 2016 04:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"11/12/16 04:4:0.012 P.M.", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC"}})
	d = append(d, [][]string{{"5/27/2017 11:42PM CDT", "Saturday May 27, 2017 23:42:00.000000000 -0500 CDT"}})
	d = append(d, [][]string{{"06/1/2017 11:42 -0700 PDT", "Thursday June 1, 2017 11:42:00.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"2016-11-26 16:26 CDT -0600", "Saturday November 26, 2016 16:26:00.000000000 -0600 CDT"}})
	d = append(d, [][]string{{"2016/11/26 16:2:3 PDT -0700", "Saturday November 26, 2016 16:02:03.000000000 -0700 PDT"}})
	d = append(d, [][]string{{"June 12th, 2016 4:26 PM", "Sunday June 12, 2016 16:26:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"05.03.2017", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"2017.3.5", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC"}})
	d = append(d, [][]string{{"6/27/2017 23:26:01 -0500 CDT", "Tuesday June 27, 2017 23:26:01.000000000 -0500 CDT"}})
	d = append(d, [][]string{{"23:26:01 -0500 CDT", "Saturday January 1, 0000 23:26:01.000000000 -0500 CDT"}})
	d = append(d, [][]string{{"11-26-2016 16:26 -0600 CST", "Saturday November 26, 2016 16:26:00.000000000 -0600 CST"}})

	return d
}

func getDateTimeSamples() []string {
	dateTime := make([]string, 0)

	dateTime = append(dateTime, "2016-11-26 16:26")
	dateTime = append(dateTime, "2016-11-26 16:26:05")
	dateTime = append(dateTime, "2016-11-26 16:6:5")
	dateTime = append(dateTime, "2016-1-3 16:6")
	dateTime = append(dateTime, "2016-12-23 2:16")
	dateTime = append(dateTime, "2016-2-21 2:6")
	dateTime = append(dateTime, "2016-12-3 2:16AM")
	dateTime = append(dateTime, "2016-12-3 2:6AM")
	dateTime = append(dateTime, "2016-2-23 11:6AM")
	dateTime = append(dateTime, "2016-1-13 11:16AM")
	dateTime = append(dateTime, "1 June 2017 11:16AM")
	dateTime = append(dateTime, "1 Jan 2017 11:16AM")
	dateTime = append(dateTime, "Friday June 2, 2017 21:5 -0600 CDT")
	dateTime = append(dateTime, "November 12, 2016")
	dateTime = append(dateTime, "Monday 11/12/2016 4:26 PM")
	dateTime = append(dateTime, "June 1st, 2017 4:26 PM")
	dateTime = append(dateTime, "June 3rd, 2017 4:26 PM")
	dateTime = append(dateTime, "June 12th, 2016 4:26 PM")
	dateTime = append(dateTime, "7-6-16 9:30AM")
	dateTime = append(dateTime, "2016-11-26 16:26 -0600")
	dateTime = append(dateTime, "5/27/2017 11:42PM CDT")
	dateTime = append(dateTime, "12/2/2017 11:42PM CST")
	dateTime = append(dateTime, "2016-11-26 16:26 CDT -0600")

	return dateTime

}

func printSuccessfulTimeParseResults(dtf dt.FormatDateTimeUtility) {

	FmtDateTimeEverything := "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
	fmt.Println()
	fmt.Println("--------------------------------------------------------")
	fmt.Println("Successful Time Parse Operation!")
	fmt.Println("Original Input Date Time:", dtf.OriginalDateTimeStringIn)
	fmt.Println("Formatted Input Date Time: ", dtf.FormattedDateTimeStringIn)
	fmt.Println("Length of Original Input Time String: ", len(dtf.OriginalDateTimeStringIn))
	fmt.Println("Length of Processed Time String: ", len(dtf.FormattedDateTimeStringIn))
	fmt.Println("Time Format Map Key: ", dtf.SelectedMapIdx)
	fmt.Println("Time Format Selected: ", dtf.SelectedFormat)
	fmt.Println("Selected Time Format Source: ", dtf.SelectedFormatSource)
	fmt.Println("Parsed time.Time:", dtf.DateTimeOut)
	fmt.Println("Parsed Time with Everything Format: ", dtf.DateTimeOut.Format(FmtDateTimeEverything))
	fmt.Println("Detailed Search Pattern: ")
	lDs := len(dtf.DictSearches)
	for i := 0; i < lDs; i++ {
		fmt.Println("Index Searched: ", dtf.DictSearches[i][0][0], "  Number of Searches per Index: ", dtf.DictSearches[i][0][1])
	}
	fmt.Println()
	fmt.Println("Total Number of Searches Performed: ", dtf.TotalNoOfDictSearches)
	fmt.Println("--------------------------------------------------------")
	fmt.Println()
}

func printTimeParseErrorResults(dtf dt.FormatDateTimeUtility, err error) {
	nu := testlibs.NumStrUtility{}
	fmt.Println("Time Parse Failed - Error: ", err.Error())
	fmt.Println()
	fmt.Println("        Original Date Time String: ", dtf.OriginalDateTimeStringIn)
	fmt.Println("       Formatted Date Time String: ", dtf.FormattedDateTimeStringIn)
	fmt.Println("Length Formatted Date Time String: ", len(dtf.FormattedDateTimeStringIn))
	fmt.Println("         Total Number of Searches: ", nu.DLimInt(dtf.TotalNoOfDictSearches, ','))
	fmt.Println("Detailed Search Pattern: ")
	lDs := len(dtf.DictSearches)
	for i := 0; i < lDs; i++ {
		fmt.Println("Index Searched: ", dtf.DictSearches[i][0][0], "  Number of Searches per Index: ", dtf.DictSearches[i][0][1])
	}
	return

}

func getZDateTimeSamples() []string {
	d := make([]string, 0)
	// FmtDateTimeEverything := "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
	d = append(d, "Saturday 11/12/2016 4:26 PM", "Saturday November 12, 2016 16:26:00.000000000 +0000 UTC")
	d = append(d, "7-6-16 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC")
	d = append(d, "7-6-2016 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC")
	d = append(d, "7-06-2016 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC")
	d = append(d, "07-6-2016 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC")
	d = append(d, "07-06-2016 9:30AM", "Wednesday July 6, 2016 09:30:00.000000000 +0000 UTC")
	d = append(d, "November 12, 2016", "Saturday November 12, 2016 00:00:00.000000000 +0000 UTC")
	d = append(d, "12 Nov 2016", "Saturday November 12, 2016 00:00:00.000000000 +0000 UTC")
	d = append(d, "12 November 2016", "Saturday November 12, 2016 00:00:00.000000000 +0000 UTC")
	d = append(d, "November 12, 11:26pm -0600 CST 2016", "Saturday November 12, 2016 23:26:00.000000000 -0600 CST")
	d = append(d, "12 November 2016 23:26:00 -0600 CST", "Saturday November 12, 2016 23:26:00.000000000 -0600 CST")
	d = append(d, "November 12, 2016 11:6pm +0000 UTC", "Saturday November 12, 2016 23:06:00.000000000 +0000 UTC")
	d = append(d, "November 12, 2016 11:6 p m +0000 UTC", "Saturday November 12, 2016 23:06:00.000000000 +0000 UTC")
	d = append(d, "November 12, 2016 1:6pm +0000 UTC", "Saturday November 12, 2016 13:06:00.000000000 +0000 UTC")
	d = append(d, "November 12, 2016 1:06pm -0500 EST", "Saturday November 12, 2016 13:06:00.000000000 -0500 EST")
	d = append(d, "2016-11-12 13:6 -0500 EST", "Saturday November 12, 2016 13:06:00.000000000 -0500 EST")
	d = append(d, "5/31/2017 23:2:17 -0700 PDT", "Wednesday May 31, 2017 23:02:17.000000000 -0700 PDT")
	d = append(d, "2016-11-12 23:26:00 +0000 UTC", "Saturday November 12, 2016 23:26:00.000000000 +0000 UTC")
	d = append(d, "2016-11-12 23:26:00Z", "Saturday November 12, 2016 23:26:00.000000000 +0000 UTC")
	d = append(d, "2017-6-12 11:26 p.m. Z", "Monday June 12, 2017 23:26:00.000000000 +0000 UTC")
	d = append(d, "2017-11-26 16:26 -0600 CST", "Sunday November 26, 2017 16:26:00.000000000 -0600 CST")
	d = append(d, "2017-6-5 17:16 +0100 BST", "Monday June 5, 2017 17:16:00.000000000 +0100 BST")
	d = append(d, "2017-6-05 17:16 +0100 BST", "Monday June 5, 2017 17:16:00.000000000 +0100 BST")
	d = append(d, "2017-06-5 17:16 +0100 BST", "Monday June 5, 2017 17:16:00.000000000 +0100 BST")
	d = append(d, "2017-06-05 17:16 +0100 BST", "Monday June 5, 2017 17:16:00.000000000 +0100 BST")
	d = append(d, "11/12/16 4:26 PM", "Saturday November 12, 2016 16:26:00.000000000 +0000 UTC")
	d = append(d, "11/12/16 4:4 P.M.", "Saturday November 12, 2016 16:04:00.000000000 +0000 UTC")
	d = append(d, "11/12/16 4:4:4.012 AM", "Saturday November 12, 2016 04:04:04.012000000 +0000 UTC")
	d = append(d, "11/1/16 4:4:04.012 A.M.", "Tuesday November 1, 2016 04:04:04.012000000 +0000 UTC")
	d = append(d, "11/1/2016 4:4:04.012 A.M.", "Tuesday November 1, 2016 04:04:04.012000000 +0000 UTC")
	d = append(d, "11/1/2016 4:4:04.012 A.M. ", "Tuesday November 1, 2016 04:04:04.012000000 +0000 UTC")
	d = append(d, "Monday June 5, 2017 17:24:46.064223400 -0500 CDT", "Monday June 5, 2017 17:24:46.064223400 -0500 CDT")
	d = append(d, "6-5-2017 17:30:17 -0700 PDT", "Monday June 5, 2017 17:30:17.000000000 -0700 PDT")
	d = append(d, "06-05-2017 17:30:17 -0700 PDT", "Monday June 5, 2017 17:30:17.000000000 -0700 PDT")
	d = append(d, "06-5-2017 17:30:17 -0700 PDT", "Monday June 5, 2017 17:30:17.000000000 -0700 PDT")
	d = append(d, "11/12/16 4:04:0.012 PM", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC")
	d = append(d, "11/2/16 04:04:0.012 PM", "Wednesday November 2, 2016 16:04:00.012000000 +0000 UTC")
	d = append(d, "11/12/16 4:04:00.012 PM", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC")
	d = append(d, "11/12/16 04:4:0.012 PM", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC")
	d = append(d, "11/12/16 04:04:00.012 AM", "Saturday November 12, 2016 04:04:00.012000000 +0000 UTC")
	d = append(d, "11/12/16 04:04:00.012 am", "Saturday November 12, 2016 04:04:00.012000000 +0000 UTC")
	d = append(d, "11/12/16 04:04:00.012 A.M.", "Saturday November 12, 2016 04:04:00.012000000 +0000 UTC")
	d = append(d, "11/12/16 04:4:0.012 P.M.", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC")
	d = append(d, "11/12/2016 04:4:0.012 P.M.", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC")
	d = append(d, "11/12/2016 04:4:0.012 PM.", "Saturday November 12, 2016 16:04:00.012000000 +0000 UTC")
	d = append(d, "5/27/2017 11:42PM CDT", "Saturday May 27, 2017 23:42:00.000000000 -0500 CDT")
	d = append(d, "06/1/2017 11:42 -0700 PDT", "Thursday June 1, 2017 11:42:00.000000000 -0700 PDT")
	d = append(d, "2016-11-26 16:26 CDT -0600", "Saturday November 26, 2016 16:26:00.000000000 -0600 CDT")
	d = append(d, "2016/11/26 16:2:3 PDT -0700", "Saturday November 26, 2016 16:02:03.000000000 -0700 PDT")
	d = append(d, "June 12th, 2016 4:26 PM", "Sunday June 12, 2016 16:26:00.000000000 +0000 UTC")
	d = append(d, "05.03.2017", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC")
	d = append(d, "5.03.2017", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC")
	d = append(d, "5.3.2017", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC")
	d = append(d, "5.3.'17", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC")
	d = append(d, "2017.3.5", "Sunday March 5, 2017 00:00:00.000000000 +0000 UTC")
	d = append(d, "6/27/2017 23:26:01 -0500 CDT", "Tuesday June 27, 2017 23:26:01.000000000 -0500 CDT")
	d = append(d, "23:26:01 -0500 CDT", "Saturday January 1, 0000 23:26:01.000000000 -0500 CDT")
	d = append(d, "11-26-2016 16:26 -0600 CST", "Saturday November 26, 2016 16:26:00.000000000 -0600 CST")
	d = append(d, "11-26-2016 16:26:0 -0600 CST", "Saturday November 26, 2016 16:26:00.000000000 -0600 CST")
	d = append(d, "Monday June 5th2017 17:24:46.064223400 -0500 CDT", "Monday June 5, 2017 17:24:46.064223400 -0500 CDT")
	d = append(d, "5/27/2017 11:42PMCDT", "Saturday May 27, 2017 23:42:00.000000000 -0500 CDT")
	d = append(d, "06/1/2017 11:42 PM-0700 PDT", "Thursday June 1, 2017 23:42:00.000000000 -0700 PDT")
	d = append(d, "06/1/2017 11:42:00   PM  -0700 PDT", "Thursday June 1, 2017 23:42:00.000000000 -0700 PDT")
	d = append(d, "June 1st, 2017 11:42:00PM -0700 PDT", "Thursday June 1, 2017 23:42:00.000000000 -0700 PDT")
	d = append(d, "June 2nd 2017 11:42:00PM -0700 PDT", "Friday June 2, 2017 23:42:00.000000000 -0700 PDT")
	d = append(d, "June 3rd, 2017 11:42:00PM -0700 PDT", "Saturday June 3, 2017 23:42:00.000000000 -0700 PDT")

	return d
}
