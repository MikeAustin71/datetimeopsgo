package datetime

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

// WriteDateTimeFormatsToFileDto - Used to output date time formats to
// a file. Reference method FormatDateTimeUtility.WriteAllFormatsInMemoryToFile()
//
type WriteDateTimeFormatsToFileDto struct {
	OutputPathFileName             string
	NumberOfFormatsGenerated       int
	NumberOfFormatMapKeysGenerated int
	FileWriteStartTime             time.Time
	FileWriteEndTime               time.Time
	ElapsedTimeForFileWriteOps     string
}

// ReadDateTimeFormatsFromFileDto - Used to read Date Time Formats
// from a file rather then generating those formats directly to memory.
//
// Reference method FormatDateTimeUtility.LoadAllFormatsFromFileIntoMemory()
//
type ReadDateTimeFormatsFromFileDto struct {
	PathFileName                   string
	NumberOfFormatsGenerated       int
	NumberOfFormatMapKeysGenerated int
	FileReadStartTime              time.Time
	FileReadEndTime                time.Time
	ElapsedTimeForFileReadOps      string
}

// formatDateTimeGenerator
type formatDateTimeGenerator struct {
	dayOfWeek            string
	dayOfWeekSeparator   string
	mthDay               string
	mthDayYear           string
	afterMthDaySeparator string
	dateTimeSeparator    string
	timeElement          string
	offsetSeparator      string
	offsetElement        string
	timeZoneSeparator    string
	timeZoneElement      string
	year                 string
}

// SearchStrings - Used by FormatDateTimeUtility
type SearchStrings struct {
	PreTrimSearchStrs [][][]string
	TimeFmtRegEx      [][][]string
}

// parseDateTimeDto
type parseDateTimeDto struct {
	isSuccessful              bool
	formattedDateTimeStringIn string
	selectedMapIdx            int
	selectedFormat            string
	totalNoOfDictSearches     int
	dateTimeOut               time.Time
	err                       error
}

/*

	'FormatDateTimeUtility' is part of the date time operations library. The source code repository
 	for this file is located at:
					https://github.com/MikeAustin71/datetimeopsgo.git

	The location of this source file is:
					MikeAustin71\datetimeopsgo\datetime\formatdatetimeutility.go

Background

==========================


	This utility is designed to convert date time strings into specific
	date time values; that is type 'time.Time' structures. The methods
	provided address the problem of converting date time strings entered
	by users which may not follow a single date time formatting standard.

	Users tend to key-enter dates and times using a variety of different
	formats depending on their national or cultural norms and the their
	personal preferences for configuring date-times. As a result users
	don't tend to follow any consistent standards when key-entering dates
	and times.

	To address this issue, the 'FormatDateTimeUtility' generates about
	1.5-million possible format patterns and applies these to date times
	submitted by users in non-standard string formats.

	On my machine, it requires about two seconds to generate and configure
	in memory the 1.5-million format maps. Thereafter, date time strings
	are usually parsed in under 35-milliseconds.

	The utility methods are found in the file:

			datetimeopsgo/datetime/formatdatetimeutility.go

	The two most useful methods are CreateAllFormatsInMemory() and
	ParseDateTimeString().

	CreateAllFormatsInMemory() creates the 1.5-million possible formats
	in memory and must be run be one begins parsing date time strings.

	ParseDateTimeString() receives the date time string and parses
	said string into a time.Time value. The method uses a an algorithm
	based on input string length and concurrent search operations.

	Other methods are provided which allow one to write the 1.5-million
	format maps to a file. Conversely, one call also read format maps
	into memory from a disk file.

Overview and General Usage

==========================

  Type 'FormatDateUtility' is designed for one purpose, to analyze date strings passed in
  by the user and convert these strings to valid numeric date time values. Given the number
	of date time formats used by different countries and cultures across the globe, this can
	be a complex task.

	Use of the date time format conversion feature requires three function calls:

	1. Create the type instance
					dtf := FormatDateTimeUtility{}

  2. Create All Formats In Memory-  Currently this method generates
 			approximately 1.5-million permutations of Date Time Formats
			and stores them in memory from which they will be used to analyze
			and evaluate user entered date strings. This process currently
			takes about 1.5-seconds on my machine.

					dtf.CreateAllFormatsInMemory()

  3. Parse the user generated input string into a valid numeric date time.

			dateTime, err := dtf.ParseDateTimeString(dtString, "")

FormatDateTimeUtility Structure

===============================
*/
type FormatDateTimeUtility struct {
	OriginalDateTimeStringIn  string
	FormattedDateTimeStringIn string
	FormatMap                 map[int]map[string]int
	SelectedMapIdx            int
	SelectedFormat            string
	SelectedFormatSource      string
	DictSearches              [][][]int
	TotalNoOfDictSearches     int
	DateTimeOut               time.Time
	NumOfFormatsGenerated     int
	FormatSearchReplaceStrs   SearchStrings
}

// CreateAllFormatsInMemory - Currently this method generates
// approximately 1.5-million permutations of Date Time Formats
// and stores them in memory as a series of "maps of maps" using
// the field: FormatDateTimeUtility.FormatMap. This process currently
// takes about 1.5-seconds on my machine.
//
// 54-map keys are currently generated and used to access the format
// strings. Access is keyed on length of the date time string one is
// attempting to parse.
func (dtf *FormatDateTimeUtility) CreateAllFormatsInMemory() (err error) {

	dtf.FormatMap = make(map[int]map[string]int)
	dtf.NumOfFormatsGenerated = 0
	dtf.assemblePreDefinedFormats()
	err2 := dtf.assembleMthDayYearFmts()

	if err2 != nil {
		ePrefix := "FormatDateTimeUtility.CreateAllFormatsInMemory() "
		err = fmt.Errorf(ePrefix+
			"Error returned by dtf.assembleMthDayYearFmts(). Error='%v' ",
			err2.Error())
		return err
	}

	dtf.assembleEdgeCaseFormats()

	dtf.FormatSearchReplaceStrs.PreTrimSearchStrs = dtf.getPreTrimSearchStrings()
	dtf.FormatSearchReplaceStrs.TimeFmtRegEx = dtf.getTimeFmtRegEx()

	return
}

