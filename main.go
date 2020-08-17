package main

import (
	"github.com/thaitanloi365/go-monitor/cmd"
)

// @title Go Monitor api docs
// @version 1.0
// @description This is a go-monitor server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email thaitanloi365@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8080
// @schemes http https
func main() {
	cmd.Execute()

}
