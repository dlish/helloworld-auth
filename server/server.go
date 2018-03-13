package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/dlish/helloworld-auth/db"
	_ "github.com/lib/pq"
)

func NewServer(connection *sql.DB) *http.Server {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/api/auth/", handleRequest(db.NewAccount(connection)))
	return &server
}

func handleRequest(t db.CrudRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request, repo db.CrudRepository) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = repo.Fetch(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(repo, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
