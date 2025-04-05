package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var store *URLStore

func main() {
	// Initialize the URL store
	store = NewURLStore("urls.json")

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/r/", redirectHandler)

	// Start the web server
	fmt.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Home handler to display the form
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the template and check for errors
	tmpl := `
		<html>
			<head><title>URL Shortener</title></head>
			<body>
				<h1>URL Shortener</h1>
				<form action="/shorten" method="POST">
					<label for="url">Enter URL to shorten:</label><br>
					<input type="text" id="url" name="url" required><br><br>
					<button type="submit">Shorten URL</button>
				</form>
			</body>
		</html>
	`
	t, err := template.New("home").Parse(tmpl) // Correcting the variable assignment
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil) // Execute the template with no dynamic data
}

// Shorten handler to process the URL and provide the shortened link
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		longURL := r.FormValue("url")

		// Generate a short URL
		slug := store.Add(longURL)
		shortURL := fmt.Sprintf("http://localhost:8080/r/%s", slug)

		// Set the Content-Type to HTML
		w.Header().Set("Content-Type", "text/html")

		// Render a complete HTML page with the shortened URL
		htmlResponse := fmt.Sprintf(`
			<html>
			<head><title>Shortened URL</title></head>
			<body>
				<h1>Shortened URL</h1>
				<p>Your shortened URL is: <a href='%s'>%s</a></p>
			</body>
			</html>
		`, shortURL, shortURL)

		// Send the response
		fmt.Fprint(w, htmlResponse)
	}
}

// Redirect handler to resolve short URLs and redirect users
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the slug from the URL path
	slug := strings.TrimPrefix(r.URL.Path, "/r/")

	// Look up the long URL associated with the slug
	longURL, found := store.Get(slug)
	if !found {
		http.NotFound(w, r)
		return
	}

	// Redirect the user to the original long URL
	http.Redirect(w, r, longURL, http.StatusFound)
}
