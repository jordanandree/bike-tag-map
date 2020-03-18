package reddit

import (
	"io/ioutil"
	"log"

	"encoding/json"
	"net/http"

	"github.com/tidwall/gjson"
)

// Post is a reddit post
type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Permalink string `json:"permalink"`
	Selftext  string `json:"selftext,omitempty"`
}

type APIerror struct {
	Message string `json:"message"`
	Error   int    `json:"error,omitempty"`
}

var baseURL string = "https://www.reddit.com"
var useragent string = "bike-tag-map/1.0"

func (post *Post) isSelftext() bool {
	return post.Selftext != ""
}

func GetPosts() []Post {
	var posts []Post
	var APIerror APIerror

	client := &http.Client{}

	req, err := http.NewRequest("GET", baseURL+"/search.json?q=bike+tag+AND+subreddit:babike&limit=100&sort=new", nil)
	req.Header.Set("User-Agent", useragent)
	if err != nil {
		log.Printf("Error creating new request: %v", err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error fetching response: %v", err.Error())
	}

	jsonBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading JSON data from response: %v", err.Error())
	}

	json.Unmarshal(jsonBody, &APIerror)
	if APIerror.Error > 0 {
		log.Printf("API Error: [%d] %v", APIerror.Error, APIerror.Message)
	}

	rawPosts := gjson.Get(string(jsonBody), "data.children.#.data")
	err = json.Unmarshal([]byte(rawPosts.String()), &posts)
	if err != nil {
		log.Printf("Error parsing JSON response: %v", err.Error())
	}

	return posts
}
