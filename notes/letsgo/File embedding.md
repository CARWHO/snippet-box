- File embedding is used in web apps to bake content (e.g. static folder) into compiler. This is useful because we can ship one binary as static becomes part of .exe
	- No runtime file dependencies

Syntax:
- package ui
- import "embed"
- //go:embed "static" "html"
- var Files embed.FS

...

- fileServer := http.FileServer(ui.Files)
- mux.Handle("GET /static/", fileServer)