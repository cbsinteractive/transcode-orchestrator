package serialization

import (
    "encoding/json"
    "github.com/mitchellh/mapstructure"
    "reflect"
    "time"
)

const TimeLayoutDateTime = "2006-01-02T15:04:05Z"
const TimeLayoutDate = "2006-01-02"

func customDecode(sourceType reflect.Type, targetType reflect.Type, data interface{}) (interface{}, error) {
	if dataAsString, isString := data.(string); isString {
		if targetType.String() == "time.Time" {
			t, err := time.Parse(TimeLayoutDateTime, dataAsString)
			if err != nil {
				return time.Parse(TimeLayoutDate, dataAsString)
			}
			return t, nil
		}
		if targetType.String() == "*time.Time" {
			t, err := time.Parse(TimeLayoutDateTime, dataAsString)
			if err != nil {
				t, err = time.Parse(TimeLayoutDate, dataAsString)
				return &t, err
			}
			return &t, nil
		}
	}
	return data, nil
}

func getDecoder(result interface{}) (*mapstructure.Decoder, error) {
	return mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook:       customDecode,
		Result:           result,
		WeaklyTypedInput: false})
}

func Decode(input interface{}, output interface{}) error {
	decoder, err := getDecoder(output)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

func Serialize(model interface{}) []byte {
	if model == nil {
		return nil
	}

	value := reflect.ValueOf(model)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			t := value.Field(i).Type()
			if t.String() == "time.Time" {
				elem := value.Field(i)
				if elem.CanSet() {
					timeValue, _ := value.Field(i).Interface().(time.Time)
					timeValue = timeValue.UTC().Round(time.Second)
					elem.Set(reflect.ValueOf(timeValue))
				}
			}
			if t.String() == "*time.Time" {
				elem := value.Field(i).Elem()
				if elem.CanSet() {
					timeValue, _ := elem.Interface().(time.Time)
					timeValue = timeValue.UTC().Round(time.Second)
					elem.Set(reflect.ValueOf(timeValue))
				}
			}
		}
	}

	ser, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return ser
}
