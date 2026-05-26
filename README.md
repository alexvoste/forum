# forgezero.pub

A terminal-oriented discussion space built for systems engineers and low-level developers. No JavaScript. No tracking. Pure HTML over a minimalist Go backend.

## Philosophy
Infrastructure follows the Plan 9 principle of simplified resource access. Access is identity-based, prioritizing cryptographic keys over traditional credentials. You own the hardware; you own the access.

## Architecture
- Backend: Go (Standard library + net/http)
- Database: SQLite (via WASM-based CGO-less driver)
- Interface: Server-Side Rendered (SSR) HTML
- Authentication: Ed25519 SSH Public Keys (preferred) / Password Fallback
- Protocol: HTTP/1.1 (designed for 9p-style mounting in future iterations)

## Development
Project structure follows the internal/ package pattern to enforce strict encapsulation of logic.

- /internal/render: Template management and SSR logic
- /internal/routes: Request multiplexing and API definitions
- /internal/db: Database connectivity and schema management
- /public: Static assets and HTML templates

## Build
Requires Go 1.22+.

$ go mod download
$ go build -o fz-forum main.go
$ ./fz-forum

## License
MIT. All contributions must adhere to the minimalist design constraints.

