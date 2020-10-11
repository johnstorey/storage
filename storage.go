package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo variables.

type DataStore struct {
	host         string
	port         string
	user         string
	password     string
	rootuser     string
	rootpassword string
	client       *mongo.Client
}

func (ds *DataStore) init() {
	if ds.host == "" {
		e := ds.populateConfiguration()
		if e != nil {
			fmt.Println("Failed to configure datastore for use: ", e)
		}

		// Setup database connection.
		// TODO need to add a disconnect method that main() can call on defer.
		var err error
		// TODO find out why this does not need a username and password
		clientOptions := options.Client().ApplyURI("mongodb://" + ds.host + ":" + ds.port)
		ds.client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		// Check the database connection.
		err = ds.client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (ds *DataStore) disconnect() error {
	// if err := ds.client.Disconnect(); err != nil {
	//panic(err)
	//}
	return nil
}

func (ds *DataStore) populateConfiguration() error {
	ds.host = os.Getenv("MONGOHOST")
	ds.port = os.Getenv("MONGOPORT")
	ds.user = os.Getenv("MONGOUSER")
	ds.password = os.Getenv("MONGOUSERPASSWORD")
	ds.rootuser = os.Getenv("MONGOROOTUSER")
	ds.rootpassword = os.Getenv("MONGOROOTPASSWORD")

	return nil
}

func (ds *DataStore) numEnvironments() (int, error) {
	ds.init()
	return 0, nil
}

// TODO determine how to convert the insertResult back to a *MesmerEnvironment and return it.
func (ds *DataStore) saveEnvironment(env *MesmerEnvironment) error {
	ds.init()
	envCollection := ds.client.Database("mesbot").Collection("environments")
	if envCollection == nil {
		return errors.New("No collection 'environments' found in database.")
	}
	res, err := envCollection.InsertOne(context.TODO(), env)
	if err != nil {
		log.Fatal(err)
	}

	// The only change should be the id.
	env.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (ds *DataStore) updateEnvironment(env *MesmerEnvironment) error {
	ds.init()
	envCollection := ds.client.Database("mesbot").Collection("environments")
	if envCollection == nil {
		return errors.New("No collection 'environments' found in database.")
	}

	// Update the document with all elements to simplify the code.
	var setElements bson.D
	setElements = append(setElements, bson.E{"name", env.Name})
	setElements = append(setElements, bson.E{"mesmernodes", env.MesmerNodes})
	setMap := bson.D{{"$set", setElements}}
	_, err := envCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": env.ID},
		setMap,
	)
	if err != nil {
		log.Fatal(err)
	}
	return nil

}

func (ds *DataStore) getEnvironment(id int32) (MesmerEnvironment, error) {
	ds.init()
	return MesmerEnvironment{}, nil
}

func (ds *DataStore) findEnvironmentByID(id primitive.ObjectID) (*MesmerEnvironment, error) {
	ds.init()
	envCollection := ds.client.Database("mesbot").Collection("environments")
	convertedID, err := primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		panic(err)
	}

	filter := bson.D{{"_id", convertedID}}
	var foundEnv bson.M
	err = envCollection.FindOne(context.TODO(), filter).Decode(&foundEnv)

	// Convert Mongo document to a MesmerEnvironment struct.
	var newEnv *MesmerEnvironment
	newEnv = ds.mapMesmerEnvironmentFromMongo(foundEnv)
	return newEnv, nil
}

func (ds *DataStore) findEnvironmentByName(name string) (*MesmerEnvironment, error) {
	ds.init()
	envCollection := ds.client.Database("mesbot").Collection("environments")
	filter := bson.D{{"name", name}}
	//var environment MesmerEnvironment

	var environment bson.M
	envCollection.FindOne(context.TODO(), filter).Decode(&environment)

	// Take the map returned by Mongo and convert to the MesmerEnvironment format.
	var emptyEnv *MesmerEnvironment
	emptyEnv = ds.mapMesmerEnvironmentFromMongo(environment)
	return emptyEnv, nil
}

func (ds *DataStore) NodesForEnvironment(env *MesmerEnvironment) ([]MesmerNode, error) {
	return nil, nil
}

func (ds *DataStore) AndroidNodesForEnvironment(env *MesmerEnvironment) ([]MesmerNode, error) {
	return nil, nil
}

func (ds *DataStore) IOSNodesForEnvironment(env *MesmerEnvironment) ([]MesmerNode, error) {
	return nil, nil
}

func (ds *DataStore) emptyCollection(collection string) error {
	ds.init()
	targetCollection := ds.client.Database("mesbot").Collection(collection)
	if targetCollection == nil {
		return errors.New("Attempt to empty nonexistent collection " + collection)
	}
	_, err := targetCollection.DeleteMany(context.TODO(), bson.D{})
	return err
}

func (ds *DataStore) findNodesByEnvironmentID(id primitive.ObjectID) (*[]MesmerNode, error) {
	ds.init()
	nodeCollection := ds.client.Database("mesbot").Collection("nodes")
	convertedID, err := primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		panic(err)
	}

	filter := bson.D{{"env", convertedID}}
	cursor, err := nodeCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	// Go through the results.
	for cursor.Next(context.TODO()) {
		var nodeResultHolder MesmerNode

		err := cursor.Decode(&nodeResultHolder)
		if err != nil {
			log.Fatal(err)
		}

		// Convert result to mesmer node.
		//result := convertBSONToStruct(
	}

	return &[]MesmerNode{}, nil
}

func (ds *DataStore) mapMesmerEnvironmentFromMongo(source primitive.M) *MesmerEnvironment {
	var newEnv MesmerEnvironment
	var results []MesmerNode

	for k, _ := range source {

		switch k {
		case "_id":
			newEnv.ID = source["_id"].(primitive.ObjectID)
		case "name":
			newEnv.Name = source["name"].(string)
		case "mesmernodes":
			// Guard against there being no nodes
			if source["mesmernodes"] == nil {
				break
			}

			for _, v2 := range source["mesmernodes"].(primitive.A) {
				var newNode MesmerNode
				for k3, v3 := range v2.(primitive.M) {
					switch k3 {
					case "host":
						newNode.Host = v3.(string)
					case "ip":
						newNode.IP = v3.(string)
					case "nodetype":
						newNode.NodeType = v3.(string)
					}
				}
				results = append(results, newNode)
			}
		}
	}

	newEnv.MesmerNodes = results
	return &newEnv
}
