package main

import (
	"github.com/nerney/dappy"
	"log"
)

func main() {

	//create a new client
	client := dappy.New(dappy.Options{
		BaseDN:       "CN=Users,DC=Company",
		Filter:       "sAMAccountName",
		BasePassword: "basePassword",
		BaseUser:     "baseUsername",
		URL:          "ldap.directory.com:389",
	})

	//username and password to authenticate
	username := "jdoe"
	password := "pass1234"

	//attempt the authentication
	err := client.Authenticate(username, password)

	//see the results
	if err != nil {
		log.Print(err)
	} else {
		log.Print("user successfully authenticated!")
	}

	//get a user entry
	user, err := client.GetUserEntry(username)
	if err == nil {
		user.PrettyPrint(2)
	}
}
