package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/success", successHandler)

	// Menjalankan server HTTP pada port 8080
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if isValidLogin(username, password) {
			// Menyimpan informasi sesi pengguna setelah login berhasil
			session, _ := store.Get(r, "session-name")
			session.Values["username"] = username
			session.Save(r, w)

			http.Redirect(w, r, "/success", http.StatusFound)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.ServeFile(w, r, "login.html")
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil informasi sesi pengguna dari cookie sesi
	session, _ := store.Get(r, "session-name")
	username, ok := session.Values["username"].(string)
	if !ok {
		// Jika tidak ada informasi sesi, arahkan pengguna kembali ke halaman login
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintf(w, "Login berhasil! Selamat datang, %s!", username)
}

func isValidLogin(username, password string) bool {
	// Implementasi validasi login di sini
	// Contoh sederhana: jika username dan password adalah "admin", maka login berhasil
	return username == "admin" && password == "admin"
}
