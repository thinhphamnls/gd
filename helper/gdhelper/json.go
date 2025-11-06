package gdhelper

import (
	"encoding/json"
	"errors"
	"fmt"
)

func UnmarshalJSON(value []byte, dest interface{}) error {
	if err := json.Unmarshal(value, dest); err != nil {
		return errors.New(fmt.Sprintf("failed to unmarshal JSON into %T, %v", dest, err))
	}
	return nil
}

func MarshalJSON(dest interface{}) ([]byte, error) {
	jByte, err := json.Marshal(dest)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to marshal JSON into %T, %v", dest, err))
	}
	return jByte, nil
}
