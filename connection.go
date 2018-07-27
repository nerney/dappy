package dappy

import (
	"net"
	"time"

	"gopkg.in/ldap.v2"
)

// wrapper for ldap connection
type connection interface {
	Bind(username, password string) error
	Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error)
}

// establishes ldap connection
func connect(url string) connection {
	tcpConnection, err := net.DialTimeout("tcp", url, time.Second*10)
	if err != nil {
		panic(errCouldNotConnect)
	}
	ldapConnection := ldap.NewConn(tcpConnection, false)
	ldapConnection.Start()
	return ldapConnection
}
