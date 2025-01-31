package helper

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mathrand "math/rand"
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	pbManagement "go-sdap/src/proto/management"
	pbSdap "go-sdap/src/proto/sdap"
)

const (
	passwordLength = 8
	lowercase      = "abcdefghijklmnopqrstuvwxyz"
	uppercase      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits         = "0123456789"
	specialChars   = "!@#$%^&*()-_=+[]{}|;:,.<>?/~`"
)

func ManagementCharacteristicToJSON(characteristic pbManagement.Characteristic) (string, error) {
	switch characteristic {
	case pbManagement.Characteristic_ADDRESS:
		return "address", nil
	case pbManagement.Characteristic_USERNAME:
		return "username", nil
	case pbManagement.Characteristic_COMMON_NAME:
		return "commonName", nil
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
		return "other_characteristics", nil
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

func SdapCharacteristicToJSON(characteristic pbSdap.Characteristic) (string, error) {
	switch characteristic {
	case pbSdap.Characteristic_ADDRESS:
		return "address", nil
	case pbSdap.Characteristic_COMMON_NAME:
		return "commonName", nil
	case pbSdap.Characteristic_COMPANY_ROLE:
		return "companyRole", nil
	case pbSdap.Characteristic_EMPLOYEE_NUMBER:
		return "employeeNumber", nil
	case pbSdap.Characteristic_FIRST_NAME:
		return "firstName", nil
	case pbSdap.Characteristic_LAST_NAME:
		return "lastName", nil
	case pbSdap.Characteristic_MEMBER_OF:
		return "memberOf", nil
	case pbSdap.Characteristic_OTHER:
		return "other_characteristics", nil
	case pbSdap.Characteristic_PHONE_NUMBER:
		return "phoneNumber", nil
	case pbSdap.Characteristic_REPORTS_TO:
		return "reportsTo", nil
	case pbSdap.Characteristic_TEAM:
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

func GeneratePassword() string {
	// Generate password, at least 8 characters long,
	// with at least one uppercase, one lowercase, one digit and one special character

	allChars := lowercase + uppercase + digits + specialChars
	password := make([]byte, passwordLength)

	// Ensure the password contains at least one character from each category
	password[0] = lowercase[randInt(len(lowercase))]
	password[1] = uppercase[randInt(len(uppercase))]
	password[2] = digits[randInt(len(digits))]
	password[3] = specialChars[randInt(len(specialChars))]

	// Fill the rest of the password with random characters from all categories
	for i := 4; i < passwordLength; i++ {
		password[i] = allChars[randInt(len(allChars))]
	}

	mathrand.Shuffle(passwordLength, func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

func ValidatePassword(password string) bool {
	// Validate password, at least 8 characters long,
	// with at least one uppercase, one lowercase, one digit and one special character

	if len(password) < passwordLength {
		return false
	}

	lowercaseFound := false
	uppercaseFound := false
	digitFound := false
	specialCharFound := false

	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			lowercaseFound = true
		} else if char >= 'A' && char <= 'Z' {
			uppercaseFound = true
		} else if char >= '0' && char <= '9' {
			digitFound = true
		} else {
			specialCharFound = true
		}
	}

	return lowercaseFound && uppercaseFound && digitFound && specialCharFound
}

func randInt(max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}

func GenerateToken() string {
	token := make([]byte, 32)
	for i := range token {
		token[i] = byte(randInt(256))
	}
	return fmt.Sprintf("%x", token)
}

func StringPtr(s string) *string {
	return &s
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
