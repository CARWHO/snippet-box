# Snippetbox

Go web app from Let's Go by Alex Edwards. Paste some text, give it an expiry, share the link.

This was a Go learning exercise before I went and built a more proper web app. Every line of code here was written by hand, nothing generated.

What's in it:

- Routing with the stdlib mux, method based and wildcard routes
- MySQL through `database/sql`, models own all the DB access
- html/template with a template cache, static files embedded in the binary
- Form parsing, decoding and validation
- Sessions with scs, signup and login with bcrypt
- CSRF protection (nosurf), TLS, secure headers, structured logging
- Middleware chains with alice
- Unit and integration tests, with mocks for the models

## Screenshots

Home, latest snippets

![home](docs/screenshots/home.png)

Create a snippet, with validation

![create](docs/screenshots/create-validation.png)

View a snippet straight after creating it

![view](docs/screenshots/view-flash.png)

Signup

![signup](docs/screenshots/signup.png)

## Running it

Needs MySQL with a `snippetbox` database and a `web` user (password `snippetbox`). Schema is in `internal/models/testdata/setup.sql`, plus a `sessions` table for scs. `tls/` is gitignored, so generate a self signed cert first:

```
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

Move `cert.pem` and `key.pem` into `tls/`, then:

```
go run ./cmd/web
```

Serves on https://localhost:4000. Flags are `-addr` and `-dsn`.

Tests:

```
go test ./...
```

## Notes

Chapter notes are in `notes/letsgo/`. I wrote these in my own words from my understanding of each section, then went back through them at the end to make sure I actually understood everything I'd built.
