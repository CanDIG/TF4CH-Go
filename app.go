package main

import (
	"flag"
	"log"
	"net/http"
)

//hosting variables defined in flags
var localhost string
var localport string
var mongo string
var keycloakhost string
var keycloakport string
var server string
var keycloakserver string

var goTest bool // true if unit tests are running

//Arrays gives the arrays to print out
var Arrays *Records

func main() {
	flag.StringVar(&localport, "p", "3000", "Specify which port to use")
	flag.StringVar(&mongo, "d", "localhost", "Specify where the mongodb is located")
	flag.StringVar(&localhost, "host", "localhost", "Specify the name of the host")
	flag.Parse()

	Arrays = GetAllConfs()

	server = "http://" + localhost + ":" + localport
	//for docker
	DatabaseInit("mcfly", "mongodb://"+mongo+":27017/")
	//for localhost
	//DatabaseInit("mcfly", "mongodb://localhost:27017/")
	router := NewRouter()
	//Stats hosting on the constant port
	log.Fatal(http.ListenAndServe(":"+localport, router))
}
