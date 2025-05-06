package main
//1121
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	_ "github.com/lib/pq"
)

const dbUser = "admin"           // Hardcoded credential
const dbPass = "SuperSecret123"  // Hardcoded credential

func vulnerableSQL(w http.ResponseWriter, r *http.Request) {
	userInput := r.URL.Query().Get("id")
	db, _ := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=testdb sslmode=disable", dbUser, dbPass))

	// 🚨 SQL Injection vulnerability
	query := fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", userInput)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "DB error", 500)
		return
	}
	defer rows.Close()

	fmt.Fprintf(w, "Executed query: %s\n", query)
}

func vulnerableCommand(w http.ResponseWriter, r *http.Request) {
	cmd := r.URL.Query().Get("cmd")

	// 🚨 Command Injection vulnerability
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		http.Error(w, "Command error", 500)
		return
	}
	w.Write(out)
}

func main() {
	// 🚨 Insecure HTTP server (no TLS)
	http.HandleFunc("/sql", vulnerableSQL)
	http.HandleFunc("/cmd", vulnerableCommand)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
