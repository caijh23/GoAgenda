package service

import (
    //"net/http"

    "github.com/urfave/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    mx.HandleFunc("/v1/user/getkey",getKeyHandler(formatter)).Methods("GET")
    mx.HandleFunc("/v1/users",listAllUserHandler(formatter)).Methods("GET")
    mx.HandleFunc("/v1/users",createUserHandler(formatter)).Methods("POST")
    mx.HandleFunc("/v1/users/{id}",updateUserHandler(formatter)).Methods("PUT")
    mx.HandleFunc("/v1/users/{id}",deleteUserHandler(formatter)).Methods("DELETE")
}
