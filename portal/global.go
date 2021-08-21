package portal

// Global Variables
const (
	// System Start Up
	BindIP = "0.0.0.0"
	Port   = ":8081"
	// Database
	MySQLDBName   = "portal_system"
	MySQLUsername = "root"
	MySQLHostName = "localhost"
	MySQLPort     = 3306
	MySQLPassword = "Your Password Here"
	// Report Level
	OK    = 0
	DEBUG = 1
	ERROR = 2
)

var ReportingLevel = OK

var levelMap = map[int]string{
	OK:    "[   OK   ]   ",
	DEBUG: "[  DEBUG  ]   ",
	ERROR: "[  ERROR  ]   ",
}
