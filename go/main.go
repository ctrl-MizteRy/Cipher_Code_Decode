package main

import (
	"cipher/http-request"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("Could not get into the home directorry %v", home))
	}
	path := filepath.Join(home, "Desktop/Projects/HTML/Cipher/front")
	fs := http.FileServer(http.Dir(path))
	http.Handle("/", fs)
	http.HandleFunc("/process", httprequest.HandleMessage)
	log.Fatal(http.ListenAndServe("https://ctrl-miztery.github.io/Cipher_Code_Decode/", nil))
}