// LoadAllFormatsFromFileIntoMemory - Loads all date time formats from a specified
// text file into memory. The formats are stored in FormatDateTimeUtility.FormatMap.
// This is an alternative means of reading all available date time formats into
// memory so that they may be used to parse time strings. This method assumes that
// the text file containing the format strings was originally created by method
// FormatDateTimeUtility.WriteAllFormatsInMemoryToFile() which employs a specific
// fixed length format which can then be read back into memory.
func (dtf *FormatDateTimeUtility) LoadAllFormatsFromFileIntoMemory(
	pathFileName string) (ReadDateTimeFormatsFromFileDto, error) {

	frDto := ReadDateTimeFormatsFromFileDto{}
	frDto.PathFileName = pathFileName
	frDto.FileReadStartTime = time.Now()
	dtf.FormatMap = make(map[int]map[string]int)
	dtf.NumOfFormatsGenerated = 0

	fmtFile, err := os.Open(pathFileName)

	if err != nil {
		return frDto, fmt.Errorf("LoadAllFormatsFromFileIntoMemory- Error Opening File: %v - Error: %v", pathFileName, err.Error())
	}

	const bufLen = 2000
	lastBufIdx := 0
	var buffer []byte
	var outRecordBuff []byte
	IsEOF := false
	idx := 0
	isPartialRec := false
	buffer = make([]byte, bufLen)

	// Read File Operation
	// REDO:
	for IsEOF == false {

		bytesRead, err := fmtFile.Read(buffer)

		if err != nil {
			IsEOF = true
		}

		idx = 0

		lastBufIdx = bytesRead - 1

		// Begin Read Record Operation
		for bytesRead > 0 {

			if !isPartialRec {
				outRecordBuff = make([]byte, 0)
			} else {
				isPartialRec = false
			}

			for i := idx; i <= lastBufIdx; i++ {
				// Extract one record from buffer and process
				if buffer[i] == '\n' {
					idx = i + 1
					break
				}

				outRecordBuff = append(outRecordBuff, buffer[i])

				if i == lastBufIdx {
					isPartialRec = true
					idx = 0
				}
			}

			// Break up the record into
			// two fields, int Length and
			// string Format.
			lOutBuff := len(outRecordBuff)

			if isPartialRec || lOutBuff < 7 {
				isPartialRec = true
				break
			}

			lenField := make([]byte, 7)

			for j := 0; j < len(outRecordBuff); j++ {

				lenField[j] = outRecordBuff[j]

			}

			s := string(lenField)

			lFmt, err := strconv.Atoi(s)

			if err != nil {
				_ = fmtFile.Close()
				return frDto, fmt.Errorf(
					"LoadAllFormatsFromFileIntoMemory - Error converting Format Length field from file. Length = %v. Idx= %v. Format Count: %v",
					s, idx, frDto.NumberOfFormatsGenerated)
			}

			fmtFieldLastIdx := 7 + lFmt

			if lOutBuff < fmtFieldLastIdx+1 {
				_ = fmtFile.Close()
				return frDto, fmt.Errorf(
					"LoadAllFormatsFromFileIntoMemory - Found corrupted Output Buffer. Buffer Length %v, Length Field = %v, Output Buffer= %s Format Count: %v",
					lOutBuff, lFmt, string(outRecordBuff), frDto.NumberOfFormatsGenerated)
			}

			fmtField := make([]byte, lFmt)

			for k := 8; k <= fmtFieldLastIdx && k < len(outRecordBuff); k++ {
				fmtField[k-8] = outRecordBuff[k]
			}

			fmtStr := string(fmtField)

			// Populate FormatDateTimeUtility.FormatMap
			if dtf.FormatMap[lFmt] == nil {
				dtf.FormatMap[lFmt] = make(map[string]int)
				frDto.NumberOfFormatMapKeysGenerated++
			}

			if dtf.FormatMap[lFmt][fmtStr] == 0 {
				dtf.FormatMap[lFmt][fmtStr] = lFmt
				dtf.NumOfFormatsGenerated++
				frDto.NumberOfFormatsGenerated++
			}

			if idx > lastBufIdx {
				break
			}
		}
	}

	frDto.FileReadEndTime = time.Now()
	frDto.NumberOfFormatMapKeysGenerated = len(dtf.FormatMap)
	du := DurationTriad{}
	err = du.SetStartEndTimes(
		frDto.FileReadStartTime,
		frDto.FileReadEndTime,
		TDurCalc.StdYearMth(),
		"Local",
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		_ = fmtFile.Close()
		return ReadDateTimeFormatsFromFileDto{},
			fmt.Errorf("LoadAllFormatsFromFileIntoMemory - Error SetStartEndTimesTz() - %v",
				err.Error())
	}

	frDto.ElapsedTimeForFileReadOps = du.BaseTime.GetYearMthDaysTimeStr()

	_ = fmtFile.Close()

	return frDto, nil
}

// WriteAllFormatsInMemoryToFile - Writes all Format Data contained in
// FormatDateTimeUtility.FormatMap field to a specified output file in
// text format. Currently, about 1.4-million formats are generated and
// written to the output file.
//
// IMPORTANT! - Before you call this method, the Format Maps must
// first be created in memory. Call FormatDateTimeUtility.CreateAllFormatsInMemory()
// first, before calling this method.
func (dtf *FormatDateTimeUtility) WriteAllFormatsInMemoryToFile(outputPathFileName string) (WriteDateTimeFormatsToFileDto, error) {

	ePrefix := "FormatDateTimeUtility.WriteAllFormatsInMemoryToFile() "

	fwDto := WriteDateTimeFormatsToFileDto{}

	fwDto.FileWriteStartTime = time.Now()
	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		return fwDto, errors.New(ePrefix +
			"Error - There are NO Formats in Memory -  FormatMap length == 0. " +
			"You MUST call CreateAllFormatsInMemory() first!")
	}

	outF, err := os.Create(outputPathFileName)

	if err != nil {
		return fwDto,
			fmt.Errorf(ePrefix+
				"Error - Failed create output file %v. Error: %v",
				outputPathFileName, err.Error())
	}

	var keys []int
	for k := range dtf.FormatMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {

		fwDto.NumberOfFormatMapKeysGenerated++

		for keyFmt := range dtf.FormatMap[k] {
			fwDto.NumberOfFormatsGenerated++
			_, err := outF.WriteString(fmt.Sprintf("%07d %s\n", k, keyFmt))

			if err != nil {
				_ = outF.Close()
				return WriteDateTimeFormatsToFileDto{},
					fmt.Errorf(ePrefix+"Error writing Format "+
						"data to output file %v. Error: %v", outputPathFileName, err.Error())
			}
		}
	}

	err = outF.Sync()

	if err != nil {
		_ = outF.Close()
		return WriteDateTimeFormatsToFileDto{},
			fmt.Errorf(ePrefix+"Error returned by  outF.Sync()"+
				"Error: %v", err.Error())
	}

	fwDto.FileWriteEndTime = time.Now()

	du := DurationTriad{}

	err = du.SetStartEndTimes(
		fwDto.FileWriteStartTime,
		fwDto.FileWriteEndTime,
		TDurCalc.StdYearMth(),
		"Local",
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		_ = outF.Close()
		return WriteDateTimeFormatsToFileDto{},
			fmt.Errorf(ePrefix+
				"Error Setting Start End Times for Duration Calculation Error: %v",
				err.Error())
	}

	fwDto.OutputPathFileName = outputPathFileName

	fwDto.ElapsedTimeForFileWriteOps = du.BaseTime.GetYearMthDaysTimeStr()

	err = outF.Close()

	if err != nil {
		return WriteDateTimeFormatsToFileDto{},
			fmt.Errorf(ePrefix+
				"Error returned by outF.Close(). Error: %v",
				err.Error())
	}

	return fwDto, nil
}

