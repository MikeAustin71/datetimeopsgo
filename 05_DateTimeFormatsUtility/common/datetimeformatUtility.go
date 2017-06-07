package common

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type DateTimeWriteFormatsToFileDto struct {
	OutputPathFileName             string
	NumberOfFormatsGenerated       int
	NumberOfFormatMapKeysGenerated int
	FileWriteStartTime             time.Time
	FileWriteEndTime               time.Time
	ElapsedTimeForFileWriteOps     string
}

type DateTimeReadFormatsFromFileDto struct {
	PathFileName                   string
	NumberOfFormatsGenerated       int
	NumberOfFormatMapKeysGenerated int
	FileReadStartTime              time.Time
	FileReadEndTime                time.Time
	ElapsedTimeForFileReadOps      string
}

type DateTimeFormatRecord struct {
	FmtLength int
	FormatStr string
}

type DateTimeFormatGenerator struct {
	DayOfWeek            string
	DayOfWeekSeparator   string
	MthDay               string
	MthDayYear           string
	AfterMthDaySeparator string
	DateTimeSeparator    string
	TimeElement          string
	OffsetSeparator      string
	OffsetElement        string
	TimeZoneSeparator    string
	TimeZoneElement      string
	Year                 string
}

type SearchStrings struct {
	FirstSecondThirdSearchStrs [][][]string
	PreTrimSearchStrs          [][][]string
	TimeFmtRegEx               [][][]string
}

type DateTimeFormatUtility struct {
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
// the field: DateTimeFormatUtility.FormatMap. This process currently
// takes about 1.5-seconds on my machine.
//
// 54-map keys are currently generated and used to access the format
// strings. Access is keyed on length of the date time string one is
// attempting to parse.
func (dtf *DateTimeFormatUtility) CreateAllFormatsInMemory() (err error) {

	dtf.FormatMap = make(map[int]map[string]int)
	dtf.NumOfFormatsGenerated = 0
	dtf.assemblePreDefinedFormats()
	dtf.assembleMthDayYearFmts()
	dtf.assembleEdgeCaseFormats()
	dtf.FormatSearchReplaceStrs.FirstSecondThirdSearchStrs = dtf.getFirstSecondThirdSearchStrs()
	dtf.FormatSearchReplaceStrs.PreTrimSearchStrs = dtf.getPreTrimSearchStrings()
	dtf.FormatSearchReplaceStrs.TimeFmtRegEx = dtf.getTimeFmtRegEx()

	return
}

// LoadAllFormatsFromFileIntoMemory - Loads all date time formats from a specified
// text file into memory. The formats are stored in DateTimeFormatUtility.FormatMap.
// This is an alternative means of reading all available date time formats into
// memory so that they may be used to parse time strings. This method assumes that
// the text file containing the format strings was originally created by method
// DateTimeFormatUtility.WriteAllFormatsInMemoryToFile() which employs a specific
// fixed length format which can then be read back into memory.
func (dtf *DateTimeFormatUtility) LoadAllFormatsFromFileIntoMemory(pathFileName string) (DateTimeReadFormatsFromFileDto, error) {

	frDto := DateTimeReadFormatsFromFileDto{}
	frDto.PathFileName = pathFileName
	frDto.FileReadStartTime = time.Now()
	dtf.FormatMap = make(map[int]map[string]int)
	dtf.NumOfFormatsGenerated = 0

	fmtFile, err := os.Open(pathFileName)

	if err != nil {
		return frDto, fmt.Errorf("LoadAllFormatsFromFileIntoMemory- Error Opening File: %v - Error: %v", pathFileName, err.Error())
	}

	defer fmtFile.Close()
	const bufLen int = 2000
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

			for j := 0; j < 7; j++ {

				lenField[j] = outRecordBuff[j]
			}

			s := string(lenField)

			lFmt, err := strconv.Atoi(s)

			if err != nil {
				return frDto, fmt.Errorf(
					"Error converting Format Length field from file. Length = %v. Idx= %v. Format Count: %v",
					s, idx, frDto.NumberOfFormatsGenerated)
			}

			fmtFieldLastIdx := 7 + lFmt

			if lOutBuff < fmtFieldLastIdx+1 {
				return frDto, fmt.Errorf(
					"Found corrupted Output Buffer. Buffer Length %v, Length Field = %v, Output Buffer= %s Format Count: %v",
					lOutBuff, lFmt, string(outRecordBuff), frDto.NumberOfFormatsGenerated)
			}

			fmtField := make([]byte, lFmt)

			for k := 8; k <= fmtFieldLastIdx; k++ {
				fmtField[k-8] = outRecordBuff[k]
			}

			fmtStr := string(fmtField)

			// Populate DateTimeFormatUtility.FormatMap
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
	du := DurationUtility{}
	etFileWrite, _ := du.GetElapsedTime(frDto.FileReadStartTime, frDto.FileReadEndTime)
	frDto.ElapsedTimeForFileReadOps = etFileWrite.DurationStr

	return frDto, nil
}

// WriteAllFormatsInMemoryToFile - Writes all Format Data contained in
// DateTimeFormatUtility.FormatMap field to a specified output file in
// text format. Currently, about 1.4-million formats are generated and
// written to the output file.
//
// IMPORTANT! - Before you call this method, the Format Maps must
// first be created in memory. Call DateTimeFormatUtility.CreateAllFormatsInMemory()
// first, before calling this method.
func (dtf *DateTimeFormatUtility) WriteAllFormatsInMemoryToFile(outputPathFileName string) (DateTimeWriteFormatsToFileDto, error) {

	fwDto := DateTimeWriteFormatsToFileDto{}

	fwDto.FileWriteStartTime = time.Now()
	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		return fwDto, errors.New("WriteAllFormatsInMemoryToFile() Error - There are NO Formats in Memory -  FormatMap length == 0. You MUST call CreateAllFormatsInMemory() first!")
	}

	outF, err := os.Create(outputPathFileName)

	if err != nil {
		return fwDto, fmt.Errorf("WriteAllFormatsInMemoryToFile() Error - Failed create output file %v. Error: %v", outputPathFileName, err.Error())
	}

	defer outF.Close()

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
				return fwDto, fmt.Errorf("WriteAllFormatsInMemoryToFile() Error writing Format data to output file %v. Error: %v", outputPathFileName, err.Error())
			}
		}
	}

	outF.Sync()

	du := DurationUtility{}

	fwDto.FileWriteEndTime = time.Now()

	etFileWrite, _ := du.GetElapsedTime(fwDto.FileWriteStartTime, fwDto.FileWriteEndTime)

	fwDto.OutputPathFileName = outputPathFileName
	fwDto.ElapsedTimeForFileWriteOps = etFileWrite.DurationStr

	return fwDto, nil
}

