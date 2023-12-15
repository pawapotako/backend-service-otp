package util

import (
	"encoding/json"
	"regexp"
)

func IsMobileNumber(input string) bool {
	// Define the regular expression for a mobile number
	// This regex assumes a simple format with optional '+' and digits
	regex := `^\+?\d+$`

	// Compile the regular expression
	re := regexp.MustCompile(regex)

	// Match the input against the regex
	return re.MatchString(input)
}

func CompressToJsonBytes(obj any) []byte {
	raw, _ := json.Marshal(obj)
	return raw
}

func UnCompressFromJsonBytes[T any](data []byte) (T, error) {
	var obj T
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return obj, err
	}
	return obj, nil
}
