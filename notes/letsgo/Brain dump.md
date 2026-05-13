The book 'Let's Go!' explains how to create a maintainable, secure and fast web app. It explains project layout, routes, session management, databases, user authentication, logging/ error management and testing.

**Project layout:**
- cmd/web/: main.go, handlers.go, routes. go, 
- /ui: .html template files
- /internal
	- /assert: general test helpers 
	- /models models are in their own package. (code that owns data + database access)
		- /mocks: mock versions of insert, get and latest methods for snippets
		- /testdata: sql scripts to run at start and end of db operation
		- errors.go, users.go, snippets.go  
	- /validator
		- validates common actions (maxchars, notblank, etc)
- /tls 
	- Used for storing files for https certificate (cert + key)

**Routes:**
Used to render certain content on page. Routes call handler methods. E.g. /snippet/view uses snippetView handler method which renders specified template. Handlers can be POST, GET, PUT, PATCH, DELETE.

**Session management:**
Used for keeping track of user state. (cookies locate what session user is working on). Sessions are not explicitly tied to auth, guest user can use 'add to cart' on site.

**Middleware:**
Used to run code that touches many parts of codebase on every request without polluting handler code. Similar to helper code but middleware wraps handlers and gets called automatically on matching request in a chain. Middleware chain is a chain of middleware operations wrapped around final handler. Request goes in, response comes out.

**Databases:**
Used to store content that should be kept for all time (e.g. what a user added to their cart). When accessing database, run predefined SQL statement to carry out action. 

**User authentication:**
Used to place gates on what users have access to. User password and email stored in database. Password is stored in form of hashed, encrypted password. Helper functions check if user is authenticated. 

**Logging and error handling:**
Used to clarify issues/ milestones in code operation. Use slog package for consistent formatting/ handling of all logs. Error handling is done using nil and error.Is() check then returning 1/0, nil/err depending on cases. Use errors.go for custom errors.

**Testing:**
Used to confirm code works as expected. Basic unit testing, mock e2e route tests, etc... Unit tests confirm a function returns correct content, done by with struct for necessary content + what wanted outcome and loop to run tests. Mock e2e is done by creating mock package w/ mock methods to mock db/ handler responses (creating duplicate app/ TLS server running middleware, full router, etc)

