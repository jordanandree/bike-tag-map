package main

import (
	"encoding/json"
	"log"

	"net/http"

	"bike-tag-map/internal/reddit"
)

func postsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	posts := reddit.GetPosts()
	json, err := json.Marshal(posts)

	if err != nil {
		log.Printf("Error encoding JSON: %v", err.Error())
	}

	w.Write(json)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/posts", postsHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