// WriteFormatStatsToFile - This method writes data to a text file.
// The text file is small, currently about 3-kilobytes in size.
// The data output to the text file describes the size of the
// slices contained in dtf.FormatMap.
//
// IMPORTANT! - Before you call this method, the Format Maps must
// first be created in memory. Call FormatDateTimeUtility.CreateAllFormatsInMemory()
// first, before calling this method.
func (dtf *FormatDateTimeUtility) WriteFormatStatsToFile(
	outputPathFileName string) (WriteDateTimeFormatsToFileDto, error) {

	ePrefix := "FormatDateTimeUtility.WriteFormatStatsToFile() "
	outputDto := WriteDateTimeFormatsToFileDto{}
	outputDto.OutputPathFileName = outputPathFileName
	outputDto.FileWriteStartTime = time.Now()

	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		return outputDto,
			errors.New(ePrefix +
				"Error - There are NO Formats in Memory -  FormatMap length == 0. " +
				"You MUST call CreateAllFormatsInMemory() first!")
	}

	f, err := os.Create(outputPathFileName)

	if err != nil {
		return outputDto,
			fmt.Errorf(ePrefix+"Output File Create Error: %v ", err.Error())
	}

	var keys []int
	for k := range dtf.FormatMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	numOfKeys := 0
	numOfFormats := 0
	mapLen := 0

	_, err = f.WriteString("Length - Number Of Formats\n")

	if err != nil {
		_ = f.Close()
		return outputDto,
			fmt.Errorf(ePrefix+
				"Error Writing Header To Output File! Error: '%v'", err.Error())
	}

	for _, k := range keys {

		numOfKeys++
		mapLen = len(dtf.FormatMap[k])
		numOfFormats += mapLen

		_, err = f.WriteString(fmt.Sprintf("%06d%18d\n", k, mapLen))

		if err != nil {
			_ = f.Close()
			return outputDto,
				fmt.Errorf("Error Writing Stats to Output File! mapLen=%v Error: %v", mapLen, err.Error())
		}

	}

	outputDto.FileWriteEndTime = time.Now()
	du := DurationTriad{}
	err = du.SetStartEndTimes(
		outputDto.FileWriteStartTime,
		outputDto.FileWriteEndTime,
		TDurCalc.StdYearMth(),
		"Local",
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		_ = f.Close()
		return outputDto,
			fmt.Errorf("Error Calculating Duration with SetStartEndTimesTz() Error: %v", err.Error())
	}

	outputDto.ElapsedTimeForFileWriteOps = du.BaseTime.GetYearMthDaysTimeStr()
	outputDto.NumberOfFormatsGenerated = numOfFormats
	outputDto.NumberOfFormatMapKeysGenerated = numOfKeys

	err = f.Close()

	if err != nil {
		return outputDto,
			fmt.Errorf(ePrefix + "Error returned by f.Close()")
	}

	return outputDto, nil
}

// Empty - Sets all data fields in current FormatDateTimeUtility
// to their uninitialized or 'Empty' state.
func (dtf *FormatDateTimeUtility) Empty() {
	dtf.OriginalDateTimeStringIn = ""
	dtf.FormattedDateTimeStringIn = ""
	dtf.DateTimeOut = time.Time{}
	dtf.SelectedFormat = ""
	dtf.SelectedFormatSource = ""
	dtf.SelectedMapIdx = -1
	dtf.DictSearches = make([][][]int, 0)
	dtf.TotalNoOfDictSearches = 0

}

// ParseDateTimeString - Parses date time strings passed into the method. If a
// format is passed to the this method as the second parameter, the method will
// use this format first in an attempt to convert the date time string to a time.Time
// structure. If the 'probableFormat' parameter is empty or if it fails to convert
// the time string to a valid time.Time value, this method will run the date time
// string against 1.4-million possible date time string formats in an effort to
// successfully convert the date time string into a valid time.Time value.
func (dtf *FormatDateTimeUtility) ParseDateTimeString(dateTimeStr string, probableFormat string) (time.Time, error) {

	ePrefix := "FormatDateTimeUtility.ParseDateTimeString() "

	if dateTimeStr == "" {
		return time.Time{}, errors.New("Empty Time String!")
	}

	dtf.Empty()

	xtimeStr := dtf.replaceMultipleStrSequence(dateTimeStr, dtf.FormatSearchReplaceStrs.PreTrimSearchStrs)
	xtimeStr = dtf.replaceDateSuffixStThRd(xtimeStr)
	xtimeStr = dtf.reformatSingleTimeString(xtimeStr, dtf.FormatSearchReplaceStrs.TimeFmtRegEx)
	xtimeStr = dtf.replaceAMPM(xtimeStr)

	ftimeStr, err := dtf.trimEndMultiple(xtimeStr, ' ')

	if err != nil {
		return time.Time{}, err
	}

	if probableFormat != "" {
		t, err := time.Parse(probableFormat, ftimeStr)

		if err == nil {
			dtf.SelectedFormat = probableFormat
			dtf.SelectedFormatSource = "User Provided"
			dtf.SelectedMapIdx = -1
			dtf.DateTimeOut = t
			dtf.OriginalDateTimeStringIn = dateTimeStr
			dtf.FormattedDateTimeStringIn = ftimeStr
			dtf.TotalNoOfDictSearches = 1
			dtf.DictSearches = append(dtf.DictSearches, [][]int{{0, 1}})

			return t, nil
		}

	}

	if len(dtf.FormatMap) == 0 {

		err = dtf.CreateAllFormatsInMemory()

		if err != nil {
			return time.Time{},
				fmt.Errorf(ePrefix+"Error returned by dtf.CreateAllFormatsInMemory(). "+
					"Error='%v' ", err.Error())
		}

		if len(dtf.FormatMap) == 0 {
			return time.Time{},
				errors.New(ePrefix +
					"Format Map is EMPTY! Load Formats into FormatDateTimeUtility.FormatMap first!")
		}

	}

	lenStr := len(ftimeStr)

	lenSequence := make([][]int, 0)

	lenTests := []int{
		lenStr - 2,
		lenStr - 3,
		lenStr - 1,
		lenStr,
		lenStr + 1,
		lenStr + 2,
		lenStr + 3,
		lenStr - 4,
		lenStr + 4,
		lenStr - 5,
		lenStr + 5,
		lenStr - 6,
		lenStr + 6,
		lenStr - 7,
		lenStr + 7,
		lenStr - 8,
		lenStr + 8,
		lenStr - 9,
		lenStr + 9,
	}

	dtf.TotalNoOfDictSearches = 0
	dtf.OriginalDateTimeStringIn = dateTimeStr
	dtf.FormattedDateTimeStringIn = ftimeStr
	dtf.DateTimeOut = time.Time{}

	threshold := 5

	if lenStr >= 30 {
		threshold = 8
	}

	ary := make([]int, 0)
	for i := 0; i < len(lenTests); i++ {

		if dtf.FormatMap[lenTests[i]] != nil {
			ary = append(ary, lenTests[i])
		}

		if len(ary) == threshold {
			lenSequence = append(lenSequence, ary)
			ary = make([]int, 0)
		}

	}

	if len(lenSequence) > 0 {
		lenSequence = append(lenSequence, ary)
	}

	for j := 0; j < len(lenSequence); j++ {

		if dtf.doParseRun(lenSequence[j], ftimeStr) {
			return dtf.DateTimeOut, nil
		}

	}

	return time.Time{},
		errors.New(ePrefix + "Failed to locate correct time format!")
}

