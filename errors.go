package dappy

import "errors"

// dappy errors
var (
	errCouldNotConnect = errors.New("godap: could not establish connection")
	errEmptyPassword   = errors.New("godap: empty password")
	errEmptyUsername   = errors.New("godap: empty username")
	errCouldNotBind    = errors.New("godap: could not perform initial bind")
	errSearch          = errors.New("godap: error performing search")
	errNotFound        = errors.New("godap: user not found in directory")
	errCouldNotAuth    = errors.New("godap: could not authenticate user")
)
