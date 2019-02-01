package router

import (
	"github.com/paulsouri/cryptian/handler"
	"github.com/gorilla/mux"
	"net/http"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route
// TODO: Add routes to create API endpoints
var routes = Routes{
	Route{
		"News",
		"GET",
		"/api/news",
		handler.NewsHandler,
	},
	Route{
		"Meetup",
		"GET",
		"/api/meetup",
		handler.MeetupHandler,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
