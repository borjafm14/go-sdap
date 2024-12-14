package helper

import (
	"fmt"
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	pbManagement "go-sdap/src/proto/management"
)

func CharacteristicToJSON(characteristic pbManagement.Characteristic) (string, error) {
	switch characteristic {
	case pbManagement.Characteristic_ADDRESS:
		return "address", nil
	case pbManagement.Characteristic_COMMON_NAME:
		return "username", nil
	case pbManagement.Characteristic_COMPANY_ROLE:
		return "companyRole", nil
	case pbManagement.Characteristic_EMPLOYEE_NUMBER:
		return "employeeNumber", nil
	case pbManagement.Characteristic_FIRST_NAME:
		return "firstName", nil
	case pbManagement.Characteristic_LAST_NAME:
		return "lastName", nil
	case pbManagement.Characteristic_MEMBER_OF:
		return "memberOf", nil
	case pbManagement.Characteristic_OTHER:
		return "other", nil
	case pbManagement.Characteristic_PHONE_NUMBER:
		return "phoneNumber", nil
	case pbManagement.Characteristic_REPORTS_TO:
		return "reportsTo", nil
	case pbManagement.Characteristic_TEAM:
		return "team", nil
	default:
		return "", fmt.Errorf("unknown characteristic: %v", characteristic)
	}
}

func GenerateNextUsername(usernameBase, highestUsername string) string {
	// Extract the number in highestUsername
	re := regexp.MustCompile(fmt.Sprintf("^%s(\\d*)$", usernameBase))
	matches := re.FindStringSubmatch(highestUsername)

	if len(matches) > 1 && matches[1] != "" {
		// Increment the number
		highestNumber, _ := strconv.Atoi(matches[1])
		return fmt.Sprintf("%s%d", usernameBase, highestNumber+1)
	}

	// If there is no number in highestUsername, start from 1
	return fmt.Sprintf("%s1", usernameBase)
}

func SanitizeName(name string) string {
	validChars := []rune{}
	for _, char := range name {

		// Discard special chars
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			validChars = append(validChars, char)
		}
	}
	return string(validChars)
}

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
