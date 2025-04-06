package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var store *URLStore

func main() {
	store = NewURLStore("urls.json")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/r/", redirectHandler)
	http.HandleFunc("/list", listHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{}
	for slug := range store.All() {
		data[slug] = fmt.Sprintf("http://localhost:8080/r/%s", slug)
	}
	templates.ExecuteTemplate(w, "list.html", data)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		longURL := r.FormValue("url")

		// Validate URL
		_, err := url.ParseRequestURI(longURL)
		if err != nil {
			message := template.HTML("<span style='color:red;'>Invalid URL format. Please try again.</span>")
			templates.ExecuteTemplate(w, "home.html", message)
			return
		}

		existingSlug := store.FindSlugByURL(longURL)
		if existingSlug != "" {
			message := template.HTML(fmt.Sprintf("This URL has already been shortened: <a href='http://localhost:8080/r/%s'>http://localhost:8080/r/%s</a>", existingSlug, existingSlug))
			templates.ExecuteTemplate(w, "home.html", message)
			return
		}

		slug := store.Add(longURL)
		fmt.Printf("Shortened URL: http://localhost:8080/r/%s\n", slug)
		http.Redirect(w, r, "/list", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/r/"):]
	if longURL, ok := store.Get(slug); ok {
		http.Redirect(w, r, longURL, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}
