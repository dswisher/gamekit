package sprites

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetectFormat_HashFormat(t *testing.T) {
	jsonData := `{
		"frames": {
			"sprite1": {"frame": {"x": 0, "y": 0, "w": 32, "h": 32}}
		},
		"meta": {"image": "test.png"}
	}`

	format, err := detectFormat([]byte(jsonData))
	assert.NoError(t, err)
	assert.Equal(t, "texture-packer-hash", format)
}

func TestDetectFormat_ArrayFormat(t *testing.T) {
	jsonData := `{
		"frames": [
			{"filename": "sprite1.png", "frame": {"x": 0, "y": 0, "w": 32, "h": 32}}
		],
		"meta": {"image": "test.png"}
	}`

	format, err := detectFormat([]byte(jsonData))
	assert.NoError(t, err)
	assert.Equal(t, "texture-packer-array", format)
}

func TestDetectFormat_EmptyHash(t *testing.T) {
	jsonData := `{
		"frames": {},
		"meta": {"image": "test.png"}
	}`

	format, err := detectFormat([]byte(jsonData))
	assert.NoError(t, err)
	assert.Equal(t, "texture-packer-hash", format)
}

func TestDetectFormat_EmptyArray(t *testing.T) {
	jsonData := `{
		"frames": [],
		"meta": {"image": "test.png"}
	}`

	format, err := detectFormat([]byte(jsonData))
	assert.NoError(t, err)
	assert.Equal(t, "texture-packer-array", format)
}

func TestDetectFormat_InvalidJSON(t *testing.T) {
	jsonData := `not valid json`

	format, err := detectFormat([]byte(jsonData))
	assert.Error(t, err)
	assert.Empty(t, format)
	assert.Contains(t, err.Error(), "failed to parse JSON")
}

func TestDetectFormat_MissingFrames(t *testing.T) {
	jsonData := `{
		"meta": {"image": "test.png"}
	}`

	format, err := detectFormat([]byte(jsonData))
	assert.Error(t, err)
	assert.Empty(t, format)
	assert.Contains(t, err.Error(), "does not contain 'frames' field")
}

func TestDetectFormat_NonObjectFrames(t *testing.T) {
	// Test with a primitive value for frames
	jsonData := `{
		"frames": "unexpected string",
		"meta": {"image": "test.png"}
	}`

	format, err := detectFormat([]byte(jsonData))
	assert.Error(t, err)
	assert.Empty(t, format)
	assert.Contains(t, err.Error(), "unable to detect format")
}

func TestDetectFormat_RealHashFile(t *testing.T) {
	data, err := os.ReadFile("testdata/texture-packer-hash.json")
	require.NoError(t, err)

	format, err := detectFormat(data)
	assert.NoError(t, err)
	assert.Equal(t, "texture-packer-hash", format)
}

func TestDetectFormat_RealArrayFile(t *testing.T) {
	data, err := os.ReadFile("testdata/texture-packer-array.json")
	require.NoError(t, err)

	format, err := detectFormat(data)
	assert.NoError(t, err)
	assert.Equal(t, "texture-packer-array", format)
}
