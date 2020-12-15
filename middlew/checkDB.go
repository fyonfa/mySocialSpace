package middlew

import (
	"github.com/fyonfa/mySocialSpace/bd"
	"net/http"
)
//CheckDB is the middleware that allows me to know the status of the data base
func CheckDB(next http.HandlerFunc) http.HandlerFunc{
	//return anonymous function, middleware
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() ==0{
			http.Error(w, "Connection lost with DB", 500)
			return
		}
		//if doesnt give error, we pass all the objects from w, r
		next.ServeHTTP(w,r)
	}
}
