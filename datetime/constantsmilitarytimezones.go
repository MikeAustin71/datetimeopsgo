package datetime

import (
	"sync"
)

// MilitaryTimeZoneData - Provides thread safe access to Military
// Time Zone, location and nomenclature data.
//
// For information on Military Time Zones, reference:
//
//     https://en.wikipedia.org/wiki/List_of_military_time_zones
//     http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//     https://www.timeanddate.com/time/zones/military
//     https://www.timeanddate.com/worldclock/timezone/alpha
//     https://www.timeanddate.com/time/map/
//
// Military time zones are commonly used in military operations and
// aviation as well as at sea. They are also known as nautical or
// maritime time zones.
//
// The 'J' (Juliet) Time Zone is occasionally used to refer to the observer's
// local time. Note that Time Zone 'J' (Juliet) is not listed below.
//
//    Time Zone       Time Zone        Equivalent IANA          UTC
//   Abbreviation       Name              Time Zone            Offset
//   ------------     --------          ---------------        ------
//
//       A        Alpha Time Zone         Etc/GMT-1            UTC +1
//       B        Bravo Time Zone         Etc/GMT-2            UTC +2
//       C        Charlie Time Zone       Etc/GMT-3            UTC +3
//       D        Delta Time Zone         Etc/GMT-4            UTC +4
//       E        Echo Time Zone          Etc/GMT-5            UTC +5
//       F        Foxtrot Time Zone       Etc/GMT-6            UTC +6
//       G        Golf Time Zone          Etc/GMT-7            UTC +7
//       H        Hotel Time Zone         Etc/GMT-8            UTC +8
//       I        India Time Zone         Etc/GMT-9            UTC +9
//       K        Kilo Time Zone          Etc/GMT-10           UTC +10
//       L        Lima Time Zone          Etc/GMT-11           UTC +11
//       M        Mike Time Zone          Etc/GMT-12           UTC +12
//       N        November Time Zone      Etc/GMT+1            UTC -1
//       O        Oscar Time Zone         Etc/GMT+2            UTC -2
//       P        Papa Time Zone          Etc/GMT+3            UTC -3
//       Q        Quebec Time Zone        Etc/GMT+4            UTC -4
//       R        Romeo Time Zone         Etc/GMT+5            UTC -5
//       S        Sierra Time Zone        Etc/GMT+6            UTC -6
//       T        Tango Time Zone         Etc/GMT+7            UTC -7
//       U        Uniform Time Zone       Etc/GMT+8            UTC -8
//       V        Victor Time Zone        Etc/GMT+9            UTC -9
//       W        Whiskey Time Zone       Etc/GMT+10           UTC -10
//       X        X-ray Time Zone         Etc/GMT+11           UTC -11
//       Y        Yankee Time Zone        Etc/GMT+12           UTC -12
//       Z        Zulu Time Zone          UTC                  UTC +0
//
//  UTC     Time Zone     Time Zone
// Offset  Abbreviation   Location
// ------  ------------   -----------------------------------------------------------------------
// UTC+1         A        (France)
// UTC+2         B        (Athens, Greece)
// UTC+3         C        (Arab Standard Time, Iraq, Bahrain, Kuwait, Saudi Arabia, Yemen, Qatar)
// UTC+4         D        (Used for Moscow, Russia and Afghanistan, however, Afghanistan is
//                           technically +4:30 from UTC)
// UTC+5         E        (Pakistan, Kazakhstan, Tajikistan, Uzbekistan and Turkmenistan)
// UTC+6         F        (Bangladesh)
// UTC+7         G        (Thailand)
// UTC+8         H        (Beijing, China)
// UTC+9         I        (Tokyo, Australia)
// UTC+10        K        (Brisbane, Australia)
// UTC+11        L        (Sydney, Australia)
// UTC+12        M        (Wellington, New Zealand)
// UTC-1         N        (Azores)
// UTC-2         O        (Godthab, Greenland)
// UTC-3         P        (Buenos Aires, Argentina)
// UTC-4         Q        (Halifax, Nova Scotia)
// UTC-5         R        (EST, New York, NY)
// UTC-6         S        (CST, Dallas, TX)
// UTC-7         T        (MST, Denver, CO)
// UTC-8         U        (PST, Los Angeles, CA)
// UTC-9         V        (Juneau, AK)
// UTC-10        W        (Honolulu, HI)
// UTC-11        X        (American Samoa)
// UTC -12       Y        (e.g. Fiji)
// UTC+-0        Z        (Zulu time)
//
type MilitaryTimeZoneData struct{
	Input   string
	Output  string
	lock    sync.Mutex
}

