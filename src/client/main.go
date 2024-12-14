package main

import (
	"go-sdap/src/client/managementClient"
	"go-sdap/src/client/sdapClient"
	"go-sdap/src/proto/management"
	"log/slog"
)

func main() {
	logger := slog.Default()

	logger.Info("Creating SDAP client...")

	s := sdapClient.New()
	status, err := s.Connect("127.0.0.1", 50051, false)
	if err == nil {
		logger.Info("Connect", "status", status.String())
	}

	user, status, err := s.Authenticate("borja", "1234")
	if err == nil {
		logger.Info("Authenticate", "status", status.String())
		logger.Info("Authenticate", "user", user.String())
	}

	s.Disconnect()

	logger.Info("Creating management client...")
	m := managementClient.New()

	mstatus, err := m.Connect("127.0.0.1", 50052, false)
	if err == nil {
		logger.Info("Connect", "status", mstatus.String())
	}

	us := "borjafm14"
	pass := "1234"
	cn := ""
	ln := "Fernández"
	en := "1"
	pn := "622111111"
	ad := "Avda 1"
	cr := "Engineer"
	t := "Backend"
	chars := make(map[string]string)
	chars["Role"] = "Employee"

	var mo []string
	mo = append(mo, "employees")

	var users []*management.User

	mu := management.User{
		Username:        &us,
		Password:        &pass,
		SdapRole:        management.SDAP_ROLE_ADMINISTRATOR.Enum(),
		CommonName:      &cn,
		FirstName:       &cn,
		LastName:        &ln,
		EmployeeNumber:  &en,
		PhoneNumber:     &pn,
		Address:         &ad,
		CompanyRole:     &cr,
		Team:            &t,
		Characteristics: chars,
		MemberOf:        mo,
	}

	users = append(users, &mu)

	mstatus, err = m.AddUsers(users)
	if err == nil {
		logger.Debug("AddUsers", "status", mstatus.String())
	} else {
		logger.Error("AddUsers", "error", err)
	}

	m.Disconnect()
}
