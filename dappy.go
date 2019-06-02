// Package dappy provides an ldap client for simple ldap authentication.
package dappy

import (
	"errors"
	"fmt"
	"net"
	"time"

	"gopkg.in/ldap.v3"
)

// Client interface performs ldap auth operation
type Client interface {
	Auth(username, password string) error
}

// Config to provide a dappy client.
// All fields are required, except for Filter.
type Config struct {
	BaseDN string // base directory, ex. "CN=Users,DC=Company"
	ROUser User   // the read-only user for initial bind
	Host   string // the ldap host and port, ex. "ldap.directory.com:389"
	Filter string // defaults to "sAMAccountName" for AD
}

// User holds the name and pass required for initial read-only bind.
type User struct {
	Name string
	Pass string
}

// local struct for implementing Client interface
type client struct {
	Config
}

// Auth implementation for the Client interface
func (c client) Auth(username, password string) error {
	// establish connection
	conn, err := connect(c.Host)
	if err != nil {
		return err
	}
	defer conn.Close()

	// perform initial read-only bind
	if err = conn.Bind(c.ROUser.Name, c.ROUser.Pass); err != nil {
		return err
	}

	// find the user attempting to login
	results, err := conn.Search(ldap.NewSearchRequest(
		c.BaseDN, ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false, fmt.Sprintf("(%v=%v)", c.Filter, username),
		[]string{}, nil,
	))
	if err != nil {
		return err
	}
	if len(results.Entries) < 1 {
		return errors.New("not found")
	}

	// attempt auth
	if err = conn.Bind(results.Entries[0].DN, password); err != nil {
		return err
	}
	return nil
}

// New dappy client with the provided config
// If the configuration provided is invalid,
// or dappy is unable to connect with the config
// provided, an error will be returned
func New(config Config) (Client, error) {
	config, err := validateConfig(config)
	if err != nil {
		return nil, err
	}
	c := client{config}
	conn, err := connect(c.Host) // test connection
	if err != nil {
		return nil, err
	}
	if err = conn.Bind(c.ROUser.Name, c.ROUser.Pass); err != nil {
		return nil, err
	}
	conn.Close()
	return c, err
}

// Helper functions

// establishes a connection with an ldap host
// (the caller is expected to Close the connection when finished)
func connect(host string) (*ldap.Conn, error) {
	c, err := net.DialTimeout("tcp", host, time.Second*8)
	if err != nil {
		return nil, err
	}
	conn := ldap.NewConn(c, false)
	conn.Start()
	return conn, nil
}

// validates that all required fields were provided
// handles default value for Filter
func validateConfig(config Config) (Config, error) {
	if config.BaseDN == "" || config.Host == "" || config.ROUser.Name == "" || config.ROUser.Pass == "" {
		return Config{}, errors.New("[CONFIG] The config provided could not be validated")
	}
	if config.Filter == "" {
		config.Filter = "sAMAccountName"
	}
	return config, nil
}
