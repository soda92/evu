package diagrams

import (
	"html/template"
)

type SVGDiagram interface {
	// Wrapped returns the diagram as an SVG, including the <svg> container.
	Wrapped() template.HTML

	// Inner returns the inner markup of the SVG.
	// This allows for the <svg> container to be created manually.
	Inner() template.HTML

	// Width returns the width of the SVG.
	Width() int

	// Height returns the height of the SVG.
	Height() int
}
