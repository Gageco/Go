package main

import(
  "net/http"
  "github.com/gorilla/mux"
)

type Route struct {                   //a struct for the Routes that we can use
  Name          string
  Method        string
  Pattern       string
  HandlerFunc   http.HandlerFunc
}

type Routes []Route                   //a struct that is a collection of Route

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    router.Methods(route.Method).
          Path(route.Pattern).
          Name(route.Name).
          Handler(route.HandlerFunc)
  }

  return router
}

var routes = Routes{
  Route{
    "INDEX",
    "GET",
    "/",
    Index,
  },
  Route{
    "TimerCreate",
    "POST",
    "/timers",
    TimerCreate,
  },
  Route{
    "TimerIndex",
    "GET",
    "/timers",
    TimerIndex,
  },
  Route{
    "RandomText",
    "GET",
    "/randomText",
    RandomText,
  },
  Route{
    "StateChange",
    "POST",
    "/stateChange",
    StateChange,
  },
  Route{
    "StateChangeIndex",
    "GET",
    "/stateChange",
    StateChangeIndex,
  },
  Route{
    "AuthTestGetId",
    "POST",
    "/getID",
    AuthTest,
  },
  Route{
    "AuthTestTryId",
    "POST",
    "/useID",
    AuthTestTryId,
  },
}
