package clients

import (
	"fmt"

	models "github.com/advwacloud/datahub-edge-domain-models"
	//"datahub-edge-core/pkg/models" // temporary

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Struct that wraps an event to handle DBRefs
type MongoEvent struct {
	models.Event
}

// Custom marshaling into mongo
func (me MongoEvent) GetBSON() (interface{}, error) {
	// Turn the readings into DBRef objects
	var readings []mgo.DBRef
	for _, reading := range me.Datas {
		readings = append(readings, mgo.DBRef{Collection: DATAS_COLLECTION, Id: reading.Id})
	}

	return struct {
		ID      bson.ObjectId `bson:"_id,omitempty"`
		Device  string        `bson:"device"` // Device identifier (name or id)
		Created int64         `bson:"created"`
		Datas   []mgo.DBRef   `bson:"datas"` // List of datas
	}{
		ID:      me.ID,
		Device:  me.Device,
		Created: me.Created,
		Datas:   readings,
	}, nil
}

// Custom unmarshaling out of mongo
func (me *MongoEvent) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		ID      bson.ObjectId `bson:"_id,omitempty"`
		Device  string        `bson:"device"` // Device identifier (name or id)
		Created int64         `bson:"created"`
		Datas   []mgo.DBRef   `bson:"datas"` // List of datas
	})

	bsonErr := raw.Unmarshal(decoded)
	if bsonErr != nil {
		return bsonErr
	}

	// Copy over the non-DBRef fields
	me.ID = decoded.ID
	me.Device = decoded.Device
	me.Created = decoded.Created

	// De-reference the DBRef fields
	mc, err := getCurrentMongoClient()
	if err != nil {
		fmt.Println("Error getting a mongo client: " + err.Error())
		return err
	}

	var readings []models.Data

	// Get all of the reading objects
	for _, rRef := range decoded.Datas {
		var reading models.Data
		err := mc.Database.C(DATAS_COLLECTION).FindId(rRef.Id).One(&reading)
		if err != nil {
			return err
		}

		readings = append(readings, reading)
	}

	me.Datas = readings

	return nil
}
