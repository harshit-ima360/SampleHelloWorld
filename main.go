package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("SampleString"))
	fmt.Println(os.Getenv("SampleString2"))
}
