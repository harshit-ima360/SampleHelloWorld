package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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

	//Custom Var using go program
	os.Setenv("NotValid", "New Var Through Go")

	//Checking Existing ones
	CheckEnvironmentVariable("SampleString")
	CheckEnvironmentVariable("SampleString2")
	CheckEnvironmentVariable("NotValid")

	//Default router initialization
	router := gin.Default()

	//Endpoint to access our environment variables by Passing a name.
	router.GET("/var/:stringVar", func(c *gin.Context) {
		//Returning result
		c.JSON(200, gin.H{
			"result": CheckEnvironmentVariable(c.Param("stringVar")),
		})

	})

	//Close the server when needed
	router.GET("/closeServer", func(c *gin.Context) {
		//deferring the stop of container
		defer exec.Command("docker", "stop", "helloworldenv")
		//logging and closing the gin server by simply interrupting the program execution.
		log.Fatal("Exiting the Server")
	})

	router.Run()

}
