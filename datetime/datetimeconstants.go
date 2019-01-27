/*
package datetime provides a variety of types and associated methods designed
to manage dates, times, timezones, date time string formats and time durations.
In addition, associated methods perform date time addition and subtraction.

All types in this package rely on the IANA Time Zone Database.

Reference:

	https://www.iana.org/time-zones

*/
package datetime

import "time"

/*
 Date Time Constants


 This source file is located in source code repository:
 		'https://github.com/MikeAustin71/datetimeopsgo.git'

 This source code file is located at:
		'MikeAustin71\datetimeopsgo\datetime\datetimeconstants.go'



Overview and General Usage


This source file contains a series of constants useful in managing
date time.

Types of constants defined here include:
	1. Date Time string formats
  2. Iana Time Zone designations
  3. Common Time conversion constants

*/
const (
	// Date Time Format Constants
	// ================================================================================
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

	// FmtDateTimeUsMilitaryDate2DYr
	FmtDateTimeUsMilitaryDate2DYr = "021504Z 06"

	// FmtDateTimeYrMDayFmtStr - Year Month Day Date Format String
	FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
)

const (
	// Time Constants
	// ================================================================================
	// Note: A Nanosecond is equal to 1 one-billionth or
	//       1/1,000,000,000 of a second.
	//
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




// GregorianYearNanoSeconds - Number of Nano Seconds in a
// Gregorian Year.
//
// For the Gregorian calendar the average length of the calendar year
// (the mean year) across the complete leap cycle of 400 Years is 365.2425 days.
// The Gregorian Average Year is therefore equivalent to 365 days, 5 hours,
// 49 minutes and 12 seconds.
//
// Sources:
//	https://en.wikipedia.org/wiki/Year
//	https://en.wikipedia.org/wiki/Gregorian_calendar
GregorianYearNanoSeconds = int64(31556952000000000)
)

