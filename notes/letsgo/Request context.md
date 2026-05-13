Request context holds per-request information, deadlines, cancel signals (ctrl + c), request scoped values. Context builder  (context.WithValue/...) available via importing context package. Context getter method (request.Context()) available via importing net/http package.

Syntax:
- import (
	"net/http"
	"context"
	)

- // Keys for context are necessary for identifying specific pieces of context  
- type ContextKey string
- const isAuthenticatedContextKey = contextKey("isAuthenticated")
  
- ...
- exists, err := app.users.Exists(id)
- if err != nil {
	- ...
- }
- if exists {
	ctx := context.WithValue(r.Context(), isAuthenticatedContextKey, true)
	r = r.WithContext(ctx)
- }

Then in app to check if authed...

- func (app `*`application) is Authenticated(r `*`http.Request) bool {
	isAuth, ok := r.Context().Value(isAuthenticatedContextKey).(bool)
	...
	return isAuth
- }