package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
)

func sqlInjection(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("postgres", "user=admin password=secret dbname=testdb sslmode=disable")
	defer db.Close()

	user := r.URL.Query().Get("user")

	// 🚨 Vulnerability: Directly concatenating user input into SQL (SQL Injection)
	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", user)
	rows, _ := db.Query(query)

	fmt.Fprintf(w, "Query executed: %s", query)
	defer rows.Close()
}
