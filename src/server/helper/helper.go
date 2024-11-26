package helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtoToBSON(pb proto.Message) (bson.M, error) {
	protoJSON, err := protojson.Marshal(pb)
	if err != nil {
		return nil, err
	}

	var bsonDoc bson.M
	err = bson.UnmarshalExtJSON(protoJSON, true, &bsonDoc)

	return bsonDoc, err
}

func BSONToProto(doc bson.M, pb proto.Message) error {
	bsonBytes, err := bson.Marshal(doc)
	if err != nil {
		return err
	}

	var jsonMap map[string]interface{}
	if err := bson.UnmarshalExtJSON(bsonBytes, true, &jsonMap); err != nil {
		return err
	}

	jsonBytes, err := bson.MarshalExtJSON(jsonMap, true, true)
	if err != nil {
		return err
	}

	if err := protojson.Unmarshal(jsonBytes, pb); err != nil {
		return err
	}

	return nil
}