func (dtf *FormatDateTimeUtility) doParseRun(lenTests []int, ftimeStr string) bool {

	lenLenTests := len(lenTests)
	msg := make(chan parseDateTimeDto)
	done := uint64(0)

	isSuccessfulParse := false

	for i := 0; i < lenLenTests; i++ {

		go dtf.parseFormatMap(msg, &done, ftimeStr, lenTests[i], dtf.FormatMap[lenTests[i]])

	}

	cnt := 0

	for m := range msg {
		cnt++

		if cnt == lenLenTests {
			close(msg)
			runtime.Gosched()
		}

		dtf.DictSearches =
			append(dtf.DictSearches, [][]int{{m.selectedMapIdx, m.totalNoOfDictSearches}})
		dtf.TotalNoOfDictSearches += m.totalNoOfDictSearches

		if m.isSuccessful && !isSuccessfulParse {
			isSuccessfulParse = true
			dtf.SelectedFormat = m.selectedFormat
			dtf.SelectedFormatSource = "Format Map Dictionary"
			dtf.SelectedMapIdx = m.selectedMapIdx
			dtf.DateTimeOut = m.dateTimeOut

		}

	}

	return isSuccessfulParse
}

func (dtf *FormatDateTimeUtility) parseFormatMap(
	msg chan<- parseDateTimeDto, done *uint64, timeStr string, idx int, fmtMap map[string]int) {

	var doneTest uint64

	dto := parseDateTimeDto{}

	dto.formattedDateTimeStringIn = timeStr
	dto.selectedMapIdx = idx

	for key := range fmtMap {

		dto.totalNoOfDictSearches++

		t, err := time.Parse(key, timeStr)

		doneTest = atomic.LoadUint64(done)

		if doneTest > 0 {
			msg <- dto
			return
		}

		if err == nil {
			atomic.AddUint64(done, 1)
			dto.dateTimeOut = t
			dto.selectedFormat = key
			dto.isSuccessful = true
			msg <- dto
			runtime.Gosched()
			return
		}

	}

	msg <- dto
	return

}

func (dtf *FormatDateTimeUtility) assembleDayMthYears() error {

	dtf.FormatMap = make(map[int]map[string]int)

	fmtStr := ""

	dayOfWeek, _ := dtf.getDayOfWeekElements()
	dayOfWeekSeparators, _ := dtf.getDayOfWeekSeparator()
	mthDayYearFmts, _ := dtf.getMonthDayYearElements()

	for _, dowk := range dayOfWeek {
		for _, dowkSep := range dayOfWeekSeparators {

			for _, mmddyyy := range mthDayYearFmts {

				if dowk != "" && mmddyyy != "" {
					fmtStr = dowk + dowkSep + mmddyyy
					dtf.assignFormatStrToMap(fmtStr)
				} else if dowk != "" && mmddyyy == "" {
					dtf.assignFormatStrToMap(dowk)
				} else if dowk == "" && mmddyyy != "" {
					dtf.assignFormatStrToMap(mmddyyy)
				}
			}
		}
	}

	return nil
}

func (dtf *FormatDateTimeUtility) assembleMthDayYearFmts() error {

	dayOfWeek, _ := dtf.getDayOfWeekElements()

	dayOfWeekSeparators, _ := dtf.getDayOfWeekSeparator()

	mthDayYearFmts, _ := dtf.getMonthDayYearElements()

	dateTimeSeparators, _ := dtf.getDateTimeSeparators()

	timeFmts, _ := dtf.getTimeElements()

	offsetSeparators, _ := dtf.getTimeOffsetSeparators()

	offsetFmts, _ := dtf.getTimeOffsets()

	tzSeparators, _ := dtf.getTimeZoneSeparators()

	timeZoneFmts, _ := dtf.getTimeZoneElements()

	for _, dowk := range dayOfWeek {
		for _, dowkSep := range dayOfWeekSeparators {
			for _, mmddyyyy := range mthDayYearFmts {
				for _, dtSep := range dateTimeSeparators {
					for _, t := range timeFmts {
						for _, tOffsetSep := range offsetSeparators {
							for _, offFmt := range offsetFmts {
								for _, stdSep := range tzSeparators {
									for _, tzF := range timeZoneFmts {
										fmtGen := formatDateTimeGenerator{
											dayOfWeek:          dowk,
											dayOfWeekSeparator: dowkSep,
											mthDayYear:         mmddyyyy,
											dateTimeSeparator:  dtSep,
											timeElement:        t,
											offsetSeparator:    tOffsetSep,
											offsetElement:      offFmt,
											timeZoneSeparator:  stdSep,
											timeZoneElement:    tzF,
										}

										dtf.analyzeDofWeekMMDDYYYYTimeOffsetTz(fmtGen)
									}
								}

							}
						}
					}
				}

			}
		}

	}

	return nil
}

