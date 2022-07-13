package routes

import (
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               func(r http.ResponseWriter, w *http.Request) {},
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               func(r http.ResponseWriter, w *http.Request) {},
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               func(r http.ResponseWriter, w *http.Request) {},
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               func(r http.ResponseWriter, w *http.Request) {},
		RequiresAuthentication: false,
	},
}