// WriteFormatStatsToFile - This method writes data to a text file.
// The text file is small, currently about 3-kilobytes in size.
// The data output to the text file describes the size of the
// slices contained in dtf.FormatMap.
// IMPORTANT! - Before you call this method, the Format Maps must
// first be created in memory. Call DateTimeFormatUtility.CreateAllFormatsInMemory()
// first, before calling this method.
func (dtf *DateTimeFormatUtility) WriteFormatStatsToFile(outputPathFileName string) (DateTimeWriteFormatsToFileDto, error) {
	outputDto := DateTimeWriteFormatsToFileDto{}
	outputDto.OutputPathFileName = outputPathFileName
	outputDto.FileWriteStartTime = time.Now()

	du := DurationUtility{}

	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		return outputDto,
			errors.New("WriteFormatStatsToFile() Error - There are NO Formats in Memory -  FormatMap length == 0. You MUST call CreateAllFormatsInMemory() first!")
	}

	f, err := os.Create(outputPathFileName)

	if err != nil {
		return outputDto, errors.New("Output File Create Error:" + err.Error())
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

	_, err = f.WriteString("Length - Number Of Formats\n")

	if err != nil {
		return outputDto, errors.New("Error Writing Header To Output File! Error: " + err.Error())
	}

	for _, k := range keys {

		numOfKeys++
		mapLen = len(dtf.FormatMap[k])
		numOfFormats += mapLen

		_, err = f.WriteString(fmt.Sprintf("%06d%18d\n", k, mapLen))

		if err != nil {
			return outputDto, fmt.Errorf("Error Writing Stats to Output File! mapLen=%v Error: %v", mapLen, err.Error())
		}

	}

	outputDto.FileWriteEndTime = time.Now()

	etFileWrite, _ := du.GetElapsedTime(outputDto.FileWriteStartTime, outputDto.FileWriteEndTime)

	outputDto.ElapsedTimeForFileWriteOps = etFileWrite.DurationStr
	outputDto.NumberOfFormatsGenerated = numOfFormats
	outputDto.NumberOfFormatMapKeysGenerated = numOfKeys

	return outputDto, nil
}

