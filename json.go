// handling json
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("rESPONDING WITH 5xxx Error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"` // go does this to make sure json.marshal and other responses are .json files
	}
	// `{
	// 	"error: "something went wrong" this is what the json object would look like
	// }`

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload) // attempt to marshall any payload(object) its given into a JSON string and return it as bytes
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500) //internal server error
		return
	}
	w.Header().Add("Content-Type", "application/json") // add response header saying we are responding with json
	w.WriteHeader(code)
	w.Write(dat)
}
