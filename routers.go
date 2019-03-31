package main

import(
  "net/http"
  "github.com/gorilla/mux"
)

//Route struct parametros del route
type Route struct{
  Name        string
  Method      string
  Pattern     string
  HandleFunc  http.HandlerFunc
}

//Routes arreglo de un struact de route
type Routes []Route

//NewRouter funcion que crear un router y lo retorna para conexión a un server.
func NewRouter() *mux.Router{
  router := mux.NewRouter().StrictSlash(true)

  for _, route := range routes{
    router.Methods(route.Method).
    Path(route.Pattern).
    Name(route.Name).
    Handler(route.HandleFunc)
  }
  return router
}

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "ErrorShow",
    "GET",
    "/errorshow/{cpnid}",
    ErrorShow,
  },
  Route{
    "ErrorAdd",
    "POST",
    "/erroradd",
    ErrorAdd,
  },
  Route{
    "ErrorUpdate",
    "PUT",
    "/errorupdate/{cpnid}",
    ErrorUpdate,
  },
}
