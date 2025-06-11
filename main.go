package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	//Database connection
	var err error
	db, err = sql.Open("mysql", "root:Mohit2607@@tcp(127.0.0.1:330)/user_db")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	// Serve static files (e.g., HTML, CSS, JS files)
	fs := http.FileServer(http.Dir("./")) // serve files from the current directory
	http.Handle("/", fs)

	//Routes
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logged", loggedHandler)

	//Setting up server
	port := ":8080"
	fmt.Printf("Server starting at port: %v\n", port)
	//Starting Server - http://localhost:8080/
	err1 := http.ListenAndServe(port, nil)
	if err1 != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// Registration handler
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	//Getting data for registration
	fullname := r.FormValue("fullname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	mobile := r.FormValue("mobile")
	password := r.FormValue("password")
	gender := r.FormValue("gender")
	dob := r.FormValue("dob")
	currency := r.FormValue("currency")
	location := r.FormValue("location")
	terms := r.FormValue("terms")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Convert "terms" (checkbox) to boolean
	agreed := false
	if terms == "on" {
		agreed = true
	}

	// Insert into the database
	_, err := db.Exec(`
        INSERT INTO users (fullname, username, email, mobile, password, gender, dob, currency, location, terms)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		fullname, username, email, mobile, password, gender, dob, currency, location, agreed,
	)

	if err != nil {
		http.Error(w, "Could not register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to login
	http.Redirect(w, r, "/login.html", http.StatusSeeOther)
}

// Login Handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	//Login data
	username := r.FormValue("identifier")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	//Retrieving the stored password
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//Validating
	if storedPassword != password {
		http.Error(w, "Invaalid username or password", http.StatusUnauthorized)
		return
	}

	//Login to dashboard
	http.Redirect(w, r, "/logged", http.StatusSeeOther)
}

func loggedHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/logged.html", http.StatusSeeOther)
	fmt.Println("Logged In")
}
