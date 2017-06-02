package common

import (
	"errors"
	"fmt"
	"os"
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
	AMPMSearchStrs             map[string]string
	FirstSecondThirdSearchStrs map[string]string
}

type DateTimeFormatUtility struct {
	OriginalDateTimeStringIn  string
	FormattedDateTimeStringIn string
	FormatMap                 map[int]map[string]int
	SelectedMapIdx            int
	SelectedFormat            string
	SelectedFormatSource      string
	DictSearches              map[int]int
	TotalNoOfDictSearches     int
	DateTimeOut               time.Time
	NumOfFormatsGenerated     int
	FormatSearchReplaceStrs   SearchStrings
}

// CreateAllFormatsInMemory - Generates 11-million permutations of
// Date Time Formats and stores them in memory as a series of
// maps of maps using the field: DateTimeFormatUtility.FormatMap
func (dtf *DateTimeFormatUtility) CreateAllFormatsInMemory() (err error) {

	dtf.FormatMap = make(map[int]map[string]int)
	dtf.NumOfFormatsGenerated = 0
	dtf.assemblePreDefinedFormats()
	dtf.assembleMthDayYearFmts()
	dtf.assembleEdgeCaseFormats()
	dtf.FormatSearchReplaceStrs.AMPMSearchStrs = dtf.getAMPMSearchStrs()
	dtf.FormatSearchReplaceStrs.FirstSecondThirdSearchStrs = dtf.getFirstSecondThirdSearchStrs()

	return
}

// LoadAllFormatsFromFileIntoMemory - Loads all date time formats from a specified
// text file into memory. The formats are stored in DateTimeFormatUtility.FormatMap
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
// text format.
func (dtf *DateTimeFormatUtility) WriteAllFormatsInMemoryToFile(pathFileName string) (DateTimeWriteFormatsToFileDto, error) {

	fwDto := DateTimeWriteFormatsToFileDto{}

	fwDto.FileWriteStartTime = time.Now()
	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		return fwDto, errors.New("WriteAllFormatsInMemoryToFile() Error - There are NO Formats in Memory -  FormatMap length == 0")
	}

	outF, err := os.Create(pathFileName)

	if err != nil {
		return fwDto, fmt.Errorf("WriteAllFormatsInMemoryToFile() Error - Failed create output file %v. Error: %v", pathFileName, err.Error())
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
				return fwDto, fmt.Errorf("WriteAllFormatsInMemoryToFile() Error writing Format data to output file %v. Error: %v", pathFileName, err.Error())
			}
		}
	}

	outF.Sync()

	du := DurationUtility{}

	fwDto.FileWriteEndTime = time.Now()

	etFileWrite, _ := du.GetElapsedTime(fwDto.FileWriteStartTime, fwDto.FileWriteEndTime)

	fwDto.OutputPathFileName = pathFileName
	fwDto.ElapsedTimeForFileWriteOps = etFileWrite.DurationStr

	return fwDto, nil
}

func (dtf *DateTimeFormatUtility) extractFmtRecordFromBuffer(
	inBuff []byte, outRecordBuff []byte, idx int, lastidx int) (nextIdx int, isPartialRec bool) {

	for i := idx; i <= lastidx; i++ {

		if inBuff[i] == '\n' {
			nextIdx = i + 1
			isPartialRec = false
			return
		}

		outRecordBuff = append(outRecordBuff, inBuff[i])

	}

	nextIdx = -1
	isPartialRec = true
	return
}

func (dtf *DateTimeFormatUtility) Empty() {
	dtf.OriginalDateTimeStringIn = ""
	dtf.FormattedDateTimeStringIn = ""
	dtf.DateTimeOut = time.Time{}
	dtf.SelectedFormat = ""
	dtf.SelectedFormatSource = ""
	dtf.SelectedMapIdx = -1
	dtf.DictSearches = make(map[int]int)
	dtf.TotalNoOfDictSearches = 0

}

