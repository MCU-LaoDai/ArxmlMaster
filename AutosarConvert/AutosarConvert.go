// AutosarConvert project AutosarConvert.go
package ArxmlMaster

import (
	"Autosar321"
	"Autosar403"
	"encoding/xml"
	"io/ioutil"
)

// Arxml version enum
type ArxmlVersion int32

const (
	ArxmlVersion321 = iota
	ArxmlVersion403
	ArxmlVersionNumber
)

// seal different version of arxml
type Autosar struct {
	version ArxmlVersion
	ar321   Autosar321.AUTOSAR
	ar403   Autosar403.AUTOSAR
}

func NewAutosar(arxmlFile string, version ArxmlVersion) Autosar {
	// convert file to byte array
	bytes, err := ioutil.ReadFile(arxmlFile)
	if err != nil {
		panic(err)
	}

	// use different mapping for different version of autosar
	var ar Autosar
	ar.version = version
	switch ar.version {
	case ArxmlVersion321:
		var autosar Autosar321.AUTOSAR
		err = xml.Unmarshal(bytes, &autosar)
		if err != nil {
			panic(err)
		}
		ar.ar321 = autosar

	case ArxmlVersion403:
		var autosar Autosar403.AUTOSAR
		err = xml.Unmarshal(bytes, &autosar)
		if err != nil {
			panic(err)
		}
		ar.ar403 = autosar
	}

	return ar
}

func (ar *Autosar) Pack(outPath string) {
	// use different mapping for different version of autosar
	var bytes []byte
	var err error
	switch ar.version {
	case ArxmlVersion321:
		bytes, err = xml.MarshalIndent(ar.ar321, "", "  ")

	case ArxmlVersion403:
		bytes, err = xml.MarshalIndent(ar.ar403, "", "  ")
	}

	if err != nil {
		panic(err)
	}

	// save bytes to file
	err = ioutil.WriteFile(outPath, bytes, 0)

	if err != nil {
		panic(err)
	}
}
