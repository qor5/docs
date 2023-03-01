package docsrc

import "embed"

//go:embed assets/**.*
var Assets embed.FS

//go:embed assets/favicon.ico
var favicon []byte
