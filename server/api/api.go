package api

import (
	"log"
	"net/http"
	"os"

	"github.com/emwalker/digraph/server/dgraph"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/json-iterator/go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Server struct {
	Engine    *echo.Echo
	Schema    *graphql.Schema
	Fragments chan string
}

type errorField struct {
	Message   string        `json:"message"`
	Locations []interface{} `json:"locations"`
}

type errorOutput struct {
	Data   *string      `json:"data"`
	Errors []errorField `json:"errors"`
}

func graphqlErrors(errors []string) []byte {
	errsOut := make([]errorField, 0)
	for _, err := range errors {
		errsOut = append(errsOut, errorField{
			Message:   err,
			Locations: []interface{}{},
		})
	}

	errorJson, err := json.Marshal(&errorOutput{
		Errors: errsOut,
	})

	if err != nil {
		log.Fatal(err)
	}
	return errorJson
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params dgraph.QueryParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	builder := dgraph.ParseQuery(r.Context(), s.Schema, params)
	if !builder.IsValid() {
		w.Write(graphqlErrors(*builder.Errors))
		return
	}

	response := dgraph.Do(r.Context(), builder.Stringify())
	str := dgraph.Serialize(response.Json)

	w.Write(str)
}

func (s *Server) ListenAndServe() error {
	s.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000", "http://localhost:8000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType},
	}))

	s.Engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${method} | ${status} | ${uri} -> ${latency_human}` + "\n",
		Output: os.Stdout,
	}))

	s.Engine.POST("/graphql", echo.WrapHandler(s))
	return s.Engine.Start(":5000")
}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
