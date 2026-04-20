package sprites

import (
	"encoding/json"
	"fmt"
	"strings"
)

// tpHashSheet represents the JSON structure of the TexturePacker Hash format.
type tpHashSheet struct {
	Frames map[string]tpHashFrame `json:"frames"`
	Meta   tpHashMeta             `json:"meta"`
}

// tpHashFrame represents a single frame in the Hash format.
type tpHashFrame struct {
	Frame            tpRect `json:"frame"`
	Rotated          bool   `json:"rotated"`
	Trimmed          bool   `json:"trimmed"`
	SpriteSourceSize tpRect `json:"spriteSourceSize"`
	SourceSize       tpSize `json:"sourceSize"`
}

// tpHashMeta represents the metadata section in the Hash format.
type tpHashMeta struct {
	Image  string `json:"image"`
	Format string `json:"format"`
	Size   tpSize `json:"size"`
	Scale  string `json:"scale"`
}

// loadHashFormat parses TexturePacker Hash format JSON data and returns
// the internal sheetMetadata representation.
func loadHashFormat(data []byte) (*sheetMetadata, error) {
	var sheet tpHashSheet
	if err := json.Unmarshal(data, &sheet); err != nil {
		return nil, fmt.Errorf("failed to unmarshal hash format JSON: %w", err)
	}

	result := &sheetMetadata{
		Image:   sheet.Meta.Image,
		Format:  sheet.Meta.Format,
		Size:    sheet.Meta.Size.toPoint(),
		Scale:   sheet.Meta.Scale,
		Sprites: make(map[string]spriteMetadata, len(sheet.Frames)),
	}

	for name, frame := range sheet.Frames {
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
