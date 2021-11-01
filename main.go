package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func CheckEnvironmentVariable(VariableStr string) string {
	if valueVar, err := os.LookupEnv(VariableStr); err {
		fmt.Println(valueVar, "is stored in ", VariableStr)
		return valueVar
	} else {
		log.Println(VariableStr, " Variable not found in Envronment")
		return "Not Found"

	}

}

func main() {
	//Few GoCD Variable Accessing
	fmt.Printf("\n\nCurrent Stats\nPipeline Name\t%v\n", os.Getenv("GO_PIPELINE_NAME"))
	fmt.Println("Stage Name-\t ", os.Getenv("GO_STAGE_NAME"))
	fmt.Println("Job Name-\t ", os.Getenv("GO_JOB_NAME"))

	//Custome Var using go program
	os.Setenv("NotValid", "New Var Through Go")
	//Checking Existing ones
	CheckEnvironmentVariable("SampleString")
	CheckEnvironmentVariable("SampleString2")
	CheckEnvironmentVariable("NotValid")

	router := gin.Default()

	router.GET("/:stringVar", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"result": CheckEnvironmentVariable(c.Param("stringVar")),
		})

	})

	router.Run()

}
