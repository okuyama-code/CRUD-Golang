// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// var users = map[string]string{
// 	"john":  "password123",
// 	"alice": "qwerty456",
// }

// func authMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		username, password, ok := r.BasicAuth()

// 		if !ok || !checkCredentials(username, password) {
// 			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

// func checkCredentials(username, password string) bool {
// 	storedPassword, ok := users[username]
// 	if !ok || storedPassword != password {
// 		return false
// 	}
// 	return true
// }

// func secureHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "認証が必要なページです")
// }

// func publicHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "認証不要なページです")
// }

// func main() {
// 	http.HandleFunc("/", publicHandler)
// 	http.Handle("/secure", authMiddleware(http.HandlerFunc(secureHandler)))

// 	fmt.Println("サーバーをポート8080で起動します...")
// 	http.ListenAndServe(":8080", nil)
// }

