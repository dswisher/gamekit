package sprites

import (
	"encoding/json"
	"fmt"
)

// detectFormat inspects JSON data and returns the detected format name.
// It looks at the structure of the "frames" field:
//   - If "frames" is an object (map), it's the Hash format
//   - If "frames" is an array, it's the Array format
//
// Returns "texture-packer-hash", "texture-packer-array", or an error if
// the format cannot be determined.
func detectFormat(data []byte) (string, error) {
	// Use a generic map to inspect the structure without full unmarshaling
	var doc map[string]any
	if err := json.Unmarshal(data, &doc); err != nil {
		return "", fmt.Errorf("failed to parse JSON for format detection: %w", err)
	}

	frames, ok := doc["frames"]
	if !ok {
		return "", fmt.Errorf("JSON does not contain 'frames' field")
	}

	switch frames.(type) {
	case map[string]any:
		return "texture-packer-hash", nil
	case []any:
		return "texture-packer-array", nil
	default:
		return "", fmt.Errorf("unable to detect format: 'frames' is neither an object nor an array")
	}
}
