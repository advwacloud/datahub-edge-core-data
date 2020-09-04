package clients

import (
	"errors"
	"fmt"

	models "github.com/advwacloud/datahub-edge-domain-models"
	//"datahub-edge-core/pkg/models" // temporary

	"gopkg.in/mgo.v2/bson"
)

type DatabaseType int8 // Database type enum
const (
	MONGO DatabaseType = iota
)

type DBClient interface {
	Datas() ([]models.Data, error)

	DataById(id string) (models.Data, error)

	AddData(r models.Data) (bson.ObjectId, error)

	UpdateData(r models.Data) error

	DeleteDataById(id string) error
}

type DBConfiguration struct {
	DbType       DatabaseType
	Host         string
	Port         int
	Timeout      int
	DatabaseName string
	Username     string
	Password     string
}

var ErrNotFound error = errors.New("Item not found")
var ErrUnsupportedDatabase error = errors.New("Unsuppored database type")
var ErrInvalidObjectId error = errors.New("Invalid object ID")
var ErrNotUnique error = errors.New("Resource already exists")

// Return the dbClient interface
func NewDBClient(config DBConfiguration) (DBClient, error) {
	var dbClient DBClient
	switch config.DbType {
	case MONGO:
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

	return dbClient, nil
}