// MilTzLetterToTextName - Returns the Military Time Zone
// text name based on an input string parameter ,'milTzLetter',
// which contains a single alphabetic character identifying
// the military time zone.
//
// Example:
//  "Z" returns military time zone name "Zulu".
//  "A" returns military time zone name "Alpha".
//  "B" returns military time zone name "Bravo".
//
// This method accesses map 'militaryTzLetterToTxtNameMap'.
//
// If the input parameter 'milTzLetter' is invalid, the boolean
// value returned by this method will be set to 'false'.
//
func (milTzDat *MilitaryTimeZoneData) MilTzLetterToTextName(
	milTzLetter string) (string, bool) {

	milTzDat.lock.Lock()

	defer milTzDat.lock.Unlock()

	lockMilitaryTzLetterToTxtNameMap.Lock()

	defer lockMilitaryTzLetterToTxtNameMap.Unlock()

		result, ok := militaryTzLetterToTxtNameMap[milTzLetter]

	return result, ok
}

// MilitaryTzToIanaTz - Returns an IANA Time Zone which
// is equivalent to the military time zone identified by
// the input parameter 'milTzTextName'.
//
// Example:
//  Military Time Zone = "Charlie"
//  Returns Equivalent IANA Time Zone = "Etc/GMT-3"
//
// This method accesses map 'militaryTzToIanaTzMap'.
//
// If successful, the equivalent IANA Time Zone will be
// returned as a string.
//
// If input parameter 'milTzTextName' is invalid, the boolean
// value returned by this method will be set to false.
//
func (milTzDat *MilitaryTimeZoneData) MilitaryTzToIanaTz(
		milTzTextName string) (string, bool) {

	milTzDat.lock.Lock()

	defer milTzDat.lock.Unlock()

	lockMilitaryTzToIanaTzMap.Lock()

	defer lockMilitaryTzToIanaTzMap.Unlock()

	result, ok := militaryTzToIanaTzMap[milTzTextName]

	return result, ok
}

// UtcOffsetToMilitaryTimeZone - Returns the military time zone
// associated with the UTC offset identified by the input parameter
// 'utcOffset'.
//
// Input parameter 'utcOffset' should be formatted in accordance
// with the following examples:
//  "+0200"
//  "-0700"
//  "+1100"
//
// This method accesses map 'militaryUTCToTzMap'.
//
// If successful, this method will return a string containing the
// full Military Time Zone name. Examples of full Military Time
// Zone names are "Alpha", "Bravo", "Charlie", "Golf" etc.
//
// If the input parameter utcOffset is invalid, the boolean
// value returned by this method is set to 'false'.
//
func (milTzDat *MilitaryTimeZoneData) UtcOffsetToMilitaryTimeZone(
	utcOffset string) (string, bool) {

	milTzDat.lock.Lock()

	defer milTzDat.lock.Unlock()

	lockMilitaryUTCToTzMap.Lock()

	defer lockMilitaryUTCToTzMap.Unlock()

	result, ok := militaryUTCToTzMap[utcOffset]

	return result, ok
}

// MilitaryTzToUtc - Returns the UTC offset associated
// with a Military Time Zone Name passed in by input
// parameter 'milTzTextName'.
//
// Input parameter is a string containing the full
// Military Time Zone Name. Examples: "Alpha", "Bravo"
// "Charlie", "Delta", etc.
//
// This method accesses map 'militaryTzToUTCMap'.
//
// If successful, the returned string will contain an
// UTC offset formatted like: "+0700", "-0100", "+0000"
// etc.
//
// If the input parameter is invalid, the boolean value
// returned by this method will be set to 'false'.
//
func (milTzDat *MilitaryTimeZoneData) MilitaryTzToUtc(
	milTzTextName string) (string, bool) {

	milTzDat.lock.Lock()

	defer milTzDat.lock.Unlock()

	lockMilitaryTzToUTCMap.Lock()

	defer lockMilitaryTzToUTCMap.Unlock()

	result, ok := militaryTzToUTCMap[milTzTextName]

	return result, ok
}

