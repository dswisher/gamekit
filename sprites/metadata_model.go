package sprites

import "image"

// spriteMetadata represents a single sprite's data in a format-agnostic way.
type spriteMetadata struct {
	Name       string
	Rect       image.Rectangle // Position in the sprite sheet
	Rotated    bool            // Whether sprite is rotated
	Trimmed    bool            // Whether sprite was trimmed
	SourceSize image.Point     // Original size before trimming
}

// sheetMetadata is the internal representation of a sprite sheet.
type sheetMetadata struct {
	Image   string                    // Path to the image file
	Format  string                    // Format identifier (e.g., "RGBA8888")
	Size    image.Point               // Size of the sprite sheet
	Scale   string                    // Scale factor
	Sprites map[string]spriteMetadata // Name → sprite lookup (lowercase keys for case-insensitive lookup)
}

// common JSON structure types used across format parsers.
type tpRect struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

// toRectangle converts a tpRect to an image.Rectangle.
func (r tpRect) toRectangle() image.Rectangle {
	return image.Rect(r.X, r.Y, r.X+r.W, r.Y+r.H)
}

type tpSize struct {
	W int `json:"w"`
	H int `json:"h"`
}

// toPoint converts a tpSize to an image.Point.
func (s tpSize) toPoint() image.Point {
	return image.Pt(s.W, s.H)
}
