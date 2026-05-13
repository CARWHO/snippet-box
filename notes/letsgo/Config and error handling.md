- In the SnippetBox app, config is managed by flags (dsn, addr).
- sLog package is used to manage logging throughout application. We define logger as logger := slog.New(), then we can user via logger.Error(), logger.Info()
- Dependency containers hold all shared dependencies for the app. In main.go we create an application struct containing all shared dependencies for app, we wire it up. This allows us to use handlers as methods. 
- Routes isolated in routes.go. Method on application returns Http.handler. This keeps main short and let's middleware wrap whole mux later.

Dependency container syntax:
- type application struct {
	logger  `*`slog.Logger
	snippets `*`models.SnippetModel
	...
- }

...

app := &application {
	logger: logger,
	snippets: &models.SnippetModel{DB: db}
	...

// {DB: db} is a struct literal, creates SnippetModel in snippets case. Sets DB to our database var. Now SnippetModel methods hit out mySQL.
}

func (app `*`application) snippetCreate(...) {
	...
	app.logger.Error(...)
}

Flag syntax:
- flag := flag.String("flag name", "default val", "desc")
- // allocates string address in memory and returns ptr to it
- // therefore we need to access w/ ptr when writing to the flag
- flag.Parse()

Log syntax:
- import "log/slog"
- logger := slog.New(slow.NewTextHandler(os.Stdout, nil))
- logger.Info("text for log", "key",  val (variable))
- logger.Error(err.Error())

Routes syntax:
- func (app `*`application) routes() http.Handler {
	// create http request router (routes incoming traffic to correct handler)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.home),
	...
	return mux
- }