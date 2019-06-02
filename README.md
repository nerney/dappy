<img src=logo.svg />

# Basic LDAP client for Go

[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/nerney/dappy)
[![Build Status](https://travis-ci.org/nerney/dappy.svg?branch=master)](https://travis-ci.org/nerney/dappy)
[![codecov](https://codecov.io/gh/nerney/dappy/branch/master/graph/badge.svg)](https://codecov.io/gh/nerney/dappy)
[![Report Card](https://goreportcard.com/badge/github.com/nerney/dappy)](https://goreportcard.com/report/github.com/nerney/dappy)

LDAP is complicated. Many times, all you really need to do is authenticate users with it.
This package boils down LDAP functionality to User Authentication, that's it.

Thanks to https://github.com/go-ldap/ldap

`go get github.com/nerney/dappy`

Example:

```go
package main

import (
	"log"

	"github.com/nerney/dappy"
)

func main() {
	var client dappy.Client
	var err error

	// create a new client
	if client, err = dappy.New(dappy.Config{
		BaseDN: "dc=example,dc=com",
		Filter: "uid",
		ROUser: dappy.User{Name: "cn=read-only-admin,dc=example,dc=com", Pass: "password"},
		Host:   "ldap.forumsys.com:389",
	}); err != nil {
		panic(err)
	}

	// username and password to authenticate
	username := "tesla"
	password := "password"

	// attempt the authentication
	if err := client.Auth(username, password); err != nil {
		panic(err)
	} else {
		log.Println("Success!")
	}
}

```
