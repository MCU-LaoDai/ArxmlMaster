// Example project Autosar.go
package main

import (
	"fmt"

	"github.com/MCU-LaoDai/ArxmlMaster/Autosar403"
	"github.com/MCU-LaoDai/ArxmlMaster/Nest"
)

type Example struct {
	Packages []Package `nest:"Packages"`
}

type Package struct {
	Name string `nest:"ShortName.Text"`
}

func main() {
	// parse arxml to mapping structure
	ar, err := Autosar403.NewAutosar("example.arxml")
	fmt.Println(err)
	// show packages in arxml
	fmt.Println(ar.Packages)
	// show name index dictionary
	fmt.Println(ar.NameIndex)

	//nest is a way to extract information from mapped structure using nest tag
	var example Example
	err = Nest.Unmarshal(ar, &example)
	fmt.Println(err)
	fmt.Println(example)
}
