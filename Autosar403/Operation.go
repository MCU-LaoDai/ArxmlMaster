// Autosar403 project Operation.go
package Autosar403

import (
	"encoding/xml"
	"io/ioutil"
)

type Autosar AUTOSAR

func NewAutosar(arxmlFile string) (Autosar, error) {
	var ar Autosar

	bytes, err := ioutil.ReadFile(arxmlFile)
	if err != nil {
		return ar, err
	}

	err = xml.Unmarshal(bytes, &ar)
	if err != nil {
		panic(err)
	}

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

func (ar *Autosar) FindPackages() []AR_PACKAGE {
	var packages []AR_PACKAGE

	if ar.ArPackages != nil {
		packages = findPackages(packages, ar.ArPackages.ArPackage)
	}
	return packages
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
