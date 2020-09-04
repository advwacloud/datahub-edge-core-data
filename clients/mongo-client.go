package clients

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	models "github.com/advwacloud/datahub-edge-domain-models"
	//"datahub-edge-core/pkg/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	EVENTS_COLLECTION = "datahub_edge_event"
	DATAS_COLLECTION  = "datahub_edge_data"
)

var currentMongoClient *MongoClient // Singleton used so that MongoEvent can use it to de-reference readings

/*
 Core data client
 Has functions for interacting with the core data mongo database
*/

// Type used to sort the readings by creation date
type ByDataCreationDate []models.Data

// func (a ByDataCreationDate) Len() int           { return len(a) }
// func (a ByDataCreationDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByDataCreationDate) Less(i, j int) bool { return (a[i].Created < a[j].Created) }

type MongoClient struct {
	Session  *mgo.Session  // Mongo database session
	Database *mgo.Database // Mongo database
}

// Return a pointer to the MongoClient
func newMongoClient(config DBConfiguration) (*MongoClient, error) {
	// Create the dial info for the Mongo session
	connectionString := config.Host + ":" + strconv.Itoa(config.Port)
	fmt.Println("INFO: Connecting to mongo at: " + connectionString)
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{connectionString},
		Timeout:  time.Duration(config.Timeout) * time.Millisecond,
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
func getCurrentMongoClient() (*MongoClient, error) {
	if currentMongoClient == nil {
		return nil, errors.New("No current mongo client, please create a new client before requesting it")
	}

	return currentMongoClient, nil
}

// Get a copy of the session
func (mc *MongoClient) GetSessionCopy() *mgo.Session {
	return mc.Session.Copy()
}

// ************************ Datas ************************************8

// Return a list of Datas sorted by reading id
func (mc *MongoClient) Datas() ([]models.Data, error) {
	return mc.getDatas(nil)
}

// Post a new data
func (mc *MongoClient) AddData(r models.Data) (bson.ObjectId, error) {
	s := mc.GetSessionCopy()
	defer s.Close()

	// Get the data ready
	r.Id = bson.NewObjectId()
	//r.Created = time.Now().UnixNano() / int64(time.Millisecond)
	r.Created = time.Now()

	err := s.DB(mc.Database.Name).C(DATAS_COLLECTION).Insert(&r)
	return r.Id, err
}

// Update a data
// 404 - data cannot be found
// 409 - Value descriptor doesn't exist
// 503 - unknown issues
func (mc *MongoClient) UpdateData(r models.Data) error {
	s := mc.GetSessionCopy()
	defer s.Close()

	//r.Modified = time.Now().UnixNano() / int64(time.Millisecond)
	r.Created = time.Now()
	// Update the reading
	err := s.DB(mc.Database.Name).C(DATAS_COLLECTION).UpdateId(r.Id, r)
	if err == mgo.ErrNotFound {
		return ErrNotFound
	}

	return err
}

// Get a dara by ID
func (mc *MongoClient) DataById(id string) (models.Data, error) {
	// Check if the id is a id hex
	if !bson.IsObjectIdHex(id) {
		return models.Data{}, ErrInvalidObjectId
	}

	query := bson.M{"_id": bson.ObjectIdHex(id)}

	return mc.getData(query)
}

// Get the count of readings in Mongo
func (mc *MongoClient) DataCount() (int, error) {
	s := mc.GetSessionCopy()
	defer s.Close()

	return s.DB(mc.Database.Name).C(DATAS_COLLECTION).Find(bson.M{}).Count()
}

// Delete a reading by ID
// 404 - can't find the reading with the given id
func (mc *MongoClient) DeleteDataById(id string) error {
	// Check if the id is a bson id
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidObjectId
	}

	return mc.deleteById(id, DATAS_COLLECTION)
}

// Get datas from the database
func (mc *MongoClient) getDatas(q bson.M) ([]models.Data, error) {
	s := mc.GetSessionCopy()
	defer s.Close()

	readings := []models.Data{}
	err := s.DB(mc.Database.Name).C(DATAS_COLLECTION).Find(q).All(&readings)
	return readings, err
}

// Get a data from the database with the passed query
func (mc *MongoClient) getData(q bson.M) (models.Data, error) {
	s := mc.GetSessionCopy()
	defer s.Close()

	var res models.Data
	err := s.DB(mc.Database.Name).C(DATAS_COLLECTION).Find(q).One(&res)
	if err == mgo.ErrNotFound {
		return res, ErrNotFound
	}
	return res, err
}

// Delete from the collection based on ID
func (mc *MongoClient) deleteById(id string, col string) error {
	s := mc.GetSessionCopy()
	defer s.Close()

	// Check if id is a hexstring
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidObjectId
	}

	err := s.DB(mc.Database.Name).C(col).RemoveId(bson.ObjectIdHex(id))
	if err == mgo.ErrNotFound {
		return ErrNotFound
	}
	return err
}
