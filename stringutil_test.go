package stringutil

import (
	"fmt"
	"testing"
	. "github.com/bborbe/assert"
)

type TestStringAfterData struct {
	Content string
	Find    string
	Result  string
}

var TestStringAfterDatas = []TestStringAfterData{
	{
		Content: "",
		Find:    "",
		Result:  "",
	},
	{
		Content: "test",
		Find:    "",
		Result:  "test",
	},
	{
		Content: "test",
		Find:    "test",
		Result:  "",
	},
	{
		Content: "bla:test",
		Find:    "bla:",
		Result:  "test",
	},
	{
		Content: "test",
		Find:    "bla:",
		Result:  "",
	},
}

func TestStringAfter(t *testing.T) {
	for i, testdata := range TestStringAfterDatas {
		result := StringAfter(testdata.Content, testdata.Find)
		err := AssertThat(result, Is(testdata.Result).Message(fmt.Sprintf("result wrong in testcase %d!", i+1)))
		if err != nil {
			t.Fatal(err)
		}
	}
}

type TestStringBeforeData struct {
	Content string
	Find    string
	Result  string
}

var TestStringBeforeDatas = []TestStringBeforeData{
	{
		Content: "",
		Find:    "",
		Result:  "",
	},
	{
		Content: "test",
		Find:    "",
		Result:  "test",
	},
	{
		Content: "test",
		Find:    "test",
		Result:  "",
	},
	{
		Content: "foo:bar",
		Find:    ":bar",
		Result:  "foo",
	},
	{
		Content: "test",
		Find:    "bla:",
		Result:  "",
	},
	{
		Content: "a b c",
		Find:    " ",
		Result:  "a",
	},
}

func TestStringBefore(t *testing.T) {
	for i, testdata := range TestStringBeforeDatas {
		result := StringBefore(testdata.Content, testdata.Find)
		err := AssertThat(result, Is(testdata.Result).Message(fmt.Sprintf("result wrong in testcase %d!", i+1)))
		if err != nil {
			t.Fatal(err)
		}
	}
}

type TestTrimData struct {
	Input    string
	Expected string
}

var TestTrimDatas = []TestTrimData{
	{
		Input:    "",
		Expected: "",
	},
	{
		Input:    " hello ",
		Expected: "hello",
	},
	{
		Input:    " hello world ",
		Expected: "hello world",
	},
	{
		Input:    "\rhello world\r",
		Expected: "hello world",
	},
	{
		Input:    "\nhello world\n",
		Expected: "hello world",
	},
	{
		Input:    "\r \nhello world\n \r",
		Expected: "hello world",
	},
}

func TestTrim(t *testing.T) {
	for i, testdata := range TestTrimDatas {
		result := Trim(testdata.Input)
		err := AssertThat(result, Is(testdata.Expected).Message(fmt.Sprintf("result wrong in testcase %d!", i+1)))
		if err != nil {
			t.Fatal(err)
		}
	}
}
