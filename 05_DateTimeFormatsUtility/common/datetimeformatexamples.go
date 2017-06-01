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

	fh := FileHelper{}

	outputFile := "../format-files/datetimeformats.txt"

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

	for _, k := range keys {

		numOfKeys++

		for keyFmt, _ := range dtf.FormatMap[k] {
			numOfFormats++
			fh.WriteFileStr(fmt.Sprintf("%07d %s\n", k, keyFmt), f)
		}

	}

	endTime := time.Now()

	etFileWrite, _ := du.GetElapsedTime(endTimeGetFormats, endTime)

	et, _ := du.GetElapsedTime(startTime, endTime)
	nu := NumStrUtility{}
	fmt.Println("Formats File Write Operation Completed to file: ", outputFile)
	fmt.Println("Number Date Time formats Generated: ", nu.DLimInt(numOfFormats, ','))
	fmt.Println("Number of Map Keys Generated: ", nu.DLimInt(numOfKeys, ','))
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

func TestParseDateTime(dateTimeStr string, probableDateTimeFormat string) {

	// tDateTime := "November 12, 2016"
	// Without length filter: time was slower
	//   Elapsed Time For Time Parse:  43-Milliseconds 499-Microseconds 100-Nanoseconds

	// With length filter
	// Elapsed Time For Time Parse:  1-Milliseconds 998-Microseconds 900-Nanoseconds

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

func printTimeParseResults(dtf DateTimeFormatUtility) {

	FmtDateTimeEverything := "Monday January 2, 2006 15:04:05 -0700 MST"
	fmt.Println()
	fmt.Println("--------------------------------")
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
	for key, value := range dtf.DictSearches {
		fmt.Println("Index Searched: ", key, "  Number of Searches per Index: ", value)
	}
	fmt.Println()
	fmt.Println("Total Number of Searches Performed: ", dtf.TotalNoOfDictSearches)
}