func (dtf *FormatDateTimeUtility) assembleMthDayTimeOffsetTzYearFmts() error {

	dayOfWeek, _ := dtf.getDayOfWeekElements()

	dayOfWeekSeparators, _ := dtf.getDayOfWeekSeparator()

	mthDayElements, _ := dtf.getMonthDayElements()

	afterMthDaySeparators, _ := dtf.getAfterMthDaySeparators()

	timeFmts, _ := dtf.getTimeElements()

	offsetSeparators, _ := dtf.getTimeOffsetSeparators()

	offsetFmts, _ := dtf.getTimeOffsets()

	tzSeparators, _ := dtf.getTimeZoneSeparators()

	timeZoneFmts, _ := dtf.getTimeZoneElements()

	yearElements, _ := dtf.getYears()

	for _, dowk := range dayOfWeek {
		for _, dowkSep := range dayOfWeekSeparators {
			for _, mthDay := range mthDayElements {
				for _, afterMthDaySeparator := range afterMthDaySeparators {
					for _, t := range timeFmts {
						for _, tOffsetSep := range offsetSeparators {
							for _, offFmt := range offsetFmts {
								for _, stdSep := range tzSeparators {
									for _, tzF := range timeZoneFmts {
										for _, yearEle := range yearElements {

											fmtGen := formatDateTimeGenerator{
												dayOfWeek:            dowk,
												dayOfWeekSeparator:   dowkSep,
												mthDayYear:           mthDay,
												afterMthDaySeparator: afterMthDaySeparator,
												timeElement:          t,
												offsetSeparator:      tOffsetSep,
												offsetElement:        offFmt,
												timeZoneSeparator:    stdSep,
												timeZoneElement:      tzF,
												year:                 yearEle,
											}

											dtf.analyzeDofWeekMMDDTimeOffsetTzYYYY(fmtGen)

										}
									}
								}

							}
						}
					}
				}

			}
		}

	}

	return nil
}

func (dtf *FormatDateTimeUtility) analyzeDofWeekMMDDYYYYTimeOffsetTz(dtfGen formatDateTimeGenerator) {

	fmtStr := ""
	fmtStr2 := ""

	if dtfGen.mthDayYear == "" &&
		dtfGen.timeElement == "" {
		return
	}

	if dtfGen.dayOfWeek != "" {
		fmtStr += dtfGen.dayOfWeek
	}

	if dtfGen.mthDayYear != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.mthDayYear
		} else {
			fmtStr += dtfGen.dayOfWeekSeparator
			fmtStr += dtfGen.mthDayYear
		}
	}

	if dtfGen.timeElement != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.timeElement
		} else {
			fmtStr += dtfGen.dateTimeSeparator
			fmtStr += dtfGen.timeElement
		}
	}

	fmtStr2 = fmtStr

	if dtfGen.offsetElement != "" &&
		fmtStr != "" &&
		dtfGen.timeElement != "" {

		fmtStr += dtfGen.offsetSeparator
		fmtStr += dtfGen.offsetElement
	}

	if dtfGen.timeZoneElement != "" &&
		fmtStr != "" &&
		dtfGen.timeElement != "" {

		fmtStr += dtfGen.timeZoneSeparator
		fmtStr += dtfGen.timeZoneElement
	}

	if fmtStr != "" {
		dtf.assignFormatStrToMap(fmtStr)
	}

	// Calculate variation of format string where
	// Time Zone comes before Offset Element

	if dtfGen.timeZoneElement != "" &&
		dtfGen.timeElement != "" &&
		fmtStr2 != "" &&
		dtfGen.offsetElement != "" {

		fmtStr2 += dtfGen.timeZoneSeparator
		fmtStr2 += dtfGen.timeZoneElement
		fmtStr2 += dtfGen.offsetSeparator
		fmtStr2 += dtfGen.offsetElement
	}

	if fmtStr2 != "" {
		dtf.assignFormatStrToMap(fmtStr2)
	}

	return

}

func (dtf *FormatDateTimeUtility) assemblePreDefinedFormats() {

	preDefFmts := dtf.getPredefinedFormats()

	for _, pdf := range preDefFmts {

		dtf.assignFormatStrToMap(pdf)

	}

}

func (dtf *FormatDateTimeUtility) assembleEdgeCaseFormats() {
	edgeCases := dtf.getEdgeCases()

	for _, ecf := range edgeCases {
		dtf.assignFormatStrToMap(ecf)
	}
}

func (dtf *FormatDateTimeUtility) analyzeDofWeekMMDDTimeOffsetTzYYYY(dtfGen formatDateTimeGenerator) {

	fmtStr := ""
	fmtStr2 := ""

	if dtfGen.dayOfWeek != "" {
		fmtStr += dtfGen.dayOfWeek
	}

	if dtfGen.mthDay != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.mthDay
		} else {
			fmtStr += dtfGen.dayOfWeekSeparator
			fmtStr += dtfGen.mthDay
		}
	}

	if dtfGen.timeElement != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.timeElement
		} else {
			fmtStr += dtfGen.afterMthDaySeparator
			fmtStr += dtfGen.timeElement
		}
	}

	fmtStr2 = fmtStr

	if dtfGen.offsetElement != "" &&
		fmtStr != "" &&
		dtfGen.timeElement != "" {
		fmtStr += dtfGen.offsetSeparator
		fmtStr += dtfGen.offsetElement

	}

	if dtfGen.timeZoneElement != "" &&
		fmtStr != "" &&
		dtfGen.timeElement != "" {
		fmtStr += dtfGen.timeZoneSeparator
		fmtStr += dtfGen.timeZoneElement
	}

	if fmtStr != "" {
		dtf.assignFormatStrToMap(fmtStr)
	}

	// Calculate variation of format string where
	// Time Zone comes before Offset Element

	if dtfGen.timeZoneElement != "" &&
		fmtStr2 != "" &&
		dtfGen.timeElement != "" {
		fmtStr2 += dtfGen.timeZoneSeparator
		fmtStr2 += dtfGen.timeZoneElement
	}

	if dtfGen.offsetElement != "" &&
		fmtStr2 != "" &&
		dtfGen.timeElement != "" {
		fmtStr2 += dtfGen.offsetSeparator
		fmtStr2 += dtfGen.offsetElement

	}

	if fmtStr2 != "" {
		dtf.assignFormatStrToMap(fmtStr)
	}

	return
}

func (dtf *FormatDateTimeUtility) assignFormatStrToMap(fmtStr string) {

	l := len(fmtStr)

	if l == 0 {
		return
	}

	if dtf.FormatMap[l] == nil {
		dtf.FormatMap[l] = make(map[string]int)
	}

	if dtf.FormatMap[l][fmtStr] != 0 {
		return
	}

	dtf.FormatMap[l][fmtStr] = l
	dtf.NumOfFormatsGenerated++
}

