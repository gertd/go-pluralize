package main

import (
	"strings"
)

// TestCmd -- enum
type TestCmd uint8

// TestCmd -- enum constants
const (
	TestCmdUnknown TestCmd = 1 << iota
	TestCmdIsPlural
	TestCmdIsSingular
	TestCmdPlural
	TestCmdSingular
	TestCmdAll = TestCmdIsPlural + TestCmdIsSingular + TestCmdPlural + TestCmdSingular
)

// TestCmd -- string constants
const (
	testCmdUnknown    = "Unknown"
	testCmdIsPlural   = "IsPlural"
	testCmdIsSingular = "IsSingular"
	testCmdPlural     = "Plural"
	testCmdSingular   = "Singular"
	testCmdAll        = "All"
)

// testCmdID -- map enum to string
var testCmdID = map[TestCmd]string{
	TestCmdUnknown:    testCmdUnknown,
	TestCmdIsPlural:   testCmdIsPlural,
	TestCmdIsSingular: testCmdIsSingular,
	TestCmdPlural:     testCmdPlural,
	TestCmdSingular:   testCmdSingular,
	TestCmdAll:        testCmdAll,
}

// testCmdName -- map string to enum
var testCmdName = map[string]TestCmd{
	strings.ToLower(testCmdUnknown):    TestCmdUnknown,
	strings.ToLower(testCmdIsPlural):   TestCmdIsPlural,
	strings.ToLower(testCmdIsSingular): TestCmdIsSingular,
	strings.ToLower(testCmdPlural):     TestCmdPlural,
	strings.ToLower(testCmdSingular):   TestCmdSingular,
	strings.ToLower(testCmdAll):        TestCmdAll,
}

// String -- stringify TestCmd
func (t TestCmd) String() string {
	return testCmdID[t]
}

// Set -- set flag
func (t *TestCmd) Set(flag TestCmd) {
	*t = *t | flag
}

// Clear -- clear flag
func (t *TestCmd) Clear(flag TestCmd) {
	*t = *t &^ flag
}

// Toggle -- toggle flag state
func (t *TestCmd) Toggle(flag TestCmd) {
	*t = *t ^ flag
}

// Has -- is flag set?
func (t TestCmd) Has(flag TestCmd) bool {
	return t&flag != 0
}

// TestCmdString -- convert string reprensentation in to enum value
func TestCmdString(s string) TestCmd {
	if value, ok := testCmdName[strings.ToLower(s)]; ok {
		return value
	}
	return TestCmdUnknown
}
