package auth

import (
	"adm/api/pkg/utls/database/mongodb"
	"adm/api/pkg/utls/server"
)

type Auth struct {
}

//NewHTTP is responsible for all routes related to Auth
func NewHTTP(db *mongodb.Database, app *server.Server) {

}
