package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "something went wrong") // see line 11 json.go
}

// {
// 	"error": "something went wrong"
//   }
// add in to documentation that this is what they should see
// recommended to commit the vendor folder in Go