func (dtf *DateTimeFormatUtility) ParseDateTimeString(timeStr string, probableFormat string) (time.Time, error) {

	if timeStr == "" {
		return time.Time{}, errors.New("Empty Time String!")
	}

	dtf.Empty()

	ftimeStr, err := StringUtility{}.TrimEndMultiple(timeStr, ' ')

	ftimeStr = dtf.replaceDateTimeSequence(ftimeStr, dtf.FormatSearchReplaceStrs.AMPMSearchStrs)
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
			dtf.OriginalDateTimeStringIn = timeStr
			dtf.FormattedDateTimeStringIn = ftimeStr
			dtf.TotalNoOfDictSearches = 1
			dtf.DictSearches[0] = 1

			return t, nil
		}

	}

	if len(dtf.FormatMap) == 0 {
		return time.Time{}, errors.New("Format Map is EMPTY! Load Formats into DateTimeFormatUtility.FormatMap first!")
	}

	lenStr := len(ftimeStr)

	lenTests := [7]int{lenStr - 1, lenStr, lenStr - 2, lenStr + 1, lenStr + 2, lenStr - 3, lenStr + 3}

	dtf.TotalNoOfDictSearches = 0

	for _, lTest := range lenTests {

		result, err := dtf.parseFormatMap(ftimeStr, lTest)

		dtf.DictSearches[lTest] = result.TotalNoOfDictSearches
		dtf.TotalNoOfDictSearches += result.TotalNoOfDictSearches

		if err == nil {
			dtf.SelectedFormat = result.SelectedFormat
			dtf.SelectedFormatSource = result.SelectedFormatSource
			dtf.SelectedMapIdx = result.SelectedMapIdx
			dtf.DateTimeOut = result.DateTimeOut
			dtf.OriginalDateTimeStringIn = timeStr
			dtf.FormattedDateTimeStringIn = result.FormattedDateTimeStringIn
			return dtf.DateTimeOut, nil
		}

	}

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

func (dtf DateTimeFormatUtility) getDayOfWeekElements() ([]string, error) {
	dayOfWeek := make([]string, 0, 10)

	dayOfWeek = append(dayOfWeek, "")
	dayOfWeek = append(dayOfWeek, "Mon")
	dayOfWeek = append(dayOfWeek, "Monday")

	return dayOfWeek, nil
}

