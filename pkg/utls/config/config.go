package config

import (
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

//Conf encapsulates main struct of app
type Conf struct {
	Application *App
	DB          *Database
}

//App is responsible for all server options
type App struct {
	Port string
}

//Database is responsible for all dependencies from database
type Database struct {
	URI    string
	DBName string
}

//NewConf generates a new app
func NewConf() *Conf {
	app := &Conf{}
	err := godotenv.Load()
	envError(err)

	app.Application = configApp()
	app.DB = configDatabase()

	return app
}

func configApp() *App {
	a := &App{}

	port := os.Getenv("PORT")
	if port == "" {
		port = "2000"
	}

	a.Port = port

	return a
}

func configDatabase() *Database {
	a := &Database{}

	URI := os.Getenv("MONGODB_URI")
	if URI == "" {
		URI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DBNAME")
	if dbName == "" {
		dbName = "Default"
	}

	a.DBName = dbName
	a.URI = URI

	return a
}

func envError(err error) {
	if err != nil {
		color.Red("Aconteceu um erro ao ler o dotEnv: ", err, "\n")
	}
}
