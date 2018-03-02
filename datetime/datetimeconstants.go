package datetime

import "time"

// Date Time Constants
const (
	// FmtDateTimeSecondStr - Date Time format used
	// for file names and directory names
	FmtDateTimeSecondStr = "20060102150405"

	// FmtDateTimeNanoSecondStr - Custom Date Time Format
	FmtDateTimeNanoSecondStr = "2006-01-02 15:04:05.000000000"

	// FmtDateTimeSecText - Custom Date Time Format
	FmtDateTimeSecText = "2006-01-02 15:04:05"

	// FmtDateTimeTzNano - Outputs date time to nano seconds with associated time zone
	FmtDateTimeTzNano = "01/02/2006 15:04:05.000000000 -0700 MST"

	// FmtDateTimeTzNanoYMD - Outputs date time to nano seconds with Year-Month-Date
	FmtDateTimeTzNanoYMD = "2006-01-02 15:04:05.000000000 -0700 MST"

	// FmtDateTimeTzNanoDowYMD - Output date time to nano seconds with Year-Month-Date
	// prefixed by day of the week
	FmtDateTimeTzNanoDowYMD = "Monday 2006-01-02 15:04:05.000000000 -0700 MST"

	// FmtDateTimeTzNanoYMDDow - Output date time to nano seconds with Year-Month-Date
	// prefixed by day of the week
	FmtDateTimeTzNanoYMDDow = "2006-01-02 Monday 15:04:05.000000000 -0700 MST"

	// FmtDateTimeYMDAbbrvDowNano - Output date time to nano seconds with abbreviated
	// day of week.
	FmtDateTimeYMDAbbrvDowNano = "2006-01-02 Mon 15:04:05.000000000 -0700 MST"

	// FmtDateTimeTzSec - Outputs date time to seconds with associated time zone
	FmtDateTimeTzSec = "01/02/2006 15:04:05 -0700 MST"

	// FmtDateTimeEverything - Custom Date Time Format showing virtually
	// all elements of a date time string.
	FmtDateTimeEverything = "Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	// FmtDateTimeNeutralDateFmt - Neutral Date Format without Time Zone
	FmtDateTimeNeutralDateFmt = "2006-01-02 15:04:05.000000000"

	// FmtDateTimeMDYrFmtStr - Month Day Year Date Format String
	FmtDateTimeMDYrFmtStr = "01/02/2006 15:04:05.000000000 -0700 MST"

	// FmtDateTimeYrMDayFmtStr - Year Month Day Date Format String
	FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
)


const (
	// Note: A Nanosecond is equal to 1 one-billionth or
	//       1/1,000,000,000 of a second.

	// MilliSecondsPerSecond - Number of Milliseconds in a Second
	MilliSecondsPerSecond = int64(1000)

	// MicroSecondsPerMilliSecond - The Number of Microseconds in
	// a Millisecond.
	MicroSecondsPerMilliSecond = int64(1000)

	// NanoSecondsPerMicroSecond - The number of nanoseconds in
	// a microsecond.
	NanoSecondsPerMicroSecond = int64(1000)


	// MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
	//	 A millisecond is 1/1,000 or 1 one-thousandth of a second
	MilliSecondNanoseconds = int64(time.Millisecond)

	// MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond
	// 	A MicroSecond is 1/1,000,000 or 1 one-millionth of a second
	MicroSecondNanoseconds = int64(time.Microsecond)

	// SecondNanoseconds - Number of Nanoseconds in a Second
	SecondNanoseconds = int64(time.Second)

	// MinuteNanoSeconds - Number of Nanoseconds in a minute
	MinuteNanoSeconds = int64(time.Minute)

	// HourNanoSeconds - Number of Nanoseconds in an hour
	HourNanoSeconds = int64(time.Hour)

	// DayNanoSeconds - Number of Nanoseconds in a 24-hour day
	DayNanoSeconds = int64(24) * HourNanoSeconds

	// WeekNanoSeconds - Number of Nanoseconds in a 7-day week
	WeekNanoSeconds = int64(7) * DayNanoSeconds

	/*
		For the Gregorian calendar the average length of the calendar year
		(the mean year) across the complete leap cycle of 400 Years is 365.2425 days.
		The Gregorian Average Year is therefore equivalent to 365 days, 5 hours,
		49 minutes and 12 seconds.
		Sources:
		https://en.wikipedia.org/wiki/Year
		Source: https://en.wikipedia.org/wiki/Gregorian_calendar
	*/

	// GregorianYearNanoSeconds - Number of Nano Seconds in a
	// Gregorian Year
	GregorianYearNanoSeconds = int64(31556952000000000)
)




