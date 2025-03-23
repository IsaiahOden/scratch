package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) { // it always takes a response writer as first and pointer to http request as second
	respondWithJSON(w, 200, struct{}{}) // for right now we only care about sending 200 back, empty struct for right now in JSON
	// should respond if server is happy and running

}