func (dtf *DateTimeFormatUtility) Empty() {
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
func (dtf *DateTimeFormatUtility) ParseDateTimeString(dateTimeStr string, probableFormat string) (time.Time, error) {

	if dateTimeStr == "" {
		return time.Time{}, errors.New("Empty Time String!")
	}

	dtf.Empty()

	xtimeStr := dtf.replaceMultipleStrSequence(dateTimeStr, dtf.FormatSearchReplaceStrs.PreTrimSearchStrs)

	ftimeStr, err := StringUtility{}.TrimEndMultiple(xtimeStr, ' ')

	ftimeStr = dtf.reformatSingleTimeString(ftimeStr, dtf.FormatSearchReplaceStrs.TimeFmtRegEx)
	ftimeStr = dtf.replaceAMPM(ftimeStr)
	ftimeStr = dtf.replaceDateTimeSequence(ftimeStr, dtf.FormatSearchReplaceStrs.FirstSecondThirdSearchStrs)

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

		dtf.CreateAllFormatsInMemory()

		if len(dtf.FormatMap) == 0 {
			return time.Time{}, errors.New("Format Map is EMPTY! Load Formats into DateTimeFormatUtility.FormatMap first!")
		}

	}

	lenStr := len(ftimeStr)

	lenTests := make([]int, 0)

	// lenTests  lenStr -3 through lenStr + 3

	if lenStr-2 > 0 {
		lenTests = append(lenTests, lenStr-2)
	}

	if lenStr-1 > 0 {
		lenTests = append(lenTests, lenStr-1)
	}

	lenTests = append(lenTests, lenStr+1)

	lenTests = append(lenTests, lenStr)

	if lenStr-3 > 0 {
		lenTests = append(lenTests, lenStr-3)
	}

	lenTests = append(lenTests, lenStr+2)

	lenTests = append(lenTests, lenStr+3)

	if lenStr-4 > 0 {
		lenTests = append(lenTests, lenStr-4)
	}

	lenTests = append(lenTests, lenStr+4)

	if lenStr-5 > 0 {
		lenTests = append(lenTests, lenStr-5)
	}

	lenTests = append(lenTests, lenStr+5)

	lenLenTests := len(lenTests)

	dtf.TotalNoOfDictSearches = 0

	for i := 0; i < lenLenTests; i++ {

		result, err := dtf.parseFormatMap(ftimeStr, lenTests[i])

		dtf.DictSearches =
			append(dtf.DictSearches, [][]int{{lenTests[i], result.TotalNoOfDictSearches}})
		dtf.TotalNoOfDictSearches += result.TotalNoOfDictSearches

		if err == nil {
			dtf.SelectedFormat = result.SelectedFormat
			dtf.SelectedFormatSource = result.SelectedFormatSource
			dtf.SelectedMapIdx = result.SelectedMapIdx
			dtf.DateTimeOut = result.DateTimeOut
			dtf.OriginalDateTimeStringIn = dateTimeStr
			dtf.FormattedDateTimeStringIn = result.FormattedDateTimeStringIn
			return dtf.DateTimeOut, nil
		}

	}

	dtf.OriginalDateTimeStringIn = dateTimeStr
	dtf.FormattedDateTimeStringIn = ftimeStr
	dtf.DateTimeOut = time.Time{}

	return time.Time{}, errors.New("Failed to locate correct time format!")
}