// MilitaryTzToLocation - This method returns a brief geographical
// location description associated Military Time Zone name passed
// by input parameter 'milTzTextName'.
//
// Input parameter string 'milTzTextName' must contain a full and
// valid Military Time Zone text name like, "Alpha", "Bravo",
// "Charlie", "Zulu", etc.
//
// If input parameter 'milTzTextName' is invalid, the boolean value
// returned by this method is set to 'false'.
//
func (milTzDat *MilitaryTimeZoneData) MilitaryTzToLocation(
	milTzTextName string) (string, bool) {

	milTzDat.lock.Lock()

	defer milTzDat.lock.Unlock()

	lockMilitaryTzToLocationMap.Lock()

	defer lockMilitaryTzToLocationMap.Unlock()

	result, ok := militaryTzToLocationMap[milTzTextName]

	return result, ok
}

var lockMilitaryTzLetterToTxtNameMap sync.Mutex

var militaryTzLetterToTxtNameMap = map[string]string {
	"A" : "Alpha",
	"B" : "Bravo",
	"C" : "Charlie",
	"D" : "Delta",
	"E" : "Echo",
	"F" : "Foxtrot",
	"G" : "Golf",
	"H" : "Hotel",
	"I" : "India",
	"K" : "Kilo",
	"L" : "Lima",
	"M" : "Mike",
	"N" : "November",
	"O" : "Oscar",
	"P" : "Papa",
	"Q" : "Quebec",
	"R" : "Romeo",
	"S" : "Sierra",
	"T" : "Tango",
	"U" : "Uniform",
	"V" : "Victor",
	"W" : "Whiskey",
	"X" : "Xray",
	"Y" : "Yankee",
	"Z" : "Zulu" }

var lockMilitaryTzToIanaTzMap sync.Mutex

var militaryTzToIanaTzMap = map[string]string {
	"Alpha":    TZones.Military.Alpha(),
	"Bravo":    TZones.Military.Bravo(),
	"Charlie":  TZones.Military.Charlie(),
	"Delta":    TZones.Military.Delta(),
	"Echo":     TZones.Military.Echo(),
	"Foxtrot":  TZones.Military.Foxtrot(),
	"Golf":     TZones.Military.Golf(),
	"Hotel":    TZones.Military.Hotel(),
	"India":    TZones.Military.India(),
	"Kilo":     TZones.Military.Kilo(),
	"Lima":     TZones.Military.Lima(),
	"Mike":     TZones.Military.Mike(),
	"November": TZones.Military.November(),
	"Oscar":    TZones.Military.Oscar(),
	"Papa":     TZones.Military.Papa(),
	"Quebec":   TZones.Military.Quebec(),
	"Romeo":    TZones.Military.Romeo(),
	"Sierra":   TZones.Military.Sierra(),
	"Tango":    TZones.Military.Tango(),
	"Uniform":  TZones.Military.Uniform(),
	"Victor":   TZones.Military.Victor(),
	"Whiskey":  TZones.Military.Whiskey(),
	"Xray":     TZones.Military.Xray(),
	"Yankee":   TZones.Military.Yankee(),
	"Zulu":     TZones.Military.Zulu()}


var lockMilitaryUTCToTzMap sync.Mutex

var militaryUTCToTzMap = map[string]string{
	"+0100":    "Alpha",
	"+0200":    "Bravo",
	"+0300":    "Charlie",
	"+0400":    "Delta",
	"+0500":    "Echo",
	"+0600":    "Foxtrot",
	"+0700":    "Golf",
	"+0800":    "Hotel",
	"+0900":    "India",
	"+1000":    "Kilo",
	"+1100":    "Lima",
	"+1200":    "Mike",
	"-0100":    "November",
	"-0200":    "Oscar",
	"-0300":    "Papa",
	"-0400":    "Quebec",
	"-0430":    "Quebec",
	"-0500":    "Romeo",
	"-0600":    "Sierra",
	"-0700":    "Tango",
	"-0800":    "Uniform",
	"-0900":    "Victor",
	"-1000":    "Whiskey",
	"-1100":    "Xray",
	"-1200":    "Yankee",
	"+0000":    "Zulu" }

var lockMilitaryTzToUTCMap sync.Mutex

var militaryTzToUTCMap = map[string]string{
	"Alpha":    "+0100",
	"Bravo":    "+0200",
	"Charlie":  "+0300",
	"Delta":    "+0400",
	"Echo":     "+0500",
	"Foxtrot":  "+0600",
	"Golf":     "+0700",
	"Hotel":    "+0800",
	"India":    "+0900",
	"Kilo":     "+1000",
	"Lima":     "+1100",
	"Mike":     "+1200",
	"November": "-0100",
	"Oscar":    "-0200",
	"Papa":     "-0300",
	"Quebec":   "-0400",
	"Romeo":    "-0500",
	"Sierra":   "-0600",
	"Tango":    "-0700",
	"Uniform":  "-0800",
	"Victor":   "-0900",
	"Whiskey":  "-1000",
	"Xray":     "-1100",
	"Yankee":   "-1200",
	"Zulu":     "+0000"}

