package main

import "testing"

type VariablesTest struct {
	queryQue  string
	resultRes string
}

func TestCheckEnvironmentVariable(t *testing.T) {

	SampleTests := []VariablesTest{
		{"SampleString", "Hello, World!"},
		{"SampleString2", "Hi again Everyone"},
		//{"NotValid", ""},
	}

	for _, testIter := range SampleTests {
		if output := CheckEnvironmentVariable(testIter.queryQue); output != testIter.resultRes {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", testIter.queryQue, testIter.resultRes, output)
		}
	}
}
