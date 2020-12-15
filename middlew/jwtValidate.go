package middlew

import (
	"github.com/fyonfa/mySocialSpace/routers"
	"net/http"
)

//JWTValidate allows to validate the JWT that comes from the request
func JWTValidate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.TokenProcess(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error in Token! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
