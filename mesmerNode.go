package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MesmerNode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Environment primitive.ObjectID `bson:"env,omitempty"`
	Host        string             `bson:"host,omitempty"`
	NodeType    string             `bson:"nodetype,omitempty"`
	IP          string             `bson:"ip,omitempty"`
}