const (
	// TimeZone Constants
	// NOTE: See https://golang.org/pkg/time/#LoadLocation
	// and https://www.iana.org/time-zones to ensure that
	// the IANA Time Zone Database is properly configured
	// on your system. Note: IANA Time Zone Data base is
	// equivalent to 'tz database'.
	//
	// Reference: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	TzIanaAfricaCairo         = "Africa/Cairo"
	TzIanaAfricaJohannesburg  = "Africa/Johannesburg"
	TzIanaAfricaTripoli       = "Africa/Tripoli"
	TzIanaAfricaTunis         = "Africa/Tunis"
	TzIanaAmericaBuenosAires  = "America/Argentina/Buenos_Aires"
	TzIanaAmericaUshuaia      = "America/Argentina/Ushuaia"      // Tierra del Fuego
	TzIanaAmericaBogota       = "America/Bogota"
	TzIanaAmericaCancum       = "America/Cancun"
	TzIanaAmericaCaracas      = "America/Caracas"
	TzIanaAmericaCayman       = "America/Cayman"
	TzIanaAmericaChihuahua    = "America/Chihuahua"
	TzIanaAmericaCostaRica    = "America/Costa_Rica"
	TzIanaAmericaDawson       = "America/Dawson"         // Yukon
	TzIanaAmericaEirunepe     = "America/Eirunepe"       // Western Brazil
	TzIanaAmericaElSalvador   = "America/El_Salvador"
	TzIanaAmericaGooseBay     = "America/Goose_Bay"
  TzIanaAmericaGodthab      = "America/Godthab"        // Most of Greenland - Nuuk
	TzIanaAmericaGuatemala    = "America/Guatemala"
	TzIanaAmericaGuayaquil    = "America/Guayaquil"      // Ecuador
	TzIanaAmericaGuyana       = "America/Guyana"
	TzIanaAmericaHalifax      = "America/Halifax"
	TzIanaAmericaHavana       = "America/Havana"
	TzIanaAmericaJamaica      = "America/Jamaica"
	TzIanaAmericaLaPaz        = "America/La_Paz"
	TzIanaAmericaLima         = "America/Lima"
	TzIanaAmericaManaus       = "America/Manaus" // Amazonas East
	TzIanaAmericaMartinique   = "America/Martinique"
	TzIanaAmericaMazatlan     = "America/Mazatlan" // Baja California
	TzIanaAmericaMatamoros    = "America/Matamoros"
	TzIanaAmericaMexicoCity   = "America/Mexico_City"
	TzIanaAmericaMotevideo    = "America/Montevideo" // Uruguay
	TzIanaAmericaNassau       = "America/Nassau"     // Bahamas
	TzIanaAmericaPanama       = "America/Panama"
	TzIanaAmericaPortOfSpain  = "America/Port_of_Spain" // Grenada
	TzIanaAmericaPuertoRico   = "America/Puerto_Rico"
	TzIanaAmericaRecife       = "America/Recife"
	TzIanaAmericaSantiago     = "America/Santiago"
	TzIanaAmericaSantoDomingo = "America/Santo_Domingo"
	TzIanaAmericaSaoPaulo     = "America/Sao_Paulo"
	TzIanaAmericaScoresbysund = "America/Scoresbysund"
	TzIanaAmericaStJohns      = "America/St_Johns" // Newfoundland Labrador
	TzIanaAmericaStThomas     = "America/St_Thomas"
	TzIanaAmericaTijuana      = "America/Tijuana"
	TzIanaAmericaThule        = "America/Thule"      // Western Greenland
	TzIanaAmericaToronto      = "America/Toronto"    // Eastern - ON, QC (most areas)
	TzIanaAmericaVancouver    = "America/Vancouver"  // Pacific - BC (most areas)
	TzIanaAmericaWinnipeg     = "America/Winnipeg"   // Central - ON (west); Manitoba
	TzIanaAmericaWhitehorse   = "America/Whitehorse" // Pacific - Yukon (south)
	TzIanaAntarcticaMcMurdo   = "Antarctica/McMurdo"
	TzIanaAntarcticaSouthPole = "Pacific/Auckland"
	TzIanaAsiaAlmaty          = "Asia/Almaty"        // Kazakhstan
	TzIanaAsiaAshgabat        = "Asia/Ashgabat"      // Turkmenistan
	TzIanaAsiaBankok          = "Asia/Bangkok"
	TzIanaAsiaBaghdad         = "Asia/Baghdad"
	TzIanaAsiaBahrain         = "Asia/Bahrain"
	TzIanaAsiaBaku            = "Asia/Baku"
	TzIanaAsia                = "Asia/Bishkek"     // Kyrgyzstan
	TzIanaAsiaBrunei          = "Asia/Brunei"
	TzIanaAsiaBeirut          = "Asia/Beirut"
	TzIanaAsiaDamasucs        = "Asia/Damascus"
	TzIanaAsiaDhaka           = "Asia/Dhaka"	     // Bangladesh
	TzIanaAsiaDubai           = "Asia/Dubai"
	TzIanaAsiaDushanbe        = "Asia/Dushanbe"    // Tajikistan
	TzIanaAsiaHoChiMinh       = "Asia/Ho_Chi_Minh" // Formerly Saigon Vietnam
	TzIanaAsiaHongKong        = "Asia/Hong_Kong"
	TzIanaAsiaIndia           = "Asia/Kolkata"     // Formerly Calcutta  - India Time
	TzIanaAsiaJakarta         = "Asia/Jakarta"
	TzIanaAsiaJerusalem       = "Asia/Jerusalem"
	TzIanaAsiaKabul           = "Asia/Kabul"       // Afghanistan +4:30
	TzIanaAsiaKarachi         = "Asia/Karachi"
	TzIanaAsiaKualaLumpur     = "Asia/Kuala_Lumpur"
	TzIanaAsiaKuwait          = "Asia/Kuwait"
	TzIanaAsiaManila          = "Asia/Manila"
	TzIanaAsiaPhnomPenh       = "Asia/Phnom_Penh"
	TzIanaAsiaPyongyang       = "Asia/Pyongyang"
	TzIanaAsiaQatar           = "Asia/Qatar"
	TzIanaAsiaRangoon         = "Asia/Yangon"        // Rangoon
	TzIanaAsiaRiyadh          = "Asia/Riyadh"
	TzIanaAsiaSeoul           = "Asia/Seoul"
	TzIanaAsiaShanghai        = "Asia/Shanghai"      // Beijing time
	TzIanaAsiaSigapore        = "Asia/Singapore"
	TzIanaAsiaTaipei          = "Asia/Taipei"
  TzIanaAsiaTashkent        = "Asia/Tashkent"      // Uzbekistan
	TzIanaAsiaTehran          = "Asia/Tehran"
	TzIanaAsiaTokyo           = "Asia/Tokyo"
	TzIanaAsiaUlaanbaatar     = "Asia/Ulaanbaatar"   // Mongolia
	TzIanaAsiaVladivostok     = "Asia/Vladivostok"
	TzIanaAsiaYakutsk         = "Asia/Yakutsk"
	TzIanaAtlanticAzores      = "Atlantic/Azores"
	TzIanaAtlanticBermuda     = "Atlantic/Bermuda"
	TzIanaAtlanticCanary      = "Atlantic/Canary"
	TzIanaAtlanticCapeVerde   = "Atlantic/Cape_Verde"
	TzIanaAtlanticReykjavik   = "Atlantic/Reykjavik"
	TzIanaAtlanticStanley     = "Atlantic/Stanley"
	TzIanaAustraliaAdelaide   = "Australia/Adelaide"
	TzIanaAustraliaBrisbane   = "Australia/Brisbane"
	TzIanaAustraliaBrokenHill = "Australia/Broken_Hill"
	TzIanaAustraliaCurrie     = "Australia/Currie"
	TzIanaAustraliaDarwin     = "Australia/Darwin"
	TzIanaAustraliaHobart     = "Australia/Hobart"
	TzIanaAustraliaLindeman   = "Australia/Lindeman"
	TzIanaAustraliaLordHowe   = "Australia/Lord_Howe"
	TzIanaAustraliaMelbourne  = "Australia/Melbourne"
	TzIanaAustraliaPerth      = "Australia/Perth"
	TzIanaAustraliaSydney     = "Australia/Sydney"
	TzIanaEuropeAmsterdam     = "Europe/Amsterdam"
	TzIanaEuropeAthens        = "Europe/Athens"
	TzIanaEuropeBelfast       = "Europe/Belfast"
	TzIanaEuropeBelgrade      = "Europe/Belgrade"
	TzIanaEuropeBerlin        = "Europe/Berlin"
	TzIanaEuropeBratislava    = "Europe/Bratislava"  // Slovakia
	TzIanaEuropeBrussels      = "Europe/Brussels"
	TzIanaEuropeBucharest     = "Europe/Bucharest"
	TzIanaEuropeBudapest      = "Europe/Budapest"
	TzIanaEuropeChisinau      = "Europe/Chisinau"    // Moldova
	TzIanaEuropeCopenhagen    = "Europe/Copenhagen"
	TzIanaEuropeDublin        = "Europe/Dublin"
	TzIanaEuropeGibraltar     = "Europe/Gibraltar"
	TzIanaEuropeHelsinki      = "Europe/Helsinki"
	TzIanaEuropeIstanbul      = "Europe/Istanbul"
	TzIanaEuropeKaliningrad   = "Europe/Kaliningrad"
	TzIanaEuropeKiev          = "Europe/Kiev"
	TzIanaEuropeLisbon        = "Europe/Lisbon"
	TzIanaEuropeLondon        = "Europe/London"
	TzIanaEuropeLuxembourg    = "Europe/Luxembourg"
	TzIanaEuropeMadrid        = "Europe/Madrid"
	TzIanaEuropeMalta         = "Europe/Malta"
	TzIanaEuropeMinsk         = "Europe/Minsk"
	TzIanaEuropeMonaco        = "Europe/Monaco"
	TzIanaEuropeMoscow        = "Europe/Moscow"
	TzIanaEuropeOslo          = "Europe/Oslo"
	TzIanaEuropeParis         = "Europe/Paris"
	TzIanaEuropePodgorica     = "Europe/Podgorica"   // Montenegro
	TzIanaEuropePrague        = "Europe/Prague"
	TzIanaEuropeRiga          = "Europe/Riga"
	TzIanaEuropeRome          = "Europe/Rome"
	TzIanaEuropeSarajevo      = "Europe/Sarajevo"    // Bosnia Herzegovina
	TzIanaEuropeSkopje        = "Europe/Skopje"      // Fryo Macedonia
	TzIanaEuropeSofia         = "Europe/Sofia"
	TzIanaEuropeStockholm     = "Europe/Stockholm"
	TzIanaEuropeTirane        = "Europe/Tirane"      // Albania
	TzIanaEuropeUzhgorod      = "Europe/Uzhgorod"    // Ukraine
	TzIanaEuropeVienna        = "Europe/Vienna"
	TzIanaEuropeVilnius       = "Europe/Vilnius"
	TzIanaEuropeVolgograd     = "Europe/Volgograd"
	TzIanaEuropeWarsaw        = "Europe/Warsaw"
	TzIanaEuropeZagreb        = "Europe/Zagreb"      // Croatia
	TzIanaEuropeZurich        = "Europe/Zurich"
	TzIanaPacificAuckland     = "Pacific/Auckland"   // New Zealand Most Areas
	TzIanaPacificChuuk        = "Pacific/Chuuk"
	TzIanaPacificFiji         = "Pacific/Fiji"
	TzIanaPacificGuam         = "Pacific/Guam"
	TzIanaPacificHonolulu     = "Pacific/Honolulu"
	TzIanaPacificPortMoresby  = "Pacific/Port_Moresby"
	TzIanaPacificTahiti       = "Pacific/Tahiti"

	// TzIanaUsAlaska - USA Alaska
	TzIanaUsAlaskaAnchorage = "America/Anchorage"
	TzIanaUsAlaskaJuneau    = "America/Juneau"
	TzIanaUsAlaskaNome      = "America/Nome"
	TzIanaUsAlaskaYakutat   = "America/Yakutat"
	TzIanaUsAleutianIslands = "America/Adak"

	// TzIanaUsArizona
	TzIanaUsArizona = "America/Phoenix"

	// TzIanaUsEast - USA Eastern Time Zone
	// IANA database identifier
	TzIanaUsEast = "America/New_York"

	// TzIanaUsCentral - USA Central Time Zone
	// IANA database identifier
	TzIanaUsCentral = "America/Chicago"


	TzIanaUsDetroit = "America/Detroit"         // Eastern Michigan

	// TzIanaUsMountain - USA Mountain Time Zone
	// IANA database identifier
	TzIanaUsMountain = "America/Denver"

	// TzIanaUsPacific - USA Pacific Time Zone
	// IANA database identifier
	TzIanaUsPacific = "America/Los_Angeles"

	// TzIanaUsHawaii - USA Hawaiian Time Zone
	// IANA database identifier
	TzIanaUsHawaii = "Pacific/Honolulu"

	// TzIanaZulu - UTC Time Zone IANA database
	// identifier
	TzIanaZulu = "Etc/UCT"

	// TzIanaGMT - Alias for UTC
	TzIanaGMT = "Etc/UCT"

	// TzIanaUTC - Alias for UTC
	TzIanaUTC = "Etc/UCT"

	// TzGoLocal - Golang Local Time Zone
	// configured on host computer
	TzGoLocal = "Local"
)

