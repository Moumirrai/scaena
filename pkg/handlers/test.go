package handlers

import (
	"database/sql"
	"net/http"
)

func TestHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//send back sucess with hello world
		w.Write([]byte("Hello World"))
	}
}
