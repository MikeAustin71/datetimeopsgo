# Date Time Format Utility

This utility is designed to convert date time strings into specific
date time values; that is time.Time structures. The methods provided
address the problem of converting date time strings entered by
users which may not follow a date time formatting standard. To
address this issue, the utility generates about 1.5-million possible
format patterns and applies these to date times submitted in unknown
formats. 

On my machine, the 1.5-million format maps require about two seconds
for generation. Thereafter, date time strings are usually parsed in
under 35-milliseconds.

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