/*
 UTC Offsets
*/
var offsetsUTC = [25]int{-12,-11,-10,-9,-8,-7,-6,-5,-4,-3,-2,-1,0,1,2,3,4,5,6,7,8,9,10,11,12}

/*
Time Zone Abbreviations – Military Time Zone Names
https://www.timeanddate.com/time/zones/military

Military time zones are commonly used in aviation as well as at sea. They are also known as nautical or maritime time zones. J (Juliet Time Zone) is occasionally used to refer to the observer's local time.

Abbreviation	Time zone name 	Other names 	Offset

A	            Alpha Time Zone                 UTC +1
B	            Bravo Time Zone                 UTC +2
C	            Charlie Time Zone               UTC +3
D	            Delta Time Zone                 UTC +4
E	            Echo Time Zone                  UTC +5
F             Foxtrot Time Zone               UTC +6
G             Golf Time Zone                  UTC +7
H	            Hotel Time Zone                 UTC +8
I	            India Time Zone                 UTC +9
K	            Kilo Time Zone                  UTC +10
L	            Lima Time Zone                  UTC +11
M	            Mike Time Zone                  UTC +12
N	            November Time Zone              UTC -1
O	            Oscar Time Zone                 UTC -2
P	            Papa Time Zone                  UTC -3
Q	            Quebec Time Zone                UTC -4
R	            Romeo Time Zone                 UTC -5
S	            Sierra Time Zone                UTC -6
T	            Tango Time Zone                 UTC -7
U	            Uniform Time Zone               UTC -8
V	            Victor Time Zone                UTC -9
W	            Whiskey Time Zone 		          UTC -10
X	            X-ray Time Zone                 UTC -11
Y	            Yankee Time Zone                UTC -12
Z	            Zulu Time ZoneUTC                   +0

Military Time Code Letter Reference:

UTC -12: Y- (e.g. Fiji)
UTC-11: X (American Samoa)
UTC-10: W (Honolulu, HI)
UTC-9: V (Juneau, AK)
UTC-8: U (PST, Los Angeles, CA)
UTC-7: T (MST, Denver, CO)
UTC-6: S (CST, Dallas, TX)
UTC-5: R (EST, New York, NY)
UTC-4: Q (Halifax, Nova Scotia
UTC-3: P (Buenos Aires, Argentina)
UTC-2: O (Godthab, Greenland)
UTC-1: N (Azores)
UTC+-0: Z (Zulu time)
UTC+1: A (France)
UTC+2: B (Athens, Greece)
UTC+3: C (Arab Standard Time, Iraq, Bahrain, Kuwait, Saudi Arabia, Yemen, Qatar)
UTC+4: D (Used for Moscow, Russia and Afghanistan, however, Afghanistan is technically +4:30 from UTC)
UTC+5: E (Pakistan, Kazakhstan, Tajikistan, Uzbekistan and Turkmenistan)
UTC+6: F (Bangladesh)
UTC+7: G (Thailand)
UTC+8: H (Beijing, China)
UTC+9: I (Tokyo, Australia)
UTC+10: K (Brisbane, Australia)
UTC+11: L (Sydney, Australia)
UTC+12: M (Wellington, New Zealand)
 */
