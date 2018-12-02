package main

import (
	"context"
	"crypto/subtle"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/emwalker/digraph/loaders"
	"github.com/emwalker/digraph/models"
	"github.com/emwalker/digraph/resolvers"
	"github.com/go-webpack/webpack"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/volatiletech/sqlboiler/boil"
)

const defaultPort = "8080"

type server struct {
	basicAuthUsername string
	basicAuthPassword string
	connectionString  string
	db                *sql.DB
	devMode           bool
	logLevel          int
	port              string
	schema            graphql.ExecutableSchema
}

func newServer(port string, devMode bool, username, password string, logLevel int) *server {
	connectionString := os.Getenv("POSTGRES_CONNECTION")
	if connectionString == "" {
		panic("POSTGRES_CONNECTION not set")
	}

	db, err := sql.Open("postgres", connectionString)
	must(err)
	must(db.Ping())

	resolver := &resolvers.Resolver{DB: db}
	schema := models.NewExecutableSchema(models.Config{Resolvers: resolver})

	return &server{
		basicAuthPassword: password,
		basicAuthUsername: username,
		connectionString:  connectionString,
		db:                db,
		devMode:           devMode,
		logLevel:          logLevel,
		port:              port,
		schema:            schema,
	}
}

func main() {
	devMode := flag.Bool("dev", false, "Development mode")
	webpack.Plugin = "manifest"
	webpack.Init(*devMode)

	logLevel := flag.Int("log", 1, "Print debugging information to the console")

	flag.Parse()

	s := newServer(
		getPlaygroundPort(),
		*devMode,
		os.Getenv("BASIC_AUTH_USERNAME"),
		os.Getenv("BASIC_AUTH_PASSWORD"),
		*logLevel,
	)
	s.routes()
	s.run()
}

func must(err error) {
	if err != nil {
		log.Fatal("there was a problem: ", err)
	}
}

func (s *server) basicAuthRequired(r *http.Request) bool {
	if s.basicAuthUsername == "" && s.basicAuthPassword == "" {
		return false
	}

	user, pass, ok := r.BasicAuth()
	return !ok ||
		subtle.ConstantTimeCompare([]byte(user), []byte(s.basicAuthUsername)) != 1 ||
		subtle.ConstantTimeCompare([]byte(pass), []byte(s.basicAuthPassword)) != 1
}

// https://stackoverflow.com/a/39591234/61048
func (s *server) withBasicAuth(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.basicAuthRequired(r) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Digraph"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized.\n"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// https://github.com/vektah/gqlgen-tutorials/blob/master/dataloader/graph.go
func (s *server) withLoaders(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, loaders.TopicLoaderKey, loaders.NewTopicLoader(ctx, s.db))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

const homepageTemplate = `<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="Content-Language" content="en">
    <title>Digraph</title>
    {{ asset "main.css" }}
  </head>

  <body>
    <div id="root"></div>
    {{ asset "vendors.js" }}
    {{ asset "main.js" }}
  </body>
</html>`

func (s *server) handleRoot() http.Handler {
	funcMap := map[string]interface{}{"asset": webpack.AssetHelper}
	t := template.New("homepage").Funcs(funcMap)
	template.Must(t.Parse(homepageTemplate))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	})
}

func (s *server) handleGraphqlRequest() http.Handler {
	handler := cors.Default().Handler(handler.GraphQL(s.schema))
	handler = handlers.CompressHandler(handler)
	if s.logLevel > 0 {
		handler = handlers.CombinedLoggingHandler(os.Stdout, handler)
	}
	return s.withLoaders(handler)
}

func (s *server) handleGraphqlPlayground() http.Handler {
	return handler.Playground("GraphQL playground", "/graphql")
}

func (s *server) handleHealthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
}

func (s *server) handleStaticFiles() http.Handler {
	fs := http.FileServer(http.Dir("public/webpack"))
	return http.StripPrefix("/static", fs)
}

func (s *server) routes() {
	http.Handle("/static/", s.withBasicAuth(s.handleStaticFiles()))
	http.Handle("/graphql", s.withBasicAuth(s.handleGraphqlRequest()))
	http.Handle("/playground", s.withBasicAuth(s.handleGraphqlPlayground()))
	http.Handle("/_ah/health", s.handleHealthCheck())
	http.Handle("/", s.withBasicAuth(s.handleRoot()))
}

func (s *server) run() {
	log.Printf("Running server with log level %d", s.logLevel)
	if s.logLevel > 1 {
		boil.DebugMode = true
	}

	log.Printf("Connect to http://localhost:%s/playground for the GraphQL playground", s.port)
	log.Printf("Listening on port %s", s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}

func getPlaygroundPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}