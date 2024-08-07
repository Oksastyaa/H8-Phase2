package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Start the web server
	// http.ListenAndServe(":8080", nil)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: setupRoute(),
	}
	// print server address
	fmt.Println("Server is running on: ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

}

func setupRoute() *http.ServeMux {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Handle the / route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		write, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", string(rune(write)))
	})

	//handle post request
	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Parse the form data
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Initialize a variable to store the formatted key-value pairs
			var msg string
			for key, values := range r.Form {
				for _, value := range values {
					msg += fmt.Sprintf("%s: %s\n", key, value)
				}
			}

			// Write the response
			_, err = w.Write([]byte(msg))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Set headers and status code
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Return the ServeMux
	return mux
}
