package mongodb

import (
	"context"
	"time"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DbInter All Methods of DB
type DbInter interface {
	GetDatabase(string) *mongo.Database
}

//Database has all dependecies of mongodb
type Database struct {
	Client *mongo.Client
	DB     *mongo.Database
}

//NewDb returns Database instance
func NewDb(uri string, dbName string) *Database {
	db := &Database{}
	color.Yellow("MONGODB -----------> Conectando com o MongoDb...")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	dbError(err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = client.Connect(ctx)
	dbError(err)
	color.Green("\nMONGODB -------------> Mongodb conectado com sucesso! \n")

	db.Client = client
	db.DB = db.GetDatabase(dbName)
	return db
}

//GetDatabase return the right database to connect
func (d *Database) GetDatabase(name string) *mongo.Database {
	db := d.Client.Database(name)

	return db
}

func dbError(err error) {
	if err != nil {
		color.Red("\nMONGODB --------------------> Aconteceu um erro ao conectar o mongoDb: ", err)
	}
}
