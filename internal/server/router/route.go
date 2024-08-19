package router

import (
	"github.com/gorilla/mux"
	"go-keycloak-example1/controller"
	"go-keycloak-example1/server/middleware"
	"net/http"
)

func New(ctrl *controller.Controller, ma *middleware.Authentication) *mux.Router {
	// create a root router
	router := mux.NewRouter().StrictSlash(true)

	//noAuthRouter := router.PathPrefix("/").Subrouter()

	// add a subrouter based on matcher func
	// note, routers are processed one by one in order, so that if one of the routing matches other won't be processed
	//noAuthRouter := router.PathPrefix("/").MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
	//	return r.Header.Get("Authorization") == ""
	//}).Subrouter()
	//
	//// add one more subrouter for the authenticated service methods
	//authRouter := router.PathPrefix("/secure").MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
	//	return r.Header.Get("Authorization") != ""
	//}).Subrouter()

	// map url routes to controller's methods
	authRoute := router.Methods(http.MethodPost, http.MethodGet).Subrouter()
	authRoute.HandleFunc("/test", ctrl.Login)
	authRoute.HandleFunc("/docs", ctrl.GetDocs)

	// use middleware
	authRoute.Use(ma.VerifyToken)

	noAuthRoute := router.PathPrefix("/login").Subrouter()
	noAuthRoute.HandleFunc("", ctrl.Login)

	return router
}
