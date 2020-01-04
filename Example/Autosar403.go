// Example project Autosar.go
package main

import (
	"fmt"

	"github.com/MCU-LaoDai/ArxmlMaster/Autosar403"
)

func main() {
	// parse arxml to mapping structure
	ar, err := Autosar403.NewAutosar("example.arxml")
	fmt.Println(err)
	// get packages in arxml
	fmt.Println(ar.Packages)
	// get name index dictionary
	fmt.Println(ar.NameIndex)
}