func (dtf *FormatDateTimeUtility) getDayOfWeekElements() ([]string, error) {
	dayOfWeek := make([]string, 0, 10)

	dayOfWeek = append(dayOfWeek, "")
	dayOfWeek = append(dayOfWeek, "Mon")
	dayOfWeek = append(dayOfWeek, "Monday")

	return dayOfWeek, nil
}

func (dtf *FormatDateTimeUtility) getDayOfWeekSeparator() ([]string, error) {
	dayOfWeekSeparator := make([]string, 0, 1024)

	dayOfWeekSeparator = append(dayOfWeekSeparator, " ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, ", ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, " - ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, "-")

	return dayOfWeekSeparator, nil
}

func (dtf *FormatDateTimeUtility) getMonthDayYearElements() ([]string, error) {
	mthDayYr := make([]string, 0, 1024)

	mthDayYr = append(mthDayYr, "2006-1-2")
	mthDayYr = append(mthDayYr, "2006 1 2")
	mthDayYr = append(mthDayYr, "1-2-06")
	mthDayYr = append(mthDayYr, "1-2-2006")

	mthDayYr = append(mthDayYr, "1 2 06")
	mthDayYr = append(mthDayYr, "1 2 2006")

	mthDayYr = append(mthDayYr, "Jan-2-06")
	mthDayYr = append(mthDayYr, "Jan 2 06")
	mthDayYr = append(mthDayYr, "Jan _2 06")
	mthDayYr = append(mthDayYr, "Jan-2-2006")
	mthDayYr = append(mthDayYr, "Jan 2 2006")
	mthDayYr = append(mthDayYr, "Jan _2 2006")

	mthDayYr = append(mthDayYr, "January-2-06")
	mthDayYr = append(mthDayYr, "January 2 06")
	mthDayYr = append(mthDayYr, "January _2 06")
	mthDayYr = append(mthDayYr, "January-2-2006")
	mthDayYr = append(mthDayYr, "January 2 2006")
	mthDayYr = append(mthDayYr, "January _2 2006")

	// European Date Formats DD.MM.YYYY
	mthDayYr = append(mthDayYr, "2.1.06")
	mthDayYr = append(mthDayYr, "2.1.2006")
	mthDayYr = append(mthDayYr, "2.1.'06")

	// Standard Dates with Dot Delimiters
	mthDayYr = append(mthDayYr, "2006.1.2")

	mthDayYr = append(mthDayYr, "2-January-2006")
	mthDayYr = append(mthDayYr, "2-January-06")
	mthDayYr = append(mthDayYr, "2 January 06")
	mthDayYr = append(mthDayYr, "2 January 2006")

	mthDayYr = append(mthDayYr, "2-Jan-2006")
	mthDayYr = append(mthDayYr, "2-Jan-06")
	mthDayYr = append(mthDayYr, "2 Jan 06")
	mthDayYr = append(mthDayYr, "2 Jan 2006")

	mthDayYr = append(mthDayYr, "20060102")
	mthDayYr = append(mthDayYr, "01022006")
	mthDayYr = append(mthDayYr, "")

	return mthDayYr, nil
}

func (dtf *FormatDateTimeUtility) getMonthDayElements() ([]string, error) {
	mthDayElements := make([]string, 0, 124)

	mthDayElements = append(mthDayElements, "Jan 2")
	mthDayElements = append(mthDayElements, "January 2")
	mthDayElements = append(mthDayElements, "Jan _2")
	mthDayElements = append(mthDayElements, "January _2")
	mthDayElements = append(mthDayElements, "1-2")
	mthDayElements = append(mthDayElements, "1-_2")
	mthDayElements = append(mthDayElements, "1 2")
	mthDayElements = append(mthDayElements, "1-2")

	mthDayElements = append(mthDayElements, "01 02")
	mthDayElements = append(mthDayElements, "01 _2")
	mthDayElements = append(mthDayElements, "01-02")

	mthDayElements = append(mthDayElements, "0102")
	// European Format Day Month
	mthDayElements = append(mthDayElements, "02.01")
	mthDayElements = append(mthDayElements, "2.1")
	mthDayElements = append(mthDayElements, "02.1")
	mthDayElements = append(mthDayElements, "2.01")

	return mthDayElements, nil
}

func (dtf *FormatDateTimeUtility) getYears() ([]string, error) {
	yearElements := make([]string, 0, 10)

	yearElements = append(yearElements, "2006")
	yearElements = append(yearElements, "06")
	yearElements = append(yearElements, "'06")

	return yearElements, nil
}

func (dtf *FormatDateTimeUtility) getAfterMthDaySeparators() ([]string, error) {
	mthDayAfterSeparators := make([]string, 0, 10)

	mthDayAfterSeparators = append(mthDayAfterSeparators, " ")
	mthDayAfterSeparators = append(mthDayAfterSeparators, ", ")
	mthDayAfterSeparators = append(mthDayAfterSeparators, ":")
	mthDayAfterSeparators = append(mthDayAfterSeparators, "T")
	mthDayAfterSeparators = append(mthDayAfterSeparators, "")

	return mthDayAfterSeparators, nil

}

func (dtf *FormatDateTimeUtility) getStandardSeparators() ([]string, error) {
	standardSeparators := make([]string, 0, 10)

	standardSeparators = append(standardSeparators, " ")
	standardSeparators = append(standardSeparators, "")

	return standardSeparators, nil
}

func (dtf *FormatDateTimeUtility) getDateTimeSeparators() ([]string, error) {
	dtTimeSeparators := make([]string, 0, 10)

	dtTimeSeparators = append(dtTimeSeparators, " ")
	dtTimeSeparators = append(dtTimeSeparators, ":")
	dtTimeSeparators = append(dtTimeSeparators, "T")
	dtTimeSeparators = append(dtTimeSeparators, "")

	return dtTimeSeparators, nil
}

func (dtf *FormatDateTimeUtility) getTimeElements() ([]string, error) {
	timeElements := make([]string, 0, 512)

	timeElements = append(timeElements, "15:04:05")
	timeElements = append(timeElements, "15:04")
	timeElements = append(timeElements, "15:04:05.000")
	timeElements = append(timeElements, "15:04:05.000000")
	timeElements = append(timeElements, "15:04:05.000000000")

	timeElements = append(timeElements, "03:04:05 pm")
	timeElements = append(timeElements, "03:04 pm")
	timeElements = append(timeElements, "03:04:05.000 pm")
	timeElements = append(timeElements, "03:04:05.000000 pm")
	timeElements = append(timeElements, "03:04:05.000000000 pm")

	timeElements = append(timeElements, "")

	return timeElements, nil
}

func (dtf *FormatDateTimeUtility) getTimeOffsets() ([]string, error) {
	timeOffsetElements := make([]string, 0, 20)

	timeOffsetElements = append(timeOffsetElements, "-0700")
	timeOffsetElements = append(timeOffsetElements, "-07:00")
	timeOffsetElements = append(timeOffsetElements, "-07")
	timeOffsetElements = append(timeOffsetElements, "Z0700")
	timeOffsetElements = append(timeOffsetElements, "Z07:00")
	timeOffsetElements = append(timeOffsetElements, "Z07")
	timeOffsetElements = append(timeOffsetElements, "")

	return timeOffsetElements, nil
}

func (dtf *FormatDateTimeUtility) getTimeOffsetSeparators() ([]string, error) {
	timeOffsetSeparators := make([]string, 0, 20)

	timeOffsetSeparators = append(timeOffsetSeparators, " ")
	timeOffsetSeparators = append(timeOffsetSeparators, "-")
	timeOffsetSeparators = append(timeOffsetSeparators, "")

	return timeOffsetSeparators, nil
}

func (dtf *FormatDateTimeUtility) getTimeZoneElements() ([]string, error) {
	tzElements := make([]string, 0, 20)

	tzElements = append(tzElements, "MST")
	tzElements = append(tzElements, "-07")
	tzElements = append(tzElements, "-0700")
	tzElements = append(tzElements, "")

	return tzElements, nil
}

func (dtf *FormatDateTimeUtility) getTimeZoneSeparators() ([]string, error) {
	tzElements := make([]string, 0, 20)

	tzElements = append(tzElements, " ")
	tzElements = append(tzElements, "-")
	tzElements = append(tzElements, "")

	return tzElements, nil
}

func (dtf *FormatDateTimeUtility) getPredefinedFormats() []string {

	preDefinedFormats := make([]string, 0, 20)

	preDefinedFormats = append(preDefinedFormats, time.ANSIC)
	preDefinedFormats = append(preDefinedFormats, time.UnixDate)
	preDefinedFormats = append(preDefinedFormats, time.RubyDate)
	preDefinedFormats = append(preDefinedFormats, time.RFC822)
	preDefinedFormats = append(preDefinedFormats, time.RFC822Z)
	preDefinedFormats = append(preDefinedFormats, time.RFC850)
	preDefinedFormats = append(preDefinedFormats, time.RFC1123)
	preDefinedFormats = append(preDefinedFormats, time.RFC1123Z)
	preDefinedFormats = append(preDefinedFormats, time.RFC3339)
	preDefinedFormats = append(preDefinedFormats, time.RFC3339Nano)
	preDefinedFormats = append(preDefinedFormats, time.Kitchen)
	preDefinedFormats = append(preDefinedFormats, time.Stamp)
	preDefinedFormats = append(preDefinedFormats, time.StampMilli)
	preDefinedFormats = append(preDefinedFormats, time.StampMicro)
	preDefinedFormats = append(preDefinedFormats, time.StampNano)

	return preDefinedFormats
}

func (dtf *FormatDateTimeUtility) getEdgeCases() []string {
	// FmtDateTimeEverything = "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
	edgeCases := make([]string, 0, 20)

	edgeCases = append(edgeCases, "2006-01-02 15:04:05.000000000 -0700 -07")
	edgeCases = append(edgeCases, "01/02/2006 15:04:05.000000000 -0700 -07")
	edgeCases = append(edgeCases, "01-02-2006 15:04:05.000000000 -0700 -07")
	edgeCases = append(edgeCases, "01.02.2006 15:04:05.000000000 -0700 -07")

	edgeCases = append(edgeCases, "Monday January 2 15:04:05 -0700 MST 2006")

	edgeCases = append(edgeCases, "Mon January 2 15:04:05 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2 15:04:05 -0700 MST 2006")
	edgeCases = append(edgeCases, "January 2 15:04:05 -0700 MST 2006")

	edgeCases = append(edgeCases, "Monday January 2 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "January 2 15:04 -0700 MST 2006")

	edgeCases = append(edgeCases, "January 2 03:04 pm -0700 MST 2006")
	edgeCases = append(edgeCases, "January 2 03:04:05 pm -0700 MST 2006")

	edgeCases = append(edgeCases, "15:04:05 -0700 MST")
	edgeCases = append(edgeCases, "3:04:05 pm -0700 MST")
	edgeCases = append(edgeCases, "15:04 -0700 MST")
	edgeCases = append(edgeCases, "3:04 pm -0700 MST")

	edgeCases = append(edgeCases, "15:04:05 -0700")
	edgeCases = append(edgeCases, "3:04:05 pm -0700")
	edgeCases = append(edgeCases, "15:04 -0700")
	edgeCases = append(edgeCases, "3:04 pm -0700")

	edgeCases = append(edgeCases, "15:04:05")
	edgeCases = append(edgeCases, "3:04:05 pm")
	edgeCases = append(edgeCases, "15:04")
	edgeCases = append(edgeCases, "3:04 pm")

	return edgeCases
}

func (dtf *FormatDateTimeUtility) getPreTrimSearchStrings() [][][]string {
	d := make([][][]string, 0)
	d = append(d, [][]string{{",", " ", "-1"}})
	d = append(d, [][]string{{"/", " ", "-1"}})
	d = append(d, [][]string{{"\\", " ", "-1"}})
	d = append(d, [][]string{{"*", " ", "-1"}})
	d = append(d, [][]string{{"-hrs", ":", "1"}})
	d = append(d, [][]string{{"-mins", ":", "1"}})
	d = append(d, [][]string{{"-secs", "", "1"}})
	d = append(d, [][]string{{"-min", ":", "1"}})
	d = append(d, [][]string{{"-sec", "", "1"}})

	d = append(d, [][]string{{"-Hrs", ":", "1"}})
	d = append(d, [][]string{{"-Mins", ":", "1"}})
	d = append(d, [][]string{{"-Secs", "", "1"}})
	d = append(d, [][]string{{"-Min", ":", "1"}})
	d = append(d, [][]string{{"-Sec", "", "1"}})

	return d
}

func (dtf *FormatDateTimeUtility) getTimeFmtRegEx() [][][]string {
	d := make([][][]string, 0)
	d = append(d, [][]string{{"\\d\\d:\\d\\d:\\d\\d", "%02d:%02d:%02d"}}) // 2:2:2
	d = append(d, [][]string{{"\\d\\d:\\d:\\d\\d", "%02d:%02d:%02d"}})    // 2:1:2
	d = append(d, [][]string{{"\\d\\d:\\d\\d:\\d", "%02d:%02d:%02d"}})    // 2:2:1
	d = append(d, [][]string{{"\\d\\d:\\d:\\d", "%02d:%02d:%02d"}})       // 2:1:1
	d = append(d, [][]string{{"\\d:\\d\\d:\\d\\d", "%02d:%02d:%02d"}})    // 1:2:2
	d = append(d, [][]string{{"\\d:\\d:\\d\\d", "%02d:%02d:%02d"}})       // 1:1:2
	d = append(d, [][]string{{"\\d:\\d\\d:\\d", "%02d:%02d:%02d"}})       // 1:2:1
	d = append(d, [][]string{{"\\d:\\d:\\d", "%02d:%02d:%02d"}})          // 1:1:1

	/*

	   1-	1	1	1
	   2-	1	1	2
	   3-	1	2	2
	   4-	1	2	1
	   5-	2	2	2
	   6-	2	1	1
	   7-	2	2	1
	   8-	2	1	2

	*/

	d = append(d, [][]string{{"\\d\\d:\\d\\d", "%02d:%02d"}}) // 2:2
	d = append(d, [][]string{{"\\d\\d:\\d", "%02d:%02d"}})    // 2:1
	d = append(d, [][]string{{"\\d:\\d\\d", "%02d:%02d"}})    // 1:2
	d = append(d, [][]string{{"\\d:\\d", "%02d:%02d"}})       // 1:1

	/*
	   1- 1:1
	   2- 1:2
	   3- 2:1
	   4- 2:2

	*/

	return d
}

func (dtf *FormatDateTimeUtility) reformatSingleTimeString(targetStr string, regExes [][][]string) string {

	max := len(regExes)

	for i := 0; i < max; i++ {
		re := regexp.MustCompile(regExes[i][0][0])

		idx := re.FindStringIndex(targetStr)

		if idx == nil {
			continue
		}

		s := []byte(targetStr)

		sExtract := string(s[idx[0]:idx[1]])

		timeElements := strings.Split(sExtract, ":")

		lTElements := len(timeElements)

		replaceStr := ""

		for i := 0; i < lTElements; i++ {

			iE, err := strconv.Atoi(timeElements[i])

			if err != nil {
				panic(fmt.Errorf("reformatSingleTimeString() Error converint Time Element %v to ASCII. Error- %v", i, err.Error()))
			}

			if i > 0 {
				replaceStr += ":"
			}

			replaceStr += fmt.Sprintf("%02d", iE)
		}

		return strings.Replace(targetStr, sExtract, replaceStr, 1)

	}

	return targetStr
}

func (dtf *FormatDateTimeUtility) replaceMultipleStrSequence(targetStr string, replaceMap [][][]string) string {

	max := len(replaceMap)

	for i := 0; i < max; i++ {
		if strings.Contains(targetStr, replaceMap[i][0][0]) {
			instances, err := strconv.Atoi(replaceMap[i][0][2])

			if err != nil {
				instances = 1
			}

			targetStr = strings.Replace(targetStr, replaceMap[i][0][0], replaceMap[i][0][1], instances)
		}

	}

	return targetStr
}

func (dtf *FormatDateTimeUtility) replaceAMPM(targetStr string) string {
	d := make([][][]string, 0)

	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)a[.]*\\s{0,4}(?i)m[.]*", " am "}})
	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)p[.]*\\s{0,4}(?i)m[.]*", " pm "}})

	lD := len(d)

	for i := 0; i < lD; i++ {
		r, err := regexp.Compile(d[i][0][0])

		if err != nil {
			panic(fmt.Errorf("replaceAMPM() Regex failed to Compile. regex== %v. Error: %v", d[i][0][0], err.Error()))
		}

		bTargetStr := []byte(targetStr)

		loc := r.FindIndex(bTargetStr)

		if loc == nil {
			continue
		}

		// Found regex expression

		foundEx := string(bTargetStr[loc[0]+1 : loc[1]])

		return strings.Replace(targetStr, foundEx, d[i][0][1], 1)

	}

	return targetStr
}

