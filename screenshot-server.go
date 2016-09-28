package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"encoding/json"
	"fmt"
	"html"
)

/* Route handlers */
func Index(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello, %q", html.EscapeString(req.URL.Path))
}

func TodoIndex(rw http.ResponseWriter, req *http.Request) {

	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	rw.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(rw).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoDetail(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todoId := vars["todoId"]
	fmt.Fprintf(rw, "Todo detail: "+todoId)
}

/* Routes */
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoDetail,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	/* Bind each defined route to the Gorilla router */
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

/* Todo model */
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoDetail)
	log.Fatal(http.ListenAndServe(":8080", router))
}
