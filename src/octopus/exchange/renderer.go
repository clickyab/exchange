package exchange

import (
	"io"
)

// Renderer is a way to render the ad into system
type Renderer interface {
	// Render is a function to handle rendering of bunch of ads into a
	// output stream
	Render(Impression, map[string]Advertise, io.Writer) error
}
