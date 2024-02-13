package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
}

// InitHttpClient init internal url for upcoming request.
func InitHttpClient() {
	router.HandleFunc(SearchTrackByIDOrName, GetTrackerByIDOrNameHandler).Methods("GET")
	router.HandleFunc(SearchTrackByArtist, GetTrackerByArtistNameHandler).Methods("GET")
	router.HandleFunc(SearchSpotify, SearchSpotifyAPI).Methods("GET")
	http.Handle("/", router)

	fmt.Println("Server listening on :", 8080)
	if err := http.ListenAndServe(HostAndPort, nil); err != nil {
		fmt.Println(err)
	}
}

// SearchSpotifyAPI search track data from spotify apis and insert into DB.
func SearchSpotifyAPI(response http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("query")
	res := searchTracks(query)
	if req != nil && len(res.Tracks.Items) > 1 {
		insertDaoHandler(context.Background(), res.Tracks.Items)
	}
	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(response, "error", http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")

	if _, err := response.Write(responseJSON); err != nil {
		fmt.Println(err)
	}
}

// GetTrackerByIDOrNameHandler return tacks for a ID or given Name.
func GetTrackerByIDOrNameHandler(response http.ResponseWriter, req *http.Request) {

	searchRequest := SearchRequestByIDOrName{}
	err := json.NewDecoder(req.Body).Decode(&searchRequest)
	if err != nil {
		fmt.Println("error while parsing request", err)
		return
	}
	res := GetTrackByIDOrName(context.Background(), searchRequest.ID, searchRequest.Name)
	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(response, "error", http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	if _, err := response.Write(responseJSON); err != nil {
		fmt.Println(err)
	}
}

// GetTrackerByArtistNameHandler return tacks for or given Artist Name.
func GetTrackerByArtistNameHandler(response http.ResponseWriter, req *http.Request) {
	searchRequest := SearchRequestByArtistName{}
	err := json.NewDecoder(req.Body).Decode(&searchRequest)
	if err != nil {
		fmt.Println("error while parsing request", err)
		return
	}
	res := GetTracksByArtistName(context.Background(), searchRequest.Name)
	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(response, "error", http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	if _, err := response.Write(responseJSON); err != nil {
		fmt.Println(err)
	}
}
