package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var jwtSecret = []byte("gojwtlogin")

type User struct {
	Username string
	Password string
}

func authenticateUser(username, password string) bool {
	// Burada, kullanıcı kimlik doğrulama işlemlerini gerçekleştirin.
	// Örneğin, kullanıcı adı ve şifreyi bir veritabanında kontrol edebilirsiniz.
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}
	var user User
	result := db.Where("username = ?", username).First(&user)
	return result.Error == nil && result.RowsAffected > 0 && user.Password == password
}

func generateJWT(username string) (string, error) {
	// Token'ın geçerlilik süresi
	expirationTime := time.Now().Add(30 * time.Minute)

	// JWT içeriği ve ayarları
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   username,
	}

	// JWT oluşturma
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomePage)
	router.HandleFunc("/login", Login).Methods("POST")
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("views"))))
	http.ListenAndServe(":8080", router)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("views/index.html")
	view.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if authenticateUser(username, password) {
		token, err := generateJWT(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "JWT oluşturulamadı.")
			return
		}
		w.Header().Set("Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)
		response := struct {
			Token string `json:"token"`
			User  User   `json:"user"`
		}{
			Token: token,
			User: User{
				Username: username,
				Password: password,
			},
		}
		view, _ := template.ParseFiles("views/login.html")
		view.Execute(w, response)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Kimlik doğrulama başarısız.")
	}
}
