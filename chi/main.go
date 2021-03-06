package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// This is the response struct that will be
// serialized and sent back
type StatusResponse struct {
	Status string `json:"status"`
	User   string `json:"user"`
}

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	// Add Content-Type header to indicate JSON response
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	body := StatusResponse{
		Status: "Hello world from chi!",
		User:   chi.URLParam(r, "user"),
	}

	serializedBody, _ := json.Marshal(body)
	_, _ = w.Write(serializedBody)
}

type RequestBody struct {
	Name string `json:"name"`
}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {
	// Read complete request body
	rawRequestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Transform into RequestBody struct
	requestBody := &RequestBody{}
	err = json.Unmarshal(rawRequestBody, requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	body := StatusResponse{
		Status: "Hello world from chi!",
		User:   requestBody.Name,
	}

	serializedBody, _ := json.Marshal(body)
	_, _ = w.Write(serializedBody)
}

func main() {
	r := chi.NewRouter()

	r.Get("/users/{user}", UserGetHandler)
	r.Post("/users", UserPostHandler)

	log.Println("Listening on :8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
