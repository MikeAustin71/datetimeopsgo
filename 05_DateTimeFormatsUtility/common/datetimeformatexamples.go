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

	dtf.GetAllDateTimeFormats()

	endTimeGetFormats := time.Now()

	etFmtOpts, _ := du.GetElapsedTime(startTime, endTimeGetFormats)
	fmt.Println("Elapsed Time For Format Map Creation: ", etFmtOpts.DurationStr)
	fmt.Println()

	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		panic(errors.New("GetAllDateTimeFormats Completed, but no formats were created! FormatMap length == 0"))
	}

	fh := FileHelper{}

	outputFile := "../formats/datetimeformats.txt"

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

		for _, fx := range dtf.FormatMap[k] {
			numOfFormats++
			fh.WriteFileStr(fmt.Sprintf("Key: %v  - %v \n", k, fx), f)
		}

	}

	endTime := time.Now()

	etFileWrite, _ := du.GetElapsedTime(endTimeGetFormats, endTime)

	et, _ := du.GetElapsedTime(startTime, endTime)

	fmt.Println("Formats File Write Operation Completed to file: ", outputFile)
	fmt.Println("Number Date Time formats Generated: ", numOfFormats)
	fmt.Println("Number of Map Keys Generated: ", numOfKeys)
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Elapsed Run Time For Write File Operations: ", etFileWrite.DurationStr)
	fmt.Println("Elapsed Run Time For All Operations: ", et.DurationStr)

}

// WriteFormatDataToFile - This method writes data to a text file.
// The text file is small, currently about 3-kilobytes in size.
// The data output to the text file describes the size of the
// slices contained in dtf.FormatMap
func WriteFormatDataToFile() {
	startTime := time.Now()

	du := DurationUtility{}

	dtf := DateTimeFormatUtility{}

	dtf.GetAllDateTimeFormats()

	endTimeGetFormats := time.Now()

	etFmtOpts, _ := du.GetElapsedTime(startTime, endTimeGetFormats)
	fmt.Println("Elapsed Time For Format Map Creation: ", etFmtOpts.DurationStr)
	fmt.Println()

	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		panic(errors.New("GetAllDateTimeFormats Completed, but no formats were created! FormatMap length == 0"))
	}

	fh := FileHelper{}

	outputFile := "../formats/fmtRun.txt"

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
	sliceLen := 0
	for _, k := range keys {

		numOfKeys++
		sliceLen = len(dtf.FormatMap[k])
		numOfFormats += sliceLen

		fh.WriteFileStr(fmt.Sprintf("Key: %v  Length Of Format Slice: %v \n", k, sliceLen), f)
	}

	endTime := time.Now()

	etFileWrite, _ := du.GetElapsedTime(endTimeGetFormats, endTime)

	et, _ := du.GetElapsedTime(startTime, endTime)

	fmt.Println("File Write Operation Completed to file: ", outputFile)
	fmt.Println("Date Time formats Generated: ", numOfFormats)
	fmt.Println("Number of Map Keys Generated: ", numOfKeys)
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

	dtf.GetAllDateTimeFormats()

	endTimeGetFormats := time.Now()

	etFmtOpts, _ := du.GetElapsedTime(startTime, endTimeGetFormats)
	fmt.Println("Elapsed Time For Format Creation: ", etFmtOpts.DurationStr)

	lFmts := NumStrUtility{}.DNumI64(int64(dtf.NumOfFormatsGenerated), ',')
	fmt.Println("Number of Formats Generated: ", lFmts)
	fmt.Println()

	startTimeParse := time.Now()

	t, err := dtf.ParseDateTimeString(dateTimeStr, probableDateTimeFormat)

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

	FmtDateTimeEverything := "Monday January 2, 2006 15:04:05 -0700 MST"
	fmt.Println()
	fmt.Println("--------------------------------")
	fmt.Println("Successful Time Parse Operation!")
	fmt.Println("Length of Input Time String: ", len(dateTimeStr))
	fmt.Println("Length of Processed Time String: ", len(dtf.DateTimeStringIn))
	fmt.Println("Time Format Map Key: ", dtf.SelectedMapIdx)
	fmt.Println("Time Format Slice Index:", dtf.SelectedSliceIdx)
	fmt.Println("Time Format Selected: ", dtf.SelectedFormat)
	fmt.Println("Original Date Time:", dateTimeStr)
	fmt.Println("Parsed time.Time:", t)
	fmt.Println("Parsed Time with Everything Format: ", t.Format(FmtDateTimeEverything))

}
