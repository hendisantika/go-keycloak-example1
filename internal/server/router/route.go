package router

import (
	"github.com/gorilla/mux"
	"go-keycloak-example1/controller"
	"go-keycloak-example1/server/middleware"
	"net/http"
)

func New(ctrl *controller.Controller, ma *middleware.Authentication) *mux.Router {
	// create a root router
	router := mux.NewRouter()

	// add a subrouter based on matcher func
	// note, routers are processed one by one in order, so that if one of the routing matches other won't be processed
	noAuthRouter := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return r.Header.Get("Authorization") == ""
	}).Subrouter()

	// add one more subrouter for the authenticated service methods
	authRouter := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return true
	}).Subrouter()

	// map url routes to controller's methods
	noAuthRouter.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		ctrl.Login(writer, request)
	}).Methods("POST")

	authRouter.HandleFunc("/docs", func(writer http.ResponseWriter, request *http.Request) {
		ctrl.GetDocs(writer, request)
	}).Methods("GET")

	authRouter.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		ctrl.GetTest(writer, request)
	}).Methods("GET")

	// apply middleware
	authRouter.Use(ma.VerifyToken)

	return router
}
