package main

import (
	asciiwebkood "asciiwebkood/functions"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var tpl = template.Must(template.ParseFiles("index.html"))

type PageData struct {
	SubmittedValue string
	OutputHTML     template.HTML
}

func main() {

	//Serving static files, mostly for CSS

	staticFileServer := http.FileServer(http.Dir("static"))

	// Serve CSS files with the correct MIME type (avoid the error in Console)
	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		staticFileServer.ServeHTTP(w, r)
	})))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/submit" && r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if r.Method == "POST" {
			// Get the value from the textarea
			message := r.FormValue("message")
			radioInput := r.FormValue("radioGroup")

			inputSplit := strings.Split(message, "\n")

			asciiValue := ""

			for j := 0; j < len(inputSplit); j++ { // In case we get a newline in the textarea
				asciiValue += asciiwebkood.AsciiProgram(inputSplit[j], radioInput)
			}

			// Create a PageData struct and populate it

			data := PageData{
				SubmittedValue: message,
				OutputHTML:     template.HTML(asciiValue), // Set the OutputHTML field with the HTML value
			}

			// Render the HTML template with the data
			err := tpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			// Render the initial HTML form
			err := tpl.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})

	// Printing data into the console
	port := "8080"
	fmt.Println("ascii-art-web")
	fmt.Printf("Server started at localhost:%s\n", port)
	fmt.Printf("CTRL+C will terminate the server\n")
	http.ListenAndServe(":"+port, nil)
}
