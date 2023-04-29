package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username, Password string
}

func main() {

	r := httprouter.New()
	r.GET("/", Anasayfa)
	r.POST("/login", Login)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("views"))))
	http.ListenAndServe(":8080", r)

}
func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, _ := template.ParseFiles("views/index.html")

	view.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}
	var user User
	username := r.FormValue("username")
	password := r.FormValue("password")
	db.First(&user, "username =?", username)
	if user.Password == password {
		view, _ := template.ParseFiles("views/login.html")
		data := User{
			Username: user.Username,
			Password: user.Password,
		}
		view.Execute(w, data)
	} else {
		http.Error(w, "Şifre Yanlış", http.StatusUnauthorized)
	}

}
