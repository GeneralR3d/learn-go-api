package middleware

import(
	"errors"
	"net/http"

	"learn-go-api/api"
	"learn-go-api/internal/tools"
	log "github.com/sirupsen/logrus"

)

//	create a custom unauthorized error
var UnAuthorizedError = errors.New("Invalid username or token!")

func Authorization(next http.Handler) http.Handler {

	//	return a http handler
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){

		//	takes in a response writer which constructs the response like setting response body or headers or status code
		// takes in a pointer to a request. The request contains all the info about the incoming request like headers, payload

		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error 

		// simply checking for existence, not checking if they are correct
		if username == "" || token == ""{
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)	// custom error handler we made, with custom error type we made
			return
		}

		//	if authorization is complete, proceed to get data from DATABASE

		//	create a new database using the interface type
		var database *tools.DatabaseInterface
		database, err = tools.NewDataBase()
		if err != nil{
			api.InternalErrorHandler(w)		
			return
		}

		//	query the database
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)	//return error if not matching

		if loginDetails == nil|| token != (*loginDetails).AuthToken{
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)	// custom error handler we made, with custom error type we made
			return
		}

		next.ServeHTTP(w,r)	//	calls the next middleware in line or calls handler function for the end point 
	})
}
