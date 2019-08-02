package route

import (
    "github.com/gorilla/mux"
    analytics "github.com/heaptracetechnology/google-analytics/analytics"
    "log"
    "net/http"
)

//Route struct
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

//Routes array
type Routes []Route

var routes = Routes{
    Route{
        "AccountList",
        "POST",
        "/accountList",
        analytics.AccountList,
    },
    Route{
        "WebPropertiesList",
        "POST",
        "/webPropertiesList",
        analytics.WebPropertiesList,
    },
    Route{
        "ProfilesList",
        "POST",
        "/profileList",
        analytics.ProfilesList,
    },
    Route{
        "RealtimeData",
        "POST",
        "/realtime",
        analytics.RealtimeData,
    },
}

//NewRouter route
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
