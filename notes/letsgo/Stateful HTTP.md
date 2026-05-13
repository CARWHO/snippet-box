- Stateful HTTP (hyper text transfer protocol) refers to managing user sessions via a session manager. 
- A session manager allows the creation and storage of user sessions. The session manager chosen for this app is: "github.com/alexedwards/scs/v2". A session is tied to a cookie, a session persists across requests (survives page reloads), a session holds stuff like: user ID, cart content, flash msgs, etc. 
- Session managers, manage tokens. These tokens identify which session belongs to which browser. These should be renewed at privilege change.
- Session != context; Middleware reads session and places into context.
- State transitioning is done: Logged out -> Logged in.
- Session managers allow Get, Put, Remove, Exists, and Pop methods

Syntax:
- import "github.com/alexedwards/scs/v2"
- sessionManager := scs.New()
- sessionManager.Store = mysqlstore.New(db)
- sessionManager.Lifetime = 12 * time.Hour
- sessionManager.Cookie.Secure = true

...

func (app `*`application) authenticate(next http.Handler) http.Handler {
	id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	if id == 0 {
	...
	return
	}
	...
}
