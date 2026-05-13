- Used to run code that touches many parts of the codebase on every request so that doesn't pollute handler code. 
- Middleware can be placed into chains to string together a number of operations (like caching, logging, etc). 
- The package 'alice' ("github.com/justinas/alice") is used to create easily readable middleware chains. .Then/ .ThenFunc() methods used to chain middleware.

Syntax:
- import "github.com/justinas/alice"
- func main() {
	- // stateful routes needing session w/o auth
	- dynamic := alice.New(middleware1, middleware2, ... middlewareN) 

- mux.handle(...)
	- // (dynamic + protected) stateful routes needing session w/ auth
	- protected := alice.New(p_middleware1, p_middleware2, ... p_middlewareN)
- }

...


func middleware1 (next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r http.Request) {
	// do something on request
	next.ServeHTTP(w,r)
	// do something on response
	})
}