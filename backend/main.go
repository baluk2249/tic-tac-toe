package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tic-tac-toe-backend/game"
)

func main() {
	g := game.NewGame()

	withLogging := func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("%s %s\n", r.Method, r.URL.Path)
			h(w, r)
		}
	}

	withCORS := func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			h(w, r)
		}
	}

	http.HandleFunc("/restart", withLogging(withCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		g.Reset()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(g)
	})))

	http.HandleFunc("/", withLogging(withCORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Tic Tac Toe Backend is running!")
	})))

	http.HandleFunc("/state", withLogging(withCORS(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(g)
	})))

	http.HandleFunc("/move", withLogging(withCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var move struct {
			Row int `json:"row"`
			Col int `json:"col"`
		}
		err := json.NewDecoder(r.Body).Decode(&move)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		ok := g.MakeMove(move.Row, move.Col)
		w.Header().Set("Content-Type", "application/json")
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid move"})
			return
		}
		json.NewEncoder(w).Encode(g)
	})))

	fmt.Println("Backend server started at http://localhost:9090/")
	http.ListenAndServe(":9090", nil)
}