func (dtf *DateTimeFormatUtility) assembleDayMthYears() error {

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

func (dtf *DateTimeFormatUtility) parseFormatMap(timeStr string, idx int) (dtResult DateTimeFormatUtility, err error) {

	if dtf.FormatMap[idx] == nil {
		err = errors.New("Time String Length not found in Format Map!")
		return
	}

	for key := range dtf.FormatMap[idx] {

		dtResult.TotalNoOfDictSearches++

		t, err := time.Parse(key, timeStr)

		if err == nil {
			dtResult.OriginalDateTimeStringIn = timeStr
			dtResult.FormattedDateTimeStringIn = timeStr
			dtResult.SelectedMapIdx = idx
			dtResult.SelectedFormatSource = "Format Dictionary"
			dtResult.DateTimeOut = t
			dtResult.SelectedFormat = key
			return dtResult, err
		}

	}

	err = errors.New("Failed to parse time string")

	return
}

func (dtf *DateTimeFormatUtility) assembleMthDayYearFmts() error {

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
										fmtGen := DateTimeFormatGenerator{
											DayOfWeek:          dowk,
											DayOfWeekSeparator: dowkSep,
											MthDayYear:         mmddyyyy,
											DateTimeSeparator:  dtSep,
											TimeElement:        t,
											OffsetSeparator:    tOffsetSep,
											OffsetElement:      offFmt,
											TimeZoneSeparator:  stdSep,
											TimeZoneElement:    tzF,
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

func (dtf *DateTimeFormatUtility) assembleMthDayTimeOffsetTzYearFmts() error {

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

											fmtGen := DateTimeFormatGenerator{
												DayOfWeek:            dowk,
												DayOfWeekSeparator:   dowkSep,
												MthDayYear:           mthDay,
												AfterMthDaySeparator: afterMthDaySeparator,
												TimeElement:          t,
												OffsetSeparator:      tOffsetSep,
												OffsetElement:        offFmt,
												TimeZoneSeparator:    stdSep,
												TimeZoneElement:      tzF,
												Year:                 yearEle,
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

func (dtf *DateTimeFormatUtility) analyzeDofWeekMMDDYYYYTimeOffsetTz(dtfGen DateTimeFormatGenerator) {

	fmtStr := ""
	fmtStr2 := ""

	if dtfGen.MthDayYear == "" &&
		dtfGen.TimeElement == "" {
		return
	}

	if dtfGen.DayOfWeek != "" {
		fmtStr += dtfGen.DayOfWeek
	}

	if dtfGen.MthDayYear != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.MthDayYear
		} else {
			fmtStr += dtfGen.DayOfWeekSeparator
			fmtStr += dtfGen.MthDayYear
		}
	}

	if dtfGen.TimeElement != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.TimeElement
		} else {
			fmtStr += dtfGen.DateTimeSeparator
			fmtStr += dtfGen.TimeElement
		}
	}

	fmtStr2 = fmtStr

	if dtfGen.OffsetElement != "" &&
		fmtStr != "" &&
		dtfGen.TimeElement != "" {

		fmtStr += dtfGen.OffsetSeparator
		fmtStr += dtfGen.OffsetElement
	}

	if dtfGen.TimeZoneElement != "" &&
		fmtStr != "" &&
		dtfGen.TimeElement != "" {

		fmtStr += dtfGen.TimeZoneSeparator
		fmtStr += dtfGen.TimeZoneElement
	}

	if fmtStr != "" {
		dtf.assignFormatStrToMap(fmtStr)
	}

	// Calculate variation of format string where
	// Time Zone comes before Offset Element

	if dtfGen.TimeZoneElement != "" &&
		dtfGen.TimeElement != "" &&
		fmtStr2 != "" &&
		dtfGen.OffsetElement != "" {

		fmtStr2 += dtfGen.TimeZoneSeparator
		fmtStr2 += dtfGen.TimeZoneElement
		fmtStr2 += dtfGen.OffsetSeparator
		fmtStr2 += dtfGen.OffsetElement
	}

	if fmtStr2 != "" {
		dtf.assignFormatStrToMap(fmtStr2)
	}

	return

}

func (dtf *DateTimeFormatUtility) assemblePreDefinedFormats() {

	preDefFmts := dtf.getPredefinedFormats()

	for _, pdf := range preDefFmts {

		dtf.assignFormatStrToMap(pdf)

	}

}

func (dtf *DateTimeFormatUtility) assembleEdgeCaseFormats() {
	edgeCases := dtf.getEdgeCases()

	for _, ecf := range edgeCases {
		dtf.assignFormatStrToMap(ecf)
	}
}

func (dtf *DateTimeFormatUtility) analyzeDofWeekMMDDTimeOffsetTzYYYY(dtfGen DateTimeFormatGenerator) {

	fmtStr := ""
	fmtStr2 := ""

	if dtfGen.DayOfWeek != "" {
		fmtStr += dtfGen.DayOfWeek
	}

	if dtfGen.MthDay != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.MthDay
		} else {
			fmtStr += dtfGen.DayOfWeekSeparator
			fmtStr += dtfGen.MthDay
		}
	}

	if dtfGen.TimeElement != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.TimeElement
		} else {
			fmtStr += dtfGen.AfterMthDaySeparator
			fmtStr += dtfGen.TimeElement
		}
	}

	fmtStr2 = fmtStr

	if dtfGen.OffsetElement != "" &&
		fmtStr != "" &&
		dtfGen.TimeElement != "" {
		fmtStr += dtfGen.OffsetSeparator
		fmtStr += dtfGen.OffsetElement

	}

	if dtfGen.TimeZoneElement != "" &&
		fmtStr != "" &&
		dtfGen.TimeElement != "" {
		fmtStr += dtfGen.TimeZoneSeparator
		fmtStr += dtfGen.TimeZoneElement
	}

	if fmtStr != "" {
		dtf.assignFormatStrToMap(fmtStr)
	}

	// Calculate variation of format string where
	// Time Zone comes before Offset Element

	if dtfGen.TimeZoneElement != "" &&
		fmtStr2 != "" &&
		dtfGen.TimeElement != "" {
		fmtStr2 += dtfGen.TimeZoneSeparator
		fmtStr2 += dtfGen.TimeZoneElement
	}

	if dtfGen.OffsetElement != "" &&
		fmtStr2 != "" &&
		dtfGen.TimeElement != "" {
		fmtStr2 += dtfGen.OffsetSeparator
		fmtStr2 += dtfGen.OffsetElement

	}

	if fmtStr2 != "" {
		dtf.assignFormatStrToMap(fmtStr)
	}

	return
}

