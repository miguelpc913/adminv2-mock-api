package helpers

import "reflect"

func CopyStructFields(src interface{}, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst).Elem()

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Type().Field(i)
		srcFieldVal := srcVal.Field(i)

		if dstFieldVal := dstVal.FieldByName(field.Name); dstFieldVal.IsValid() {
			if dstFieldVal.CanSet() {
				dstFieldVal.Set(srcFieldVal)
			}
		}
	}

	return nil
}
