package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// URL represents the structure to store URLs in the "database"
type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

//Example entry in the map (in-memory database):
/*
	d9736711 --> {
					ID: "d9736711",
					OriginalURL: "https://github.com/Prince-1501/",
					ShortURL: "d9736711",
					CreationDate: time.Now()
				}
*/

// In-memory database to store URLs using a map
var urlDB = make(map[string]URL)

// Function to create a shortened version of a given original URL
func generateShortURL(OriginalURL string) string {
	// Create a new MD5 hasher
	hasher := md5.New()
	// Write the OriginalURL to the hasher as a byte slice
	hasher.Write([]byte(OriginalURL))
	fmt.Println("hasher: ", hasher)

	// Generate the hash value
	data := hasher.Sum(nil)
	fmt.Println("hasher data: ", data)

	// Convert the hash to a human-readable hex string
	hash := hex.EncodeToString(data)
	fmt.Println("Encode to string: ", hash)
	fmt.Println("Final string: ", hash[:8])

	// Return the first 8 characters of the hex string as the short URL
	return hash[:8]
}

func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL) // Generate the short URL
	id := shortURL                            // Use the short URL as the ID for simplicity

	// Insert the new entry into the in-memory database
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

// Function to retrieve the original URL based on its short ID
func getURL(id string) (URL, error) {
	url, ok := urlDB[id] // Check if the ID exists in the map
	if !ok {
		return URL{}, errors.New("url not found") // Return an error if not found
	}
	return url, nil // Return the URL object if found
}

// Root page handler, responds with a simple "Hello World"
func RootPageURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// Handler for creating short URLs from original URLs
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"` // Extract the "url" field from the JSON request body
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	shortURL_ := createURL(data.URL)
	//fmt.Fprintf(w, shortURL)
	//create json format
	// Create a response struct to send back the short URL as JSON
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	// Set the content type to JSON
	w.Header().Set("content_Type", "application/json")
	// Encode and send the response as JSON
	json.NewEncoder(w).Encode(response)
}

// Handler to redirect short URLs to their original URLs
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the short ID from the request path
	id := r.URL.Path[len("/redirect/"):]
	// Retrieve the original URL from the database
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "url not found", http.StatusNotFound) //// Return a 404 if not found
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {

	// Register the handler function to handle all requests to the root URL ("/")

	http.HandleFunc("/", RootPageURL)                 // Root path
	http.HandleFunc("/shorten", ShortURLHandler)      // Path for creating short URLs
	http.HandleFunc("/redirect/", redirectURLHandler) //// Path for redirecting short URLs

	//start the HTTp server on port 3000

	fmt.Println("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on starting server", err) //// Print error if server fails to start
	}

}
