package api

import (
	"adm/api/pkg/api/auth"
	"adm/api/pkg/utls/config"
	"adm/api/pkg/utls/database/mongodb"
	"adm/api/pkg/utls/server"

	"github.com/fatih/color"
)

//Start initiates server and all functions
func Start(c *config.Conf) {
	app := server.NewServer()
	db = mongodb.NewDb(c.DB.URI, c.DB.DBName)

	auth.NewHTTP(db, app)

	color.Green("\nSERVER ---------------> Servidor Rodando na porta: http://localhost:" + c.Application.Port)
	app.App.Run(":" + c.Application.Port)

}
