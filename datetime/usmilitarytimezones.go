package datetime



// Time Zone Abbreviations – Military Time Zone Names
//	Reference:
//		https://www.timeanddate.com/time/zones/military
//		http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//
// Military time zones are commonly used in aviation as well as at sea.
// They are also known as nautical or maritime time zones. J (Juliet Time Zone)
// is occasionally used to refer to the observer's local time.
//
//	Abbreviation Time zone name     Other names    Offset
//
//	    A        Alpha Time Zone                   UTC +1
//	    B        Bravo Time Zone                   UTC +2
//	    C        Charlie Time Zone                 UTC +3
//	    D        Delta Time Zone                   UTC +4
//	    E        Echo Time Zone                    UTC +5
//	    F        Foxtrot Time Zone                 UTC +6
//	    G        Golf Time Zone                    UTC +7
//	    H        Hotel Time Zone                   UTC +8
//	    I        India Time Zone                   UTC +9
//	    K        Kilo Time Zone                    UTC +10
//	    L        Lima Time Zone                    UTC +11
//	    M        Mike Time Zone                    UTC +12
//	    N        November Time Zone                UTC -1
//	    O        Oscar Time Zone                   UTC -2
//	    P        Papa Time Zone                    UTC -3
//	    Q        Quebec Time Zone                  UTC -4
//	    R        Romeo Time Zone                   UTC -5
//	    S        Sierra Time Zone                  UTC -6
//	    T        Tango Time Zone                   UTC -7
//	    U        Uniform Time Zone                 UTC -8
//	    V        Victor Time Zone                  UTC -9
//	    W        Whiskey Time Zone                 UTC -10
//	    X        X-ray Time Zone                   UTC -11
//	    Y        Yankee Time Zone                  UTC -12
//	    Z        Zulu Time Zone                    UTC +0
//
// The methods associated with type 'UsMilitaryTimeZones' return the
// equivalent IANA time zones. At first this may seem confusing.
//
// For example, Military Time Zone 'L' or 'Lima' specifies UTC +11-hours.
// However, the equivalent IANA Time Zone is "Etc/GMT-11". In date time
// calculations, IANA Time Zone "Etc/GMT-11" computes as UTC +11-hours.
//	Reference:
//		https://en.wikipedia.org/wiki/Tz_database#Area
//
type UsMilitaryTimeZones	string

// Alpha - Military Time Zone 'A' equivalent to "Etc/GMT-1"
// Offset from UTC is computed at +1 hours
func (umtz UsMilitaryTimeZones) Alpha() string {return IanaTz.Etc.GMT_Minus_1()}

// Bravo - Military Time Zone 'B' equivalent to "Etc/GMT-2"
// Offset from UTC is computed at +2 hours
func (umtz UsMilitaryTimeZones) Bravo() string {return IanaTz.Etc.GMT_Minus_2()}

// Charlie - Military Time Zone 'C' equivalent to "Etc/GMT-3"
// Offset from UTC is computed at +3 hours
func (umtz UsMilitaryTimeZones) Charlie() string {return IanaTz.Etc.GMT_Minus_3()}

// Delta - Military Time Zone 'D' equivalent to "Etc/GMT-4"
// Offset from UTC is computed at +4 hours
func (umtz UsMilitaryTimeZones) Delta() string {return IanaTz.Etc.GMT_Minus_4()}

// Echo - Military Time Zone 'E' equivalent to "Etc/GMT-5"
// Offset from UTC is computed at +5 hours
func (umtz UsMilitaryTimeZones) Echo() string {return IanaTz.Etc.GMT_Minus_5()}

// Foxtrot - Military Time Zone 'F' equivalent to "Etc/GMT-6"
// Offset from UTC is computed at +6 hours
func (umtz UsMilitaryTimeZones) Foxtrot() string {return IanaTz.Etc.GMT_Minus_6()}

// Golf - Military Time Zone 'G' equivalent to "Etc/GMT-7"
// Offset from UTC is computed at +7 hours
func (umtz UsMilitaryTimeZones) Golf() string {return IanaTz.Etc.GMT_Minus_7()}

// Hotel - Military Time Zone 'H' equivalent to "Etc/GMT-8"
// Offset from UTC is computed at +8 hours
func (umtz UsMilitaryTimeZones) Hotel() string {return IanaTz.Etc.GMT_Minus_8()}

// India - Military Time Zone 'I' equivalent to "Etc/GMT-9"
// Offset from UTC is computed at +9 hours
func (umtz UsMilitaryTimeZones) India() string {return IanaTz.Etc.GMT_Minus_9()}

// Kilo - Military Time Zone 'K' equivalent to "Etc/GMT-10"
// Offset from UTC is computed at +10 hours
func (umtz UsMilitaryTimeZones) Kilo() string {return IanaTz.Etc.GMT_Minus_10()}

// Lima - Military Time Zone 'L' equivalent to "Etc/GMT-11"
// Offset from UTC is computed at +11 hours
func (umtz UsMilitaryTimeZones) Lima() string {return IanaTz.Etc.GMT_Minus_11()}

// Mike - Military Time Zone 'M' equivalent to "Etc/GMT-12"
// Offset from UTC is computed at +12 hours
func (umtz UsMilitaryTimeZones) Mike() string {return IanaTz.Etc.GMT_Minus_12()}

// November - Military Time Zone 'N' equivalent to "Etc/GMT+1"
// Offset from UTC is computed at -1 hours
func (umtz UsMilitaryTimeZones) November() string {return IanaTz.Etc.GMT_Plus_1()}

// Oscar - Military Time Zone 'O' equivalent to "Etc/GMT+2"
// Offset from UTC is computed at -2 hours
func (umtz UsMilitaryTimeZones) Oscar() string {return IanaTz.Etc.GMT_Plus_2()}

// Papa - Military Time Zone 'P' equivalent to "Etc/GMT+3"
// Offset from UTC is computed at -3 hours
func (umtz UsMilitaryTimeZones) Papa() string {return IanaTz.Etc.GMT_Plus_3()}

// Quebec - Military Time Zone 'Q' equivalent to "Etc/GMT+4"
// Offset from UTC is computed at -4 hours
func (umtz UsMilitaryTimeZones) Quebec() string {return IanaTz.Etc.GMT_Plus_4()}

// Romeo - Military Time Zone 'R' equivalent to "Etc/GMT+5"
// Offset from UTC is computed at -5 hours
func (umtz UsMilitaryTimeZones) Romeo() string {return IanaTz.Etc.GMT_Plus_5()}

// Sierra - Military Time Zone 'S' equivalent to "Etc/GMT+6"
// Offset from UTC is computed at -6 hours
func (umtz UsMilitaryTimeZones) Sierra() string {return IanaTz.Etc.GMT_Plus_6()}

