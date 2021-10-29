package main

import (
	"fmt"
	"log"
	"os"
)

func CheckEnvironmentVariable(VariableStr string) string {
	if valueVar, err := os.LookupEnv(VariableStr); err {
		fmt.Println(valueVar, "is stored in ", VariableStr)
		return valueVar
	} else {
		log.Fatal(VariableStr, " Variable not found in Envronment")
		return ""

	}

}

func main() {
	CheckEnvironmentVariable("SampleString")
	CheckEnvironmentVariable("SampleString2")
	CheckEnvironmentVariable("NotValid") //
}
