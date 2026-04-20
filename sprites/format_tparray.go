package sprites

import (
	"encoding/json"
	"fmt"
	"strings"
)

// tpArraySheet represents the JSON structure of the TexturePacker Array format.
type tpArraySheet struct {
	Frames []tpArrayFrame `json:"frames"`
	Meta   tpArrayMeta    `json:"meta"`
}

// tpArrayFrame represents a single frame in the Array format.
type tpArrayFrame struct {
	Filename         string `json:"filename"`
	Frame            tpRect `json:"frame"`
	Rotated          bool   `json:"rotated"`
	Trimmed          bool   `json:"trimmed"`
	SpriteSourceSize tpRect `json:"spriteSourceSize"`
	SourceSize       tpSize `json:"sourceSize"`
}

// tpArrayMeta represents the metadata section in the Array format.
type tpArrayMeta struct {
	Image  string `json:"image"`
	Format string `json:"format"`
	Size   tpSize `json:"size"`
	Scale  string `json:"scale"`
}

// loadArrayFormat parses TexturePacker Array format JSON data and returns
// the internal sheetMetadata representation.
func loadArrayFormat(data []byte) (*sheetMetadata, error) {
	var sheet tpArraySheet
	if err := json.Unmarshal(data, &sheet); err != nil {
		return nil, fmt.Errorf("failed to unmarshal array format JSON: %w", err)
	}

	result := &sheetMetadata{
		Image:   sheet.Meta.Image,
		Format:  sheet.Meta.Format,
		Size:    sheet.Meta.Size.toPoint(),
		Scale:   sheet.Meta.Scale,
		Sprites: make(map[string]spriteMetadata, len(sheet.Frames)),
	}

	for _, frame := range sheet.Frames {
		name := frame.Filename
		sprite := spriteMetadata{
			Name:       name,
			Rect:       frame.Frame.toRectangle(),
			Rotated:    frame.Rotated,
			Trimmed:    frame.Trimmed,
			SourceSize: frame.SourceSize.toPoint(),
		}
		// Store with lowercase key for case-insensitive lookup
		result.Sprites[strings.ToLower(name)] = sprite
	}

	return result, nil
}