func (dtf *FormatDateTimeUtility) replaceDateSuffixStThRd(targetStr string) string {
	// \d{1}\s{0,4}(?i)t\s{0,4}(?i)h
	d := make([][][]string, 0)

	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)s\\s{0,4}(?i)t", " "}})
	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)n\\s{0,4}(?i)d", " "}})
	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)r\\s{0,4}(?i)d", " "}})
	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)t\\s{0,4}(?i)h", " "}})

	lD := len(d)

	for i := 0; i < lD; i++ {
		r, err := regexp.Compile(d[i][0][0])

		if err != nil {
			panic(fmt.Errorf("replaceDateSuffixStThRd() Regex failed to Compile. regex== %v. Error: %v", d[i][0][0], err.Error()))
		}

		bTargetStr := []byte(targetStr)

		loc := r.FindIndex(bTargetStr)

		if loc == nil {
			continue
		}

		// Found regex expression

		foundEx := string(bTargetStr[loc[0]+1 : loc[1]])

		return strings.Replace(targetStr, foundEx, d[i][0][1], 1)

	}

	return targetStr
}

// TrimEndMultiple- Performs the following operations on strings:
// 1. Trims Right and Left for all instances of 'trimChar'
// 2. Within the interior of a string, multiple instances
// 		of 'trimChar' are reduce to a single instance.
func (dtf *FormatDateTimeUtility) trimEndMultiple(targetStr string, trimChar rune) (rStr string, err error) {

	if targetStr == "" {
		err = errors.New("trimEndMultiple() - Empty targetStr")
		return
	}

	fStr := []rune(targetStr)
	lenTargetStr := len(fStr)
	outputStr := make([]rune, lenTargetStr)
	lenTargetStr--
	idx := lenTargetStr
	foundFirstChar := false

	for i := lenTargetStr; i >= 0; i-- {

		if !foundFirstChar && fStr[i] == trimChar {
			continue
		}

		if i > 0 && fStr[i] == trimChar && fStr[i-1] == trimChar {
			continue
		}

		if i == 0 && fStr[i] == trimChar {
			continue
		}

		foundFirstChar = true
		outputStr[idx] = fStr[i]
		idx--
	}

	if idx != lenTargetStr {
		idx++
	}

	if outputStr[idx] == trimChar {
		idx++
	}

	result := string(outputStr[idx:])

	return result, nil

}
