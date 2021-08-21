package models

/**********
 * User Table Fields
 * username, user_password,
 * first_name, last_name
 * email, contact_number, last_login
 *
 * For struct:
 * It should be patterned in your db table users. For this,
 * we're just selecting the field starting from username
 * until email only.
 **********/
type User struct {
	Username     string
	UserPassword string
	FirstName    string
	LastName     string
	Email        string
}