var abbreviatedUsMilitaryTimeZones = [25]string{
	"Y",
	"X",
	"W",
	"V",
	"U",
	"T",
	"S",
	"R",
	"Q",
	"P",
	"O",
	"N",
	"Z",
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"K",
	"L",
	"M"}


/*
http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm

UTC+12: M (Wellington, New Zealand)

 */
var IanaUsMilitaryTimeZone = [25] string{

	"Etc/GMT+12",            // UTC -12: Y
	"Pacific/Pago_Pago",     // UTC-11: X (American Samoa)
	"Pacific/Honolulu",      // UTC-10: W (Honolulu, HI)
	"America/Juneau",        // UTC-9: V (Juneau, AK)
	"America/Los_Angeles",   // UTC-8: U (PST, Los Angeles, CA)
	"America/Denver",        // UTC-7: T (MST, Denver, CO)
	"America/Chicago",       // UTC-6: S (CST, Dallas, TX)
	"America/New_York",      // UTC-5: R (EST, New York, NY)
	"America/Halifax",       // UTC-4: Q (Halifax, Nova Scotia
	"America/Argentina/Buenos_Aires", // UTC-3: P (Buenos Aires, Argentina)
	"America/Godthab",       // UTC-2: O (Godthab, Greenland)
	"Atlantic/Azores",       // UTC-1: N (Azores)
	"Etc/UCT",               // UTC+-0: Z (Zulu time)
	"Europe/Paris",          // UTC+1: A (France)
	"Europe/Athens",         // UTC+2: B (Athens, Greece)
	"Asia/Riyadh",           // UTC+3: C (Arab Standard Time, Iraq, Bahrain,
                           //        Kuwait, Saudi Arabia, Yemen, Qatar)

	"Asia/Dubai",            // UTC+4: D (Used for Moscow, Russia and Afghanistan,
                           //        however, Afghanistan is technically +4:30 from UTC)
	"Asia/Karachi",          // UTC+5: E (Pakistan, Kazakhstan, Tajikistan,
                           //        Uzbekistan and Turkmenistan)
	"Asia/Dhaka",            // UTC+6: F (Bangladesh)
	"Asia/Bangkok",          // UTC+7: G (Thailand)
	"Asia/Shanghai",         // UTC+8: H (Beijing, China)
	"Asia/Tokyo",            // UTC+9: I (Tokyo, Australia)
	"Australia/Brisbane",    // UTC+10: K (Brisbane, Australia)
	"Australia/Sydney",      // UTC+11: L (Sydney, Australia)
	"Pacific/Auckland" }     // UTC+12: M (Wellington, New Zealand)
