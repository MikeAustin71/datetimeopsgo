package common

import (
	"errors"
	"time"
)

type DateTimeFormatGenerator struct {
	DayOfWeek          string
	DayOfWeekSeparator string
	MthDayYear         string
	DateTimeSeparator  string
	TimeElement        string
	OffsetSeparator    string
	OffsetElement      string
	TimeZoneSeparator  string
	TimeZoneElement    string
}

type DateTimeFormatUtility struct {
	DateTimeStringIn      string
	FormatMap             map[int][]string
	SelectedFormat        string
	DateTimeOut           time.Time
	NumOfFormatsGenerated int
}

func (dtf *DateTimeFormatUtility) GetAllDateTimeFormats() (err error) {

	dtf.FormatMap = make(map[int][]string)
	dtf.NumOfFormatsGenerated = 0
	dtf.assemblePreDefinedFormats()
	dtf.assembleDayMthYearFmts()

	return
}

func (dtf *DateTimeFormatUtility) ParseDateTimeString(timeStr string, probableFormat string) (t time.Time,
	idx int, lenStr int, err error) {

	if probableFormat != "" {
		t, err = time.Parse(probableFormat, timeStr)

		if err == nil {
			return
		}

	}

	lenStr = len(timeStr) - 1

	t, idx, err = dtf.parseFormatMap(timeStr, lenStr)

	if err == nil {
		return
	}

	lenStr++
	t, idx, err = dtf.parseFormatMap(timeStr, lenStr)

	if err == nil {
		return
	}

	return t, idx, lenStr, errors.New("Falied to locate correct time format!")
}

