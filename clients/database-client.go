package clients

import (
	"errors"
	"fmt"
	"os"

	// models "github.com/advwacloud/datahub-edge-domain-models"
	models "github.com/advwacloud/datahub-edge-domain-models"
)

type dbClient interface {
	AddData(r models.Data) error
}

type DBConfiguration struct {
	DbType       string
	Host         string
	Port         string
	Timeout      int
	DatabaseName string
	Username     string
	Password     string
}

var ErrNotFound error = errors.New("Item not found")
var ErrUnsupportedDatabase error = errors.New("Unsuppored database type")
var ErrInvalidObjectId error = errors.New("Invalid object ID")
var ErrNotUnique error = errors.New("Resource already exists")

var Dbc dbClient

func init() {
	// Create a database client
	Dbc, _ = newDBClient(DBConfiguration{
		DbType:       os.Getenv("DB_TYPE"),
		Host:         os.Getenv("MONGODB_HOST"),
		Port:         os.Getenv("MONGODB_PORT"),
		DatabaseName: os.Getenv("MONGODB_DATABASE_NAME"),
		Username:     os.Getenv("MONGODB_USER"),
		Password:     os.Getenv("MONGODB_PWD"),
	})
}

// Return the dbClient interface
func newDBClient(config DBConfiguration) (dbClient, error) {
	switch config.DbType {
	case "MONGO":
		// Create the mongo client
		mc, err := newMongoClient(config)
		if err != nil {
			fmt.Println("Error creating the mongo client: " + err.Error())
			return nil, err
		}
		return mc, nil
	default:
		return nil, ErrUnsupportedDatabase
	}
}
