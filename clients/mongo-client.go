package clients

import (
	"errors"
	"fmt"

	// models "github.com/advwacloud/datahub-edge-domain-models"
	models "github.com/advwacloud/datahub-edge-domain-models"

	"gopkg.in/mgo.v2"
)

const (
	DATAS_COLLECTION = "datahub_edge_data"
)

var currentMongoClient *MongoClient

// Type used to sort the readings by creation date
type ByDataCreationDate []models.Data

type MongoClient struct {
	Session  *mgo.Session  // Mongo database session
	Database *mgo.Database // Mongo database
}

// Return a pointer to the MongoClient
func newMongoClient(config DBConfiguration) (*MongoClient, error) {
	// Create the dial info for the Mongo session
	// connectionString := config.Host + ":" + strconv.Itoa(config.Port)
	connectionString := config.Host + ":" + config.Port
	fmt.Println("INFO: Connecting to mongo at: " + connectionString)
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs: []string{connectionString},
		//Timeout:  time.Duration(config.Timeout) * time.Millisecond,
		Database: config.DatabaseName,
		Username: config.Username,
		Password: config.Password,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		fmt.Println("Error dialing the mongo server: " + err.Error())
		return nil, err
	}

	mongoClient := &MongoClient{Session: session, Database: session.DB(config.DatabaseName)}
	currentMongoClient = mongoClient // Set the singleton
	return mongoClient, nil
}

// Get the current Mongo Client
func GetCurrentMongoClient() (*MongoClient, error) {
	if currentMongoClient == nil {
		return nil, errors.New("No current mongo client, please create a new client before requesting it")
	}

	return currentMongoClient, nil
}

// Get a copy of the session
func (mc *MongoClient) GetSessionCopy() *mgo.Session {
	return mc.Session.Copy()
}

// Post a new data
func (mc *MongoClient) AddData(r models.Data) error {
	s := mc.GetSessionCopy()
	defer s.Close()

	// Get the data ready
	//r.Id = bson.NewObjectId()
	//r.Created = time.Now().UnixNano() / int64(time.Millisecond)
	//r.Created = time.Now()

	err := s.DB(mc.Database.Name).C(DATAS_COLLECTION).Insert(&r)
	return err
}
