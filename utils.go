package alphavantage

import (
	"encoding/json"
	"errors"
)

func checkForError(data []byte) error {
	jsonRaw := map[string]json.RawMessage{}
	json.Unmarshal(data, &jsonRaw)

	if err := jsonRaw["Error Message"]; err != nil {
		return errors.New("Alpha Vantage API Error")
	}
	return nil
}
