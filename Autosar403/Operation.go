// Autosar403 project Operation.go
package Autosar403

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/MCU-LaoDai/ArxmlMaster"
)

type Autosar struct {
	Ar        AUTOSAR
	NameIndex map[string]interface{}
	Packages  []AR_PACKAGE
}

func NewAutosar(arxmlFile string) (Autosar, error) {
	var ar Autosar

	bytes, err := ioutil.ReadFile(arxmlFile)
	if err != nil {
		return ar, err
	}

	err = xml.Unmarshal(bytes, &ar.Ar)
	if err != nil {
		panic(err)
	}

	ar.NameIndex = make(map[string]interface{})
	ArxmlMaster.NameIndexCreate(ar.Ar, "", ar.NameIndex)
	ar.FindPackages()

	return ar, err
}

func (ar *Autosar) Pack(outPath string) error {
	bytes, err := xml.MarshalIndent(ar, "", "  ")

	if err != nil {
		return err
	}

	// save bytes to file
	err = ioutil.WriteFile(outPath, bytes, 0)
	return err
}

func (ar *Autosar) FindPackages() {
	ar.Packages = findPackages(ar.Packages, ar.Ar.ArPackages.ArPackage)
}

func findPackages(target []AR_PACKAGE, source []AR_PACKAGE) []AR_PACKAGE {
	if source != nil {
		for _, p := range source {
			target = append(target, p)
			if p.ArPackages != nil {
				target = findPackages(target, p.ArPackages.ArPackage)
			}
		}
	}
	return target
}