func (dtf *DateTimeFormatUtility) assembleDayMthYears() error {

	dtf.FormatMap = make(map[int][]string)

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

func (dtf *DateTimeFormatUtility) parseFormatMap(timeStr string, idx int) (t time.Time, ix int, err error) {

	if dtf.FormatMap[idx] == nil {
		err = errors.New("Time String Length not found in Format Map!")
		return
	}

	for i, f := range dtf.FormatMap[idx] {

		t, err = time.Parse(f, timeStr)

		if err == nil {
			return t, i, err
		}

	}

	err = errors.New("Found Format Map formats, but failed to parse time string")

	return
}

func (dtf *DateTimeFormatUtility) assembleDayMthYearFmts() error {

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

func (dtf *DateTimeFormatUtility) assemblePreDefinedFormats() {

	preDefFmts := dtf.getPredefinedFormats()

	for _, pdf := range preDefFmts {

		dtf.assignFormatStrToMap(pdf)

	}

}

func (dtf *DateTimeFormatUtility) analyzeDofWeekMMDDYYYYTimeOffsetTz(dtfGen DateTimeFormatGenerator) {

	fmtStr := ""

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

	if dtfGen.OffsetElement != "" {
		if fmtStr == "" || dtfGen.TimeElement == "" {
			return
		} else {
			fmtStr += dtfGen.OffsetSeparator
			fmtStr += dtfGen.OffsetElement
		}
	}

	if dtfGen.TimeZoneElement != "" {
		if fmtStr == "" || dtfGen.TimeElement == "" {
			return
		} else {
			fmtStr += dtfGen.TimeZoneSeparator
			fmtStr += dtfGen.TimeZoneElement
		}

	}

	dtf.assignFormatStrToMap(fmtStr)

	return
}

func (dtf *DateTimeFormatUtility) assignFormatStrToMap(fmtStr string) {

	l := len(fmtStr)

	if l == 0 {
		return
	}

	sliceCap := 25000

	if l > 19 && l < 48 {
		sliceCap = 100000
	}

	if dtf.FormatMap[l] == nil {
		dtf.FormatMap[l] = make([]string, 0, sliceCap)
	}

	dtf.FormatMap[l] = append(dtf.FormatMap[l], fmtStr)
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

	mthDayYr = append(mthDayYr, "2006-01-02")
	mthDayYr = append(mthDayYr, "2006/01/02")
	mthDayYr = append(mthDayYr, "2006.01.02")
	mthDayYr = append(mthDayYr, "2006-1-2")
	mthDayYr = append(mthDayYr, "2006/1/2")
	mthDayYr = append(mthDayYr, "2006.1.2")

	mthDayYr = append(mthDayYr, "02.01.06")
	mthDayYr = append(mthDayYr, "02.01.2006")
	mthDayYr = append(mthDayYr, "02.01.'06")

	mthDayYr = append(mthDayYr, "01-02-06")
	mthDayYr = append(mthDayYr, "01-02-2006")
	mthDayYr = append(mthDayYr, "1-2-06")
	mthDayYr = append(mthDayYr, "1-2-2006")
	mthDayYr = append(mthDayYr, "1-02-06")
	mthDayYr = append(mthDayYr, "1-02-2006")
	mthDayYr = append(mthDayYr, "01/02/06")
	mthDayYr = append(mthDayYr, "01/02/2006")
	mthDayYr = append(mthDayYr, "1/2/06")
	mthDayYr = append(mthDayYr, "1/02/06")
	mthDayYr = append(mthDayYr, "1/2/2006")
	mthDayYr = append(mthDayYr, "1/02/2006")

	mthDayYr = append(mthDayYr, "Jan-02-06")
	mthDayYr = append(mthDayYr, "Jan 02 06")
	mthDayYr = append(mthDayYr, "Jan 02, 06")
	mthDayYr = append(mthDayYr, "Jan/02/06")
	mthDayYr = append(mthDayYr, "Jan-2-06")
	mthDayYr = append(mthDayYr, "Jan 2 06")
	mthDayYr = append(mthDayYr, "Jan 2, 06")
	mthDayYr = append(mthDayYr, "Jan/2/06")
	mthDayYr = append(mthDayYr, "Jan 02 2006")
	mthDayYr = append(mthDayYr, "Jan 2 2006")
	mthDayYr = append(mthDayYr, "Jan-2-2006")
	mthDayYr = append(mthDayYr, "Jan 02, 2006")
	mthDayYr = append(mthDayYr, "Jan 2, 2006")

	mthDayYr = append(mthDayYr, "Jan _2, 2006")
	mthDayYr = append(mthDayYr, "Jan _2 2006")
	mthDayYr = append(mthDayYr, "Jan _2 06")
	mthDayYr = append(mthDayYr, "Jan _2, 06")

	mthDayYr = append(mthDayYr, "January 02, 06")
	mthDayYr = append(mthDayYr, "January 02, 2006")
	mthDayYr = append(mthDayYr, "January 02 06")
	mthDayYr = append(mthDayYr, "January 02 2006")
	mthDayYr = append(mthDayYr, "January-02-2006")
	mthDayYr = append(mthDayYr, "January-02-06")
	mthDayYr = append(mthDayYr, "January 2, 06")
	mthDayYr = append(mthDayYr, "January 2, 2006")

	mthDayYr = append(mthDayYr, "2 January, 06")
	mthDayYr = append(mthDayYr, "02 January, 06")
	mthDayYr = append(mthDayYr, "2 January, 2006")
	mthDayYr = append(mthDayYr, "02 January, 2006")
	mthDayYr = append(mthDayYr, "2 January 06")
	mthDayYr = append(mthDayYr, "02 January 06")
	mthDayYr = append(mthDayYr, "2 January 2006")
	mthDayYr = append(mthDayYr, "02 January 2006")
	mthDayYr = append(mthDayYr, "02/Jan/2006")
	mthDayYr = append(mthDayYr, "2/Jan/2006")
	mthDayYr = append(mthDayYr, "02/Jan/06")
	mthDayYr = append(mthDayYr, "2/Jan/06")
	mthDayYr = append(mthDayYr, "02 Jan 06")
	mthDayYr = append(mthDayYr, "02 Jan 2006")
	mthDayYr = append(mthDayYr, "2 Jan 06")
	mthDayYr = append(mthDayYr, "2 Jan 2006")

	mthDayYr = append(mthDayYr, "02 Jan, 06")
	mthDayYr = append(mthDayYr, "02 Jan, 2006")
	mthDayYr = append(mthDayYr, "2 Jan, 06")
	mthDayYr = append(mthDayYr, "2 Jan, 2006")

	// ? May cause problems
	mthDayYr = append(mthDayYr, "06-01-02")
	mthDayYr = append(mthDayYr, "06/01/02")
	mthDayYr = append(mthDayYr, "060102")
	mthDayYr = append(mthDayYr, "06-1-2")
	mthDayYr = append(mthDayYr, "06/1/2")

	mthDayYr = append(mthDayYr, "20060102")
	mthDayYr = append(mthDayYr, "01022006")
	mthDayYr = append(mthDayYr, "")

	return mthDayYr, nil
}

func (dtf DateTimeFormatUtility) getMonthDayElements() ([]string, error) {
	mthDayElements := make([]string, 0, 124)

	mthDayElements = append(mthDayElements, "Jan 2")
	mthDayElements = append(mthDayElements, "January 2")
	mthDayElements = append(mthDayElements, "01-02")
	mthDayElements = append(mthDayElements, "01/02")
	mthDayElements = append(mthDayElements, "1-2")
	mthDayElements = append(mthDayElements, "1/2")
	mthDayElements = append(mthDayElements, "0102")

	return mthDayElements, nil
}

func (dtf DateTimeFormatUtility) getYears() ([]string, error) {
	yearElements := make([]string, 0, 10)

	yearElements = append(yearElements, "2006")
	yearElements = append(yearElements, "06")
	yearElements = append(yearElements, "'06")

	return yearElements, nil
}

func (dtf DateTimeFormatUtility) getStandardSeparators() ([]string, error) {
	standardSeparators := make([]string, 0, 10)

	standardSeparators = append(standardSeparators, " ")

	return standardSeparators, nil
}

func (dtf DateTimeFormatUtility) getDateTimeSeparators() ([]string, error) {
	dtTimeSeparators := make([]string, 0, 10)

	dtTimeSeparators = append(dtTimeSeparators, " ")
	dtTimeSeparators = append(dtTimeSeparators, ":")
	dtTimeSeparators = append(dtTimeSeparators, "T")

	return dtTimeSeparators, nil
}

func (dtf DateTimeFormatUtility) getTimeElements() ([]string, error) {
	timeElements := make([]string, 0, 512)

	timeElements = append(timeElements, "15:04:05")
	timeElements = append(timeElements, "15:04")
	timeElements = append(timeElements, "15:04:05.000")
	timeElements = append(timeElements, "15:04:05.000000")
	timeElements = append(timeElements, "15:04:05.000000000")

	timeElements = append(timeElements, "15:4:5")
	timeElements = append(timeElements, "15:4")
	timeElements = append(timeElements, "15:4:5.000")
	timeElements = append(timeElements, "15:4:5.000000")
	timeElements = append(timeElements, "15:4:5.000000000")

	timeElements = append(timeElements, "03:04:05pm")
	timeElements = append(timeElements, "03:04:05p.m.")
	timeElements = append(timeElements, "03:04:05 pm")
	timeElements = append(timeElements, "03:04:05 p.m.")

	timeElements = append(timeElements, "03:04:05PM")
	timeElements = append(timeElements, "03:04:05P.M.")
	timeElements = append(timeElements, "03:04:05 PM")
	timeElements = append(timeElements, "03:04:05 P.M.")

	timeElements = append(timeElements, "3:4:5pm")
	timeElements = append(timeElements, "3:4:5p.m.")
	timeElements = append(timeElements, "3:4:5 pm")
	timeElements = append(timeElements, "3:4:5 p.m.")

	timeElements = append(timeElements, "3:4:5PM")
	timeElements = append(timeElements, "3:4:5P.M.")
	timeElements = append(timeElements, "3:4:5 PM")
	timeElements = append(timeElements, "3:4:5 P.M.")

	timeElements = append(timeElements, "03:04pm")
	timeElements = append(timeElements, "03:04p.m.")
	timeElements = append(timeElements, "03:04 pm")
	timeElements = append(timeElements, "03:04 p.m.")

	timeElements = append(timeElements, "03:04PM")
	timeElements = append(timeElements, "03:04P.M.")
	timeElements = append(timeElements, "03:04 PM")
	timeElements = append(timeElements, "03:04 P.M.")

	timeElements = append(timeElements, "3:4pm")
	timeElements = append(timeElements, "3:4p.m.")
	timeElements = append(timeElements, "3:4 pm")
	timeElements = append(timeElements, "3:4 p.m.")

	timeElements = append(timeElements, "3:4PM")
	timeElements = append(timeElements, "3:4P.M.")
	timeElements = append(timeElements, "3:4 PM")
	timeElements = append(timeElements, "3:4 P.M.")

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
