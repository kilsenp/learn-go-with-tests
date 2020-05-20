package main

import (
	"github.com/kilsenp/application"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {

	store, close, error = FileSystemPlayerStoreFromFile(dbFileName)

	if error != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	} else {
		log.Printf("server started on port 5000")
	}

}
