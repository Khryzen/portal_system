package handler

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/rmarasigan/portal_system/models"
	"github.com/rmarasigan/portal_system/portal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func PortalHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/portal/")
	path := r.URL.Path

	if strings.HasPrefix(path, "users") {
		UsersHandler(w, r)
		return
	}

	if strings.HasPrefix(path, "profile") {
		ProfileHandler(w, r)
		return
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	data["Title"] = "Portal System | Users"

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", portal.MySQLUsername, portal.MySQLPassword, portal.MySQLHostName, portal.MySQLPort, portal.MySQLDBName)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// A list of user will be selected
	// so it would be an array type.
	user := []models.User{}

	// Select all users
	db.Find(&user)
	data["Users"] = user

	// Parsing templates from files
	tmp := template.Must(template.ParseFiles("./templates/users.html"))
	tmp.Execute(w, data)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	data["Title"] = "Portal System | Profile"

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", portal.MySQLUsername, portal.MySQLPassword, portal.MySQLHostName, portal.MySQLPort, portal.MySQLDBName)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	user := models.User{}

	// https://gorm.io/docs/query.html#String-Conditions
	db.Where("username = ?", "admin").Find(&user)
	data["User"] = user

	// Parsing templates from files
	tmp := template.Must(template.ParseFiles("./templates/profile.html"))
	tmp.Execute(w, data)
}
