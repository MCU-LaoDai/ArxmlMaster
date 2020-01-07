// Nest project Nest.go
package Nest

import (
	"errors"
	"reflect"
	"strings"
)

func Unmarshal(src interface{}, tgt interface{}) error {
	tgtVal := reflect.ValueOf(tgt)
	if tgtVal.Type().Kind() != reflect.Ptr {
		return errors.New("tgt should be a pointer")
	}

	tgtVal = tgtVal.Elem()
	output := unmarshal(reflect.ValueOf(src), tgtVal.Type())
	tgtVal.Set(output)

	return nil
}

func unmarshal(srcVal reflect.Value, tgtTyp reflect.Type) reflect.Value {
	output := reflect.New(tgtTyp).Elem()
	srcTyp := srcVal.Type()

	if srcTyp.Kind() == reflect.Slice {
		if tgtTyp.Kind() == reflect.Slice {
			sliceNum := srcVal.Len()
			for i := 0; i < sliceNum; i++ {
				outputNew := unmarshal(srcVal.Index(i), tgtTyp.Elem())
				output = reflect.Append(output, outputNew)
			}
		}
	} else if srcTyp.Kind() == reflect.Struct {
		tgtNum := tgtTyp.NumField()

		// iterate all target member
		for i := 0; i < tgtNum; i++ {
			tgtTypChild := tgtTyp.Field(i)
			outputChild := output.Field(i)
			// get member nest tag
			tag := tgtTypChild.Tag.Get("nest")
			tags := strings.Split(tag, ".")
			resultVal := srcVal

			for _, t := range tags {
				resultVal = getValueFromSingleTag(resultVal, string(t))
			}

			if tgtTypChild.Type.Kind() == reflect.String {
				if resultVal.IsValid() == true {
					outputChild.Set(resultVal)
				}
			} else {
				if resultVal.IsValid() == true {
					outputChild.Set(unmarshal(resultVal, tgtTypChild.Type))
				}
			}
		}
	} else {
		output.Set(srcVal)
	}
	return output
}

func getValueFromSingleTag(srcVal reflect.Value, tag string) reflect.Value {
	if srcVal.IsValid() == false {
		return srcVal
	}

	typ := srcVal.Type()
	val := srcVal

	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch typ.Kind() {
	case reflect.Struct:
		result := val.FieldByName(tag)
		return result

	case reflect.Slice:
		fieldChild := val.Index(0).FieldByName(tag)
		results := reflect.MakeSlice(reflect.SliceOf(fieldChild.Type()), 0, 0)
		for i := 0; i < val.Len(); i++ {
			result := val.Index(i).FieldByName(tag)
			results = reflect.Append(results, result)
		}
		return results

	default:
		return val.FieldByName(tag)
	}
}