func (dtf *DateTimeFormatUtility) assignFormatStrToMap(fmtStr string) {

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

func (dtf *DateTimeFormatUtility) getDayOfWeekElements() ([]string, error) {
	dayOfWeek := make([]string, 0, 10)

	dayOfWeek = append(dayOfWeek, "")
	dayOfWeek = append(dayOfWeek, "Mon")
	dayOfWeek = append(dayOfWeek, "Monday")

	return dayOfWeek, nil
}

func (dtf *DateTimeFormatUtility) getDayOfWeekSeparator() ([]string, error) {
	dayOfWeekSeparator := make([]string, 0, 1024)

	dayOfWeekSeparator = append(dayOfWeekSeparator, " ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, ", ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, " - ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, "-")

	return dayOfWeekSeparator, nil
}

func (dtf *DateTimeFormatUtility) getMonthDayYearElements() ([]string, error) {
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

func (dtf *DateTimeFormatUtility) getMonthDayElements() ([]string, error) {
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

func (dtf *DateTimeFormatUtility) getYears() ([]string, error) {
	yearElements := make([]string, 0, 10)

	yearElements = append(yearElements, "2006")
	yearElements = append(yearElements, "06")
	yearElements = append(yearElements, "'06")

	return yearElements, nil
}

func (dtf *DateTimeFormatUtility) getAfterMthDaySeparators() ([]string, error) {
	mthDayAfterSeparators := make([]string, 0, 10)

	mthDayAfterSeparators = append(mthDayAfterSeparators, " ")
	mthDayAfterSeparators = append(mthDayAfterSeparators, ", ")
	mthDayAfterSeparators = append(mthDayAfterSeparators, ":")
	mthDayAfterSeparators = append(mthDayAfterSeparators, "T")
	mthDayAfterSeparators = append(mthDayAfterSeparators, "")

	return mthDayAfterSeparators, nil

}

func (dtf *DateTimeFormatUtility) getStandardSeparators() ([]string, error) {
	standardSeparators := make([]string, 0, 10)

	standardSeparators = append(standardSeparators, " ")
	standardSeparators = append(standardSeparators, "")

	return standardSeparators, nil
}

func (dtf *DateTimeFormatUtility) getDateTimeSeparators() ([]string, error) {
	dtTimeSeparators := make([]string, 0, 10)

	dtTimeSeparators = append(dtTimeSeparators, " ")
	dtTimeSeparators = append(dtTimeSeparators, ":")
	dtTimeSeparators = append(dtTimeSeparators, "T")
	dtTimeSeparators = append(dtTimeSeparators, "")

	return dtTimeSeparators, nil
}

func (dtf *DateTimeFormatUtility) getTimeElements() ([]string, error) {
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

func (dtf *DateTimeFormatUtility) getTimeOffsets() ([]string, error) {
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

func (dtf *DateTimeFormatUtility) getTimeOffsetSeparators() ([]string, error) {
	timeOffsetSeparators := make([]string, 0, 20)

	timeOffsetSeparators = append(timeOffsetSeparators, " ")
	timeOffsetSeparators = append(timeOffsetSeparators, "-")
	timeOffsetSeparators = append(timeOffsetSeparators, "")

	return timeOffsetSeparators, nil
}

func (dtf *DateTimeFormatUtility) getTimeZoneElements() ([]string, error) {
	tzElements := make([]string, 0, 20)

	tzElements = append(tzElements, "MST")
	tzElements = append(tzElements, "")

	return tzElements, nil
}

func (dtf *DateTimeFormatUtility) getTimeZoneSeparators() ([]string, error) {
	tzElements := make([]string, 0, 20)

	tzElements = append(tzElements, " ")
	tzElements = append(tzElements, "-")
	tzElements = append(tzElements, "")

	return tzElements, nil
}

func (dtf *DateTimeFormatUtility) getPredefinedFormats() []string {

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

func (dtf *DateTimeFormatUtility) getEdgeCases() []string {
	// FmtDateTimeEverything = "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
	edgeCases := make([]string, 0, 20)

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

func (dtf *DateTimeFormatUtility) getFirstSecondThirdSearchStrs() [][][]string {
	d := make([][][]string, 0)

	d = append(d, [][]string{{"1ST", "1"}})
	d = append(d, [][]string{{"2ND", "2"}})
	d = append(d, [][]string{{"3RD", "3"}})
	d = append(d, [][]string{{"4TH", "4"}})
	d = append(d, [][]string{{"5TH", "5"}})
	d = append(d, [][]string{{"6TH", "6"}})
	d = append(d, [][]string{{"7TH", "7"}})
	d = append(d, [][]string{{"8TH", "8"}})
	d = append(d, [][]string{{"9TH", "9"}})
	d = append(d, [][]string{{"10TH", "10"}})
	d = append(d, [][]string{{"11TH", "11"}})
	d = append(d, [][]string{{"12TH", "12"}})
	d = append(d, [][]string{{"13TH", "13"}})
	d = append(d, [][]string{{"14TH", "14"}})
	d = append(d, [][]string{{"15TH", "15"}})
	d = append(d, [][]string{{"16TH", "16"}})
	d = append(d, [][]string{{"17TH", "17"}})
	d = append(d, [][]string{{"18TH", "18"}})
	d = append(d, [][]string{{"19TH", "9"}})
	d = append(d, [][]string{{"20TH", "20"}})
	d = append(d, [][]string{{"21ST", "21"}})
	d = append(d, [][]string{{"22ND", "22"}})
	d = append(d, [][]string{{"23RD", "23"}})
	d = append(d, [][]string{{"24TH", "24"}})
	d = append(d, [][]string{{"25TH", "25"}})
	d = append(d, [][]string{{"26TH", "26"}})
	d = append(d, [][]string{{"27TH", "27"}})
	d = append(d, [][]string{{"28TH", "28"}})
	d = append(d, [][]string{{"29TH", "29"}})
	d = append(d, [][]string{{"30TH", "30"}})
	d = append(d, [][]string{{"31ST", "31"}})
	d = append(d, [][]string{{"1st", "1"}})
	d = append(d, [][]string{{"2nd", "2"}})
	d = append(d, [][]string{{"3rd", "3"}})
	d = append(d, [][]string{{"4th", "4"}})
	d = append(d, [][]string{{"5th", "5"}})
	d = append(d, [][]string{{"6th", "6"}})
	d = append(d, [][]string{{"7th", "7"}})
	d = append(d, [][]string{{"8th", "8"}})
	d = append(d, [][]string{{"9th", "9"}})
	d = append(d, [][]string{{"10th", "10"}})
	d = append(d, [][]string{{"11th", "11"}})
	d = append(d, [][]string{{"12th", "12"}})
	d = append(d, [][]string{{"13th", "13"}})
	d = append(d, [][]string{{"14th", "14"}})
	d = append(d, [][]string{{"15th", "15"}})
	d = append(d, [][]string{{"16th", "16"}})
	d = append(d, [][]string{{"17th", "17"}})
	d = append(d, [][]string{{"18th", "18"}})
	d = append(d, [][]string{{"19th", "9"}})
	d = append(d, [][]string{{"20th", "20"}})
	d = append(d, [][]string{{"21st", "21"}})
	d = append(d, [][]string{{"22nd", "22"}})
	d = append(d, [][]string{{"23rd", "23"}})
	d = append(d, [][]string{{"24th", "24"}})
	d = append(d, [][]string{{"25th", "25"}})
	d = append(d, [][]string{{"26th", "26"}})
	d = append(d, [][]string{{"27th", "27"}})
	d = append(d, [][]string{{"28th", "28"}})
	d = append(d, [][]string{{"29th", "29"}})
	d = append(d, [][]string{{"30th", "30"}})
	d = append(d, [][]string{{"31st", "31"}})

	return d
}


func (dtf *DateTimeFormatUtility) getPreTrimSearchStrings() [][][]string {
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

func (dtf *DateTimeFormatUtility) getTimeFmtRegEx() [][][]string {
	d := make([][][]string, 0)
	d = append(d, [][]string{{"\\d\\d:\\d\\d:\\d\\d", "%02d:%02d:%02d"}}) // 2:2:2
	d = append(d, [][]string{{"\\d\\d:\\d:\\d\\d", "%02d:%02d:%02d"}})    // 2:1:2
	d = append(d, [][]string{{"\\d\\d:\\d\\d:\\d", "%02d:%02d:%02d"}})    // 2:2:1
	d = append(d, [][]string{{"\\d\\d:\\d:\\d", "%02d:%02d:%02d"}})       // 2:1:1
	d = append(d, [][]string{{"\\d:\\d\\d:\\d\\d", "%02d:%02d:%02d"}})    // 1:2:2
	d = append(d, [][]string{{"\\d:\\d:\\d\\d", "%02d:%02d:%02d"}})       // 1:1:2
	d = append(d, [][]string{{"\\d:\\d\\d:\\d", "%02d:%02d:%02d"}})    		// 1:2:1
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

func (dtf *DateTimeFormatUtility) reformatSingleTimeString(targetStr string, regExes [][][]string) string {

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

func (dtf *DateTimeFormatUtility) replaceMultipleStrSequence(targetStr string, replaceMap [][][]string) string {

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

func (dtf *DateTimeFormatUtility) replaceAMPM(targetStr string) string {
	d := make([][][]string, 0)

	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)a[.]*\\s{0,4}(?i)m[.]*", " am"}})
	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)p[.]*\\s{0,4}(?i)m[.]*", " pm"}})

	lD := len(d)

	for i:=0; i < lD; i++ {
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

		foundEx := string(bTargetStr[loc[0]+1:loc[1]])

		return strings.Replace(targetStr,foundEx,d[i][0][1],1)

	}

	return targetStr
}

func (dtf *DateTimeFormatUtility) replaceDateTimeSequence(targetStr string, replaceMap [][][]string) string {

	max := len(replaceMap)

	for i := 0; i < max; i++ {
		if strings.Contains(targetStr, replaceMap[i][0][0]) {
			return strings.Replace(targetStr, replaceMap[i][0][0], replaceMap[i][0][1], 1)
		}

	}

	return targetStr
}
