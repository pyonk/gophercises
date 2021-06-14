package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/pyonk/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filename)

	story := cyoa.LoadStory(*filename)
	h := cyoa.NewHandler(story, cyoa.WithTemplate(nil))
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
