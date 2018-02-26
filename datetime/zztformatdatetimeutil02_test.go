package datetime

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"time"
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

	assert.Equal(suite.T(), t1Str, actualStr,"Acutal Output NOT EQUAL To Expected!")

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


func TestDtfmtTestSuite(t *testing.T) {
	tests := new(dtfmtTestSuite)
	suite.Run(t, tests)
}