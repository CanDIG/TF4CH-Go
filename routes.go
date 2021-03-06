package main

import "net/http"

//Route object creates to keep track of routes for router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is an array of Route objects
type Routes []Route

var routes = Routes{
	//Home page
	///Authenticated
	Route{
		"Index",
		"GET",
		"/mcfly/main",
		indexHandler,
	},
	Route{
		"Upload",
		"POST",
		"/mcfly/upload",
		uploadHandler,
	},
	//List of JSON objects
	///Authenticated
	Route{
		"GenericListJSON",
		"GET",
		"/mcfly/patients",
		patientJSONHandler,
	},
	Route{
		"GenericListJSON",
		"GET",
		"/mcfly/samples",
		sampleJSONHandler,
	},
	Route{
		"GenericListJSON",
		"GET",
		"/mcfly/mutations",
		mutationJSONHandler,
	},
	Route{
		"GenericListJSON",
		"GET",
		"/mcfly/metapatients",
		patientmetaJSONHandler,
	},
	Route{
		"GenericListJSON",
		"GET",
		"/mcfly/metasamples",
		samplemetaJSONHandler,
	},
	Route{
		"GenericListJSON",
		"GET",
		"/mcfly/metamutations",
		mutationmetaJSONHandler,
	},
	Route{
		"GenericListJSON",
		"GET",
		"/mcfly/metastudy",
		studymetaJSONHandler,
	},
	Route{
		"GenericListJSON",
		"GET",
		"/mcfly/export",
		makeAllHandler,
	},
}