func (dtf DateTimeFormatUtility) getDayOfWeekSeparator() ([]string, error) {
	dayOfWeekSeparator := make([]string, 0, 1024)

	dayOfWeekSeparator = append(dayOfWeekSeparator, " ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, ", ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, " - ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, "-")

	return dayOfWeekSeparator, nil
}

func (dtf DateTimeFormatUtility) getMonthDayYearElements() ([]string, error) {
	mthDayYr := make([]string, 0, 1024)

	mthDayYr = append(mthDayYr, "2006-1-2")
	mthDayYr = append(mthDayYr, "2006/1/2")
	mthDayYr = append(mthDayYr, "1-2-06")
	mthDayYr = append(mthDayYr, "1-2-2006")
	mthDayYr = append(mthDayYr, "1/2/06")
	mthDayYr = append(mthDayYr, "1/2/2006")

	mthDayYr = append(mthDayYr, "Jan-2-06")
	mthDayYr = append(mthDayYr, "Jan 2 06")
	mthDayYr = append(mthDayYr, "Jan 2, 06")
	mthDayYr = append(mthDayYr, "Jan/2/06")
	mthDayYr = append(mthDayYr, "Jan _2 06")
	mthDayYr = append(mthDayYr, "Jan _2, 06")
	mthDayYr = append(mthDayYr, "Jan-2-2006")
	mthDayYr = append(mthDayYr, "Jan 2 2006")
	mthDayYr = append(mthDayYr, "Jan 2, 2006")
	mthDayYr = append(mthDayYr, "Jan/2/2006")
	mthDayYr = append(mthDayYr, "Jan _2 2006")
	mthDayYr = append(mthDayYr, "Jan _2, 2006")

	mthDayYr = append(mthDayYr, "January-2-06")
	mthDayYr = append(mthDayYr, "January 2 06")
	mthDayYr = append(mthDayYr, "January 2, 06")
	mthDayYr = append(mthDayYr, "January/2/06")
	mthDayYr = append(mthDayYr, "January _2 06")
	mthDayYr = append(mthDayYr, "January _2, 06")
	mthDayYr = append(mthDayYr, "January-2-2006")
	mthDayYr = append(mthDayYr, "January 2 2006")
	mthDayYr = append(mthDayYr, "January 2, 2006")
	mthDayYr = append(mthDayYr, "January/2/2006")
	mthDayYr = append(mthDayYr, "January _2 2006")
	mthDayYr = append(mthDayYr, "January _2, 2006")

	// European Date Formats
	mthDayYr = append(mthDayYr, "2.1.06")
	mthDayYr = append(mthDayYr, "2.1.2006")
	mthDayYr = append(mthDayYr, "2.1.'06")

	mthDayYr = append(mthDayYr, "2/January/2006")
	mthDayYr = append(mthDayYr, "2/January/06")
	mthDayYr = append(mthDayYr, "2-January-2006")
	mthDayYr = append(mthDayYr, "2-January-06")
	mthDayYr = append(mthDayYr, "2 January, 06")
	mthDayYr = append(mthDayYr, "2 January, 2006")
	mthDayYr = append(mthDayYr, "2 January 06")
	mthDayYr = append(mthDayYr, "2 January 2006")

	mthDayYr = append(mthDayYr, "2/Jan/2006")
	mthDayYr = append(mthDayYr, "2/Jan/06")
	mthDayYr = append(mthDayYr, "2-Jan-2006")
	mthDayYr = append(mthDayYr, "2-Jan-06")
	mthDayYr = append(mthDayYr, "2 Jan 06")
	mthDayYr = append(mthDayYr, "2 Jan 2006")
	mthDayYr = append(mthDayYr, "2 Jan, 06")
	mthDayYr = append(mthDayYr, "2 Jan, 2006")

	mthDayYr = append(mthDayYr, "20060102")
	mthDayYr = append(mthDayYr, "01022006")
	mthDayYr = append(mthDayYr, "")

	return mthDayYr, nil
}

func (dtf DateTimeFormatUtility) getMonthDayElements() ([]string, error) {
	mthDayElements := make([]string, 0, 124)

	mthDayElements = append(mthDayElements, "Jan 2")
	mthDayElements = append(mthDayElements, "January 2")
	mthDayElements = append(mthDayElements, "Jan _2")
	mthDayElements = append(mthDayElements, "January _2")
	mthDayElements = append(mthDayElements, "1-2")
	mthDayElements = append(mthDayElements, "1-_2")
	mthDayElements = append(mthDayElements, "1/2")
	mthDayElements = append(mthDayElements, "1/_2")
	mthDayElements = append(mthDayElements, "1/2")
	mthDayElements = append(mthDayElements, "1-2")
	mthDayElements = append(mthDayElements, "0102")
	// European Format Day Month
	mthDayElements = append(mthDayElements, "02.01")
	mthDayElements = append(mthDayElements, "2.1")
	mthDayElements = append(mthDayElements, "02.1")
	mthDayElements = append(mthDayElements, "2.01")

	return mthDayElements, nil
}

func (dtf DateTimeFormatUtility) getYears() ([]string, error) {
	yearElements := make([]string, 0, 10)

	yearElements = append(yearElements, "2006")
	yearElements = append(yearElements, "06")
	yearElements = append(yearElements, "'06")

	return yearElements, nil
}

func (dtf DateTimeFormatUtility) getAfterMthDaySeparators() ([]string, error) {
	mthDayAfterSeparators := make([]string, 0, 10)

	mthDayAfterSeparators = append(mthDayAfterSeparators, " ")
	mthDayAfterSeparators = append(mthDayAfterSeparators, ", ")
	mthDayAfterSeparators = append(mthDayAfterSeparators, ":")
	mthDayAfterSeparators = append(mthDayAfterSeparators, "T")
	mthDayAfterSeparators = append(mthDayAfterSeparators, "")

	return mthDayAfterSeparators, nil

}

func (dtf DateTimeFormatUtility) getStandardSeparators() ([]string, error) {
	standardSeparators := make([]string, 0, 10)

	standardSeparators = append(standardSeparators, " ")
	standardSeparators = append(standardSeparators, "")

	return standardSeparators, nil
}

func (dtf DateTimeFormatUtility) getDateTimeSeparators() ([]string, error) {
	dtTimeSeparators := make([]string, 0, 10)

	dtTimeSeparators = append(dtTimeSeparators, " ")
	dtTimeSeparators = append(dtTimeSeparators, ":")
	dtTimeSeparators = append(dtTimeSeparators, "T")
	dtTimeSeparators = append(dtTimeSeparators, "")

	return dtTimeSeparators, nil
}

func (dtf DateTimeFormatUtility) getTimeElements() ([]string, error) {
	timeElements := make([]string, 0, 512)

	timeElements = append(timeElements, "15:4:5")
	timeElements = append(timeElements, "15:4")
	timeElements = append(timeElements, "15:4:5.000")
	timeElements = append(timeElements, "15:4:5.000000")
	timeElements = append(timeElements, "15:4:5.000000000")

	timeElements = append(timeElements, "15:04:05")
	timeElements = append(timeElements, "15:04")
	timeElements = append(timeElements, "15:04:05.000")
	timeElements = append(timeElements, "15:04:05.000000")
	timeElements = append(timeElements, "15:04:05.000000000")

	timeElements = append(timeElements, "03:04:05 pm")
	timeElements = append(timeElements, "3:4:5 pm")
	timeElements = append(timeElements, "03:04 pm")
	timeElements = append(timeElements, "3:04 pm")
	timeElements = append(timeElements, "3:4 pm")

	timeElements = append(timeElements, "")

	return timeElements, nil
}

func (dtf DateTimeFormatUtility) getTimeOffsets() ([]string, error) {
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

func (dtf DateTimeFormatUtility) getTimeOffsetSeparators() ([]string, error) {
	timeOffsetSeparators := make([]string, 0, 20)

	timeOffsetSeparators = append(timeOffsetSeparators, " ")
	timeOffsetSeparators = append(timeOffsetSeparators, "-")
	timeOffsetSeparators = append(timeOffsetSeparators, "")

	return timeOffsetSeparators, nil
}

func (dtf DateTimeFormatUtility) getTimeZoneElements() ([]string, error) {
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

	edgeCases = append(edgeCases, "Monday January 2 15:4:5 -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2, 15:4:5 -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2 15:4:5 -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2, 15:4:5 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2 15:4:5 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2, 15:4:5 -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2 15:4:5 -0700 MST, 2006")
	edgeCases = append(edgeCases, "Mon January 2 15:4:5 -0700 MST, 2006")
	edgeCases = append(edgeCases, "Jan 2 15:4:5 -0700 MST, 2006")

	edgeCases = append(edgeCases, "Monday January 2 15:4 -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2, 15:4 -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2 15:4 -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2, 15:4 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2 15:4 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2, 15:4 -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2 15:4 -0700 MST, 2006")
	edgeCases = append(edgeCases, "Mon January 2 15:4 -0700 MST, 2006")
	edgeCases = append(edgeCases, "Jan 2 15:4 -0700 MST, 2006")

	edgeCases = append(edgeCases, "Monday January 2 15:04:05 -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2, 15:04:05 -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2 15:04:05 -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2, 15:04:05 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2 15:04:05 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2, 15:04:05 -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2 15:04:05 -0700 MST, 2006")
	edgeCases = append(edgeCases, "Mon January 2 15:04:05 -0700 MST, 2006")
	edgeCases = append(edgeCases, "Jan 2 15:04:05 -0700 MST, 2006")

	edgeCases = append(edgeCases, "Monday January 2 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2, 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2, 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2, 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2 15:04 -0700 MST, 2006")
	edgeCases = append(edgeCases, "Mon January 2 15:04 -0700 MST, 2006")
	edgeCases = append(edgeCases, "Jan 2 15:04 -0700 MST, 2006")

	edgeCases = append(edgeCases, "Monday January 2 3:4:5pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2, 3:4:5pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2 3:4:5pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2, 3:4:5pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2 3:4:5pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2, 3:4:5pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2 3:4:5pm -0700 MST, 2006")
	edgeCases = append(edgeCases, "Mon January 2 3:4:5pm -0700 MST, 2006")
	edgeCases = append(edgeCases, "Jan 2 3:4:5pm -0700 MST, 2006")
	edgeCases = append(edgeCases, "Monday January 2 3:4pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2, 3:4pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2 3:4pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2, 3:4pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2 3:4pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2, 3:4pm -0700 MST 2006")
	edgeCases = append(edgeCases, "Monday January 2 15:4 -0700 MST, 2006")
	edgeCases = append(edgeCases, "Mon January 2 3:4pm -0700 MST, 2006")
	edgeCases = append(edgeCases, "Jan 2 3:4pm -0700 MST, 2006")

	return edgeCases
}

func (dtu *DateTimeFormatUtility) getFirstSecondThirdSearchStrs() map[string]string {
	searchStrs := make(map[string]string)

	searchStrs["1ST"] = "1"
	searchStrs["2ND"] = "2"
	searchStrs["3RD"] = "3"
	searchStrs["4TH"] = "4"
	searchStrs["5TH"] = "5"
	searchStrs["6TH"] = "6"
	searchStrs["7TH"] = "7"
	searchStrs["8TH"] = "8"
	searchStrs["9TH"] = "9"
	searchStrs["10TH"] = "10"
	searchStrs["11TH"] = "11"
	searchStrs["12TH"] = "12"
	searchStrs["13TH"] = "13"
	searchStrs["14TH"] = "14"
	searchStrs["15TH"] = "15"
	searchStrs["16TH"] = "16"
	searchStrs["17TH"] = "17"
	searchStrs["18TH"] = "18"
	searchStrs["19TH"] = "9"
	searchStrs["20TH"] = "20"
	searchStrs["21ST"] = "21"
	searchStrs["22ND"] = "22"
	searchStrs["23RD"] = "23"
	searchStrs["24TH"] = "24"
	searchStrs["25TH"] = "25"
	searchStrs["26TH"] = "26"
	searchStrs["27TH"] = "27"
	searchStrs["28TH"] = "28"
	searchStrs["29TH"] = "29"
	searchStrs["30TH"] = "30"
	searchStrs["31ST"] = "31"
	searchStrs["1st"] = "1"
	searchStrs["2nd"] = "2"
	searchStrs["3rd"] = "3"
	searchStrs["4th"] = "4"
	searchStrs["5th"] = "5"
	searchStrs["6th"] = "6"
	searchStrs["7th"] = "7"
	searchStrs["8th"] = "8"
	searchStrs["9th"] = "9"
	searchStrs["10th"] = "10"
	searchStrs["11th"] = "11"
	searchStrs["12th"] = "12"
	searchStrs["13th"] = "13"
	searchStrs["14th"] = "14"
	searchStrs["15th"] = "15"
	searchStrs["16th"] = "16"
	searchStrs["17th"] = "17"
	searchStrs["18th"] = "18"
	searchStrs["19th"] = "9"
	searchStrs["20th"] = "20"
	searchStrs["21st"] = "21"
	searchStrs["22nd"] = "22"
	searchStrs["23rd"] = "23"
	searchStrs["24th"] = "24"
	searchStrs["25th"] = "25"
	searchStrs["26th"] = "26"
	searchStrs["27th"] = "27"
	searchStrs["28th"] = "28"
	searchStrs["29th"] = "29"
	searchStrs["30th"] = "30"
	searchStrs["31st"] = "31"

	return searchStrs
}

func (dtu *DateTimeFormatUtility) getAMPMSearchStrs() map[string]string {
	searchStrs := make(map[string]string)

	searchStrs["0 PM"] = "0 pm"
	searchStrs["1 PM"] = "1 pm"
	searchStrs["2 PM"] = "2 pm"
	searchStrs["3 PM"] = "3 pm"
	searchStrs["4 PM"] = "4 pm"
	searchStrs["5 PM"] = "5 pm"
	searchStrs["6 PM"] = "6 pm"
	searchStrs["7 PM"] = "7 pm"
	searchStrs["8 PM"] = "8 pm"
	searchStrs["9 PM"] = "9 pm"
	searchStrs["0 P.M."] = "0 pm"
	searchStrs["1 P.M."] = "1 pm"
	searchStrs["2 P.M."] = "2 pm"
	searchStrs["3 P.M."] = "3 pm"
	searchStrs["4 P.M."] = "4 pm"
	searchStrs["5 P.M."] = "5 pm"
	searchStrs["6 P.M."] = "6 pm"
	searchStrs["7 P.M."] = "7 pm"
	searchStrs["8 P.M."] = "8 pm"
	searchStrs["9 P.M."] = "9 pm"
	searchStrs["0PM"] = "0 pm"
	searchStrs["1PM"] = "1 pm"
	searchStrs["2PM"] = "2 pm"
	searchStrs["3PM"] = "3 pm"
	searchStrs["4PM"] = "4 pm"
	searchStrs["5PM"] = "5 pm"
	searchStrs["6PM"] = "6 pm"
	searchStrs["7PM"] = "7 pm"
	searchStrs["8PM"] = "8 pm"
	searchStrs["9PM"] = "9 pm"
	searchStrs["0P.M."] = "0 pm"
	searchStrs["1P.M."] = "1 pm"
	searchStrs["2P.M."] = "2 pm"
	searchStrs["3P.M."] = "3 pm"
	searchStrs["4P.M."] = "4 pm"
	searchStrs["5P.M."] = "5 pm"
	searchStrs["6P.M."] = "6 pm"
	searchStrs["7P.M."] = "7 pm"
	searchStrs["8P.M."] = "8 pm"
	searchStrs["9P.M."] = "9 pm"
	searchStrs["0 AM"] = "0 am"
	searchStrs["1 AM"] = "1 am"
	searchStrs["2 AM"] = "2 am"
	searchStrs["3 AM"] = "3 am"
	searchStrs["4 AM"] = "4 am"
	searchStrs["5 AM"] = "5 am"
	searchStrs["6 AM"] = "6 am"
	searchStrs["7 AM"] = "7 am"
	searchStrs["8 AM"] = "8 am"
	searchStrs["9 AM"] = "9 am"
	searchStrs["0 A.M."] = "0 am"
	searchStrs["1 A.M."] = "1 am"
	searchStrs["2 A.M."] = "2 am"
	searchStrs["3 A.M."] = "3 am"
	searchStrs["4 A.M."] = "4 am"
	searchStrs["5 A.M."] = "5 am"
	searchStrs["6 A.M."] = "6 am"
	searchStrs["7 A.M."] = "7 am"
	searchStrs["8 A.M."] = "8 am"
	searchStrs["9 A.M."] = "9 am"
	searchStrs["0AM"] = "0 am"
	searchStrs["1AM"] = "1 am"
	searchStrs["2AM"] = "2 am"
	searchStrs["3AM"] = "3 am"
	searchStrs["4AM"] = "4 am"
	searchStrs["5AM"] = "5 am"
	searchStrs["6AM"] = "6 am"
	searchStrs["7AM"] = "7 am"
	searchStrs["8AM"] = "8 am"
	searchStrs["9AM"] = "9 am"
	searchStrs["0A.M."] = "0 am"
	searchStrs["1A.M."] = "1 am"
	searchStrs["2A.M."] = "2 am"
	searchStrs["3A.M."] = "3 am"
	searchStrs["4A.M."] = "4 am"
	searchStrs["5A.M."] = "5 am"
	searchStrs["6A.M."] = "6 am"
	searchStrs["7A.M."] = "7 am"
	searchStrs["8A.M."] = "8 am"
	searchStrs["9A.M."] = "9 am"

	return searchStrs
}

func (dtu *DateTimeFormatUtility) replaceDateTimeSequence(targetStr string, replaceMap map[string]string) string {

	for key, replaceStr := range replaceMap {
		if strings.Contains(targetStr, key) {
			return strings.Replace(targetStr, key, replaceStr, 1)
		}
	}

	return targetStr
}
