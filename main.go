package main

import (
	_ "github.com/VictorNapoles/teste-victor-conductor/adapter/crud"
	_ "github.com/VictorNapoles/teste-victor-conductor/config"
	_ "github.com/VictorNapoles/teste-victor-conductor/config/database"
	_ "github.com/VictorNapoles/teste-victor-conductor/config/http"
	conductor "github.com/VictorNapoles/teste-victor-conductor/config/http"
	_ "github.com/VictorNapoles/teste-victor-conductor/controller"
	_ "github.com/VictorNapoles/teste-victor-conductor/docs"
	_ "github.com/VictorNapoles/teste-victor-conductor/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /conductor/v1/api
func main() {

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	conductor.GetServer().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	conductor.GetServer().Run()
}
