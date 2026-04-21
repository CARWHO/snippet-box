package ui

import (
	"embed"
)

// compiler should bake static directory into binary
// and store it in the variable Files
//
//go:embed "static" "html"
var Files embed.FS
