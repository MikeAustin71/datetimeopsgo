# Enumerations

## Listing

1. LocationNameType
     datetime/locationnametypeenum.go

2. TDurCalcType
     datetime/timedurationcalctypeenum.go
     
3. TimeZoneCategory
     datetime/timezonecategoryenum.go

4. TimeZoneClass
     datetime/timezoneclassenum.go

5. TimeZoneConversionType
     datetime/timezoneconversiontypeenum.go

6. TimeZoneType
     datetime/timezonetypeenum.go

7. TimeZoneUtcOffsetStatus
     datetime/timezoneutcoffsetstatusenum.go

## Enumeration Constants

1. LocationNameType
    None                    0
    NonConvertibleTimeZone  1
    ConvertibleTimeZone     2

2. TDurCalcType
    None             :  0
    StdYearMt        :  1
    CumMonths        :  2
    CumWeeks         :  3
    CumDays          :  4
    CumHours         :  5
    CumMinutes       :  6
    CumSeconds       :  7
    CumMilliseconds  :  8
    CumMicroseconds  :  9
    CumNanoseconds   : 10
    GregorianYears   : 11

3. TimeZoneCategory
    None       :  0
    TextName   :  1
    UtcOffset  :  2

4. TimeZoneClass
    None              :  0
    AlternateTimeZone :  1
    OriginalTimeZone  :  2

5. TimeZoneConversionType
    None      :  0
    Absolute  :  1
    Relative  :  2

6. TimeZoneType
    None      :  0
    Iana      :  1
    Local     :  2
    Military  :  3
    
7. TimeZoneUtcOffsetStatus
    None      :  0
    Static    :  1
    Variable  :  2

## Global Variable Access Listing

1. LocationNameType
    var LocNameType LocationNameType

2. TDurCalcType
    var TDurCalc TDurCalcType

3. TimeZoneCategory
    var TzCat TimeZoneCategory

4. TimeZoneClass
    var TzClass TimeZoneClass

5. TimeZoneConversionType
    var TzConvertType TimeZoneConversionType

6. TimeZoneType
    var TzType TimeZoneType
    
7. TimeZoneUtcOffsetStatus
    var TzUtcStatus TimeZoneUtcOffsetStatus