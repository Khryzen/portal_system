package main

import (
	"fmt"
	"net/http"

	"github.com/rmarasigan/portal_system/api"
	"github.com/rmarasigan/portal_system/handler"
	"github.com/rmarasigan/portal_system/models"
	"github.com/rmarasigan/portal_system/portal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func PageSetup() {
	generateAdmin()
	handle()
	handleFunc()

	// Specifying that it should listen on host and port
	http.ListenAndServe(portal.BindIP+portal.Port, nil)
}

func handleFunc() {
	http.HandleFunc("/api/", api.APIHandler)
	http.HandleFunc("/", handler.LoginHandler)
	http.HandleFunc("/portal/", handler.PortalHandler)
}

// Registers the handler for the given pattern in the DefaultServeMux
func handle() {
	// To access the folder and serve it
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
}

// If you have created a table and has no user(s) or admin user, it will create automatically an Admin account
func generateAdmin() {
	// dsn = data source name
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", portal.MySQLUsername, portal.MySQLPassword, portal.MySQLHostName, portal.MySQLPort, portal.MySQLDBName)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Instead of typing repeatedly type User struct, just create a model for that.
	// Go to models -> user.go
	user := models.User{}

	// If user (admin) not found, create a new record
	// https://gorm.io/docs/advanced_query.html#FirstOrCreate
	result := db.FirstOrCreate(&user, models.User{
		Username:     "admin",
		UserPassword: "admin",
		FirstName:    "System",
		LastName:     "Administrator",
		Email:        "admin@email.com",
	})

	if result.RowsAffected != 0 {
		portal.Print(portal.OK, "Admin account is created. Username: admin Password: admin")
	}
}
