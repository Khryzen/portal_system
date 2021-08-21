package main

import (
	"github.com/rmarasigan/portal_system/portal"
)

func main() {
	SystemStartUp()
	PageSetup()
}

func SystemStartUp() {
	portal.Print(portal.OK, "Portal System running http://%v%s/...", portal.BindIP, portal.Port)
}
