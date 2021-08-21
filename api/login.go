package api

import (
	"fmt"
	"net/http"

	"github.com/rmarasigan/portal_system/models"
	"github.com/rmarasigan/portal_system/portal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConfirmLogin : Return JSON object
func ConfirmLogin(w http.ResponseWriter, r *http.Request) {
	/*
	 * r.FormValue is the way on how to access or get your
	 * input value (html), check AJAX data code in templates
	 * -> index.html
	 */
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Checks if your input is not empty
	if username != "" && password != "" {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", portal.MySQLUsername, portal.MySQLPassword, portal.MySQLHostName, portal.MySQLPort, portal.MySQLDBName)
		db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		user := models.User{}
		// https://gorm.io/docs/query.html#String-Conditions
		result := db.Where("username = ? AND user_password = ?", username, password).Find(&user)

		if result.RowsAffected > 0 {
			// return a status with OK value
			portal.ReturnJSON(w, r, map[string]interface{}{
				"status": "ok",
				"user":   user,
			})
		} else {
			// returns a status with ERROR value
			portal.ReturnJSON(w, r, map[string]interface{}{
				"status": "error",
				"msg":    "Incorrect username or password.",
			})
		}
	}
}
