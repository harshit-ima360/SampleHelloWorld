package main

import (
	"fmt"
	"os"
	"testing"
)

type VariablesTest struct {
	queryQue  string
	resultRes string
}

func TestCheckEnvironmentVariable(t *testing.T) {
	//Few GoCD Variable Accessing
	fmt.Printf("\n\nCurrent Stats\nPipeline Name\t%v\n", os.Getenv("GO_PIPELINE_NAME"))
	fmt.Println("Stage Name-\t ", os.Getenv("GO_STAGE_NAME"))
	fmt.Println("Job Name-\t ", os.Getenv("GO_JOB_NAME"))

	//Sample Cases
	SampleTests := []VariablesTest{
		{"SampleString", "Hello, World!"},
		{"SampleString2", "Hi again Everyone"},
		//{"NotValid", ""},
	}

	//Iterating through all the cases
	for _, testIter := range SampleTests {
		if output := CheckEnvironmentVariable(testIter.queryQue); output != testIter.resultRes { //Checking for errors
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", testIter.queryQue, testIter.resultRes, output)
		}
	}
}
