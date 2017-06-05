package common

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

// WriteAllFormatsToFile - This method will write all generated
// formats to a text file in this directory structure. Be advised,
// currently, over 6-million formats are generated. The size of
// the text file on disk is approximately 270-megabytes.
func WriteAllFormatsToFile() {
	startTime := time.Now()

	du := DurationUtility{}

	dtf := DateTimeFormatUtility{}

	dtf.CreateAllFormatsInMemory()

	endTimeGetFormats := time.Now()

	etFmtOpts, _ := du.GetElapsedTime(startTime, endTimeGetFormats)
	fmt.Println("Elapsed Time For Format Map Creation: ", etFmtOpts.DurationStr)
	fmt.Println()

	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		panic(errors.New("CreateAllFormatsInMemory Completed, but no formats were created! FormatMap length == 0"))
	}

	outputFile := "../format-files/datetimeformats.txt"

	writeDto, err := dtf.WriteAllFormatsInMemoryToFile(outputFile)

	if err != nil {
		panic(err)
	}

	endTime := time.Now()

	etFileWrite, _ := du.GetElapsedTime(endTimeGetFormats, endTime)

	et, _ := du.GetElapsedTime(startTime, endTime)
	nu := NumStrUtility{}
	fmt.Println("Formats File Write Operation Completed to file: ", outputFile)
	fmt.Println("Number Date Time formats Generated: ", nu.DLimInt(writeDto.NumberOfFormatsGenerated, ','))
	fmt.Println("Number of Map Keys Generated: ", nu.DLimInt(writeDto.NumberOfFormatMapKeysGenerated, ','))
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Elapsed Run Time For Write File Operations: ", etFileWrite.DurationStr)
	fmt.Println("Elapsed Run Time For All Operations: ", et.DurationStr)

}

// WriteFormatDataToFile - This method writes data to a text file.
// The text file is small, currently about 3-kilobytes in size.
// The data output to the text file describes the size of the
// slices contained in dtf.FormatMap
func WriteFormatStatsToFile() {
	startTime := time.Now()

	du := DurationUtility{}

	dtf := DateTimeFormatUtility{}

	dtf.CreateAllFormatsInMemory()

	endTimeGetFormats := time.Now()

	etFmtOpts, _ := du.GetElapsedTime(startTime, endTimeGetFormats)
	fmt.Println("Elapsed Time For Format Map Creation: ", etFmtOpts.DurationStr)
	fmt.Println()

	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		panic(errors.New("CreateAllFormatsInMemory Completed, but no formats were created! FormatMap length == 0"))
	}

	fh := FileHelper{}

	outputFile := "../formats/fmtStats.txt"

	if fh.DoesFileExist(outputFile) {
		err := fh.DeleteDirFile(outputFile)

		if err != nil {
			panic(errors.New("Error from DeleteDirFile() Error:" + err.Error()))
		}
	}

	f, err := fh.CreateDirAndFile(outputFile)

	if err != nil {
		panic(errors.New("Error from CreateDirAndFile Error:" + err.Error()))
	}

	defer f.Close()

	var keys []int
	for k := range dtf.FormatMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	numOfKeys := 0
	numOfFormats := 0
	mapLen := 0
	fh.WriteFileStr("Length - Number Of Formats\n", f)
	for _, k := range keys {

		numOfKeys++
		mapLen = len(dtf.FormatMap[k])
		numOfFormats += mapLen

		fh.WriteFileStr(fmt.Sprintf("%06d%18d\n", k, mapLen), f)
	}

	endTime := time.Now()

	etFileWrite, _ := du.GetElapsedTime(endTimeGetFormats, endTime)

	et, _ := du.GetElapsedTime(startTime, endTime)
	nu := NumStrUtility{}
	fmt.Println("File Write Operation Completed to file: ", outputFile)
	fmt.Println("Date Time formats Generated: ", nu.DLimInt(numOfFormats, ','))
	fmt.Println("Number of Map Keys Generated: ", nu.DLimInt(numOfKeys, ','))
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Elapsed Run Time For Write File Operations: ", etFileWrite.DurationStr)
	fmt.Println("Elapsed Run Time For All Operations: ", et.DurationStr)

}

