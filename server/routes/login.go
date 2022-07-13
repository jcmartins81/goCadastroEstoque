package routes

import "net/http"

var loginRoute = Route{
	URI: "/login",
	Method: http.MethodPost,
	Function: func(w http.ResponseWriter, r *http.Request) {},
	RequiresAuthentication: false,
}