var lockMilitaryTzToLocationMap sync.Mutex

var militaryTzToLocationMap = map[string]string{
	"Alpha"    :  "France",
	"Bravo"    :  "Athens, Greece",
	"Charlie"  :  "Arab Standard Time, Iraq, Bahrain, Kuwait, Saudi Arabia, Yemen, Qatar",
	"Delta"    :  "Moscow, Russia and Afghanistan, however, Afghanistan is technically +4:30 from UTC",
	"Echo"     :  "Pakistan, Kazakhstan, Tajikistan, Uzbekistan and Turkmenistan",
	"Foxtrot"  :  "Bangladesh",
	"Golf"     :  "Thailand",
	"Hotel"    :  "Beijing, China",
	"India"    :  "Tokyo, Australia",
	"Kilo"     :  "Brisbane, Australia",
	"Lima"     :  "Sydney, Australia",
	"Mike"     :  "Wellington, New Zealand",
	"November" :  "Azores",
	"Oscar"    :  "Godthab, Greenland",
	"Papa"     :  "Buenos Aires, Argentina",
	"Quebec"   :  "Halifax, Nova Scotia",
	"Romeo"    :  "EST, New York, NY",
	"Sierra"   :  "CST, Dallas, TX",
	"Tango"    :  "MST, Denver, CO",
	"Uniform"  :  "PST, Los Angeles, CA",
	"Victor"   :  "Juneau, AK",
	"Whiskey"  :  "Honolulu, HI",
	"Xray"     :  "American Samoa",
	"Yankee"   :  "e.g. Fiji",
	"Zulu"     :  "Zulu time",
}

/*
Time Zone Abbreviations – Military Time Zone Names
https://www.timeanddate.com/time/zones/military

Military time zones are commonly used in aviation as well as at sea.
They are also known as nautical or maritime time zones. J (Juliet Time Zone)
is occasionally used to refer to the observer's local time.

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




var IanaUsMilitaryTimeZone = [25] string{

  "Europe/Paris",                    // UTC+1:  A (France)
  "Europe/Athens",                   // UTC+2:  B (Athens, Greece)
  "Asia/Riyadh",                     // UTC+3:  C (Arab Standard Time, Iraq, Bahrain,
                                     //           Kuwait, Saudi Arabia, Yemen, Qatar)
  "Asia/Dubai",                      // UTC+4:  D (Used for Moscow, Russia and Afghanistan,
                                     //           however, Afghanistan is technically +4:30 from UTC)
  "Asia/Karachi",                    // UTC+5:  E (Pakistan, Kazakhstan, Tajikistan,
                                     //           Uzbekistan and Turkmenistan)
  "Asia/Dhaka",                      // UTC+6:  F (Bangladesh)
  "Asia/Bangkok",                    // UTC+7:  G (Thailand)
  "Asia/Shanghai",                   // UTC+8:  H (Beijing, China)
  "Asia/Tokyo",                      // UTC+9:  I (Tokyo, Australia)
  "Australia/Brisbane",              // UTC+10: K (Brisbane, Australia)
  "Australia/Sydney",                // UTC+11: L (Sydney, Australia)
  "Pacific/Auckland" }               // UTC+12: M (Wellington, New Zealand)
  "Atlantic/Azores",                 // UTC-1:  N (Azores)
  "America/Godthab",                 // UTC-2:  O (Godthab, Greenland)
  "America/Argentina/Buenos_Aires",  // UTC-3:  P (Buenos Aires, Argentina)
  "America/Halifax",                 // UTC-4:  Q (Halifax, Nova Scotia
  "America/New_York",                // UTC-5:  R (EST, New York, NY)
  "America/Chicago",                 // UTC-6:  S (CST, Dallas, TX)
  "America/Denver",                  // UTC-7:  T (MST, Denver, CO)
  "America/Los_Angeles",             // UTC-8:  U (PST, Los Angeles, CA)
  "America/Juneau",                  // UTC-9:  V (Juneau, AK)
  "Pacific/Honolulu",                // UTC-10: W (Honolulu, HI)
  "Pacific/Pago_Pago",               // UTC-11: X (American Samoa)
  "Etc/GMT+12",                      // UTC -12: Y
  "Etc/UCT",                         // UTC+-0: Z (Zulu time)

*/
