package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luisgomez29/webscreenshot/screenshot"
	"log"
	"net/http"
	"net/url"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/screenshot", ScreenshotHandler).Queries("url", "{url}").Methods("GET")

	fmt.Println("server listening on port :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Response to requests
func response(w http.ResponseWriter, errorCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(data)
}

func ScreenshotHandler(w http.ResponseWriter, r *http.Request) {
	rawURL, err := url.ParseRequestURI(r.URL.Query().Get("url"))
	if err != nil {
		response(w, http.StatusBadRequest, map[string]string{
			"error": "invalid URL",
		})
		return
	}

	img, err := screenshot.GenerateScreenshot(rawURL)
	if err != nil {
		response(w, http.StatusBadRequest, map[string]string{
			"error": "invalid URL",
		})
		return
	}
	response(w, http.StatusOK, map[string]string{
		"screenshot": img,
		"url":        rawURL.String(),
	})
}
