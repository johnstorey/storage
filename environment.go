package storage

// Environment
// Another comment line to force Github Actions to execute

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ANDROID = "android" // node type android
const IOS = "ios"         // node type ios

type MesmerEnvironment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty`
	Name        string             `bson:"name,omitempty"`
	MesmerNodes []MesmerNode       `bson:"mesmernodes,omitempty`
}

func newEnvironment() *MesmerEnvironment {
	e := MesmerEnvironment{}
	return &e
}

func (env *MesmerEnvironment) setName(name string) error {
	env.Name = name
	return nil
}

func (env *MesmerEnvironment) addNode(ds *DataStore, node MesmerNode) error {
	ds.init()
	env.MesmerNodes = append(env.MesmerNodes, node)
	var newslice []MesmerNode
	newslice = append(newslice, node)
	error := ds.saveEnvironment(env)
	return error
}

func findEnvironmentByID(ds *DataStore, id primitive.ObjectID) (*MesmerEnvironment, error) {
	e, error := ds.findEnvironmentByID(id)
	return e, error
}

func findEnvironmentByName(ds *DataStore, name string) (*MesmerEnvironment, error) {
	e, error := ds.findEnvironmentByName(name)
	return e, error
}

func (env *MesmerEnvironment) nodes() []MesmerNode {
	return env.MesmerNodes
}

func (env *MesmerEnvironment) androidNodes(ds *DataStore) ([]MesmerNode, error) {
	result := []MesmerNode{}
	for i := range env.MesmerNodes {
		if env.MesmerNodes[i].NodeType == ANDROID {
			result = append(result, env.MesmerNodes[i])
		}
	}
	return result, nil
}

func (env *MesmerEnvironment) iOSNodes(ds *DataStore) ([]MesmerNode, error) {
	result := []MesmerNode{}
	for i := range env.MesmerNodes {
		if env.MesmerNodes[i].NodeType == ANDROID {
			result = append(result, env.MesmerNodes[i])
		}
	}
	return result, nil
}

func (env *MesmerEnvironment) save(ds *DataStore) error {
	err := ds.saveEnvironment(env)
	return err
}

func (env *MesmerEnvironment) update(ds *DataStore) error {
	err := ds.updateEnvironment(env)
	return err
}

func (env *MesmerEnvironment) toStrings() []string {
	var output []string
	for _, v := range env.MesmerNodes {
		output = append(output, fmt.Sprintf("%s %s %s", v.NodeType, v.Host, v.IP))
	}
	return output
}

func (env *MesmerEnvironment) removeNodeByHost(host string) {
	// TODO is this really the most efficient way? Review when more awake.
	var results []MesmerNode
	for _, v := range env.MesmerNodes {
		if v.Host != host {
			results = append(results, v)
		}
	}
	env.MesmerNodes = results
}
