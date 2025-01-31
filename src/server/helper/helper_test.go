package helper_test

import (
	"testing"

	pbManagement "go-sdap/src/proto/management"
	pbSdap "go-sdap/src/proto/sdap"
	"go-sdap/src/server/helper"

	"google.golang.org/protobuf/proto"
)

func TestManagementCharacteristicToJSON(t *testing.T) {
	tests := []struct {
		input    pbManagement.Characteristic
		expected string
	}{
		{pbManagement.Characteristic_ADDRESS, "address"},
		{pbManagement.Characteristic_USERNAME, "username"},
		{pbManagement.Characteristic_COMMON_NAME, "commonName"},
		{pbManagement.Characteristic_COMPANY_ROLE, "companyRole"},
		{pbManagement.Characteristic_EMPLOYEE_NUMBER, "employeeNumber"},
		{pbManagement.Characteristic_FIRST_NAME, "firstName"},
		{pbManagement.Characteristic_LAST_NAME, "lastName"},
		{pbManagement.Characteristic_MEMBER_OF, "memberOf"},
		{pbManagement.Characteristic_OTHER, "other_characteristics"},
		{pbManagement.Characteristic_PHONE_NUMBER, "phoneNumber"},
		{pbManagement.Characteristic_REPORTS_TO, "reportsTo"},
		{pbManagement.Characteristic_TEAM, "team"},
	}

	for _, test := range tests {
		result, err := helper.ManagementCharacteristicToJSON(test.input)
		if err != nil || result != test.expected {
			t.Errorf("expected %s, got %s, error: %v", test.expected, result, err)
		}
	}
}

func TestSdapCharacteristicToJSON(t *testing.T) {
	tests := []struct {
		input    pbSdap.Characteristic
		expected string
	}{
		{pbSdap.Characteristic_ADDRESS, "address"},
		{pbSdap.Characteristic_COMMON_NAME, "commonName"},
		{pbSdap.Characteristic_COMPANY_ROLE, "companyRole"},
		{pbSdap.Characteristic_EMPLOYEE_NUMBER, "employeeNumber"},
		{pbSdap.Characteristic_FIRST_NAME, "firstName"},
		{pbSdap.Characteristic_LAST_NAME, "lastName"},
		{pbSdap.Characteristic_MEMBER_OF, "memberOf"},
		{pbSdap.Characteristic_OTHER, "other_characteristics"},
		{pbSdap.Characteristic_PHONE_NUMBER, "phoneNumber"},
		{pbSdap.Characteristic_REPORTS_TO, "reportsTo"},
		{pbSdap.Characteristic_TEAM, "team"},
	}

	for _, test := range tests {
		result, err := helper.SdapCharacteristicToJSON(test.input)
		if err != nil || result != test.expected {
			t.Errorf("expected %s, got %s, error: %v", test.expected, result, err)
		}
	}
}

func TestGenerateNextUsername(t *testing.T) {
	tests := []struct {
		usernameBase    string
		highestUsername string
		expected        string
	}{
		{"user", "user1", "user2"},
		{"user", "user2", "user3"},
		{"user", "user", "user1"},
	}

	for _, test := range tests {
		result := helper.GenerateNextUsername(test.usernameBase, test.highestUsername)
		if result != test.expected {
			t.Errorf("expected %s, got %s", test.expected, result)
		}
	}
}

func TestSanitizeName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"John Doe", "JohnDoe"},
		{"Jane_Doe", "JaneDoe"},
		{"User123", "User"},
	}

	for _, test := range tests {
		result := helper.SanitizeName(test.input)
		if result != test.expected {
			t.Errorf("expected %s, got %s", test.expected, result)
		}
	}
}

func TestGeneratePassword(t *testing.T) {
	password := helper.GeneratePassword()
	if !helper.ValidatePassword(password) {
		t.Errorf("generated password is invalid: %s", password)
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Password1!", true},
		{"password", false},
		{"PASSWORD", false},
		{"12345678", false},
		{"Pass1!", false},
	}

	for _, test := range tests {
		result := helper.ValidatePassword(test.input)
		if result != test.expected {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestGenerateToken(t *testing.T) {
	token := helper.GenerateToken()
	if len(token) != 64 {
		t.Errorf("expected token length 64, got %d", len(token))
	}
}

func TestProtoToBSON(t *testing.T) {
	var pb proto.Message
	bsonDoc, err := helper.ProtoToBSON(pb)
	if err != nil {
		t.Errorf("error converting proto to BSON: %v", err)
	}
	if bsonDoc == nil {
		t.Errorf("expected non-nil BSON document")
	}
}
