# Running Tests

## Run Basic Tests
Open a command prompt in this directory (MikeAustin71\datetimeopsgo\zztests)
and run the following command:

### `go test -v > xxtests.txt`

This will generate a '.txt' in the current directory which contains all 
test results.

## Resources
 ##### `http://codesamplez.com/development/golang-unit-testing`

## Dependencies

### testify package
These tests make use of the 'testify' package. The package
can be installed from:

   `https://github.com/stretchr/testify`

## Running Tests with code coverage

First pull down and install the `cover` package.
 
  `go get golang.org/x/tools/cmd/cover`
  
Next, follow the test execution protocol.  
  
### Test Execution with Code Coverage

  `go test -cover -v > xx_tests.txt`  
     

### Cover Profile

Generate the code coverage detail:

  `go test -coverprofile=xx_coverage.out`


The following provides for code coverage display in your
browser. Run this on the command line:

  `go tool cover -html=xx_coverage.out`