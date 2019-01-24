package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	// note package order matters here, the inits are called in sequence
	// so if you need models before db this is where that happens
	"github.com/DIS-SIN/gobasicbits/bbmodels"
	"github.com/DIS-SIN/gobasicbits/bbdatabase"
)

// note you must define the return value otherwise it acts like you just cant get it
func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = "8080"	
	return port	
}

// also lower case function first letter means "private"/unexported, if you 
// use a captial letter here, magic dictates you will export the func and 
// it can be used. Capitalization matters.
func getServerEnvironment() string {
	env := os.Getenv("ENVIRONMENT")
	if "" == env {
	  env = "development"
	}	
	return env
}

// were using dotenv to get the .env files
func InitEnvironment() {
	env := getServerEnvironment()
	godotenv.Load(".env." + env + ".local")

	if os.Getenv("ENVIRONMENT") == "production" {
		fmt.Println("Init "+os.Getenv("ENVIRONMENT")+" Webserver")
	}
	if os.Getenv("ENVIRONMENT") == "development" {
		fmt.Println("Init "+os.Getenv("ENVIRONMENT")+" Webserver")
	}
}

func main() {
	// grab our env vars
	InitEnvironment()

	// init of the packages (models and db) iu implicitly done, 
	// Go fires up the inits for packages when you import them
	// this is only done once regardles of import. So you can use this for any of the one
	// time setups

	// do dummy functions to show use
	// note the capitalization
	bbmodels.NoOp()
	bbdatabase.NoOp()
	// grab our templates from the dir and assign to pointer
	InitalizeGuiTemplates()
	// fire up the router and start serving!
	CreateRouter( getServerPort() ) 	
}

