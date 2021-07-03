package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/michael-grace/onlyfrogs/controllers"
	m "github.com/michael-grace/onlyfrogs/models"
)

func main() {

	noDBFlag := flag.Bool("no-db", false, "Whether to use internal memory instead of a database connection")
	flag.Parse()

	if *noDBFlag {
		m.NewOnlyFrogsSession(&m.TestSuite{})
	} else {
		m.NewOnlyFrogsSession(&m.Database{})
	}

	r := mux.NewRouter()

	r.HandleFunc("/post", c.AddPost).Methods("POST")
	r.HandleFunc("/post/{uuid}", c.GetPost) // Do we want this?
	r.HandleFunc("/posts", c.GetAllPosts)   // Replace when following exists

	r.HandleFunc("/score", c.AddScore).Methods("POST")

	r.HandleFunc("/follow", c.Follow).Methods("POST")
	r.HandleFunc("/unfollow", c.Unfollow).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":3000", r)

}
