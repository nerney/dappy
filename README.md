<p align="left"><img src="logo/horizontal.png" alt="dappy" height="150px"></p>

## Basic LDAP client for Go

[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/nerney/dappy)
[![Build Status](https://travis-ci.org/nerney/dappy.svg?branch=master)](https://travis-ci.org/nerney/dappy)
[![codecov](https://codecov.io/gh/nerney/dappy/branch/master/graph/badge.svg)](https://codecov.io/gh/nerney/dappy)
[![Report Card](https://goreportcard.com/badge/github.com/nerney/dappy)](https://goreportcard.com/report/github.com/nerney/dappy)

LDAP is complicated. Many times, all you really need to do is authenticate users with it or fetch a user entry.
This package boils down LDAP functionality to User Authentication and Entry retrieval. 

Thanks to https://github.com/go-ldap/ldap

`go get github.com/nerney/dappy`

Example:

```go
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
```
