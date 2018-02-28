package datetime

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"time"
	"fmt"
)


type dtfmtTestSuite struct {

	suite.Suite
	DtFmt 			FormatDateTimeUtility
}

func (suite *dtfmtTestSuite) SetupSuite() {
	suite.DtFmt = FormatDateTimeUtility{}
	suite.DtFmt.CreateAllFormatsInMemory()
}

func (suite *dtfmtTestSuite) TearDownSuite() {
	suite.DtFmt = FormatDateTimeUtility{}
}

func (suite *dtfmtTestSuite) SetupTest() {

}

func (suite *dtfmtTestSuite) TearDownTest() {

}

func (suite *dtfmtTestSuite) TestParseDateTimeString01() {
	// Testing Vladivostok type time zones
	t1Str := "2018-02-25 16:28:52.515539300 +1000 +1000"

	t1, err := suite.DtFmt.ParseDateTimeString(t1Str, "")

	assert.Nil(suite.T(),err,"Error:")
	//assert.Nil(suite.T(), err, fmt.Sprintf("Error returned by suite.DtFmt.ParseDateTimeString(t1Str, \"\"). Error='%v'", err.Error()))

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	actualStr := t1.Format(fmtstr)

	assert.Equal(suite.T(), t1Str, actualStr,"Acutal Output NOT EQUAL To Expected!")

}

func (suite *dtfmtTestSuite) TestParseDateTimeString02() {
	// Testing Vladivostok type time zones
	t1Str := "2018-02-25 16:28:52.515539300 +1000 +1000"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1Expected, err := time.Parse("2006-01-02 15:04:05.000000000 -0700 -0700", t1Str)

	assert.Nil(suite.T(), err,"Error from time.Parse(fmtstr, t1Str).")

	t1, err := suite.DtFmt.ParseDateTimeString(t1Str, "")

	assert.Nil(suite.T(),err,"Error:")

	actualStr := t1.Format(fmtstr)

	assert.Equal(suite.T(), t1Str, actualStr,"Actual Output NOT EQUAL To Expected!")

	assert.True(suite.T(), t1.Equal(t1Expected),"t1 NOT EQUAL to t1Expected")

}

func (suite *dtfmtTestSuite) TestParseDateTimeString03() {
	// Testing Vladivostok type time zones
	t1Str := "2018-02-25 16:28:52.515539300 +1000 +10"

	t1, err := suite.DtFmt.ParseDateTimeString(t1Str, "")

	assert.Nil(suite.T(),err,"Error from suite.DtFmt.ParseDateTimeString(t1Str, \"\")")

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	actualStr := t1.Format(fmtstr)

	expectedStr := "2018-02-25 16:28:52.515539300 +1000 +1000"

	assert.Equal(suite.T(), expectedStr, actualStr,"Acutal Output NOT EQUAL To Expected!")

	t1Expected, err := time.Parse("2006-01-02 15:04:05.000000000 -0700 -07", t1Str)

	assert.Nil(suite.T(), err, "Error from time.Parse(fmtstr, t1Str)")

	assert.True(suite.T(), t1.Equal(t1Expected),"t1 NOT EQUAL to t1Expected")
}

func (suite *dtfmtTestSuite) TestParseDateTimeString04() {
// Testing Iana Vladivostok Time Zone Conversions
fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"


// expectedStr := dtz.String() == "2018-02-25 16:28:52.515539300 +1000 +10"

expectedStr := "2018-02-25 16:28:52.515539300 +1000 +1000"

t1, err := suite.DtFmt.ParseDateTimeString(expectedStr, "")

assert.Nil(suite.T(), err, "Error from suite.DtFmt.ParseDateTimeString(expectedStr, \"\")")


dtz, err := DateTzDto{}.NewDateTimeElements(2018,02,25,16,28,52,515539300,TzIanaAsiaVladivostok, fmtstr)

assert.Nil(suite.T(), err, "Error from DateTzDto{}.NewDateTimeElements")

t2, err := TimeZoneDto{}.New(t1, TzIanaAsiaVladivostok, fmtstr)

assert.Nil(suite.T(), err, "Error from TimeZoneDto{}.New(t1, TzIanaAsiaVladivostok, fmtstr)")


s := fmt.Sprintf("t2.TimeOut.DateTime NOT EQUAL to dtz.DateTime. t2.TimeOut.DateTime='%v'  dtz.DateTime='%v'", t2.TimeOut.DateTime.Format(fmtstr), dtz.DateTime.Format(fmtstr))
assert.True(suite.T(), t2.TimeOut.DateTime.Equal(dtz.DateTime),s)


actualStr := t1.Format(fmtstr)

assert.Equal(suite.T(), expectedStr , actualStr,"Actual DateTime Output NOT EQUAL To Expected!")

}

func TestDtfmtTestSuite(t *testing.T) {
	tests := new(dtfmtTestSuite)
	suite.Run(t, tests)
}