package utils

import "encoding/json"

func MapToStruct[T any](src map[string]any, dest T) error {
	dataB, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataB, dest)
}
