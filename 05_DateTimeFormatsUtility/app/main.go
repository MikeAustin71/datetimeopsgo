package main

import (
	"MikeAustin71/datetimeopsgo/05_DateTimeFormatsUtility/common"
	"errors"
	"fmt"
	"time"
)

/*

import (
	"MikeAustin71/datetimeopsgo/05_DateTimeFormatsUtility/common"
	"fmt"
)

*/

func main() {

	startTime := time.Now()

	dtf := common.DateTimeFormatUtility{}

	dtf.GetAllDateTimeFormats()

	lFmts := len(dtf.Formats)

	if lFmts < 1 {
		panic(errors.New("GetAllDateTimeFormats Completed, but no formats were created!"))
	}

	fh := common.FileHelper{}

	outputFile := "./formats/fmtRun.txt"

	if fh.DoesFileExist(outputFile) {
		err := fh.DeleteDirFile(outputFile)

		if err != nil {
			panic(errors.New("Error from DeleteDirFile() Error:" + err.Error()))
		}
	}

	f, err := fh.CreateDirAndFile(outputFile)

	if err != nil {
		panic(errors.New("Error from CreateDirAndFile Error:" + err.Error()))
	}

	defer f.Close()

	for _, s := range dtf.Formats {
		fh.WriteFileStr(s+"\n", f)
	}

	endTime := time.Now()

	du := common.DurationUtility{}

	et, _ := du.GetElapsedTime(startTime, endTime)

	fmt.Println(fmt.Sprintf("File Write Complete. %v formats written to file: %v ", lFmts, outputFile))
	fmt.Println("Elapsed Run Time: ", et.NanosecStr)

}
