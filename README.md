# godap

Basic LDAP Authenticator for Go

[![Build Status](https://travis-ci.org/nerney/godap.svg?branch=master)](https://travis-ci.org/nerney/godap)
[![Report Card](https://goreportcard.com/badge/github.com/nerney/godap)](https://goreportcard.com/report/github.com/nerney/godap)
[![codecov](https://codecov.io/gh/nerney/godap/branch/master/graph/badge.svg)](https://codecov.io/gh/nerney/godap)

LDAP is complicated. Many times, all you really need to do is authenticate users with it.
This package boils down LDAP functionality to one thing: User Authentication.

```
go get github.com/nerney/godap
```

Example:

```go
package main

import (
	"godap"
	"log"
)

func main() {

	//create a new client
	client := godap.New(godap.Options{
		BaseDN:   "CN=Users,DC=Company",
		Filter:   "sAMAccountName",
		Password: "username",
		Username: "password",
		URL:      "ldap.directory.com:389",
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
}
```
