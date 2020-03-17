package main

import (
	"encoding/json"
	"log"

	"net/http"

	"bike-tag-map/internal/reddit"
)

func main() {
	posts := reddit.GetPosts()
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		json, err := json.Marshal(posts)

		if err != nil {
			log.Printf("Error encoding JSON: %v", err.Error())
		}

		w.Write(json)
	})

	http.ListenAndServe(":8080", nil)
}
