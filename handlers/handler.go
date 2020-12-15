package handlers

import (
	"github.com/fyonfa/mySocialSpace/middlew"
	"github.com/fyonfa/mySocialSpace/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors" //permits thats allow the API to be accessible from anywhere, local not much problem but remote, deployed
	"log"
	"net/http"
	"os"
)

//Handling sets our port and server starts to listen
func Handling() {
	router := mux.NewRouter() //captures the HTTP, response and request, check if there is information in header...

	//first check that there is an endpoint created for register, it has to match that it arrives POST type
	//detected that POST arrives with /register, middleware takes action
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("Post")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("Post")
	router.HandleFunc("/viewprofile", middlew.CheckDB(middlew.JWTValidate(routers.ViewProfile))).Methods("Get")
	router.HandleFunc("/modifyProfile", middlew.CheckDB(middlew.JWTValidate(routers.ModifyUserProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.JWTValidate(routers.RecordTweet))).Methods("POST")
	router.HandleFunc("/readTweet", middlew.CheckDB(middlew.JWTValidate(routers.ReadTweets))).Methods("GET")

	//open the port, see if in the OS the PORT is already created
	PORT := os.Getenv("PORT")
	if PORT == "" { //if port is nothing I will create it
		PORT = "8080"
	}
	//cors takes over of the web request
	handler := cors.AllowAll().Handler(router)        //allow to anyone, permits, depending IP, etc...
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //:8080, sets wht port has to listen and we put our handler

}
