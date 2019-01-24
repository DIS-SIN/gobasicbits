package bbmodels

import (
    "fmt"
)

func init() {
	// we will use this later
	fmt.Println("Init: Basic Models Initialized!")
}
// dummy function to show usage of package compilation
func NoOp() {
    i := "no-op success"
    fmt.Println("bbmodels: " + i)        
}