func TestParseSampleDateTimes() {
	startTime := time.Now()

	du := DurationUtility{}

	dtf := DateTimeFormatUtility{}

	dtf.CreateAllFormatsInMemory()

	endTimeGetFormats := time.Now()

	etFmtOpts, _ := du.GetElapsedTime(startTime, endTimeGetFormats)
	fmt.Println("********************************************************")
	fmt.Println("Elapsed Time For Format Creation: ", etFmtOpts.DurationStr)

	lFmts := NumStrUtility{}.DLimInt(dtf.NumOfFormatsGenerated, ',')
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
func TestParseDateTime(dtf DateTimeFormatUtility, dateTimeStr string, probableDateTimeFormat string) {

	du := DurationUtility{}

	startTimeParse := time.Now()

	_, err := dtf.ParseDateTimeString(dateTimeStr, probableDateTimeFormat)

	endTimeParse := time.Now()
	fmt.Println()
	etFmtOpts, _ := du.GetElapsedTime(startTimeParse, endTimeParse)
	fmt.Println("Elapsed Time For Time Parse: ", etFmtOpts.DurationStr)

	if err != nil {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Printf("Failure attempting to format date time: %v/n", dateTimeStr)
		fmt.Println("Time Parse Failed - Error: ", err.Error())
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		return
	}

	dtf.OriginalDateTimeStringIn = dateTimeStr

	printTimeParseResults(dtf)

}

func TestParseDateTimeCreateFormatsInMemory(dateTimeStr string, probableDateTimeFormat string) {

	startTime := time.Now()

	du := DurationUtility{}

	dtf := DateTimeFormatUtility{}

	dtf.CreateAllFormatsInMemory()

	endTimeGetFormats := time.Now()

	etFmtOpts, _ := du.GetElapsedTime(startTime, endTimeGetFormats)
	fmt.Println("Elapsed Time For Format Creation: ", etFmtOpts.DurationStr)

	lFmts := NumStrUtility{}.DLimInt(dtf.NumOfFormatsGenerated, ',')
	fmt.Println("Number of Formats Generated: ", lFmts)
	fmt.Println()

	startTimeParse := time.Now()

	_, err := dtf.ParseDateTimeString(dateTimeStr, probableDateTimeFormat)

	endTimeParse := time.Now()
	fmt.Println()
	etFmtOpts, _ = du.GetElapsedTime(startTimeParse, endTimeParse)
	fmt.Println("Elapsed Time For Time Parse: ", etFmtOpts.DurationStr)
	fmt.Println("Actual Duration Value: ", etFmtOpts.TimeDuration)

	etFmtOpts, _ = du.GetElapsedTime(startTime, endTimeParse)
	fmt.Println("Total Elapsed Time For All Operations: ", etFmtOpts.DurationStr)
	fmt.Println()

	if err != nil {
		fmt.Println("Time Parse Failed - Error: ", err.Error())
		return
	}

	dtf.OriginalDateTimeStringIn = dateTimeStr

	printTimeParseResults(dtf)
}

func TestParseDateTimeFromFile(dateTimeStr string, probableDateTimeFormat string) {

	startTime := time.Now()

	du := DurationUtility{}

	dtf := DateTimeFormatUtility{}

	drDto, err := dtf.LoadAllFormatsFromFileIntoMemory("../format-files/datetimeformats.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("Elapsed Time For File Read and Format Creation: ", drDto.ElapsedTimeForFileReadOps)

	nu := NumStrUtility{}

	fmt.Println("Number of Formats Generated: ", nu.DLimInt(drDto.NumberOfFormatsGenerated, ','))
	fmt.Println("   Number of Keys Generated: ", nu.DLimInt(drDto.NumberOfFormatMapKeysGenerated, ','))

	startTimeParse := time.Now()

	_, err = dtf.ParseDateTimeString(dateTimeStr, probableDateTimeFormat)

	endTimeParse := time.Now()
	fmt.Println()
	etFmtOpts, _ := du.GetElapsedTime(startTimeParse, endTimeParse)

	fmt.Println("Elapsed Time For Time Parse: ", etFmtOpts.DurationStr)
	fmt.Println("Actual Duration Value: ", etFmtOpts.TimeDuration)

	etFmtOpts, _ = du.GetElapsedTime(startTime, endTimeParse)
	fmt.Println("Total Elapsed Time For All Operations: ", etFmtOpts.DurationStr)
	fmt.Println()

	if err != nil {
		fmt.Println("Time Parse Failed - Error: ", err.Error())
		return
	}

	dtf.OriginalDateTimeStringIn = dateTimeStr

	printTimeParseResults(dtf)

}

/*
func TestMatchTimeFormats() {
	tf1:= "\\d:\\d:\\d"


}
*/

func TestLoadandWriteFileAllFormats() {
	dtf := DateTimeFormatUtility{}
	fmtFile := "D:/go/work/src/MikeAustin71/datetimeopsgo/05_DateTimeFormatsUtility/format-files/TestRead.txt"
	dtoR, err := dtf.LoadAllFormatsFromFileIntoMemory(fmtFile)

	if err != nil {
		panic(err)
	}

	nu := NumStrUtility{}
	dtFmt := "2006-01-02 15:04:05.000000000"
	fmt.Println("Results of LoadAllFormatsFromFileIntoMemory")
	fmt.Println("-------------------------------------------")
	fmt.Println("  Target Read File: ", dtoR.PathFileName)
	fmt.Println("   Read Start Time: ", dtoR.FileReadStartTime.Format(dtFmt))
	fmt.Println("     Read End Time: ", dtoR.FileReadEndTime.Format(dtFmt))
	fmt.Println("      Elapsed Time: ", dtoR.ElapsedTimeForFileReadOps)
	fmt.Println("Number of Map Keys: ", nu.DLimInt(dtoR.NumberOfFormatMapKeysGenerated, ','))
	fmt.Println(" Number of Formats: ", nu.DLimInt(dtoR.NumberOfFormatsGenerated, ','))

	outputFile := "D:/go/work/src/MikeAustin71/datetimeopsgo/05_DateTimeFormatsUtility/format-files/TestOutput.txt"
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

func printTimeParseResults(dtf DateTimeFormatUtility) {

	FmtDateTimeEverything := "Monday January 2, 2006 15:04:05 -0700 MST"
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
