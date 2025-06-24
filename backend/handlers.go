package main

import (
	"io"
	"net/http"
	"os"
)

func DetailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	tmdbUrl := "https://api.themoviedb.org/3/movie/" + id + "?api_key=" + os.Getenv("TMDB_API_KEY")
	proxyRequest(tmdbUrl, w)
}

func proxyRequest(url string, w http.ResponseWriter) {
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch", 500)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
