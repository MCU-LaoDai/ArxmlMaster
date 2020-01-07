// AutosarConvert project AutosarConvert.go
package ArxmlMaster

import (
	"reflect"
)

// Arxml version enum
type ArxmlVersion int32

const (
	ArxmlVersionNone = iota
	ArxmlVersion321
	ArxmlVersion403
)

// create name index
func NameIndexCreate(node interface{}, nameParent string,
	nameIndex map[string]interface{}) {
	if node == nil {
		return
	}

	typ := reflect.TypeOf(node)
	value := reflect.ValueOf(node)

	if typ.Kind() == reflect.Ptr {
		if value.IsNil() {
			return
		}
		typ = typ.Elem()
		value = value.Elem()
	}

	switch typ.Kind() {
	case reflect.Slice:
		sliceNum := value.Len()
		for i := 0; i < sliceNum; i++ {
			NameIndexCreate(value.Index(i).Interface(), nameParent, nameIndex)
		}

	case reflect.Struct:
		fieldNum := typ.NumField()

		for i := 0; i < fieldNum; i++ {
			if typ.Field(i).Name == "ShortName" {
				nameNode := value.Field(i)
				shortName := nameNode.FieldByName("Text").String()
				nameParent += "/" + shortName
				nameIndex[nameParent] = nameNode.Interface()
			}
		}

		for i := 0; i < fieldNum; i++ {
			NameIndexCreate(value.Field(i).Interface(), nameParent, nameIndex)
		}

	default:

	}
}
