The foundations of web apps are: Routes, handlers and a http web server. Routes are spaces where we can display content (URL pattern matched against incoming request path), handlers control what we serve to each page and the http web server is how we serve this content to users. 

More foundational knowledge on webapps:
- Wildcard routing can be used to enable dynamic routing (pass variables in to change routes) E.g. snippet/view/{ID} -> snippet/view/4
- Method based routing is restricting our application to only respond to http requests w/ appropriate http method (GET/PUT/POST/PATCH/HEAD/OPTIONS)
- Routers, middleware and handlers are all the same shape. All are single interface = whole http layer composes by passing handlers around.
- http.Handler is a go interface w/ one method (ServeHTTP(w, http.ResponseWriter r http.Request)). Anything that implements this method is a handler.
	- w: Where you write responses (header, status, body)
	- r: Where you read requests (method, URL, headers)
