// Example project Autosar.go
package main

import (
	"fmt"

	"github.com/MCU-LaoDai/ArxmlMaster/Autosar403"
)

func main() {
	ar, err := Autosar403.NewAutosar("example.arxml")
	fmt.Println(err)
	packages := ar.FindPackages()
	fmt.Println(len(packages))
	for _, p := range packages {
		fmt.Println(p.ShortName.Text)
	}
}
