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
func testCmdID(t TestCmd) string {

	innerTestCmdID := map[TestCmd]string{

		TestCmdUnknown:    testCmdUnknown,
		TestCmdIsPlural:   testCmdIsPlural,
		TestCmdIsSingular: testCmdIsSingular,
		TestCmdPlural:     testCmdPlural,
		TestCmdSingular:   testCmdSingular,
		TestCmdAll:        testCmdAll,
	}

	return innerTestCmdID[t]
}

// testCmdName -- map string to enum
// func testCmdName() map[string]TestCmd {

// 	return map[string]TestCmd{
// 		strings.ToLower(testCmdUnknown):    TestCmdUnknown,
// 		strings.ToLower(testCmdIsPlural):   TestCmdIsPlural,
// 		strings.ToLower(testCmdIsSingular): TestCmdIsSingular,
// 		strings.ToLower(testCmdPlural):     TestCmdPlural,
// 		strings.ToLower(testCmdSingular):   TestCmdSingular,
// 		strings.ToLower(testCmdAll):        TestCmdAll,
// 	}
// }

// testCmdName -- map string to enum value
func testCmdName(s string) TestCmd {

	f := func() func(s string) TestCmd {

		innerTestCmdName := map[string]TestCmd{
			strings.ToLower(testCmdUnknown):    TestCmdUnknown,
			strings.ToLower(testCmdIsPlural):   TestCmdIsPlural,
			strings.ToLower(testCmdIsSingular): TestCmdIsSingular,
			strings.ToLower(testCmdPlural):     TestCmdPlural,
			strings.ToLower(testCmdSingular):   TestCmdSingular,
			strings.ToLower(testCmdAll):        TestCmdAll,
		}

		inner := func(s2 string) TestCmd {

			if value, ok := innerTestCmdName[strings.ToLower(s2)]; ok {
				return value
			}
			return TestCmdUnknown
		}
		return inner
	}

	return f()(s)
}

// String -- stringify TestCmd
func (t TestCmd) String() string {
	return testCmdID(t)
}

// Set -- set flag
func (t *TestCmd) Set(flag TestCmd) {
	*t |= flag
}

// Clear -- clear flag
func (t *TestCmd) Clear(flag TestCmd) {
	*t &^= flag
}

// Toggle -- toggle flag state
func (t *TestCmd) Toggle(flag TestCmd) {
	*t ^= flag
}

// Has -- is flag set?
func (t TestCmd) Has(flag TestCmd) bool {
	return t&flag != 0
}

// TestCmdString -- convert string reprensentation in to enum value
func TestCmdString(s string) TestCmd {
	return testCmdName(s)
}