// NOTE: See https://golang.org/pkg/time/#LoadLocation
// and https://www.iana.org/time-zones to ensure that
// the IANA Time Zone Database is properly configured
// on your system. Note: IANA Time Zone Data base is
// equivalent to 'tz database'.
//
// Reference: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
const (

	TzIanaAfricaCairo = "Africa/Cairo"
	TzIanaAfricaJohannesburg = "Africa/Johannesburg"
	TzIanaAfricaTripoli = "Africa/Tripoli"
	TzIanaAfricaTunis = "Africa/Tunis"
	TzIanaAmericaJamaica = "America/Jamaica"
	TzIanaAmericaMatamoros = "America/Matamoros"
	TzIanaAmericaMexicoCity = "America/Mexico_City"
	TzIanaAmericaPuertoRico = "America/Puerto_Rico"
	TzIanaAmericaTijuana = "America/Tijuana"
	TzIanaAmericaBuenosAires = "America/Argentina/Buenos_Aires"
	TzIanaAmericaBogota = "America/Bogota"
	TzIanaAmericaCancum = "America/Cancun"
	TzIanaAmericaCaracas = "America/Caracas"
	TzIanaAmericaCostaRica = "America/Costa_Rica"
	TzIanaAmericaElSalvador = "America/El_Salvador"
	TzIanaAmericaGooseBay = "America/Goose_Bay"
	TzIanaAmericaPortOfSpain ="America/Port_of_Spain" // Grenada
	TzIanaAmericaGuatemala = "America/Guatemala"
	TzIanaAmericaGuayaquil = "America/Guayaquil" // Ecuador
	TzIanaAmericaGuyana = "America/Guyana"
	TzIanaAmericaHalifax = "America/Halifax"
	TzIanaAmericaHavana = "America/Havana"
	TzIanaAmericaLaPaz = "America/La_Paz"
	TzIanaAmericaLima = "America/Lima"
	TzIanaAmericaManaus = "America/Manaus"  // Amazonas East
	TzIanaAmericaMartinique = "America/Martinique"
	TzIanaAmericaMazatlan = "America/Mazatlan" // Baja California
	TzIanaAmericaMotevideo = "America/Montevideo" //Uruguay
	TzIanaAmericaNassau = "America/Nassau" // Bahamas
	TzIanaAmericaPanama = "America/Panama"
	TzIanaAmericaRecife = "America/Recife"
	TzIanaAmericaSantiago = "America/Santiago"
	TzIanaAmericaSantoDomingo = "America/Santo_Domingo"
	TzIanaAmericaSaoPaulo = "America/Sao_Paulo"
	TzIanaAmericaStJohns = "America/St_Johns" // Newfoundland Labrador
	TzIanaAmericaStThomas = "America/St_Thomas"
	TzIanaAmericaThule = "America/Thule"
	TzIanaAmericaToronto = "America/Toronto" // Eastern - ON, QC (most areas)
	TzIanaAmericaVancouver = "America/Vancouver" // Pacific - BC (most areas)
	TzIanaAmericaWinnipeg = "America/Winnipeg" // Central - ON (west); Manitoba
	TzIanaAmericaWhitehorse = "America/Whitehorse" // Pacific - Yukon (south)
	TzIanaAntarcticaMcMurdo = "Antarctica/McMurdo"
	TzIanaAntarcticaSoutPole = "Pacific/Auckland"
	TzIanaAsiaBankok = "Asia/Bangkok"
	TzIanaAsiaBaghdad = "Asia/Baghdad"
	TzIanaAsiaBahrain = "Asia/Bahrain"
	TzIanaAsiaBaku = "Asia/Baku"
	TzIanaAsia = "Asia/Brunei"
	TzIanaAsiaBeirut = "Asia/Beirut"
	TzIanaAsiaDamasucs = "Asia/Damascus"
	TzIanaAsiaDubai = "Asia/Dubai"
	TzIanaAsiaHoChiMinh = "Asia/Ho_Chi_Minh"			// Saigon Vietnam
	TzIanaAsiaHongKong = "Asia/Hong_Kong"
	TzIanaAsiaIndia = "Asia/Kolkata"						// Formerly Calcutta  - India Time
	TzIanaAsiaJakarta = "Asia/Jakarta"
	TzIanaAsiaJerusalem = "Asia/Jerusalem"
	TzIanaAsiaKabul = "Asia/Kabul"
	TzIanaAsiaKarachi = "Asia/Karachi"
	TzIanaAsiaKualaLumpur = "Asia/Kuala_Lumpur"
	TzIanaAsiaKuwait = "Asia/Kuwait"
	TzIanaAsiaManila = "Asia/Manila"
	TzIanaAsiaPhnomPenh = "Asia/Phnom_Penh"
	TzIanaAsiaPyongyang = "Asia/Pyongyang"
	TzIanaAsiaQatar = "Asia/Qatar"
	TzIanaAsiaRangoon = "Asia/Yangon" 							// Rangoon
	TzIanaAsiaRiyadh = "Asia/Riyadh"
	TzIanaAsiaSeoul = "Asia/Seoul"
	TzIanaAsiaShanghai = "Asia/Shanghai"						// Beijing time
	TzIanaAsiaSigapore = "Asia/Singapore"
	TzIanaAsiaTaipei = "Asia/Taipei"
	TzIanaAsiaTehran = "Asia/Tehran"
	TzIanaAsiaTokyo = "Asia/Tokyo"
	TzIanaAsiaVladivostok = "Asia/Vladivostok"
	TzIanaAtlanticAzores = "Atlantic/Azores"
	TzIanaAtlanticBermuda = "Atlantic/Bermuda"
	TzIanaAtlanticCanary = "Atlantic/Canary"
	TzIanaAtlanticCapeVerde = "Atlantic/Cape_Verde"
	TzIanaAtlanticReykjavik = "Atlantic/Reykjavik"
	TzIanaAtlanticStanley = "Atlantic/Stanley"
	TzIanaAustraliaDarwin = "Australia/Darwin"
	TzIanaAustraliaMelbourne = "Australia/Melbourne"
	TzIanaAustraliaSydney = "Australia/Sydney"
	TzIanaAustraliaPerth = "Australia/Perth"
	TzIanaEuropeAmsterdam = "Europe/Amsterdam"
	TzIanaEuropeAthens = "Europe/Athens"
	TzIanaEuropeBelgrade = "Europe/Belgrade"
	TzIanaEuropeBerlin = "Europe/Berlin"
	TzIanaEuropeBrussels = "Europe/Brussels"
	TzIanaEuropeBucharest = "Europe/Bucharest"
	TzIanaEuropeBudapest = "Europe/Budapest"
	TzIanaEuropeCopenhagen = "Europe/Copenhagen"
	TzIanaEuropeDublin = "Europe/Dublin"
	TzIanaEuropeGibraltar = "Europe/Gibraltar"
	TzIanaEuropeHelsinki = "Europe/Helsinki"
	TzIanaEuropeIstanbul = "Europe/Istanbul"
	TzIanaEuropeKiev = "Europe/Kiev"
	TzIanaEuropeLisbon = "Europe/Lisbon"
	TzIanaEuropeLondon = "Europe/London"
	TzIanaEuropeLuxembourg = "Europe/Luxembourg"
	TzIanaEuropeMadrid = "Europe/Madrid"
	TzIanaEuropeMalta = "Europe/Malta"
	TzIanaEuropeMinsk = "Europe/Minsk"
	TzIanaEuropeMonaco = "Europe/Monaco"
	TzIanaEuropeMoscow = "Europe/Moscow"
	TzIanaEuropeOslo = "Europe/Oslo"
	TzIanaEuropeParis = "Europe/Paris"
	TzIanaEuropePrague = "Europe/Prague"
	TzIanaEuropeRiga = "Europe/Riga"
	TzIanaEuropeRome = "Europe/Rome"
	TzIanaEuropeSofia = "Europe/Sofia"
	TzIanaEuropeStockholm = "Europe/Stockholm"
	TzIanaEuropeVienna = "Europe/Vienna"
	TzIanaEuropeVilnius = "Europe/Vilnius"
	TzIanaEuropeWarsaw = "Europe/Warsaw"
	TzIanaEuropeZurich ="Europe/Zurich"
	TzIanaPacificAuckland = "Pacific/Auckland"
	TzIanaPacificFiji = "Pacific/Fiji"
	TzIanaPacificGuam = "Pacific/Guam"
	TzIanaPacificHonolulu = "Pacific/Honolulu"
	TzIanaPacificPortMoresby = "Pacific/Port_Moresby"
	TzIanaPacificTahiti = "Pacific/Tahiti"

	// TzIanaUsAlaska - USA Alaska
	TzIanaUsAlaskaAnchorage =  "America/Anchorage"
	TzIanaUsAlaskaJuneau = "America/Juneau"
	TzIanaUsAlaskaNome = "America/Nome"
	TzIanaUsAlaskaYakutat = "America/Yakutat"
	TzIanaUsAleutianIslands = "America/Adak"

	// TzIanaUsArizona
	TzIanaUsArizona = "America/Phoenix"

	// TzIanaUsEast - USA Eastern Time Zone
	// IANA database identifier
	TzIanaUsEast = "America/New_York"

	// TzIanaUsCentral - USA Central Time Zone
	// IANA database identifier
	TzIanaUsCentral = "America/Chicago"

	// TzIanaUsMountain - USA Mountain Time Zone
	// IANA database identifier
	TzIanaUsMountain = "America/Denver"

	// TzIanaUsPacific - USA Pacific Time Zone
	// IANA database identifier
	TzIanaUsPacific  = "America/Los_Angeles"

	// TzIanaUsHawaii - USA Hawaiian Time Zone
	// IANA database identifier
	TzIanaUsHawaii = "Pacific/Honolulu"

	// TzIanaZulu - UTC Time Zone IANA database
	// identifier
	TzIanaZulu = "Etc/UCT"

	// TzIanaGMT - Alias for UTC
	TzIanaGMT	= "Etc/UCT"

	// TzIanaUTC - Alias for UTC
	TzIanaUTC = "Etc/UCT"


)

